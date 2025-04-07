package gormclass

// Use returns the model (`mod`) and its associated columns (`cls`), suitable for queries or operations that need both.
// Use 返回模型（`mod`）、关联的列（`cls`），适用于需要同时获取模型和列数据的查询或操作。
func Use[MOD ModelCols[CLS], CLS any](one MOD) (MOD, CLS) {
	return one, one.Columns()
}

// Umc returns the model (mod) and the associated columns (cls), functioning the same as the Use function.
// Umc 返回模型（mod）和关联的列（cls），功能与 Use 函数相同。
func Umc[MOD ModelCols[CLS], CLS any](one MOD) (MOD, CLS) {
	return one, one.Columns()
}

// Cls returns the column information (`cls`), useful when only column data is needed.
// Cls 返回列信息（`cls`），适用于仅需要列数据的场景。
func Cls[MOD ModelCols[CLS], CLS any](one MOD) CLS {
	return one.Columns()
}

// Usc returns a slice of models (`MOD`) and the associated columns (`cls`), suitable for queries returning multiple models (e.g., `Find` queries).
// Usc 返回多个模型（`MOD`）、关联的列（`cls`），适用于返回多个模型的查询（如 `Find` 查询）。
func Usc[MOD ModelCols[CLS], CLS any](one MOD) ([]MOD, CLS) {
	return []MOD{}, one.Columns()
}

// Msc returns the model (`mod`), the model slice (`[]MOD`), and the associated columns (`cls`), useful for queries requiring both model and column data.
// Msc 返回模型（`mod`）、模型切片（`[]MOD`）、关联的列（`cls`），适用于需要模型和列数据的查询。
func Msc[MOD ModelCols[CLS], CLS any](one MOD) (MOD, []MOD, CLS) {
	return one, []MOD{}, one.Columns()
}

// Nsc returns the model (`mod`), a slice of models (`[]MOD`) with a specified initial capacity, and the associated columns (`cls`),
// Nsc 返回模型（`mod`）、具有指定初始容量的模型切片（`[]MOD`）和关联的列（`cls`）。
func Nsc[MOD ModelCols[CLS], CLS any](one MOD, cap int) (MOD, []MOD, CLS) {
	return one, make([]MOD, 0, cap), one.Columns()
}

// One returns the model (mod), ensuring type safety by checking whether the argument is a pointer type at compile-time.
// One 返回模型（mod），通过编译时检查确保类型安全。
func One[MOD ModelCols[CLS], CLS any](one MOD) MOD {
	return one // 按照 gormcngen 的默认规则，类型只会生成 func (*X) Columns() XColumns {} 这样的成员函数
}

// Ums returns a slice of models (MOD), useful for queries that expect a slice of models (e.g., Find queries).
// Ums 返回模型（mod）切片，适用于需要模型切片的查询（例如 Find 查询）。
func Ums[MOD ModelCols[CLS], CLS any](MOD) []MOD {
	return []MOD{}
}

// Uss returns an empty slice of models (MOD), typically used for initialization or preparing for future object population without needing the columns (CLS).
// Uss 返回一个空的模型（mod）切片，通常用于初始化或为未来填充对象做准备，无需关联列（cls）。
func Uss[MOD ModelCols[CLS], CLS any]() []MOD {
	return []MOD{}
}

// Usn returns a slice of models (MOD) with a specified initial capacity, optimizing memory allocation based on the expected number of objects (MOD).
// Usn 返回一个具有指定初始容量的模型（mod）切片，优化内存分配以适应预期的对象数量（MOD）。
func Usn[MOD ModelCols[CLS], CLS any](cap int) []MOD {
	return make([]MOD, 0, cap)
}
