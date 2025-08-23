package models

import (
	"gorm.io/gorm"
)

// Author 作者表
type Author struct {
	gorm.Model

	Name    string `gorm:"not null;index"`
	Email   string `gorm:"unique;index"`
	Bio     string `gorm:"type:text"`
	Country string `gorm:"index"`
}

func (*Author) TableName() string {
	return "authors"
}

// Book 书籍表
type Book struct {
	gorm.Model

	Title       string  `gorm:"not null;index"`
	ISBN        string  `gorm:"unique;not null;index"`
	Price       float64 `gorm:"type:decimal(10,2)"`
	PublishedAt string  `gorm:"index"`
	Status      string  `gorm:"index;default:'PUBLISHED'"`

	// 外键
	AuthorID uint `gorm:"not null;index"`
}

func (*Book) TableName() string {
	return "books"
}
