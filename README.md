[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormrepo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormrepo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormrepo)](https://pkg.go.dev/github.com/yyle88/gormrepo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormrepo/main.svg)](https://coveralls.io/github/yyle88/gormrepo?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormrepo.svg)](https://github.com/yyle88/gormrepo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormrepo)](https://goreportcard.com/report/github.com/yyle88/gormrepo)

# 🚀 GORM Ecosystem - Enterprise-Grade Type-Safe Database Operations

**gormrepo** is the centerpiece of a complete GORM ecosystem, delivering **type-safe**, **enterprise-grade**, and **very efficient** database operations for Go developers.

> 🌟 **Combining the best of Java MyBatis Plus + Python SQLAlchemy, designed with Go's next-generation ORM toolchain**

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## 🎯 Ecosystem Core Values

### ✨ Compile-Time Type Safe
- **Zero runtime errors**: Catch every column name and type errors at compile time
- **Refactoring-compatible**: Field renames auto update every reference
- **IDE intelligence**: Complete code completion and type checking

### 🔄 Intelligent Code Generation
- **AST precision**: Smart code generation based on syntax trees
- **Zero maintenance cost**: Auto-generate and update column constants
- **Progressive updates**: Preserve existing code structure

### 🌍 Native Language Support
- **Chinese field names**: Support Chinese and native languages for business fields
- **Automatic conversion**: Smart generation of database-compatible column mappings
- **Worldwide-compatible**: Reduce barriers for non-English developers

### 🏢 Enterprise Repo Pattern
- **CRUD encapsulation**: Out-of-the-box common database operations
- **Pagination support**: Built-in pagination, counting, and sorting
- **Scope isolation**: Elegant temp variable management

---

## 🏗️ Architecture Overview

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

## 📦 Ecosystem Components

### 🔹 [gormcnm](https://github.com/yyle88/gormcnm) - Type-Safe Column Foundation
**Core Value**: Eliminate hardcoded column names, achieve compile-time type safe
- `ColumnName[T]` generic type definition
- Complete SQL operations: `Eq()`, `Gt()`, `Lt()`, `In()`, `Between()`, etc.
- Expression building: `ExprAdd()`, `ExprSub()`, `ExprMul()`, etc.

### 🔹 [gormcngen](https://github.com/yyle88/gormcngen) - Smart Code Generation
**Core Value**: Auto-generate `Columns()` methods with zero maintenance
- AST syntax tree analysis and precise operations
- Auto-generate column structs and methods
- Support custom column mappings and embedded fields

### 🔹 [gormrepo](https://github.com/yyle88/gormrepo) - Enterprise Repo Pattern ⭐
**Core Value**: Streamline GORM operations with enterprise-grade experience
- Generic repo pattern `GormRepo[MOD, CLS]`
- Flexible condition building
- Complete pagination, counting, and existence checks

### 🔹 [gormmom](https://github.com/yyle88/gormmom) - Native Language Support
**Core Value**: Smart tag generation supporting native language programming
- AST-based automatic tag generation and updates
- Intelligent column name conversion strategies
- Automatic index name correction

### 🔹 [gormzhcn](https://github.com/go-zwbc/gormzhcn) - Chinese Programming Interface
**Core Value**: Complete Chinese API for native Chinese development
- Pure Chinese method and type names (`T编码器`, `T表结构`, `T配置项`)
- Chinese field name support (`V名称`, `V性别`, `V年龄`)
- Built on gormmom with complete ecosystem integration

---

## 🚀 Quick Start

### Installation

```bash
go get github.com/yyle88/gormrepo
```

### Complete Usage Flow

#### 1. Define The Model (Supporting Native Fields)

```go
type User struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"uniqueIndex" cnm:"username"` 
    Nickname string `gorm:"index" cnm:"nickname"`
    Age      int    `cnm:"age"`
}
```

#### 2. Auto-Generate Column Structs (gormcngen)

```go
// Auto-generated
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

#### 3. Type-Safe Repo Operations (gormrepo Core Features)

