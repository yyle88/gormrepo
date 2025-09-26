// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: gormcnm.gen_test.go:44 -> models_test.TestGenerateColumns
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Author) Columns() *AuthorColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *Author) TableColumns(decoration gormcnm.ColumnNameDecoration) *AuthorColumns {
	return &AuthorColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:        gormcnm.Cmn(c.ID, "id", decoration),
		CreatedAt: gormcnm.Cmn(c.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(c.UpdatedAt, "updated_at", decoration),
		DeletedAt: gormcnm.Cmn(c.DeletedAt, "deleted_at", decoration),
		Name:      gormcnm.Cmn(c.Name, "name", decoration),
		Email:     gormcnm.Cmn(c.Email, "email", decoration),
		Bio:       gormcnm.Cmn(c.Bio, "bio", decoration),
		Country:   gormcnm.Cmn(c.Country, "country", decoration),
	}
}

type AuthorColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Name      gormcnm.ColumnName[string]
	Email     gormcnm.ColumnName[string]
	Bio       gormcnm.ColumnName[string]
	Country   gormcnm.ColumnName[string]
}

func (c *Book) Columns() *BookColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *Book) TableColumns(decoration gormcnm.ColumnNameDecoration) *BookColumns {
	return &BookColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:          gormcnm.Cmn(c.ID, "id", decoration),
		CreatedAt:   gormcnm.Cmn(c.CreatedAt, "created_at", decoration),
		UpdatedAt:   gormcnm.Cmn(c.UpdatedAt, "updated_at", decoration),
		DeletedAt:   gormcnm.Cmn(c.DeletedAt, "deleted_at", decoration),
		Title:       gormcnm.Cmn(c.Title, "title", decoration),
		ISBN:        gormcnm.Cmn(c.ISBN, "isbn", decoration),
		Price:       gormcnm.Cmn(c.Price, "price", decoration),
		PublishedAt: gormcnm.Cmn(c.PublishedAt, "published_at", decoration),
		Status:      gormcnm.Cmn(c.Status, "status", decoration),
		AuthorID:    gormcnm.Cmn(c.AuthorID, "author_id", decoration),
	}
}

type BookColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID          gormcnm.ColumnName[uint]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
	DeletedAt   gormcnm.ColumnName[gorm.DeletedAt]
	Title       gormcnm.ColumnName[string]
	ISBN        gormcnm.ColumnName[string]
	Price       gormcnm.ColumnName[float64]
	PublishedAt gormcnm.ColumnName[string]
	Status      gormcnm.ColumnName[string]
	AuthorID    gormcnm.ColumnName[uint]
}
