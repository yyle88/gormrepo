package models

import (
	"time"

	"gorm.io/gorm"
)

// Category 分类模型
type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null;uniqueIndex"`
	Code      string `gorm:"size:50;uniqueIndex"`
	Sort      int    `gorm:"default:0;index"`
	Status    string `gorm:"size:20;default:'active';index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (*Category) TableName() string {
	return "categories"
}
