package gormrepo

import (
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

// Use returns the database(db), model(mod), and associated columns(cls)
// Convenient function to get all three components in one call
//
// Use 返回数据库(db)、模型(mod)和关联的列(cls)
// 便捷函数，一次调用获取三个组件
func Use[MOD gormclass.ModelCols[CLS], CLS any](db *gorm.DB, one MOD) (*gorm.DB, MOD, CLS) {
	one, cls := gormclass.Use(one)
	return db, one, cls
}

// Umc returns the database(db), model(mod), and associated columns(cls)
// Alias to Use, provides the same function with a shorter name
//
// Umc 返回数据库(db)、模型(mod)和关联的列(cls)
// 是 Use 的别名，以更短的名称提供相同功能
func Umc[MOD gormclass.ModelCols[CLS], CLS any](db *gorm.DB, one MOD) (*gorm.DB, MOD, CLS) {
	one, cls := gormclass.Umc(one)
	return db, one, cls
}
