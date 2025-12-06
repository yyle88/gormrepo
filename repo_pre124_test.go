//go:build !go1.24

package gormrepo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

// TestRepoVersion shows which Repo version is being used
// 显示当前使用的 Repo 版本
func TestRepoVersion(t *testing.T) {
	t.Log("Running pre-1.24 test: struct embedding approach (no type alias)")
	t.Log("运行 1.24 之前测试：结构体嵌入方案（无类型别名）")
}

// TestRepo_Base tests Repo with struct embedding (Go < 1.24)
// Go 1.24 之前使用结构体嵌入方案，此测试验证兼容版本
func TestRepo_Base(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	ctx := context.Background()
	db := caseDB
	{
		res, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	{
		res, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
	}
}
