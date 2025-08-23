package example02_test

import (
	"database/sql" // Added for sqlDB.Close()
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"           // Added
	"github.com/yyle88/gormrepo/gormclass" // Keep gormclass for TestCompare
	"github.com/yyle88/gormrepo/internal/examples/example02/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer func(sqlDB *sql.DB) { // Added defer for sqlDB.Close()
		err := sqlDB.Close()
		if err != nil {
			panic(err)
		}
	}(sqlDB)

	err = db.AutoMigrate(&models.Account{}) // Changed done.Done to require.NoError
	if err != nil {
		panic(err)
	}

	caseDB = db
	m.Run()
}

func TestCompare(t *testing.T) {
	for i := 0; i < 10; i++ {
		const count = 100000
		one := &models.Account{}
		{
			stm := time.Now()
			for i := 0; i < count; i++ {
				gormclass.Umc(one) //由于不使用缓存所以每次都要计算，但实际也能看到影响是特别小的
			}
			t.Log("--0--", time.Since(stm))
		}
		{
			stm := time.Now()
			for i := 0; i < count; i++ {
				models.UmcV1(one) //由于使用缓存所以这里只计算一次，目前看来性能提升幅度比较有限，而且涉及到DB的操作瓶颈都在DB那边
			}
			t.Log("--1--", time.Since(stm))
		}
		{
			stm := time.Now()
			for i := 0; i < count; i++ {
				models.UmcV2(one) //由于使用缓存所以这里只计算一次，目前看来性能提升幅度比较有限，而且涉及到DB的操作瓶颈都在DB那边
			}
			t.Log("--2--", time.Since(stm))
		}
		{
			stm := time.Now()
			for i := 0; i < count; i++ {
				models.UmcV3(one) //使用的缓存不同，这两种缓存方案几乎没有性能差异
			}
			t.Log("--3--", time.Since(stm))
		}
	}
}

func TestAccount(t *testing.T) {
	var db = caseDB
	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &models.Account{})) // Added repo initialization

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

	// Refactor First call
	resA, err := repo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("abc"))
	})
	require.NoError(t, err)
	require.NotNil(t, resA)
	require.Equal(t, "abc", resA.Username)
	t.Log("select res.username:", resA.Username)

	// Refactor Second First call
	resB, err := repo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("aaa"))
	})
	require.NoError(t, err)
	require.NotNil(t, resB)
	require.Equal(t, "aaa", resB.Username)
	t.Log("select res.username:", resB.Username)
}
