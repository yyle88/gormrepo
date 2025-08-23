package example10_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormjoin"
	"github.com/yyle88/gormrepo/gormtablerepo"
	"github.com/yyle88/gormrepo/internal/examples/example10/internal/models"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	db := rese.P1(gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// 自动迁移所有表
	done.Done(db.AutoMigrate(
		&models.Author{},
		&models.Book{},
	))

	// 创建作者
	authors := []*models.Author{
		{Name: "J.R.R. Tolkien", Email: "tolkien@middle-earth.com", Bio: "Creator of Middle-earth", Country: "UK"},
		{Name: "Isaac Asimov", Email: "asimov@foundation.com", Bio: "Science fiction master", Country: "USA"},
	}

	for _, author := range authors {
		done.Done(db.Create(author).Error)
	}

	// 创建书籍
	books := []*models.Book{
		{
			Title:       "The Lord of the Rings",
			ISBN:        "978-0544003415",
			Price:       35.00,
			PublishedAt: "1954-07-29",
			AuthorID:    authors[0].ID,
		},
		{
			Title:       "Foundation",
			ISBN:        "978-0553293357",
			Price:       15.99,
			PublishedAt: "1951-05-01",
			AuthorID:    authors[1].ID,
		},
	}

	for _, book := range books {
		done.Done(db.Create(book).Error)
	}

	testDB = db
	m.Run()
}

// TestJoinQuery 演示连表查询
func TestJoinQuery(t *testing.T) {
	// 查询包含作者信息的书籍（JOIN查询）
	type BookWithAuthor struct {
		BookTitle  string `gorm:"column:book_title"`
		AuthorName string `gorm:"column:author_name"`
	}

	bookRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Book{}))
	authorRepo := gormtablerepo.NewTableRepo(gormclass.UseTable(&models.Author{}))

	var books []*BookWithAuthor
	bookCls := bookRepo.TableColumns()
	authorCls := authorRepo.TableColumns()

	var common = &gormcnm.ColumnOperationClass{}

	err := testDB.Table(bookRepo.GetTableName()).
		Select(common.MergeStmts(
			bookCls.Title.AsName(gormcnm.ColumnName[string]("book_title")),
			authorCls.Name.AsName(gormcnm.ColumnName[string]("author_name")),
		)).
		Joins(gormjoin.LEFTJOIN(bookRepo, authorRepo).On(func(book *models.BookColumns, author *models.AuthorColumns) []string {
			return []string{author.ID.OnEq(book.AuthorID)}
		})).
		Scan(&books).Error
	require.NoError(t, err)
	require.Greater(t, len(books), 0)

	t.Logf("Found %d books with authors", len(books))
	for _, book := range books {
		t.Logf("  Book: %s by %s",
			book.BookTitle, book.AuthorName)
	}
}
