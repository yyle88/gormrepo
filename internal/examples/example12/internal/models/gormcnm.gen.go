package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *SaleRecord) Columns() *SaleRecordColumns {
	return &SaleRecordColumns{
		ID:          gormcnm.Cnm(c.ID, "id"),
		CreatedAt:   gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt:   gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt:   gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		ProductName: gormcnm.Cnm(c.ProductName, "product_name"),
		Category:    gormcnm.Cnm(c.Category, "category"),
		Region:      gormcnm.Cnm(c.Region, "region"),
		SaleDate:    gormcnm.Cnm(c.SaleDate, "sale_date"),
		Quantity:    gormcnm.Cnm(c.Quantity, "quantity"),
		UnitPrice:   gormcnm.Cnm(c.UnitPrice, "unit_price"),
		TotalAmount: gormcnm.Cnm(c.TotalAmount, "total_amount"),
		SalesRep:    gormcnm.Cnm(c.SalesRep, "sales_rep"),
		Channel:     gormcnm.Cnm(c.Channel, "channel"),
	}
}

type SaleRecordColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID          gormcnm.ColumnName[uint]
	CreatedAt   gormcnm.ColumnName[time.Time]
	UpdatedAt   gormcnm.ColumnName[time.Time]
	DeletedAt   gormcnm.ColumnName[gorm.DeletedAt]
	ProductName gormcnm.ColumnName[string]
	Category    gormcnm.ColumnName[string]
	Region      gormcnm.ColumnName[string]
	SaleDate    gormcnm.ColumnName[time.Time]
	Quantity    gormcnm.ColumnName[int]
	UnitPrice   gormcnm.ColumnName[float64]
	TotalAmount gormcnm.ColumnName[float64]
	SalesRep    gormcnm.ColumnName[string]
	Channel     gormcnm.ColumnName[string]
}
