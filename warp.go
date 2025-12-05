package gormrepo

import "context"

// Gorm converts a GormRepo to a GormWrap, sharing the same DB and CLS instances
// Returns a new GormWrap instance enabling chainable operations
//
// Gorm 将 GormRepo 转换为 GormWrap，共享相同的 DB 和 CLS 实例
// 返回新的 GormWrap 实例以支持链式操作
func (repo *GormRepo[MOD, CLS]) Gorm() *GormWrap[MOD, CLS] {
	return NewGormWrap(repo.db, (*MOD)(nil), repo.cls)
}

// Repo converts a GormWrap to a GormRepo, sharing the same DB and CLS instances
// Returns a new GormRepo instance enabling chainable operations
//
// Repo 将 GormWrap 转换为 GormRepo，共享相同的 DB 和 CLS 实例
// 返回新的 GormRepo 实例以支持链式操作
func (wrap *GormWrap[MOD, CLS]) Repo() *GormRepo[MOD, CLS] {
	return NewGormRepo(wrap.db, (*MOD)(nil), wrap.cls)
}

// Mold sets the default model template (MOD) on the GormRepo DB instance
// Returns a new GormRepo instance enabling chainable operations
//
// Mold 在 GormRepo 的 DB 实例上设置默认模型模板 (MOD)
// 返回新的 GormRepo 实例以支持链式操作
func (repo *GormRepo[MOD, CLS]) Mold() *GormRepo[MOD, CLS] {
	return NewGormRepo(repo.db.Model((*MOD)(nil)), (*MOD)(nil), repo.cls)
}

// Mold sets the default model template (MOD) on the GormWrap DB instance
// Returns a new GormWrap instance enabling chainable operations
//
// Mold 在 GormWrap 的 DB 实例上设置默认模型模板 (MOD)
// 返回新的 GormWrap 实例以支持链式操作
func (wrap *GormWrap[MOD, CLS]) Mold() *GormWrap[MOD, CLS] {
	return NewGormWrap(wrap.db.Model((*MOD)(nil)), (*MOD)(nil), wrap.cls)
}

// WithContext sets the context on the GormRepo DB instance
// Returns a new GormRepo instance enabling chainable operations
//
// WithContext 在 GormRepo 的 DB 实例上设置上下文
// 返回新的 GormRepo 实例以支持链式操作
func (repo *GormRepo[MOD, CLS]) WithContext(ctx context.Context) *GormRepo[MOD, CLS] {
	return NewGormRepo(repo.db.WithContext(ctx), (*MOD)(nil), repo.cls)
}

// WithContext sets the context on the GormWrap DB instance
// Returns a new GormWrap instance enabling chainable operations
//
// WithContext 在 GormWrap 的 DB 实例上设置上下文
// 返回新的 GormWrap 实例以支持链式操作
func (wrap *GormWrap[MOD, CLS]) WithContext(ctx context.Context) *GormWrap[MOD, CLS] {
	return NewGormWrap(wrap.db.WithContext(ctx), (*MOD)(nil), wrap.cls)
}
