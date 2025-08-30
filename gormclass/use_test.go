package gormclass_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Account struct {
	gorm.Model
	Username string `gorm:"unique;"`
	Password string `gorm:"size:255;"`
	Nickname string `gorm:"column:nickname;"`
}

func (*Account) TableName() string {
	return "accounts"
}

func (a *Account) Columns() *AccountColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Account) TableColumns(decoration gormcnm.ColumnNameDecoration) *AccountColumns {
	return &AccountColumns{
		ID:        gormcnm.Cmn(a.ID, "id", decoration),
		CreatedAt: gormcnm.Cmn(a.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(a.UpdatedAt, "updated_at", decoration),
		DeletedAt: gormcnm.Cmn(a.DeletedAt, "deleted_at", decoration),
		Username:  gormcnm.Cmn(a.Username, "username", decoration),
		Password:  gormcnm.Cmn(a.Password, "password", decoration),
		Nickname:  gormcnm.Cmn(a.Nickname, "nickname", decoration),
	}
}

type AccountColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Username  gormcnm.ColumnName[string]
	Password  gormcnm.ColumnName[string]
	Nickname  gormcnm.ColumnName[string]
}

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"unique;"`
	Age       int       `gorm:"type:int32;index:idx_example_age;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (*Example) TableName() string {
	return "examples"
}

func (a *Example) Columns() *ExampleColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Example) TableColumns(decoration gormcnm.ColumnNameDecoration) *ExampleColumns {
	return &ExampleColumns{
		ID:        gormcnm.Cmn(a.ID, "id", decoration),
		Name:      gormcnm.Cmn(a.Name, "name", decoration),
		Age:       gormcnm.Cmn(a.Age, "age", decoration),
		CreatedAt: gormcnm.Cmn(a.CreatedAt, "created_at", decoration),
		UpdatedAt: gormcnm.Cmn(a.UpdatedAt, "updated_at", decoration),
	}
}

type ExampleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Age       gormcnm.ColumnName[int]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}

