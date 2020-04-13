package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DropOwnedStmt struct {
	Behavior DropBehavior
	Roles    *ast.List
}

func (n *DropOwnedStmt) Pos() int {
	return 0
}
