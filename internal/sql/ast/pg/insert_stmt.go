package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type InsertStmt struct {
	Cols             *ast.List
	OnConflictClause *OnConflictClause
	Override         OverridingKind
	Relation         *RangeVar
	ReturningList    *ast.List
	SelectStmt       ast.Node
	WithClause       *WithClause
}

func (n *InsertStmt) Pos() int {
	return 0
}
