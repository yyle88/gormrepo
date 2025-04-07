package example2models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Username string `gorm:"unique;"`
	Password string `gorm:"size:255;"`
	Nickname string `gorm:"column:nickname;"`
}

func (*Account) TableName() string {
	return "accounts"
}
