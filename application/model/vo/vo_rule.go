package vo

import (
	"umaru-server/application/model/dao"
)

type RuleGetListRequest struct {
}
type RuleGetListResponseItem struct {
	dao.Rule
	UseRegex    bool `json:"use_regex"`
	SmartFilter bool `json:"smart_filter"`
}
type RuleGetListResponse struct {
	Items []RuleGetListResponseItem `json:"items"`
}

type RuleCreateRequest struct {
	Name           string `json:"name"`             // 名称
	MustContain    string `json:"must_contain"`     // 必须包含
	MustNotContain string `json:"must_not_contain"` // 必须不包含
	EpisodeFilter  string `json:"episode_filter"`   // 剧集过滤
	UseRegex       bool   `json:"use_regex"`        // 正则表达式
	SmartFilter    bool   `json:"smart_filter"`     // 智能剧集过滤
}

func (r RuleCreateRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if r.Name == "" {
		v.Add("规则名称不能为空")
		ok = true
	}
	msg = v.Message()
	return
}

type RuleCreateResponse struct {
}
type RuleUpdateRequest struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`             // 名称
	MustContain    string `json:"must_contain"`     // 必须包含
	MustNotContain string `json:"must_not_contain"` // 必须不包含
	EpisodeFilter  string `json:"episode_filter"`   // 剧集过滤
	UseRegex       bool   `json:"use_regex"`        // 正则表达式
	SmartFilter    bool   `json:"smart_filter"`     // 智能剧集过滤
}

func (r RuleUpdateRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if r.Name == "" {
		v.Add("规则名称不能为空")
		ok = true
	}
	msg = v.Message()
	return
}

type RuleUpdateResponse struct {
}
type RuleDeleteListRequest struct {
	IdList []int64 `json:"id_list"`
}

func (r RuleDeleteListRequest) ValidateError() (msg string, ok bool) {
	v := NewValidateMsg()
	if len(r.IdList) == 0 {
		v.Add("ID列表不能为空")
		ok = true
	}
	for _, item := range r.IdList {
		if item == 0 {
			v.Add("ID不能为空")
			ok = true
			break
		}
	}
	msg = v.Message()
	return
}

type RuleDeleteListResponse struct {
}
