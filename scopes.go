package gormrepo

import (
	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// ScopeFunction is a type alias representing a function that modifies a GORM DB instance
// Used with db.Scopes() to set custom conditions
// See: https://github.com/go-gorm/gorm/blob/c44405a25b0fb15c20265e672b8632b8774793ca/chainable_api.go#L376
//
// ScopeFunction 是表示修改 GORM DB 实例的函数的类型别名
// 与 db.Scopes() 配合使用以设置自定义条件
// 参见：https://github.com/go-gorm/gorm/blob/c44405a25b0fb15c20265e672b8632b8774793ca/chainable_api.go#L376
type ScopeFunction = func(db *gorm.DB) *gorm.DB

// NewScope creates a GORM scope function that applies a custom where condition
// Returns a ScopeFunction that can be used with db.Scopes()
//
// NewScope 创建应用自定义 where 条件的 GORM scope 函数
// 返回可与 db.Scopes() 配合使用的 ScopeFunction
func (repo *BaseRepo[MOD, CLS]) NewScope(where func(db *gorm.DB, cls CLS) *gorm.DB) ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return where(db, repo.cls)
	}
}

// NewWhereScope creates a GORM scope function that applies a custom where condition
// Alias to NewScope, providing a more explicit name when scoping where clauses
//
// NewWhereScope 创建应用自定义 where 条件的 GORM scope 函数
// 是 NewScope 的别名，在设置 where 子句作用域时提供更明确的名称
func (repo *BaseRepo[MOD, CLS]) NewWhereScope(where func(db *gorm.DB, cls CLS) *gorm.DB) ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return where(db, repo.cls)
	}
}

// NewOrderScope creates a GORM scope function that applies an ordering condition
// Returns a ScopeFunction that can be used with db.Scopes() to set the ordering
//
// NewOrderScope 创建应用排序条件的 GORM scope 函数
// 返回可与 db.Scopes() 配合使用以设置排序的 ScopeFunction
func (repo *BaseRepo[MOD, CLS]) NewOrderScope(ordering func(cls CLS) gormcnm.OrderByBottle) ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(string(ordering(repo.cls)))
	}
}

// Pagination defines parameters enabling paginated queries
// Contains limit (max records) and offset (records to skip)
//
// Pagination 定义分页查询参数
// 包含 limit（最大记录数）和 offset（跳过的记录数）
type Pagination struct {
	Limit  int // Max records to retrieve // 最大检索记录数
	Offset int // Records to skip // 跳过的记录数
}

// Scope creates a GORM scope function that applies pagination parameters
// Returns a ScopeFunction that sets limit and offset during pagination
//
// Scope 创建应用分页参数的 GORM scope 函数
// 返回在分页时设置 limit 和 offset 的 ScopeFunction
func (p *Pagination) Scope() ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return db.Limit(p.Limit).Offset(p.Offset)
	}
}

// NewPaginateScope creates a GORM scope function that applies pagination and ordering
// Combines ordering with limit and offset in a single scope
//
// NewPaginateScope 创建应用分页和排序的 GORM scope 函数
// 在单个 scope 中组合排序、limit 和 offset
func (repo *BaseRepo[MOD, CLS]) NewPaginateScope(ordering func(cls CLS) gormcnm.OrderByBottle, page *Pagination) ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(string(ordering(repo.cls))).Limit(page.Limit).Offset(page.Offset)
	}
}
