package example08_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example08/internal/models"
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

	done.Done(db.AutoMigrate(&models.Post{}))

	// 插入大量测试帖子数据
	posts := make([]*models.Post, 0, 25)
	categories := []string{"tech", "life", "travel", "food", "sports"}
	authors := []uint{1, 2, 3, 4, 5}

	for i := 1; i <= 25; i++ {
		post := &models.Post{
			Title:     fmt.Sprintf("Post Title %02d", i),
			Content:   fmt.Sprintf("This is the content of post %d", i),
			AuthorID:  authors[i%len(authors)],
			Category:  categories[i%len(categories)],
			Tags:      fmt.Sprintf("tag%d,common", i%3+1),
			ViewCount: i * 10,
			LikeCount: i * 2,
			Status:    "published",
		}
		if i > 20 {
			post.Status = "draft"
		}
		posts = append(posts, post)
	}

	for _, post := range posts {
		done.Done(db.Create(post).Error)
	}

	testDB = db
	m.Run()
}

// TestBasicPagination 演示基础分页查询
func TestBasicPagination(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Post{}))

	// 1. 第一页数据 (每页5条)
	page1, err := repo.FindPage(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("asc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 5},
	)
	require.NoError(t, err)
	require.Len(t, page1, 5)
	require.Equal(t, "Post Title 01", page1[0].Title)
	require.Equal(t, "Post Title 05", page1[4].Title)
	t.Logf("Page 1: %d posts", len(page1))

	// 2. 第二页数据
	page2, err := repo.FindPage(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("asc")
		},
		&gormrepo.Pagination{Offset: 5, Limit: 5},
	)
	require.NoError(t, err)
	require.Len(t, page2, 5)
	require.Equal(t, "Post Title 06", page2[0].Title)
	require.Equal(t, "Post Title 10", page2[4].Title)
	t.Logf("Page 2: %d posts", len(page2))

	// 3. 最后一页数据
	lastPage, err := repo.FindPage(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("asc")
		},
		&gormrepo.Pagination{Offset: 15, Limit: 5},
	)
	require.NoError(t, err)
	require.Len(t, lastPage, 5)
	t.Logf("Last page: %d posts", len(lastPage))
}

// TestPaginationWithCount 演示带总数的分页查询
func TestPaginationWithCount(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Post{}))

	// 1. 分页查询并获取总数
	posts, totalCount, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("desc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 10},
	)
	require.NoError(t, err)
	require.Len(t, posts, 10)
	require.Equal(t, int64(20), totalCount) // 应该有20条已发布的帖子

	t.Logf("Found %d posts, total: %d", len(posts), totalCount)
	t.Logf("First post: %s", posts[0].Title)
	t.Logf("Last post: %s", posts[9].Title)

	// 2. 计算分页信息
	pageSize := 10
	totalPages := (totalCount + int64(pageSize) - 1) / int64(pageSize)
	currentPage := 1
	hasNextPage := currentPage < int(totalPages)
	hasPrevPage := currentPage > 1

	t.Logf("Pagination info:")
	t.Logf("  Current page: %d", currentPage)
	t.Logf("  Total pages: %d", totalPages)
	t.Logf("  Has next page: %v", hasNextPage)
	t.Logf("  Has prev page: %v", hasPrevPage)

	require.Equal(t, int64(2), totalPages)
	require.True(t, hasNextPage)
	require.False(t, hasPrevPage)
}

// TestFilteredPagination 演示带条件过滤的分页查询
func TestFilteredPagination(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Post{}))

	// 1. 按分类分页查询
	techPosts, techCount, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Category.Eq("tech")).Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ViewCount.OrderByBottle("desc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 3},
	)
	require.NoError(t, err)
	require.Greater(t, len(techPosts), 0)
	t.Logf("Tech posts: %d, total: %d", len(techPosts), techCount)

	// 2. 按作者分页查询
	authorPosts, authorCount, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.AuthorID.Eq(1)).Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.LikeCount.OrderByBottle("desc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 5},
	)
	require.NoError(t, err)
	require.Greater(t, len(authorPosts), 0)
	t.Logf("Author 1 posts: %d, total: %d", len(authorPosts), authorCount)

	// 3. 按浏览量范围分页查询
	popularPosts, popularCount, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.ViewCount.Gt(100)).Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ViewCount.OrderByBottle("desc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 5},
	)
	require.NoError(t, err)
	require.Greater(t, len(popularPosts), 0)
	t.Logf("Popular posts (>100 views): %d, total: %d", len(popularPosts), popularCount)

	// 验证排序正确性
	for i := 1; i < len(popularPosts); i++ {
		require.GreaterOrEqual(t, popularPosts[i-1].ViewCount, popularPosts[i].ViewCount)
	}
}

