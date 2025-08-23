package example04_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example04/internal/models"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	// 创建内存数据库
	db := rese.P1(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// 自动迁移两个表
	done.Done(db.AutoMigrate(&models.Category{}, &models.Product{}))

	// 插入测试分类数据
	categories := []*models.Category{
		{Name: "电子产品", Code: "electronics", Sort: 1, Status: "active"},
		{Name: "服装", Code: "clothing", Sort: 2, Status: "active"},
		{Name: "图书", Code: "books", Sort: 3, Status: "inactive"},
	}
	for _, category := range categories {
		done.Done(db.Create(category).Error)
	}

	// 插入测试产品数据
	products := []*models.Product{
		{Name: "iPhone 15", Price: 999.99, Stock: 50, CategoryID: 1, Status: "active"},
		{Name: "MacBook Pro", Price: 1999.99, Stock: 20, CategoryID: 1, Status: "active"},
		{Name: "T恤", Price: 29.99, Stock: 100, CategoryID: 2, Status: "active"},
		{Name: "牛仔裤", Price: 79.99, Stock: 0, CategoryID: 2, Status: "inactive"},
		{Name: "Go语言圣经", Price: 59.99, Stock: 30, CategoryID: 3, Status: "active"},
	}
	for _, product := range products {
		done.Done(db.Create(product).Error)
	}

	testDB = db
	m.Run()
}

// TestTwoModelsBasicOperations 演示双模型的基础操作
func TestTwoModelsBasicOperations(t *testing.T) {
	// 1. 分类相关操作
	categoryRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Category{}))

	// 查询活跃分类
	activeCategories, err := categoryRepo.Find(func(db *gorm.DB, cls *models.CategoryColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("active")).Order(cls.Sort.OrderByBottle("asc").Ox())
	})
	require.NoError(t, err)
	require.Len(t, activeCategories, 2)
	require.Equal(t, "电子产品", activeCategories[0].Name)
	require.Equal(t, "服装", activeCategories[1].Name)
	t.Logf("Found %d active categories", len(activeCategories))

	// 2. 产品相关操作
	productRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Product{}))

	// 查询高价产品
	expensiveProducts, err := productRepo.Find(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		return db.Where(cls.Price.Gt(100.0)).Order(cls.Price.OrderByBottle("desc").Ox())
	})
	require.NoError(t, err)
	require.Len(t, expensiveProducts, 2)
	require.Equal(t, "MacBook Pro", expensiveProducts[0].Name)
	require.Equal(t, "iPhone 15", expensiveProducts[1].Name)
	t.Logf("Found %d expensive products", len(expensiveProducts))

	// 3. 查询特定分类的产品
	electronicsProducts, err := productRepo.Find(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		return db.Where(cls.CategoryID.Eq(1))
	})
	require.NoError(t, err)
	require.Len(t, electronicsProducts, 2)
	t.Logf("Electronics category has %d products", len(electronicsProducts))
}

// TestComplexQueries 演示复杂查询功能
func TestComplexQueries(t *testing.T) {
	productRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Product{}))

	// 1. 价格区间查询
	midRangeProducts, err := productRepo.Find(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		stmt, arg1, arg2 := cls.Price.BetweenAND(50.0, 100.0)
		return db.Where(stmt, arg1, arg2)
	})
	require.NoError(t, err)
	require.Len(t, midRangeProducts, 2) // T恤 和 Go语言圣经
	t.Logf("Mid-range products: %d", len(midRangeProducts))

	// 2. 多条件组合查询 - 有库存且活跃的产品
	availableProducts, err := productRepo.Find(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		return db.Where(cls.Stock.Gt(0)).Where(cls.Status.Eq("active"))
	})
	require.NoError(t, err)
	require.Len(t, availableProducts, 4) // 排除库存为0的牛仔裤
	t.Logf("Available products: %d", len(availableProducts))

	// 3. 名称模糊查询
	iphoneProducts, err := productRepo.Find(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		return db.Where(cls.Name.Like("%iPhone%"))
	})
	require.NoError(t, err)
	require.Len(t, iphoneProducts, 1)
	require.Equal(t, "iPhone 15", iphoneProducts[0].Name)
	t.Logf("iPhone products: %d", len(iphoneProducts))

	// 4. IN 查询 - 特定分类的产品
	targetCategories := []uint{1, 2}
	categoryProducts, err := productRepo.Find(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		return db.Where(cls.CategoryID.In(targetCategories))
	})
	require.NoError(t, err)
	require.Len(t, categoryProducts, 4) // 电子产品2个 + 服装2个
	t.Logf("Products in target categories: %d", len(categoryProducts))
}

// TestCountAndExist 演示计数和存在性检查
func TestCountAndExist(t *testing.T) {
	productRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Product{}))
	categoryRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Category{}))

	// 1. 计算每个分类的产品数量
	electronicsCount, err := productRepo.Count(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		return db.Where(cls.CategoryID.Eq(1))
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), electronicsCount)
	t.Logf("Electronics products count: %d", electronicsCount)

	// 2. 检查是否存在特定价格的产品
	exists, err := productRepo.Exist(func(db *gorm.DB, cls *models.ProductColumns) *gorm.DB {
		return db.Where(cls.Price.Gt(1500.0))
	})
	require.NoError(t, err)
	require.True(t, exists) // MacBook Pro 价格 > 1500
	t.Logf("Expensive product exists: %v", exists)

	// 3. 检查是否存在特定编码的分类
	exists, err = categoryRepo.Exist(func(db *gorm.DB, cls *models.CategoryColumns) *gorm.DB {
		return db.Where(cls.Code.Eq("electronics"))
	})
	require.NoError(t, err)
	require.True(t, exists)
	t.Logf("Electronics category exists: %v", exists)

	// 4. 统计活跃分类数量
	activeCategoryCount, err := categoryRepo.Count(func(db *gorm.DB, cls *models.CategoryColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("active"))
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), activeCategoryCount)
	t.Logf("Active categories count: %d", activeCategoryCount)
}
