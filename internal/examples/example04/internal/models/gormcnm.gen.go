package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// 这个文件通过 gormcnm.gen_test.go 生成

func (*Product) Columns() *ProductColumns {
	return &ProductColumns{
		ID:         "id",
		Name:       "name",
		Price:      "price",
		Stock:      "stock",
		CategoryID: "category_id",
		Status:     "status",
		CreatedAt:  "created_at",
		UpdatedAt:  "updated_at",
		DeletedAt:  "deleted_at",
	}
}

type ProductColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
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

func (*Category) Columns() *CategoryColumns {
	return &CategoryColumns{
		ID:        "id",
		Name:      "name",
		Code:      "code",
		Sort:      "sort",
		Status:    "status",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
	}
}

type CategoryColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	Name      gormcnm.ColumnName[string]
	Code      gormcnm.ColumnName[string]
	Sort      gormcnm.ColumnName[int]
	Status    gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
}
