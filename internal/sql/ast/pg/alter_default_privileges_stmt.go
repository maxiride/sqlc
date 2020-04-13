package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterDefaultPrivilegesStmt struct {
	Action  *GrantStmt
	Options *ast.List
}

func (n *AlterDefaultPrivilegesStmt) Pos() int {
	return 0
}
