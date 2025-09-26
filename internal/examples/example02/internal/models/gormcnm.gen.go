// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: gormcnm.gen_test.go:43 -> models_test.TestGenerateColumns
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Account) Columns() *AccountColumns {
	return &AccountColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		Username:  gormcnm.Cnm(c.Username, "username"),
		Password:  gormcnm.Cnm(c.Password, "password"),
		Nickname:  gormcnm.Cnm(c.Nickname, "nickname"),
	}
}

type AccountColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Username  gormcnm.ColumnName[string]
	Password  gormcnm.ColumnName[string]
	Nickname  gormcnm.ColumnName[string]
}

func (c *Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		Name:      gormcnm.Cnm(c.Name, "name"),
		Age:       gormcnm.Cnm(c.Age, "age"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
	}
}

type ExampleColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Age       gormcnm.ColumnName[int]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
