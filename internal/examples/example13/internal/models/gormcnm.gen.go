// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: gormcnm.gen_test.go:43 -> models_test.TestGenerateColumns
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Account) Columns() *AccountColumns {
	return &AccountColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:            gormcnm.Cnm(c.ID, "id"),
		CreatedAt:     gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt:     gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt:     gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		AccountNumber: gormcnm.Cnm(c.AccountNumber, "account_number"),
		AccountName:   gormcnm.Cnm(c.AccountName, "account_name"),
		Balance:       gormcnm.Cnm(c.Balance, "balance"),
		AccountType:   gormcnm.Cnm(c.AccountType, "account_type"),
		Status:        gormcnm.Cnm(c.Status, "status"),
		BankCode:      gormcnm.Cnm(c.BankCode, "bank_code"),
		BranchCode:    gormcnm.Cnm(c.BranchCode, "branch_code"),
	}
}

type AccountColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
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

func (c *Transaction) Columns() *TransactionColumns {
	return &TransactionColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:                gormcnm.Cnm(c.ID, "id"),
		CreatedAt:         gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt:         gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt:         gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		TransactionID:     gormcnm.Cnm(c.TransactionID, "transaction_id"),
		FromAccountNumber: gormcnm.Cnm(c.FromAccountNumber, "from_account_number"),
		ToAccountNumber:   gormcnm.Cnm(c.ToAccountNumber, "to_account_number"),
		Amount:            gormcnm.Cnm(c.Amount, "amount"),
		TransactionType:   gormcnm.Cnm(c.TransactionType, "transaction_type"),
		Description:       gormcnm.Cnm(c.Description, "description"),
		Status:            gormcnm.Cnm(c.Status, "status"),
		Reference:         gormcnm.Cnm(c.Reference, "reference"),
	}
}

type TransactionColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
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
