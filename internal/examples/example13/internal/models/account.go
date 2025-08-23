package models

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AccountNumber string  `gorm:"unique;not null;index"`
	AccountName   string  `gorm:"not null"`
	Balance       float64 `gorm:"type:decimal(15,2);not null;default:0"`
	AccountType   string  `gorm:"index;default:'SAVINGS'"`
	Status        string  `gorm:"index;default:'ACTIVE'"`
	BankCode      string  `gorm:"index"`
	BranchCode    string  `gorm:"index"`
}

func (*Account) TableName() string {
	return "accounts"
}

type Transaction struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TransactionID     string  `gorm:"unique;not null;index"`
	FromAccountNumber string  `gorm:"index"`
	ToAccountNumber   string  `gorm:"index"`
	Amount            float64 `gorm:"type:decimal(15,2);not null"`
	TransactionType   string  `gorm:"not null;index"`
	Description       string
	Status            string `gorm:"index;default:'PENDING'"`
	Reference         string `gorm:"index"`
}

func (*Transaction) TableName() string {
	return "transactions"
}
