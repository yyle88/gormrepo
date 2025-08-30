package example05_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example05/internal/models"
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

	done.Done(db.AutoMigrate(&models.Article{}))

	// 插入测试文章
	articles := []*models.Article{
		{Title: "Go语言入门", Content: "Go语言基础教程", AuthorID: 1, Status: "draft", ViewCount: 0, LikeCount: 0},
		{Title: "数据结构详解", Content: "数据结构和算法", AuthorID: 2, Status: "published", ViewCount: 100, LikeCount: 10},
		{Title: "Web开发实战", Content: "全栈Web开发", AuthorID: 1, Status: "published", ViewCount: 200, LikeCount: 25},
	}
	for _, article := range articles {
		if article.Status == "published" {
			now := time.Now()
			article.PublishedAt = &now
		}
		done.Done(db.Create(article).Error)
	}

	testDB = db
	m.Run()
}

// TestSingleFieldUpdate 演示单字段更新
func TestSingleFieldUpdate(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Article{}))

	// 1. 更新单个字段 - 修改标题
	err := repo.Update(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Title.Eq("Go语言入门"))
	}, func(cls *models.ArticleColumns) (string, interface{}) {
		return cls.Title.Kv("Go语言进阶教程")
	})
	require.NoError(t, err)

	// 验证更新结果
	article, err := repo.First(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.AuthorID.Eq(1)).Where(cls.Status.Eq("draft"))
	})
	require.NoError(t, err)
	require.Equal(t, "Go语言进阶教程", article.Title)
	t.Logf("Updated title: %s", article.Title)

	// 2. 更新计数字段 - 增加浏览量
	err = repo.Update(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Title.Eq("数据结构详解"))
	}, func(cls *models.ArticleColumns) (string, interface{}) {
		return cls.ViewCount.Kv(150)
	})
	require.NoError(t, err)

	// 验证计数更新
	article, err = repo.First(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Title.Eq("数据结构详解"))
	})
	require.NoError(t, err)
	require.Equal(t, 150, article.ViewCount)
	t.Logf("Updated view count: %d", article.ViewCount)
}

// TestMultipleFieldsUpdate 演示多字段批量更新
func TestMultipleFieldsUpdate(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Article{}))

	// 1. 批量更新多个字段 - 发布文章
	now := time.Now()
	err := repo.Updates(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("draft"))
	}, func(cls *models.ArticleColumns) map[string]interface{} {
		return cls.
			Kw(cls.Status.Kv("published")).
			Kw(cls.PublishedAt.Kv(&now)).
			AsMap()
	})
	require.NoError(t, err)

	// 验证批量更新
	publishedArticles, err := repo.Find(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("published"))
	})
	require.NoError(t, err)
	require.Len(t, publishedArticles, 3) // 现在应该有3篇发布的文章
	t.Logf("Published articles count: %d", len(publishedArticles))

	// 2. 条件批量更新 - 增加特定作者文章的点赞数
	err = repo.Updates(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.AuthorID.Eq(1))
	}, func(cls *models.ArticleColumns) map[string]interface{} {
		return cls.
			Kw(cls.LikeCount.Kv(50)).
			AsMap()
	})
	require.NoError(t, err)

	// 验证作者文章更新
	authorArticles, err := repo.Find(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.AuthorID.Eq(1))
	})
	require.NoError(t, err)
	for _, article := range authorArticles {
		require.Equal(t, 50, article.LikeCount)
	}
	t.Logf("Author 1 articles updated, like count: 50")
}

// TestConditionalUpdate 演示条件更新
func TestConditionalUpdate(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Article{}))

	// 1. 基于浏览量的条件更新 - 热门文章加标记
	err := repo.Updates(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.ViewCount.Gt(100))
	}, func(cls *models.ArticleColumns) map[string]interface{} {
		return map[string]interface{}{
			cls.Title.Name(): gorm.Expr("CONCAT(title, ' [热门]')"),
		}
	})
	require.NoError(t, err)

	// 验证条件更新
	hotArticles, err := repo.Find(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Title.Like("%[热门]%"))
	})
	require.NoError(t, err)
	require.Greater(t, len(hotArticles), 0)
	t.Logf("Hot articles marked: %d", len(hotArticles))

	// 2. 基于状态和时间的复合条件更新
	err = repo.Update(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("published")).Where(cls.PublishedAt.IsNotNull())
	}, func(cls *models.ArticleColumns) (string, interface{}) {
		return cls.Content.Kv("内容已更新")
	})
	require.NoError(t, err)

	// 验证复合条件更新
	updatedCount, err := repo.Count(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Content.Eq("内容已更新"))
	})
	require.NoError(t, err)
	require.Greater(t, updatedCount, int64(0))
	t.Logf("Updated articles with new content: %d", updatedCount)
}

// TestIncrementUpdate 演示递增更新
func TestIncrementUpdate(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Article{}))

	// 获取更新前的数据
	beforeArticle, err := repo.First(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.Title.Like("%Web开发%"))
	})
	require.NoError(t, err)
	originalViewCount := beforeArticle.ViewCount
	originalLikeCount := beforeArticle.LikeCount

	// 1. 使用表达式递增浏览量
	err = repo.Updates(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(beforeArticle.ID))
	}, func(cls *models.ArticleColumns) map[string]interface{} {
		key, expr := cls.ViewCount.KeAdd(1)
		return map[string]interface{}{key: expr}
	})
	require.NoError(t, err)

	// 2. 同时递增多个计数字段
	err = repo.Updates(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(beforeArticle.ID))
	}, func(cls *models.ArticleColumns) map[string]interface{} {
		viewKey, viewExpr := cls.ViewCount.KeAdd(5)
		likeKey, likeExpr := cls.LikeCount.KeAdd(2)
		return map[string]interface{}{viewKey: viewExpr, likeKey: likeExpr}
	})
	require.NoError(t, err)

	// 验证递增结果
	afterArticle, err := repo.First(func(db *gorm.DB, cls *models.ArticleColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(beforeArticle.ID))
	})
	require.NoError(t, err)
	require.Equal(t, originalViewCount+1+5, afterArticle.ViewCount) // +1 + 5
	require.Equal(t, originalLikeCount+2, afterArticle.LikeCount)   // +2
	t.Logf("Incremental update: view_count %d->%d, like_count %d->%d",
		originalViewCount, afterArticle.ViewCount,
		originalLikeCount, afterArticle.LikeCount)
}
