package example06_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example06/internal/models"
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

	done.Done(db.AutoMigrate(&models.Order{}))

	// 插入测试订单数据
	orders := []*models.Order{
		{OrderNo: "ORD001", UserID: 1, Amount: 99.99, Status: "pending", PayMethod: "alipay"},
		{OrderNo: "ORD002", UserID: 1, Amount: 199.99, Status: "paid", PayMethod: "wechat"},
		{OrderNo: "ORD003", UserID: 2, Amount: 299.99, Status: "paid", PayMethod: "alipay"},
		{OrderNo: "ORD004", UserID: 2, Amount: 399.99, Status: "shipped", PayMethod: "card"},
		{OrderNo: "ORD005", UserID: 3, Amount: 499.99, Status: "delivered", PayMethod: "alipay"},
		{OrderNo: "ORD006", UserID: 3, Amount: 599.99, Status: "cancelled", PayMethod: "wechat"},
		{OrderNo: "ORD007", UserID: 1, Amount: 99.99, Status: "refunded", PayMethod: "card"},
	}
	for _, order := range orders {
		done.Done(db.Create(order).Error)
	}

	testDB = db
	m.Run()
}

// TestBasicCountOperations 演示基础计数操作
func TestBasicCountOperations(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Order{}))

	// 1. 统计所有订单数量
	totalCount, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db
	})
	require.NoError(t, err)
	require.Equal(t, int64(7), totalCount)
	t.Logf("Total orders count: %d", totalCount)

	// 2. 按状态统计订单数量
	paidCount, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("paid"))
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), paidCount)
	t.Logf("Paid orders count: %d", paidCount)

	// 3. 按用户统计订单数量
	user1Count, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.UserID.Eq(1))
	})
	require.NoError(t, err)
	require.Equal(t, int64(3), user1Count)
	t.Logf("User 1 orders count: %d", user1Count)

	// 4. 按支付方式统计
	alipayCount, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.PayMethod.Eq("alipay"))
	})
	require.NoError(t, err)
	require.Equal(t, int64(3), alipayCount)
	t.Logf("Alipay orders count: %d", alipayCount)
}

// TestConditionalCountOperations 演示条件计数操作
func TestConditionalCountOperations(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Order{}))

	// 1. 统计金额范围内的订单
	midRangeCount, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		stmt, arg1, arg2 := cls.Amount.BetweenAND(200.0, 400.0)
		return db.Where(stmt, arg1, arg2)
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), midRangeCount) // ORD003 和 ORD004
	t.Logf("Mid-range orders count: %d", midRangeCount)

	// 2. 统计高价值订单
	highValueCount, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.Amount.Gt(300.0))
	})
	require.NoError(t, err)
	require.Equal(t, int64(3), highValueCount)
	t.Logf("High value orders count: %d", highValueCount)

	// 3. 复合条件计数 - 特定用户的特定状态订单
	complexCount, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.UserID.Eq(1)).Where(cls.Status.In([]string{"paid", "pending"}))
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), complexCount)
	t.Logf("User 1 paid/pending orders count: %d", complexCount)

	// 4. 统计非取消状态的订单
	activeCount, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.Status.Ne("cancelled"))
	})
	require.NoError(t, err)
	require.Equal(t, int64(6), activeCount)
	t.Logf("Active orders count: %d", activeCount)
}

// TestExistenceOperations 演示存在性检查操作
func TestExistenceOperations(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Order{}))

	// 1. 检查特定订单号是否存在
	exists, err := repo.Exist(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.OrderNo.Eq("ORD001"))
	})
	require.NoError(t, err)
	require.True(t, exists)
	t.Logf("Order ORD001 exists: %v", exists)

	// 2. 检查不存在的订单号
	exists, err = repo.Exist(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.OrderNo.Eq("ORD999"))
	})
	require.NoError(t, err)
	require.False(t, exists)
	t.Logf("Order ORD999 exists: %v", exists)

	// 3. 检查特定用户是否有订单
	exists, err = repo.Exist(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.UserID.Eq(4))
	})
	require.NoError(t, err)
	require.False(t, exists)
	t.Logf("User 4 has orders: %v", exists)

	// 4. 检查是否存在高价值订单
	exists, err = repo.Exist(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.Amount.Gt(500.0))
	})
	require.NoError(t, err)
	require.True(t, exists)
	t.Logf("High value orders exist: %v", exists)

	// 5. 检查特定支付方式和状态的组合是否存在
	exists, err = repo.Exist(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.PayMethod.Eq("card")).Where(cls.Status.Eq("shipped"))
	})
	require.NoError(t, err)
	require.True(t, exists)
	t.Logf("Card payment shipped orders exist: %v", exists)
}

