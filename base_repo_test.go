package gormrepo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

// TestRepo_Repo tests BaseRepo.Repo method to create GormRepo with db connection
// TestRepo_Repo 测试 BaseRepo.Repo 方法创建带数据库连接的 GormRepo
func TestRepo_Repo(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{}))

	t.Run("demo1", func(t *testing.T) {
		res, err := repo.Repo(caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	})

	t.Run("demo2", func(t *testing.T) {
		require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
			res, err := repo.Repo(db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
				return db.Where(cls.Username.Eq("demo-2-username"))
			})
			require.NoError(t, err)
			require.Equal(t, "demo-2-nickname", res.Nickname)
			return nil
		}))
	})
}

// TestRepo_Gorm tests BaseRepo.Gorm method to create GormWrap with db connection
// TestRepo_Gorm 测试 BaseRepo.Gorm 方法创建带数据库连接的 GormWrap
func TestRepo_Gorm(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{}))

	t.Run("demo1", func(t *testing.T) {
		var account Account
		require.NoError(t, repo.Gorm(caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		}, &account).Error)
		require.Equal(t, "demo-1-nickname", account.Nickname)
	})

	t.Run("demo2", func(t *testing.T) {
		require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
			var account Account
			require.NoError(t, repo.Gorm(db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
				return db.Where(cls.Username.Eq("demo-2-username"))
			}, &account).Error)
			require.Equal(t, "demo-2-nickname", account.Nickname)
			return nil
		}))
	})
}

// TestRepo_With tests BaseRepo.With method to create GormRepo with context
// TestRepo_With 测试 BaseRepo.With 方法创建带上下文的 GormRepo
func TestRepo_With(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{}))

	ctx := context.Background()
	res, err := repo.With(ctx, caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})
	require.NoError(t, err)
	require.Equal(t, "demo-1-nickname", res.Nickname)
}

// TestRepo_Wrap tests BaseRepo.Wrap method to create GormWrap with context
// TestRepo_Wrap 测试 BaseRepo.Wrap 方法创建带上下文的 GormWrap
func TestRepo_Wrap(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{}))

	ctx := context.Background()
	var account Account
	require.NoError(t, repo.Wrap(ctx, caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	}, &account).Error)
	require.Equal(t, "demo-1-nickname", account.Nickname)
}
