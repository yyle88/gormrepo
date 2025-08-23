package models_test

import (
	"testing"

	"github.com/yyle88/gormrepo/internal/examples/example02/internal/models"
)

func TestUmcV1(t *testing.T) {
	one, cls := models.UmcV1(&models.Account{})
	t.Log(one.TableName())
	t.Log(cls.Username)
	t.Log(cls.Password)
}

func TestUmcV2(t *testing.T) {
	one, cls := models.UmcV2(&models.Account{})
	t.Log(one.TableName())
	t.Log(cls.Username)
	t.Log(cls.Password)
}

func TestUmcV3(t *testing.T) {
	one, cls := models.UmcV3(&models.Account{})
	t.Log(one.TableName())
	t.Log(cls.Username)
	t.Log(cls.Password)
}
