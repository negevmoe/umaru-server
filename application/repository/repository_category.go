package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"time"
	"umaru-server/application/model/dto"
)

func (r repositoryImpl) CategoryInsert(db IDB, req dto.CategoryInsertRequest) (res dto.CategoryInsertResponse, err error) {
	sql := `insert into category (name, origin, create_time, update_time) values (?,?,?,?)`
	_, err = db.Exec(sql, req.Category.Name, req.Category.Origin, req.Category.CreateTime, req.Category.UpdateTime)
	if err != nil {
		log.Error("添加分类失败", zap.Error(err),
			zap.String("name", req.Category.Name),
			zap.Int64("origin", req.Category.Origin),
			zap.Time("create_time", time.Unix(req.Category.CreateTime, 0)),
			zap.Time("update_time", time.Unix(req.Category.UpdateTime, 0)),
		)
		return
	}
	return
}

func (r repositoryImpl) CategorySelectList(db IDB, req dto.CategorySelectListRequest) (res dto.CategorySelectListResponse, err error) {
	var sql string
	var values []any
	if len(req.IdList) > 0 {
		sql, values, err = sqlx.In(`select id, name, origin, create_time, update_time from category where id in (?) `, req.IdList)
	} else {
		sql = `select id,name,origin,create_time,update_time from category where true `
		values = make([]any, 0)
	}

	if req.Origin > 0 {
		sql += ` and origin = ? `
		values = append(values, req.Origin)
	}

	err = db.Select(&res.Items, sql, values...)
	if err != nil {
		log.Error("查询分类失败", zap.Error(err))
		return
	}
	return
}

func (r repositoryImpl) CategoryDelete(db IDB, req dto.CategoryDeleteRequest) (res dto.CategoryDeleteResponse, err error) {
	sql := `delete from category where id = ? `
	_, err = db.Exec(sql, req.Id)
	if err != nil {
		log.Error("删除分类失败", zap.Error(err), zap.Int64("id", req.Id))
		return
	}
	return
}

func (r repositoryImpl) CategoryUpdate(db IDB, req dto.CategoryUpdateRequest) (res dto.CategoryUpdateResponse, err error) {
	sql := ` update category set name = ? where id = ?`
	_, err = db.Exec(sql, req.Category.Name, req.Category.Id)
	if err != nil {
		log.Error("更新分类失败", zap.Error(err), zap.Int64("id", req.Category.Id), zap.String("name", req.Category.Name))
		return
	}
	return
}

func (r repositoryImpl) CategorySelect(db IDB, req dto.CategorySelectRequest) (res dto.CategorySelectResponse, err error) {
	var values []any
	sql := ` select id, name, origin, create_time, update_time from category where true `

	if req.Id > 0 {
		sql += ` and id = ? `
		values = append(values, req.Id)
	}
	if req.Name != "" {
		sql += ` and name = ? `
		values = append(values, req.Name)
	}

	queryx, err := db.Queryx(sql, values...)
	if err != nil {
		log.Error("获取分类失败", zap.Error(err),
			zap.Int64("id", req.Id),
			zap.String("name", req.Name),
		)
		return
	}
	defer queryx.Close()

	if queryx.Next() {
		_ = queryx.StructScan(&res.Category)
	}
	return
}
