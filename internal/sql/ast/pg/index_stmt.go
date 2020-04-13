package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type IndexStmt struct {
	AccessMethod   *string
	Concurrent     bool
	Deferrable     bool
	ExcludeOpNames *ast.List
	Idxcomment     *string
	Idxname        *string
	IfNotExists    bool
	IndexOid       Oid
	IndexParams    *ast.List
	Initdeferred   bool
	Isconstraint   bool
	OldNode        Oid
	Options        *ast.List
	Primary        bool
	Relation       *RangeVar
	TableSpace     *string
	Transformed    bool
	Unique         bool
	WhereClause    ast.Node
}

func (n *IndexStmt) Pos() int {
	return 0
}
