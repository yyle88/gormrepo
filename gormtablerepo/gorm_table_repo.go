package gormtablerepo

type TableRepo[MOD any, CLS any] struct {
	tableName string
	tbColumns CLS
}

func NewTableRepo[MOD any, CLS any](_ *MOD, tableName string, tbColumns CLS) *TableRepo[MOD, CLS] {
	return &TableRepo[MOD, CLS]{
		tableName: tableName,
		tbColumns: tbColumns,
	}
}

func (repo *TableRepo[MOD, CLS]) GetTableName() string {
	return repo.tableName
}

func (repo *TableRepo[MOD, CLS]) TableColumns() CLS {
	return repo.tbColumns
}

func (repo *TableRepo[MOD, CLS]) BuildColumns(run func(cls CLS) []string) []string {
	return run(repo.tbColumns)
}
