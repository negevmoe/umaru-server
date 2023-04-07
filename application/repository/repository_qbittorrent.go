package repository

import (
	"encoding/json"
	"go.uber.org/zap"
	"umaru/application/model/dto"
)

func (r repositoryImpl) QBLogin(req dto.QBLoginRequest) (res dto.QBLoginResponse, err error) {
	err = reqHandler(qb.R().SetFormData(map[string]string{
		"username": req.Username,
		"password": req.Password,
	}).Post("/auth/login"))
	if err != nil {
		log.Error("qbittorrent登录失败", zap.Error(err),
			zap.String("username", req.Username),
			zap.String("password", req.Password),
		)
	}
	return
}

func (r repositoryImpl) QBLogSelectList(req dto.QBLogSelectListRequest) (res dto.QBLogSelectListResponse, err error) {
	err = reqHandler(qb.R().SetResult(&res.Items).Get("/log/main"))
	if err != nil {
		log.Error("获取qbittorrent日志失败", zap.Error(err))
		return
	}
	return
}

func (r repositoryImpl) QBCategoryInsert(req dto.QBCategoryInsertRequest) (res dto.QBCategoryInsertResponse, err error) {
	err = reqHandler(qb.R().SetFormData(map[string]string{
		"category": req.Category,
		"savePath": req.SavePath,
	}).Post("/torrents/createCategory"))
	if err != nil {
		log.Error("创建qbittorrent分类失败", zap.Error(err),
			zap.String("category", req.Category),
			zap.String("savePath", req.SavePath),
		)
		return
	}
	return
}

func (r repositoryImpl) QBRuleSet(req dto.QBRuleSetRequest) (res dto.QBRuleSetResponse, err error) {
	ruleDef, err := json.Marshal(req.RuleDef)
	if err != nil {
		log.Error("qbittorrent 设置规则失败", zap.Error(err), zap.String("rule_name", req.RuleName), zap.Any("rule_def", req.RuleDef))
		return
	}
	err = reqHandler(qb.R().SetQueryParams(map[string]string{
		"ruleName": req.RuleName,
		"ruleDef":  string(ruleDef),
	}).Get("/rss/setRule"))
	if err != nil {
		log.Error("qbittorrent 设置规则失败", zap.Error(err))
		return
	}
	return
}

func (r repositoryImpl) QBRuleDelete(req dto.QBRuleDeleteRequest) (res dto.QBRuleDeleteResponse, err error) {
	err = reqHandler(qb.R().SetQueryParams(map[string]string{
		"ruleName": req.RuleName,
	}).Get("/rss/removeRule"))

	if err != nil {
		log.Error("qbittorrent 删除规则失败", zap.Error(err), zap.String("rule_name", req.RuleName))
		return
	}
	return
}

func (r repositoryImpl) QBRssInsert(req dto.QBRssInsertRequest) (res dto.QBRssInsertResponse, err error) {
	err = reqHandler(qb.R().SetFormData(map[string]string{
		"url":  req.Url,
		"path": req.Path,
	}).Post("/rss/addFeed"))
	if err != nil {
		log.Error("qbittorrent 添加RSS失败", zap.Error(err),
			zap.String("url", req.Url),
			zap.String("path", req.Path),
		)
		return
	}
	return
}

func (r repositoryImpl) QBRssDelete(req dto.QBRssDeleteRequest) (res dto.QBRssDeleteResponse, err error) {
	err = reqHandler(qb.R().SetQueryParams(map[string]string{
		"path": req.Path,
	}).Get("/rss/removeItem"))
	if err != nil {
		log.Error("qbittorrent 删除RSS失败", zap.Error(err),
			zap.String("path", req.Path),
		)
		return
	}
	return
}

func (r repositoryImpl) QBTorrentInsertList(req dto.QBTorrentInsertListRequest) (res dto.QBTorrentInsertListResponse, err error) {
	err = reqHandler(qb.R().SetFormDataAnyType(map[string]interface{}{
		"urls":     req.Urls,
		"category": req.Category,
		"savepath": req.SavePath,
	}).Post("/torrents/add"))
	if err != nil {
		log.Error("qbittorrent 添加种子失败", zap.Error(err),
			zap.String("urls", req.Urls),
			zap.String("category", req.Category),
			zap.String("savepath", req.SavePath),
		)
		return
	}
	return
}

func (r repositoryImpl) QBCategorySelectList(req dto.QBCategorySelectListRequest) (res dto.QBCategorySelectListResponse, err error) {
	m := make(map[string]dto.QBCategory)

	err = reqHandler(qb.R().SetResult(&m).Get("/torrents/categories"))
	if err != nil {
		log.Error("qbittorrent 获取分类失败", zap.Error(err))
		return
	}
	res.Map = m
	return
}

func (r repositoryImpl) QBRssFolderSelectList(req dto.QBRssFolderSelectListRequest) (res dto.QBRssFolderSelectListResponse, err error) {
	m := make(map[string]any)
	err = reqHandler(qb.R().SetResult(&m).Get("/rss/items"))
	if err != nil {
		log.Error("qbittorrent 获取RSS文件夹失败", zap.Error(err))
		return
	}

	res.FolderMap = make(map[string]struct{})
	res.FolderList = make([]string, 0, len(m))

	for folder := range m {
		res.FolderList = append(res.FolderList, folder)
		res.FolderMap[folder] = struct{}{}
	}

	return
}

func (r repositoryImpl) QBRssFolderInsert(req dto.QBRssFolderInsertRequest) (res dto.QBRssFolderInsertResponse, err error) {
	err = reqHandler(qb.R().SetQueryParams(map[string]string{
		"path": req.Path,
	}).Get("/rss/addFolder"))
	if err != nil {
		log.Error("qbittorrent 创建RSS文件夹失败", zap.Error(err),
			zap.String("path", req.Path),
		)
		return
	}
	return
}
