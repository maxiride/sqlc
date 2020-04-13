package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RowExpr struct {
	Args      *ast.List
	Colnames  *ast.List
	Location  int
	RowFormat CoercionForm
	RowTypeid Oid
	Xpr       ast.Node
}

func (n *RowExpr) Pos() int {
	return n.Location
}
