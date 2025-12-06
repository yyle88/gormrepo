package gormrepo_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

// TestMain sets up the test database and runs all tests
// TestMain 设置测试数据库并运行所有测试
func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	must.Done(db.AutoMigrate(&Account{}))

	done.Done(db.Save(&Account{
		Model:    gorm.Model{},
		Username: "demo-1-username",
		Password: "demo-1-password",
		Nickname: "demo-1-nickname",
	}).Error)
	done.Done(db.Save(&Account{
		Model:    gorm.Model{},
		Username: "demo-2-username",
		Password: "demo-2-password",
		Nickname: "demo-2-nickname",
	}).Error)

	caseDB = db
	m.Run()
}

// TestGormRepo_First tests the First method to find a single record
// TestGormRepo_First 测试 First 方法查找单条记录
func TestGormRepo_First(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))

	{
		res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	{
		res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
	}
}

// TestGormRepo_FirstE tests the FirstE method with structured error handling
// TestGormRepo_FirstE 测试带结构化错误处理的 FirstE 方法
func TestGormRepo_FirstE(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))

	{
		res, erb := repo.FirstE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.Nil(t, erb)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	{
		res, erb := repo.FirstE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-x-username"))
		})
		require.NotNil(t, erb)
		require.ErrorIs(t, erb.Cause, gorm.ErrRecordNotFound)
		require.True(t, erb.NotExist)
		require.Nil(t, res)
	}
}

// TestGormRepo_Where tests the Where method to apply custom conditions
// TestGormRepo_Where 测试 Where 方法应用自定义条件
func TestGormRepo_Where(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))

	{
		var nicknames []string
		db := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).
				Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"})).
				Distinct(cls.Nickname.Name())
		})
		require.NoError(t, db.Find(&nicknames).Error)
		t.Log(nicknames)
		require.Len(t, nicknames, 2)
		sort.Strings(nicknames)
		require.Equal(t, []string{"demo-1-nickname", "demo-2-nickname"}, nicknames)
	}
	{
		var nickname string
		db := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).
				Select("MIN(" + cls.Nickname.Name() + ")").
				Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
		})
		require.NoError(t, db.First(&nickname).Error)
		t.Log(nickname)
		require.Equal(t, "demo-1-nickname", nickname)
	}
}

// TestGormRepo_Exist tests the Exist method to check record existence
// TestGormRepo_Exist 测试 Exist 方法检查记录是否存在
func TestGormRepo_Exist(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))

	{
		exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.True(t, exist)
	}

	{
		exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-x-username"))
		})
		require.NoError(t, err)
		require.False(t, exist)
	}
}

// TestGormRepo_Find tests the Find method to retrieve multiple records
// TestGormRepo_Find 测试 Find 方法检索多条记录
func TestGormRepo_Find(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	accounts, err := repo.Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Like("demo-%-username"))
	})
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

// TestGormRepo_FindN tests the FindN method to retrieve limited records
// TestGormRepo_FindN 测试 FindN 方法检索有限数量的记录
func TestGormRepo_FindN(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	accounts, err := repo.FindN(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	}, 2)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

// TestGormRepo_FindC tests the FindC method with custom paging and count
// TestGormRepo_FindC 测试带自定义分页和计数的 FindC 方法
func TestGormRepo_FindC(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	accounts, count, err := repo.FindC(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	}, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Offset(1).Limit(2).Order(cls.Username.OrderByBottle("asc").Orders())
	})
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 1)
	require.Equal(t, "demo-2-username", accounts[0].Username)
	require.Equal(t, "demo-2-nickname", accounts[0].Nickname)
	require.Equal(t, int64(2), count)
	t.Log(neatjsons.S(accounts))
}

// TestGormRepo_FindPageAndCount tests paginated search with total count
// TestGormRepo_FindPageAndCount 测试带总数的分页搜索
func TestGormRepo_FindPageAndCount(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	accounts, count, err := repo.FindPageAndCount(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	}, func(cls *AccountColumns) gormcnm.OrderByBottle {
		return cls.Username.OrderByBottle("asc")
	}, &gormrepo.Pagination{
		Offset: 1,
		Limit:  2,
	})
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 1)
	require.Equal(t, "demo-2-username", accounts[0].Username)
	require.Equal(t, "demo-2-nickname", accounts[0].Nickname)
	require.Equal(t, int64(2), count)
	t.Log(neatjsons.S(accounts))
}

