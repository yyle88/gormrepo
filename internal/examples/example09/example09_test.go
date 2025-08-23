package example09_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example09/internal/models"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	db := rese.P1(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&models.Employee{}))

	testDB = db
	m.Run()
}

// TestBatchInsert 演示批量插入
func TestBatchInsert(t *testing.T) {
	// 创建多个员工记录
	employees := []*models.Employee{
		{EmployeeID: "EMP001", Name: "Alice Johnson", Department: "Engineering", Position: "Senior Developer", Email: "alice@company.com", Salary: 95000, Manager: "MGR001", HireYear: 2020},
		{EmployeeID: "EMP002", Name: "Bob Smith", Department: "Engineering", Position: "Developer", Email: "bob@company.com", Salary: 75000, Manager: "MGR001", HireYear: 2021},
		{EmployeeID: "EMP003", Name: "Charlie Brown", Department: "Marketing", Position: "Manager", Email: "charlie@company.com", Salary: 85000, Manager: "MGR002", HireYear: 2019},
		{EmployeeID: "EMP004", Name: "Diana Prince", Department: "HR", Position: "Specialist", Email: "diana@company.com", Salary: 65000, Manager: "MGR003", HireYear: 2022},
		{EmployeeID: "EMP005", Name: "Eve Wilson", Department: "Engineering", Position: "Junior Developer", Email: "eve@company.com", Salary: 55000, Manager: "MGR001", HireYear: 2023},
	}

	// 批量插入
	result := testDB.CreateInBatches(employees, 3) // 每批3条记录
	require.NoError(t, result.Error)
	require.Equal(t, int64(5), result.RowsAffected)

	t.Logf("Successfully inserted %d employees in batches", len(employees))

	// 验证插入成功
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Employee{}))
	var count int64
	err := repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Model(&models.Employee{})
	}).Count(&count).Error
	require.NoError(t, err)
	require.Equal(t, int64(5), count)
}

// TestBatchUpdate 演示批量更新
func TestBatchUpdate(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Employee{}))

	// 1. 批量更新所有工程部员工的状态
	var employee models.Employee
	cls := employee.Columns()
	err := repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Where(cls.Department.Eq("Engineering")).Model(&models.Employee{})
	}).Updates(cls.Status.Kw("PROMOTED").AsMap()).Error
	require.NoError(t, err)

	// 验证更新结果
	var promotedCount int64
	err = repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("PROMOTED")).Model(&models.Employee{})
	}).Count(&promotedCount).Error
	require.NoError(t, err)
	require.Equal(t, int64(3), promotedCount) // 工程部有3个员工

	t.Logf("Promoted %d engineering employees", promotedCount)

	// 2. 批量调整2021年入职员工的薪资
	err = repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Where(cls.HireYear.Eq(2021)).Model(&models.Employee{})
	}).Updates(map[string]interface{}{
		"salary": gorm.Expr("salary * ?", 1.1), // 涨薪10%
	}).Error
	require.NoError(t, err)

	// 验证薪资调整
	emp, err := repo.First(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Where(cls.EmployeeID.Eq("EMP002"))
	})
	require.NoError(t, err)
	require.Equal(t, float64(82500), emp.Salary) // 75000 * 1.1 = 82500

	t.Logf("Updated salary for 2021 hires, EMP002 now earns $%.0f", emp.Salary)
}

// TestBatchDelete 演示批量删除（软删除）
func TestBatchDelete(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Employee{}))

	// 获取删除前的总数
	var beforeCount int64
	err := repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Model(&models.Employee{})
	}).Count(&beforeCount).Error
	require.NoError(t, err)

	// 批量删除HR部门的员工
	err = repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Where(cls.Department.Eq("HR"))
	}).Delete(&models.Employee{}).Error
	require.NoError(t, err)

	// 验证软删除结果
	var afterCount int64
	err = repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Model(&models.Employee{})
	}).Count(&afterCount).Error
	require.NoError(t, err)

	require.Equal(t, beforeCount-1, afterCount) // HR部门有1个员工

	t.Logf("Soft deleted HR employees. Before: %d, After: %d", beforeCount, afterCount)

	// 验证软删除的记录仍在数据库中（使用Unscoped）
	var unscowdCount int64
	var employee models.Employee
	cls := employee.Columns()
	err = testDB.Unscoped().Model(&models.Employee{}).Where(cls.Department.Eq("HR")).Count(&unscowdCount).Error
	require.NoError(t, err)
	require.Equal(t, int64(1), unscowdCount)

	t.Logf("Unscoped count for HR department: %d (soft deleted records still exist)", unscowdCount)
}

// TestBatchFind 演示批量查找和处理
func TestBatchFind(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Employee{}))

	// 1. 分批处理所有员工
	var processedCount int
	batchSize := 2

	for offset := 0; ; offset += batchSize {
		employees, err := repo.FindPage(
			func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
				return db.Where(cls.Status.NotEq("DELETED"))
			},
			func(cls *models.EmployeeColumns) gormcnm.OrderByBottle {
				return cls.ID.OrderByBottle("asc")
			},
			&gormrepo.Pagination{Offset: offset, Limit: batchSize},
		)
		require.NoError(t, err)

		if len(employees) == 0 {
			break // 没有更多记录
		}

		// 模拟处理每批员工
		for _, emp := range employees {
			t.Logf("Processing employee: %s (%s)", emp.Name, emp.EmployeeID)
			processedCount++
		}

		if len(employees) < batchSize {
			break // 最后一批
		}
	}

	t.Logf("Processed %d employees in batches of %d", processedCount, batchSize)
	require.Greater(t, processedCount, 0)
}

