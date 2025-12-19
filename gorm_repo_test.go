package gormrepo_test

import (
	"sort"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/tests"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// setupDemoData creates demo data in database
// setupDemoData 在数据库中创建演示数据
func setupDemoData(t *testing.T, db *gorm.DB) {
	t.Log("--- setup demo data BEGIN ---")

	t.Log("auto migrate Account table")
	must.Done(db.AutoMigrate(&Account{}))

	account1 := &Account{
		Model:    gorm.Model{},
		Username: "demo-1-username",
		Password: "demo-1-password",
		Nickname: "demo-1-nickname",
	}
	t.Log("save account-1")
	done.Done(db.Save(account1).Error)
	t.Log("account-1:", neatjsons.S(account1))

	account2 := &Account{
		Model:    gorm.Model{},
		Username: "demo-2-username",
		Password: "demo-2-password",
		Nickname: "demo-2-nickname",
	}
	t.Log("save account-2")
	done.Done(db.Save(account2).Error)
	t.Log("account-2:", neatjsons.S(account2))

	t.Log("--- setup demo data READY ---")
}

// TestGormRepo_First tests the First method to find a single record
// TestGormRepo_First 测试 First 方法查找单条记录
func TestGormRepo_First(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Umc(db, &Account{}))

	t.Run("case-1", func(t *testing.T) {
		res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	})

	t.Run("case-2", func(t *testing.T) {
		res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
	})
}

// TestGormRepo_FirstE tests the FirstE method with structured error handling
// TestGormRepo_FirstE 测试带结构化错误处理的 FirstE 方法
func TestGormRepo_FirstE(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Umc(db, &Account{}))

	t.Run("case-1", func(t *testing.T) {
		res, erb := repo.FirstE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.Nil(t, erb)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	})

	t.Run("case-2", func(t *testing.T) {
		res, erb := repo.FirstE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-x-username"))
		})
		require.NotNil(t, erb)
		require.ErrorIs(t, erb.Cause, gorm.ErrRecordNotFound)
		require.True(t, erb.NotExist)
		require.Nil(t, res)
	})
}

// TestGormRepo_Where tests the Where method to apply custom conditions
// TestGormRepo_Where 测试 Where 方法应用自定义条件
func TestGormRepo_Where(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Umc(db, &Account{}))

	t.Run("case-1", func(t *testing.T) {
		var nicknames []string
		dbx := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).
				Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"})).
				Distinct(cls.Nickname.Name())
		})
		require.NoError(t, dbx.Find(&nicknames).Error)
		t.Log(nicknames)
		require.Len(t, nicknames, 2)
		sort.Strings(nicknames)
		require.Equal(t, []string{"demo-1-nickname", "demo-2-nickname"}, nicknames)
	})

	t.Run("case-2", func(t *testing.T) {
		var nickname string
		dbx := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).
				Select("MIN(" + cls.Nickname.Name() + ")").
				Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
		})
		require.NoError(t, dbx.First(&nickname).Error)
		t.Log(nickname)
		require.Equal(t, "demo-1-nickname", nickname)
	})
}

// TestGormRepo_Exist tests the Exist method to check record existence
// TestGormRepo_Exist 测试 Exist 方法检查记录是否存在
func TestGormRepo_Exist(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Umc(db, &Account{}))

	t.Run("case-1", func(t *testing.T) {
		exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.True(t, exist)
	})

	t.Run("case-2", func(t *testing.T) {
		exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-x-username"))
		})
		require.NoError(t, err)
		require.False(t, exist)
	})
}

// TestGormRepo_Find tests the Find method to retrieve multiple records
// TestGormRepo_Find 测试 Find 方法检索多条记录
func TestGormRepo_Find(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

	count, err := repo.Count(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), count)
}

