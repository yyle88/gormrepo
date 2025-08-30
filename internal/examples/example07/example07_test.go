package example07_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example07/internal/models"
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

	done.Done(db.AutoMigrate(&models.Book{}))

	// 插入测试图书数据
	books := []*models.Book{
		{Title: "Go Programming", Author: "John Doe", ISBN: "978-1", Price: 59.99, PublishYear: 2020, Rating: 4.5, Sales: 1000, Category: "tech"},
		{Title: "Python Mastery", Author: "Jane Smith", ISBN: "978-2", Price: 49.99, PublishYear: 2019, Rating: 4.2, Sales: 1500, Category: "tech"},
		{Title: "Web Development", Author: "Bob Wilson", ISBN: "978-3", Price: 69.99, PublishYear: 2021, Rating: 4.7, Sales: 800, Category: "tech"},
		{Title: "Data Science", Author: "Alice Brown", ISBN: "978-4", Price: 79.99, PublishYear: 2022, Rating: 4.8, Sales: 600, Category: "tech"},
		{Title: "Machine Learning", Author: "Charlie Davis", ISBN: "978-5", Price: 89.99, PublishYear: 2021, Rating: 4.6, Sales: 750, Category: "tech"},
		{Title: "Database Design", Author: "David Lee", ISBN: "978-6", Price: 54.99, PublishYear: 2018, Rating: 4.3, Sales: 900, Category: "tech"},
		{Title: "Clean Code", Author: "Robert Martin", ISBN: "978-7", Price: 45.99, PublishYear: 2008, Rating: 4.9, Sales: 2000, Category: "tech"},
		{Title: "Design Patterns", Author: "Gang of Four", ISBN: "978-8", Price: 65.99, PublishYear: 1994, Rating: 4.6, Sales: 1200, Category: "tech"},
	}

	for _, book := range books {
		done.Done(db.Create(book).Error)
	}

	testDB = db
	m.Run()
}

// TestBasicSorting 演示基础排序功能
func TestBasicSorting(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Book{}))

	// 1. 使用Where构建带排序的查询
	sortedBooks := make([]*models.Book, 0)
	err := repo.Where(func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
		return db.Where(cls.Category.Eq("tech")).Order(cls.Price.Ob("asc").Ox())
	}).Find(&sortedBooks).Error
	require.NoError(t, err)

	require.Greater(t, len(sortedBooks), 0)
	t.Logf("Books sorted by price (ASC): %d books", len(sortedBooks))

	// 验证排序正确性
	for i := 1; i < len(sortedBooks); i++ {
		require.LessOrEqual(t, sortedBooks[i-1].Price, sortedBooks[i].Price)
	}

	// 2. 按评分降序排序
	err = repo.Where(func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
		return db.Where(cls.Category.Eq("tech")).Order(cls.Rating.Ob("desc").Ox())
	}).Find(&sortedBooks).Error
	require.NoError(t, err)

	t.Logf("Books sorted by rating (DESC): %d books", len(sortedBooks))
	for i := 1; i < len(sortedBooks); i++ {
		require.GreaterOrEqual(t, sortedBooks[i-1].Rating, sortedBooks[i].Rating)
	}
}

// TestMultipleFieldSorting 演示多字段排序
func TestMultipleFieldSorting(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Book{}))

	// 使用 FindPage 方法进行多字段排序
	books, err := repo.FindPage(
		func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
			return db.Where(cls.Category.Eq("tech"))
		},
		func(cls *models.BookColumns) gormcnm.OrderByBottle {
			// 先按出版年份降序，再按价格升序
			return cls.PublishYear.OrderByBottle("desc").Ob(cls.Price.Ob("asc"))
		},
		&gormrepo.Pagination{Offset: 0, Limit: 10},
	)
	require.NoError(t, err)
	require.Greater(t, len(books), 0)

	t.Logf("Multi-field sorting (year DESC, price ASC):")
	for i, book := range books {
		t.Logf("  %d. %s (%d) - $%.2f", i+1, book.Title, book.PublishYear, book.Price)
	}

	// 验证排序正确性
	for i := 1; i < len(books); i++ {
		if books[i-1].PublishYear == books[i].PublishYear {
			require.LessOrEqual(t, books[i-1].Price, books[i].Price)
		} else {
			require.GreaterOrEqual(t, books[i-1].PublishYear, books[i].PublishYear)
		}
	}
}

