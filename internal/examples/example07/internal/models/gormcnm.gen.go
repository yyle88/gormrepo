package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// 通过 gormcnm.gen_test.go 生成

func (*Book) Columns() *BookColumns {
	return &BookColumns{
		ID:          "id",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
		Title:       "title",
		Author:      "author",
		ISBN:        "isbn",
		Price:       "price",
		PublishYear: "publish_year",
		Rating:      "rating",
		Sales:       "sales",
		Category:    "category",
		Status:      "status",
	}
}

type BookColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID          gormcnm.ColumnName[uint]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
	DeletedAt   gormcnm.ColumnName[gorm.DeletedAt]
	Title       gormcnm.ColumnName[string]
	Author      gormcnm.ColumnName[string]
	ISBN        gormcnm.ColumnName[string]
	Price       gormcnm.ColumnName[float64]
	PublishYear gormcnm.ColumnName[int]
	Rating      gormcnm.ColumnName[float32]
	Sales       gormcnm.ColumnName[int]
	Category    gormcnm.ColumnName[string]
	Status      gormcnm.ColumnName[string]
}