func TestGormRepo_Update(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	username := uuid.New().String()
	require.NoError(t, db.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	username := uuid.New().String()
	require.NoError(t, db.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	username := uuid.New().String()
	require.NoError(t, db.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormRepo(gormrepo.Umc(db, &Account{}))

	t.Run("case-1", func(t *testing.T) {
		var passwords []string
		require.NoError(t, repo.Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"})).Distinct(cls.Password.Name()).Find(&passwords)
		}))
		t.Log(passwords)
		require.Len(t, passwords, 2)
		sort.Strings(passwords)
		require.Equal(t, []string{"demo-1-password", "demo-2-password"}, passwords)
	})

	t.Run("case-2", func(t *testing.T) {
		var nickname string
		require.NoError(t, repo.Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"})).Select("MAX(" + cls.Nickname.Name() + ")").First(&nickname)
		}))
		t.Log(nickname)
		require.Equal(t, "demo-2-nickname", nickname)
	})
}

func TestGormRepo_Create(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

func TestGormRepo_Save(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormRepo_Creates tests batch insert of multiple records
// TestGormRepo_Creates 测试批量插入多条记录
func TestGormRepo_Creates(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

		accounts := []*Account{
			newAccount(uuid.New().String()),
			newAccount(uuid.New().String()),
			newAccount(uuid.New().String()),
		}
		require.NoError(t, repo.Creates(accounts))

		// Verify all records have IDs assigned
		// 验证所有记录都分配了 ID
		for _, account := range accounts {
			require.NotZero(t, account.ID)
		}

		// Verify all records exist in database
		// 验证所有记录都存在于数据库中
		count, err := repo.Count(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.In([]string{
				accounts[0].Username,
				accounts[1].Username,
				accounts[2].Username,
			}))
		})
		require.NoError(t, err)
		require.Equal(t, int64(3), count)
	})
}

// TestGormRepo_Saves tests batch insert or update of multiple records
// TestGormRepo_Saves 测试批量插入或更新多条记录
func TestGormRepo_Saves(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

		// First save - should create new records
		// 首次保存 - 应该创建新记录
		accounts := []*Account{
			newAccount(uuid.New().String()),
			newAccount(uuid.New().String()),
		}
		require.NoError(t, repo.Saves(accounts))

		for _, account := range accounts {
			require.NotZero(t, account.ID)
		}

		// Update nicknames and save again - should update existing records
		// 更新 nickname 并再次保存 - 应该更新现有记录
		newNickname0 := uuid.New().String()
		newNickname1 := uuid.New().String()
		accounts[0].Nickname = newNickname0
		accounts[1].Nickname = newNickname1
		require.NoError(t, repo.Saves(accounts))

		// Verify updates
		// 验证更新
		res0, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.ID.Eq(accounts[0].ID))
		})
		require.NoError(t, err)
		require.Equal(t, newNickname0, res0.Nickname)

		res1, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.ID.Eq(accounts[1].ID))
		})
		require.NoError(t, err)
		require.Equal(t, newNickname1, res1.Nickname)
	})
}

func TestGormRepo_Delete(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

func TestGormRepo_DeleteW(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

func TestGormRepo_DeleteM(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormRepo_UpdatesO tests update using object primary key as condition
// O = Object, updates record located by object's primary key
//
// TestGormRepo_UpdatesO 测试使用 object 主键作为条件的更新
// O = Object，通过 object 的主键定位要更新的记录
func TestGormRepo_UpdatesO(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormRepo_UpdatesC tests update using combined conditions: object primary key plus where clause
// C = Combined, uses both object primary key and where conditions
//
// TestGormRepo_UpdatesC 测试使用组合条件的更新：object 主键加上 where 子句
// C = Combined，同时使用 object 主键和 where 条件进行精确定位
func TestGormRepo_UpdatesC(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormRepo_Clauses tests Clauses method for upsert operations
// TestGormRepo_Clauses 测试 Clauses 方法的 upsert 操作
func TestGormRepo_Clauses(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

		// First create
		// 首次创建
		account1 := &Account{
			Username: uuid.New().String(),
			Password: uuid.New().String(),
			Nickname: uuid.New().String(),
		}
		require.NoError(t, repo.Create(account1))
		require.NotZero(t, account1.ID)

		// Upsert with Clauses - should update nickname
		// 使用 Clauses 进行 upsert - 应该更新 nickname
		account2 := &Account{
			Username: account1.Username,
			Password: uuid.New().String(),
			Nickname: uuid.New().String(),
		}
		cls := account2.Columns()
		require.NoError(t, repo.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: cls.Username.Name()}},
			DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
		}).Create(account2))

		// Verify the nickname was updated
		// 验证 nickname 已被更新
		res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(account1.Username))
		})
		require.NoError(t, err)
		require.Equal(t, account2.Username, res.Username)
		require.Equal(t, account2.Nickname, res.Nickname)
		require.Equal(t, account1.Password, res.Password)
	})
}

