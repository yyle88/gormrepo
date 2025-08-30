package example12_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo/internal/examples/example12/internal/models"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&models.SaleRecord{}))

	// 插入测试销售数据
	now := time.Now()
	records := []*models.SaleRecord{
		{ProductName: "iPhone 15", Category: "Electronics", Region: "North", SaleDate: now.AddDate(0, -2, 0), Quantity: 10, UnitPrice: 999.99, TotalAmount: 9999.90, SalesRep: "Alice", Channel: "Online"},
		{ProductName: "iPhone 15", Category: "Electronics", Region: "South", SaleDate: now.AddDate(0, -2, -5), Quantity: 15, UnitPrice: 999.99, TotalAmount: 14999.85, SalesRep: "Bob", Channel: "Store"},
		{ProductName: "MacBook Pro", Category: "Electronics", Region: "North", SaleDate: now.AddDate(0, -1, -10), Quantity: 5, UnitPrice: 1999.99, TotalAmount: 9999.95, SalesRep: "Alice", Channel: "Online"},
		{ProductName: "iPad", Category: "Electronics", Region: "West", SaleDate: now.AddDate(0, -1, -3), Quantity: 20, UnitPrice: 599.99, TotalAmount: 11999.80, SalesRep: "Charlie", Channel: "Store"},
		{ProductName: "AirPods", Category: "Electronics", Region: "East", SaleDate: now.AddDate(0, -1, 0), Quantity: 30, UnitPrice: 249.99, TotalAmount: 7499.70, SalesRep: "Alice", Channel: "Online"},
		{ProductName: "T-Shirt", Category: "Clothing", Region: "North", SaleDate: now.AddDate(0, -2, -2), Quantity: 50, UnitPrice: 29.99, TotalAmount: 1499.50, SalesRep: "David", Channel: "Store"},
		{ProductName: "Jeans", Category: "Clothing", Region: "South", SaleDate: now.AddDate(0, -1, -5), Quantity: 25, UnitPrice: 79.99, TotalAmount: 1999.75, SalesRep: "Eve", Channel: "Online"},
		{ProductName: "Sneakers", Category: "Clothing", Region: "West", SaleDate: now.AddDate(0, -1, -1), Quantity: 15, UnitPrice: 129.99, TotalAmount: 1949.85, SalesRep: "David", Channel: "Store"},
	}

	for _, record := range records {
		done.Done(db.Create(record).Error)
	}

	testDB = db
	m.Run()
}

// TestBasicAggregation 演示基础聚合查询
func TestBasicAggregation(t *testing.T) {
	// 使用原始 SQL 进行聚合查询
	type CategoryStats struct {
		Category    string  `json:"category"`
		TotalSales  float64 `json:"total_sales"`
		TotalQty    int64   `json:"total_qty"`
		AvgPrice    float64 `json:"avg_price"`
		RecordCount int64   `json:"record_count"`
	}

	var results []CategoryStats
	var saleRecord models.SaleRecord
	cls := saleRecord.Columns()
	var common = &gormcnm.ColumnOperationClass{}

	err := testDB.Model(&models.SaleRecord{}).
		Select(common.MergeStmts(
			cls.Category.Name(),
			cls.TotalAmount.COALESCE().SumStmt("total_sales"),
			cls.Quantity.COALESCE().SumStmt("total_qty"),
			cls.UnitPrice.COALESCE().AvgStmt("avg_price"),
			common.CountStmt("record_count"),
		)).
		Group(cls.Category.Name()).
		Order(cls.TotalAmount.OrderByBottle("desc").Ox()).
		Find(&results).Error
	require.NoError(t, err)
	require.Greater(t, len(results), 0)

	t.Logf("Category-wise Sales Summary:")
	for i, result := range results {
		t.Logf("  %d. %s: Sales=$%.2f, Qty=%d, AvgPrice=$%.2f, Records=%d",
			i+1, result.Category, result.TotalSales, result.TotalQty, result.AvgPrice, result.RecordCount)
	}
}

