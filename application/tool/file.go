package tool

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func PathJoin(p ...string) (res string) {
	arr := make([]string, 0, len(p))
	for index, item := range p {
		if index == 0 {
			if strings.HasPrefix(item, "/") {
				arr = append(arr, "")
			}
		}
		v := strings.TrimRight(item, "/")
		v = strings.TrimLeft(v, "/")
		arr = append(arr, v)
	}
	return strings.Join(arr, "/")
}

type IsExistsResult struct {
	Exist bool
	IsDir bool
}

func IsExists(path string) (ret IsExistsResult, err error) {
	stat, err := os.Stat(path)
	// 存在
	if err == nil {
		ret.Exist = true
		// 是目录
		if stat.IsDir() {
			ret.IsDir = true
		} else {
			ret.IsDir = false
		}
		return
	}
	// 如果是不存在错误
	if os.IsNotExist(err) {
		ret.Exist = false
		err = nil
		return
	}

	return
}

// RemoveEmptyDirAll 清空空目录
func RemoveEmptyDirAll(path string, deleteRoot bool) (ret []string, err error) {
	// 目录是否存在
	exists, err := IsExists(path)
	if err != nil {
		return
	}
	if !exists.Exist || !exists.IsDir {
		return
	}

	// 获取目录下的所有空目录
	checkDirList := make([]string, 0)
	err = filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			checkDirList = append(checkDirList, path)
		}
		return nil
	})
	if err != nil {
		return
	}

	// 判断是否删除参数目录
	endIndex := 1
	if deleteRoot {
		endIndex = 0
	}

	for i := len(checkDirList) - 1; i >= endIndex; i-- {
		dir, er := ioutil.ReadDir(checkDirList[i])
		if er != nil {
			err = er
			return
		}
		if len(dir) == 0 {
			if err = os.Remove(checkDirList[i]); err != nil {
				return
			}
			ret = append(ret, checkDirList[i])
		}
	}
	return
}
