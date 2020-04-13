package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ReassignOwnedStmt struct {
	Newrole *RoleSpec
	Roles   *ast.List
}

func (n *ReassignOwnedStmt) Pos() int {
	return 0
}
