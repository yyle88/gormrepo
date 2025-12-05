package models_test

import (
	"testing"

	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormrepo/internal/examples/example02/internal/models"
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
	// Retrieve the absolute path of the source file based on current test file location
	// 根据当前测试文件位置获取源文件的绝对路径
	absPath := osmustexist.FILE(runtestpath.SrcPath(t))
	t.Log(absPath)

	// Define data objects for column generation - supports both instance and non-instance types
	// 定义需要生成列的模型对象 - 支持指针类型和非指针类型
	objects := []any{
		&models.Account{},
		&models.Example{},
	}

	// Configure generation options with latest best practices
	// 使用最新最佳实践配置生成选项
	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable column class names like ExampleColumns // 生成可导出的列类名称如 ExampleColumns
		WithColumnsMethodRecvName("c").  // Set receiver name for column methods // 设置列方法的接收器名称
		WithColumnsCheckFieldType(true)  // Enable field type checking for type safe // 启用字段类型检查以获得更好的类型安全

	// Create configuration and generate code to target file
	// 创建配置并生成代码到目标文件
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen() // Generate code to "gormcnm.gen.go" file // 生成代码到 "gormcnm.gen.go" 文件
}
