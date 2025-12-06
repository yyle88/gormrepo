package gormrepo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/neatjson/neatjsons"
	"gorm.io/gorm"
)

// TestRepo_NewScope tests NewScope via matching Accounts via username
// TestRepo_NewScope 测试通过用户名匹配 Account 的 NewScope 方法
func TestRepo_NewScope(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{})) // Init repo with Account // 使用 Account 初始化仓储

	// Create scope to match via username="demo-1-username"
	// 创建 scope 通过 username="demo-1-username" 匹配
	scope := repo.NewScope(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})

	// Query first matching Account
	// 查询第一个匹配的 Account
	var account Account
	err := caseDB.Scopes(scope).First(&account).Error
	require.NoError(t, err) // Check no error // 检查无错误

	t.Log(neatjsons.S(account)) // Log result // 打印结果

	require.Equal(t, "demo-1-nickname", account.Nickname) // Check nickname // 检查昵称
}

// TestRepo_Paginate tests new-paginate-scope with sorting and pagination
// TestRepo_Paginate 测试带排序和分页的 new-paginate-scope
func TestRepo_Paginate(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{})) // Init repo with Account // 使用 Account 初始化仓储

	// Create scope to match via username in ("demo-1-username", "demo-2-username")
	// 创建 scope 通过 username 在 ("demo-1-username", "demo-2-username") 中匹配
	condScope := repo.NewScope(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	})

	// Create scope to sort via username (desc) nickname (asc) with limit=10, offset=0
	// 创建 scope 按 username (desc) nickname (asc) 排序，limit=10, offset=0
	pageScope := repo.NewPaginateScope(func(cls *AccountColumns) gormcnm.OrderByBottle {
		return cls.Username.Ob("desc").Ob(cls.Nickname.Ob("asc"))
	}, &gormrepo.Pagination{
		Limit:  10,
		Offset: 0,
	})

	// Query Accounts with pagination
	// 带分页查询 Accounts
	var accounts []*Account
	err := caseDB.Scopes(condScope, pageScope).Find(&accounts).Error
	require.NoError(t, err) // Check no error // 检查无错误

	t.Log(neatjsons.S(accounts)) // Log results // 打印结果

	require.Len(t, accounts, 2)
	require.Equal(t, "demo-2-username", accounts[0].Username) // Check username // 检查用户名
	require.Equal(t, "demo-1-username", accounts[1].Username) // Check username // 检查用户名
}
