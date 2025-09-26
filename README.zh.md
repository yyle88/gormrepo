[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormrepo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormrepo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormrepo)](https://pkg.go.dev/github.com/yyle88/gormrepo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormrepo/main.svg)](https://coveralls.io/github/yyle88/gormrepo?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormrepo.svg)](https://github.com/yyle88/gormrepo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormrepo)](https://goreportcard.com/report/github.com/yyle88/gormrepo)

# 🚀 GORM 生态系统 - 企业级类型安全数据库操作

**gormrepo** 是完整 GORM 生态系统的**核心组件**，给 Go 开发者提供**类型安全**、**企业级**和**高效**的数据库操作。

> 🌟 **融合 Java MyBatis Plus + Python SQLAlchemy 的精华，与 Go 下一代 ORM 工具链精心设计**

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 🎯 生态系统核心价值

### ✨ 编译时类型安全
- **零运行时错误**：在编译时捕获所有列名和类型错误
- **重构兼容**：字段重命名自动更新所有引用
- **IDE 智能提示**：完整的代码补全和类型检查

### 🔄 智能代码生成
- **AST 级别精度**：基于语法树的智能代码生成
- **零维护成本**：自动生成和更新列常量
- **增量更新**：保留现有代码结构

### 🌍 原生语言支持
- **中文字段名**：支持中文和其他原生语言的业务字段
- **自动转换**：智能生成数据库兼容的列映射
- **国际化兼容**：降低非英语开发者的入门门槛

### 🏢 企业仓储模式
- **CRUD 封装**：开箱即用的常用数据库操作
- **分页支持**：内置分页、计数和排序功能
- **作用域隔离**：优雅的临时变量管理

---

## 🏗️ 架构概览

```
┌─────────────────────────────────────────────────────────────────────┐
│                    GORM Type-Safe Ecosystem                         │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐              │
│  │  gormzhcn   │    │  gormmom    │    │  gormrepo   │              │
│  │ Chinese API │───▶│ Native Lang │───▶│  Package    │─────┐        │
│  │  Localize   │    │  Smart Tags │    │  Pattern    │     │        │
│  └─────────────┘    └─────────────┘    └─────────────┘     │        │
│         │                   │                              │        │
│         │                   ▼                              ▼        │
│         │            ┌─────────────┐              ┌─────────────┐   │
│         │            │ gormcngen   │              │Application  │   │
│         │            │Code Generate│─────────────▶│Custom Code  │   │
│         │            │AST Operation│              │             │   │
│         │            └─────────────┘              └─────────────┘   │
│         │                   │                              ▲        │
│         │                   ▼                              │        │
│         └────────────▶┌─────────────┐◄─────────────────────┘        │
│                       │   GORMCNM   │                               │
│                       │ FOUNDATION  │                               │
│                       │ Type-Safe   │                               │
│                       │ Core Logic  │                               │
│                       └─────────────┘                               │
│                              │                                      │
│                              ▼                                      │
│                       ┌─────────────┐                               │
│                       │    GORM     │                               │
│                       │  Database   │                               │
│                       └─────────────┘                               │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## 📦 生态系统组件

### 🔹 [gormcnm](https://github.com/yyle88/gormcnm) - 类型安全列基础
**核心价值**：消除硬编码列名，实现编译时类型安全
- `ColumnName[T]` 泛型类型定义
- 完整 SQL 操作：`Eq()`、`Gt()`、`Lt()`、`In()`、`Between()` 等
- 表达式构建：`ExprAdd()`、`ExprSub()`、`ExprMul()` 等

### 🔹 [gormcngen](https://github.com/yyle88/gormcngen) - 智能代码生成
**核心价值**：零维护的自动生成 `Columns()` 方法
- AST 语法树分析和精确操作
- 自动生成列结构体和方法
- 支持自定义列映射和嵌入字段

### 🔹 [gormrepo](https://github.com/yyle88/gormrepo) - 企业仓储模式 ⭐
**核心价值**：简化 GORM 操作，提供企业级体验
- 泛型仓储模式 `GormRepo[MOD, CLS]`
- 函数式条件构建
- 完整的分页、计数和存在性检查

### 🔹 [gormmom](https://github.com/yyle88/gormmom) - 原生语言支持
**核心价值**：支持原生语言编程的智能标签生成
- 基于 AST 的自动标签生成和更新
- 智能列名转换策略
- 自动索引名修正

### 🔹 [gormzhcn](https://github.com/go-zwbc/gormzhcn) - 中文编程接口
**核心价值**：完整的中文 API，原生中文开发
- 纯中文方法和类型名（`T编码器`、`T表结构`、`T配置项`）
- 中文字段名支持（`V名称`、`V性别`、`V年龄`）
- 基于 gormmom 构建，完整生态集成

---

## 🚀 快速开始

### 安装

```bash
go get github.com/yyle88/gormrepo
```

### 完整使用流程

#### 1. 定义模型（支持原生字段）

```go
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex" cnm:"username"` 
    Nickname string `gorm:"index" cnm:"nickname"`
    Age      int    `cnm:"age"`
}
```

#### 2. 自动生成列结构体（gormcngen）

```go
// 自动生成
func (*User) Columns() *UserColumns {
    return &UserColumns{
        ID:       "id",
        Username: "username",
        Nickname: "nickname",
        Age:      "age",
    }
}

type UserColumns struct {
    ID       gormcnm.ColumnName[uint]
    Username gormcnm.ColumnName[string]
    Nickname gormcnm.ColumnName[string]
    Age      gormcnm.ColumnName[int]
}
```

#### 3. 类型安全的仓储操作（gormrepo 核心功能）

```go
// 创建仓储
repo := gormrepo.NewGormRepo(db, &User{}, (&User{}).Columns())

