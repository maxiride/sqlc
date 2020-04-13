package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type UpdateStmt struct {
	FromClause    *ast.List
	Relation      *RangeVar
	ReturningList *ast.List
	TargetList    *ast.List
	WhereClause   ast.Node
	WithClause    *WithClause
}

func (n *UpdateStmt) Pos() int {
	return 0
}
