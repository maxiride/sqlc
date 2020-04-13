package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DropStmt struct {
	Behavior   DropBehavior
	Concurrent bool
	MissingOk  bool
	Objects    *ast.List
	RemoveType ObjectType
}

func (n *DropStmt) Pos() int {
	return 0
}
