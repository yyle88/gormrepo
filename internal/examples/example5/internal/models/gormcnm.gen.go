package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (a *Guest) Columns() *GuestColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Guest) TableColumns(decoration gormcnm.ColumnNameDecoration) *GuestColumns {
	return &GuestColumns{
		ID:        gormcnm.Cmn(a.ID, "id", decoration),
		CreatedAt: gormcnm.Cmn(a.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(a.UpdatedAt, "updated_at", decoration),
		DeletedAt: gormcnm.Cmn(a.DeletedAt, "deleted_at", decoration),
		Username:  gormcnm.Cmn(a.Username, "username", decoration),
		Nickname:  gormcnm.Cmn(a.Nickname, "nickname", decoration),
		Phone:     gormcnm.Cmn(a.Phone, "phone", decoration),
		Email:     gormcnm.Cmn(a.Email, "email", decoration),
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

func (a *Order) Columns() *OrderColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		ID:          gormcnm.Cmn(a.ID, "id", decoration),
		CreatedAt:   gormcnm.Cmn(a.CreatedAt, "created_at", decoration),
		UpdatedAt:   gormcnm.Cmn(a.UpdatedAt, "updated_at", decoration),
		DeletedAt:   gormcnm.Cmn(a.DeletedAt, "deleted_at", decoration),
		GuestID:     gormcnm.Cmn(a.GuestID, "guest_id", decoration),
		ProductName: gormcnm.Cmn(a.ProductName, "product_name", decoration),
		Amount:      gormcnm.Cmn(a.Amount, "amount", decoration),
		Cost:        gormcnm.Cmn(a.Cost, "cost", decoration),
		Address:     gormcnm.Cmn(a.Address, "address", decoration),
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
