package models_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormrepo/internal/examples/example10/internal/models"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

//go:generate go test -v -run TestGenerate

func TestGenerate(t *testing.T) {
	absPath := runtestpath.SrcPath(t)
	require.True(t, osmustexist.IsFile(absPath))

	objects := []any{
		&models.Author{},
		&models.Book{},
	}
	options := gormcngen.NewOptions().
		WithColumnClassExportable(true).
		WithColumnsMethodRecvName("p").
		WithColumnsCheckFieldType(true).
		WithIsGenFuncTableColumns(true)
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen()
}
