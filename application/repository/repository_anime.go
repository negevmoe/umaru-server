package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
	"umaru/application/enum"
	"umaru/application/model/dto"
)

func (r repositoryImpl) AnimeSelect(db IDB, req dto.AnimeSelectRequest) (res dto.AnimeSelectResponse, err error) {
	var values []any
	sql := `
select id,
       bangumi_id,
       category_id,
       title,
       season,
       cover,
       total,
       rss_url,
       rss_path,
       play_time,
       create_time,
       update_time
from anime where true
`
	if req.Id > 0 {
		sql += ` and id = ? `
		values = append(values, req.Id)
	}
	if req.BangumiId > 0 {
		sql += ` and bangumi_id = ? `
		values = append(values, req.BangumiId)
	}

	queryx, err := db.Queryx(sql, values...)
	if err != nil {
		log.Error("获取番剧详情失败", zap.Error(err), zap.Int64("id", req.Id), zap.Int64("bangumi_id", req.BangumiId))
		return
	}
	defer queryx.Close()

	if queryx.Next() {
		_ = queryx.StructScan(&res.Anime)
	}

	return
}

func (r repositoryImpl) AnimeSelectList(db IDB, req dto.AnimeSelectListRequest) (res dto.AnimeSelectListResponse, err error) {
	var values []any
	sql := `
select id,
       bangumi_id,
       category_id,
       title,
       season,
       cover,
       total,
       rss_url,
       rss_path,
       play_time,
       create_time,
       update_time
from anime where true 
`
	if req.CategoryId != 0 {
		sql += ` and category_id = ? `
		values = append(values, req.CategoryId)
	}

	err = db.Select(&res.Items, sql, values...)
	if err != nil {
		log.Error("获取番剧列表失败", zap.Error(err),
			zap.Int64("category_id", req.CategoryId),
		)
		return
	}

	return
}

func (r repositoryImpl) AnimeInfoViewSelect(db IDB, req dto.AnimeInfoViewSelectRequest) (res dto.AnimeInfoViewSelectResponse, err error) {
	var values []any
	sql := `
select id,
       bangumi_id,
       category_id,
       category_name,
       title,
       season,
       cover,
       total,
       rss_url,
       rss_path,
       play_time,
       create_time,
       update_time
from anime_info_view where true 
`

	if req.Id > 0 {
		sql += ` and id = ? `
		values = append(values, req.Id)
	}
	if req.BangumiId > 0 {
		sql += ` and bangumi_id = ? `
		values = append(values, req.BangumiId)
	}

	queryx, err := db.Queryx(sql, values...)
	if err != nil {
		log.Error("获取番剧详情失败", zap.Error(err), zap.Int64("id", req.Id), zap.Int64("bangumi_id", req.BangumiId))
		return
	}
	defer queryx.Close()

	if queryx.Next() {
		_ = queryx.StructScan(&res.AnimeInfo)
	}
	return
}

func (r repositoryImpl) AnimeInfoViewSelectList(db IDB, req dto.AnimeInfoViewSelectListRequest) (res dto.AnimeInfoViewSelectListResponse, err error) {
	l := log.With(
		zap.String("title", req.Title),
		zap.Int64("category_id", req.CategoryId),
		zap.Time("play_start_time", time.Unix(req.PlayStartTime, 0)),
		zap.Time("play_end_time", time.Unix(req.PlayEndTime, 0)),
		zap.Time("add_start_time", time.Unix(req.AddStartTime, 0)),
		zap.Time("add_end_time", time.Unix(req.AddEndTime, 0)),
	)

	var values []any
	sql := `
select id,
       bangumi_id,
       category_id,
       category_name,
       title,
       season,
       cover,
       total,
       rss_url,
       rss_path,
       play_time,
       create_time,
       update_time
from anime_info_view where true 
`
	if req.CategoryId != 0 {
		sql += ` and category_id = ? `
		values = append(values, req.CategoryId)
	}

	if req.PlayStartTime > 0 {
		sql += ` and play_time >= ? `
		values = append(values, req.PlayStartTime)
	}

	if req.PlayEndTime > 0 {
		sql += ` and play_time <= ? `
		values = append(values, req.PlayEndTime)
	}
	if req.AddStartTime > 0 {
		sql += ` and create_time >= ? `
		values = append(values, req.AddStartTime)
	}

	if req.AddEndTime > 0 {
		sql += ` and create_time <= ? `
		values = append(values, req.AddEndTime)
	}

	if req.Title != "" {
		sql += ` and title like '%'||?||'%' `
		values = append(values, req.Title)
	}

	err = db.Get(&res.Total, r.count(sql), values...)
	if err != nil {
		l.Error("获取番剧列表失败", zap.Error(err))
		return
	}

	switch req.Sort {
	case enum.CREATE_TIME_DESC:
		sql += ` order by create_time desc,id desc`
	case enum.CREATE_TIME_ASC:
		sql += ` order by create_time asc,id desc `
	case enum.PLAY_TIME_DESC:
		sql += ` order by play_time desc,id desc `
	case enum.PLAY_TIME_ASC:
		sql += ` order by play_time asc,id desc `
	}

	if req.Limit > 0 {
		sql += ` limit ? offset ? `
		values = append(values, req.Limit, req.Offset)
	}

	err = db.Select(&res.Items, sql, values...)
	if err != nil {
		l.Error("获取番剧列表失败", zap.Error(err))
		return
	}

	return
}

