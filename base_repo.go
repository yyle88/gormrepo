// Package gormrepo provides repository pattern implementation with GORM
// Offers type-safe generic repository wrappers and query builders
// Supports fluent API design with chainable operations
//
// gormrepo 提供基于 GORM 的仓储模式实现
// 提供类型安全的泛型仓储封装和查询构建器
// 支持流畅的 API 设计和链式操作
package gormrepo

import (
	"context"

	"gorm.io/gorm"
)

// BaseRepo is the base repository holding column definitions without database connection
// Provides factory methods to create GormRepo or GormWrap instances with database connection
// Uses generics to ensure type safety between model and column definitions
//
// BaseRepo 是持有列定义但不持有数据库连接的基础仓储
// 提供工厂方法来创建带数据库连接的 GormRepo 或 GormWrap 实例
// 使用泛型确保模型和列定义之间的类型安全
type BaseRepo[MOD any, CLS any] struct {
	cls CLS // Column definitions // 列定义
}

// NewBaseRepo creates a new BaseRepo instance with column definitions
// The model parameter is used to infer the type, actual value is not used
//
// NewBaseRepo 使用列定义创建新的 BaseRepo 实例
// model 参数用于类型推断，实际值不使用
func NewBaseRepo[MOD any, CLS any](_ *MOD, cls CLS) *BaseRepo[MOD, CLS] {
	return &BaseRepo[MOD, CLS]{
		cls: cls,
	}
}

// Repo creates a GormRepo instance with the given database connection
// Returns a repository with error-returning methods
//
// Repo 使用给定的数据库连接创建 GormRepo 实例
// 返回带有错误返回方法的仓储
func (repo *BaseRepo[MOD, CLS]) Repo(db *gorm.DB) *GormRepo[MOD, CLS] {
	return NewGormRepo(db, (*MOD)(nil), repo.cls)
}

// Gorm creates a GormWrap instance with the given database connection
// Returns a wrapper with *gorm.DB returning methods
//
// Gorm 使用给定的数据库连接创建 GormWrap 实例
// 返回带有 *gorm.DB 返回方法的封装器
func (repo *BaseRepo[MOD, CLS]) Gorm(db *gorm.DB) *GormWrap[MOD, CLS] {
	return NewGormWrap(db, (*MOD)(nil), repo.cls)
}

// With creates a GormRepo instance with context attached to database connection
// Convenient method combining context setting and repository creation
//
// With 创建带有上下文附加到数据库连接的 GormRepo 实例
// 组合上下文设置和仓储创建的便捷方法
func (repo *BaseRepo[MOD, CLS]) With(ctx context.Context, db *gorm.DB) *GormRepo[MOD, CLS] {
	return repo.Repo(db.WithContext(ctx))
}

// Wrap creates a GormWrap instance with context attached to database connection
// Convenient method combining context setting and wrapper creation
//
// Wrap 创建带有上下文附加到数据库连接的 GormWrap 实例
// 组合上下文设置和封装器创建的便捷方法
func (repo *BaseRepo[MOD, CLS]) Wrap(ctx context.Context, db *gorm.DB) *GormWrap[MOD, CLS] {
	return repo.Gorm(db.WithContext(ctx))
}
