package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CoalesceExpr struct {
	Args           *ast.List
	Coalescecollid Oid
	Coalescetype   Oid
	Location       int
	Xpr            ast.Node
}

func (n *CoalesceExpr) Pos() int {
	return n.Location
}
