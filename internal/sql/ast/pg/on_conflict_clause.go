package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type OnConflictClause struct {
	Action      OnConflictAction
	Infer       *InferClause
	Location    int
	TargetList  *ast.List
	WhereClause ast.Node
}

func (n *OnConflictClause) Pos() int {
	return n.Location
}
