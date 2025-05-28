package gormrepo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

func TestRepo_Repo(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

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

func TestRepo_Gorm(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

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

func TestRepo_With(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	ctx := context.Background()
	res, err := repo.With(ctx, caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})
	require.NoError(t, err)
	require.Equal(t, "demo-1-nickname", res.Nickname)
}

func TestRepo_Wrap(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	ctx := context.Background()
	var account Account
	require.NoError(t, repo.Wrap(ctx, caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	}, &account).Error)
	require.Equal(t, "demo-1-nickname", account.Nickname)
}
