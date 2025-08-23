package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (*User) Columns() *UserColumns {
	return &UserColumns{
		ID:        "id",
		Name:      "name",
		Email:     "email",
		Age:       "age",
		Status:    "status",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
	}
}

type UserColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	Name      gormcnm.ColumnName[string]
	Email     gormcnm.ColumnName[string]
	Age       gormcnm.ColumnName[int]
	Status    gormcnm.ColumnName[string]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
}
