package gormrepo

// Repo is a generic type alias to BaseRepo
// Go 1.24+ has native generic type alias support
//
// Repo 是 BaseRepo 的泛型类型别名
// Go 1.24+ 原生支持泛型类型别名
type Repo[MOD any, CLS any] = BaseRepo[MOD, CLS]

// NewRepo creates a new Repo instance with MOD and CLS definitions
// The param is used to deduce the type, its value is not used
//
// NewRepo 使用 MOD 和 CLS 定义创建新的 Repo 实例
// 参数用于类型推断，其值不使用
func NewRepo[MOD any, CLS any](one *MOD, cls CLS) *Repo[MOD, CLS] {
	return NewBaseRepo[MOD, CLS](one, cls)
}
