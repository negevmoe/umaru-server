package vo

import (
	"umaru/application/model/dao"
)

type CategoryCreateRequest struct {
	Name string `db:"name"` // 分类名称
}

func (c CategoryCreateRequest) ValidateError() (msg string, ok bool) {
	if c.Name == "" {
		msg = "分类名称不能为空"
		ok = true
	}
	return
}

type CategoryCreateResponse struct{}

type CategoryGetListRequest struct{}
type CategoryGetListResponse struct {
	Items []dao.Category `json:"items"`
}
type CategoryDeleteRequest struct {
	Id int64 `json:"id"`
}

func (c CategoryDeleteRequest) ValidateError() (msg string, ok bool) {
	if c.Id <= 0 {
		msg = "分类ID不能为空"
		ok = true
		return
	}
	if c.Id == 1 {
		msg = "默认分类不能删除"
		ok = true
		return
	}
	return
}

type CategoryDeleteResponse struct {
}
type CategoryUpdateRequest struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (c CategoryUpdateRequest) ValidateError() (msg string, ok bool) {
	if c.Id <= 0 {
		msg = "分类ID不能为空"
		ok = true
		return
	}

	if c.Name == "" {
		msg += "分类名不能为空"
		ok = true
		return
	}
	return
}

type CategoryUpdateResponse struct {
}
