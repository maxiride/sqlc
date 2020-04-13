package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DeleteStmt struct {
	Relation      *RangeVar
	ReturningList *ast.List
	UsingClause   *ast.List
	WhereClause   ast.Node
	WithClause    *WithClause
}

func (n *DeleteStmt) Pos() int {
	return 0
}
