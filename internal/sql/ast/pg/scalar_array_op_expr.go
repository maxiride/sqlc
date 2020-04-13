package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ScalarArrayOpExpr struct {
	Args        *ast.List
	Inputcollid Oid
	Location    int
	Opfuncid    Oid
	Opno        Oid
	UseOr       bool
	Xpr         ast.Node
}

func (n *ScalarArrayOpExpr) Pos() int {
	return n.Location
}
