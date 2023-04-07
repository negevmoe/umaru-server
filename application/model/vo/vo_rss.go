package vo

import "umaru/application/model/dao"

type ParseRssRequest struct {
	Url            string `json:"url"`              // rss链接
	MustContain    string `json:"must_contain"`     // 必须包含
	MustNotContain string `json:"must_not_contain"` // 必须不包含
	UseRegex       int    `json:"use_regex"`        // 正则表达式 1:true 2:false 默认2
	EpisodeFilter  string `json:"episode_filter"`   // 剧集过滤
	SmartFilter    int    `json:"smart_filter"`     // 智能剧集过滤 1:true 2:false 默认2
}

type ParseRssResponse struct {
	Feed dao.Feed `json:"feed"`
}

type MikanGetRssListRequest struct {
	Name string `json:"name" form:"name"`
}

func (a MikanGetRssListRequest) ValidateError() (msg string, ok bool) {
	if a.Name == "" {
		msg = "番剧名称不能为空"
		ok = true
		return
	}
	return
}

type MikanGetRssListResponse struct {
	Items []dao.MikanRss `json:"items"`
}
