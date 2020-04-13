package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterRoleStmt struct {
	Action  int
	Options *ast.List
	Role    *RoleSpec
}

func (n *AlterRoleStmt) Pos() int {
	return 0
}
