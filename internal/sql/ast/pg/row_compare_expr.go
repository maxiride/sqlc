package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RowCompareExpr struct {
	Inputcollids *ast.List
	Largs        *ast.List
	Opfamilies   *ast.List
	Opnos        *ast.List
	Rargs        *ast.List
	Rctype       RowCompareType
	Xpr          ast.Node
}

func (n *RowCompareExpr) Pos() int {
	return 0
}
