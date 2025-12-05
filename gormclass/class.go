package gormclass

import "github.com/yyle88/gormcnm"

// GormTable is used with models that implement the TableName method to return the table name.
// GormTable 用于实现 TableName 方法以返回表名的模型。
type GormTable interface {
	TableName() string
}

// ModelCols is used with models that implement the Columns method to return associated columns (cls).
// ModelCols 用于实现 Columns 方法以返回关联列（cls）的模型。
type ModelCols[CLS any] interface {
	Columns() CLS
}

// ModelClass is used with models that implement the Columns method to return associated columns (cls) and the TableName method to return the table name.
// ModelClass 用于实现 Columns 方法以返回关联列（cls）和 TableName 方法以返回表名的模型。
type ModelClass[CLS any] interface {
	TableName() string
	Columns() CLS
}

// TableCols is used with models that implement the TableColumns method to return associated columns (cls) with decoration.
// TableCols 用于实现 TableColumns 方法以返回带装饰器的关联列（cls）的模型。
type TableCols[CLS any] interface {
	TableColumns(gormcnm.ColumnNameDecoration) CLS
}

// TableClass is used with models that implement the TableColumns method to return associated columns (cls) and the TableName method to return the table name.
// TableClass 用于实现 TableColumns 方法以返回关联列（cls）和 TableName 方法以返回表名的模型。
type TableClass[CLS any] interface {
	TableName() string
	TableColumns(gormcnm.ColumnNameDecoration) CLS
}
