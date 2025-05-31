package example4

import (
	"github.com/yyle88/gormrepo/internal/examples/example4/internal/models"
)

type User struct {
	models.User
	Orders []*Order
}

type Order struct {
	models.Order
	Products []*models.Product
}
