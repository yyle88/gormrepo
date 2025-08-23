package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (*SaleRecord) Columns() *SaleRecordColumns {
	return &SaleRecordColumns{
		ID:          "id",
		CreatedAt:   "created_at",
		UpdatedAt:   "updated_at",
		DeletedAt:   "deleted_at",
		ProductName: "product_name",
		Category:    "category",
		Region:      "region",
		SaleDate:    "sale_date",
		Quantity:    "quantity",
		UnitPrice:   "unit_price",
		TotalAmount: "total_amount",
		SalesRep:    "sales_rep",
		Channel:     "channel",
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
