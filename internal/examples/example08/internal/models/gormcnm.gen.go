// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: gormcnm.gen_test.go:42 -> models_test.TestGenerateColumns
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Post) Columns() *PostColumns {
	return &PostColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
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
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
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