// TestRegionalAggregation 演示区域聚合查询
func TestRegionalAggregation(t *testing.T) {
	type RegionStats struct {
		Region      string  `json:"region"`
		TotalSales  float64 `json:"total_sales"`
		TotalQty    int64   `json:"total_qty"`
		RecordCount int64   `json:"record_count"`
	}

	var regionStats []RegionStats
	err := testDB.Model(&models.SaleRecord{}).
		Select("region, SUM(total_amount) as total_sales, SUM(quantity) as total_qty, COUNT(*) as record_count").
		Group("region").
		Order("total_sales DESC").
		Find(&regionStats).Error
	require.NoError(t, err)
	require.Greater(t, len(regionStats), 0)

	t.Logf("Regional Sales Summary:")
	for i, stat := range regionStats {
		t.Logf("  %d. %s: Sales=$%.2f, Qty=%d, Records=%d",
			i+1, stat.Region, stat.TotalSales, stat.TotalQty, stat.RecordCount)
	}
}

// TestSalesRepPerformance 演示销售代表业绩统计
func TestSalesRepPerformance(t *testing.T) {
	type SalesRepStats struct {
		SalesRep    string  `json:"sales_rep"`
		TotalSales  float64 `json:"total_sales"`
		TotalQty    int64   `json:"total_qty"`
		RecordCount int64   `json:"record_count"`
		MaxSale     float64 `json:"max_sale"`
		MinSale     float64 `json:"min_sale"`
	}

	var repStats []SalesRepStats
	err := testDB.Model(&models.SaleRecord{}).
		Select("sales_rep, SUM(total_amount) as total_sales, SUM(quantity) as total_qty, COUNT(*) as record_count, MAX(total_amount) as max_sale, MIN(total_amount) as min_sale").
		Group("sales_rep").
		Order("total_sales DESC").
		Find(&repStats).Error
	require.NoError(t, err)
	require.Greater(t, len(repStats), 0)

	t.Logf("Sales Representative Performance:")
	for i, stat := range repStats {
		t.Logf("  %d. %s: Sales=$%.2f, Qty=%d, Records=%d, Max=$%.2f, Min=$%.2f",
			i+1, stat.SalesRep, stat.TotalSales, stat.TotalQty, stat.RecordCount, stat.MaxSale, stat.MinSale)
	}
}

// TestMonthlyTrends 演示月度趋势分析
func TestMonthlyTrends(t *testing.T) {
	type MonthlyStats struct {
		Month         string  `json:"month"`
		TotalSales    float64 `json:"total_sales"`
		TotalQty      int64   `json:"total_qty"`
		RecordCount   int64   `json:"record_count"`
		AvgOrderValue float64 `json:"avg_order_value"`
	}

	var monthlyStats []MonthlyStats
	// SQLite 的日期函数
	err := testDB.Model(&models.SaleRecord{}).
		Select("strftime('%Y-%m', sale_date) as month, SUM(total_amount) as total_sales, SUM(quantity) as total_qty, COUNT(*) as record_count, AVG(total_amount) as avg_order_value").
		Group("strftime('%Y-%m', sale_date)").
		Order("month DESC").
		Find(&monthlyStats).Error
	require.NoError(t, err)
	require.Greater(t, len(monthlyStats), 0)

	t.Logf("Monthly Sales Trends:")
	for i, stat := range monthlyStats {
		t.Logf("  %d. %s: Sales=$%.2f, Qty=%d, Records=%d, AOV=$%.2f",
			i+1, stat.Month, stat.TotalSales, stat.TotalQty, stat.RecordCount, stat.AvgOrderValue)
	}
}

