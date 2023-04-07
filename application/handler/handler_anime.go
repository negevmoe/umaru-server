package handler

import (
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
	"umaru-server/application/model/dao"
	"umaru-server/application/model/dto"
	"umaru-server/application/model/vo"
	"umaru-server/application/setting"
	"umaru-server/application/usecase"
)

/*

根据种子hash操作不行,RSS无法实时获取哈希, 手动添加提取麻烦


1. 创建接口 (种子列表和RSS连接必填一个), 种子列表qb提交, 创建番剧
2. bangumi番剧详情 调用番剧详情接口与 获取番剧视频列表接口 可知RSS订阅情况与本地视频情况
3. 如果RSS有订阅, 必须先取消当前订阅(选项:是否删除qbittorrent下载的文件)(qbittorrent删除,rss清空)
4. 如果RSS无订阅, 可以订阅
5. 手动添加种子
6. 下面有视频列表,可以进行删除
*/

/*
番剧管理
title season  total bangumi_url rule_name category_name play_time create_time update_time  operation: 添加种子 更新 删除
详: rss_url video_list

bangumi_id 唯一
更新: 如果修改了title season category_id 其中任何一个 都要移动原视频
*/

func (s handlerImpl) AnimeCreate(req vo.AnimeCreateRequest) (res vo.AnimeCreateResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}
	// 检查分类是否存在
	category, err := repo.CategorySelect(db, dto.CategorySelectRequest{
		Id: req.CategoryId,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "添加失败", err)
		return
	}
	if category.Category.Id == 0 {
		err = vo.ErrorNew(400, "分类不存在", "")
		return
	}

	// 如果season为0 尝试提取季
	if req.Season == 0 {
		season, success := usecase.GetSeason(req.Title)
		if success {
			req.Season = season
		} else {
			log.Warn("番剧季信息提取失败,已设置为第一季", zap.String("title", req.Title))
			req.Season = 1
		}
	}

	// 检查 title+season是否已存在
	tsExists, err := repo.AnimeSelectByTitleAndSeason(db, dto.AnimeSelectByTitleAndSeasonRequest{
		Title:  req.Title,
		Season: req.Season,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "添加失败", err)
		return
	}
	if tsExists.Anime.Id > 0 {
		err = vo.ErrorNew(400, "番剧已存在", "")
		return
	}

	// 如果有bangumi id 检查是否已订阅
	bExists, err := repo.AnimeSelect(db, dto.AnimeSelectRequest{
		BangumiId: req.BangumiId,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "添加失败", err)
		return
	}
	if bExists.Anime.Id > 0 {
		err = vo.ErrorNew(400, "番剧已订阅", "")
		return
	}

	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		err = vo.ErrorNew(500, "添加失败", "")
		return
	}
	defer tx.Rollback()
	// 入库
	now := time.Now().Unix()
	rssPath := usecase.GetRssPath(req.Title)
	ret, err := repo.AnimeInsert(tx, dto.AnimeInsertRequest{
		Anime: dao.Anime{
			BangumiId:  req.BangumiId,
			CategoryId: req.CategoryId,
			Title:      req.Title,
			Season:     req.Season,
			Cover:      req.Cover,
			Total:      req.Total,
			RssUrl:     req.RssUrl,
			RssPath:    rssPath,
			PlayTime:   req.PlayTime,
			CreateTime: now,
			UpdateTime: now,
		},
	})
	if err != nil {
		err = vo.ErrorWrap(500, "添加失败", err)
		return
	}
	res.Id = ret.Id

	// 手动添加的种子
	if len(req.TorrentList) != 0 {
		_, err = repo.QBTorrentInsertList(dto.QBTorrentInsertListRequest{
			Urls:     strings.Join(req.TorrentList, "\n"),
			Category: setting.QB_CATEGORY,
			SavePath: usecase.GetQbDownloadPath(ret.Id),
		})
		if err != nil {
			err = vo.ErrorWrap(500, "添加失败", err)
			return
		}
	}

	// RSS订阅
	if req.RssUrl != "" {
		_, err = repo.QBRssInsert(dto.QBRssInsertRequest{
			Url:  req.RssUrl,
			Path: rssPath,
		})
		if err != nil {
			err = vo.ErrorWrap(500, "添加失败", err)
			return
		}

		_, err = repo.QBRuleSet(dto.QBRuleSetRequest{
			RuleName: rssPath,
			RuleDef: dao.RuleDef{
				Enabled:          true,
				MustContain:      req.MustContain,
				MustNotContain:   req.MustNotContain,
				UseRegex:         req.UseRegex,
				EpisodeFilter:    req.EpisodeFilter,
				SmartFilter:      req.SmartFilter,
				AffectedFeeds:    []string{req.RssUrl},
				AddPaused:        false,
				AssignedCategory: setting.QB_CATEGORY,
				SavePath:         usecase.GetQbDownloadPath(ret.Id),
			},
		})
		if err != nil {
			err = vo.ErrorWrap(500, "添加失败", err)
			return
		}
	}

	// 提交
	err = tx.Commit()
	if err != nil {
		err = vo.ErrorWrap(500, "添加失败", err)
		return
	}
	return
}

