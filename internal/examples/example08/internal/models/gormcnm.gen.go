package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Post) Columns() *PostColumns {
	return &PostColumns{
		ID:        gormcnm.Cnm(c.ID, "id"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		Title:     gormcnm.Cnm(c.Title, "title"),
		Content:   gormcnm.Cnm(c.Content, "content"),
		AuthorID:  gormcnm.Cnm(c.AuthorID, "author_id"),
		Category:  gormcnm.Cnm(c.Category, "category"),
		Tags:      gormcnm.Cnm(c.Tags, "tags"),
		ViewCount: gormcnm.Cnm(c.ViewCount, "view_count"),
		LikeCount: gormcnm.Cnm(c.LikeCount, "like_count"),
		Status:    gormcnm.Cnm(c.Status, "status"),
	}
}

type PostColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Title     gormcnm.ColumnName[string]
	Content   gormcnm.ColumnName[string]
	AuthorID  gormcnm.ColumnName[uint]
	Category  gormcnm.ColumnName[string]
	Tags      gormcnm.ColumnName[string]
	ViewCount gormcnm.ColumnName[int]
	LikeCount gormcnm.ColumnName[int]
	Status    gormcnm.ColumnName[string]
}
