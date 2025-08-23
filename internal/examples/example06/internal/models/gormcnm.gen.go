package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

// 通过 gormcnm.gen_test.go 生成

func (*Order) Columns() *OrderColumns {
	return &OrderColumns{
		ID:        "id",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
		DeletedAt: "deleted_at",
		OrderNo:   "order_no",
		UserID:    "user_id",
		Amount:    "amount",
		Status:    "status",
		PayMethod: "pay_method",
		Remark:    "remark",
	}
}

type OrderColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	OrderNo   gormcnm.ColumnName[string]
	UserID    gormcnm.ColumnName[uint]
	Amount    gormcnm.ColumnName[float64]
	Status    gormcnm.ColumnName[string]
	PayMethod gormcnm.ColumnName[string]
	Remark    gormcnm.ColumnName[string]
}
