package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ArrayCoerceExpr struct {
	Arg          ast.Node
	Coerceformat CoercionForm
	Elemfuncid   Oid
	IsExplicit   bool
	Location     int
	Resultcollid Oid
	Resulttype   Oid
	Resulttypmod int32
	Xpr          ast.Node
}

func (n *ArrayCoerceExpr) Pos() int {
	return n.Location
}