// TestGormRepo_Clause tests Clause method for type-safe upsert operations
// TestGormRepo_Clause 测试 Clause 方法的类型安全 upsert 操作
func TestGormRepo_Clause(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

		// First create
		// 首次创建
		account1 := &Account{
			Username: uuid.New().String(),
			Password: uuid.New().String(),
			Nickname: uuid.New().String(),
		}
		require.NoError(t, repo.Create(account1))
		require.NotZero(t, account1.ID)

		// Upsert with Clause - type-safe column names
		// 使用 Clause 进行 upsert - 类型安全的列名
		account2 := &Account{
			Username: account1.Username,
			Password: uuid.New().String(),
			Nickname: uuid.New().String(),
		}
		require.NoError(t, repo.Clause(func(cls *AccountColumns) clause.Expression {
			return clause.OnConflict{
				Columns:   []clause.Column{{Name: cls.Username.Name()}},
				DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
			}
		}).Create(account2))

		// Verify the nickname was updated
		// 验证 nickname 已被更新
		res, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(account1.Username))
		})
		require.NoError(t, err)
		require.Equal(t, account2.Username, res.Username)
		require.Equal(t, account2.Nickname, res.Nickname)
		require.Equal(t, account1.Password, res.Password)
	})
}

// TestGormRepo_Clause_Creates tests Clause + Creates for batch upsert operations
// TestGormRepo_Clause_Creates 测试 Clause + Creates 的批量 upsert 操作
func TestGormRepo_Clause_Creates(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormRepo(gormrepo.Use(db, &Account{}))

		// Prepare nicknames
		// 准备 nicknames
		preNick1 := uuid.New().String()
		preNick2 := uuid.New().String()
		preNick3 := uuid.New().String()

		// First create some accounts
		// 首先创建一些账户
		accounts := []*Account{
			{Username: uuid.New().String(), Password: uuid.New().String(), Nickname: preNick1},
			{Username: uuid.New().String(), Password: uuid.New().String(), Nickname: preNick2},
			{Username: uuid.New().String(), Password: uuid.New().String(), Nickname: preNick3},
		}
		require.NoError(t, repo.Creates(accounts))

		newNick1 := uuid.New().String()
		newNick2 := uuid.New().String()
		newNick4 := uuid.New().String()

		// Batch upsert with Clause - update nicknames
		// 使用 Clause 进行批量 upsert - 更新 nicknames
		upsertAccounts := []*Account{
			{Username: accounts[0].Username, Password: uuid.New().String(), Nickname: newNick1},
			{Username: accounts[1].Username, Password: uuid.New().String(), Nickname: newNick2},
			{Username: uuid.New().String(), Password: uuid.New().String(), Nickname: newNick4},
		}
		require.NoError(t, repo.Clause(func(cls *AccountColumns) clause.Expression {
			return clause.OnConflict{
				Columns:   []clause.Column{{Name: cls.Username.Name()}},
				DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
			}
		}).Creates(upsertAccounts))

		// Verify updates
		// 验证更新结果
		res1, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(accounts[0].Username))
		})
		require.NoError(t, err)
		require.Equal(t, newNick1, res1.Nickname)
		require.Equal(t, accounts[0].Password, res1.Password)

		res2, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(accounts[1].Username))
		})
		require.NoError(t, err)
		require.Equal(t, newNick2, res2.Nickname)

		// Verify new insert
		// 验证新插入的记录
		res4, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(upsertAccounts[2].Username))
		})
		require.NoError(t, err)
		require.Equal(t, newNick4, res4.Nickname)

		// Verify original account not affected
		// 验证原始账户未受影响
		res3, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(accounts[2].Username))
		})
		require.NoError(t, err)
		require.Equal(t, preNick3, res3.Nickname)
	})
}
