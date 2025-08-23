package models

import (
	"time"

	"github.com/yyle88/gormcnm"
	"gorm.io/gorm"
)

func (*Employee) Columns() *EmployeeColumns {
	return &EmployeeColumns{
		ID:         "id",
		CreatedAt:  "created_at",
		UpdatedAt:  "updated_at",
		DeletedAt:  "deleted_at",
		EmployeeID: "employee_id",
		Name:       "name",
		Department: "department",
		Position:   "position",
		Email:      "email",
		Salary:     "salary",
		Status:     "status",
		Manager:    "manager",
		HireYear:   "hire_year",
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
