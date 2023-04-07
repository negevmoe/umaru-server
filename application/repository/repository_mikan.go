package repository

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"umaru-server/application/model/dao"
	"umaru-server/application/model/dto"
)

func (r repositoryImpl) MikanInfoSelect(req dto.MikanInfoSelectRequest) (res dto.MikanInfoSelectResponse, err error) {
	request, err := http.NewRequest("GET", "https://mikanani.me/Home/Search", nil)
	if err != nil {
		log.Error("获取mikan信息失败", zap.Error(err), zap.String("name", req.Name))
		return
	}

	q := request.URL.Query()
	q.Add("searchstr", req.Name)
	request.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Error("获取mikan信息失败,请求失败,", zap.Error(err), zap.String("name", req.Name))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("http request status:%d", resp.StatusCode))
		log.Error("获取mikan信息失败,请求失败", zap.Error(err), zap.String("name", req.Name))
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error("获取mikan信息失败,css选择器解析html失败,", zap.Error(err), zap.String("name", req.Name))
		return
	}

	doc.Find("div.central-container ul li").Each(func(i int, s *goquery.Selection) {
		if i > 0 {
			return
		}
		u, found := s.Find("a").Attr("href")
		if found {
			res.Url = "https://mikanani.me" + u
		}
	})

	return
}

func (r repositoryImpl) MikanRssSelectList(req dto.MikanRssSelectListRequest) (res dto.MikanRssSelectListResponse, err error) {
	fmt.Println(req.Url)
	resp, err := http.Get(req.Url)
	if err != nil {
		log.Error("获取mikan rss 失败", zap.Error(err), zap.String("mikan_url", req.Url))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("http request status:%d", resp.StatusCode))
		log.Error("获取mikan rss 失败", zap.Error(err), zap.String("mikan_url", req.Url))
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Error("获取mikan rss 失败,css选择器解析html失败,", zap.Error(err), zap.String("mikan_url", req.Url))
		return
	}

	res.Items = make([]dao.MikanRss, 0)
	doc.Find("div.leftbar-nav ul li").Each(func(i int, s *goquery.Selection) {
		el := s.Find("a")

		// 获取字幕组名称
		name := el.Text()
		// 获取字幕组ID
		idStr, found := el.Attr("data-anchor")
		if !found {
			return
		}
		id := strings.TrimPrefix(idStr, "#")
		// 获取mikan id
		arr := strings.Split(req.Url, "/")
		mid := arr[len(arr)-1]
		// 获取rss
		rss := fmt.Sprintf("https://mikanani.me/RSS/Bangumi?bangumiId=%s&subgroupid=%s", mid, id)

		res.Items = append(res.Items, dao.MikanRss{
			Id:   id,
			Name: name,
			Rss:  rss,
		})
	})

	return
}
