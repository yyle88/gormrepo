package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// 通过 gormcnm.gen_test.go 生成

func (*Post) Columns() *PostColumns {
	return &PostColumns{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Title:     "title",
		Content:   "content",
		AuthorID:  "author_id",
		Category:  "category",
		Tags:      "tags",
		ViewCount: "view_count",
		LikeCount: "like_count",
		Status:    "status",
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
