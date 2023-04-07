package handler

import (
	"go.uber.org/zap"
	"time"
	"umaru/application/model/vo"
)

func (s handlerImpl) CronGetList(req vo.CronGetListRequest) (res vo.CronGetListResponse, err error) {
	list := CronServer.List()
	res.Items = list
	return
}

func (s handlerImpl) CronUpdate(req vo.CronUpdateRequest) (res vo.CronUpdateResponse, err error) {
	err = CronServer.UpdateCron(req.Id, req.Sep)
	if err != nil {
		err = vo.ErrorWrap(500, "更新定时任务失败", err)
		return
	}
	return
}

func (s handlerImpl) CronRunIt(req vo.CronRunItRequest) (res vo.CronRunItResponse, err error) {
	c, err := CronServer.Get(req.Id)
	if err != nil {
		err = vo.ErrorWrap(500, "定时任务运行失败", err)
		return
	}

	err = CronServer.RunIt(req.Id)
	if err != nil {
		err = vo.ErrorWrap(500, "定时任务运行失败", err)
		return
	}

	log.Info("运行定时任务", zap.String("name", c.Name), zap.Time("time", time.Now()))
	return
}

func (s handlerImpl) CronStop(req vo.CronStopRequest) (res vo.CronStopResponse, err error) {
	CronServer.Stop()
	log.Info("定时任务已停止", zap.Time("time", time.Now()))
	return
}

func (s handlerImpl) CronStart(req vo.CronStartRequest) (res vo.CronStartResponse, err error) {
	CronServer.Start()
	log.Info("定时任务已启动", zap.Time("time", time.Now()))
	return
}
