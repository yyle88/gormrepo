package models

import (
	"time"

	"gorm.io/gorm"
)

// Product 产品模型
type Product struct {
	ID         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"size:200;not null;index"`
	Price      float64 `gorm:"type:decimal(10,2);not null"`
	Stock      int     `gorm:"default:0;index"`
	CategoryID uint    `gorm:"index"`
	Status     string  `gorm:"size:20;default:'active';index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (*Product) TableName() string {
	return "products"
}
