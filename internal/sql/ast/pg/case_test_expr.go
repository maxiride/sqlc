package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CaseTestExpr struct {
	Collation Oid
	TypeId    Oid
	TypeMod   int32
	Xpr       ast.Node
}

func (n *CaseTestExpr) Pos() int {
	return 0
}
