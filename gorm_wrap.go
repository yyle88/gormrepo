package gormrepo

import "gorm.io/gorm"

type GormWrap[MOD any, CLS any] struct {
	db  *gorm.DB
	cls CLS
}

func NewGormWrap[MOD any, CLS any](db *gorm.DB, _ *MOD, cls CLS) *GormWrap[MOD, CLS] {
	return &GormWrap[MOD, CLS]{
		db:  db,
		cls: cls,
	}
}

func (repo *GormWrap[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *MOD) *gorm.DB {
	return where(repo.db, repo.cls).First(dest)
}

func (repo *GormWrap[MOD, CLS]) Where(where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return where(repo.db, repo.cls)
}

func (repo *GormWrap[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *[]*MOD) *gorm.DB {
	return where(repo.db, repo.cls).Find(dest)
}

func (repo *GormWrap[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) *gorm.DB {
	return where(repo.db, repo.cls).Model((*MOD)(nil)).Update(valueFunc(repo.cls))
}

func (repo *GormWrap[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, mapValues func(cls CLS) map[string]interface{}) *gorm.DB {
	return where(repo.db, repo.cls).Model((*MOD)(nil)).Updates(mapValues(repo.cls))
}

func (repo *GormWrap[MOD, CLS]) Invoke(clsRun func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return clsRun(repo.db, repo.cls)
}
