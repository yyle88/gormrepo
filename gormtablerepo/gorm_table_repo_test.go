package gormtablerepo_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormcnm/gormcnmstub"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormtablerepo"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	ID   uint
	Name string
	Rank int
}

func (*Student) TableName() string {
	return "students"
}

func (a *Student) Columns() *StudentColumns {
	return a.TableColumns(gormcnm.NewPlainDecoration())
}

func (a *Student) TableColumns(decoration gormcnm.ColumnNameDecoration) *StudentColumns {
	return &StudentColumns{
		ID:   gormcnm.Cmn(a.ID, "id", decoration),
		Name: gormcnm.Cmn(a.Name, "name", decoration),
		Rank: gormcnm.Cmn(a.Rank, "rank", decoration),
	}
}

type StudentColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID   gormcnm.ColumnName[uint]
	Name gormcnm.ColumnName[string]
	Rank gormcnm.ColumnName[int]
}

// Tests the generation of columns for models.
// 测试模型列的生成。
func TestGenerateColumns(t *testing.T) {
	absPath := runpath.Path() // Retrieve the absolute path of the source file based on the current test file's location
	// 获取当前测试文件位置基础上的源文件绝对路径
	t.Log(absPath)

	// Verify the existence of the target file. The file should be created manually to ensure it can be located by the code.
	// 检查目标文件是否存在。文件应手动创建，确保代码能够找到它。
	require.True(t, osmustexist.IsFile(absPath))

	// List the models for which columns will be generated. Both pointer and non-pointer types are supported.
	// 设置需要生成列的模型，这里支持指针类型和非指针类型。
	objects := []any{&Student{}}

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

func TestNewTableRepo(t *testing.T) {
	repo := gormtablerepo.NewTableRepo(gormclass.UseTable(&Student{}))
	t.Log(repo.GetTableName())
	require.Equal(t, "students", repo.GetTableName())
}

func TestTableRepo_Gorm_Repo(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&Student{}))

	must.Done(db.Save(&Student{
		ID:   0,
		Name: "A",
		Rank: 100,
	}).Error)
	must.Done(db.Save(&Student{
		ID:   0,
		Name: "B",
		Rank: 85,
	}).Error)

	repo := gormtablerepo.NewTableRepo(gormclass.UseTable(&Student{}))
	t.Run("case-1", func(t *testing.T) {
		studentA, err := repo.Repo(db).First(func(db *gorm.DB, cls *StudentColumns) *gorm.DB {
			return db.Where(cls.Name.Eq("A"))
		})
		require.NoError(t, err)
		require.Equal(t, "A", studentA.Name)
		require.Equal(t, 100, studentA.Rank)
	})
	t.Run("case-2", func(t *testing.T) {
		studentB, err := repo.Repo(db).First(func(db *gorm.DB, cls *StudentColumns) *gorm.DB {
			return db.Where(cls.Name.Eq("B"))
		})
		require.NoError(t, err)
		require.Equal(t, "B", studentB.Name)
		require.Equal(t, 85, studentB.Rank)
	})
}

func TestTableRepo_BuildColumns(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&Student{}))

	must.Done(db.Save(&Student{ID: 0, Name: "A", Rank: 100}).Error)
	must.Done(db.Save(&Student{ID: 0, Name: "B", Rank: 85}).Error)
	must.Done(db.Save(&Student{ID: 0, Name: "C", Rank: 90}).Error)

	type Result struct {
		Name  string
		Score int
	}
	var result = &Result{}
	repo := gormtablerepo.NewTableRepo(gormclass.UseTable(&Student{}))
	err := repo.Repo(db).Mold().Invoke(func(db *gorm.DB, cls *StudentColumns) *gorm.DB {
		return db.Select(
			gormcnmstub.MergeSlices(repo.BuildColumns(func(cls *StudentColumns) []string {
				return []string{
					cls.Name.AsName(gormcnm.Cnm(result.Name, "name")),
					cls.Rank.AsName(gormcnm.Cnm(result.Score, "score")),
				}
			}))).Where(cls.Name.Eq("B")).First(result)
	})
	require.NoError(t, err)
	require.Equal(t, "B", result.Name)
	require.Equal(t, 85, result.Score)
}
