// Package gormclasscache_test provides comprehensive testing for GORM column caching features
// This package tests three different caching approaches (UmcV1, UmcV2, UmcV3) with various cache backends
// Includes code generation tests and performance validation for type-safe column operations
//
// gormclasscache_test 包为 GORM 列缓存功能提供全面测试
// 此包测试三种不同的缓存实现（UmcV1、UmcV2、UmcV3）和各种缓存后端
// 包括代码生成测试和类型安全列操作的性能验证
package gormclasscache_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo/gormclasscache"
	"github.com/yyle88/mutexmap"
	"github.com/yyle88/mutexmap/cachemap"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

// TestModel represents a sample database model for testing caching features
// Contains basic user account information with three essential fields
// Used as the primary test subject for all caching approaches
//
// TestModel 代表用于测试缓存功能的示例数据库模型
// 包含基本用户账户信息的三个基本字段
// 用作所有缓存实现的主要测试对象
type TestModel struct {
	ID       uint   // Primary key ID // 主键标识符
	Username string // User login username // 用户登录名
	Password string // User password data // 用户密码哈希
}

// TableName gets the database table name for TestModel
// Implements GORM's Tabler interface for explicit table naming
// Returns "test_models" following standard naming conventions
//
// TableName 返回 TestModel 的数据库表名
// 实现 GORM 的 Tabler 接口用于显式表命名
// 返回 "test_models" 遵循标准命名约定
func (*TestModel) TableName() string {
	return "test_models"
}

// Columns gets type-safe column definitions for TestModel
// Auto-generated method that provides type-safe column access
// Each field maps to its corresponding database column with correct types
//
// Columns 返回 TestModel 的类型安全列定义
// 自动生成的方法提供强类型的列访问
// 每个字段都映射到其对应的数据库列并具有正确的类型

func (c *TestModel) Columns() *TestModelColumns {
	return &TestModelColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:       gormcnm.Cnm(c.ID, "id"),
		Username: gormcnm.Cnm(c.Username, "username"),
		Password: gormcnm.Cnm(c.Password, "password"),
	}
}

// TestModelColumns contains type-safe column definitions for TestModel
// Each field provides compile-time type checking and SQL operation methods
// Enables fluent query building with full IDE support and error checking
//
// TestModelColumns 代表 TestModel 的类型安全列定义
// 每个字段提供编译时类型检查和 SQL 操作方法
// 支持流畅的查询构建，完整的 IDE 支持和错误预防

type TestModelColumns struct {
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID       gormcnm.ColumnName[uint]
	Username gormcnm.ColumnName[string]
	Password gormcnm.ColumnName[string]
}

// ===== Code generation logic for above model =====
// ===== 上述模型的代码生成逻辑 =====

// TestGenerateColumns creates type-safe column definitions for TestModel
// Uses gormcngen to auto-generate Columns() method and TestModelColumns struct
// Configured for optimal type safety and IDE integration without embedded operations
// Run with: go generate ./... or go test -v -run TestGenerateColumns
//
// TestGenerateColumns 为 TestModel 生成类型安全的列定义
// 使用 gormcngen 自动生成 Columns() 方法和 TestModelColumns 结构体
// 配置为最佳类型安全和 IDE 集成，不嵌入操作函数
// 运行命令：go generate ./... 或 go test -v -run TestGenerateColumns
//
//go:generate go test -v -run TestGenerateColumns
func TestGenerateColumns(t *testing.T) {
	// Step 1: Get current file path for code generation target
	// 步骤1：获取当前文件路径作为代码生成目标
	absPath := osmustexist.FILE(runpath.CurrentPath())
	t.Log("Generation target file:", absPath) // Log target file path // 记录目标文件路径

	// Step 2: Define model objects for column generation
	// Support both pointer and value model types
	// 步骤2：定义用于生成列的模型对象
	// 支持指针和非指针模型类型
	objects := []any{
		&TestModel{}, // TestModel instance for analysis // TestModel 实例用于分析
	}

	// Step 3: Configure generation options with enterprise-grade settings
	// Optimized for type safety, IDE support, and clean code generation
	// 步骤3：使用企业级设置配置生成选项
	// 针对类型安全、IDE 支持和清洁代码生成进行优化
	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable column class names // 生成可导出的列类名称
		WithColumnsMethodRecvName("c").  // Set receiver name "c" for column methods // 为列方法设置接收器名称 "c"
		WithColumnsCheckFieldType(true). // Enable compile-time field type checking // 启用编译时字段类型检查
		WithEmbedColumnOperations(false) // Disable embedded operations for clean API // 禁用嵌入操作以获得更清洁的 API

	// Step 4: Create generation configuration and execute code generation
	// Creates optimized column definitions into this test file
	// 步骤4：创建生成配置并执行代码生成
	// 直接在此测试文件中生成优化的列定义
	cfg := gormcngen.NewConfigs(objects, options, absPath).
		WithIsGenPreventEdit(false)
	cfg.Gen() // Execute code generation to current file // 执行代码生成到当前文件
}

