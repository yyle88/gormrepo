package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Book) Columns() *BookColumns {
	return &BookColumns{
		ID:          gormcnm.Cnm(c.ID, "id"),
		CreatedAt:   gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt:   gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt:   gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		Title:       gormcnm.Cnm(c.Title, "title"),
		Author:      gormcnm.Cnm(c.Author, "author"),
		ISBN:        gormcnm.Cnm(c.ISBN, "isbn"),
		Price:       gormcnm.Cnm(c.Price, "price"),
		PublishYear: gormcnm.Cnm(c.PublishYear, "publish_year"),
		Rating:      gormcnm.Cnm(c.Rating, "rating"),
		Sales:       gormcnm.Cnm(c.Sales, "sales"),
		Category:    gormcnm.Cnm(c.Category, "category"),
		Status:      gormcnm.Cnm(c.Status, "status"),
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
