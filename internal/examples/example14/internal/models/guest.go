package models

import "gorm.io/gorm"

type Guest struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Nickname string
	Phone    string
	Email    string
}

func (*Guest) TableName() string {
	return "guests"
}
