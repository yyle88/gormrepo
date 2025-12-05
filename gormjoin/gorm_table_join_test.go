package gormjoin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormjoin"
	"github.com/yyle88/gormrepo/gormtablerepo"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath"
)

type Guest struct {
	ID   uint
	Name string
}

func (*Guest) TableName() string {
	return "guests"
}

func (a *Guest) Columns() *GuestColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Guest) TableColumns(decoration gormcnm.ColumnNameDecoration) *GuestColumns {
	return &GuestColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:   gormcnm.Cmn(a.ID, "id", decoration),
		Name: gormcnm.Cmn(a.Name, "name", decoration),
	}
}

type GuestColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
}

type Order struct {
	ID      uint
	GuestID uint
	Amount  float64
}

func (*Order) TableName() string {
	return "orders"
}

func (a *Order) Columns() *OrderColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Order) TableColumns(decoration gormcnm.ColumnNameDecoration) *OrderColumns {
	return &OrderColumns{
		// Auto-generated: column mapping in table operations. DO NOT EDIT. // 自动生成：表操作的列映射。请勿编辑。
		ID:      gormcnm.Cmn(a.ID, "id", decoration),
		GuestID: gormcnm.Cmn(a.GuestID, "guest_id", decoration),
		Amount:  gormcnm.Cmn(a.Amount, "amount", decoration),
	}
}

type OrderColumns struct {
	// Auto-generated: embedding operation functions to make it simple to use. DO NOT EDIT. // 自动生成：嵌入操作函数便于使用。请勿编辑。
	gormcnm.ColumnOperationClass
	// Auto-generated: column names and types in database table. DO NOT EDIT. // 自动生成：数据库表的列名和类型。请勿编辑。
	ID      gormcnm.ColumnName[uint]
	GuestID gormcnm.ColumnName[uint]
	Amount  gormcnm.ColumnName[float64]
}

// Tests the generation of columns for models.
// 测试模型列的生成。
func TestGenerateColumns(t *testing.T) {
	absPath := runpath.Path() // Retrieve the absolute path of the source file based on the current test file's location
	// 获取当前测试文件位置基础上的源文件绝对路径
	t.Log(absPath)

	// Check the existence of the target file. The file should be created beforehand to ensure it can be located.
	// 检查目标文件是否存在。文件应手动创建，确保代码能够找到。
	require.True(t, osmustexist.IsFile(absPath))

	// List the models to have columns generated. Both instance and non-instance types are supported.
	// 设置需要生成列的模型，这里支持地址类型和非地址类型。
	objects := []any{&Guest{}, &Order{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable struct names (e.g., ExampleColumns) // 生成可导出的结构体名称（例如 ExampleColumns）
		WithColumnsMethodRecvName("a").
		WithColumnsCheckFieldType(true).
		WithIsGenFuncTableColumns(true)

	// Configure code generation settings
	// 配置代码生成设置
	cfg := gormcngen.NewConfigs(objects, options, absPath).
		WithIsGenPreventEdit(false)
	cfg.Gen() // Generate and write the code to the target location (e.g., "gormcnm.gen.go") // 生成并将代码写入目标位置（例如 "gormcnm.gen.go"）
}

func TestTableJoin(t *testing.T) {
	{
		repo1 := gormtablerepo.NewTableRepo(gormclass.UseTable(&Guest{}))
		repo2 := gormtablerepo.NewTableRepo(gormclass.UseTable(&Order{}))
		res := gormjoin.RIGHTJOIN(repo1, repo2).On(func(cls1 *GuestColumns, cls2 *OrderColumns) []string {
			return []string{
				cls2.GuestID.OnEq(cls1.ID),
			}
		})
		t.Log(res)
		require.Equal(t, "RIGHT JOIN orders ON orders.guest_id=guests.id", res)
	}
	{
		repo1 := gormtablerepo.NewTableRepo(gormclass.UseTable(&Order{}))
		repo2 := gormtablerepo.NewTableRepo(gormclass.UseTable(&Guest{}))
		res := gormjoin.INNERJOIN(repo1, repo2).On(func(cls1 *OrderColumns, cls2 *GuestColumns) []string {
			return []string{
				cls1.GuestID.OnEq(cls2.ID),
			}
		})
		t.Log(res)
		require.Equal(t, "INNER JOIN guests ON orders.guest_id=guests.id", res)
	}
}
