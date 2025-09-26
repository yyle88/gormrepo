package gormrepo

import (
	"gorm.io/gorm"
)

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

func (wrap *GormWrap[MOD, CLS]) Create(one *MOD) *gorm.DB {
	return wrap.db.Create(one)
}

func (wrap *GormWrap[MOD, CLS]) Save(one *MOD) *gorm.DB {
	return wrap.db.Save(one)
}

func (wrap *GormWrap[MOD, CLS]) Delete(one *MOD) *gorm.DB {
	// 使用 GORM Delete 时，Delete 的参数 one 不允许传 nil，因此 GORM 内部需要有效指针进行反射
	// When using GORM Delete, param one cannot be nil, as GORM requires valid instance for internal reflection
	return wrap.db.Delete(one)
}

func (wrap *GormWrap[MOD, CLS]) DeleteW(where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	// GORM Delete 需要有效指针，不能传 nil，否则报错 "invalid value, should be pointer to struct or slice"
	// GORM Delete needs valid instance, not nil, otherwise error "invalid value, should be instance to struct or slice"
	return where(wrap.db, wrap.cls).Delete(new(MOD))
}

func (wrap *GormWrap[MOD, CLS]) DeleteM(one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	// 使用 GORM Delete 时，Delete 的参数 one 不允许传 nil，因此 GORM 内部需要有效指针进行反射
	// When using GORM Delete, param one cannot be nil, as GORM requires valid instance for internal reflection
	return where(wrap.db, wrap.cls).Delete(one)
}
