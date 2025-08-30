package example02_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example02/internal/models"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	must.Done(db.AutoMigrate(&models.Account{}))

	caseDB = db
	m.Run()
}

func TestAccount(t *testing.T) {
	var db = caseDB
	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &models.Account{}))

	account1 := &models.Account{
		Username: "abc",
		Password: "123",
		Nickname: "xyz",
	}
	account2 := &models.Account{
		Username: "aaa",
		Password: "111",
		Nickname: "xxx",
	}

	require.NoError(t, db.Create(account1).Error)
	require.NoError(t, db.Create(account2).Error)

	resA, err := repo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("abc"))
	})
	require.NoError(t, err)
	require.NotNil(t, resA)
	require.Equal(t, "abc", resA.Username)
	t.Log("select res.username:", resA.Username)

	resB, err := repo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("aaa"))
	})
	require.NoError(t, err)
	require.NotNil(t, resB)
	require.Equal(t, "aaa", resB.Username)
	t.Log("select res.username:", resB.Username)
}
