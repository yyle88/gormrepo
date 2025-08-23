package models

import (
	"time"

	"gorm.io/gorm"
)

// Article 文章模型 - 演示更新操作
type Article struct {
	ID          uint       `gorm:"primaryKey"`
	Title       string     `gorm:"size:200;not null;index"`
	Content     string     `gorm:"type:text"`
	AuthorID    uint       `gorm:"index"`
	Status      string     `gorm:"size:20;default:'draft';index"`
	ViewCount   int        `gorm:"default:0"`
	LikeCount   int        `gorm:"default:0"`
	PublishedAt *time.Time `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (*Article) TableName() string {
	return "articles"
}
