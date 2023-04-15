package repository

import (
	"errors"
	"github.com/imroc/req/v3"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"syscall"
	"time"
	"umaru-server/application/global"
	"umaru-server/application/model/dao"
	"umaru-server/application/model/dto"
	"umaru-server/application/setting"
)

var Repo IRepository = repositoryImpl{}
var log *zap.Logger
var qb *req.Client

func Init() {
	log = global.Log
	qb = global.QB
	Repo.initSqlite(global.Sqlite)
	Repo.initQbittorrent(global.Sqlite)
}

type repositoryImpl struct{}

type IDB interface {
	sqlx.Queryer
	sqlx.Execer
	sqlx.Preparer
	sqlx.QueryerContext
	sqlx.ExecerContext
	sqlx.PreparerContext

	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
}

type IRepository interface {
	initSqlite(db *sqlx.DB)
	initQbittorrent(db *sqlx.DB)
	count(sql string) string // 构建count语句

	AnimeSelect(db IDB, req dto.AnimeSelectRequest) (res dto.AnimeSelectResponse, err error)                                                 // 番剧 获取详情
	AnimeSelectByTitleAndSeason(db IDB, req dto.AnimeSelectByTitleAndSeasonRequest) (res dto.AnimeSelectByTitleAndSeasonResponse, err error) // 番剧 根据标题和季查询
	AnimeSelectList(db IDB, req dto.AnimeSelectListRequest) (res dto.AnimeSelectListResponse, err error)                                     // 番剧 获取列表
	AnimeInfoViewSelect(db IDB, req dto.AnimeInfoViewSelectRequest) (res dto.AnimeInfoViewSelectResponse, err error)                         // 番剧 获取详细信息
	AnimeInfoViewSelectList(db IDB, req dto.AnimeInfoViewSelectListRequest) (res dto.AnimeInfoViewSelectListResponse, err error)             // 番剧 获取详细信息列表
	AnimeInsert(db IDB, req dto.AnimeInsertRequest) (res dto.AnimeInsertResponse, err error)                                                 // 番剧 插入
	AnimeDelete(db IDB, req dto.AnimeDeleteRequest) (res dto.AnimeDeleteResponse, err error)                                                 // 番剧 删除
	AnimeUpdate(db IDB, req dto.AnimeUpdateRequest) (res dto.AnimeUpdateResponse, err error)                                                 // 番剧 更新
	AnimeRssUrlUpdate(db IDB, req dto.AnimeRssUrlUpdateRequest) (res dto.AnimeRssUrlUpdateResponse, err error)                               // 番剧 更新RSS
	AnimeCategoryUpdate(db IDB, req dto.AnimeCategoryUpdateRequest) (res dto.AnimeCategoryUpdateResponse, err error)                         // 番剧 批量更新分类ID                                              // 番剧 更新

	RuleSelectList(db IDB, req dto.RuleSelectListRequest) (res dto.RuleSelectListResponse, err error)       // 下载规则 获取列表
	RuleSelectByName(db IDB, req dto.RuleSelectByNameRequest) (res dto.RuleSelectByNameResponse, err error) // 下载规则 根据名称查询
	RuleUpdate(db IDB, req dto.RuleUpdateRequest) (res dto.RuleUpdateResponse, err error)                   // 下载规则 更新
	RuleDeleteList(db IDB, req dto.RuleDeleteListRequest) (res dto.RuleDeleteListResponse, err error)       // 下载规则 批量删除
	RuleInsert(db IDB, req dto.RuleInsertRequest) (res dto.RuleInsertResponse, err error)                   // 下载规则 插入

	CategorySelect(db IDB, req dto.CategorySelectRequest) (res dto.CategorySelectResponse, err error)             // 分类 获取详情
	CategoryInsert(db IDB, req dto.CategoryInsertRequest) (res dto.CategoryInsertResponse, err error)             // 分类 添加
	CategorySelectList(db IDB, req dto.CategorySelectListRequest) (res dto.CategorySelectListResponse, err error) // 分类 获取列表
	CategoryDelete(db IDB, req dto.CategoryDeleteRequest) (res dto.CategoryDeleteResponse, err error)             // 分类 删除
	CategoryUpdate(db IDB, req dto.CategoryUpdateRequest) (res dto.CategoryUpdateResponse, err error)             // 分类 更新

	QBLogin(req dto.QBLoginRequest) (res dto.QBLoginResponse, err error)                                           // qbittorrent 登录
	QBLogSelectList(req dto.QBLogSelectListRequest) (res dto.QBLogSelectListResponse, err error)                   // qbittorrent 获取日志
	QBCategoryInsert(req dto.QBCategoryInsertRequest) (res dto.QBCategoryInsertResponse, err error)                // qbittorrent 创建分类
	QBCategorySelectList(req dto.QBCategorySelectListRequest) (res dto.QBCategorySelectListResponse, err error)    // qbittorrent 获取分类列表
	QBRuleSet(req dto.QBRuleSetRequest) (res dto.QBRuleSetResponse, err error)                                     // qbittorrent 添加/更新下载规则
	QBRuleDelete(req dto.QBRuleDeleteRequest) (res dto.QBRuleDeleteResponse, err error)                            // qbittorrent 删除下载规则
	QBRssInsert(req dto.QBRssInsertRequest) (res dto.QBRssInsertResponse, err error)                               // qbittorrent 添加RSS
	QBRssDelete(req dto.QBRssDeleteRequest) (res dto.QBRssDeleteResponse, err error)                               // qbittorrent 删除RSS
	QBRssFolderSelectList(req dto.QBRssFolderSelectListRequest) (res dto.QBRssFolderSelectListResponse, err error) // qbittorrent 获取RSS目录
	QBRssFolderInsert(req dto.QBRssFolderInsertRequest) (res dto.QBRssFolderInsertResponse, err error)             // qbittorrent 创建RSS目录
	QBTorrentInsertList(req dto.QBTorrentInsertListRequest) (res dto.QBTorrentInsertListResponse, err error)       // qbittorrent 批量添加种子

	MikanInfoSelect(req dto.MikanInfoSelectRequest) (res dto.MikanInfoSelectResponse, err error)          // mikan 获取mikan信息
	MikanRssSelectList(req dto.MikanRssSelectListRequest) (res dto.MikanRssSelectListResponse, err error) // mikan 获取rss
}

