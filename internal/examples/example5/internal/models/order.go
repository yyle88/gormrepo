package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	GuestID     uint    `gorm:"index:idx_orders_guest_id_product_name"`
	ProductName string  `gorm:"index:idx_orders_guest_id_product_name"`
	Amount      int     `gorm:"not null"`
	Cost        float64 `gorm:"not null"`
	Address     string  `gorm:"not null"`
}

func (*Order) TableName() string {
	return "orders"
}
