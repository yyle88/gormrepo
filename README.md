[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormrepo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormrepo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormrepo)](https://pkg.go.dev/github.com/yyle88/gormrepo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormrepo/main.svg)](https://coveralls.io/github/yyle88/gormrepo?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.24+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormrepo.svg)](https://github.com/yyle88/gormrepo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormrepo)](https://goreportcard.com/report/github.com/yyle88/gormrepo)

# 🚀 GORM Ecosystem - Enterprise-Grade Type-Safe Database Operations

**gormrepo** is the centerpiece of a complete GORM ecosystem, delivering **type-safe**, **enterprise-grade**, and **quite efficient** database operations to Go developers.

> 🌟 **Combining the best of Java MyBatis Plus + Python SQLAlchemy, designed with Go's next-generation ORM toolchain**

---

## Ecosystem

![GORM Type-Safe Ecosystem](https://github.com/yyle88/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

---

## 🔄 Tech Comparison

| Ecosystem             | Java MyBatis Plus  | Python SQLAlchemy | Go GORM Ecosystem   |
| --------------------- | ------------------ | ----------------- | ------------------- |
| **Type-Safe Columns** | `Example::getName` | `Example.name`    | `cls.Name.Eq()`     |
| **Code Generation**   | ✅ Plugin support  | ✅ Reflection     | ✅ AST precision    |
| **Repo Pattern**      | ✅ BaseMapper      | ✅ Session API    | ✅ GormRepo         |
| **Native Language**   | 🟡 Limited         | 🟡 Limited        | ✅ Complete support |

---

## 🚀 Quick Start

### Installation

```bash
go get github.com/yyle88/gormrepo
```

### Complete Usage Flow

#### 1. Define The Model (Supporting Native Fields)

```go
type Account struct {
    ID       uint   `gorm:"primaryKey"`
    Accountname string `gorm:"uniqueIndex" cnm:"accountname"`
    Nickname string `gorm:"index" cnm:"nickname"`
    Age      int    `cnm:"age"`
}
```

#### 2. Auto-Generate Column Structs (gormcngen)

```go
// Auto-generated
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

#### 3. Create Repo

```go
// Create repo with columns
repo := gormrepo.NewRepo(&Account{}, (&Account{}).Columns())

// Concise approach with gormrepo/gormclass
repo := gormrepo.NewRepo(gormclass.Use(&Account{}))
```

#### 4. How to Select

```go
// Classic GORM
err := db.Where("name = ?", "alice").First(&account).Error

// gormrepo - First
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Accountname.Eq("alice"))
})

// gormrepo - FirstE (returns ErrorOrNotExist)
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

// gormrepo - Where conditions: Eq, Ne, Gt, Lt, Gte, Lte
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

// gormrepo - Or conditions
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

#### Select Operations

| Method             | Parameters                                  | Returns                  | Description                          |
| ------------------ | ------------------------------------------- | ------------------------ | ------------------------------------ |
| `First`            | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `*MOD, error`            | Find first matching record           |
| `FirstE`           | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `*MOD, *ErrorOrNotExist` | Find first with not-exist indication |
| `Find`             | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `[]*MOD, error`          | Find each matching records           |
| `FindPage`         | `where, ordering, pagination`               | `[]*MOD, error`          | Paginated search                     |
| `FindPageAndCount` | `where, ordering, pagination`               | `[]*MOD, int64, error`   | Paginated search with record count   |
| `Count`            | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `int64, error`           | Count matching records               |
| `Exist`            | `where func(db *gorm.DB, cls CLS) *gorm.DB` | `bool, error`            | Check if records exist               |

#### 5. How to Create

```go
// Classic GORM
err := db.Create(&Account{Accountname: "bob", Nickname: "Bob", Age: 25}).Error

// gormrepo
err := repo.With(ctx, db).Create(&Account{Accountname: "bob", Nickname: "Bob", Age: 25})
```

#### Create Operations

| Method    | Parameters    | Returns | Description                 |
| --------- | ------------- | ------- | --------------------------- |
| `Create`  | `one *MOD`    | `error` | Create new record           |
| `Creates` | `ones []*MOD` | `error` | Batch create records        |
| `Save`    | `one *MOD`    | `error` | Insert/update record        |
| `Saves`   | `ones []*MOD` | `error` | Batch insert/update records |

#### 6. How to Update

