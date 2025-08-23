package models

import (
	"time"

	"gorm.io/gorm"
)

type SaleRecord struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ProductName string    `gorm:"not null;index"`
	Category    string    `gorm:"index"`
	Region      string    `gorm:"index"`
	SaleDate    time.Time `gorm:"index"`
	Quantity    int       `gorm:"not null"`
	UnitPrice   float64   `gorm:"type:decimal(10,2);not null"`
	TotalAmount float64   `gorm:"type:decimal(10,2);not null"`
	SalesRep    string    `gorm:"index"`
	Channel     string    `gorm:"index"`
}

func (*SaleRecord) TableName() string {
	return "sale_records"
}
