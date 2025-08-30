package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Product) Columns() *ProductColumns {
	return &ProductColumns{
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

func (c *Category) Columns() *CategoryColumns {
	return &CategoryColumns{
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
