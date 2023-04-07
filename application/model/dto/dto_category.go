package dto

import "umaru-server/application/model/dao"

type CategorySelectRequest struct {
	Id   int64  `json:"id"`   // ID
	Name string `json:"name"` // 名称
}
type CategorySelectResponse struct {
	Category dao.Category
}

type CategoryInsertRequest struct {
	Category dao.Category
}
type CategoryInsertResponse struct {
}

type CategorySelectListRequest struct {
	IdList []int64
	Origin int64
}
type CategorySelectListResponse struct {
	Items []dao.Category
}

type CategoryDeleteRequest struct {
	Id int64
}
type CategoryDeleteResponse struct {
}

type CategoryUpdateRequest struct {
	Category dao.Category
}
type CategoryUpdateResponse struct {
}
