package example03_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example03/internal/models"
	"github.com/yyle88/must"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	// 创建内存数据库
	dsn := fmt.Sprintf("file:db-%s?mode=memory&cache=shared", uuid.New().String())
	db := rese.P1(gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	// 自动迁移表结构
	done.Done(db.AutoMigrate(&models.User{}))

	// 插入测试数据
	testUsers := []*models.User{
		{Name: "张三", Email: "zhangsan@example.com", Age: 25, Status: "active"},
		{Name: "李四", Email: "lisi@example.com", Age: 30, Status: "active"},
		{Name: "王五", Email: "wangwu@example.com", Age: 35, Status: "inactive"},
	}
	for _, user := range testUsers {
		must.Done(db.Create(user).Error)
	}

	testDB = db
	m.Run()
}

// TestBasicGormRepoUsage 演示 gormrepo 的基础使用方法
func TestBasicGormRepoUsage(t *testing.T) {
	// 1. 创建 GormRepo 实例 - 这是使用 gormrepo 的核心
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.User{}))

	// 2. First - 查询单条记录
	user, err := repo.First(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Name.Eq("张三"))
	})
	require.NoError(t, err)
	require.Equal(t, "张三", user.Name)
	require.Equal(t, "zhangsan@example.com", user.Email)
	t.Logf("Found user: %+v", user)

	// 3. Find - 查询多条记录
	users, err := repo.Find(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Status.Eq("active"))
	})
	require.NoError(t, err)
	require.Len(t, users, 2) // 应该有2个active用户
	t.Logf("Found %d active users", len(users))

	// 4. Count - 计算记录数量
	count, err := repo.Count(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Age.Gt(25))
	})
	require.NoError(t, err)
	require.Equal(t, int64(2), count) // 年龄>25的用户有2个
	t.Logf("Count of users with age > 25: %d", count)

	// 5. Exist - 检查记录是否存在
	exists, err := repo.Exist(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Email.Eq("zhangsan@example.com"))
	})
	require.NoError(t, err)
	require.True(t, exists)
	t.Logf("User with email zhangsan@example.com exists: %v", exists)

	// 6. Where - 构建查询，可用于复杂查询
	var customUser models.User
	err = repo.Where(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Name.Like("%三%")).Order(cls.ID.OrderByBottle("desc").Ox())
	}).First(&customUser).Error
	require.NoError(t, err)
	require.Equal(t, "张三", customUser.Name)
	t.Logf("Custom query result: %+v", customUser)
}

// TestUpdateOperations 演示更新操作
func TestUpdateOperations(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.User{}))

	// 1. Update - 更新单个字段
	err := repo.Update(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Name.Eq("张三"))
	}, func(cls *models.UserColumns) (string, interface{}) {
		return cls.Age.Kv(26) // 更新年龄为26
	})
	require.NoError(t, err)

	// 验证更新结果
	user, err := repo.First(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Name.Eq("张三"))
	})
	require.NoError(t, err)
	require.Equal(t, 26, user.Age)
	t.Logf("Updated user age: %d", user.Age)

	// 2. Updates - 更新多个字段
	err = repo.Updates(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Name.Eq("李四"))
	}, func(cls *models.UserColumns) map[string]interface{} {
		return cls.
			Kw(cls.Age.Kv(31)).       // 更新年龄
			Kw(cls.Status.Kv("vip")). // 更新状态
			AsMap()
	})
	require.NoError(t, err)

	// 验证批量更新结果
	user, err = repo.First(func(db *gorm.DB, cls *models.UserColumns) *gorm.DB {
		return db.Where(cls.Name.Eq("李四"))
	})
	require.NoError(t, err)
	require.Equal(t, 31, user.Age)
	require.Equal(t, "vip", user.Status)
	t.Logf("Batch updated user: Age=%d, Status=%s", user.Age, user.Status)
}
