package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ArrayExpr struct {
	ArrayCollid   Oid
	ArrayTypeid   Oid
	ElementTypeid Oid
	Elements      *ast.List
	Location      int
	Multidims     bool
	Xpr           ast.Node
}

func (n *ArrayExpr) Pos() int {
	return n.Location
}
