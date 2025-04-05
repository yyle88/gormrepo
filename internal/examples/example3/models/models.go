package models

import "github.com/yyle88/gormcnm"

type User struct {
	ID   uint
	Name string
}

func (*User) TableName() string {
	return "users"
}

func (a *User) Columns() *UserColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *User) TableColumns(decoration gormcnm.ColumnNameDecoration) *UserColumns {
	return &UserColumns{
		ID:   gormcnm.Cmn(a.ID, "id", decoration),
		Name: gormcnm.Cmn(a.Name, "name", decoration),
	}
}

type UserColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

type Order struct {
	ID     uint
	UserID uint
	Amount float64
}

func (*Order) TableName() string {
	return "orders"
}

func (a *Order) Columns() *OrderColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		ID:     gormcnm.Cmn(a.ID, "id", decoration),
		UserID: gormcnm.Cmn(a.UserID, "user_id", decoration),
		Amount: gormcnm.Cmn(a.Amount, "amount", decoration),
	}
}

type OrderColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID     gormcnm.ColumnName[uint]
	UserID gormcnm.ColumnName[uint]
	Amount gormcnm.ColumnName[float64]
}
