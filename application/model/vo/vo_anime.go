package vo

import (
	"umaru-server/application/model/dao"
)

type AnimeCreateRequest struct {
	BangumiId      int64    `json:"bangumi_id"`       // bangumi id
	Title          string   `json:"title"`            // 标题 not null
	CategoryId     int64    `json:"category_id"`      // 分类ID not null
	Season         int64    `json:"season"`           // 季
	Cover          string   `json:"cover"`            // 封面图
	Total          int64    `json:"total"`            // 总集数
	RssUrl         string   `json:"rss_url"`          // RSS链接
	PlayTime       int64    `json:"play_time"`        // 放送时间
	TorrentList    []string `json:"torrent_list"`     // 种子列表
	MustContain    string   `json:"must_contain"`     // 必须包含
	MustNotContain string   `json:"must_not_contain"` // 必须不包含
	EpisodeFilter  string   `json:"episode_filter"`   // 剧集过滤
	UseRegex       bool     `json:"use_regex"`        // 使用正则
	SmartFilter    bool     `json:"smart_filter"`     // 智能剧集过滤
}

func (a AnimeCreateRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if a.Title == "" {
		v.Add("标题不能为空")
		ok = true
	}

	if a.CategoryId == 0 {
		v.Add("分类不能为空")
		ok = true
	}
	if a.Season < 0 {
		v.Add("季必须大于0")
		ok = true
	}

	if a.RssUrl == "" && len(a.TorrentList) == 0 {
		v.Add("rss链接或种子列表至少一个不能为空")
		ok = true
	}

	for _, item := range a.TorrentList {
		if item == "" {
			v.Add("种子url不能为空")
			ok = true
			break
		}
	}

	msg = v.Message()
	return
}

type AnimeCreateResponse struct {
	Id int64 `json:"id"`
}
type AnimeGetRequest struct {
	Id        int64 `json:"id" form:"id"`
	BangumiId int64 `json:"bangumi_id" form:"bangumi_id"`
}

func (a AnimeGetRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if a.Id == 0 && a.BangumiId == 0 {
		v.Add("缺少id或bangumi_id")
		ok = true
	}

	if a.Id != 0 && a.Id < 0 {
		v.Add("ID必须大于0")
		ok = true
	}
	if a.BangumiId != 0 && a.BangumiId < 0 {
		v.Add("bangumi id 必须大于0")
		ok = true
	}

	msg = v.Message()
	return
}

type AnimeGetResponse struct {
	Anime dao.AnimeInfoView `json:"anime"`
}
type AnimeGetListRequest struct {
	Page          int64  `json:"page" form:"page"`
	Size          int64  `json:"size" form:"size"`
	Title         string `json:"title" form:"title"`                     // 名称模糊搜索
	CategoryId    int64  `json:"category_id" form:"category_id"`         // 分类搜索
	Sort          string `json:"sort" form:"sort"`                       // 排序字段
	PlayStartTime int64  `json:"play_start_time" form:"play_start_time"` // 放送时间范围的开始时间
	PlayEndTime   int64  `json:"play_end_time" form:"play_end_time"`     // 放送时间范围的结束时间
	AddStartTime  int64  `json:"add_start_time" form:"add_start_time"`   // 添加时间范围的开始时间
	AddEndTime    int64  `json:"add_end_time" form:"add_end_time"`       // 添加时间范围的结束时间
}

func (a *AnimeGetListRequest) ValidateError() (msg string, ok bool) {
	if a.Page < 0 {
		a.Page = 1
	}

	v := NewValidateMsg()
	if a.Size < 0 {
		v.Add("分页size不能小于0")
	}

	msg = v.Message()
	ok = v.HasMessage()
	return
}

type AnimeGetListResponse struct {
	Items []dao.AnimeInfoView `json:"items"`
	Total int64               `json:"total"`
}
type AnimeDeleteRequest struct {
	Id int64 `json:"id"`
}

type AnimeDeleteResponse struct {
}

type AnimeUpdateRequest struct {
	Id         int64  `json:"id"`          // ID
	Title      string `json:"title"`       // 标题
	Season     int64  `json:"season"`      // 季
	CategoryId int64  `json:"category_id"` // 分类ID
	Total      int64  `json:"total"`       // 总集数
	PlayTime   int64  `json:"play_time"`   // 放送时间
}

func (a AnimeUpdateRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if a.Id == 0 {
		v.Add("ID不能为空")
	}
	if a.Title == "" {
		v.Add("标题不能为空")
	}

	if a.Season < 0 {
		v.Add("季不能小于1")
	}

	if a.CategoryId <= 0 {
		v.Add("分类错误")
	}
	if a.Total < 0 {
		v.Add("集数不能小于0")
	}

	msg = v.Message()
	ok = v.HasMessage()
	return
}

type AnimeUpdateResponse struct {
}
type AnimeVideoGetListRequest struct {
	Id int64 `json:"id" form:"id"` // ID
}

type AnimeVideoGetListResponse struct {
	Items []dao.Video `json:"items"`
}

type AnimeRssCancelRequest struct {
	Id int64 `json:"id" form:"id"`
}

func (a AnimeRssCancelRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if a.Id <= 0 {
		v.Add("番剧ID不能为空")
	}
	msg = v.Message()
	ok = v.HasMessage()
	return
}

type AnimeRssCancelResponse struct {
}

type AnimeRssAddRequest struct {
	Id             int64  `json:"id" form:"id"`
	RssUrl         string `json:"rss_url" form:"rss_url"`
	MustContain    string `json:"must_contain"`     // 必须包含
	MustNotContain string `json:"must_not_contain"` // 必须不包含
	EpisodeFilter  string `json:"episode_filter"`   // 剧集过滤
	UseRegex       bool   `json:"use_regex"`        // 使用正则
	SmartFilter    bool   `json:"smart_filter"`     // 智能剧集过滤
}

func (a AnimeRssAddRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if a.Id <= 0 {
		v.Add("番剧ID不能为空")
	}
	msg = v.Message()
	ok = v.HasMessage()
	return
}

type AnimeRssAddResponse struct {
}
