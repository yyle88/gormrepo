package gormtablerepo_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"github.com/yyle88/gormrepo/gormtablerepo"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestTableRepo_Repo(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&Student{}))

	must.Done(db.Save(&Student{ID: 0, Name: "A", Rank: 100}).Error)
	must.Done(db.Save(&Student{ID: 0, Name: "B", Rank: 85}).Error)
	must.Done(db.Save(&Student{ID: 0, Name: "C", Rank: 90}).Error)

	repo := gormtablerepo.NewTableRepo(gormclass.UseTable(&Student{}))
	students, count, err := repo.Repo(db).FindPageAndCount(func(db *gorm.DB, cls *StudentColumns) *gorm.DB {
		return db.Where(cls.Name.In([]string{"A", "B", "C"}))
	}, func(cls *StudentColumns) gormcnm.OrderByBottle {
		return cls.Rank.Ob("desc")
	}, &gormrepo.Pagination{
		Limit:  3,
		Offset: 1,
	})
	require.NoError(t, err)
	require.Equal(t, int64(3), count) // Expecting 3 students in aggregate count without pagination
	require.Len(t, students, 2)       // Expecting 2 students in the pagination result
	require.Equal(t, "C", students[0].Name)
	require.Equal(t, "B", students[1].Name)
}

func TestTableRepo_Gorm(t *testing.T) {
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&Student{}))

	must.Done(db.Save(&Student{ID: 0, Name: "A", Rank: 100}).Error)
	must.Done(db.Save(&Student{ID: 0, Name: "B", Rank: 85}).Error)
	must.Done(db.Save(&Student{ID: 0, Name: "C", Rank: 90}).Error)

	repo := gormtablerepo.NewTableRepo(gormclass.UseTable(&Student{}))
	var students []*Student
	pagination := &gormrepo.Pagination{
		Limit:  3,
		Offset: 1,
	}
	err := repo.Gorm(db).Find(func(db *gorm.DB, cls *StudentColumns) *gorm.DB {
		return db.Where(cls.Name.In([]string{"A", "B", "C"})).
			Order(cls.Rank.Ob("desc").Ox()).
			Scopes(pagination.Scope())
	}, &students).Error
	require.NoError(t, err)
	var count int64
	err = repo.Gorm(db).Mold().Invoke(func(db *gorm.DB, cls *StudentColumns) *gorm.DB {
		return db.Where(cls.Name.In([]string{"A", "B", "C"})).Count(&count)
	}).Error
	require.NoError(t, err)
	require.Equal(t, int64(3), count) // Expecting 3 students in aggregate count without pagination
	require.Len(t, students, 2)       // Expecting 2 students in the pagination result
	require.Equal(t, "C", students[0].Name)
	require.Equal(t, "B", students[1].Name)
}