// 类型安全查询 - 编译时验证，零运行时错误
user, err := repo.First(func(db *gorm.DB, cls *UserColumns) *gorm.DB {
    return db.Where(cls.Username.Eq("alice")).
              Where(cls.Age.Gte(18))
})

// 批量查询
users, err := repo.Find(func(db *gorm.DB, cls *UserColumns) *gorm.DB {
    return db.Where(cls.Nickname.Like("%admin%"))
})

// 分页查询带总数
users, total, err := repo.FindPageAndCount(
    func(db *gorm.DB, cls *UserColumns) *gorm.DB {
        return db.Where(cls.Age.Between(18, 65))
    },
    func(cls *UserColumns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// 类型安全更新
err = repo.Updates(
    func(db *gorm.DB, cls *UserColumns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *UserColumns) map[string]interface{} {
        return cls.Kw(cls.Age.Kv(25)).
                  Kw(cls.Nickname.Kv("NewNick")).
                  AsMap()
    },
)
```

---

## 💡 核心优势对比

| 特性 | 经典 GORM | GORM 生态系统 |
|---------|-----------------|----------------|
| **取消硬编码** | ❌ "name", "email" 字符串 | ✅ 类型安全列访问 |
| **防止拼写错误** | ❌ 运行时 SQL 错误 | ✅ 编译时错误检测 |
| **防止类型错误** | ❌ 错误的类型赋值 | ✅ 泛型类型强制检查 |
| **重构支持** | ❌ 手动查找替换 | ✅ IDE 自动重构 |
| **原生语言** | ❌ 仅英文字段 | ✅ 支持中文等语言 |
| **代码生成** | ❌ 手动维护 | ✅ AST 智能生成 |

### 经典 vs 生态系统方法

```go
// ❌ 经典方法：容易出错，难以维护
db.Where("username = ?", "alice").
   Where("age >= ?", 18).
   First(&user)

// ✅ 生态系统：类型安全，IDE 智能提示
repo.First(func(db *gorm.DB, cls *UserColumns) *gorm.DB {
    return db.Where(cls.Username.Eq("alice")).
              Where(cls.Age.Gte(18))
})
```

---

## 🔧 GormRepo API 文档

### 查询操作
| 方法 | 参数 | 返回值 | 描述 |
|--------|-----------|---------|-------------|
| `First` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `*MOD, error` | 查询第一个匹配记录 |
| `Find` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `[]*MOD, error` | 查询所有匹配记录 |
| `FindPage` | `where, ordering, pagination` | `[]*MOD, error` | 分页查询 |
| `FindPageAndCount` | `where, ordering, pagination` | `[]*MOD, int64, error` | 分页查询带总数 |
| `Count` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `int64, error` | 统计匹配记录数 |
| `Exist` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `bool, error` | 检查记录是否存在 |

### 创建操作
| 方法 | 参数 | 返回值 | 描述 |
|--------|-----------|---------|-------------|
| `Create` | `one *MOD` | `error` | 创建新记录 |
| `Save` | `one *MOD` | `error` | 插入或更新记录 |

### 更新操作
| 方法 | 参数 | 返回值 | 描述 |
|--------|-----------|---------|-------------|
| `Update` | `where, valueFunc` | `error` | 更新单个字段 |
| `Updates` | `where, mapValues` | `error` | 更新多个字段 |

### 删除操作
| 方法 | 参数 | 返回值 | 描述 |
|--------|-----------|---------|-------------|
| `Delete` | `one *MOD` | `error` | 根据实体删除记录 |
| `DeleteW` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | 根据条件删除记录 |
| `DeleteM` | `one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | 根据实体和条件删除记录 |

---

## 🌟 企业应用场景

### 🏢 大型项目数据库标准化
- 统一类型安全操作标准
- 减少代码审查中的低级错误
- 提高团队协作效率

### 🌍 国际化项目支持
- 原生字段名降低业务理解门槛
- 自动生成标准数据库列名
- 支持多语言团队协作

### ⚡ 快速开发与维护
- 零配置代码生成
- IDE 智能提示和重构支持
- 减少 90% 重复 CRUD 代码

---

## 🔄 技术对比

| 生态系统 | Java MyBatis Plus | Python SQLAlchemy | Go GORM 生态系统 |
|-----------|------------------|-------------------|-------------------|
| **类型安全列** | `Example::getName` | `Example.name` | `cls.Name.Eq()` |
| **代码生成** | ✅ 插件支持 | ✅ 反射机制 | ✅ AST 精度 |
| **仓储模式** | ✅ BaseMapper | ✅ Session API | ✅ GormRepo |
| **原生语言** | 🟡 有限支持 | 🟡 有限支持 | ✅ 完整支持 |

---

## 📝 完整示例

查看 [examples](internal/examples) 目录获取完整集成示例。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 许可证类型

MIT 许可证。详见 [LICENSE](LICENSE)。

---

## 🤝 项目贡献

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **发现问题？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **功能建议？** 创建 issue 讨论您的想法
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
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
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Pull Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Pull Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormrepo.svg?variant=adaptive)](https://starchart.cc/yyle88/gormrepo)

---

## 🔗 相关项目

- 🏗️ **[gormcnm](https://github.com/yyle88/gormcnm)** - 类型安全列基础库
- 🤖 **[gormcngen](https://github.com/yyle88/gormcngen)** - 智能代码生成
- 🏢 **[gormrepo](https://github.com/yyle88/gormrepo)** - 企业仓储模式
- 🌍 **[gormmom](https://github.com/yyle88/gormmom)** - 原生语言编程支持
