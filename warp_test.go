package gormrepo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

func TestRepo_GormRepo_Morm(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		var nickname string
		require.NoError(t, repo.Repo(caseDB).Morm().WhereE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		}, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-1-nickname", nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		var nickname string
		require.NoError(t, repo.Repo(db).Morm().WhereE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		}, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-2-nickname", nickname)
		return nil
	}))
}
