package setting

import (
	"os"
	"strconv"
)

var (
	SERVER_DEBUG                 = getBool("SERVER_DEBUG", false)                 // DEBUG模式
	SERVER_PORT                  = getInt64("SERVER_PORT", 8001)                  // 服务端口
	SERVER_TOKEN_EXPIRATION_TIME = getInt64("SERVER_TOKEN_EXPIRATION_TIME", 3600) // token过期时间
	SERVER_USERNAME              = getString("SERVER_USERNAME", "admin")          // web 用户名
	SERVER_PASSWORD              = getString("SERVER_PASSWORD", "adminadmin")     // web 密码
	SERVER_LOG_DIR               = getString("SERVER_LOG_DIR", "/var/log/umaru")  // 日志存储路径
	SOURCE_PATH                  = getString("SOURCE_PATH", "/downloads")         // 资源目录 (下载目录)
	MEDIA_PATH                   = getString("MEDIA_PATH", "/media")              // 媒体目录 (硬链接目录)
	DB_PATH                      = getString("DB_PATH", "umaru.db")               // sqlite 路径
	DB_MAX_CONNS                 = getInt64("DB_MAX_CONNS", 30)                   // 最大连接数
	QB_URL                       = getString("QB_URL", "http://localhost:7999")   // qb的web url
	QB_USERNAME                  = getString("QB_USERNAME", "admin")              // qb的web 用户
	QB_PASSWORD                  = getString("QB_PASSWORD", "adminadmin")         // qb的web 密码
	QB_CATEGORY                  = getString("QB_CATEGORY", "umaru")              // qb中的下载分类
	QB_RSS_FOLDER                = getString("QB_RSS_FOLDER", "umaru")            // qb中的rss文件夹
	QB_DOWNLOAD_PATH             = getString("QB_DOWNLOAD_PATH", "/downloads")    // QB下载目录 (QB下载文件时指定的路径)
)

func getString(key string, value string) string {
	env, ok := os.LookupEnv(key)
	if !ok {
		return value
	}
	return env
}

func getInt64(key string, value int64) int64 {
	env, ok := os.LookupEnv(key)
	if !ok {
		return value
	}
	parseInt, err := strconv.ParseInt(env, 10, 64)
	if err != nil {
		return value
	}
	return parseInt
}

func getBool(key string, value bool) bool {
	env, ok := os.LookupEnv(key)
	if !ok {
		return value
	}

	parseBool, err := strconv.ParseBool(env)
	if err != nil {
		return value
	}
	return parseBool

}
