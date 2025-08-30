package example14_test

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormjoin"
	"github.com/yyle88/gormrepo/gormtablerepo"
	"github.com/yyle88/gormrepo/internal/examples/example14"
	"github.com/yyle88/gormrepo/internal/examples/example14/internal/models"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
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

	done.Done(db.AutoMigrate(&models.Guest{}, &models.Order{}))

	for idx := 0; idx < 3; idx++ {
		// 插入数据
		username := "username" + strconv.Itoa(idx+1)
		guest := example14.Guest{
			Guest: models.Guest{
				Model:    gorm.Model{},
				Username: username,
				Nickname: "nickname" + strconv.Itoa(idx+1),
				Phone:    "phone" + strconv.Itoa(idx+1),
				Email:    "email" + strconv.Itoa(idx+1) + "@example.com",
			},
			Orders: []*models.Order{},
		}
		for odx := 0; odx < 3; odx++ {
			guest.Orders = append(guest.Orders, &models.Order{
				Model:       gorm.Model{},
				GuestID:     0,
				ProductName: "product" + strconv.Itoa(odx+1) + "(" + username + ")",
				Amount:      1 + rand.IntN(1000),
				Cost:        float64(rand.IntN(1000)) + rand.Float64(),
				Address:     "address" + strconv.Itoa(odx+1) + "(" + username + ")",
			})
		}
		must.Done(db.Create(&guest).Error)
	}

	caseDB = db
	m.Run()
}

type GuestOrderView struct {
	GuestUsername string
	GuestNickname string
	ProductName   string
	Amount        int
	Cost          float64
}

func TestExample(t *testing.T) {
	view := &GuestOrderView{}

	db := caseDB
	guestRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Guest{}))
	orderRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Order{}))

	//这是使用名称的逻辑
	var results []*GuestOrderView
	require.NoError(t, db.Table(guestRepo.GetTableName()).
		Select(gormcnmstub.MergeSlices(
			guestRepo.BuildColumns(func(cls *models.GuestColumns) []string {
				return []string{
					cls.Username.AsName(gormcnm.Cnm(view.GuestUsername, "guest_username")),
					cls.Nickname.AsName(gormcnm.Cnm(view.GuestNickname, "guest_nickname")),
				}
			}),
			orderRepo.BuildColumns(func(cls *models.OrderColumns) []string {
				return []string{
					cls.ProductName.AsName(gormcnm.Cnm(view.ProductName, "product_name")),
					cls.Amount.AsName(gormcnm.Cnm(view.Amount, "amount")),
					cls.Cost.AsName(gormcnm.Cnm(view.Cost, "cost")),
				}
			}),
		)).
		Joins(gormjoin.LEFTJOIN(guestRepo, orderRepo).On(func(gc *models.GuestColumns, oc *models.OrderColumns) []string {
			return []string{
				oc.GuestID.OnEq(gc.ID),
			}
		})).
		Scopes(guestRepo.Base().NewWhereScope(func(db *gorm.DB, cls *models.GuestColumns) *gorm.DB {
			return db.Where(cls.ID.Gte(2))
		})).
		Scopes(guestRepo.Base().NewOrderScope(func(cls *models.GuestColumns) gormcnm.OrderByBottle {
			return cls.ID.Ob("asc")
		})).
		Scopes(orderRepo.Base().NewOrderScope(func(cls *models.OrderColumns) gormcnm.OrderByBottle {
			return cls.ID.Ob("asc")
		})).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	require.Equal(t, 6, len(results))
	require.Equal(t, "username2", results[0].GuestUsername)
	require.Equal(t, "nickname2", results[0].GuestNickname)
	require.Equal(t, "product1(username2)", results[0].ProductName)
	require.Positive(t, results[0].Amount)
	require.Positive(t, results[0].Cost)

	require.Equal(t, "username3", results[5].GuestUsername)
	require.Equal(t, "nickname3", results[5].GuestNickname)
	require.Equal(t, "product3(username3)", results[5].ProductName)
	require.Positive(t, results[5].Amount)
	require.Positive(t, results[5].Cost)
}

func TestExample2(t *testing.T) { // 查询并预加载
	db := caseDB

	var guests []*example14.Guest
	require.NoError(t, db.Preload("Orders").Where("guests.id >= ?", 2).Find(&guests).Error)

	var results []*GuestOrderView
	for _, guest := range guests {
		for _, order := range guest.Orders {
			results = append(results, &GuestOrderView{
				GuestUsername: guest.Username,
				GuestNickname: guest.Nickname,
				ProductName:   order.ProductName,
				Amount:        order.Amount,
				Cost:          order.Cost,
			})
		}
	}
	t.Log(neatjsons.S(results))
	require.Equal(t, 6, len(results))
	require.Equal(t, "username2", results[0].GuestUsername)
	require.Equal(t, "nickname2", results[0].GuestNickname)
	require.Equal(t, "product1(username2)", results[0].ProductName)
	require.Positive(t, results[0].Amount)
	require.Positive(t, results[0].Cost)

	require.Equal(t, "username3", results[5].GuestUsername)
	require.Equal(t, "nickname3", results[5].GuestNickname)
	require.Equal(t, "product3(username3)", results[5].ProductName)
	require.Positive(t, results[5].Amount)
	require.Positive(t, results[5].Cost)
}
