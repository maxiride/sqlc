package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ConvertRowtypeExpr struct {
	Arg           ast.Node
	Convertformat CoercionForm
	Location      int
	Resulttype    Oid
	Xpr           ast.Node
}

func (n *ConvertRowtypeExpr) Pos() int {
	return n.Location
}
