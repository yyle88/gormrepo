package gormrepo_test

import (
	"context"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/internal/tests"
	"github.com/yyle88/rese"
	"gorm.io/gorm"
)

// TestRepoVersion shows which Repo version is being used
// 显示当前使用的 Repo 版本
func TestRepoVersion(t *testing.T) {
	t.Log("Running generic type alias version (Go 1.24+ native support)")
	t.Log(runtime.Version())
}

// TestRepo_Base tests Repo with generic type alias (Go 1.24+ native support)
// Go 1.24+ 原生支持泛型类型别名，此测试验证类型别名版本
func TestRepo_Base(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	ctx := context.Background()

	t.Run("case-1", func(t *testing.T) {
		res, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	})

	t.Run("case-2", func(t *testing.T) {
		res, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
	})
}
