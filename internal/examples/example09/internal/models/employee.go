package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model

	EmployeeID string  `gorm:"unique;not null;index"`
	Name       string  `gorm:"not null"`
	Department string  `gorm:"index"`
	Position   string  `gorm:"index"`
	Email      string  `gorm:"unique;index"`
	Salary     float64 `gorm:"type:decimal(10,2)"`
	Status     string  `gorm:"index;default:'ACTIVE'"`
	Manager    string  `gorm:"index"`
	HireYear   int     `gorm:"index"`
}

func (*Employee) TableName() string {
	return "employees"
}
