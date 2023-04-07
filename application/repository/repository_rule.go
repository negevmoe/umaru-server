package repository

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"umaru/application/model/dto"
)

func (r repositoryImpl) RuleSelectList(db IDB, req dto.RuleSelectListRequest) (res dto.RuleSelectListResponse, err error) {
	sql := `
select id,
       name,
       must_contain,
       must_not_contain,
       use_regex,
       episode_filter,
       smart_filter,
       create_time,
       update_time
from rule;
`
	err = db.Select(&res.Items, sql)
	if err != nil {
		log.Error("获取规则列表失败", zap.Error(err))
		return
	}
	return
}

func (r repositoryImpl) RuleSelectByName(db IDB, req dto.RuleSelectByNameRequest) (res dto.RuleSelectByNameResponse, err error) {
	sql := `
select id,
       name,
       must_contain,
       must_not_contain,
       use_regex,
       episode_filter,
       smart_filter,
       create_time,
       update_time
from rule where name = ? 
`
	queryx, err := db.Queryx(sql, req.Name)
	if err != nil {
		log.Error("获取规则失败", zap.Error(err))
		return
	}
	defer queryx.Close()

	if queryx.Next() {
		_ = queryx.StructScan(&res.Rule)
	}

	return
}

func (r repositoryImpl) RuleUpdate(db IDB, req dto.RuleUpdateRequest) (res dto.RuleUpdateResponse, err error) {
	sql := `
update rule
set name=?,
    must_contain=?,
    must_not_contain=?,
    use_regex=?,
    episode_filter=?,
    smart_filter=?,
    update_time=?
where id = ?;
`
	if _, err = db.Exec(sql,
		req.Rule.Name,
		req.Rule.MustContain,
		req.Rule.MustNotContain,
		req.Rule.UseRegex,
		req.Rule.EpisodeFilter,
		req.Rule.SmartFilter,
		req.Rule.UpdateTime,
		req.Rule.Id,
	); err != nil {
		log.Error("更新规则失败", zap.Error(err),
			zap.String("name", req.Rule.Name),
			zap.String("must_contain", req.Rule.MustContain),
			zap.String("must_not_contain", req.Rule.MustNotContain),
			zap.Int("use_regex", req.Rule.UseRegex),
			zap.String("episode_filter", req.Rule.EpisodeFilter),
			zap.Int("smart_filter", req.Rule.SmartFilter),
			zap.Int64("update_time", req.Rule.UpdateTime),
			zap.Int64("id", req.Rule.Id),
		)
		return
	}
	return
}

func (r repositoryImpl) RuleDeleteList(db IDB, req dto.RuleDeleteListRequest) (res dto.RuleDeleteListResponse, err error) {
	sql, values, err := sqlx.In(`delete from rule where id in (?)`, req.IdList)
	if err != nil {
		log.Error("删除规则失败", zap.Error(err))
		return
	}

	if _, err = db.Exec(sql, values...); err != nil {
		log.Error("删除规则失败", zap.Error(err))
		return
	}

	return
}

func (r repositoryImpl) RuleInsert(db IDB, req dto.RuleInsertRequest) (res dto.RuleInsertResponse, err error) {
	sql := `
insert into rule (name,
                  must_contain,
                  must_not_contain,
                  use_regex,
                  episode_filter,
                  smart_filter,
                  create_time,
                  update_time)
VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`
	if _, err = db.Exec(sql,
		req.Rule.Name,
		req.Rule.MustContain,
		req.Rule.MustNotContain,
		req.Rule.UseRegex,
		req.Rule.EpisodeFilter,
		req.Rule.SmartFilter,
		req.Rule.CreateTime,
		req.Rule.UpdateTime,
	); err != nil {
		log.Error("删除规则失败", zap.Error(err),
			zap.String("name", req.Rule.Name),
			zap.String("must_contain", req.Rule.MustContain),
			zap.String("must_not_contain", req.Rule.MustNotContain),
			zap.Int("use_regex", req.Rule.UseRegex),
			zap.String("episode_filter", req.Rule.EpisodeFilter),
			zap.Int("smart_filter", req.Rule.SmartFilter),
			zap.Int64("update_time", req.Rule.UpdateTime),
		)
		return
	}

	return
}
