package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type LockStmt struct {
	Mode      int
	Nowait    bool
	Relations *ast.List
}

func (n *LockStmt) Pos() int {
	return 0
}
