package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type MinMaxExpr struct {
	Args         *ast.List
	Inputcollid  Oid
	Location     int
	Minmaxcollid Oid
	Minmaxtype   Oid
	Op           MinMaxOp
	Xpr          ast.Node
}

func (n *MinMaxExpr) Pos() int {
	return n.Location
}