// TestChannelComparison 演示销售渠道对比
func TestChannelComparison(t *testing.T) {
	type ChannelStats struct {
		Channel      string  `json:"channel"`
		TotalSales   float64 `json:"total_sales"`
		TotalQty     int64   `json:"total_qty"`
		RecordCount  int64   `json:"record_count"`
		AvgUnitPrice float64 `json:"avg_unit_price"`
	}

	var channelStats []ChannelStats
	err := testDB.Model(&models.SaleRecord{}).
		Select("channel, SUM(total_amount) as total_sales, SUM(quantity) as total_qty, COUNT(*) as record_count, AVG(unit_price) as avg_unit_price").
		Group("channel").
		Order("total_sales DESC").
		Find(&channelStats).Error
	require.NoError(t, err)
	require.Greater(t, len(channelStats), 0)

	t.Logf("Sales Channel Comparison:")
	for i, stat := range channelStats {
		t.Logf("  %d. %s: Sales=$%.2f, Qty=%d, Records=%d, AvgPrice=$%.2f",
			i+1, stat.Channel, stat.TotalSales, stat.TotalQty, stat.RecordCount, stat.AvgUnitPrice)
	}

	// 计算渠道占比
	var totalSales float64
	for _, stat := range channelStats {
		totalSales += stat.TotalSales
	}

	t.Logf("Channel Sales Percentage:")
	for i, stat := range channelStats {
		percentage := (stat.TotalSales / totalSales) * 100
		t.Logf("  %d. %s: %.1f%%", i+1, stat.Channel, percentage)
	}
}

// TestTopProductsByCategory 演示按类别统计热销产品
func TestTopProductsByCategory(t *testing.T) {
	type ProductStats struct {
		Category    string  `json:"category"`
		ProductName string  `json:"product_name"`
		TotalSales  float64 `json:"total_sales"`
		TotalQty    int64   `json:"total_qty"`
		RecordCount int64   `json:"record_count"`
	}

	var productStats []ProductStats
	err := testDB.Model(&models.SaleRecord{}).
		Select("category, product_name, SUM(total_amount) as total_sales, SUM(quantity) as total_qty, COUNT(*) as record_count").
		Group("category, product_name").
		Order("category ASC, total_sales DESC").
		Find(&productStats).Error
	require.NoError(t, err)
	require.Greater(t, len(productStats), 0)

	// 按类别分组显示
	currentCategory := ""
	t.Logf("Top Products by Category:")
	for _, stat := range productStats {
		if stat.Category != currentCategory {
			currentCategory = stat.Category
			t.Logf("  === %s Category ===", currentCategory)
		}
		t.Logf("    - %s: Sales=$%.2f, Qty=%d, Records=%d",
			stat.ProductName, stat.TotalSales, stat.TotalQty, stat.RecordCount)
	}
}

// TestComplexAggregationWithFilters 演示复杂聚合查询与过滤
func TestComplexAggregationWithFilters(t *testing.T) {
	// 查询最近30天的高价值订单统计（总金额>5000）
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30)

	type HighValueStats struct {
		Category    string  `json:"category"`
		Region      string  `json:"region"`
		TotalSales  float64 `json:"total_sales"`
		RecordCount int64   `json:"record_count"`
		AvgSale     float64 `json:"avg_sale"`
	}

	var highValueStats []HighValueStats
	err := testDB.Model(&models.SaleRecord{}).
		Select("category, region, SUM(total_amount) as total_sales, COUNT(*) as record_count, AVG(total_amount) as avg_sale").
		Where("sale_date > ? AND total_amount > ?", thirtyDaysAgo, 5000.0).
		Group("category, region").
		Having("total_sales > ?", 10000.0).
		Order("total_sales DESC").
		Find(&highValueStats).Error
	require.NoError(t, err)

	t.Logf("High-Value Orders (>$5000) in Last 30 Days:")
	if len(highValueStats) > 0 {
		for i, stat := range highValueStats {
			t.Logf("  %d. %s-%s: Sales=$%.2f, Records=%d, Avg=$%.2f",
				i+1, stat.Category, stat.Region, stat.TotalSales, stat.RecordCount, stat.AvgSale)
		}
	} else {
		t.Logf("  No high-value order groups found meeting criteria")
	}

	// 统计符合条件的记录总数
	var totalHighValueCount int64
	err = testDB.Model(&models.SaleRecord{}).
		Where("sale_date > ? AND total_amount > ?", thirtyDaysAgo, 5000.0).
		Count(&totalHighValueCount).Error
	require.NoError(t, err)

	t.Logf("Total high-value orders: %d", totalHighValueCount)
}