func TestGormRepo_FindPage(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	accounts, err := repo.FindPage(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	}, func(cls *AccountColumns) gormcnm.OrderByBottle {
		return cls.Username.OrderByBottle("desc").Ob(cls.Nickname.Ob("asc"))
	}, &gormrepo.Pagination{
		Offset: 0,
		Limit:  2,
	})
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 2)
	require.Equal(t, "demo-2-username", accounts[0].Username)
	require.Equal(t, "demo-2-nickname", accounts[0].Nickname)
	require.Equal(t, "demo-1-username", accounts[1].Username)
	require.Equal(t, "demo-1-nickname", accounts[1].Nickname)
	t.Log(neatjsons.S(accounts))
}

func TestGormRepo_Count(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	count, err := repo.Count(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), count)
}

func TestGormRepo_Update(t *testing.T) {
	username := uuid.New().String()
	require.NoError(t, caseDB.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	newNickname := uuid.New().String()
	require.NoError(t, repo.Update(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) (string, interface{}) {
		return cls.Nickname.Kv(newNickname)
	}))

	res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
}

func TestGormRepo_Updates(t *testing.T) {
	username := uuid.New().String()
	require.NoError(t, caseDB.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.Updates(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) map[string]interface{} {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword)).
			AsMap() // Convert to map[string]interface{} // 转换为 map[string]interface{} 类型
	}))

	res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}

// TestGormRepo_UpdatesM tests fluent API without AsMap() call
// TestGormRepo_UpdatesM 测试流畅的 API，无需调用 AsMap()
func TestGormRepo_UpdatesM(t *testing.T) {
	username := uuid.New().String()
	require.NoError(t, caseDB.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.UpdatesM(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) gormcnm.ColumnValueMap {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword))
		// No AsMap() needed! // 不需要 AsMap()！
	}))

	res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}

func TestGormRepo_Invoke(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))

	{
		var passwords []string
		require.NoError(t, repo.Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"})).Distinct(cls.Password.Name()).Find(&passwords)
		}))
		t.Log(passwords)
		require.Len(t, passwords, 2)
		sort.Strings(passwords)
		require.Equal(t, []string{"demo-1-password", "demo-2-password"}, passwords)
	}
	{
		var nickname string
		require.NoError(t, repo.Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"})).Select("MAX(" + cls.Nickname.Name() + ")").First(&nickname)
		}))
		t.Log(nickname)
		require.Equal(t, "demo-2-nickname", nickname)
	}
}

func TestGormRepo_Create(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account))
	require.NotZero(t, account.ID)

	res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, username, res.Username)
}

func TestGormRepo_Save(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Save(account))
	require.NotZero(t, account.ID)

	newNickname := uuid.New().String()
	account.Nickname = newNickname
	require.NoError(t, repo.Save(account))

	res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
}

func TestGormRepo_Delete(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account))

	require.NoError(t, repo.Delete(account))

	exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(account.ID))
	})
	require.NoError(t, err)
	require.False(t, exist)
}

func TestGormRepo_DeleteW(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account))

	require.NoError(t, repo.DeleteW(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}))

	exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.False(t, exist)
}

func TestGormRepo_DeleteM(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account))

	require.NoError(t, repo.DeleteM(account, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}))

	exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.ID.Eq(account.ID))
	})
	require.NoError(t, err)
	require.False(t, exist)
}

// TestGormRepo_UpdatesO tests update using object primary key as condition
// O = Object, updates record located by object's primary key
//
// TestGormRepo_UpdatesO 测试使用 object 主键作为条件的更新
// O = Object，通过 object 的主键定位要更新的记录
func TestGormRepo_UpdatesO(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account))

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.UpdatesO(account, func(cls *AccountColumns) gormcnm.ColumnValueMap {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword))
	}))

	res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}

// TestGormRepo_UpdatesC tests update using combined conditions: object primary key plus where clause
// C = Combined, uses both object primary key and where conditions for precise targeting
//
// TestGormRepo_UpdatesC 测试使用组合条件的更新：object 主键加上 where 子句
// C = Combined，同时使用 object 主键和 where 条件进行精确定位
func TestGormRepo_UpdatesC(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	username := uuid.New().String()
	account := newAccount(username)
	require.NoError(t, repo.Create(account))

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	require.NoError(t, repo.UpdatesC(account, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) gormcnm.ColumnValueMap {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword))
	}))

	res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}
