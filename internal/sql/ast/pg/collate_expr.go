package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CollateExpr struct {
	Arg      ast.Node
	CollOid  Oid
	Location int
	Xpr      ast.Node
}

func (n *CollateExpr) Pos() int {
	return n.Location
}
