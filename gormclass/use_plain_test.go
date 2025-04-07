package gormclass_test

import (
	"testing"

	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestUsePlainWithAccount(t *testing.T) {
	one, cls := gormclass.UsePlain(&Account{})
	t.Log(one.TableName())
	t.Log(neatjsons.S(cls))
}

func TestUmcPlainWithExample(t *testing.T) {
	one, cls := gormclass.UmcPlain(&Example{})
	t.Log(one.TableName())
	t.Log(neatjsons.S(cls))
}
