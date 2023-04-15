package usecase

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"umaru-server/application/setting"
)

// GetLinkPath 获取硬链接路径
//	媒体目录/分类/番剧名称/S季/文件名
// @params
//	category string 分类名
//	title string 标题
//	season int64 季
//	episode int64 集
//	ext string 文件格式
// @return
//	string 媒体路径/分类/标题/S季/标题 - S季E集.ext 格式的路径
//	error 错误信息
func GetLinkPath(category string, title string, season int64, episode int64, ext string) (string, error) {

	if ext == "" {
		return "", errors.New("缺少文件格式")
	}
	if ext[0] != '.' {
		return "", errors.New("文件格式错误")
	}
	dir, err := GetLinkDir(category, title, season)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%s - S%02dE%02d%s", title, season, episode, ext)
	return filepath.Join(dir, filename), nil
}

// GetLinkDir 获取硬链接目录
//	媒体目录/分类/番剧名称/S季
// @params
//	category string 分类名
//	title string 标题
//	season int64 季
// @return
//	string 媒体路径/分类/标题/S季 格式的路径
//	error 错误信息
func GetLinkDir(category string, title string, season int64) (string, error) {
	if category == "" {
		return "", errors.New("缺少分类")
	}
	if title == "" {
		return "", errors.New("缺少标题")
	}
	if season == 0 {
		return "", errors.New("缺少季信息")
	}

	s := "S" + strconv.FormatInt(season, 10)
	return filepath.Join(setting.MEDIA_PATH, category, title, s), nil
}

// GetSourceDir 获取番剧下载目录
//	下载目录/番剧ID
// @params
//	animeId int64 番剧ID
// @return
//	string 下载目录/番剧ID 格式的文件夹路径
func GetSourceDir(animeId int64) string {
	return filepath.Join(setting.SOURCE_PATH, strconv.FormatInt(animeId, 10))
}

// GetLinkCategoryDir 获取分类的硬连接目录
//	@param categoryName string 分类名称
// @return string 分类的硬连接目录
func GetLinkCategoryDir(categoryName string) string {
	return filepath.Join(setting.MEDIA_PATH, categoryName)
}

// IsFileExists 文件是否存在
//	存在返回true 不存在返回false
// @params
//	path string 文件路径
// @return
//	bool 文件是否存在
func IsFileExists(path string) (ok bool, err error) {
	ok = false
	_, err = os.Stat(path)
	if err == nil {
		ok = true
		return
	}
	if os.IsNotExist(err) {
		ok = false
		err = nil
		return
	}
	ok = false
	return
}
