package restful

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"umaru-server/application/handler"
	"umaru-server/application/model/vo"
	"umaru-server/application/setting"
)

func Run() error {
	if setting.SERVER_DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.POST("/api/user/login", UserLogin) // 登录

	router.GET("/api/link", MiddleAuth(), Link) // 硬连接

	router.POST("/api/anime", MiddleAuth(), AnimeCreate)                 // 创建番剧(必须订阅rss或上传种子)
	router.GET("/api/anime", MiddleAuth(), AnimeGet)                     // 获取番剧详情
	router.GET("/api/anime_list", MiddleAuth(), AnimeGetList)            // 获取番剧列表
	router.DELETE("/api/anime", MiddleAuth(), AnimeDelete)               // 删除番剧
	router.PUT("/api/anime", MiddleAuth(), AnimeUpdate)                  // 更新番剧信息
	router.GET("/api/anime/rss/cancel", MiddleAuth(), AnimeRssCancel)    // 取消订阅
	router.POST("/api/anime/rss", MiddleAuth(), AnimeRssAdd)             // 添加rss
	router.GET("/api/anime/video_list", MiddleAuth(), AnimeVideoGetList) // 获取番剧的视频列表
	router.GET("/api/anime/rss", MiddleAuth(), MikanGetRssList)          // 获取番剧的mikan信息

	router.POST("/api/category", MiddleAuth(), CategoryCreate)      // 创建分类
	router.GET("/api/category_list", MiddleAuth(), CategoryGetList) // 获取分类列表
	router.DELETE("/api/category", MiddleAuth(), CategoryDelete)    // 删除分类
	router.PUT("/api/category", MiddleAuth(), CategoryUpdate)       // 更新分类

	router.GET("/api/rule_list", MiddleAuth(), RuleGetList)       // 获取规则列表
	router.POST("/api/rule", MiddleAuth(), RuleCreate)            // 创建规则
	router.PUT("/api/rule", MiddleAuth(), RuleUpdate)             // 更新规则
	router.DELETE("/api/rule_list", MiddleAuth(), RuleDeleteList) // 批量删除规则

	router.GET("/api/qb/logs", MiddleAuth(), GetQbLogs)   // 获取qbittorrent日志
	router.POST("/api/rss/parse", MiddleAuth(), ParseRss) // 解析rss

	router.GET("/api/cron_list", MiddleAuth(), CronGetList)  // 定时任务 获取任务列表
	router.PUT("/api/cron/update", MiddleAuth(), CronUpdate) // 定时任务 更新
	router.GET("/api/cron/run", MiddleAuth(), CronRunIt)     // 定时任务 立即运行一次
	router.GET("/api/cron/stop", MiddleAuth(), CronStop)     // 定时任务 停止
	router.GET("/api/cron/start", MiddleAuth(), CronStart)   // 定时任务 开始

	router.LoadHTMLGlob("./dist/*.html")
	router.Static("/assets", "./dist/assets")
	router.StaticFile("/", "dist/index.html")
	return router.Run(fmt.Sprintf(":%d", setting.SERVER_PORT))
}

func MikanGetRssList(ctx *gin.Context) {
	var req vo.MikanGetRssListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.MikanGetRssList(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}

func ParseRss(ctx *gin.Context) {
	var req vo.ParseRssRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.ParseRss(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}

func UserLogin(ctx *gin.Context) {
	var req vo.UserLoginRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.UserLogin(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func Link(ctx *gin.Context) {
	var req vo.LinkRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.Link(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func AnimeCreate(ctx *gin.Context) {
	var req vo.AnimeCreateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeCreate(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func AnimeGet(ctx *gin.Context) {
	var req vo.AnimeGetRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeGet(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func AnimeGetList(ctx *gin.Context) {
	var req vo.AnimeGetListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeGetList(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func AnimeDelete(ctx *gin.Context) {
	var req vo.AnimeDeleteRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeDelete(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func AnimeUpdate(ctx *gin.Context) {
	var req vo.AnimeUpdateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeUpdate(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}

func AnimeRssCancel(ctx *gin.Context) {
	var req vo.AnimeRssCancelRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeRssCancel(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func AnimeRssAdd(ctx *gin.Context) {
	var req vo.AnimeRssAddRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeRssAdd(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}

func AnimeVideoGetList(ctx *gin.Context) {
	var req vo.AnimeVideoGetListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.AnimeVideoGetList(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func CategoryCreate(ctx *gin.Context) {
	var req vo.CategoryCreateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CategoryCreate(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func CategoryGetList(ctx *gin.Context) {
	var req vo.CategoryGetListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CategoryGetList(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}

func CategoryDelete(ctx *gin.Context) {
	var req vo.CategoryDeleteRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CategoryDelete(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func CategoryUpdate(ctx *gin.Context) {
	var req vo.CategoryUpdateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CategoryUpdate(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func RuleGetList(ctx *gin.Context) {
	var req vo.RuleGetListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.RuleGetList(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func RuleCreate(ctx *gin.Context) {
	var req vo.RuleCreateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}
	res, err := handler.Server.RuleCreate(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func RuleUpdate(ctx *gin.Context) {
	var req vo.RuleUpdateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.RuleUpdate(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func RuleDeleteList(ctx *gin.Context) {
	var req vo.RuleDeleteListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.RuleDeleteList(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}

func GetQbLogs(ctx *gin.Context) {
	res, err := handler.Server.GetQbLogs(vo.GetQbLogsRequest{})
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}

func CronGetList(ctx *gin.Context) {
	var req vo.CronGetListRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CronGetList(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func CronUpdate(ctx *gin.Context) {
	var req vo.CronUpdateRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CronUpdate(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func CronRunIt(ctx *gin.Context) {
	var req vo.CronRunItRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}
	res, err := handler.Server.CronRunIt(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func CronStop(ctx *gin.Context) {
	var req vo.CronStopRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CronStop(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
func CronStart(ctx *gin.Context) {
	var req vo.CronStartRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.Error(ctx, vo.ErrorBind(err))
		return
	}

	res, err := handler.Server.CronStart(req)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, res)
}
