package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type OpExpr struct {
	Args         *ast.List
	Inputcollid  Oid
	Location     int
	Opcollid     Oid
	Opfuncid     Oid
	Opno         Oid
	Opresulttype Oid
	Opretset     bool
	Xpr          ast.Node
}

func (n *OpExpr) Pos() int {
	return n.Location
}
