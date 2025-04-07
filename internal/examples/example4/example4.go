package example4

import models "github.com/yyle88/gormrepo/internal/examples/example4/example4models"

type User struct {
	models.User
	Orders []*Order
}

type Order struct {
	models.Order
	Products []*models.Product
}
