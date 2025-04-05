package gormrepo

func (repo *GormRepo[MOD, CLS]) Gorm() *GormWrap[MOD, CLS] {
	return NewGormWrap(repo.db, (*MOD)(nil), repo.cls)
}

func (repo *GormWrap[MOD, CLS]) Morm() *GormWrap[MOD, CLS] {
	return NewGormWrap(repo.db.Model((*MOD)(nil)), (*MOD)(nil), repo.cls)
}

func (repo *GormWrap[MOD, CLS]) Repo() *GormRepo[MOD, CLS] {
	return NewGormRepo(repo.db, (*MOD)(nil), repo.cls)
}

func (repo *GormRepo[MOD, CLS]) Morm() *GormRepo[MOD, CLS] {
	return NewGormRepo(repo.db.Model((*MOD)(nil)), (*MOD)(nil), repo.cls)
}
