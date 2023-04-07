package dto

import "umaru/application/model/dao"

type QBCategorySelectListRequest struct {
}

type QBCategory struct {
	Name     string `json:"name"`
	SavePath string `json:"savePath"`
}
type QBCategorySelectListResponse struct {
	Map map[string]QBCategory
}

type QBCategoryInsertRequest struct {
	Category string `json:"category"`
	SavePath string `json:"savePath"`
}
type QBCategoryInsertResponse struct {
}

type QBRuleSelectListRequest struct {
}
type QBRuleSelectListResponse struct {
}

type QBRuleSetRequest struct {
	RuleName string      `json:"ruleName"`
	RuleDef  dao.RuleDef `json:"ruleDef"`
}
type QBRuleSetResponse struct {
}

type QBRuleDeleteRequest struct {
	RuleName string
}
type QBRuleDeleteResponse struct {
}

type QBRssInsertRequest struct {
	Url  string
	Path string
}
type QBRssInsertResponse struct {
}

type QBRssDeleteRequest struct {
	Path string
}
type QBRssDeleteResponse struct {
}

type QBLogSelectListRequest struct {
}
type QBLogSelectListResponse struct {
	Items []dao.QbLog
}

type QBTorrentInsertListRequest struct {
	Urls     string
	Category string
	SavePath string
}
type QBTorrentInsertListResponse struct {
}

type QBLoginRequest struct {
	Username string
	Password string
}
type QBLoginResponse struct {
}

type QBRssFolderSelectListRequest struct {
}
type QBRssFolderSelectListResponse struct {
	FolderList []string
	FolderMap  map[string]struct{}
}
type QBRssFolderInsertRequest struct {
	Path string
}
type QBRssFolderInsertResponse struct {
}
