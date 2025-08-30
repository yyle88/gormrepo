package models_test

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormrepo/internal/examples/example11/internal/models"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
)

// Auto generate columns with go generate command
// Support execution via: go generate ./...
// Delete this comment block if auto generation is not needed
//
// 使用 go generate 命令自动生成列定义
// 支持通过以下命令执行：go generate ./...
// 如果不需要自动生成，可以删除此注释块
//
//go:generate go test -v -run TestGenerateColumns
func TestGenerateColumns(t *testing.T) {
	// Generate columns code to separate file instead of mixing with models
	// 将列代码生成到独立文件，而不是与模型混合
	// Generate to gormcnm.gen.go file matching test file name pattern
	// 生成到与测试文件名模式匹配的 gormcnm.gen.go 文件
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	// Define model objects for column generation - supports both pointer and non-pointer types
	// 定义需要生成列的模型对象 - 支持指针类型和非指针类型
	objects := []any{
		&models.User{},
		&models.Order{},
	}

	// Configure generation options with latest best practices
	// 使用最新最佳实践配置生成选项
	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable column class names like ExampleColumns // 生成可导出的列类名称如 ExampleColumns
		WithColumnsMethodRecvName("c").  // Set receiver name for column methods // 设置列方法的接收器名称
		WithColumnsCheckFieldType(true). // Enable field type checking for type safe // 启用字段类型检查以获得更好的类型安全
		WithIsGenFuncTableColumns(true)  // Generate table column functions for join operations // 生成表列函数用于连接操作

	// Create configuration and generate code to target file
	// 创建配置并生成代码到目标文件
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen() // Generate code to "gormcnm.gen.go" file // 生成代码到 "gormcnm.gen.go" 文件
}
