package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterUserMappingStmt struct {
	Options    *ast.List
	Servername *string
	User       *RoleSpec
}

func (n *AlterUserMappingStmt) Pos() int {
	return 0
}
