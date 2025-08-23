package models_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormrepo/internal/examples/example04/internal/models"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

//go:generate go test -v -run TestGenerate
func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	t.Log(absPath)

	require.True(t, osmustexist.IsFile(absPath))

	// 生成两个模型的列定义
	objects := []any{&models.Product{}, &models.Category{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true)

	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen()
}
