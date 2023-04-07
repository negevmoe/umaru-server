package usecase

import (
	"regexp"
	"strconv"
)

/*
从qb下载后的文件标题中获取集信息
只支持数字
*/

var episodeNumberReg = regexp.MustCompile("\\[(\\d+)\\]|\\s(\\d+)\\s|\\s第(\\d+)[话集]")

// GetEpisode 获取季信息
//	数字: [4] 04 4 第4话 第4集
// @params
//	title string 标题
// @return
//	int64 集
//	ok 是否成功提取
func GetEpisode(title string) (res int64, ok bool) {
	title = PreHandleTitle(title)
	match := episodeNumberReg.FindStringSubmatch(title)
	if len(match) == 0 {
		return 0, false
	}

	for i := 1; i < len(match); i++ {
		// 跳过没有匹配成功的 () 分组
		if match[i] == "" {
			continue
		}
		number, err := strconv.ParseInt(match[i], 10, 64)
		if err != nil {
			continue
		}

		return number, true
	}

	return 0, false
}