// AnimeGet 获取番剧详情
func (s handlerImpl) AnimeGet(req vo.AnimeGetRequest) (res vo.AnimeGetResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	animeRet, err := repo.AnimeInfoViewSelect(db, dto.AnimeInfoViewSelectRequest{
		Id:        req.Id,
		BangumiId: req.BangumiId,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "获取番剧失败", err)
		return
	}
	res.Anime = animeRet.AnimeInfo
	return
}

// AnimeGetList 获取番剧列表
func (s handlerImpl) AnimeGetList(req vo.AnimeGetListRequest) (res vo.AnimeGetListResponse, err error) {
	ret, err := repo.AnimeInfoViewSelectList(db, dto.AnimeInfoViewSelectListRequest{
		Limit:         req.Size,
		Offset:        req.Size * (req.Page - 1),
		Title:         req.Title,
		CategoryId:    req.CategoryId,
		Sort:          req.Sort,
		PlayStartTime: req.PlayStartTime,
		PlayEndTime:   req.PlayEndTime,
		AddStartTime:  req.AddStartTime,
		AddEndTime:    req.AddEndTime,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "获取番剧列表失败", err)
		return
	}

	res.Total = ret.Total
	res.Items = ret.Items
	if res.Items == nil {
		res.Items = make([]dao.AnimeInfoView, 0)
	}
	return
}

// AnimeDelete 删除番剧
//	1. 获取番剧信息
//	2. 找到视频目录
//	3. 开启事务-> 删除数据 -> 删除视频
//	注: 删除数据,硬链接的视频,QB中的订阅信息. 但没有删除订阅下载的源视频
//	执行后 硬链接不会再次执行,种子还可以继续保种
//	如果qb下载的也想删除,自行到qbittorrent中进行删除
func (s handlerImpl) AnimeDelete(req vo.AnimeDeleteRequest) (res vo.AnimeDeleteResponse, err error) {
	// 硬连接锁
	LinkLock.Lock()
	defer LinkLock.Unlock()

	// 获取番剧信息
	ret, err := repo.AnimeInfoViewSelect(db, dto.AnimeInfoViewSelectRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "删除番剧失败", err)
		return
	}

	anime := ret.AnimeInfo
	if anime.Id == 0 {
		log.Error("番剧不存在", zap.Int64("id", req.Id))
		err = vo.ErrorNew(500, "删除番剧失败", "番剧不存在")
		return
	}

	// 获取硬链接目录
	dir, err := usecase.GetLinkDir(anime.CategoryName, anime.Title, anime.Season)
	if err != nil {
		log.Error("删除番剧失败", zap.Error(err),
			zap.String("category_name", anime.CategoryName),
			zap.String("title", anime.Title),
			zap.Int64("season", anime.Season),
		)
		err = vo.ErrorWrap(500, "删除番剧失败", err)
		return
	}

	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		log.Error("删除番剧失败,开启事务失败", zap.Error(err))
		err = vo.ErrorWrap(500, "删除番剧失败", err)
		return
	}
	defer tx.Rollback()

	// 删除数据
	_, err = repo.AnimeDelete(tx, dto.AnimeDeleteRequest{Id: req.Id})
	if err != nil {
		err = vo.ErrorWrap(500, "删除番剧失败", err)
		return
	}

	// 删除硬连接视频目录
	err = os.RemoveAll(dir)
	if err != nil {
		log.Error("删除番剧视频失败", zap.Error(err),
			zap.String("dir", dir),
		)
		return
	}

	// 删除RSS
	if anime.RssPath != "" {
		_, err = repo.QBRssDelete(dto.QBRssDeleteRequest{
			Path: anime.RssPath,
		})
		if err != nil {
			if !strings.Contains(err.Error(), "项目不存在") {
				err = vo.ErrorWrap(500, "删除失败", err)
				return
			}
		}
	}

	// 删除规则
	_, err = repo.QBRuleDelete(dto.QBRuleDeleteRequest{
		RuleName: anime.RssPath,
	})
	if err != nil {
		log.Error("删除下载规则失败", zap.Error(err), zap.String("rule_name", anime.Title))
		return
	}

	err = tx.Commit()
	if err != nil {
		err = vo.ErrorWrap(500, "删除番剧失败", err)
		return
	}
	return
}

