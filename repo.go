package gormrepo

import (
	"context"

	"gorm.io/gorm"
)

type Repo[MOD any, CLS any] struct {
	cls CLS
}

func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return &Repo[MOD, CLS]{
		cls: cls,
	}
}

func (repo *Repo[MOD, CLS]) Repo(db *gorm.DB) *GormRepo[MOD, CLS] {
	return NewGormRepo(db, (*MOD)(nil), repo.cls)
}

func (repo *Repo[MOD, CLS]) Gorm(db *gorm.DB) *GormWrap[MOD, CLS] {
	return NewGormWrap(db, (*MOD)(nil), repo.cls)
}

func (repo *Repo[MOD, CLS]) With(db *gorm.DB, ctx context.Context) *GormRepo[MOD, CLS] {
	return repo.Repo(db.WithContext(ctx))
}

func (repo *Repo[MOD, CLS]) Wrap(db *gorm.DB, ctx context.Context) *GormWrap[MOD, CLS] {
	return repo.Gorm(db.WithContext(ctx))
}
