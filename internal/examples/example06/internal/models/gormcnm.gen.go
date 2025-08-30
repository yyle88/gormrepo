package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Order) Columns() *OrderColumns {
	return &OrderColumns{
		ID:        gormcnm.Cnm(c.ID, "id"),
		CreatedAt: gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		OrderNo:   gormcnm.Cnm(c.OrderNo, "order_no"),
		UserID:    gormcnm.Cnm(c.UserID, "user_id"),
		Amount:    gormcnm.Cnm(c.Amount, "amount"),
		Status:    gormcnm.Cnm(c.Status, "status"),
		PayMethod: gormcnm.Cnm(c.PayMethod, "pay_method"),
		Remark:    gormcnm.Cnm(c.Remark, "remark"),
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
