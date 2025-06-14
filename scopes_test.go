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

// TestRepo_NewScope tests NewScope by filtering Accounts by username.
func TestRepo_NewScope(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{})) // Init repo with Account

	// Create scope to filter by username="demo-1-username"
	scope := repo.NewScope(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})

	// Query first matching Account
	var account Account
	err := caseDB.Scopes(scope).First(&account).Error
	require.NoError(t, err) // Check no error

	t.Log(neatjsons.S(account)) // Log result

	require.Equal(t, "demo-1-nickname", account.Nickname) // Check nickname
}

// TestRepo_Paginate tests new-paginate-scope with ordering and pagination.
func TestRepo_Paginate(t *testing.T) {
	repo := gormrepo.NewBaseRepo(gormclass.Use(&Account{})) // Init repo with Account

	// Create scope to filter by username in ("demo-1-username", "demo-2-username")
	condScope := repo.NewScope(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	})

	// Create scope to order by username (desc) nickname (asc) with limit=10, offset=0
	pageScope := repo.NewPaginateScope(func(cls *AccountColumns) gormcnm.OrderByBottle {
		return cls.Username.Ob("desc").Ob(cls.Nickname.Ob("asc"))
	}, &gormrepo.Pagination{
		Limit:  10,
		Offset: 0,
	})

	// Query Accounts with pagination
	var accounts []*Account
	err := caseDB.Scopes(condScope, pageScope).Find(&accounts).Error
	require.NoError(t, err) // Check no error

	t.Log(neatjsons.S(accounts)) // Log results

	require.Len(t, accounts, 2)
	require.Equal(t, "demo-2-username", accounts[0].Username) // Check username
	require.Equal(t, "demo-1-username", accounts[1].Username) // Check username
}
