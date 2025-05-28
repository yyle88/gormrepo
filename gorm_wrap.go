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

func (wrap *GormWrap[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *MOD) *gorm.DB {
	return where(wrap.db, wrap.cls).First(dest)
}

func (wrap *GormWrap[MOD, CLS]) Where(where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return where(wrap.db, wrap.cls)
}

func (wrap *GormWrap[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *[]*MOD) *gorm.DB {
	return where(wrap.db, wrap.cls).Find(dest)
}

func (wrap *GormWrap[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) *gorm.DB {
	return where(wrap.db, wrap.cls).Model((*MOD)(nil)).Update(valueFunc(wrap.cls))
}

func (wrap *GormWrap[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, mapValues func(cls CLS) map[string]interface{}) *gorm.DB {
	return where(wrap.db, wrap.cls).Model((*MOD)(nil)).Updates(mapValues(wrap.cls))
}

func (wrap *GormWrap[MOD, CLS]) Invoke(clsRun func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return clsRun(wrap.db, wrap.cls)
}
