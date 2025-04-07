package example2models

import (
	"testing"
)

func TestUmcV1(t *testing.T) {
	one, cls := UmcV1(&Account{})
	t.Log(one.TableName())
	t.Log(cls.Username)
	t.Log(cls.Nickname)
	t.Log(cls.Password)
}

func TestUmcV2(t *testing.T) {
	one, cls := UmcV2(&Account{})
	t.Log(one.TableName())
	t.Log(cls.Username)
	t.Log(cls.Nickname)
	t.Log(cls.Password)
}

func TestUmcV3(t *testing.T) {
	one, cls := UmcV3(&Account{})
	t.Log(one.TableName())
	t.Log(cls.Username)
	t.Log(cls.Nickname)
	t.Log(cls.Password)
}
