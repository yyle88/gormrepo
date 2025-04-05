package example2

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/internal/examples/example2/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(&models.Account{}))

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

	var resA models.Account
	if one, cls := models.UmcV2(&models.Account{}); cls.OK() {
		require.NoError(t, db.Table(one.TableName()).Where(cls.Username.Eq("abc")).First(&resA).Error)
		require.Equal(t, "abc", resA.Username)
	}
	t.Log("select res.username:", resA.Username)

	var resB models.Account
	if one, cls := models.UmcV2(&models.Account{}); cls.OK() {
		require.NoError(t, db.Table(one.TableName()).Where(cls.Username.Eq("aaa")).First(&resB).Error)
		require.Equal(t, "aaa", resB.Username)
	}
	t.Log("select res.username:", resB.Username)
}
