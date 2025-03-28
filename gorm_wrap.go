package gormrepo

import "gorm.io/gorm"

type GormWrap[MOD any, CLS any] struct {
	db  *gorm.DB
	mod *MOD
	cls CLS
}

func NewGormWrap[MOD any, CLS any](db *gorm.DB, _ *MOD, cls CLS) *GormWrap[MOD, CLS] {
	return &GormWrap[MOD, CLS]{
		db:  db,
		mod: nil, // 这里就是设置个空值避免共享对象
		cls: cls,
	}
}

func (repo *GormWrap[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *MOD) *gorm.DB {
	return where(repo.db, repo.cls).First(dest)
}

func (repo *GormWrap[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *[]*MOD) *gorm.DB {
	return where(repo.db, repo.cls).Find(dest)
}

func (repo *GormWrap[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) *gorm.DB {
	column, value := valueFunc(repo.cls)
	return where(repo.db, repo.cls).Model((*MOD)(nil)).Update(column, value)
}

func (repo *GormWrap[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, valuesFunc func(cls CLS) map[string]interface{}) *gorm.DB {
	mp := valuesFunc(repo.cls)
	return where(repo.db, repo.cls).Model((*MOD)(nil)).Updates(mp)
}