// TestLimitQueries 演示限制查询功能
func TestLimitQueries(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Book{}))

	// 1. 获取最贵的3本书
	top3Expensive, err := repo.FindN(
		func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
			return db.Where(cls.Category.Eq("tech")).Order(cls.Price.Ob("desc").Ox())
		},
		3,
	)
	require.NoError(t, err)
	require.Len(t, top3Expensive, 3)

	t.Logf("Top 3 expensive books:")
	for i, book := range top3Expensive {
		t.Logf("  %d. %s - $%.2f", i+1, book.Title, book.Price)
	}

	// 2. 获取最新出版的5本书
	latest5Books, err := repo.FindN(
		func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
			return db.Where(cls.Category.Eq("tech")).Order(cls.PublishYear.Ob("desc").Ob(cls.ID.Ob("desc")).Ox())
		},
		5,
	)
	require.NoError(t, err)
	require.Len(t, latest5Books, 5)

	t.Logf("Latest 5 books:")
	for i, book := range latest5Books {
		t.Logf("  %d. %s (%d)", i+1, book.Title, book.PublishYear)
	}

	// 3. 获取评分最高的书（只要1本）
	topRatedBook, err := repo.First(func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
		return db.Where(cls.Category.Eq("tech")).Order(cls.Rating.OrderByBottle("desc").Ox())
	})
	require.NoError(t, err)
	require.NotNil(t, topRatedBook)

	t.Logf("Top rated book: %s (%.1f stars)", topRatedBook.Title, topRatedBook.Rating)
}

// TestComplexSortingAndFiltering 演示复杂排序和过滤
func TestComplexSortingAndFiltering(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Book{}))

	// 1. 筛选2020年后出版且价格<70的书，按销量降序排列
	books, count, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
			return db.Where(cls.PublishYear.Gt(2019)).Where(cls.Price.Lt(70.0))
		},
		func(cls *models.BookColumns) gormcnm.OrderByBottle {
			return cls.Sales.OrderByBottle("desc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 5},
	)
	require.NoError(t, err)
	require.Greater(t, len(books), 0)

	t.Logf("Books after 2019, price < $70, sorted by sales:")
	t.Logf("Found %d books, total matching: %d", len(books), count)
	for i, book := range books {
		t.Logf("  %d. %s (%d) - $%.2f, %d sales",
			i+1, book.Title, book.PublishYear, book.Price, book.Sales)
	}

	// 2. 多条件复杂查询：高评分且高销量的书
	premiumBooks, err := repo.Find(func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
		return db.Where(cls.Rating.Gt(4.5)).
			Where(cls.Sales.Gt(800)).
			Order(cls.Rating.Ob("desc").Ob(cls.Sales.Ob("desc")).Ox())
	})
	require.NoError(t, err)

	t.Logf("Premium books (rating > 4.5, sales > 800):")
	for i, book := range premiumBooks {
		t.Logf("  %d. %s - %.1f stars, %d sales",
			i+1, book.Title, book.Rating, book.Sales)
	}
}

// TestRangeQueriesWithSorting 演示范围查询和排序
func TestRangeQueriesWithSorting(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Book{}))

	// 1. 价格区间查询，按性价比排序（评分/价格）
	midRangeBooks := make([]*models.Book, 0)
	err := repo.Where(func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
		stmt, arg1, arg2 := cls.Price.BetweenAND(50.0, 70.0)
		return db.Where(stmt, arg1, arg2).
			Order("(rating / price) DESC")
	}).Find(&midRangeBooks).Error
	require.NoError(t, err)

	t.Logf("Mid-range books ($50-$70) sorted by value ratio:")
	for i, book := range midRangeBooks {
		ratio := float64(book.Rating) / book.Price
		t.Logf("  %d. %s - $%.2f, %.1f stars (ratio: %.3f)",
			i+1, book.Title, book.Price, book.Rating, ratio)
	}

	// 2. 按年代分组查询最新的书
	decades := []int{2020, 2010, 2000, 1990}
	for _, decade := range decades {
		books, err := repo.FindN(
			func(db *gorm.DB, cls *models.BookColumns) *gorm.DB {
				condition := gormcnm.Qx(cls.PublishYear.Gte(decade)).AND(gormcnm.Qx(cls.PublishYear.Lt(decade + 10)))
				return db.Scopes(condition.Scope()).
					Order(cls.PublishYear.Ob("desc").Ob(cls.Rating.Ob("desc")).Ox())
			},
			2, // 每个年代最多2本
		)
		require.NoError(t, err)

		if len(books) > 0 {
			t.Logf("Best books from %ds:", decade)
			for i, book := range books {
				t.Logf("  %d. %s (%d) - %.1f stars",
					i+1, book.Title, book.PublishYear, book.Rating)
			}
		}
	}
}
