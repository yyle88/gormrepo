package models_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	models_pkg "github.com/yyle88/gormrepo/internal/examples/example05/internal/models"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

//go:generate go test -v -run TestGenerate
func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	t.Log(absPath)

	require.True(t, osmustexist.IsFile(absPath))

	objects := []any{&models_pkg.Article{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true)

	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen()
}
