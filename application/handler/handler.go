package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"umaru-server/application/global"
	"umaru-server/application/model/dto"
	"umaru-server/application/model/vo"
	"umaru-server/application/repository"
	"umaru-server/application/tool"
	"umaru-server/application/usecase"
)

var db *sqlx.DB
var log *zap.Logger
var repo = repository.Repo

var Server IHandler = handlerImpl{}

type IHandler interface {
	UserLogin(req vo.UserLoginRequest) (res vo.UserLoginResponse, err error)
	Link(req vo.LinkRequest) (res vo.LinkResponse, err error)
	ParseRss(req vo.ParseRssRequest) (res vo.ParseRssResponse, err error)          // 解析rss
	CronGetList(req vo.CronGetListRequest) (res vo.CronGetListResponse, err error) // 获取定时任务列表
	CronUpdate(req vo.CronUpdateRequest) (res vo.CronUpdateResponse, err error)    // 更新定时任务
	CronRunIt(req vo.CronRunItRequest) (res vo.CronRunItResponse, err error)       // 立即运行一次定时任务
	CronStop(req vo.CronStopRequest) (res vo.CronStopResponse, err error)          // 停止定时任务
	CronStart(req vo.CronStartRequest) (res vo.CronStartResponse, err error)       // 开始定时任务

	AnimeCreate(req vo.AnimeCreateRequest) (res vo.AnimeCreateResponse, err error)                   // 创建番剧
	AnimeGet(req vo.AnimeGetRequest) (res vo.AnimeGetResponse, err error)                            // 获取番剧详情
	AnimeGetList(req vo.AnimeGetListRequest) (res vo.AnimeGetListResponse, err error)                // 获取番剧列表
	AnimeDelete(req vo.AnimeDeleteRequest) (res vo.AnimeDeleteResponse, err error)                   // 删除番剧
	AnimeUpdate(req vo.AnimeUpdateRequest) (res vo.AnimeUpdateResponse, err error)                   // 更新番剧
	AnimeVideoGetList(req vo.AnimeVideoGetListRequest) (res vo.AnimeVideoGetListResponse, err error) // 获取番剧下载的视频列表
	MikanGetRssList(req vo.MikanGetRssListRequest) (res vo.MikanGetRssListResponse, err error)       // 获取mikan的字幕组信息
	AnimeRssCancel(req vo.AnimeRssCancelRequest) (res vo.AnimeRssCancelResponse, err error)          //
	AnimeRssAdd(req vo.AnimeRssAddRequest) (res vo.AnimeRssAddResponse, err error)                   //

	CategoryCreate(req vo.CategoryCreateRequest) (res vo.CategoryCreateResponse, err error)    // 创建分类
	CategoryGetList(req vo.CategoryGetListRequest) (res vo.CategoryGetListResponse, err error) // 获取分类列表
	CategoryDelete(req vo.CategoryDeleteRequest) (res vo.CategoryDeleteResponse, err error)    // 批量删除分类
	CategoryUpdate(req vo.CategoryUpdateRequest) (res vo.CategoryUpdateResponse, err error)    // 更新分类
	RuleGetList(req vo.RuleGetListRequest) (res vo.RuleGetListResponse, err error)             // 获取下载规则列表
	RuleCreate(req vo.RuleCreateRequest) (res vo.RuleCreateResponse, err error)                // 创建下载规则
	RuleUpdate(req vo.RuleUpdateRequest) (res vo.RuleUpdateResponse, err error)                // 更新下载规则
	RuleDeleteList(req vo.RuleDeleteListRequest) (res vo.RuleDeleteListResponse, err error)    // 批量删除下载规则
	GetQbLogs(req vo.GetQbLogsRequest) (res vo.GetQbLogsResponse, err error)                   // 获取QB日志
}

type handlerImpl struct{}

func Init() {
	db = global.Sqlite
	log = global.Log

	CronServer.Init()
	CronServer.Start()
}

func (s handlerImpl) UserLogin(req vo.UserLoginRequest) (res vo.UserLoginResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	fmt.Println(req.Username)
	fmt.Println(req.Password)
	res.Token = "token"
	return
}

func (s handlerImpl) Link(req vo.LinkRequest) (res vo.LinkResponse, err error) {
	Link()
	return
}

// GetQbLogs 获取qbittorrent日志
func (s handlerImpl) GetQbLogs(req vo.GetQbLogsRequest) (res vo.GetQbLogsResponse, err error) {
	ret, err := repo.QBLogSelectList(dto.QBLogSelectListRequest{})
	if err != nil {
		err = vo.ErrorWrap(500, "获取qbittorrent日志失败", err)
		return
	}
	res.Items = ret.Items
	return
}

func ClearEmptyDir() {
	LinkLock.Lock()
	defer LinkLock.Unlock()
	// 清理媒体文件夹下的空目录
	categoryListRet, err := repo.CategorySelectList(db, dto.CategorySelectListRequest{})
	if err != nil {
		log.Error("定期清理空目录失败,获取分类列表失败", zap.Error(err))
		return
	}

	for _, category := range categoryListRet.Items {
		list, err := tool.RemoveEmptyDirAll(usecase.GetLinkCategoryDir(category.Name), false)
		if err != nil {
			log.Error("定期清理空目录失败,获取分类列表失败", zap.Error(err))
			return
		}

		for _, item := range list {
			log.Info("检测到空目录,已清理", zap.String("path", item))
		}
	}
}
