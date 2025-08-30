package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Article) Columns() *ArticleColumns {
	return &ArticleColumns{
		ID:          gormcnm.Cnm(c.ID, "id"),
		Title:       gormcnm.Cnm(c.Title, "title"),
		Content:     gormcnm.Cnm(c.Content, "content"),
		AuthorID:    gormcnm.Cnm(c.AuthorID, "author_id"),
		Status:      gormcnm.Cnm(c.Status, "status"),
		ViewCount:   gormcnm.Cnm(c.ViewCount, "view_count"),
		LikeCount:   gormcnm.Cnm(c.LikeCount, "like_count"),
		PublishedAt: gormcnm.Cnm(c.PublishedAt, "published_at"),
		CreatedAt:   gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt:   gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt:   gormcnm.Cnm(c.DeletedAt, "deleted_at"),
	}
}

type ArticleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID          gormcnm.ColumnName[uint]
	Title       gormcnm.ColumnName[string]
	Content     gormcnm.ColumnName[string]
	AuthorID    gormcnm.ColumnName[uint]
	Status      gormcnm.ColumnName[string]
	ViewCount   gormcnm.ColumnName[int]
	LikeCount   gormcnm.ColumnName[int]
	PublishedAt gormcnm.ColumnName[*time.Time]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
	DeletedAt   gormcnm.ColumnName[gorm.DeletedAt]
}
