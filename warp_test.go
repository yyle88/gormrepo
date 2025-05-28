package gormrepo_test

import (
	"context"
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

func TestGormWrap_Repo(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	res, err := repo.Gorm(caseDB).Repo().First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})
	require.NoError(t, err)
	require.Equal(t, "demo-1-nickname", res.Nickname)
}

func TestGormRepo_Mold(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		var nickname string
		require.NoError(t, repo.Repo(caseDB).Mold().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username")).Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-1-nickname", nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		var nickname string
		require.NoError(t, repo.Repo(db).Mold().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username")).Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-2-nickname", nickname)
		return nil
	}))
}

func TestGormWrap_Mold(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		var nickname string
		require.NoError(t, repo.Gorm(caseDB).Mold().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username")).Select(string(cls.Nickname)).First(&nickname)
		}).Error)
		require.Equal(t, "demo-1-nickname", nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		var nickname string
		require.NoError(t, repo.Gorm(db).Mold().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username")).Select(string(cls.Nickname)).First(&nickname)
		}).Error)
		require.Equal(t, "demo-2-nickname", nickname)
		return nil
	}))
}

func TestGormRepo_WithContext(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))
	ctx := context.Background()
	type contextKeyType struct{}
	ctx = context.WithValue(ctx, contextKeyType{}, "value-abc")
	{
		res, err := repo.Repo(caseDB).WithContext(ctx).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			value := ctx.Value(contextKeyType{})
			require.NotNil(t, value)
			rawValue, ok := value.(string)
			require.True(t, ok)
			require.Equal(t, "value-abc", rawValue)

			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}
}

func TestGormWrap_WithContext(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))
	ctx := context.Background()
	type contextKeyType struct{}
	ctx = context.WithValue(ctx, contextKeyType{}, "value-abc")
	{
		var account Account
		require.NoError(t, repo.Gorm(caseDB).WithContext(ctx).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			value := ctx.Value(contextKeyType{})
			require.NotNil(t, value)
			rawValue, ok := value.(string)
			require.True(t, ok)
			require.Equal(t, "value-abc", rawValue)

			return db.Where(cls.Username.Eq("demo-1-username"))
		}, &account).Error)
		require.Equal(t, "demo-1-nickname", account.Nickname)
	}
}
