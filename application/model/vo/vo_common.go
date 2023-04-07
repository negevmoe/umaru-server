package vo

import (
	"github.com/robfig/cron/v3"
	"umaru/application/model/dao"
)

type UserLoginRequest struct {
	Username string `json:"username" ` // 用户名
	Password string `json:"password" ` // 密码
}

func (receiver UserLoginRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if receiver.Username == "" {
		v.Add("用户名不能为空")
		ok = true
	}
	if receiver.Password == "" {
		v.Add("密码不能为空")
		ok = true
	}
	msg = v.Message()
	return
}

type UserLoginResponse struct {
	Token string `json:"token"` // token
}

type LinkRequest struct{}
type LinkResponse struct{}

type GetQbLogsRequest struct {
}
type GetQbLogsResponse struct {
	Items []dao.QbLog `json:"items"`
}

type CronGetListRequest struct {
}
type CronGetListResponse struct {
	Items []dao.Cron `json:"items"`
}
type CronUpdateRequest struct {
	Id  cron.EntryID `json:"id"`
	Sep string       `json:"sep"`
}
type CronUpdateResponse struct {
}

type CronRunItRequest struct {
	Id cron.EntryID `json:"id" form:"id"`
}
type CronRunItResponse struct {
}
type CronStopRequest struct {
}
type CronStopResponse struct {
}
type CronStartRequest struct {
}
type CronStartResponse struct {
}