//可以更新 标题,季,分类,总集数,分类,放送时间
//其中
//只更新标题时,创建 新番剧名称/S季 的文件夹, 原番剧名称/S季 目录移动到 新番剧名称/S季 ; 更新标题时要判断 新标题+原季是否存在
//`setting.MEDIA_PATH/分类/原番剧名称/S季/重命名文件` ==> `setting.MEDIA_PATH/分类/新番剧名称/S季/重命名文件`
//只更新季时, 新建季文件夹, 原季文件夹 移动到 新季文件夹; 更新季时要判断 标题+新季是否存在
//`setting.MEDIA_PATH/分类/番剧名称/原季/重命名文件` ==> `setting.MEDIA_PATH/分类/番剧名称/新季/重命名文件`
//
//标题与季更新时, 创建 新番剧名称/新季 的文件夹, 原番剧名称/原季 目录移动到

// AnimeUpdate 修改番剧信息
//	如果修改的番剧有bangumi_id值,则不允许修改标题
func (s handlerImpl) AnimeUpdate(req vo.AnimeUpdateRequest) (res vo.AnimeUpdateResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	// 如果season为0 尝试提取季
	if req.Season == 0 {
		season, success := usecase.GetSeason(req.Title)
		if success {
			req.Season = season
		} else {
			log.Warn("番剧季信息提取失败,已设置为第一季", zap.String("title", req.Title))
			req.Season = 1
		}
	}

	// 检查番剧是否存在
	oldRet, err := repo.AnimeInfoViewSelect(db, dto.AnimeInfoViewSelectRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "番剧更新失败", err)
		return
	}
	old := oldRet.AnimeInfo
	if old.Id == 0 {
		err = vo.ErrorNew(400, "番剧不存在", "")
		return
	}
	// 检查 名称+季 是否存在
	existsRet, err := repo.AnimeSelectByTitleAndSeason(db,
		dto.AnimeSelectByTitleAndSeasonRequest{
			Title:  req.Title,
			Season: req.Season,
		},
	)
	if err != nil {
		err = vo.ErrorWrap(500, "番剧更新失败", err)
		return
	}
	// 如果存在且存在的番剧的ID不是请求的ID
	if existsRet.Anime.Id > 0 && existsRet.Anime.Id != req.Id {
		err = vo.ErrorNew(400, fmt.Sprintf("番剧 %s %d 已存在", req.Title, req.Season), "")
		return
	}

	// 检查分类是否存在
	categoryRet, err := repo.CategorySelect(db, dto.CategorySelectRequest{
		Id: req.CategoryId,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "番剧更新失败", err)
		return
	}
	category := categoryRet.Category
	if category.Id == 0 {
		err = vo.ErrorNew(400, "番剧更新失败", "分类不存在")
		return
	}

	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		err = vo.ErrorWrap(500, "番剧更新失败", err)
		return
	}
	defer tx.Rollback()

	// 更新数据
	updateAnime := dao.Anime{
		Id:         req.Id,
		CategoryId: req.CategoryId,
		Title:      req.Title,
		Season:     req.Season,
		Total:      req.Total,
		PlayTime:   req.PlayTime,
		UpdateTime: time.Now().Unix(),
	}
	_, err = repo.AnimeUpdate(tx, dto.AnimeUpdateRequest{Anime: updateAnime})
	if err != nil {
		err = vo.ErrorWrap(500, "番剧更新失败", err)
		return
	}

	// 如果有改动分类/名称/季 任何一个 删除原硬连接目录 执行硬连接
	if req.Title != old.Title || req.Season != old.Season && req.CategoryId != req.CategoryId {
		LinkLock.Lock()
		oldDir, er := usecase.GetLinkDir(old.CategoryName, old.Title, old.Season)
		if er != nil {
			LinkLock.Unlock()
			err = vo.ErrorWrap(500, "番剧更新失败", er)
			return
		}

		err = os.RemoveAll(oldDir)
		if err != nil {
			LinkLock.Unlock()
			err = vo.ErrorWrap(500, "番剧更新失败", err)
			return
		}
		LinkLock.Unlock()
		go Link()
	}
	_ = tx.Commit()
	return
}

