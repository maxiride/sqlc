package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type TruncateStmt struct {
	Behavior    DropBehavior
	Relations   *ast.List
	RestartSeqs bool
}

func (n *TruncateStmt) Pos() int {
	return 0
}
