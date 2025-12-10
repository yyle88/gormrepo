// Package gormrepo implements repositories pattern with GORM
// Offers type-safe generic repositories and chainable operations
// Supports fluent API design with method chaining
//
// gormrepo 基于 GORM 实现仓储模式
// 提供类型安全的泛型仓储和链式操作
// 支持流畅的 API 设计和方法链
package gormrepo

import (
	"context"

	"gorm.io/gorm"
)

// BaseRepo is the base repo holding CLS definitions without database connection
// Provides methods to create GormRepo/GormWrap instances with database connection
// Uses generics to ensure type-safe MOD and CLS definitions
//
// BaseRepo 是持有 CLS 定义但不持有数据库连接的基础仓储
// 提供方法来创建带数据库连接的 GormRepo/GormWrap 实例
// 使用泛型确保 MOD 和 CLS 定义的类型安全
type BaseRepo[MOD any, CLS any] struct {
	cls CLS // Column definitions // 列定义
}

// NewBaseRepo creates a new BaseRepo instance with CLS definitions
// The MOD param is used to deduce the type, its value is not used
//
// NewBaseRepo 使用 CLS 定义创建新的 BaseRepo 实例
// MOD 参数用于类型推断，其值不使用
func NewBaseRepo[MOD any, CLS any](_ *MOD, cls CLS) *BaseRepo[MOD, CLS] {
	return &BaseRepo[MOD, CLS]{
		cls: cls,
	}
}

// Repo creates a GormRepo instance with the given database connection
// GormRepo methods have (T, error) signatures
//
// Repo 使用给定的数据库连接创建 GormRepo 实例
// GormRepo 方法返回 (T, error) 签名
func (repo *BaseRepo[MOD, CLS]) Repo(db *gorm.DB) *GormRepo[MOD, CLS] {
	return NewGormRepo(db, (*MOD)(nil), repo.cls)
}

// Gorm creates a GormWrap instance with the given database connection
// GormWrap methods have *gorm.DB signatures
//
// Gorm 使用给定的数据库连接创建 GormWrap 实例
// GormWrap 方法返回 *gorm.DB 签名
func (repo *BaseRepo[MOD, CLS]) Gorm(db *gorm.DB) *GormWrap[MOD, CLS] {
	return NewGormWrap(db, (*MOD)(nil), repo.cls)
}

// With creates a GormRepo instance with context attached to database connection
// Combines context setting and GormRepo creation in one call
//
// With 创建带有上下文附加到数据库连接的 GormRepo 实例
// 在一次调用中组合上下文设置和 GormRepo 创建
func (repo *BaseRepo[MOD, CLS]) With(ctx context.Context, db *gorm.DB) *GormRepo[MOD, CLS] {
	return repo.Repo(db.WithContext(ctx))
}

// Wrap creates a GormWrap instance with context attached to database connection
// Combines context setting and GormWrap creation in one call
//
// Wrap 创建带有上下文附加到数据库连接的 GormWrap 实例
// 在一次调用中组合上下文设置和 GormWrap 创建
func (repo *BaseRepo[MOD, CLS]) Wrap(ctx context.Context, db *gorm.DB) *GormWrap[MOD, CLS] {
	return repo.Gorm(db.WithContext(ctx))
}
