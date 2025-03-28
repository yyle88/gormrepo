package gormrepo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

func TestRepo_Repo(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		res, err := repo.Repo(caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		res, err := repo.Repo(db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
		return nil
	}))
}

func TestRepo_Gorm(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		var account Account
		require.NoError(t, repo.Gorm(caseDB).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		}, &account).Error)
		require.Equal(t, "demo-1-nickname", account.Nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		var account Account
		require.NoError(t, repo.Gorm(db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		}, &account).Error)
		require.Equal(t, "demo-2-nickname", account.Nickname)
		return nil
	}))
}
