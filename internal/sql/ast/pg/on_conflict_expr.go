package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type OnConflictExpr struct {
	Action          OnConflictAction
	ArbiterElems    *ast.List
	ArbiterWhere    ast.Node
	Constraint      Oid
	ExclRelIndex    int
	ExclRelTlist    *ast.List
	OnConflictSet   *ast.List
	OnConflictWhere ast.Node
}

func (n *OnConflictExpr) Pos() int {
	return 0
}
