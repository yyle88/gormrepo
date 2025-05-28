package gormrepo

import (
	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// ScopeFunction is a type alias for a function that modifies a GORM DB instance,
// used with db.Scopes() to apply custom query conditions.
// See: https://github.com/go-gorm/gorm/blob/c44405a25b0fb15c20265e672b8632b8774793ca/chainable_api.go#L376
type ScopeFunction = func(db *gorm.DB) *gorm.DB

// NewScope creates a GORM scope function that applies a custom where condition
// based on the provided CLS type and the repository's cls instance.
// Parameters:
//   - where: A function that takes a GORM DB instance and CLS type, returning a modified DB instance with applied conditions.
//
// Returns:
//   - A ScopeFunction that can be used with db.Scopes() to apply the where condition.
func (repo *Repo[MOD, CLS]) NewScope(where func(db *gorm.DB, cls CLS) *gorm.DB) ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return where(db, repo.cls)
	}
}

// Pagination defines parameters for paginated queries, including limit and offset.
// Fields:
//   - Limit: The maximum number of records to retrieve.
//   - Offset: The number of records to skip before retrieving.
type Pagination struct {
	Limit  int
	Offset int
}

// NewPaginateScope creates a GORM scope function that applies pagination and ordering
// to a query based on the provided orderBy function and pagination parameters.
// Parameters:
//   - orderBy: A function that takes a CLS instance and returns a gormcnm.OrderByBottle for query ordering.
//   - page: A Pagination struct specifying limit and offset for pagination.
//
// Returns:
//   - A ScopeFunction that can be used with db.Scopes() to apply ordering, limit, and offset.
func (repo *Repo[MOD, CLS]) NewPaginateScope(orderByFunc func(cls CLS) gormcnm.OrderByBottle, pageParam *Pagination) ScopeFunction {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(string(orderByFunc(repo.cls))).Limit(pageParam.Limit).Offset(pageParam.Offset)
	}
}
