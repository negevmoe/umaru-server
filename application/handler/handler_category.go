package handler

import (
	"os"
	"time"
	"umaru-server/application/model/dao"
	"umaru-server/application/model/dto"
	"umaru-server/application/model/vo"
	"umaru-server/application/usecase"
)

func (s handlerImpl) CategoryCreate(req vo.CategoryCreateRequest) (res vo.CategoryCreateResponse, err error) {
	// 参数校验
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	// 检查分类名称是否存在
	categoryRet, err := repo.CategorySelect(db, dto.CategorySelectRequest{Name: req.Name})
	if err != nil {
		err = vo.ErrorWrap(500, "创建分类失败", err)
		return
	}

	if categoryRet.Category.Id > 0 {
		err = vo.ErrorNew(400, "分类名称已存在", req.Name+" 分类名称已存在")
		return
	}

	// 插入
	now := time.Now().Unix()
	_, err = repo.CategoryInsert(db, dto.CategoryInsertRequest{
		Category: dao.Category{
			Name:       req.Name,
			Origin:     2,
			CreateTime: now,
			UpdateTime: now,
		},
	})
	if err != nil {
		err = vo.ErrorWrap(500, "创建分类失败", err)
		return
	}
	return
}

func (s handlerImpl) CategoryGetList(req vo.CategoryGetListRequest) (res vo.CategoryGetListResponse, err error) {
	listRet, err := repo.CategorySelectList(db, dto.CategorySelectListRequest{})
	if err != nil {
		err = vo.ErrorWrap(500, "获取分类列表失败", err)
		return
	}

	res.Items = listRet.Items
	return
}

func (s handlerImpl) CategoryDelete(req vo.CategoryDeleteRequest) (res vo.CategoryDeleteResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}

	// 查询分类是否存在
	categorySelectRet, err := repo.CategorySelect(db, dto.CategorySelectRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "分类删除失败", err)
		return
	}
	if categorySelectRet.Category.Id == 0 {
		err = vo.ErrorNew(400, "分类删除失败", "分类不存在")
		return
	}
	if categorySelectRet.Category.Origin == 1 {
		err = vo.ErrorNew(400, "无法删除内置的分类", "无法删除内置的分类")
		return
	}
	// 获取原分类下的所有番剧
	animeListRet, err := repo.AnimeInfoViewSelectList(db, dto.AnimeInfoViewSelectListRequest{
		CategoryId: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "分类删除失败", err)
		return
	}

	// 获取原分类番剧的ID列表
	n := len(animeListRet.Items)
	animeIdList := make([]int64, 0, n)
	for _, item := range animeListRet.Items {
		animeIdList = append(animeIdList, item.Id)
	}
	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		err = vo.ErrorWrap(500, "分类删除失败", err)
		return
	}
	defer tx.Rollback()
	// 删除分类
	_, err = repo.CategoryDelete(tx, dto.CategoryDeleteRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "分类删除失败", err)
		return
	}
	if len(animeIdList) > 0 {
		// 更新番剧的分类ID为默认
		_, err = repo.AnimeCategoryUpdate(tx, dto.AnimeCategoryUpdateRequest{
			IdList:     animeIdList,
			CategoryId: 1,
		})
		if err != nil {
			err = vo.ErrorWrap(500, "分类删除失败", err)
			return
		}
	}

	// 移动视频文件夹到默认分类的文件夹
	LinkLock.Lock()
	err = os.RemoveAll(usecase.GetLinkCategoryDir(categorySelectRet.Category.Name))
	if err != nil {
		err = vo.ErrorWrap(500, "分类删除失败", err)
		return
	}
	LinkLock.Unlock()
	go Link()
	_ = tx.Commit()
	return
}

func (s handlerImpl) CategoryUpdate(req vo.CategoryUpdateRequest) (res vo.CategoryUpdateResponse, err error) {
	if msg, ok := req.ValidateError(); ok {
		err = vo.ErrorValidate(msg)
		return
	}
	// 检查分类是否存在
	exists, err := repo.CategorySelect(db, dto.CategorySelectRequest{
		Id: req.Id,
	})
	if err != nil {
		err = vo.ErrorWrap(500, "更新失败", err)
		return
	}
	category := exists.Category

	if category.Id == 0 {
		err = vo.ErrorNew(400, "分类不存在", "分类不存在")
		return
	}
	// 检查分类名是否重复
	exists, err = repo.CategorySelect(db, dto.CategorySelectRequest{
		Name: req.Name,
	})
	if exists.Category.Name == req.Name && exists.Category.Id != req.Id {
		err = vo.ErrorNew(400, "分类名称已存在", "分类名称已存在")
		return
	}

	// 开启事务
	tx, err := db.Beginx()
	if err != nil {
		err = vo.ErrorWrap(500, "分类删除失败", err)
		return
	}
	defer tx.Rollback()
	now := time.Now().Unix()

	// 更新分类
	_, err = repo.CategoryUpdate(tx, dto.CategoryUpdateRequest{
		Category: dao.Category{
			Id:         req.Id,
			Name:       req.Name,
			UpdateTime: now,
		},
	})
	if err != nil {
		err = vo.ErrorWrap(500, "分类更新失败", err)
		return
	}

	// 删除旧分类硬连接目录
	LinkLock.Lock()
	err = os.RemoveAll(usecase.GetLinkCategoryDir(category.Name))
	if err != nil {
		LinkLock.Unlock()
		err = vo.ErrorWrap(500, "分类删除失败", err)
		return
	}
	LinkLock.Unlock()
	// 硬连接
	go Link()

	_ = tx.Commit()
	return
}
