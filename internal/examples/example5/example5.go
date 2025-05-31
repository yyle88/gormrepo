package example5

import "github.com/yyle88/gormrepo/internal/examples/example5/internal/models"

type Guest struct {
	models.Guest
	Orders []*models.Order
}