// ===== Global cache instances for testing different approaches =====
// ===== 用于测试不同实现的全局缓存实例 =====

// cache1 provides cachemap-based caching for UmcV1 testing
// Uses mutexmap/cachemap with capacity of 1 for testing
// Supports concurrent access with built-in locking mechanisms
//
// cache1 为 UmcV1 测试提供基于 cachemap 的缓存
// 使用容量为 1 的 mutexmap/cachemap 用于测试目的
// 支持具有内置锁定机制的并发访问
var cache1 = cachemap.NewMap[string, interface{}](1)

// cache2 provides mutexmap-based caching for UmcV2 testing
// Uses direct mutexmap with capacity of 1 for testing
// Offers alternate concurrent access patterns compared to cachemap
//
// cache2 为 UmcV2 测试提供基于 mutexmap 的缓存
// 使用容量为 1 的直接 mutexmap 用于测试目的
// 与 cachemap 相比提供替代的并发访问模式
var cache2 = mutexmap.NewMap[string, interface{}](1)

// cache3 provides sync.Map-based caching for UmcV3 testing
// Uses Go's built-in concurrent map implementation
// Optimized for scenarios with frequent reads and few writes
//
// cache3 为 UmcV3 测试提供基于 sync.Map 的缓存
// 使用 Go 的内置并发 map 实现
// 针对频繁读取和不频繁写入的场景进行优化
var cache3 = &sync.Map{}

// TestUmcV1 validates cachemap-based caching implementation (UmcV1)
// Tests both initial cache population and subsequent cache access
// Verifies type safety, cache features, and correct column mappings
//
// TestUmcV1 验证基于 cachemap 的缓存实现（UmcV1）
// 测试初始缓存填充和后续缓存检索
// 验证类型安全、缓存功能和正确的列映射
func TestUmcV1(t *testing.T) {
	t.Run("TEST-1", func(t *testing.T) {
		// Test cache population with column data
		// 测试用列数据填充缓存
		one, cls := gormclasscache.UmcV1(&TestModel{}, cache1)

		// Verify model instance is returned
		// 验证模型实例返回
		require.NotNil(t, one, "Model instance should not be nil")                                 // 模型实例不应为空
		require.Equal(t, "test_models", one.TableName(), "Table name should match expected value") // 表名应匹配预期值

		// Verify columns struct is generated and cached
		// 验证列结构体生成和缓存
		require.NotNil(t, cls, "Columns struct should not be nil") // 列结构体不应为空
		t.Log("✓ Table name:", one.TableName())                    // Log table name info // 记录表名信息
		t.Log("✓ Username column:", cls.Username)                  // Log username column mapping // 记录用户名列映射
		t.Log("✓ Password column:", cls.Password)                  // Log password column mapping // 记录密码列映射
	})

	t.Run("TEST-2", func(t *testing.T) {
		// Test cache access without regeneration
		// 测试缓存检索而不重新生成
		one, cls := gormclasscache.UmcV1(&TestModel{}, cache1)

		// Verify results maintain accuracy
		// 验证结果保持准确性
		require.NotNil(t, one, "Model instance should not be nil")                          // 模型实例不应为空
		require.NotNil(t, cls, "Columns struct should not be nil")                          // 列结构体不应为空
		require.Equal(t, "test_models", one.TableName(), "Table name should be consistent") // 表名应保持一致
	})
}

