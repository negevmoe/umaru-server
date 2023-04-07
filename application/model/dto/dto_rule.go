package dto

import "umaru/application/model/dao"

type RuleSelectListRequest struct {
}
type RuleSelectListResponse struct {
	Items []dao.Rule
}

type RuleSelectByNameRequest struct {
	Name string
}
type RuleSelectByNameResponse struct {
	Rule dao.Rule
}

type RuleUpdateRequest struct {
	Rule dao.Rule
}
type RuleUpdateResponse struct {
}

type RuleDeleteListRequest struct {
	IdList []int64
}
type RuleDeleteListResponse struct {
}

type RuleInsertRequest struct {
	Rule dao.Rule
}
type RuleInsertResponse struct {
}
