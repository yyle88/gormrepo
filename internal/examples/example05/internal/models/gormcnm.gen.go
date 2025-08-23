package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// 通过 gormcnm.gen_test.go 生成

func (*Article) Columns() *ArticleColumns {
	return &ArticleColumns{
		ID:          "id",
		Title:       "title",
		Content:     "content",
		AuthorID:    "author_id",
		Status:      "status",
		ViewCount:   "view_count",
		LikeCount:   "like_count",
		PublishedAt: "published_at",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
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