func (r repositoryImpl) AnimeInsert(db IDB, req dto.AnimeInsertRequest) (res dto.AnimeInsertResponse, err error) {

	sql := `
insert into anime (
   bangumi_id, 
   category_id, 
   title, 
   season, 
   cover, 
   total, 
   rss_url,
   rss_path,
   play_time, 
   create_time, 
   update_time
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)
returning id
`

	err = db.Get(&res.Id, sql,
		req.Anime.BangumiId,
		req.Anime.CategoryId,
		req.Anime.Title,
		req.Anime.Season,
		req.Anime.Cover,
		req.Anime.Total,
		req.Anime.RssUrl,
		req.Anime.RssPath,
		req.Anime.PlayTime,
		req.Anime.CreateTime,
		req.Anime.UpdateTime,
	)
	if err != nil {
		log.Error("添加番剧失败", zap.Error(err),
			zap.Int64("bangumi_id", req.Anime.BangumiId),
			zap.Int64("category_id", req.Anime.CategoryId),
			zap.String("title", req.Anime.Title),
			zap.Int64("season", req.Anime.Season),
			zap.String("cover", req.Anime.Cover),
			zap.Int64("total", req.Anime.Total),
			zap.String("rss_url", req.Anime.RssUrl),
			zap.String("rss_path", req.Anime.RssPath),
			zap.Time("play_time", time.Unix(req.Anime.PlayTime, 0)),
			zap.Time("create_time", time.Unix(req.Anime.CreateTime, 0)),
			zap.Time("update_time", time.Unix(req.Anime.UpdateTime, 0)),
		)
		return
	}
	return
}

func (r repositoryImpl) AnimeDelete(db IDB, req dto.AnimeDeleteRequest) (res dto.AnimeDeleteResponse, err error) {
	sql := ` delete from anime where id = ? `
	_, err = db.Exec(sql, req.Id)
	if err != nil {
		log.Error("删除番剧失败", zap.Error(err), zap.Int64("id", req.Id))
		return
	}
	return
}

func (r repositoryImpl) AnimeUpdate(db IDB, req dto.AnimeUpdateRequest) (res dto.AnimeUpdateResponse, err error) {
	sql := `
update anime set
category_id = ?,
title = ?,
season = ?,
total = ?,
play_time = ?,
update_time = ?
where id = ?;
`
	_, err = db.Exec(sql,
		req.Anime.CategoryId,
		req.Anime.Title,
		req.Anime.Season,
		req.Anime.Total,
		req.Anime.PlayTime,
		req.Anime.UpdateTime,
		req.Anime.Id,
	)
	if err != nil {
		log.Error("更新番剧失败", zap.Error(err),
			zap.Int64("id", req.Anime.Id),
			zap.Int64("category_id", req.Anime.CategoryId),
			zap.String("title", req.Anime.Title),
			zap.Int64("season", req.Anime.Season),
			zap.Int64("total", req.Anime.Total),
			zap.Time("play_time", time.Unix(req.Anime.PlayTime, 0)),
			zap.Time("update_time", time.Unix(req.Anime.UpdateTime, 0)),
		)
		return
	}

	return
}

func (r repositoryImpl) AnimeSelectByTitleAndSeason(db IDB, req dto.AnimeSelectByTitleAndSeasonRequest) (res dto.AnimeSelectByTitleAndSeasonResponse, err error) {
	sql := `
select id,
       bangumi_id,
       category_id,
       title,
       season,
       cover,
       total,
       rss_url,
       rss_path,
       play_time,
       create_time,
       update_time
from anime where title = ? and season = ?
`

	queryx, err := db.Queryx(sql, req.Title, req.Season)
	if err != nil {
		log.Error("获取番剧信息失败", zap.Error(err),
			zap.String("title", req.Title),
			zap.Int64("season", req.Season),
		)
		return
	}
	defer queryx.Close()

	if queryx.Next() {
		_ = queryx.StructScan(&res.Anime)
	}
	return
}

func (r repositoryImpl) AnimeCategoryUpdate(db IDB, req dto.AnimeCategoryUpdateRequest) (res dto.AnimeCategoryUpdateResponse, err error) {
	sql := `
update anime
set category_id = ?
where id in (?)
`
	sql, values, err := sqlx.In(sql, req.CategoryId, req.IdList)
	if err != nil {
		log.Error("番剧更新分类ID失败", zap.Error(err), zap.Int64s("id_list", req.IdList), zap.Int64("category_id", req.CategoryId))
		return
	}
	_, err = db.Exec(sql, values...)
	if err != nil {
		log.Error("番剧更新分类ID失败", zap.Error(err), zap.Int64s("id_list", req.IdList), zap.Int64("category_id", req.CategoryId))
		return
	}
	return
}

func (r repositoryImpl) AnimeRssUrlUpdate(db IDB, req dto.AnimeRssUrlUpdateRequest) (res dto.AnimeRssUrlUpdateResponse, err error) {
	sql := `update anime set rss_url = ?, rss_path = ? where id = ?`
	_, err = db.Exec(sql, req.RssUrl, req.RssPath, req.Id)
	if err != nil {
		log.Error("番剧RSS更新失败", zap.Error(err),
			zap.Int64("id", req.Id),
			zap.String("rss_url", req.RssUrl),
			zap.String("rss_path", req.RssPath),
		)
		return
	}
	return
}
