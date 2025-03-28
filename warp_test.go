package gormrepo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

func TestGormRepo_Gorm(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	var account Account
	require.NoError(t, repo.Repo(caseDB).Gorm().First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	}, &account).Error)
	require.Equal(t, "demo-1-nickname", account.Nickname)
}

func TestGormRepo_Morm(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		var nickname string
		require.NoError(t, repo.Repo(caseDB).Morm().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username")).Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-1-nickname", nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		var nickname string
		require.NoError(t, repo.Repo(db).Morm().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username")).Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-2-nickname", nickname)
		return nil
	}))
}

func TestGormWrap_Morm(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		var nickname string
		require.NoError(t, repo.Gorm(caseDB).Morm().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username")).Select(string(cls.Nickname)).First(&nickname)
		}).Error)
		require.Equal(t, "demo-1-nickname", nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		var nickname string
		require.NoError(t, repo.Gorm(db).Morm().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username")).Select(string(cls.Nickname)).First(&nickname)
		}).Error)
		require.Equal(t, "demo-2-nickname", nickname)
		return nil
	}))
}

func TestGormWrap_Repo(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	res, err := repo.Gorm(caseDB).Repo().First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})
	require.NoError(t, err)
	require.Equal(t, "demo-1-nickname", res.Nickname)
}
