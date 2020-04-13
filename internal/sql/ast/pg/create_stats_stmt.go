package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateStatsStmt struct {
	Defnames    *ast.List
	Exprs       *ast.List
	IfNotExists bool
	Relations   *ast.List
	StatTypes   *ast.List
}

func (n *CreateStatsStmt) Pos() int {
	return 0
}
