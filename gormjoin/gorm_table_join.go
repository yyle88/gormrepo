package gormjoin

import (
	"strings"

	"github.com/yyle88/gormrepo/gormtablerepo"
	"gorm.io/gorm/clause"
)

type TableJoin[
	MOD1 any, CLS1 any,
	MOD2 any, CLS2 any,
] struct {
	repo1     *gormtablerepo.TableRepo[MOD1, CLS1]
	whichJoin clause.JoinType
	repo2     *gormtablerepo.TableRepo[MOD2, CLS2]
}

func JOIN[MOD1 any, CLS1 any, MOD2 any, CLS2 any](
	repo1 *gormtablerepo.TableRepo[MOD1, CLS1],
	whichJoin clause.JoinType,
	repo2 *gormtablerepo.TableRepo[MOD2, CLS2],
) *TableJoin[MOD1, CLS1, MOD2, CLS2] {
	return &TableJoin[MOD1, CLS1, MOD2, CLS2]{
		repo1:     repo1,
		whichJoin: whichJoin,
		repo2:     repo2,
	}
}

func LEFTJOIN[MOD1 any, CLS1 any, MOD2 any, CLS2 any](repo1 *gormtablerepo.TableRepo[MOD1, CLS1], repo2 *gormtablerepo.TableRepo[MOD2, CLS2]) *TableJoin[MOD1, CLS1, MOD2, CLS2] {
	return JOIN(repo1, clause.LeftJoin, repo2)
}

func RIGHTJOIN[MOD1 any, CLS1 any, MOD2 any, CLS2 any](repo1 *gormtablerepo.TableRepo[MOD1, CLS1], repo2 *gormtablerepo.TableRepo[MOD2, CLS2]) *TableJoin[MOD1, CLS1, MOD2, CLS2] {
	return JOIN(repo1, clause.RightJoin, repo2)
}

func INNERJOIN[MOD1 any, CLS1 any, MOD2 any, CLS2 any](repo1 *gormtablerepo.TableRepo[MOD1, CLS1], repo2 *gormtablerepo.TableRepo[MOD2, CLS2]) *TableJoin[MOD1, CLS1, MOD2, CLS2] {
	return JOIN(repo1, clause.InnerJoin, repo2)
}

func CROSSJOIN[MOD1 any, CLS1 any, MOD2 any, CLS2 any](repo1 *gormtablerepo.TableRepo[MOD1, CLS1], repo2 *gormtablerepo.TableRepo[MOD2, CLS2]) *TableJoin[MOD1, CLS1, MOD2, CLS2] {
	return JOIN(repo1, clause.CrossJoin, repo2)
}

func (op *TableJoin[MOD1, CLS1, MOD2, CLS2]) On(onFunc func(cls1 CLS1, cls2 CLS2) []string) string {
	return string(op.whichJoin) + " JOIN " + op.repo2.GetTableName() + " ON " + strings.Join(onFunc(op.repo1.TableColumns(), op.repo2.TableColumns()), " AND ")
}
