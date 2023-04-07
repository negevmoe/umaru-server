package dto

import "umaru/application/model/dao"

type MikanInfoSelectRequest struct {
	Name string
}
type MikanInfoSelectResponse struct {
	Url string
}
type MikanRssSelectListRequest struct {
	Url string
}
type MikanRssSelectListResponse struct {
	Items []dao.MikanRss `json:"items"`
}
