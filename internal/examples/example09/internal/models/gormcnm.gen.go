package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (c *Employee) Columns() *EmployeeColumns {
	return &EmployeeColumns{
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
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
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
