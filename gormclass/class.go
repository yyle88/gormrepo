package gormclass

import "github.com/yyle88/gormcnm"

// GormClass is used for models that implement the Columns method to return associated columns (cls).
// GormClass 用于实现 Columns 方法以返回关联列（cls）的模型。
type GormClass[CLS any] interface {
	Columns() CLS
}

// GormTable is used for models that implement the TableName method to return the table name.
// GormTable 用于实现 TableName 方法以返回表名的模型。
type GormTable interface {
	TableName() string
}

// ClassType is used for models that implement the Columns method to return associated columns (cls) and the TableName method to return the table name.
// ClassType 用于实现 Columns 方法以返回关联列（cls）和 TableName 方法以返回表名的模型。
type ClassType[CLS any] interface {
	TableName() string
	Columns() CLS
}

// TableClass is used for models that implement the TableColumns method to return associated columns (cls) and the TableName method to return the table name.
// TableClass 用于实现 TableColumns 方法以返回关联列（cls）和 TableName 方法以返回表名的模型。
type TableClass[CLS any] interface {
	TableName() string
	TableColumns(gormcnm.ColumnNameDecoration) CLS
}

// TableCols is used for models that implement the TableColumns method to return associated columns (cls) with decoration.
// TableCols 用于实现 TableColumns 方法以返回带装饰器的关联列（cls）的模型。
type TableCols[CLS any] interface {
	TableColumns(gormcnm.ColumnNameDecoration) CLS
}
