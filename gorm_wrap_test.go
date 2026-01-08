package gormrepo_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/tests"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestGormWrap_First(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormWrap(gormrepo.Umc(db, &Account{}))

	t.Run("case-1", func(t *testing.T) {
		var account Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		}, &account).Error)
		require.Equal(t, "demo-1-nickname", account.Nickname)
	})

	t.Run("case-2", func(t *testing.T) {
		var account Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		}, &account).Error)
		require.Equal(t, "demo-2-nickname", account.Nickname)
	})
}

func TestGormWrap_Where(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormWrap(gormrepo.Umc(db, &Account{}))

	dbx := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})
	var account Account
	require.NoError(t, dbx.First(&account).Error)
	require.Equal(t, "demo-1-nickname", account.Nickname)
}

func TestGormWrap_Find(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	setupDemoData(t, db)

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

	var accounts []*Account
	require.NoError(t, repo.Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Like("demo-%-username"))
	}, &accounts).Error)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestGormWrap_Update(t *testing.T) {
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	username := uuid.New().String()
	require.NoError(t, db.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	username := uuid.New().String()
	require.NoError(t, db.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	username := uuid.New().String()
	require.NoError(t, db.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

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
	db := tests.NewMemDB(t)
	defer rese.F0(rese.P1(db.DB()).Close)
	must.Done(db.AutoMigrate(&Account{}))

	username := uuid.New().String()
	require.NoError(t, db.Save(newAccount(username)).Error)

	repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

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
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

func TestGormWrap_Save(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormWrap_Creates tests batch insert of multiple records
// TestGormWrap_Creates 测试批量插入多条记录
func TestGormWrap_Creates(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

		accounts := []*Account{
			newAccount(uuid.New().String()),
			newAccount(uuid.New().String()),
			newAccount(uuid.New().String()),
		}
		require.NoError(t, repo.Creates(accounts).Error)

		// Verify all records have IDs assigned
		// 验证所有记录都分配了 ID
		for _, account := range accounts {
			require.NotZero(t, account.ID)
		}

		// Verify all records exist in database
		// 验证所有记录都存在于数据库中
		var count int64
		require.NoError(t, repo.Mold().Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.In([]string{
				accounts[0].Username,
				accounts[1].Username,
				accounts[2].Username,
			}))
		}).Count(&count).Error)
		require.Equal(t, int64(3), count)
	})
}

// TestGormWrap_CreateInBatches tests batch insert with batch size control
// TestGormWrap_CreateInBatches 测试带批次大小控制的分批插入
func TestGormWrap_CreateInBatches(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

		// Create 10 accounts with batch size 3
		// 创建 10 个账户，每批 3 条
		accounts := make([]*Account, 10)
		for i := 0; i < 10; i++ {
			accounts[i] = newAccount(uuid.New().String())
		}
		require.NoError(t, repo.CreateInBatches(accounts, 3).Error)

		// Verify all records have IDs assigned
		// 验证所有记录都分配了 ID
		for _, account := range accounts {
			require.NotZero(t, account.ID)
		}

		// Verify all records exist in database
		// 验证所有记录都存在于数据库中
		usernames := make([]string, len(accounts))
		for i, account := range accounts {
			usernames[i] = account.Username
		}
		var count int64
		require.NoError(t, repo.Mold().Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.In(usernames))
		}).Count(&count).Error)
		require.Equal(t, int64(10), count)
	})
}

// TestGormWrap_Saves tests batch insert or update of multiple records
// TestGormWrap_Saves 测试批量插入或更新多条记录
func TestGormWrap_Saves(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

		// First save - should create new records
		// 首次保存 - 应该创建新记录
		accounts := []*Account{
			newAccount(uuid.New().String()),
			newAccount(uuid.New().String()),
		}
		require.NoError(t, repo.Saves(accounts).Error)

		for _, account := range accounts {
			require.NotZero(t, account.ID)
		}

		// Update nicknames and save again - should update existing records
		// 更新 nickname 并再次保存 - 应该更新现有记录
		newNickname0 := uuid.New().String()
		newNickname1 := uuid.New().String()
		accounts[0].Nickname = newNickname0
		accounts[1].Nickname = newNickname1
		require.NoError(t, repo.Saves(accounts).Error)

		// Verify updates
		// 验证更新
		var res0 Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.ID.Eq(accounts[0].ID))
		}, &res0).Error)
		require.Equal(t, newNickname0, res0.Nickname)

		var res1 Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.ID.Eq(accounts[1].ID))
		}, &res1).Error)
		require.Equal(t, newNickname1, res1.Nickname)
	})
}

