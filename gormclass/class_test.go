package gormclass_test

import (
	"testing"

	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo/gormclass"
)

func TestImplementGormClass(t *testing.T) {
	var x gormclass.GormClass[*AccountColumns] = &Account{}
	c := x.Columns()
	t.Log(c.Username)
	t.Log(c.Nickname)
	t.Log(c.Password)
}

func TestImplementClassType(t *testing.T) {
	var x gormclass.ClassType[*AccountColumns] = &Account{}
	t.Log(x.TableName())
	c := x.Columns()
	t.Log(c.Username)
	t.Log(c.Nickname)
	t.Log(c.Password)
}

func TestImplementTableClass(t *testing.T) {
	var x gormclass.TableClass[*AccountColumns] = &Account{}
	t.Log(x.TableName())
	c := x.TableColumns(gormcnm.NewTableDecoration(x.TableName()))
	t.Log(c.Username)
	t.Log(c.Nickname)
	t.Log(c.Password)
}
