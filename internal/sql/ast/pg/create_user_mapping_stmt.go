package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateUserMappingStmt struct {
	IfNotExists bool
	Options     *ast.List
	Servername  *string
	User        *RoleSpec
}

func (n *CreateUserMappingStmt) Pos() int {
	return 0
}
