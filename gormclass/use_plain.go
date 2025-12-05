package gormclass

import "github.com/yyle88/gormcnm"

// UsePlain returns the struct (`mod`) and its associated columns (`cls`) without table name prefix, suitable in queries/operations that need both.
// UsePlain 返回模型（`mod`）和不带表名前缀的关联列（`cls`），适用于需要同时获取模型和列数据的查询或操作。
func UsePlain[MOD TableCols[CLS], CLS any](one MOD) (MOD, CLS) {
	return one, one.TableColumns(gormcnm.NewPlainDecoration())
}

// UmcPlain returns the struct (mod) and the associated columns (cls) without table name prefix, functioning the same as the UsePlain function.
// UmcPlain 返回模型（mod）和不带表名前缀的关联列（cls），功能与 UsePlain 函数相同。
func UmcPlain[MOD TableCols[CLS], CLS any](one MOD) (MOD, CLS) {
	return one, one.TableColumns(gormcnm.NewPlainDecoration())
}
