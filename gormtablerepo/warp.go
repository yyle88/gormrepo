package gormtablerepo

import (
	"github.com/yyle88/gormrepo"
	"gorm.io/gorm"
)

func (repo *TableRepo[MOD, CLS]) Base() *gormrepo.BaseRepo[MOD, CLS] {
	return gormrepo.NewBaseRepo((*MOD)(nil), repo.tbColumns)
}

func (repo *TableRepo[MOD, CLS]) Repo(db *gorm.DB) *gormrepo.GormRepo[MOD, CLS] {
	return repo.Base().Repo(db)
}

func (repo *TableRepo[MOD, CLS]) Gorm(db *gorm.DB) *gormrepo.GormWrap[MOD, CLS] {
	return repo.Base().Gorm(db)
}