func TestGormWrap_Delete(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

func TestGormWrap_DeleteW(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

func TestGormWrap_DeleteM(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormWrap_UpdatesO tests update using object primary key as condition
// O = Object, updates record located by object's primary key
//
// TestGormWrap_UpdatesO 测试使用 object 主键作为条件的更新
// O = Object，通过 object 的主键定位要更新的记录
func TestGormWrap_UpdatesO(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormWrap_UpdatesC tests update using combined conditions: object primary key plus where clause
// C = Combined, uses both object primary key and where conditions
//
// TestGormWrap_UpdatesC 测试使用组合条件的更新：object 主键加上 where 子句
// C = Combined，同时使用 object 主键和 where 条件进行精确定位
func TestGormWrap_UpdatesC(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
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
	})
}

// TestGormWrap_Clauses tests Clauses method for upsert operations
// TestGormWrap_Clauses 测试 Clauses 方法的 upsert 操作
func TestGormWrap_Clauses(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

		// First create
		// 首次创建
		account1 := &Account{
			Username: uuid.New().String(),
			Password: uuid.New().String(),
			Nickname: uuid.New().String(),
		}
		require.NoError(t, repo.Create(account1).Error)
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
		}).Create(account2).Error)

		// Verify the nickname was updated
		// 验证 nickname 已被更新
		var res Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(account1.Username))
		}, &res).Error)
		require.Equal(t, account2.Username, res.Username)
		require.Equal(t, account2.Nickname, res.Nickname)
		require.Equal(t, account1.Password, res.Password)
	})
}

// TestGormWrap_Clause tests Clause method for type-safe upsert operations
// TestGormWrap_Clause 测试 Clause 方法的类型安全 upsert 操作
func TestGormWrap_Clause(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

		// First create
		// 首次创建
		account1 := &Account{
			Username: uuid.New().String(),
			Password: uuid.New().String(),
			Nickname: uuid.New().String(),
		}
		require.NoError(t, repo.Create(account1).Error)
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
		}).Create(account2).Error)

		// Verify the nickname was updated
		// 验证 nickname 已被更新
		var res Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(account1.Username))
		}, &res).Error)
		require.Equal(t, account2.Username, res.Username)
		require.Equal(t, account2.Nickname, res.Nickname)
		require.Equal(t, account1.Password, res.Password)
	})
}

// TestGormWrap_Clause_Creates tests Clause + Creates for batch upsert operations
// TestGormWrap_Clause_Creates 测试 Clause + Creates 的批量 upsert 操作
func TestGormWrap_Clause_Creates(t *testing.T) {
	tests.NewDBRun(t, func(db *gorm.DB) {
		must.Done(db.AutoMigrate(&Account{}))

		repo := gormrepo.NewGormWrap(gormrepo.Use(db, &Account{}))

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
		require.NoError(t, repo.Creates(accounts).Error)

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
		}).Creates(upsertAccounts).Error)

		// Verify updates
		// 验证更新结果
		var res1 Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(accounts[0].Username))
		}, &res1).Error)
		require.Equal(t, newNick1, res1.Nickname)
		require.Equal(t, accounts[0].Password, res1.Password)

		var res2 Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(accounts[1].Username))
		}, &res2).Error)
		require.Equal(t, newNick2, res2.Nickname)

		// Verify new insert
		// 验证新插入的记录
		var res4 Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(upsertAccounts[2].Username))
		}, &res4).Error)
		require.Equal(t, newNick4, res4.Nickname)

		// Verify original account not affected
		// 验证原始账户未受影响
		var res3 Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq(accounts[2].Username))
		}, &res3).Error)
		require.Equal(t, preNick3, res3.Nickname)
	})
}
