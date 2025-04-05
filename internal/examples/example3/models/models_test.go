package models

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

func TestGenerateColumns(t *testing.T) {
	absPath := runtestpath.SrcPath(t) // Retrieve the absolute path of the source file based on the current test file's location
	// 获取当前测试文件位置基础上的源文件绝对路径
	t.Log(absPath)

	// Verify the existence of the target file. The file should be created manually to ensure it can be located by the code.
	// 检查目标文件是否存在。文件应手动创建，确保代码能够找到它。
	require.True(t, osmustexist.IsFile(absPath))

	// List the models for which columns will be generated. Both pointer and non-pointer types are supported.
	// 设置需要生成列的模型，这里支持指针类型和非指针类型。
	objects := []any{&User{}, &Order{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable struct names (e.g., ExampleColumns) // 生成可导出的结构体名称（例如 ExampleColumns）
		WithColumnsMethodRecvName("a").
		WithColumnsCheckFieldType(true).
		WithIsGenFuncTableColumns(true)

	// Configure code generation settings
	// 配置代码生成设置
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen() // Generate and write the code to the target location (e.g., "gormcnm.gen.go") // 生成并将代码写入目标位置（例如 "gormcnm.gen.go"）
}
