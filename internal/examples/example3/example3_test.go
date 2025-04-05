package example3

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormjoin"
	"github.com/yyle88/gormrepo/gormtablerepo"
	"github.com/yyle88/gormrepo/internal/examples/example3/models"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&models.User{}, &models.Order{}))

	const userCount = 10
	users := make([]*models.User, 0, userCount)
	for idx := 0; idx < userCount; idx++ {
		users = append(users, &models.User{
			ID:   0,
			Name: "name" + strconv.Itoa(idx+1),
		})
	}
	done.Done(db.Create(&users).Error)

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
	done.Done(db.Create(&orders).Error)

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
	var results []*UserOrder
	require.NoError(t, db.Table("users").
		Select("users.id as user_id, users.name as user_name, orders.id as order_id, orders.amount as order_amount").
		Joins("left join orders on orders.user_id = users.id").
		Order("users.id asc, orders.id asc").
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

	userOrder := &UserOrder{}
	userColumns := userRepo.TableColumns()
	orderColumns := orderRepo.TableColumns()

	//这是使用名称的逻辑
	var results []*UserOrder
	require.NoError(t, db.Table(userRepo.GetTableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.AsName(gormcnm.Cnm(userOrder.UserID, "user_id")),
			userColumns.Name.AsName(gormcnm.Cnm(userOrder.UserName, "user_name")),
			orderColumns.ID.AsName(gormcnm.Cnm(userOrder.OrderID, "order_id")),
			orderColumns.Amount.AsName(gormcnm.Cnm(userOrder.OrderAmount, "order_amount")),
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
