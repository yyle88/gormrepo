//go:build go1.23

package gormrepo

// Repo is a type alias to BaseRepo (Go 1.23+ feature)
// Since BaseRepo has more explicit meaning, we renamed it to BaseRepo
// But Repo is shorter and more convenient, so we use type alias
//
// Repo 是 BaseRepo 的类型别名（Go 1.23+ 特性）
// 由于 BaseRepo 具有更明确的语义，因此重命名为 BaseRepo
// 但 Repo 更短更便于使用，因此使用类型别名
type Repo[MOD any, CLS any] = BaseRepo[MOD, CLS]

// NewRepo creates a new Repo instance with column definitions
// The model parameter is used to infer the type, actual value is not used
//
// NewRepo 使用列定义创建新的 Repo 实例
// model 参数用于类型推断，实际值不使用
func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return NewBaseRepo[MOD, CLS]((*MOD)(nil), cls)
}
