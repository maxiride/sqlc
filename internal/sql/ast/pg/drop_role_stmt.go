package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DropRoleStmt struct {
	MissingOk bool
	Roles     *ast.List
}

func (n *DropRoleStmt) Pos() int {
	return 0
}