// Tests the generation of columns for models.
// 测试模型列的生成。
func TestGenerateColumns(t *testing.T) {
	// Retrieve the absolute path of the source file based on the current test file's location
	// 获取当前测试文件位置基础上的源文件绝对路径
	absPath := runpath.Path()
	t.Log(absPath)

	// Verify the existence of the target file. The file should be created manually to ensure it can be located by the code.
	// 检查目标文件是否存在。文件应手动创建，确保代码能够找到它。
	require.True(t, osmustexist.IsFile(absPath))

	// List the models for which columns will be generated. Both pointer and non-pointer types are supported.
	// 设置需要生成列的模型，这里支持指针类型和非指针类型。
	objects := []any{&Account{}, &Example{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable struct names (e.g., ExampleColumns) // 生成可导出的结构体名称（例如 ExampleColumns）
		WithColumnsMethodRecvName("a").
		WithColumnsCheckFieldType(true).
		WithIsGenFuncTableColumns(true)

	// Configure code generation settings
	// 配置代码生成设置
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	// Generate and write the code to the target location (e.g., "gormcnm.gen.go")
	// 生成并将代码写入目标位置（例如 "gormcnm.gen.go"）
	cfg.Gen()
}

func TestUseWithAccount(t *testing.T) {
	if account, cls := gormclass.Use(&Account{}); cls.OK() {
		t.Logf("TableName: %s", account.TableName())
		t.Logf("Columns: %s", neatjsons.S(cls))
	}
}

func TestUseWithExample(t *testing.T) {
	if example, cls := gormclass.Use(&Example{}); cls.OK() {
		t.Logf("TableName: %s", example.TableName())
		t.Logf("Columns: %s", neatjsons.S(cls))
	}
}

func TestAccountAndExample(t *testing.T) {
	var account Account
	if cls := account.Columns(); cls.OK() {
		t.Log(cls.Username)
	}

	var example Example
	if cls := example.Columns(); cls.OK() {
		t.Log(cls.Name)
	}

	t.Logf("Account TableName: %s", account.TableName())
	t.Logf("Example TableName: %s", example.TableName())
}

func TestColumnsWithAccount(t *testing.T) {
	cls := gormclass.Cls(&Account{})
	require.True(t, cls.OK())
	t.Logf("Account Columns: %s", neatjsons.S(cls))
}

func TestOnePointerOutput(t *testing.T) {
	{
		var account Account
		one := gormclass.One(&account)
		require.Equal(t, reflect.Ptr, reflect.TypeOf(one).Kind())
	}

	{
		example := Example{}
		one := gormclass.One(&example)
		require.Equal(t, reflect.Ptr, reflect.TypeOf(one).Kind())
	}
}

func TestUmsWithExample(t *testing.T) {
	examples := gormclass.Ums(&Example{})
	t.Logf("Ums result: %s", neatjsons.S(examples))
	t.Logf("result cap: %d", cap(examples))
}

func TestUssWithExample(t *testing.T) {
	examples := gormclass.Uss[*Example]()
	t.Logf("Uss result: %s", neatjsons.S(examples))
	t.Logf("result cap: %d", cap(examples))
}

func TestUsnWithExample(t *testing.T) {
	examples := gormclass.Usn[*Example](100)
	t.Logf("Usn result: %s", neatjsons.S(examples))
	t.Logf("result cap: %d", cap(examples))
}

func TestUscWithExample(t *testing.T) {
	examples, cls := gormclass.Usc(&Example{})
	require.True(t, cls.OK())
	t.Logf("Usc result: %s", neatjsons.S(examples))
}

func TestMscWithExample(t *testing.T) {
	one, examples, cls := gormclass.Msc(&Example{})
	require.True(t, cls.OK())
	t.Logf("Msc TableName: %s", one.TableName())
	t.Logf("Msc examples: %s", neatjsons.S(examples))
	t.Logf("result cap: %d", cap(examples))
}

func TestNscWithExample(t *testing.T) {
	one, examples, cls := gormclass.Nsc(&Example{}, 32)
	require.True(t, cls.OK())
	t.Logf("Msc TableName: %s", one.TableName())
	t.Logf("Msc examples: %s", neatjsons.S(examples))
	t.Logf("result cap: %d", cap(examples))
	require.Equal(t, 32, cap(examples))
}

func TestExample(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&Example{}))

	example1 := &Example{
		ID:        0,
		Name:      "aaa",
		Age:       1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	example2 := &Example{
		ID:        0,
		Name:      "bbb",
		Age:       2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	require.NoError(t, db.Create(example1).Error)
	require.NoError(t, db.Create(example2).Error)

	var resA Example
	if cls := gormclass.Cls(&Example{}); cls.OK() {
		require.NoError(t, db.Table(resA.TableName()).Where(cls.Name.Eq("aaa")).First(&resA).Error)
		require.Equal(t, "aaa", resA.Name)
	}
	t.Log("res.name:", resA.Name)

	var maxAge int
	if one, cls := gormclass.Use(&Example{}); cls.OK() {
		require.NoError(t, db.Model(one).Where(cls.Age.Gt(0)).Select(cls.Age.COALESCE().MaxStmt("age_alias")).First(&maxAge).Error)
		require.Equal(t, 2, maxAge)
	}
	t.Log("max_age:", maxAge)

	if one, cls := gormclass.Use(&Example{}); cls.OK() {
		require.NoError(t, db.Model(one).Where(cls.Name.Eq("bbb")).Update(cls.Age.Kv(18)).Error)
		require.Equal(t, 18, one.Age)
	}

	var resB Example
	if cls := resB.Columns(); cls.OK() {
		require.NoError(t, db.Table(resB.TableName()).Where(cls.Name.Eq("bbb")).Update(cls.Age.KeAdd(2)).Error)

		require.NoError(t, db.Table(resB.TableName()).Where(cls.Name.Eq("bbb")).First(&resB).Error)
		require.Equal(t, 20, resB.Age)
	}
	t.Log(resB)
}
