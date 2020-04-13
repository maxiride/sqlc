package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type VariableSetStmt struct {
	Args    *ast.List
	IsLocal bool
	Kind    VariableSetKind
	Name    *string
}

func (n *VariableSetStmt) Pos() int {
	return 0
}
