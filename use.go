package gormrepo

import (
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

// Use returns the database(db) model (mod) and the associated columns (cls).
// Use 返回 数据库(db) 模型（mod）和关联的列（cls）。
func Use[MOD gormclass.GormClass[CLS], CLS any](db *gorm.DB, one MOD) (*gorm.DB, MOD, CLS) {
	one, cls := gormclass.Use(one)
	return db, one, cls
}

// Umc returns the database(db) model (mod) and the associated columns (cls), functioning the same as the Use function.
// Umc 返回 数据库(db) 模型（mod）和关联的列（cls），功能与 Use 函数相同。
func Umc[MOD gormclass.GormClass[CLS], CLS any](db *gorm.DB, one MOD) (*gorm.DB, MOD, CLS) {
	one, cls := gormclass.Umc(one)
	return db, one, cls
}
