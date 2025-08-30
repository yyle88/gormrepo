package example01_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/internal/examples/example01/internal/models"
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

	done.Done(db.AutoMigrate(&models.Example{}))

	caseDB = db
	m.Run()
}

func TestExample(t *testing.T) {
	var db = caseDB

	example1 := &models.Example{
		ID:        0,
		Name:      "aaa",
		Age:       1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	example2 := &models.Example{
		ID:        0,
		Name:      "bbb",
		Age:       2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	require.NoError(t, db.Create(example1).Error)
	require.NoError(t, db.Create(example2).Error)

	var resA models.Example
	if cls := gormclass.Cls(&models.Example{}); cls.OK() {
		require.NoError(t, db.Table(resA.TableName()).Where(cls.Name.Eq("aaa")).First(&resA).Error)
		require.Equal(t, "aaa", resA.Name)
	}
	t.Log("select res.name:", resA.Name)

	var maxAge int
	if one, cls := gormclass.Use(&models.Example{}); cls.OK() {
		require.NoError(t, db.Model(one).Where(cls.Age.Gt(0)).Select(cls.Age.COALESCE().MaxStmt("age_alias")).First(&maxAge).Error)
		require.Equal(t, 2, maxAge)
	}
	t.Log("max_age:", maxAge)

	if one, cls := gormclass.Use(&models.Example{}); cls.OK() {
		require.NoError(t, db.Model(one).Where(cls.Name.Eq("bbb")).Update(cls.Age.Kv(18)).Error)
		require.Equal(t, 18, one.Age)
	}

	var resB models.Example
	if cls := resB.Columns(); cls.OK() {
		require.NoError(t, db.Table(resB.TableName()).Where(cls.Name.Eq("bbb")).Update(cls.Age.KeAdd(2)).Error)

		require.NoError(t, db.Table(resB.TableName()).Where(cls.Name.Eq("bbb")).First(&resB).Error)
		require.Equal(t, 20, resB.Age)
	}
	t.Log(resB)
}