// TestUmcV2 validates mutexmap-based caching implementation (UmcV2)
// Tests alternate caching backend with direct mutexmap usage
// Compares performance and behavior against cachemap method
//
// TestUmcV2 验证基于 mutexmap 的缓存实现（UmcV2）
// 测试使用直接 mutexmap 的替代缓存后端
// 与 cachemap 方法进行性能和行为比较
func TestUmcV2(t *testing.T) {
	t.Run("TEST-1", func(t *testing.T) {
		// Test mutexmap cache population with column definitions
		// 测试用列定义填充 mutexmap 缓存
		one, cls := gormclasscache.UmcV2(&TestModel{}, cache2)

		// Validate mutexmap caching features
		// 验证 mutexmap 缓存功能
		require.NotNil(t, one, "Model instance should not be nil")                                 // 模型实例不应为空
		require.Equal(t, "test_models", one.TableName(), "Table name should match expected value") // 表名应匹配预期值
		require.NotNil(t, cls, "Columns struct should not be nil")                                 // 列结构体不应为空
		t.Log("✓ Table name:", one.TableName())                                                    // Log table name info // 记录表名信息
		t.Log("✓ Username column:", cls.Username)                                                  // Log username column mapping // 记录用户名列映射
		t.Log("✓ Password column:", cls.Password)                                                  // Log password column mapping // 记录密码列映射
	})

	t.Run("TEST-2", func(t *testing.T) {
		// Test efficient access from mutexmap cache
		// 测试高效地从 mutexmap 缓存中检索
		one, cls := gormclasscache.UmcV2(&TestModel{}, cache2)

		// Verify results from mutexmap backend
		// 验证来自 mutexmap 后端的结果
		require.NotNil(t, one, "Model instance should not be nil")                          // 模型实例不应为空
		require.NotNil(t, cls, "Columns struct should not be nil")                          // 列结构体不应为空
		require.Equal(t, "test_models", one.TableName(), "Table name should be consistent") // 表名应保持一致
	})
}

// TestUmcV3 validates sync.Map-based caching implementation (UmcV3)
// Tests Go's built-in concurrent map for column caching
// Evaluates performance characteristics of sync.Map vs custom approaches
//
// TestUmcV3 验证基于 sync.Map 的缓存实现（UmcV3）
// 测试 Go 的内置并发 map 用于列缓存
// 评估 sync.Map 与自定义实现的性能特征
func TestUmcV3(t *testing.T) {
	t.Run("TEST-1", func(t *testing.T) {
		// Test sync.Map cache population with type-safe columns
		// 测试用类型安全的列填充 sync.Map 缓存
		one, cls := gormclasscache.UmcV3(&TestModel{}, cache3)

		// Validate sync.Map caching features
		// 验证 sync.Map 缓存功能
		require.NotNil(t, one, "Model instance should not be nil")                                 // 模型实例不应为空
		require.Equal(t, "test_models", one.TableName(), "Table name should match expected value") // 表名应匹配预期值
		require.NotNil(t, cls, "Columns struct should not be nil")                                 // 列结构体不应为空
		t.Log("✓ Table name:", one.TableName())                                                    // Log table name info // 记录表名信息
		t.Log("✓ Username column:", cls.Username)                                                  // Log username column mapping // 记录用户名列映射
		t.Log("✓ Password column:", cls.Password)                                                  // Log password column mapping // 记录密码列映射
	})

	t.Run("TEST-2", func(t *testing.T) {
		// Test optimal performance access from sync.Map cache
		// 测试以最佳性能从 sync.Map 缓存中检索
		one, cls := gormclasscache.UmcV3(&TestModel{}, cache3)

		// Verify results from sync.Map backend
		// 验证来自 sync.Map 后端的结果
		require.NotNil(t, one, "Model instance should not be nil")                          // 模型实例不应为空
		require.NotNil(t, cls, "Columns struct should not be nil")                          // 列结构体不应为空
		require.Equal(t, "test_models", one.TableName(), "Table name should be consistent") // 表名应保持一致
	})
}
