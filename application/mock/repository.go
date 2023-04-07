package mock

import (
	"umaru/application/model/dao"
	"umaru/application/model/dto"
	"umaru/application/repository"
)

type Repository struct {
}

func (r Repository) AnimeSelect(db repository.IDB, id int64, bangumiId int64) (res dao.Anime, err error) {

	//TODO implement me
	panic("implement me")
}

func (r Repository) AnimeSelectByTitleAndSeason(db repository.IDB, title string, season int64) (res dao.Anime, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) AnimeSelectList(db repository.IDB, params dto.AnimeSelectListParams) (res []dao.Anime, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) AnimeInfoViewSelect(db repository.IDB, id int64, bangumiId int64) (res dao.AnimeInfoView, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) AnimeInfoViewSelectList(db repository.IDB, params dto.AnimeSelectListParams) (res []dao.AnimeInfoView, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) AnimeInsert(db repository.IDB, anime dao.Anime) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) AnimeDelete(db repository.IDB, id int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) AnimeUpdate(db repository.IDB, anime dao.Anime) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) RuleSelectList(db repository.IDB) (res []dao.Rule, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) RuleSelectByName(db repository.IDB, name string) (res dao.Rule, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) RuleUpdate(db repository.IDB, rule dao.Rule) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) RuleDeleteList(db repository.IDB, idList []int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) RuleInsert(db repository.IDB, rule dao.Rule) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CategorySelect(db repository.IDB, params dto.CategorySelectParams) (res dao.Category, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CategoryInsert(db repository.IDB, category dao.Category) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CategorySelectList(db repository.IDB) (res []dao.Category, err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CategoryDelete(db repository.IDB, id int64) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) CategoryUpdate(db repository.IDB, category dao.Category) (err error) {
	//TODO implement me
	panic("implement me")
}

func (r Repository) QBCategoryInsert() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) QBRuleSet() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) QBRuleDelete() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) QBRssInsert() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) QBRssDelete() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) QBLogSelectList() {
	//TODO implement me
	panic("implement me")
}

func (r Repository) QBTorrentInsertList() {
	//TODO implement me
	panic("implement me")
}
