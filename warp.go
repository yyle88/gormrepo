package gormrepo

import "context"

// Gorm converts a GormRepo to a GormWrap, sharing the same DB and CLS instances.
// Returns a new GormWrap instance for chainable operations.
func (repo *GormRepo[MOD, CLS]) Gorm() *GormWrap[MOD, CLS] {
	return NewGormWrap(repo.db, (*MOD)(nil), repo.cls)
}

// Repo converts a GormWrap to a GormRepo, sharing the same DB and CLS instances.
// Returns a new GormRepo instance for chainable operations.
func (wrap *GormWrap[MOD, CLS]) Repo() *GormRepo[MOD, CLS] {
	return NewGormRepo(wrap.db, (*MOD)(nil), wrap.cls)
}

// Mold sets the default struct template (MOD) for the GormRepo, setting it to the DB instance.
// Returns a new GormRepo instance for chainable operations.
func (repo *GormRepo[MOD, CLS]) Mold() *GormRepo[MOD, CLS] {
	return NewGormRepo(repo.db.Model((*MOD)(nil)), (*MOD)(nil), repo.cls)
}

// Mold sets the default struct template (MOD) for the GormWrap, setting it to the DB instance.
// Returns a new GormWrap instance for chainable operations.
func (wrap *GormWrap[MOD, CLS]) Mold() *GormWrap[MOD, CLS] {
	return NewGormWrap(wrap.db.Model((*MOD)(nil)), (*MOD)(nil), wrap.cls)
}

// WithContext sets the context for the GormRepo, applying it to the DB instance.
// Parameters:
//   - ctx: The context to set for database operations.
//
// Returns a new GormRepo instance for chainable operations.
func (repo *GormRepo[MOD, CLS]) WithContext(ctx context.Context) *GormRepo[MOD, CLS] {
	return NewGormRepo(repo.db.WithContext(ctx), (*MOD)(nil), repo.cls)
}

// WithContext sets the context for the GormWrap, setting it to the DB instance.
// Parameters:
//   - ctx: The context to set for database operations.
//
// Returns a new GormWrap instance for chainable operations.
func (wrap *GormWrap[MOD, CLS]) WithContext(ctx context.Context) *GormWrap[MOD, CLS] {
	return NewGormWrap(wrap.db.WithContext(ctx), (*MOD)(nil), wrap.cls)
}
