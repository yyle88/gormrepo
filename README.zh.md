[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormrepo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormrepo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormrepo)](https://pkg.go.dev/github.com/yyle88/gormrepo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormrepo/main.svg)](https://coveralls.io/github/yyle88/gormrepo?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormrepo.svg)](https://github.com/yyle88/gormrepo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormrepo)](https://goreportcard.com/report/github.com/yyle88/gormrepo)

# 🚀 GORM 生态系统 - 企业级类型安全数据库操作

**gormrepo** 是完整 GORM 生态系统的**核心组件**，给 Go 开发者提供**类型安全**、**企业级**和**高效**的数据库操作。

> 🌟 **融合 Java MyBatis Plus + Python SQLAlchemy 的精华，与 Go 下一代 ORM 工具链精心设计**

---

## 生态系统

![GORM Type-Safe Ecosystem](https://github.com/yyle88/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 🔄 技术对比

| 生态系统      | Java MyBatis Plus  | Python SQLAlchemy | Go GORM 生态系统    |
|-----------|--------------------|-------------------|-----------------|
| **类型安全列** | `Example::getName` | `Example.name`    | `cls.Name.Eq()` |
| **代码生成**  | ✅ 插件支持             | ✅ 反射机制            | ✅ AST 精度        |
| **仓储模式**  | ✅ BaseMapper       | ✅ Session API     | ✅ GormRepo      |
| **原生语言**  | 🟡 有限支持            | 🟡 有限支持           | ✅ 完整支持          |

---

## 🚀 快速开始

### 安装

```bash
go get github.com/yyle88/gormrepo
```

### 完整使用流程

#### 1. 定义模型（支持原生字段）

```go
type Account struct {
    ID       uint   `gorm:"primaryKey"`
    Accountname string `gorm:"uniqueIndex" cnm:"accountname"`
    Nickname string `gorm:"index" cnm:"nickname"`
    Age      int    `cnm:"age"`
}
```

#### 2. 自动生成列结构体（gormcngen）

```go
// 自动生成
func (*Account) Columns() *AccountColumns {
    return &AccountColumns{
        ID:       "id",
        Accountname: "accountname",
        Nickname: "nickname",
        Age:      "age",
    }
}

type AccountColumns struct {
    ID       gormcnm.ColumnName[uint]
    Accountname gormcnm.ColumnName[string]
    Nickname gormcnm.ColumnName[string]
    Age      gormcnm.ColumnName[int]
}
```

#### 3. 创建 Repo

```go
// 创建 repo，传入列定义
repo := gormrepo.NewRepo(&Account{}, (&Account{}).Columns())

// gormrepo/gormclass 简洁写法
repo := gormrepo.NewRepo(gormclass.Use(&Account{}))
```

#### 4. 如何查询

```go
// 经典 GORM
err := db.Where("name = ?", "alice").First(&account).Error

// gormrepo - First
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Accountname.Eq("alice"))
})

// gormrepo - FirstE (返回 ErrorOrNotExist)
account, erb := repo.With(ctx, db).FirstE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Accountname.Eq("alice"))
})

// gormrepo - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18))
})

// gormrepo - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.Age.Gte(18))
    },
    func(cls *AccountColumns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// gormrepo - FindPageAndCount
accounts, count, err := repo.With(ctx, db).FindPageAndCount(
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.Age.Between(18, 65))
    },
    func(cls *AccountColumns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// gormrepo - Count
count, err := repo.With(ctx, db).Count(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18))
})

// gormrepo - Exist
exist, err := repo.With(ctx, db).Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Accountname.Eq("alice"))
})

// gormrepo - Where 条件: Eq, Ne, Gt, Lt, Gte, Lte
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gt(18)).    // age > 18
              Where(cls.Age.Lt(60)).    // age < 60
              Where(cls.Age.Ne(30))     // age != 30
})

// gormrepo - Like / NotLike
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Nickname.Like("%test%")).
              Where(cls.Accountname.NotLike("%admin%"))
})

// gormrepo - In / NotIn
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.ID.In([]uint{1, 2, 3})).
              Where(cls.Age.NotIn([]int{18, 19, 20}))
})

// gormrepo - Between / NotBetween
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Between(20, 40)).
              Where(cls.ID.NotBetween(100, 200))
})

// gormrepo - IsNull / IsNotNull
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Nickname.IsNotNull())
})

// gormrepo - Or 条件
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18)).
              Or(cls.Nickname.Eq("vip"))
})

// gormrepo - Order / Select
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18)).
              Order(cls.Age.Name() + " DESC").
              Select(cls.ID.Name(), cls.Accountname.Name())
})
```

#### 查询操作

| 方法                 | 参数                                          | 返回值                      | 描述           |
|--------------------|---------------------------------------------|--------------------------|--------------|
| `First`            | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `*MOD, error`            | 查询第一个匹配记录    |
| `FirstE`           | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `*MOD, *ErrorOrNotExist` | 查询首条或返回不存在标识 |
| `Find`             | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `[]*MOD, error`          | 查询所有匹配记录     |
| `FindPage`         | `where, ordering, pagination`               | `[]*MOD, error`          | 分页查询         |
| `FindPageAndCount` | `where, ordering, pagination`               | `[]*MOD, int64, error`   | 分页查询带总数      |
| `Count`            | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `int64, error`           | 统计匹配记录数      |
| `Exist`            | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `bool, error`            | 检查记录是否存在     |

#### 5. 如何创建

```go
// 经典 GORM
err := db.Create(&Account{Accountname: "bob", Nickname: "Bob", Age: 25}).Error

// gormrepo
err := repo.With(ctx, db).Create(&Account{Accountname: "bob", Nickname: "Bob", Age: 25})
```

#### 创建操作

| 方法       | 参数         | 返回值     | 描述      |
|----------|------------|---------|---------|
| `Create` | `one *MOD` | `error` | 创建新记录   |
| `Save`   | `one *MOD` | `error` | 插入或更新记录 |

#### 6. 如何更新

```go
// 经典 GORM
err := db.Model(&Account{}).Where("id = ?", 1).Updates(map[string]interface{}{"age": 26}).Error

// gormrepo - Updates
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *AccountColumns) map[string]interface{} {
        return cls.Kw(cls.Age.Kv(26)).
                  Kw(cls.Nickname.Kv("NewNick")).
                  AsMap()
    },
)

// gormrepo - UpdatesM (使用 ColumnValueMap)
err := repo.With(ctx, db).UpdatesM(
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *AccountColumns) *gormcnm.ColumnValueMap {
        return cls.Kw(cls.Age.Kv(26)).
                  Kw(cls.Nickname.Kv("NewNick"))
    },
)

// gormrepo - UpdatesO (通过主键)
account := &Account{ID: 1}
err := repo.With(ctx, db).UpdatesO(account, func(cls *AccountColumns) *gormcnm.ColumnValueMap {
    return cls.Kw(cls.Age.Kv(26))
})

// gormrepo - UpdatesC (组合条件)
account := &Account{ID: 1}
err := repo.With(ctx, db).UpdatesC(account,
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.Age.Lt(30))
    },
    func(cls *AccountColumns) *gormcnm.ColumnValueMap {
        return cls.Kw(cls.Age.Kv(26))
    },
)
```

#### 更新操作

| 方法         | 参数                         | 返回值     | 描述                           |
|------------|----------------------------|---------|------------------------------|
| `Update`   | `where, valueFunc`         | `error` | 更新单个字段                       |
| `Updates`  | `where, mapValues`         | `error` | 更新多个字段                       |
| `UpdatesM` | `where, newValues`         | `error` | 使用 ColumnValueMap 更新 (M=Map) |
| `UpdatesO` | `object, newValues`        | `error` | 通过主键更新 (O=Object)            |
| `UpdatesC` | `object, where, newValues` | `error` | 组合条件更新 (C=Combined)          |

#### 7. 如何删除

```go
// 经典 GORM
err := db.Where("id = ?", 1).Delete(&Account{}).Error

// gormrepo - Delete (通过实例)
account := &Account{ID: 1}
err := repo.With(ctx, db).Delete(account)

// gormrepo - DeleteW (通过条件)
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})

// gormrepo - DeleteM (实例 + 条件)
account := &Account{ID: 1}
err := repo.With(ctx, db).DeleteM(account, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Lt(30))
})
```

#### 删除操作

| 方法        | 参数                                                    | 返回值     | 描述        |
|-----------|-------------------------------------------------------|---------|-----------|
| `Delete`  | `one *MOD`                                            | `error` | 通过实例删除记录  |
| `DeleteW` | `where func(db *gorm.DB, cls CLS) *gorm.DB`           | `error` | 通过条件删除记录  |
| `DeleteM` | `one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | 实例加条件删除记录 |

#### 8. 自定义操作

遇到未涵盖的操作，使用 `Invoke`：

```go
// Invoke - 批量更新
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18)).Update(cls.Nickname.Name(), "adult")
})

