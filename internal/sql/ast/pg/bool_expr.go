package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type BoolExpr struct {
	Args     *ast.List
	Boolop   BoolExprType
	Location int
	Xpr      ast.Node
}

func (n *BoolExpr) Pos() int {
	return n.Location
}
