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

func (c *Product) Columns() *ProductColumns {
	return &ProductColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:         gormcnm.Cnm(c.ID, "id"),
		Name:       gormcnm.Cnm(c.Name, "name"),
		Price:      gormcnm.Cnm(c.Price, "price"),
		Stock:      gormcnm.Cnm(c.Stock, "stock"),
		CategoryID: gormcnm.Cnm(c.CategoryID, "category_id"),
		Status:     gormcnm.Cnm(c.Status, "status"),
		CreatedAt:  gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt:  gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt:  gormcnm.Cnm(c.DeletedAt, "deleted_at"),
	}
}

type ProductColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID         gormcnm.ColumnName[uint]
	Name       gormcnm.ColumnName[string]
	Price      gormcnm.ColumnName[float64]
	Stock      gormcnm.ColumnName[int]
	CategoryID gormcnm.ColumnName[uint]
	Status     gormcnm.ColumnName[string]
	CreatedAt  gormcnm.ColumnName[time.Time]
	UpdatedAt  gormcnm.ColumnName[time.Time]
	DeletedAt  gormcnm.ColumnName[gorm.DeletedAt]
}

func (c *Category) Columns() *CategoryColumns {
	return &CategoryColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(c.ID, "id"),
		Name:      gormcnm.Cnm(c.Name, "name"),
		Code:      gormcnm.Cnm(c.Code, "code"),
		Sort:      gormcnm.Cnm(c.Sort, "sort"),
		Status:    gormcnm.Cnm(c.Status, "status"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
	}
}

type CategoryColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	Name      gormcnm.ColumnName[string]
	Code      gormcnm.ColumnName[string]
	Sort      gormcnm.ColumnName[int]
	Status    gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
}
