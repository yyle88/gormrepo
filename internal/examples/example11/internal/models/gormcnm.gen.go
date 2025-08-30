package models

import "github.com/yyle88/gormcnm"

func (c *User) Columns() *UserColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *User) TableColumns(decoration gormcnm.ColumnNameDecoration) *UserColumns {
	return &UserColumns{
		ID:   gormcnm.Cmn(c.ID, "id", decoration),
		Name: gormcnm.Cmn(c.Name, "name", decoration),
	}
}

type UserColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

func (c *Order) Columns() *OrderColumns {
	return c.TableColumns(gormcnm.NewPlainDecoration())
}

func (c *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		ID:     gormcnm.Cmn(c.ID, "id", decoration),
		UserID: gormcnm.Cmn(c.UserID, "user_id", decoration),
		Amount: gormcnm.Cmn(c.Amount, "amount", decoration),
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