// TestMultipleSortPagination 演示多字段排序的分页查询
func TestMultipleSortPagination(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Post{}))

	// 1. 多字段排序：先按分类，再按浏览量降序
	posts, count, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.Category.OrderByBottle("asc").Ob(cls.ViewCount.Ob("desc"))
		},
		&gormrepo.Pagination{Offset: 0, Limit: 8},
	)
	require.NoError(t, err)
	require.Len(t, posts, 8)
	require.Greater(t, count, int64(0))

	t.Logf("Multi-sort pagination: %d posts, total: %d", len(posts), count)
	for i, post := range posts {
		t.Logf("  %d. [%s] %s (views: %d)", i+1, post.Category, post.Title, post.ViewCount)
	}

	// 2. 复杂排序：先按点赞数降序，再按ID升序
	posts2, _, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.LikeCount.OrderByBottle("desc").Ob(cls.ID.Ob("asc"))
		},
		&gormrepo.Pagination{Offset: 0, Limit: 5},
	)
	require.NoError(t, err)
	require.Len(t, posts2, 5)

	t.Logf("Like-based sorting:")
	for i, post := range posts2 {
		t.Logf("  %d. %s (likes: %d, ID: %d)", i+1, post.Title, post.LikeCount, post.ID)
	}

	// 验证排序正确性
	for i := 1; i < len(posts2); i++ {
		if posts2[i-1].LikeCount == posts2[i].LikeCount {
			require.Less(t, posts2[i-1].ID, posts2[i].ID)
		} else {
			require.Greater(t, posts2[i-1].LikeCount, posts2[i].LikeCount)
		}
	}
}

// TestPaginationEdgeCases 演示分页查询的边界情况
func TestPaginationEdgeCases(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Post{}))

	// 1. 超出范围的分页查询
	posts, count, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("asc")
		},
		&gormrepo.Pagination{Offset: 100, Limit: 10}, // 超出数据范围
	)
	require.NoError(t, err)
	require.Len(t, posts, 0)            // 应该返回空数组
	require.Greater(t, count, int64(0)) // 但总数应该大于0
	t.Logf("Out of range pagination: %d posts, total: %d", len(posts), count)

	// 2. 大页面尺寸
	largePage, _, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("asc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 100}, // 大于实际数据量
	)
	require.NoError(t, err)
	require.Equal(t, 20, len(largePage)) // 应该返回所有已发布的帖子
	t.Logf("Large page size: %d posts returned", len(largePage))

	// 3. 零页面尺寸处理
	zeroPosts, zeroCount, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("asc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 0},
	)
	require.NoError(t, err)
	require.Len(t, zeroPosts, 0)
	require.Greater(t, zeroCount, int64(0))
	t.Logf("Zero limit pagination: %d posts, total: %d", len(zeroPosts), zeroCount)

	// 4. 不存在的条件分页
	nonexistentPosts, nonexistentCount, err := repo.FindPageAndCount(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Category.Eq("nonexistent"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("asc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 10},
	)
	require.NoError(t, err)
	require.Len(t, nonexistentPosts, 0)
	require.Equal(t, int64(0), nonexistentCount)
	t.Logf("Nonexistent condition: %d posts, total: %d", len(nonexistentPosts), nonexistentCount)
}

// TestPaginationPerformance 演示分页查询的性能相关功能
func TestPaginationPerformance(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Post{}))

	// 1. 仅获取数据不计算总数（更快）
	fastPosts, err := repo.FindPage(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published"))
		},
		func(cls *models.PostColumns) gormcnm.OrderByBottle {
			return cls.ID.OrderByBottle("desc")
		},
		&gormrepo.Pagination{Offset: 0, Limit: 5},
	)
	require.NoError(t, err)
	require.Len(t, fastPosts, 5)
	t.Logf("Fast pagination (no count): %d posts", len(fastPosts))

	// 2. 限制结果数量查询
	limitedPosts, err := repo.FindN(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Status.Eq("published")).Order(cls.ViewCount.Ob("desc").Ox())
		},
		3, // 限制3条记录
	)
	require.NoError(t, err)
	require.Len(t, limitedPosts, 3)
	t.Logf("Limited query: %d posts", len(limitedPosts))

	// 验证是按浏览量降序排列
	for i := 1; i < len(limitedPosts); i++ {
		require.GreaterOrEqual(t, limitedPosts[i-1].ViewCount, limitedPosts[i].ViewCount)
	}

	// 3. 使用 FindC 方法同时获取数据和总数
	customPosts, customCount, err := repo.FindC(
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Where(cls.Category.Eq("tech"))
		},
		func(db *gorm.DB, cls *models.PostColumns) *gorm.DB {
			return db.Order(cls.LikeCount.Ob("desc").Ox()).Limit(3).Offset(0)
		},
	)
	require.NoError(t, err)
	require.Greater(t, len(customPosts), 0)
	require.Greater(t, customCount, int64(0))
	t.Logf("Custom FindC: %d posts, total: %d", len(customPosts), customCount)
}
