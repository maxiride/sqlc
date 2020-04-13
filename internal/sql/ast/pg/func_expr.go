package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type FuncExpr struct {
	Args           *ast.List
	Funccollid     Oid
	Funcformat     CoercionForm
	Funcid         Oid
	Funcresulttype Oid
	Funcretset     bool
	Funcvariadic   bool
	Inputcollid    Oid
	Location       int
	Xpr            ast.Node
}

func (n *FuncExpr) Pos() int {
	return n.Location
}
