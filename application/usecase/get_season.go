package usecase

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
从rss标题中获取季信息
中文和罗马只支持 10 以内
*/

// GetSeason 获取季信息
//	数字: s4 S4 第4季 第4期
//	中文: 第四季 第四期
//	罗马: IV 第IV季 第IV期
// @params
//	title string 标题
// @return
//	int64 季
//	ok 是否成功提取
func GetSeason(title string) (res int64, ok bool) {
	title = PreHandleTitle(title)
	res, ok = findNumberSeason(title)
	if ok {
		return
	}
	res, ok = findZhSeason(title)
	if ok {
		return
	}
	res, ok = findRomeSeason(title)
	if ok {
		return
	}
	return 0, false
}

var seasonNumberReg = regexp.MustCompile("\\ss(\\d+)|S(\\d+)|第?(\\d+)[季期]")
var seasonRomeReg = regexp.MustCompile("\\s([IVX]+)|\\s第([IVX]+)[季期]")
var seasonZhReg = regexp.MustCompile("\\s第([\u4e00\u4e8c\u4e09\u56db\u4e94\u516d\u4e03\u516b\u4e5d\u5341]+)[季期]")

var seasonZhMap = map[string]int64{
	"一": 1,
	"二": 2,
	"三": 3,
	"四": 4,
	"五": 5,
	"六": 6,
	"七": 7,
	"八": 8,
	"九": 9,
	"十": 10,
}

var seasonRomeMap = map[string]int64{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

func findNumberSeason(title string) (res int64, ok bool) {
	match := seasonNumberReg.FindStringSubmatch(title)
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

func findZhSeason(title string) (res int64, ok bool) {
	match := seasonZhReg.FindStringSubmatch(title)
	if len(match) == 0 {
		return 0, false
	}

	fmt.Println(match)
	for i := 1; i < len(match); i++ {
		// 跳过没有匹配成功的 () 分组
		if match[i] == "" {
			continue
		}
		res, ok = seasonZhMap[match[i]]
		return
	}
	return 0, false
}

func findRomeSeason(title string) (res int64, ok bool) {
	match := seasonRomeReg.FindStringSubmatch(title)
	if len(match) == 0 {
		return 0, false
	}

	for i, v := range match {
		fmt.Println(i, v)
	}
	for i := 1; i < len(match); i++ {
		// 跳过没有匹配成功的 () 分组
		if match[i] == "" {
			continue
		}
		res, ok = seasonRomeMap[match[i]]
		return
	}
	return 0, false
}
