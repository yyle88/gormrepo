package example2models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (*Account) Columns() *AccountColumns {
	return &AccountColumns{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		Username:  "username",
		Password:  "password",
		Nickname:  "nickname",
	}
}

type AccountColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Username  gormcnm.ColumnName[string]
	Password  gormcnm.ColumnName[string]
	Nickname  gormcnm.ColumnName[string]
}

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        "id",
		Name:      "name",
		Age:       "age",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type ExampleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Age       gormcnm.ColumnName[int]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
