package handler

import (
	"github.com/mmcdole/gofeed"
	"go.uber.org/zap"
	"strconv"
	"umaru-server/application/model/dao"
	"umaru-server/application/model/dto"
	"umaru-server/application/model/vo"
)

// ParseRss 解析rss
func (s handlerImpl) ParseRss(req vo.ParseRssRequest) (res vo.ParseRssResponse, err error) {

	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(req.Url)
	if err != nil {
		log.Error("rss url解析失败", zap.Error(err), zap.String("url", req.Url))
		err = vo.ErrorWrap(500, "获取rss失败", err)
		return
	}

	var result dao.Feed
	result.Title = feed.Title
	result.Desc = feed.Description

	for _, item := range feed.Items {
		length, _ := strconv.Atoi(item.Enclosures[0].Length)
		result.Items = append(result.Items, dao.FeedItem{
			Title:   item.Title,
			Desc:    item.Description,
			PubDate: item.Published,
			Url:     item.Enclosures[0].URL,
			Length:  length,
		})
	}

	res.Feed = result
	return
}

func (s handlerImpl) MikanGetRssList(req vo.MikanGetRssListRequest) (res vo.MikanGetRssListResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	info, err := repo.MikanInfoSelect(dto.MikanInfoSelectRequest{
		Name: req.Name,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "获取mikan信息失败", err)
		return
	}

	if info.Url == "" {
		res.Items = make([]dao.MikanRss, 0)
		return
	}

	rssList, err := repo.MikanRssSelectList(dto.MikanRssSelectListRequest{
		Url: info.Url,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "获取mikan信息失败", err)
		return
	}
	res.Items = rssList.Items
	return
}
