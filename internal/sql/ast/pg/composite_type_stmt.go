package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CompositeTypeStmt struct {
	Coldeflist *ast.List
	Typevar    *RangeVar
}

func (n *CompositeTypeStmt) Pos() int {
	return 0
}
