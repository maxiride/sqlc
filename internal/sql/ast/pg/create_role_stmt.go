package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateRoleStmt struct {
	Options  *ast.List
	Role     *string
	StmtType RoleStmtType
}

func (n *CreateRoleStmt) Pos() int {
	return 0
}
