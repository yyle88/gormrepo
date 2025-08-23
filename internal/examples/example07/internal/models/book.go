package models

import "gorm.io/gorm"

// Book 图书模型 - 演示排序和限制查询
type Book struct {
	gorm.Model
	Title       string  `gorm:"size:200;not null;index"`
	Author      string  `gorm:"size:100;not null;index"`
	ISBN        string  `gorm:"size:20;uniqueIndex"`
	Price       float64 `gorm:"type:decimal(8,2);not null;index"`
	PublishYear int     `gorm:"index"`
	Rating      float32 `gorm:"type:decimal(3,1);index"`
	Sales       int     `gorm:"default:0;index"`
	Category    string  `gorm:"size:50;index"`
	Status      string  `gorm:"size:20;default:'available';index"`
}

func (*Book) TableName() string {
	return "books"
}
