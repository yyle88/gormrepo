package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (*Account) Columns() *AccountColumns {
	return &AccountColumns{
		ID:            "id",
		CreatedAt:     "created_at",
		UpdatedAt:     "updated_at",
		DeletedAt:     "deleted_at",
		AccountNumber: "account_number",
		AccountName:   "account_name",
		Balance:       "balance",
		AccountType:   "account_type",
		Status:        "status",
		BankCode:      "bank_code",
		BranchCode:    "branch_code",
	}
}

type AccountColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID            gormcnm.ColumnName[uint]
	CreatedAt     gormcnm.ColumnName[time.Time]
	UpdatedAt     gormcnm.ColumnName[time.Time]
	DeletedAt     gormcnm.ColumnName[gorm.DeletedAt]
	AccountNumber gormcnm.ColumnName[string]
	AccountName   gormcnm.ColumnName[string]
	Balance       gormcnm.ColumnName[float64]
	AccountType   gormcnm.ColumnName[string]
	Status        gormcnm.ColumnName[string]
	BankCode      gormcnm.ColumnName[string]
	BranchCode    gormcnm.ColumnName[string]
}

func (*Transaction) Columns() *TransactionColumns {
	return &TransactionColumns{
		ID:                "id",
		CreatedAt:         "created_at",
		UpdatedAt:         "updated_at",
		DeletedAt:         "deleted_at",
		TransactionID:     "transaction_id",
		FromAccountNumber: "from_account_number",
		ToAccountNumber:   "to_account_number",
		Amount:            "amount",
		TransactionType:   "transaction_type",
		Description:       "description",
		Status:            "status",
		Reference:         "reference",
	}
}

type TransactionColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID                gormcnm.ColumnName[uint]
	CreatedAt         gormcnm.ColumnName[time.Time]
	UpdatedAt         gormcnm.ColumnName[time.Time]
	DeletedAt         gormcnm.ColumnName[gorm.DeletedAt]
	TransactionID     gormcnm.ColumnName[string]
	FromAccountNumber gormcnm.ColumnName[string]
	ToAccountNumber   gormcnm.ColumnName[string]
	Amount            gormcnm.ColumnName[float64]
	TransactionType   gormcnm.ColumnName[string]
	Description       gormcnm.ColumnName[string]
	Status            gormcnm.ColumnName[string]
	Reference         gormcnm.ColumnName[string]
}