```go
// Create repo
repo := gormrepo.NewGormRepo(db, &User{}, (&User{}).Columns())

// Type-safe queries - compile-time validation, zero runtime errors
user, err := repo.First(func(db *gorm.DB, cls *UserColumns) *gorm.DB {
    return db.Where(cls.Username.Eq("alice")).
              Where(cls.Age.Gte(18))
})

// Batch queries
users, err := repo.Find(func(db *gorm.DB, cls *UserColumns) *gorm.DB {
    return db.Where(cls.Nickname.Like("%admin%"))
})

// Pagination with count
users, total, err := repo.FindPageAndCount(
    func(db *gorm.DB, cls *UserColumns) *gorm.DB {
        return db.Where(cls.Age.Between(18, 65))
    },
    func(cls *UserColumns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// Type-safe updates
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

## 💡 Core Advantages Comparison

| Feature | Classic GORM | GORM Ecosystem |
|---------|-----------------|----------------|
| **Hardcoded Strings** | ❌ "name", "email" literals | ✅ Type-safe column access |
| **Typo Prevention** | ❌ Runtime SQL errors | ✅ Compile-time error detection |
| **Type Validation** | ❌ Wrong type assignments | ✅ Generic type enforcement |
| **Refactoring** | ❌ Hand-written find-replace | ✅ IDE auto-update |
| **Native Language** | ❌ English fields just | ✅ Support Chinese/others |
| **Code Generation** | ❌ Hand-written maintenance | ✅ AST smart generation |

### Classic vs Ecosystem Approach

```go
// ❌ Classic: Error-prone, hard to maintain
db.Where("username = ?", "alice").
   Where("age >= ?", 18).
   First(&user)

// ✅ Ecosystem: Type-safe, IDE intelligent hints
repo.First(func(db *gorm.DB, cls *UserColumns) *gorm.DB {
    return db.Where(cls.Username.Eq("alice")).
              Where(cls.Age.Gte(18))
})
```

---

## 🔧 GormRepo API Documentation

### Query Operations
| Method | Parameters | Returns | Description |
|--------|-----------|---------|-------------|
| `First` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `*MOD, error` | Query first matching record |
| `Find` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `[]*MOD, error` | Query every matching records |
| `FindPage` | `where, ordering, pagination` | `[]*MOD, error` | Paginated search |
| `FindPageAndCount` | `where, ordering, pagination` | `[]*MOD, int64, error` | Paginated search with total count |
| `Count` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `int64, error` | Count matching records |
| `Exist` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `bool, error` | Check if records exist |

### Create Operations
| Method | Parameters | Returns | Description |
|--------|-----------|---------|-------------|
| `Create` | `one *MOD` | `error` | Create new record |
| `Save` | `one *MOD` | `error` | Insert or update record |

### Update Operations
| Method | Parameters | Returns | Description |
|--------|-----------|---------|-------------|
| `Update` | `where, valueFunc` | `error` | Update single field |
| `Updates` | `where, mapValues` | `error` | Update multiple fields |

### Delete Operations
| Method | Parameters | Returns | Description |
|--------|-----------|---------|-------------|
| `Delete` | `one *MOD` | `error` | Delete record by entity |
| `DeleteW` | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | Delete by conditions |
| `DeleteM` | `one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | Delete item with conditions |

---

## 🌟 Enterprise Use Cases

### 🏢 Large Project Database Standardization
- Unified type-safe operation standards
- Reduce low-grade errors in code reviews
- Improve team collaboration speed

### 🌍 Worldwide Project Support
- Native field names reduce business understanding barriers
- Auto-generate standard database column names
- Support multi-language team collaboration

### ⚡ Rapid Development & Maintenance
- Zero-configuration code generation
- IDE intelligent hints and refactoring support
- Reduce 90% repetitive CRUD code

---

## 🔄 Tech Comparison

| Ecosystem | Java MyBatis Plus | Python SQLAlchemy | Go GORM Ecosystem |
|-----------|------------------|-------------------|-------------------|
| **Type-Safe Columns** | `Example::getName` | `Example.name` | `cls.Name.Eq()` |
| **Code Generation** | ✅ Plugin support | ✅ Reflection | ✅ AST precision |
| **Repo Pattern** | ✅ BaseMapper | ✅ Session API | ✅ GormRepo |
| **Native Language** | 🟡 Limited | 🟡 Limited | ✅ Complete support |

---

## 📝 Complete Examples

Check [examples](internal/examples) DIR for complete integration examples.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-06 04:53:24.895249 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a bug?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a pull request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![starring](https://starchart.cc/yyle88/gormrepo.svg?variant=adaptive)](https://starchart.cc/yyle88/gormrepo)

---

## 🔗 Related Projects

- 🏗️ **[gormcnm](https://github.com/yyle88/gormcnm)** - Type-safe column foundation
- 🤖 **[gormcngen](https://github.com/yyle88/gormcngen)** - Smart code generation
- 🏢 **[gormrepo](https://github.com/yyle88/gormrepo)** - Enterprise repo pattern
- 🌍 **[gormmom](https://github.com/yyle88/gormmom)** - Native language programming