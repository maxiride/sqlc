package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type TableSampleClause struct {
	Args       *ast.List
	Repeatable ast.Node
	Tsmhandler Oid
}

func (n *TableSampleClause) Pos() int {
	return 0
}
