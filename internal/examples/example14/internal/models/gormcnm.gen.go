package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Guest) Columns() *GuestColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *Guest) TableColumns(decoration gormcnm.ColumnNameDecoration) *GuestColumns {
	return &GuestColumns{
		ID:        gormcnm.Cmn(c.ID, "id", decoration),
		CreatedAt: gormcnm.Cmn(c.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(c.UpdatedAt, "updated_at", decoration),
		DeletedAt: gormcnm.Cmn(c.DeletedAt, "deleted_at", decoration),
		Username:  gormcnm.Cmn(c.Username, "username", decoration),
		Nickname:  gormcnm.Cmn(c.Nickname, "nickname", decoration),
		Phone:     gormcnm.Cmn(c.Phone, "phone", decoration),
		Email:     gormcnm.Cmn(c.Email, "email", decoration),
	}
}

type GuestColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Username  gormcnm.ColumnName[string]
	Nickname  gormcnm.ColumnName[string]
	Phone     gormcnm.ColumnName[string]
	Email     gormcnm.ColumnName[string]
}

func (c *Order) Columns() *OrderColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		ID:          gormcnm.Cmn(c.ID, "id", decoration),
		CreatedAt:   gormcnm.Cmn(c.CreatedAt, "created_at", decoration),
		UpdatedAt:   gormcnm.Cmn(c.UpdatedAt, "updated_at", decoration),
		DeletedAt:   gormcnm.Cmn(c.DeletedAt, "deleted_at", decoration),
		GuestID:     gormcnm.Cmn(c.GuestID, "guest_id", decoration),
		ProductName: gormcnm.Cmn(c.ProductName, "product_name", decoration),
		Amount:      gormcnm.Cmn(c.Amount, "amount", decoration),
		Cost:        gormcnm.Cmn(c.Cost, "cost", decoration),
		Address:     gormcnm.Cmn(c.Address, "address", decoration),
	}
}

type OrderColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID          gormcnm.ColumnName[uint]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
	DeletedAt   gormcnm.ColumnName[gorm.DeletedAt]
	GuestID     gormcnm.ColumnName[uint]
	ProductName gormcnm.ColumnName[string]
	Amount      gormcnm.ColumnName[int]
	Cost        gormcnm.ColumnName[float64]
	Address     gormcnm.ColumnName[string]
}