func reqHandler(res *req.Response, err error) error {
	if err != nil {
		return err
	}
	if res.GetStatusCode() != 200 {
		return errors.New(res.String())
	}
	return nil
}

func (r repositoryImpl) initSqlite(db *sqlx.DB) {
	var err error
	log.Info("初始化数据库", zap.String("db_path", setting.DB_PATH))

	errMsg := "初始化数据库失败"

	// 创建表
	if _, err = db.Exec(global.SQL); err != nil {
		log.Fatal(errMsg, zap.Error(err))
	}

	// 获取默认分类列表
	selectRet, err := r.CategorySelectList(db, dto.CategorySelectListRequest{
		IdList: []int64{1, 2, 3, 4},
		Origin: 1,
	})
	if err != nil {
		log.Fatal(errMsg, zap.Error(err))
	}

	now := time.Now().Unix()
	initSet := map[int64]dao.Category{
		1: {
			Id:         1,
			Name:       "未分类",
			Origin:     1,
			CreateTime: now,
			UpdateTime: now,
		},
		2: {
			Id:         2,
			Name:       "TV",
			Origin:     1,
			CreateTime: now,
			UpdateTime: now,
		},
		3: {
			Id:         3,
			Name:       "剧场版",
			Origin:     1,
			CreateTime: now,
			UpdateTime: now,
		},
		4: {
			Id:         4,
			Name:       "OVA",
			Origin:     1,
			CreateTime: now,
			UpdateTime: now,
		},
	}
	for _, item := range selectRet.Items {
		if _, ok := initSet[item.Id]; ok {
			delete(initSet, item.Id)
		}
	}

	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		log.Fatal(errMsg, zap.Error(err))
	}
	defer func() { _ = tx.Rollback() }()

	// prepare
	stmt, err := tx.Preparex(` insert or ignore into category (id,name,origin,create_time,update_time) values (?,?,?,?,?) `)
	if err != nil {
		log.Fatal(errMsg, zap.Error(err))
	}

	mask := syscall.Umask(0)
	defer syscall.Umask(mask)

	for _, item := range initSet {
		if _, err = stmt.Exec(item.Id, item.Name, item.Origin, item.CreateTime, item.UpdateTime); err != nil {
			log.Fatal(errMsg, zap.Error(err))
		}

		if err = os.MkdirAll(filepath.Join(setting.MEDIA_PATH, item.Name), 0766); err != nil {
			log.Fatal("创建分类目录失败", zap.Error(err))
		}
	}
	// commit
	if err = tx.Commit(); err != nil {
		log.Fatal("创建分类失败", zap.Error(err))
	}

	log.Info("数据库初始化成功")
}
func (r repositoryImpl) initQbittorrent(db *sqlx.DB) {
	var err error

	log.Info("登录qbittorrent")
	_, err = r.QBLogin(dto.QBLoginRequest{
		Username: setting.QB_USERNAME,
		Password: setting.QB_PASSWORD,
	})
	if err != nil {
		log.Fatal("qbittorrent 登录失败", zap.Error(err))
	}
	log.Info("qbittorrent 登录成功")

	// 获取qb分类
	categories, err := r.QBCategorySelectList(dto.QBCategorySelectListRequest{})
	if err != nil {
		log.Fatal("qbittorrent 获取分类失败", zap.Error(err))
	}

	// 如果分类不存在则创建
	if _, ok := categories.Map[setting.QB_CATEGORY]; !ok {
		log.Info("qbittorrent 初始化分类")
		_, err = r.QBCategoryInsert(dto.QBCategoryInsertRequest{
			Category: setting.QB_CATEGORY,
			SavePath: setting.QB_DOWNLOAD_PATH,
		})
		if err != nil {
			log.Fatal("qbittorrent 初始化分类失败", zap.Error(err),
				zap.String("category", setting.QB_CATEGORY),
				zap.String("qb_download_path", setting.QB_DOWNLOAD_PATH),
			)
		}
		log.Info("qbittorrent 初始化分类成功",
			zap.String("category", setting.QB_CATEGORY),
			zap.String("qb_download_path", setting.QB_DOWNLOAD_PATH))
	}

	res, err := r.QBRssFolderSelectList(dto.QBRssFolderSelectListRequest{})
	if err != nil {
		log.Fatal("qbittorrent 初始化rss目录信息失败", zap.Error(err))
	}
	_, ok := res.FolderMap[setting.QB_RSS_FOLDER]
	if !ok {
		log.Info("qbittorrent 初始化rss目录")
		_, err = r.QBRssFolderInsert(dto.QBRssFolderInsertRequest{
			Path: setting.QB_RSS_FOLDER,
		})
		if err != nil {
			log.Fatal("qbittorrent 初始化rss目录失败", zap.Error(err))
		}
		log.Info("qbittorrent 初始化rss目录成功")
	}
}

func (r repositoryImpl) count(sql string) string {
	return `select count(*) as total from ( ` + sql + `) `
}
