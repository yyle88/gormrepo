package example11_test

import (
	"database/sql" // Added for sqlDB.Close()
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormjoin"
	"github.com/yyle88/gormrepo/gormtablerepo"
	"github.com/yyle88/gormrepo/internal/examples/example11/internal/models"
	"github.com/yyle88/neatjson/neatjsons"
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

	err = db.AutoMigrate(&models.User{}, &models.Order{}) // Changed done.Done to require.NoError
	if err != nil {
		panic(err)
	}

	const userCount = 10
	users := make([]*models.User, 0, userCount)
	for idx := 0; idx < userCount; idx++ {
		users = append(users, &models.User{
			ID:   0,
			Name: "name" + strconv.Itoa(idx+1),
		})
	}
	err = db.Create(&users).Error // Changed done.Done to require.NoError
	if err != nil {
		panic(err)
	}

	const orderCount = 20
	orders := make([]*models.Order, 0, orderCount)
	for idx := 0; idx < orderCount; idx++ {
		userID := users[rand.IntN(len(users))].ID

		orders = append(orders, &models.Order{
			ID:     0,
			UserID: userID,
			Amount: float64(rand.IntN(1000)) + rand.Float64(),
		})
	}
	err = db.Create(&orders).Error // Changed done.Done to require.NoError
	if err != nil {
		panic(err)
	}

	caseDB = db
	m.Run()
}

func TestTableJoin(t *testing.T) {
	expected0Text := neatjsons.S(selectFunc0(t, caseDB))
	expected1Text := neatjsons.S(selectFunc1(t, caseDB))
	//确保两者结果相同
	require.Equal(t, expected0Text, expected1Text)
}

// 这是比较常规的逻辑
func selectFunc0(t *testing.T, db *gorm.DB) []*UserOrder {
	userRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.User{}))
	orderRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Order{}))

	userColumns := userRepo.TableColumns()
	orderColumns := orderRepo.TableColumns()
	var common = &gormcnm.ColumnOperationClass{} // Added common

	var results []*UserOrder
	require.NoError(t, db.Table(userRepo.GetTableName()).
		Select(common.MergeStmts( // Refactored Select
			userColumns.ID.AsName(gormcnm.ColumnName[uint]("user_id")),
			userColumns.Name.AsName(gormcnm.ColumnName[string]("user_name")),
			orderColumns.ID.AsName(gormcnm.ColumnName[uint]("order_id")),
			orderColumns.Amount.AsName(gormcnm.ColumnName[float64]("order_amount")),
		)).
		Joins(gormjoin.LEFTJOIN(userRepo, orderRepo).On(func(uc *models.UserColumns, oc *models.OrderColumns) []string {
			return []string{oc.UserID.OnEq(uc.ID)}
		})).
		Order(userColumns.ID.Ob("asc").Ob(orderColumns.ID.Ob("asc")).Ox()). // Refactored Order
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

type UserOrder struct {
	UserID      uint
	UserName    string
	OrderID     uint
	OrderAmount float64
}

func selectFunc1(t *testing.T, db *gorm.DB) []*UserOrder {
	userRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.User{}))
	orderRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Order{}))

	userColumns := userRepo.TableColumns()
	orderColumns := orderRepo.TableColumns()
	var common = &gormcnm.ColumnOperationClass{} // Added common

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(userRepo.GetTableName()).
		Select(common.MergeStmts( // Replaced gormcnmstub.MergeStmts with common.MergeStmts
			userColumns.ID.AsName(gormcnm.ColumnName[uint]("user_id")),
			userColumns.Name.AsName(gormcnm.ColumnName[string]("user_name")),
			orderColumns.ID.AsName(gormcnm.ColumnName[uint]("order_id")),
			orderColumns.Amount.AsName(gormcnm.ColumnName[float64]("order_amount")),
		)).
		Joins(gormjoin.LEFTJOIN(userRepo, orderRepo).On(func(uc *models.UserColumns, oc *models.OrderColumns) []string {
			return []string{
				oc.UserID.OnEq(uc.ID),
			}
		})).
		Order(userColumns.ID.Ob("asc").Ob(orderColumns.ID.Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}
