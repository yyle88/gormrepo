package gormrepo

import (
	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// GormRepo is the main repo struct with database connection and column definitions
// Provides CRUD operations with (T, error) return signatures
// All methods accept type-safe where functions using column definitions
//
// GormRepo 是带有数据库连接和列定义的主仓储结构体
// 提供返回 (T, error) 签名的 CRUD 操作
// 所有方法接受使用列定义的类型安全 where 函数
type GormRepo[MOD any, CLS any] struct {
	db  *gorm.DB // Database connection // 数据库连接
	cls CLS      // Column definitions // 列定义
}

// NewGormRepo creates a new GormRepo instance with database connection and column definitions
// The MOD param is used to deduce the type, its value is not used
//
// NewGormRepo 使用数据库连接和列定义创建新的 GormRepo 实例
// MOD 参数用于类型推断，其值不使用
func NewGormRepo[MOD any, CLS any](db *gorm.DB, _ *MOD, cls CLS) *GormRepo[MOD, CLS] {
	return &GormRepo[MOD, CLS]{
		db:  db,
		cls: cls,
	}
}

// First finds the first record matching the where condition
// Returns the found record or error if not found or query fails
//
// First 查找符合 where 条件的第一条记录
// 返回找到的记录，如果未找到或查询失败则返回错误
func (repo *GormRepo[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB) (*MOD, error) {
	var result = new(MOD)
	if err := repo.Gorm().First(where, result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// FirstE finds the first record with structured error handling
// Returns ErrorOrNotExist to distinguish between not found and other errors
//
// FirstE 查找第一条记录，带有结构化错误处理
// 返回 ErrorOrNotExist 以区分记录未找到和其他错误
func (repo *GormRepo[MOD, CLS]) FirstE(where func(db *gorm.DB, cls CLS) *gorm.DB) (*MOD, *ErrorOrNotExist) {
	var result = new(MOD)
	if err := repo.Gorm().First(where, result).Error; err != nil {
		return nil, NewErrorOrNotExist(err)
	}
	return result, nil
}

// Where applies the where condition and returns the gorm.DB to enable chaining
// Use when you need custom operations not provided by GormRepo
//
// Where 应用 where 条件并返回 gorm.DB 以便链式调用
// 当需要 GormRepo 未提供的自定义操作时使用
func (repo *GormRepo[MOD, CLS]) Where(where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return where(repo.db, repo.cls)
}

// Exist checks if any record matches the where condition
// Returns true if at least one record exists, false otherwise
//
// Exist 检查是否存在符合 where 条件的记录
// 如果至少存在一条记录则返回 true，否则返回 false
func (repo *GormRepo[MOD, CLS]) Exist(where func(db *gorm.DB, cls CLS) *gorm.DB) (bool, error) {
	var exists bool
	if err := where(repo.db, repo.cls).Model((*MOD)(nil)).Select("1").Limit(1).Find(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

// Find retrieves all records matching the where condition
// Returns slice of records or error if query fails
//
// Find 检索所有符合 where 条件的记录
// 返回记录切片，如果查询失败则返回错误
func (repo *GormRepo[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB) ([]*MOD, error) {
	var results []*MOD
	if err := repo.Gorm().Find(where, &results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// FindN retrieves records matching the where condition with size limit
// Returns up to size records
//
// FindN 检索有限数量的符合 where 条件的记录
// 最多返回 size 条记录
func (repo *GormRepo[MOD, CLS]) FindN(where func(db *gorm.DB, cls CLS) *gorm.DB, size int) ([]*MOD, error) {
	var results = make([]*MOD, 0, size)
	if err := where(repo.db, repo.cls).Limit(size).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// FindC retrieves records with custom paging and returns total count
// Executes two queries: one with paging, one without to get total count
//
// FindC 使用自定义分页检索记录并返回总数
// 执行两个查询：一个带分页，一个不带以获取总数
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

// FindPageAndCount retrieves paginated records with ordering and returns total count
// Combines pagination with count query in single method call
//
// FindPageAndCount 使用排序检索分页记录并返回总数
// 在单个方法调用中组合分页和计数查询
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

// FindPage retrieves paginated records with ordering
// Uses Pagination struct to specify offset and limit
//
// FindPage 使用排序检索分页记录
// 使用 Pagination 结构体指定偏移和限制
func (repo *GormRepo[MOD, CLS]) FindPage(where func(db *gorm.DB, cls CLS) *gorm.DB, ordering func(cls CLS) gormcnm.OrderByBottle, page *Pagination) ([]*MOD, error) {
	db := where(repo.db, repo.cls)
	// GORM method just receives a few types, so we convert it to string
	// GORM 方法只接受几种类型，因此我们将其转换为字符串
	db = db.Order(string(ordering(repo.cls)))
	db = db.Limit(page.Limit).Offset(page.Offset)
	var results = make([]*MOD, 0, page.Limit)
	if err := db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// Count returns the number of records matching the where condition
//
// Count 返回符合 where 条件的记录数量
func (repo *GormRepo[MOD, CLS]) Count(where func(db *gorm.DB, cls CLS) *gorm.DB) (int64, error) {
	var count int64
	if err := where(repo.db, repo.cls).Model((*MOD)(nil)).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// Update updates a single column for records matching the where condition
// Uses valueFunc to specify column name and value
//
// Update 更新符合 where 条件的记录的单个列
// 使用 valueFunc 指定列名和值
func (repo *GormRepo[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) error {
	if err := repo.Gorm().Update(where, valueFunc).Error; err != nil {
		return err
	}
	return nil
}

// Updates updates multiple columns for records matching the where condition
// Uses mapValues to specify column-value pairs
//
// Updates 更新符合 where 条件的记录的多个列
// 使用 mapValues 指定列值对
func (repo *GormRepo[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, mapValues func(cls CLS) map[string]interface{}) error {
	if err := repo.Gorm().Updates(where, mapValues).Error; err != nil {
		return err
	}
	return nil
}

// UpdatesM updates multiple columns using ColumnValueMap, provides fluent API without AsMap() call
// Usage: repo.UpdatesM(where, func(cls CLS) gormcnm.ColumnValueMap { return cls.Kw(cls.Name.Kv("x"))... })
//
// UpdatesM 使用 ColumnValueMap 更新多列，提供流畅的 API，无需调用 AsMap()
// 用法：repo.UpdatesM(where, func(cls CLS) gormcnm.ColumnValueMap { return cls.Kw(cls.Name.Kv("x"))... })
func (repo *GormRepo[MOD, CLS]) UpdatesM(where func(db *gorm.DB, cls CLS) *gorm.DB, newValues func(cls CLS) gormcnm.ColumnValueMap) error {
	return repo.Updates(where, func(cls CLS) map[string]interface{} {
		return newValues(cls).AsMap()
	})
}

// UpdatesO updates object using primary key as condition, with ColumnValueMap for values
// O = Object, the object must have valid primary key value, GORM uses it to locate the record
//
// UpdatesO 使用主键作为条件更新对象，使用 ColumnValueMap 指定更新值
// O = Object，object 必须有有效的主键值，GORM 会用它来定位要更新的记录
func (repo *GormRepo[MOD, CLS]) UpdatesO(object *MOD, newValues func(cls CLS) gormcnm.ColumnValueMap) error {
	if err := repo.db.Model(object).Updates(newValues(repo.cls).AsMap()).Error; err != nil {
		return err
	}
	return nil
}

// UpdatesC updates object using combined conditions: primary key from object plus where clause
// C = Combined, uses both object primary key and where conditions for precise targeting
//
// UpdatesC 使用组合条件更新对象：object 的主键加上 where 子句
// C = Combined，同时使用 object 主键和 where 条件进行精确定位
func (repo *GormRepo[MOD, CLS]) UpdatesC(object *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB, newValues func(cls CLS) gormcnm.ColumnValueMap) error {
	if err := where(repo.db.Model(object), repo.cls).Updates(newValues(repo.cls).AsMap()).Error; err != nil {
		return err
	}
	return nil
}

// Invoke executes a custom operation using the database connection and column definitions
// Returns the custom operation's resulting error
//
// Invoke 使用数据库连接和列定义执行自定义操作
// 返回自定义操作产生的错误
func (repo *GormRepo[MOD, CLS]) Invoke(clsRun func(db *gorm.DB, cls CLS) *gorm.DB) error {
	if err := clsRun(repo.db, repo.cls).Error; err != nil {
		return err
	}
	return nil
}

// Create inserts a new record into the database
// The record's primary key will be populated after creation
//
// Create 向数据库插入一条新记录
// 创建后记录的主键将被填充
func (repo *GormRepo[MOD, CLS]) Create(one *MOD) error {
	if err := repo.db.Create(one).Error; err != nil {
		return err
	}
	return nil
}

// Save inserts or updates a record based on primary key
// If primary key is zero value, creates new record; otherwise updates existing
//
// Save 根据主键插入或更新记录
// 如果主键是零值，创建新记录；否则更新现有记录
func (repo *GormRepo[MOD, CLS]) Save(one *MOD) error {
	if err := repo.db.Save(one).Error; err != nil {
		return err
	}
	return nil
}

// Delete deletes the given record using its primary key
// When using GORM Delete, param one cannot be nil, as GORM requires valid instance
//
// Delete 使用主键删除给定记录
// 使用 GORM Delete 时，参数 one 不能为 nil，因为 GORM 需要有效实例
func (repo *GormRepo[MOD, CLS]) Delete(one *MOD) error {
	if err := repo.db.Delete(one).Error; err != nil {
		return err
	}
	return nil
}

// DeleteW deletes records matching the where condition
// W = Where, uses where condition instead of object primary key
//
// DeleteW 删除符合 where 条件的记录
// W = Where，使用 where 条件而不是对象主键
func (repo *GormRepo[MOD, CLS]) DeleteW(where func(db *gorm.DB, cls CLS) *gorm.DB) error {
	// GORM Delete needs valid instance, not nil, otherwise gets ErrInvalidValue
	// GORM Delete 需要有效实例，不能为 nil，否则报 ErrInvalidValue 错误
	if err := where(repo.db, repo.cls).Delete(new(MOD)).Error; err != nil {
		return err
	}
	return nil
}

// DeleteM deletes the given object with additional where condition
// M = Model + Where, combines object with where condition
//
// DeleteM 删除给定对象并附加 where 条件
// M = Model + Where，组合对象和 where 条件
func (repo *GormRepo[MOD, CLS]) DeleteM(one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB) error {
	if err := where(repo.db, repo.cls).Delete(one).Error; err != nil {
		return err
	}
	return nil
}
