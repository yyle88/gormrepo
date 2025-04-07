package example4_test

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
	"github.com/yyle88/gormrepo/internal/examples/example4"
	models "github.com/yyle88/gormrepo/internal/examples/example4/example4models"
	"github.com/yyle88/must"
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

	done.Done(db.AutoMigrate(&models.User{}, &models.Order{}, &models.Product{}))

	for idx := 0; idx < 5; idx++ {
		// 插入数据
		user := example4.User{
			User: models.User{
				ID:   0,
				Name: "name" + strconv.Itoa(idx+1),
			},
			Orders: []*example4.Order{
				&example4.Order{
					Order: models.Order{
						ID:     0,
						UserID: 0,
						Amount: float64(rand.IntN(1000)) + rand.Float64(),
					},
					Products: []*models.Product{
						&models.Product{
							ID:      0,
							OrderID: 0,
							Name:    "Laptop",
						},
						&models.Product{
							ID:      0,
							OrderID: 0,
							Name:    "Mouse",
						},
					},
				},
				&example4.Order{
					Order: models.Order{
						ID:     0,
						UserID: 0,
						Amount: float64(rand.IntN(1000)) + rand.Float64(),
					},
					Products: []*models.Product{
						&models.Product{
							ID:      0,
							OrderID: 0,
							Name:    "Phone",
						},
						&models.Product{
							ID:      0,
							OrderID: 0,
							Name:    "Macbook",
						},
					},
				},
			},
		}
		must.Done(db.Create(&user).Error)
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
func selectFunc0(t *testing.T, db *gorm.DB) []*UserOrderProduct {
	var results []*UserOrderProduct
	require.NoError(t, db.Table("users").
		Select("users.id as user_id, users.name as user_name, "+
			"orders.id as order_id, orders.amount as order_amount, "+
			"products.id as product_id, products.name as product_name").
		Joins("left join orders on orders.user_id = users.id").
		Joins("left join products on products.order_id = orders.id").
		Order("users.id asc, orders.id asc, products.id asc").
		Where("products.Name in ?", []string{"Laptop", "Mouse", "Phone"}).
		Where("users.id >= ?", 2).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

type UserOrderProduct struct {
	UserID      uint
	UserName    string
	OrderID     uint
	OrderAmount float64
	ProductID   uint
	ProductName string
}

func selectFunc1(t *testing.T, db *gorm.DB) []*UserOrderProduct {
	userRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.User{}))
	orderRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Order{}))
	productRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Product{}))

	userOrder := &UserOrderProduct{}
	userColumns := userRepo.TableColumns()
	orderColumns := orderRepo.TableColumns()
	productColumns := productRepo.TableColumns()

	//这是使用名称的逻辑
	var results []*UserOrderProduct
	require.NoError(t, db.Table(userRepo.GetTableName()).
		Select(gormcnmstub.MergeStmts(
			userColumns.ID.AsName(gormcnm.Cnm(userOrder.UserID, "user_id")),
			userColumns.Name.AsName(gormcnm.Cnm(userOrder.UserName, "user_name")),
			orderColumns.ID.AsName(gormcnm.Cnm(userOrder.OrderID, "order_id")),
			orderColumns.Amount.AsName(gormcnm.Cnm(userOrder.OrderAmount, "order_amount")),
			productColumns.ID.AsName(gormcnm.Cnm(userOrder.ProductID, "product_id")),
			productColumns.Name.AsName(gormcnm.Cnm(userOrder.ProductName, "product_name")),
		)).
		Joins(gormjoin.LEFTJOIN(userRepo, orderRepo).On(func(uc *models.UserColumns, oc *models.OrderColumns) []string {
			return []string{
				oc.UserID.OnEq(uc.ID),
			}
		})).
		Joins(gormjoin.LEFTJOIN(orderRepo, productRepo).On(func(oc *models.OrderColumns, pc *models.ProductColumns) []string {
			return []string{
				pc.OrderID.OnEq(oc.ID),
			}
		})).
		Where(productColumns.Name.In([]string{"Laptop", "Mouse", "Phone"})).
		Where(userColumns.ID.Gte(2)).
		Order(userColumns.ID.Ob("asc").
			Ob(orderColumns.ID.Ob("asc")).
			Ob(productColumns.ID.Ob("asc")).Ox()).
		Scan(&results).Error)
	t.Log(neatjsons.S(results))
	return results
}

func TestPreload3t(t *testing.T) {
	expected0Text := neatjsons.S(selectFunc0(t, caseDB))
	//确保两者结果相同
	require.Equal(t, expected0Text, neatjsons.S(selectFunc2(t, caseDB)))
	//确保两者结果相同
	require.Equal(t, expected0Text, neatjsons.S(selectFunc3(t, caseDB)))
}

func selectFunc2(t *testing.T, db *gorm.DB) []*UserOrderProduct {
	// 查询并预加载
	var users []*example4.User
	db.Preload("Orders").Preload("Orders.Products", func(db *gorm.DB) *gorm.DB {
		return db.Where("name IN ?", []string{"Laptop", "Mouse", "Phone"})
	}).Where("users.id >= ?", 2).Find(&users)

	var results []*UserOrderProduct
	for _, user := range users {
		for _, order := range user.Orders {
			for _, product := range order.Products {
				results = append(results, &UserOrderProduct{
					UserID:      user.ID,
					UserName:    user.Name,
					OrderID:     order.ID,
					OrderAmount: order.Amount,
					ProductID:   product.ID,
					ProductName: product.Name,
				})
			}
		}
	}
	t.Log(neatjsons.S(results))
	return results
}

func selectFunc3(t *testing.T, db *gorm.DB) []*UserOrderProduct {
	type OrderWithProducts struct {
		models.Order
		SubProducts []*models.Product `gorm:"foreignKey:OrderID;references:ID"` //这个注释非常关键
	}

	type UserWithOrdersWithProducts struct {
		models.User
		SubOrders []*OrderWithProducts `gorm:"foreignKey:UserID;references:ID"` //这个注释非常关键
	}

	// 查询并预加载
	var users []*UserWithOrdersWithProducts
	// 这里要使用成员属性名称，而不是别的比如 Orders 和 Orders.Products
	db.Preload("SubOrders").Preload("SubOrders.SubProducts", func(db *gorm.DB) *gorm.DB {
		return db.Where("name IN ?", []string{"Laptop", "Mouse", "Phone"})
	}).Where("users.id >= ?", 2).Find(&users)

	var results []*UserOrderProduct
	for _, user := range users {
		for _, order := range user.SubOrders {
			for _, product := range order.SubProducts {
				results = append(results, &UserOrderProduct{
					UserID:      user.ID,
					UserName:    user.Name,
					OrderID:     order.ID,
					OrderAmount: order.Amount,
					ProductID:   product.ID,
					ProductName: product.Name,
				})
			}
		}
	}
	t.Log(neatjsons.S(results))
	return results
}
