package handler

import (
	"time"
	"umaru-server/application/model/dao"
	"umaru-server/application/model/dto"
	"umaru-server/application/model/vo"
)

func (s handlerImpl) RuleGetList(req vo.RuleGetListRequest) (res vo.RuleGetListResponse, err error) {
	ret, err := repo.RuleSelectList(db, dto.RuleSelectListRequest{})
	if err != nil {
		err = vo.ErrorWrap(500, "获取规则列表失败", err)
		return
	}
	res.Items = make([]vo.RuleGetListResponseItem, 0, len(ret.Items))
	for _, item := range ret.Items {
		res.Items = append(res.Items, vo.RuleGetListResponseItem{
			Rule:        item,
			UseRegex:    item.UseRegex == 1,
			SmartFilter: item.SmartFilter == 1,
		})
	}
	return
}

func (s handlerImpl) RuleCreate(req vo.RuleCreateRequest) (res vo.RuleCreateResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	now := time.Now().Unix()
	rule := dao.Rule{
		Name:           req.Name,
		MustContain:    req.MustContain,
		MustNotContain: req.MustNotContain,
		UseRegex:       2,
		EpisodeFilter:  req.EpisodeFilter,
		SmartFilter:    2,
		CreateTime:     now,
		UpdateTime:     now,
	}
	if req.UseRegex {
		rule.UseRegex = 1
	}
	if req.SmartFilter {
		rule.SmartFilter = 1
	}
	_, err = repo.RuleInsert(db, dto.RuleInsertRequest{Rule: rule})
	if err != nil {
		err = vo.ErrorWrap(500, "创建规则失败", err)
		return
	}
	return
}

func (s handlerImpl) RuleUpdate(req vo.RuleUpdateRequest) (res vo.RuleUpdateResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	now := time.Now().Unix()
	rule := dao.Rule{
		Id:             req.Id,
		Name:           req.Name,
		MustContain:    req.MustContain,
		MustNotContain: req.MustNotContain,
		UseRegex:       2,
		EpisodeFilter:  req.EpisodeFilter,
		SmartFilter:    2,
		UpdateTime:     now,
	}
	if req.UseRegex {
		rule.UseRegex = 1
	}
	if req.SmartFilter {
		rule.SmartFilter = 1
	}
	_, err = repo.RuleUpdate(db, dto.RuleUpdateRequest{Rule: rule})
	if err != nil {
		err = vo.ErrorWrap(500, "更新规则失败", err)
		return
	}
	return
}

func (s handlerImpl) RuleDeleteList(req vo.RuleDeleteListRequest) (res vo.RuleDeleteListResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	_, err = repo.RuleDeleteList(db, dto.RuleDeleteListRequest{IdList: req.IdList})
	if err != nil {
		err = vo.ErrorWrap(500, "删除规则失败", err)
		return
	}
	return
}
