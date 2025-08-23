package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型 - 演示 gormrepo 基础功能
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null;index"`
	Email     string `gorm:"size:100;uniqueIndex"`
	Age       int    `gorm:"default:0;index"`
	Status    string `gorm:"size:20;default:'active';index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (*User) TableName() string {
	return "users"
}