```go
// Classic GORM
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

// gormrepo - UpdatesM (using ColumnValueMap)
err := repo.With(ctx, db).UpdatesM(
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *AccountColumns) *gormcnm.ColumnValueMap {
        return cls.Kw(cls.Age.Kv(26)).
                  Kw(cls.Nickname.Kv("NewNick"))
    },
)

// gormrepo - UpdatesO (by primary key)
account := &Account{ID: 1}
err := repo.With(ctx, db).UpdatesO(account, func(cls *AccountColumns) *gormcnm.ColumnValueMap {
    return cls.Kw(cls.Age.Kv(26))
})

// gormrepo - UpdatesC (combined conditions)
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

#### Update Operations

| Method     | Parameters                 | Returns | Description                                  |
| ---------- | -------------------------- | ------- | -------------------------------------------- |
| `Update`   | `where, valueFunc`         | `error` | Update single field                          |
| `Updates`  | `where, mapValues`         | `error` | Update multiple fields                       |
| `UpdatesM` | `where, newValues`         | `error` | Update with ColumnValueMap (M=Map)           |
| `UpdatesO` | `object, newValues`        | `error` | Update via primary key (O=Object)            |
| `UpdatesC` | `object, where, newValues` | `error` | Update with combined conditions (C=Combined) |

#### 7. How to Delete

```go
// Classic GORM
err := db.Where("id = ?", 1).Delete(&Account{}).Error

// gormrepo - Delete (by instance)
account := &Account{ID: 1}
err := repo.With(ctx, db).Delete(account)

// gormrepo - DeleteW (by where conditions)
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})

// gormrepo - DeleteM (instance + conditions)
account := &Account{ID: 1}
err := repo.With(ctx, db).DeleteM(account, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Lt(30))
})
```

#### Delete Operations

| Method    | Parameters                                            | Returns | Description                 |
| --------- | ----------------------------------------------------- | ------- | --------------------------- |
| `Delete`  | `one *MOD`                                            | `error` | Delete record via instance  |
| `DeleteW` | `where func(db *gorm.DB, cls CLS) *gorm.DB`           | `error` | Delete via conditions       |
| `DeleteM` | `one *MOD, where func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | Delete item with conditions |

#### 8. Custom Operations

When above methods do not fit, use `Invoke`:

```go
// Invoke - batch update
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18)).Update(cls.Nickname.Name(), "adult")
})

// Invoke - increment field
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.ID.Eq(1)).
        Update(cls.Age.Name(), gorm.Expr(cls.Age.Name()+" + ?", 1))
})

// Invoke - select specific columns
var names []string
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18)).Pluck(cls.Accountname.Name(), &names)
})

// Invoke - complex conditions
err := repo.With(ctx, db).Invoke(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Between(18, 60)).
        Where(cls.Nickname.IsNotNull()).
        Updates(map[string]interface{}{
            cls.Nickname.Name(): "active",
        })
})
```

#### Invoke Operations

| Method   | Parameters                                   | Returns | Description              |
| -------- | -------------------------------------------- | ------- | ------------------------ |
| `Invoke` | `clsRun func(db *gorm.DB, cls CLS) *gorm.DB` | `error` | Execute custom operation |

#### 9. Upsert with Clauses

Use `Clauses` / `Clause` to handle upsert (insert / update on conflict):

```go
// Clauses - pass clause.Expression
cls := account.Columns()
err := repo.With(ctx, db).Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: cls.Username.Name()}},
    DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
}).Create(&account)

// Clause - type-safe with column definitions
err := repo.With(ctx, db).Clause(func(cls *AccountColumns) clause.Expression {
    return clause.OnConflict{
        Columns:   []clause.Column{{Name: cls.Username.Name()}},
        DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
    }
}).Create(&account)

// Batch upsert with Clause + Creates
accounts := []*Account{
    {Username: "user1", Nickname: "nick1"},
    {Username: "user2", Nickname: "nick2"},
}
err := repo.With(ctx, db).Clause(func(cls *AccountColumns) clause.Expression {
    return clause.OnConflict{
        Columns:   []clause.Column{{Name: cls.Username.Name()}},
        DoUpdates: clause.AssignmentColumns([]string{cls.Nickname.Name()}),
    }
}).Creates(accounts)
```

#### Clause Operations

| Method    | Parameters                        | Returns     | Description                          |
| --------- | --------------------------------- | ----------- | ------------------------------------ |
| `Clauses` | `clauses ...clause.Expression`    | `*GormRepo` | Add clauses and return new repo      |
| `Clause`  | `func(cls CLS) clause.Expression` | `*GormRepo` | Build clause with column definitions |

---

## 📝 Complete Examples

Check [examples](internal/examples) DIR with complete integration examples.

---

## Related Projects

Explore the complete GORM ecosystem with these integrated packages:

### Core Ecosystem

- **[gormcnm](https://github.com/yyle88/gormcnm)** - GORM foundation providing type-safe column queries operations
- **[gormcngen](https://github.com/yyle88/gormcngen)** - Code generation engine using AST, enables type-safe GORM operations
- **[gormrepo](https://github.com/yyle88/gormrepo)** - Repo pattern implementation with GORM best practices (this project)
- **[gormmom](https://github.com/yyle88/gormmom)** - Native language GORM tag generation engine with smart column naming
- **[gormzhcn](https://github.com/go-zwbc/gormzhcn)** - Complete Chinese programming interface with GORM

Each package targets different aspects of GORM development, from localization to type-safe operations and code generation.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
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
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![Stargazers](https://starchart.cc/yyle88/gormrepo.svg?variant=adaptive)](https://starchart.cc/yyle88/gormrepo)
