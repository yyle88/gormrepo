package example1models

import (
	"time"
)

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"unique;"`
	Age       int       `gorm:"type:int32;index:idx_example_age;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (*Example) TableName() string {
	return "examples"
}
