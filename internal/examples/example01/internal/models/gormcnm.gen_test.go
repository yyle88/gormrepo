package models_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormrepo/internal/examples/example01/internal/models"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

// 这句能让你的代码配合 go generate ./... 执行，假如不需要可以删除这句注释
//
//go:generate go test -v -run TestGenerate
func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t) //根据当前测试文件的路径找到其对应的源文件路径
	t.Log(absPath)

	//出于安全起见，需要需要判断目标文件是已经存在的，需要手动创建该文件，让代码能找到此文件
	require.True(t, osmustexist.IsFile(absPath))

	//在这里写下你要生成的 models 的对象列表，指针类型或非指针类型都是可以的，选中生成模型
	objects := []any{&models.Account{}, &models.Example{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true) //中间类型名称的样式为可导出的 ExampleColumns

	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen() //将会把生成后的代码写到目标位置，即 "gormcnm.gen.go" 这个文件里
}
