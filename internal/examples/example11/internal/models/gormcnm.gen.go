// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: gormcnm.gen_test.go:46 -> models_test.TestGenerateColumns
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import "github.com/yyle88/gormcnm"

func (c *User) Columns() *UserColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *User) TableColumns(decoration gormcnm.ColumnNameDecoration) *UserColumns {
	return &UserColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:   gormcnm.Cmn(c.ID, "id", decoration),
		Name: gormcnm.Cmn(c.Name, "name", decoration),
	}
}

type UserColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

func (c *Order) Columns() *OrderColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:     gormcnm.Cmn(c.ID, "id", decoration),
		UserID: gormcnm.Cmn(c.UserID, "user_id", decoration),
		Amount: gormcnm.Cmn(c.Amount, "amount", decoration),
	}
}

type OrderColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID     gormcnm.ColumnName[uint]
	UserID gormcnm.ColumnName[uint]
	Amount gormcnm.ColumnName[float64]
}
