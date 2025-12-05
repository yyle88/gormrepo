//go:build !go1.23

package gormrepo

// Repo is a compatibility version using struct embedding (Go < 1.23)
// This version is used when Go version is below 1.23
// Go 1.23+ uses type alias in repo.go instead
//
// Repo 是使用结构体嵌入的兼容版本（Go < 1.23）
// 当 Go 版本低于 1.23 时使用此版本
// Go 1.23+ 使用 repo.go 中的类型别名版本
type Repo[MOD any, CLS any] struct {
	*BaseRepo[MOD, CLS]
}

// NewRepo creates a new Repo instance with column definitions
// The model parameter is used to infer the type, actual value is not used
//
// NewRepo 使用列定义创建新的 Repo 实例
// model 参数用于类型推断，实际值不使用
func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return &Repo[MOD, CLS]{
		BaseRepo: NewBaseRepo[MOD, CLS]((*MOD)(nil), cls),
	}
}