// Invoke - 字段自增
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.ID.Eq(1)).
        Update(cls.Age.Name(), gorm.Expr(cls.Age.Name()+" + ?", 1))
})

// Invoke - 选取特定列
var names []string
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18)).Pluck(cls.Accountname.Name(), &names)
})

// Invoke - 复杂条件
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Between(18, 60)).
        Where(cls.Nickname.IsNotNull()).
        Updates(map[string]interface{}{
            cls.Nickname.Name(): "active",
        })
})
```

#### 调用操作

| 方法       | 参数                                           | 返回值     | 描述      |
|----------|----------------------------------------------|---------|---------|
| `Invoke` | `clsRun func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | 执行自定义操作 |

#### 9. 使用 Clauses 实现 Upsert

使用 `Clauses` 或 `Clause` 实现 upsert（冲突时插入或更新）：

```go
// Clauses - 直接传入 clause.Expression
cls := account.Columns()
err := repo.With(ctx, db).Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: cls.Username.Name()}},
    DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
}).Create(&account)

// Clause - 类型安全，使用列定义构建
err := repo.With(ctx, db).Clause(func(cls *AccountColumns) clause.Expression {
    return clause.OnConflict{
        Columns:   []clause.Column{{Name: cls.Username.Name()}},
        DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
    }
}).Create(&account)
```