// TestCountByGroups 演示分组计数操作
func TestCountByGroups(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Order{}))

	// 1. 按状态分组统计
	statusCounts := make(map[string]int64)
	statuses := []string{"pending", "paid", "shipped", "delivered", "cancelled", "refunded"}

	for _, status := range statuses {
		count, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
			return db.Where(cls.Status.Eq(status))
		})
		require.NoError(t, err)
		statusCounts[status] = count
	}

	t.Logf("Orders by status:")
	for status, count := range statusCounts {
		if count > 0 {
			t.Logf("  %s: %d", status, count)
		}
	}

	// 2. 按用户分组统计
	userCounts := make(map[uint]int64)
	userIDs := []uint{1, 2, 3}

	for _, userID := range userIDs {
		count, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
			return db.Where(cls.UserID.Eq(userID))
		})
		require.NoError(t, err)
		userCounts[userID] = count
	}

	t.Logf("Orders by user:")
	for userID, count := range userCounts {
		t.Logf("  User %d: %d orders", userID, count)
	}

	// 3. 按支付方式分组统计
	payMethodCounts := make(map[string]int64)
	payMethods := []string{"alipay", "wechat", "card"}

	for _, method := range payMethods {
		count, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
			return db.Where(cls.PayMethod.Eq(method))
		})
		require.NoError(t, err)
		payMethodCounts[method] = count
	}

	t.Logf("Orders by payment method:")
	for method, count := range payMethodCounts {
		t.Logf("  %s: %d", method, count)
	}
}

// TestComplexCountAndExist 演示复杂的计数和存在性检查
func TestComplexCountAndExist(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Order{}))

	// 1. 检查每个用户是否都有订单
	for userID := uint(1); userID <= 5; userID++ {
		exists, err := repo.Exist(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
			return db.Where(cls.UserID.Eq(userID))
		})
		require.NoError(t, err)

		if exists {
			count, err := repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
				return db.Where(cls.UserID.Eq(userID))
			})
			require.NoError(t, err)
			t.Logf("User %d: has %d orders", userID, count)
		} else {
			t.Logf("User %d: no orders", userID)
		}
	}

	// 2. 统计业务指标
	metrics := struct {
		TotalOrders     int64
		PendingOrders   int64
		CompletedOrders int64
		TotalRevenue    float64
		AvgOrderValue   float64
	}{}

	// 总订单数
	metrics.TotalOrders, _ = repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db
	})

	// 待处理订单数
	metrics.PendingOrders, _ = repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("pending"))
	})

	// 已完成订单数 (delivered)
	metrics.CompletedOrders, _ = repo.Count(func(db *gorm.DB, cls *models.OrderColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("delivered"))
	})

	// 计算总收入和平均订单价值 (需要用原生 SQL)
	var revenue struct {
		Total float64
		Avg   float64
	}
	var order models.Order
	cls := order.Columns()
	testDB.Model(&models.Order{}).
		Select(fmt.Sprintf("SUM(%s) as total, AVG(%s) as avg", cls.Amount.Name(), cls.Amount.Name())).
		Where(cls.Status.NotIn([]string{"cancelled", "refunded"})).
		Scan(&revenue)

	metrics.TotalRevenue = revenue.Total
	metrics.AvgOrderValue = revenue.Avg

	t.Logf("Business Metrics:")
	t.Logf("  Total Orders: %d", metrics.TotalOrders)
	t.Logf("  Pending Orders: %d", metrics.PendingOrders)
	t.Logf("  Completed Orders: %d", metrics.CompletedOrders)
	t.Logf("  Total Revenue: %.2f", metrics.TotalRevenue)
	t.Logf("  Average Order Value: %.2f", metrics.AvgOrderValue)

	// 验证一些基本假设
	require.Greater(t, metrics.TotalOrders, int64(0))
	require.Greater(t, metrics.TotalRevenue, 0.0)
	require.Greater(t, metrics.AvgOrderValue, 0.0)
}
