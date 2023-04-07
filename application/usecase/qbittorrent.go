package usecase

import (
	"strconv"
	"umaru/application/setting"
	"umaru/application/tool"
)

// GetRssPath 获取qbittorrent添加rss时的path
//	@param title string 番剧标题
//	@return string rss的path
func GetRssPath(title string) string {
	return setting.QB_RSS_FOLDER + "\\" + title
}

// GetQbDownloadPath 获取qbittorrent下载路径
//	@param animeId int64 番剧数据库ID
//	@return string qbittorrent下载路径(文件夹)
func GetQbDownloadPath(animeId int64) string {
	return tool.PathJoin(setting.QB_DOWNLOAD_PATH, strconv.FormatInt(animeId, 10))
}
