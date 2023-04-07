package usecase

import (
	"regexp"
	"strings"
)

var preReg = regexp.MustCompile("\\[\\s+")
var suffixReg = regexp.MustCompile("\\s+\\]")

// PreHandleTitle 预处理
//
// @params
//	title string 标题
// @return
//	string 处理后的标题
func PreHandleTitle(title string) string {
	// 替换所有中文中括号为英文中括号
	title = strings.ReplaceAll(title, "【", "[")
	title = strings.ReplaceAll(title, "】", "]")
	// 去除中括号内周围的空格
	title = preReg.ReplaceAllString(title, "[")
	title = suffixReg.ReplaceAllString(title, "]")
	return title
}
