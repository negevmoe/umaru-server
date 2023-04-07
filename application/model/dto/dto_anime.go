package dto

import "umaru/application/model/dao"

type AnimeSelectRequest struct {
	Id        int64
	BangumiId int64
}
type AnimeSelectResponse struct {
	Anime dao.Anime
}

type AnimeSelectByTitleAndSeasonRequest struct {
	Title  string
	Season int64
}
type AnimeSelectByTitleAndSeasonResponse struct {
	Anime dao.Anime
}

type AnimeSelectListRequest struct {
	CategoryId int64 `json:"category_id" form:"category_id"` // 分类搜索
}
type AnimeSelectListResponse struct {
	Items []dao.Anime
}

type AnimeInfoViewSelectRequest struct {
	Id        int64
	BangumiId int64
}
type AnimeInfoViewSelectResponse struct {
	AnimeInfo dao.AnimeInfoView
}

type AnimeInfoViewSelectListRequest struct {
	Limit         int64
	Offset        int64
	Title         string // 名称模糊搜索
	CategoryId    int64  // 分类搜索
	Sort          string // 排序
	PlayStartTime int64  // 放送时间范围的开始时间
	PlayEndTime   int64  // 放送时间范围的结束时间
	AddStartTime  int64  // 添加时间范围的开始时间
	AddEndTime    int64  // 添加时间范围的结束时间
}
type AnimeInfoViewSelectListResponse struct {
	Items []dao.AnimeInfoView
	Total int64 `db:"total"`
}

type AnimeInsertRequest struct {
	Anime dao.Anime
}
type AnimeInsertResponse struct {
	Id int64 `db:"id"`
}

type AnimeDeleteRequest struct {
	Id int64
}
type AnimeDeleteResponse struct {
}

type AnimeUpdateRequest struct {
	Anime dao.Anime
}
type AnimeUpdateResponse struct {
}

type AnimeCategoryUpdateRequest struct {
	IdList     []int64
	CategoryId int64
}
type AnimeCategoryUpdateResponse struct {
}

type AnimeRssUrlUpdateRequest struct {
	Id      int64
	RssUrl  string
	RssPath string
}
type AnimeRssUrlUpdateResponse struct {
}
