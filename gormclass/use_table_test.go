package gormclass_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/neatjson/neatjsons"
)

func TestTableUseWithAccount(t *testing.T) {
	one, tableName, cls := gormclass.UseTable(&Account{})
	t.Log(one.TableName())
	t.Log(tableName)
	require.Equal(t, one.TableName(), tableName)
	t.Log(neatjsons.S(cls))
}

func TestTableUmcWithExample(t *testing.T) {
	one, tableName, cls := gormclass.UmcTable(&Example{})
	t.Log(one.TableName())
	t.Log(tableName)
	require.Equal(t, one.TableName(), tableName)
	t.Log(neatjsons.S(cls))
}