func (s handlerImpl) AnimeVideoGetList(req vo.AnimeVideoGetListRequest) (res vo.AnimeVideoGetListResponse, err error) {
	animeRet, err := repo.AnimeInfoViewSelect(db, dto.AnimeInfoViewSelectRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "获取视频失败", err)
		return
	}
	anime := animeRet.AnimeInfo
	if anime.Id == 0 {
		err = vo.ErrorNew(400, "番剧不存在", "")
		return
	}

	dir, err := usecase.GetLinkDir(anime.CategoryName, anime.Title, anime.Season)
	if err != nil {
		log.Error("获取硬连接目录失败", zap.Error(err),
			zap.String("category_name", anime.CategoryName),
			zap.String("title", anime.Title),
			zap.Int64("season", anime.Season),
		)
		err = vo.ErrorWrap(500, "获取视频失败", err)
		return
	}

	var list []dao.Video

	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			log.Error("获取番剧的视频信息失败", zap.Error(err),
				zap.Int64("anime_id", anime.Id),
				zap.String("path", path),
			)
			return err
		}

		list = append(list, dao.Video{
			AnimeId:    anime.Id,
			Path:       path,
			Filename:   info.Name(),
			Size:       info.Size(),
			UpdateTime: info.ModTime().Unix(),
		})
		return nil
	})

	if err != nil {
		err = vo.ErrorWrap(500, "获取视频列表失败", err)
		return
	}

	res.Items = list
	return
}

func (s handlerImpl) AnimeRssCancel(req vo.AnimeRssCancelRequest) (res vo.AnimeRssCancelResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	// 获取番剧信息
	animeSelect, err := repo.AnimeSelect(db, dto.AnimeSelectRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "取消订阅失败", err)
		return
	}
	anime := animeSelect.Anime
	if anime.Id == 0 {
		err = vo.ErrorNew(500, "取消订阅失败", "番剧不存在")
		return
	}

	if anime.RssUrl == "" {
		return
	}

	tx, err := db.Beginx()
	if err != nil {
		err = vo.ErrorWrap(500, "取消订阅失败", err)
		return
	}
	defer tx.Rollback()

	// 更新数据库
	_, err = repo.AnimeRssUrlUpdate(tx, dto.AnimeRssUrlUpdateRequest{
		Id:      req.Id,
		RssUrl:  "",
		RssPath: "",
	})
	if err != nil {
		err = vo.ErrorWrap(500, "取消订阅失败", err)
		return
	}

	// QB 删除规则
	_, err = repo.QBRuleDelete(dto.QBRuleDeleteRequest{
		RuleName: anime.RssPath,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "取消订阅失败", err)
		return
	}
	// QB 删除rss
	_, err = repo.QBRssDelete(dto.QBRssDeleteRequest{
		Path: anime.RssPath,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "取消订阅失败", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = vo.ErrorWrap(500, "取消订阅失败", err)
		return
	}

	return
}

func (s handlerImpl) AnimeRssAdd(req vo.AnimeRssAddRequest) (res vo.AnimeRssAddResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	// 获取番剧信息
	animeSelect, err := repo.AnimeSelect(db, dto.AnimeSelectRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "订阅失败", err)
		return
	}
	anime := animeSelect.Anime
	if anime.Id == 0 {
		err = vo.ErrorNew(500, "订阅失败", "番剧不存在")
		return
	}

	if anime.RssUrl != "" {
		err = vo.ErrorNew(500, "订阅失败", "番剧已订阅rss")
		return
	}

	tx, err := db.Beginx()
	if err != nil {
		err = vo.ErrorWrap(500, "订阅失败", err)
		return
	}
	defer tx.Rollback()

	rssPath := usecase.GetRssPath(anime.Title)
	// 更新数据库
	_, err = repo.AnimeRssUrlUpdate(tx, dto.AnimeRssUrlUpdateRequest{
		Id:      req.Id,
		RssUrl:  req.RssUrl,
		RssPath: rssPath,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "订阅失败", err)
		return
	}

	// QB 添加规则
	_, err = repo.QBRuleSet(dto.QBRuleSetRequest{
		RuleName: rssPath,
		RuleDef: dao.RuleDef{
			Enabled:          true,
			MustContain:      req.MustContain,
			MustNotContain:   req.MustNotContain,
			UseRegex:         req.UseRegex,
			EpisodeFilter:    req.EpisodeFilter,
			SmartFilter:      req.SmartFilter,
			AffectedFeeds:    []string{req.RssUrl},
			AddPaused:        false,
			AssignedCategory: setting.QB_CATEGORY,
			SavePath:         usecase.GetQbDownloadPath(anime.Id),
		},
	})
	if err != nil {
		err = vo.ErrorWrap(500, "订阅失败", err)
		return
	}
	// QB 添加rss
	_, err = repo.QBRssInsert(dto.QBRssInsertRequest{
		Url:  req.RssUrl,
		Path: rssPath,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "订阅失败", err)
		return
	}

	err = tx.Commit()
	if err != nil {
		err = vo.ErrorWrap(500, "订阅失败", err)
		return
	}

	return
}
