package gormrepo

import (
	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

type GormRepo[MOD any, CLS any] struct {
	db  *gorm.DB
	cls CLS
}

func NewGormRepo[MOD any, CLS any](db *gorm.DB, _ *MOD, cls CLS) *GormRepo[MOD, CLS] {
	return &GormRepo[MOD, CLS]{
		db:  db,
		cls: cls,
	}
}

func (repo *GormRepo[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB) (*MOD, error) {
	var result = new(MOD)
	if err := repo.Gorm().First(where, result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *GormRepo[MOD, CLS]) FirstE(where func(db *gorm.DB, cls CLS) *gorm.DB) (*MOD, *ErrorOrNotExist) {
	var result = new(MOD)
	if err := repo.Gorm().First(where, result).Error; err != nil {
		return nil, NewErrorOrNotExist(err)
	}
	return result, nil
}

func (repo *GormRepo[MOD, CLS]) Where(where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return where(repo.db, repo.cls)
}

func (repo *GormRepo[MOD, CLS]) Exist(where func(db *gorm.DB, cls CLS) *gorm.DB) (bool, error) {
	var exists bool
	if err := where(repo.db, repo.cls).Model((*MOD)(nil)).Select("1").Limit(1).Find(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (repo *GormRepo[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB) ([]*MOD, error) {
	var results []*MOD
	if err := repo.Gorm().Find(where, &results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *GormRepo[MOD, CLS]) FindN(where func(db *gorm.DB, cls CLS) *gorm.DB, size int) ([]*MOD, error) {
	var results = make([]*MOD, 0, size)
	if err := where(repo.db, repo.cls).Limit(size).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *GormRepo[MOD, CLS]) FindC(where func(db *gorm.DB, cls CLS) *gorm.DB, paging func(db *gorm.DB, cls CLS) *gorm.DB) ([]*MOD, int64, error) {
	var results []*MOD
	{
		db := where(repo.db, repo.cls)
		db = paging(db, repo.cls)
		if err := db.Find(&results).Error; err != nil {
			return nil, 0, err
		}
	}
	var count int64
	{
		db := repo.db.Model((*MOD)(nil))
		if err := where(db, repo.cls).Count(&count).Error; err != nil {
			return nil, 0, err
		}
	}
	return results, count, nil
}

func (repo *GormRepo[MOD, CLS]) FindPageAndCount(where func(db *gorm.DB, cls CLS) *gorm.DB, ordering func(cls CLS) gormcnm.OrderByBottle, page *Pagination) ([]*MOD, int64, error) {
	results, err := repo.FindPage(where, ordering, page)
	if err != nil {
		return nil, 0, err
	}
	db := repo.db.Model((*MOD)(nil))
	var count int64
	if err := where(db, repo.cls).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return results, count, nil
}

func (repo *GormRepo[MOD, CLS]) FindPage(where func(db *gorm.DB, cls CLS) *gorm.DB, ordering func(cls CLS) gormcnm.OrderByBottle, page *Pagination) ([]*MOD, error) {
	db := where(repo.db, repo.cls)
	db = db.Order(string(ordering(repo.cls))) // gorm method just receives a few types, so we convert it to string.
	db = db.Limit(page.Limit).Offset(page.Offset)
	var results = make([]*MOD, 0, page.Limit)
	if err := db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *GormRepo[MOD, CLS]) Count(where func(db *gorm.DB, cls CLS) *gorm.DB) (int64, error) {
	var count int64
	if err := where(repo.db, repo.cls).Model((*MOD)(nil)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *GormRepo[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) error {
	if err := repo.Gorm().Update(where, valueFunc).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, mapValues func(cls CLS) map[string]interface{}) error {
	if err := repo.Gorm().Updates(where, mapValues).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) Invoke(clsRun func(db *gorm.DB, cls CLS) *gorm.DB) error {
	if err := clsRun(repo.db, repo.cls).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) Create(one *MOD) error {
	if err := repo.db.Create(one).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) Save(one *MOD) error {
	if err := repo.db.Save(one).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) Delete(one *MOD) error {
	// 使用 GORM Delete 时，Delete 的参数 one 不允许传 nil，因此 GORM 内部需要有效指针进行反射
	// When using GORM Delete, param one cannot be nil, as GORM requires valid instance for internal reflection
	if err := repo.db.Delete(one).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) DeleteW(where func(db *gorm.DB, cls CLS) *gorm.DB) error {
	// GORM Delete 需要有效指针，不能传 nil，否则报错 "invalid value, should be pointer to struct or slice"
	// GORM Delete needs valid instance, not nil, otherwise error "invalid value, should be instance to struct or slice"
	if err := where(repo.db, repo.cls).Delete(new(MOD)).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) DeleteM(one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB) error {
	// 使用 GORM Delete 时，Delete 的参数 one 不允许传 nil，因此 GORM 内部需要有效指针进行反射
	// When using GORM Delete, param one cannot be nil, as GORM requires valid instance for internal reflection
	if err := where(repo.db, repo.cls).Delete(one).Error; err != nil {
		return err
	}
	return nil
}
