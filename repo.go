package gormrepo

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepo[MOD any, CLS any] struct {
	cls CLS
}

func NewBaseRepo[MOD any, CLS any](_ *MOD, cls CLS) *BaseRepo[MOD, CLS] {
	return &BaseRepo[MOD, CLS]{
		cls: cls,
	}
}

func (repo *BaseRepo[MOD, CLS]) Repo(db *gorm.DB) *GormRepo[MOD, CLS] {
	return NewGormRepo(db, (*MOD)(nil), repo.cls)
}

func (repo *BaseRepo[MOD, CLS]) Gorm(db *gorm.DB) *GormWrap[MOD, CLS] {
	return NewGormWrap(db, (*MOD)(nil), repo.cls)
}

func (repo *BaseRepo[MOD, CLS]) With(ctx context.Context, db *gorm.DB) *GormRepo[MOD, CLS] {
	return repo.Repo(db.WithContext(ctx))
}

func (repo *BaseRepo[MOD, CLS]) Wrap(ctx context.Context, db *gorm.DB) *GormWrap[MOD, CLS] {
	return repo.Gorm(db.WithContext(ctx))
}
