package models

import (
	"gorm.io/gorm"
)

// Order 订单模型 - 使用 gorm.Model 继承，演示计数和存在性检查
type Order struct {
	gorm.Model
	OrderNo   string  `gorm:"size:50;uniqueIndex;not null"`
	UserID    uint    `gorm:"index;not null"`
	Amount    float64 `gorm:"type:decimal(10,2);not null"`
	Status    string  `gorm:"size:20;default:'pending';index"`
	PayMethod string  `gorm:"size:20;index"`
	Remark    string  `gorm:"size:500"`
}

func (*Order) TableName() string {
	return "orders"
}
