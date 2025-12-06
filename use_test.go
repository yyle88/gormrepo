package gormrepo_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
	"gorm.io/gorm"
)

// Account is the test struct with common fields
// Account 是带有常用字段的测试结构体
type Account struct {
	gorm.Model
	Username string `gorm:"unique;"`
	Password string `gorm:"size:255;"`
	Nickname string `gorm:"column:nickname;"`
}

// newAccount creates a new Account instance with the given username
// newAccount 使用给定的用户名创建新的 Account 实例
func newAccount(username string) *Account {
	return &Account{
		Model:    gorm.Model{},
		Username: username,
		Password: uuid.New().String(),
		Nickname: uuid.New().String(),
	}
}

// TableName returns the database table name
// TableName 返回数据库表名
func (*Account) TableName() string {
	return "accounts"
}

// Columns returns the column definitions with type-safe column names
// Columns 返回带有类型安全列名的列定义

func (a *Account) Columns() *AccountColumns {
	return &AccountColumns{
		// Auto-generated: column names and types mapping. DO NOT EDIT. // 自动生成：列名和类型映射。请勿编辑。
		ID:        gormcnm.Cnm(a.ID, "id"),
		CreatedAt: gormcnm.Cnm(a.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(a.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(a.DeletedAt, "deleted_at"),
		Username:  gormcnm.Cnm(a.Username, "username"),
		Password:  gormcnm.Cnm(a.Password, "password"),
		Nickname:  gormcnm.Cnm(a.Nickname, "nickname"),
	}
}

// AccountColumns contains type-safe column definitions
// AccountColumns 包含类型安全的列定义

type AccountColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Username  gormcnm.ColumnName[string]
	Password  gormcnm.ColumnName[string]
	Nickname  gormcnm.ColumnName[string]
}

// TestGenerateColumns tests the column generation
// TestGenerateColumns 测试列生成
func TestGenerateColumns(t *testing.T) {
	absPath := runpath.Path() // Retrieve the absolute path of the source file based on the current test file's location
	// 获取当前测试文件位置基础上的源文件绝对路径
	t.Log(absPath)

	// Check the existence of the target file. The file should be created beforehand to ensure it can be located.
	// 检查目标文件是否存在。文件应手动创建，确保代码能够找到。
	require.True(t, osmustexist.IsFile(absPath))

	// List the models to have columns generated. Both instance and non-instance types are supported.
	// 设置需要生成列的模型，这里支持指针类型和非指针类型。
	objects := []any{&Account{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable struct names (e.g., ExampleColumns) // 生成可导出的结构体名称（例如 ExampleColumns）
		WithColumnsMethodRecvName("a").
		WithColumnsCheckFieldType(true)

	// Configure code generation settings
	// 配置代码生成设置
	cfg := gormcngen.NewConfigs(objects, options, absPath).
		WithIsGenPreventEdit(false)
	cfg.Gen() // Generate and write the code to the target location (e.g., "gormcnm.gen.go") // 生成并将代码写入目标位置（例如 "gormcnm.gen.go"）
}

// TestUse tests the Use function to get db, model, and columns
// TestUse 测试 Use 函数获取 db、model 和 columns
func TestUse(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))
	require.NotNil(t, repo)
}

// TestUmc tests the Umc function (alias to Use)
// TestUmc 测试 Umc 函数（Use 的别名）
func TestUmc(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))
	require.NotNil(t, repo)
}
