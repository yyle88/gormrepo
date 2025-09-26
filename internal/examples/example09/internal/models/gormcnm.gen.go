// Code generated using gormcngen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/gormcngen
// Generated from: gormcnm.gen_test.go:42 -> models_test.TestGenerateColumns
// ========== GORMCNGEN:DO-NOT-EDIT-MARKER:END ==========

package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Employee) Columns() *EmployeeColumns {
	return &EmployeeColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:         gormcnm.Cnm(c.ID, "id"),
		CreatedAt:  gormcnm.Cnm(c.CreatedAt, "created_at"),
		UpdatedAt:  gormcnm.Cnm(c.UpdatedAt, "updated_at"),
		DeletedAt:  gormcnm.Cnm(c.DeletedAt, "deleted_at"),
		EmployeeID: gormcnm.Cnm(c.EmployeeID, "employee_id"),
		Name:       gormcnm.Cnm(c.Name, "name"),
		Department: gormcnm.Cnm(c.Department, "department"),
		Position:   gormcnm.Cnm(c.Position, "position"),
		Email:      gormcnm.Cnm(c.Email, "email"),
		Salary:     gormcnm.Cnm(c.Salary, "salary"),
		Status:     gormcnm.Cnm(c.Status, "status"),
		Manager:    gormcnm.Cnm(c.Manager, "manager"),
		HireYear:   gormcnm.Cnm(c.HireYear, "hire_year"),
	}
}

type EmployeeColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID         gormcnm.ColumnName[uint]
	CreatedAt  gormcnm.ColumnName[time.Time]
	UpdatedAt  gormcnm.ColumnName[time.Time]
	DeletedAt  gormcnm.ColumnName[gorm.DeletedAt]
	EmployeeID gormcnm.ColumnName[string]
	Name       gormcnm.ColumnName[string]
	Department gormcnm.ColumnName[string]
	Position   gormcnm.ColumnName[string]
	Email      gormcnm.ColumnName[string]
	Salary     gormcnm.ColumnName[float64]
	Status     gormcnm.ColumnName[string]
	Manager    gormcnm.ColumnName[string]
	HireYear   gormcnm.ColumnName[int]
}
