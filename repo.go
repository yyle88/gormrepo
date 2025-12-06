//go:build go1.24

package gormrepo

// Repo is a generic type alias to BaseRepo
// Generic type alias was introduced in Go 1.23 but needs GOEXPERIMENT=aliastypeparams
// Go 1.24+ has native support without experiment flag
// Using go1.24 build constraint to avoid: "generic type alias requires GOEXPERIMENT=aliastypeparams"
//
// Repo 是 BaseRepo 的泛型类型别名
// 泛型类型别名在 Go 1.23 引入，但需要设置 GOEXPERIMENT=aliastypeparams
// Go 1.24+ 原生支持泛型类型别名，无需实验标志
// 使用 go1.24 构建约束以避免报错："generic type alias requires GOEXPERIMENT=aliastypeparams"
type Repo[MOD any, CLS any] = BaseRepo[MOD, CLS]

// NewRepo creates a new Repo instance with column definitions
// The model param is used to deduce the type, actual value is not used
//
// NewRepo 使用列定义创建新的 Repo 实例
// model 参数用于类型推断，实际值不使用
func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return NewBaseRepo[MOD, CLS]((*MOD)(nil), cls)
}