// TestBatchOperationsWithConditions 演示基于条件的批量操作
func TestBatchOperationsWithConditions(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Employee{}))

	// 1. 按部门分组进行批量操作
	departments := []string{"Engineering", "Marketing"}

	for _, dept := range departments {
		// 获取该部门的员工数量
		var count int64
		err := repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
			return db.Where(cls.Department.Eq(dept)).Model(&models.Employee{})
		}).Count(&count).Error
		require.NoError(t, err)

		if count > 0 {
			// 为该部门员工添加部门前缀到邮箱
			err = repo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
				return db.Where(cls.Department.Eq(dept)).Model(&models.Employee{})
			}).Updates(map[string]interface{}{
				"email": gorm.Expr("REPLACE(email, '@company.com', ?)", fmt.Sprintf("@%s.company.com", dept)),
			}).Error
			require.NoError(t, err)

			t.Logf("Updated %d employees in %s department", count, dept)
		}
	}

	// 验证邮箱更新
	engEmployee, err := repo.First(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Where(cls.EmployeeID.Eq("EMP001"))
	})
	require.NoError(t, err)
	require.Contains(t, engEmployee.Email, "@Engineering.company.com")

	t.Logf("Engineering employee email updated to: %s", engEmployee.Email)
}

// TestBulkInsertWithDuplicateHandling 演示处理重复记录的批量插入
func TestBulkInsertWithDuplicateHandling(t *testing.T) {
	// 尝试插入一些员工，包括重复的EmployeeID
	newEmployees := []*models.Employee{
		{EmployeeID: "EMP006", Name: "Frank Miller", Department: "Sales", Position: "Representative", Email: "frank@company.com", Salary: 50000, Manager: "MGR004", HireYear: 2023},
		{EmployeeID: "EMP007", Name: "Grace Lee", Department: "Sales", Position: "Manager", Email: "grace@company.com", Salary: 80000, Manager: "MGR004", HireYear: 2020},
		{EmployeeID: "EMP001", Name: "Alice Johnson Updated", Department: "Engineering", Position: "Team Lead", Email: "alice.new@company.com", Salary: 105000, Manager: "MGR001", HireYear: 2020}, // 重复的EmployeeID
	}

	// 分别处理每个员工，跳过重复的
	insertedCount := 0
	skippedCount := 0

	for _, emp := range newEmployees {
		result := testDB.Create(emp)
		if result.Error != nil {
			t.Logf("Skipped duplicate employee: %s (%s)", emp.Name, emp.EmployeeID)
			skippedCount++
		} else {
			t.Logf("Inserted new employee: %s (%s)", emp.Name, emp.EmployeeID)
			insertedCount++
		}
	}

	require.Equal(t, 2, insertedCount) // EMP006 和 EMP007 应该成功插入
	require.Equal(t, 1, skippedCount)  // EMP001 应该被跳过

	t.Logf("Bulk insert completed: %d inserted, %d skipped", insertedCount, skippedCount)
}

// TestBatchTransactionOperation 演示批量事务操作
func TestBatchTransactionOperation(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Employee{}))

	// 在事务中执行批量年终奖发放
	err := testDB.Transaction(func(tx *gorm.DB) error {
		txRepo := gormrepo.NewGormRepo(gormrepo.Use(tx, &models.Employee{}))

		// 1. 为所有活跃员工增加年终奖字段（通过更新备注）
		err := txRepo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
			return db.Where(cls.Status.NotEq("INACTIVE")).Model(&models.Employee{})
		}).Updates(map[string]interface{}{
			"salary": gorm.Expr("salary + ?", 5000), // 每人5000年终奖
		}).Error
		if err != nil {
			return err
		}

		// 2. 验证没有员工的薪资超过公司上限
		var highSalaryCount int64
		err = txRepo.Where(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
			return db.Where(cls.Salary.Gt(150000)).Model(&models.Employee{})
		}).Count(&highSalaryCount).Error
		if err != nil {
			return err
		}

		if highSalaryCount > 0 {
			return fmt.Errorf("bonus would cause %d employees to exceed salary cap", highSalaryCount)
		}

		return nil
	})

	require.NoError(t, err)

	// 验证所有活跃员工都收到了年终奖
	employees, err := repo.Find(func(db *gorm.DB, cls *models.EmployeeColumns) *gorm.DB {
		return db.Where(cls.Status.NotEq("INACTIVE"))
	})
	require.NoError(t, err)

	t.Logf("Year-end bonus distributed to %d active employees", len(employees))
	for _, emp := range employees {
		t.Logf("  %s (%s): $%.0f", emp.Name, emp.EmployeeID, emp.Salary)
	}
}
