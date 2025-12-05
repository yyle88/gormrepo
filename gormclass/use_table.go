package gormclass

import "github.com/yyle88/gormcnm"

// UseTable returns the struct (`mod`) and its associated columns (`cls`) with table name prefix, suitable in queries/operations that need both.
// UseTable 返回模型（`mod`）和带有表名前缀的关联列（`cls`），适用于需要同时获取模型和列数据的查询或操作。
func UseTable[MOD TableClass[CLS], CLS any](one MOD) (MOD, string, CLS) {
	return one, one.TableName(), one.TableColumns(gormcnm.NewTableDecoration(one.TableName()))
}

// UmcTable returns the struct (mod) and the associated columns (cls) with table name prefix, functioning the same as the UseTable function.
// UmcTable 返回模型（mod）和带有表名前缀的关联列（cls），功能与 UseTable 函数相同。
func UmcTable[MOD TableClass[CLS], CLS any](one MOD) (MOD, string, CLS) {
	return one, one.TableName(), one.TableColumns(gormcnm.NewTableDecoration(one.TableName()))
}
