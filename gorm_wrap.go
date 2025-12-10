package gormrepo

import (
	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// GormWrap is a GORM database connection with column definitions
// Returns *gorm.DB from all operations that enables method chaining
// Provides more primitive access compared to GormRepo
//
// GormWrap 是带有列定义的 GORM 数据库连接封装
// 所有操作返回 *gorm.DB 以便方法链式调用
// 相比 GormRepo 提供更底层的访问
type GormWrap[MOD any, CLS any] struct {
	db  *gorm.DB // Database connection // 数据库连接
	cls CLS      // Column definitions // 列定义
}

// NewGormWrap creates a new GormWrap instance with database connection and column definitions
// The MOD param is used to deduce the type, its value is not used
//
// NewGormWrap 使用数据库连接和列定义创建新的 GormWrap 实例
// MOD 参数用于类型推断，其值不使用
func NewGormWrap[MOD any, CLS any](db *gorm.DB, _ *MOD, cls CLS) *GormWrap[MOD, CLS] {
	return &GormWrap[MOD, CLS]{
		db:  db,
		cls: cls,
	}
}

// First finds the first record matching the where condition
// Returns *gorm.DB for checking errors via .Error field
//
// First 查找符合 where 条件的第一条记录
// 返回 *gorm.DB 以便通过 .Error 字段检查错误
func (wrap *GormWrap[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *MOD) *gorm.DB {
	return where(wrap.db, wrap.cls).First(dest)
}

// Where applies the where condition and returns the gorm.DB to enable chaining
// Use when you need custom operations not provided by GormWrap
//
// Where 应用 where 条件并返回 gorm.DB 以便链式调用
// 当需要 GormWrap 未提供的自定义操作时使用
func (wrap *GormWrap[MOD, CLS]) Where(where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return where(wrap.db, wrap.cls)
}

// Find retrieves all records matching the where condition into dest slice
// Returns *gorm.DB for checking errors via .Error field
//
// Find 检索所有符合 where 条件的记录到 dest 切片
// 返回 *gorm.DB 以便通过 .Error 字段检查错误
func (wrap *GormWrap[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *[]*MOD) *gorm.DB {
	return where(wrap.db, wrap.cls).Find(dest)
}

// Update updates a single column for records matching the where condition
// Uses valueFunc to specify column name and value
//
// Update 更新符合 where 条件的记录的单个列
// 使用 valueFunc 指定列名和值
func (wrap *GormWrap[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) *gorm.DB {
	return where(wrap.db, wrap.cls).Model((*MOD)(nil)).Update(valueFunc(wrap.cls))
}

// Updates updates multiple columns for records matching the where condition
// Uses mapValues to specify column-value pairs
//
// Updates 更新符合 where 条件的记录的多个列
// 使用 mapValues 指定列值对
func (wrap *GormWrap[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, mapValues func(cls CLS) map[string]interface{}) *gorm.DB {
	return where(wrap.db, wrap.cls).Model((*MOD)(nil)).Updates(mapValues(wrap.cls))
}

// UpdatesM updates multiple columns using ColumnValueMap, provides fluent API without AsMap() call
// Usage: wrap.UpdatesM(where, func(cls CLS) gormcnm.ColumnValueMap { return cls.Kw(cls.Name.Kv("x"))... })
//
// UpdatesM 使用 ColumnValueMap 更新多列，提供流畅的 API，无需调用 AsMap()
// 用法：wrap.UpdatesM(where, func(cls CLS) gormcnm.ColumnValueMap { return cls.Kw(cls.Name.Kv("x"))... })
func (wrap *GormWrap[MOD, CLS]) UpdatesM(where func(db *gorm.DB, cls CLS) *gorm.DB, newValues func(cls CLS) gormcnm.ColumnValueMap) *gorm.DB {
	return wrap.Updates(where, func(cls CLS) map[string]interface{} {
		return newValues(cls).AsMap()
	})
}

// UpdatesO updates object using primary key as condition, with ColumnValueMap for values
// O = Object, the object must have valid primary key value, GORM uses it to locate the record
//
// UpdatesO 使用主键作为条件更新对象，使用 ColumnValueMap 指定更新值
// O = Object，object 必须有有效的主键值，GORM 会用它来定位要更新的记录
func (wrap *GormWrap[MOD, CLS]) UpdatesO(object *MOD, newValues func(cls CLS) gormcnm.ColumnValueMap) *gorm.DB {
	return wrap.db.Model(object).Updates(newValues(wrap.cls).AsMap())
}

// UpdatesC updates object using combined conditions: primary key from object plus where clause
// C = Combined, uses both object primary key and where conditions for precise targeting
//
// UpdatesC 使用组合条件更新对象：object 的主键加上 where 子句
// C = Combined，同时使用 object 主键和 where 条件进行精确定位
func (wrap *GormWrap[MOD, CLS]) UpdatesC(object *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB, newValues func(cls CLS) gormcnm.ColumnValueMap) *gorm.DB {
	return where(wrap.db.Model(object), wrap.cls).Updates(newValues(wrap.cls).AsMap())
}

// Invoke executes a custom operation using the database connection and column definitions
// Returns *gorm.DB to check errors and enable chaining
//
// Invoke 使用数据库连接和列定义执行自定义操作
// 返回 *gorm.DB 以便检查错误和链式调用
func (wrap *GormWrap[MOD, CLS]) Invoke(clsRun func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return clsRun(wrap.db, wrap.cls)
}

// Create inserts a new record into the database
// Returns *gorm.DB for checking errors via .Error field
//
// Create 向数据库插入一条新记录
// 返回 *gorm.DB 以便通过 .Error 字段检查错误
func (wrap *GormWrap[MOD, CLS]) Create(one *MOD) *gorm.DB {
	return wrap.db.Create(one)
}

// Save inserts or updates a record based on primary key
// If primary key is zero value, creates new record; otherwise updates existing
//
// Save 根据主键插入或更新记录
// 如果主键是零值，创建新记录；否则更新现有记录
func (wrap *GormWrap[MOD, CLS]) Save(one *MOD) *gorm.DB {
	return wrap.db.Save(one)
}

// Delete deletes the given record using its primary key
// When using GORM Delete, param one cannot be nil, as GORM requires valid instance
//
// Delete 使用主键删除给定记录
// 使用 GORM Delete 时，参数 one 不能为 nil，因为 GORM 需要有效实例
func (wrap *GormWrap[MOD, CLS]) Delete(one *MOD) *gorm.DB {
	return wrap.db.Delete(one)
}

// DeleteW deletes records matching the where condition
// W = Where, uses where condition instead of object primary key
//
// DeleteW 删除符合 where 条件的记录
// W = Where，使用 where 条件而不是对象主键
func (wrap *GormWrap[MOD, CLS]) DeleteW(where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	// GORM Delete needs valid instance, not nil, otherwise gets ErrInvalidValue
	// GORM Delete 需要有效实例，不能为 nil，否则报 ErrInvalidValue 错误
	return where(wrap.db, wrap.cls).Delete(new(MOD))
}

// DeleteM deletes the given object with additional where condition
// M = Model + Where, combines object with where condition
//
// DeleteM 删除给定对象并附加 where 条件
// M = Model + Where，组合对象和 where 条件
func (wrap *GormWrap[MOD, CLS]) DeleteM(one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB) *gorm.DB {
	return where(wrap.db, wrap.cls).Delete(one)
}
