package gormrepo_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestGormWrap_First(t *testing.T) {
	repo := gormrepo.NewGormWrap(gormrepo.Umc(caseDB, &Account{}))

	{
		var account Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		}, &account).Error)
		require.Equal(t, "demo-1-nickname", account.Nickname)
	}

	{
		var account Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		}, &account).Error)
		require.Equal(t, "demo-2-nickname", account.Nickname)
	}
}

func TestGormWrap_Where(t *testing.T) {
	repo := gormrepo.NewGormWrap(gormrepo.Umc(caseDB, &Account{}))

	db := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})
	var account Account
	require.NoError(t, db.First(&account).Error)
	require.Equal(t, "demo-1-nickname", account.Nickname)
}

func TestGormWrap_Find(t *testing.T) {
	repo := gormrepo.NewGormWrap(gormrepo.Use(caseDB, &Account{}))

	var accounts []*Account
	require.NoError(t, repo.Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Like("demo-%-username"))
	}, &accounts).Error)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestGormWrap_Update(t *testing.T) {
	username := uuid.New().String()
	require.NoError(t, caseDB.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(caseDB, &Account{}))

	newNickname := uuid.New().String()
	require.NoError(t, repo.Update(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) (string, interface{}) {
		return cls.Nickname.Kv(newNickname)
	}).Error)

	var res Account
	require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, &res).Error)
	require.Equal(t, newNickname, res.Nickname)
}

func TestGormWrap_Updates(t *testing.T) {
	username := uuid.New().String()
	require.NoError(t, caseDB.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(caseDB, &Account{}))

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.Updates(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) map[string]interface{} {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword)).
			AsMap() // Convert to map[string]interface{} // 转换为 map[string]interface{} 类型
	}).Error)

	var res Account
	require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, &res).Error)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}

// TestGormWrap_UpdatesM tests fluent API without AsMap() call
// TestGormWrap_UpdatesM 测试流畅的 API，无需调用 AsMap()
func TestGormWrap_UpdatesM(t *testing.T) {
	username := uuid.New().String()
	require.NoError(t, caseDB.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(caseDB, &Account{}))

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.UpdatesM(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) gormcnm.ColumnValueMap {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword))
		// No AsMap() needed! // 不需要 AsMap()！
	}).Error)

	var res Account
	require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, &res).Error)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}

func TestGormWrap_Invoke(t *testing.T) {
	username := uuid.New().String()
	require.NoError(t, caseDB.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(caseDB, &Account{}))

	newNickname := uuid.New().String()
	require.NoError(t, repo.Mold().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username)).Update(cls.Nickname.Kv(newNickname))
	}).Error)

	var account Account
	require.NoError(t, repo.Mold().Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username)).First(&account)
	}).Error)
	require.Equal(t, newNickname, account.Nickname)
}

func TestGormWrap_Create(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account).Error)
	require.NotZero(t, account.ID)

	var res Account
	require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, &res).Error)
	require.Equal(t, username, res.Username)
}

func TestGormWrap_Save(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Save(account).Error)
	require.NotZero(t, account.ID)

	newNickname := uuid.New().String()
	account.Nickname = newNickname
	require.NoError(t, repo.Save(account).Error)

	var res Account
	require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, &res).Error)
	require.Equal(t, newNickname, res.Nickname)
}

func TestGormWrap_Delete(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account).Error)

	require.NoError(t, repo.Delete(account).Error)

	var count int64
	require.NoError(t, repo.Mold().Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(account.ID))
	}).Count(&count).Error)
	require.Equal(t, int64(0), count)
}

func TestGormWrap_DeleteW(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account).Error)

	require.NoError(t, repo.DeleteW(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}).Error)

	var count int64
	require.NoError(t, repo.Mold().Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}).Count(&count).Error)
	require.Equal(t, int64(0), count)
}

func TestGormWrap_DeleteM(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account).Error)

	require.NoError(t, repo.DeleteM(account, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}).Error)

	var count int64
	require.NoError(t, repo.Mold().Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(account.ID))
	}).Count(&count).Error)
	require.Equal(t, int64(0), count)
}

// TestGormWrap_UpdatesO tests update using object primary key as condition
// O = Object, updates record located by object's primary key
//
// TestGormWrap_UpdatesO 测试使用 object 主键作为条件的更新
// O = Object，通过 object 的主键定位要更新的记录
func TestGormWrap_UpdatesO(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account).Error)

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.UpdatesO(account, func(cls *AccountColumns) gormcnm.ColumnValueMap {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword))
	}).Error)

	var res Account
	require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, &res).Error)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}

// TestGormWrap_UpdatesC tests update using combined conditions: object primary key plus where clause
// C = Combined, uses both object primary key and where conditions for precise targeting
//
// TestGormWrap_UpdatesC 测试使用组合条件的更新：object 主键加上 where 子句
// C = Combined，同时使用 object 主键和 where 条件进行精确定位
func TestGormWrap_UpdatesC(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account).Error)

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.UpdatesC(account, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) gormcnm.ColumnValueMap {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword))
	}).Error)

	var res Account
	require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, &res).Error)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}
