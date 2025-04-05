package gormtablerepo

import (
	"github.com/yyle88/gormrepo"
	"gorm.io/gorm"
)

type TableRepo[MOD any, CLS any] struct {
	tableName string
	tbColumns CLS
}

func NewTableRepo[MOD any, CLS any](_ *MOD, tableName string, tbColumns CLS) *TableRepo[MOD, CLS] {
	return &TableRepo[MOD, CLS]{
		tableName: tableName,
		tbColumns: tbColumns,
	}
}

func (repo *TableRepo[MOD, CLS]) GetTableName() string {
	return repo.tableName
}

func (repo *TableRepo[MOD, CLS]) TableColumns() CLS {
	return repo.tbColumns
}

func (repo *TableRepo[MOD, CLS]) Repo(db *gorm.DB) *gormrepo.GormRepo[MOD, CLS] {
	return gormrepo.NewRepo((*MOD)(nil), repo.tbColumns).Repo(db)
}

func (repo *TableRepo[MOD, CLS]) Gorm(db *gorm.DB) *gormrepo.GormWrap[MOD, CLS] {
	return gormrepo.NewRepo((*MOD)(nil), repo.tbColumns).Gorm(db)
}
