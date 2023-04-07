package handler

import (
	"errors"
	"go.uber.org/zap"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"umaru/application/model/dto"
	"umaru/application/usecase"
)

type mappingT struct {
	SourcePath string
	LinkPath   string
}

var LinkLock sync.Mutex // 硬连接锁

// Link 硬链接
//	1.查询anime 根据 分类,title,season 获取资源路径列表
//	2.遍历source_dir 对比资源路径列表: source_dir不存在的 警告, 存在的加入link列表
//  3.遍历link列表执行硬链接
func Link() {
	LinkLock.Lock()
	defer LinkLock.Unlock()

	var err error
	var mappingList = make([]mappingT, 0) // 硬连接映射列表

	/* 获取番剧列表 */
	ret, err := repo.AnimeInfoViewSelectList(db, dto.AnimeInfoViewSelectListRequest{})
	if err != nil {
		log.Error("硬连接失败,获取番剧信息失败", zap.Error(err))
		return
	}
	animeList := ret.Items

	if len(animeList) == 0 {
		return
	}

	/* 根据番剧列表获取全部的硬连接视频 */
	for _, item := range animeList {
		// 获取资源目录
		dir := usecase.GetSourceDir(item.Id)

		// 遍历资源目录
		_ = filepath.WalkDir(dir, func(sourcePath string, info fs.DirEntry, err error) error {

			// 处理读取路径失败
			if err != nil {
				log.Warn("获取资源文件失败,跳过该番剧",
					zap.Error(err),
					zap.String("title", item.Title),
					zap.Int64("season", item.Season),
				)
				return err
			}

			// 跳过目录
			if info.IsDir() {
				return nil
			}

			fullFilename := info.Name()                       // 文件名,带格式尾缀 filename.ext
			ext := path.Ext(fullFilename)                     // 格式尾缀 .ext
			filename := strings.TrimSuffix(fullFilename, ext) // 文件名,无格式尾缀 filename

			// 从文件名中提取集数
			episode, ok := usecase.GetEpisode(filename)
			if !ok {
				log.Warn("提取集数失败,硬链接跳过该视频", zap.String("filename", filename))
				return nil
			}

			// 获取硬链接路径
			linkPath, err := usecase.GetLinkPath(item.CategoryName, item.Title, item.Season, episode, ext)
			if err != nil {
				log.Warn("硬链接路径生成失败,跳过该视频", zap.Error(err),
					zap.String("media_path", item.Title),
					zap.String("category", item.Title),
					zap.String("title", item.Title),
					zap.Int64("season", item.Season),
					zap.Int64("episode", episode),
					zap.String("ext", ext),
				)
				return err
			}

			// 硬连接文件已存在 跳过
			if usecase.IsFileExists(linkPath) {
				log.Debug("跳过已存在的硬连接文件", zap.String("target", linkPath))
				return nil
			}

			// 添加到映射列表
			mappingList = append(mappingList, mappingT{
				SourcePath: sourcePath,
				LinkPath:   linkPath,
			})
			return nil
		})
	}

	/* 执行硬连接 */
	for _, item := range mappingList {
		// 如果不存在则创建硬连接目录
		dir := filepath.Dir(item.LinkPath)
		err = os.MkdirAll(dir, 0666)
		if err != nil {
			log.Error("硬连接失败,创建硬连接目标目录失败", zap.Error(err), zap.String("link_dir", dir))
			return
		}
		// 硬连接
		err = os.Link(item.SourcePath, item.LinkPath)
		if err != nil {
			// 目标文件已存在则跳过
			if errors.Is(err, syscall.ERROR_ALREADY_EXISTS) {
				log.Debug("硬连接失败,目标文件已存在", zap.Error(err), zap.String("target", item.LinkPath))
				continue
			}

			log.Error("硬连接失败", zap.Error(err),
				zap.String("source_path", item.SourcePath),
				zap.String("link_path", item.LinkPath),
			)
			return
		}

		log.Info("硬连接成功",
			zap.String("source_path", item.SourcePath),
			zap.String("link_path", item.LinkPath),
		)
	}

	return
}