#### 子句操作

| 方法        | 参数                                          | 返回值         | 描述              |
|-----------|---------------------------------------------|-------------|-----------------|
| `Clauses` | `clauses ...clause.Expression`              | `*GormRepo` | 添加子句并返回新 repo   |
| `Clause`  | `func(cls CLS) clause.Expression`           | `*GormRepo` | 使用列定义构建子句       |

---

## 📝 完整示例

查看 [examples](internal/examples) 目录获取完整集成示例。

---

## 关联项目

探索完整的 GORM 生态系统集成包：

### 核心生态

- **[gormcnm](https://github.com/yyle88/gormcnm)** - GORM 基础层，提供类型安全的列操作和条件构建
- **[gormcngen](https://github.com/yyle88/gormcngen)** - 使用 AST 的代码生成引擎，支持类型安全的 GORM 操作
- **[gormrepo](https://github.com/yyle88/gormrepo)** - 仓储模式实现，遵循 GORM 最佳实践（本项目）
- **[gormmom](https://github.com/yyle88/gormmom)** - 原生语言 GORM 标签生成引擎，支持智能列名
- **[gormzhcn](https://github.com/go-zwbc/gormzhcn)** - 完整的 GORM 中文编程接口

每个包针对 GORM 开发的不同方面，从本地化到类型安全操作和代码生成。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/gormrepo.svg?variant=adaptive)](https://starchart.cc/yyle88/gormrepo)
