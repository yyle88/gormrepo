//go:build !go1.24

package gormrepo

// Repo is a compat version using struct embedding (Go < 1.24)
// Generic type alias needs Go 1.24+ (Go 1.23 needs GOEXPERIMENT=aliastypeparams)
// This struct embedding approach provides same API with small overhead
// Go 1.24+ uses type alias in repo.go instead
//
// Repo 是使用结构体嵌入的兼容版本（Go < 1.24）
// 泛型类型别名需要 Go 1.24+（Go 1.23 需要 GOEXPERIMENT=aliastypeparams）
// 结构体嵌入方式提供相同 API，仅有轻微开销
// Go 1.24+ 使用 repo.go 中的类型别名版本
type Repo[MOD any, CLS any] struct {
	*BaseRepo[MOD, CLS]
}

// NewRepo creates a new Repo instance with column definitions
// The model param is used to deduce the type, actual value is not used
//
// NewRepo 使用列定义创建新的 Repo 实例
// model 参数用于类型推断，实际值不使用
func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return &Repo[MOD, CLS]{
		BaseRepo: NewBaseRepo[MOD, CLS]((*MOD)(nil), cls),
	}
}
