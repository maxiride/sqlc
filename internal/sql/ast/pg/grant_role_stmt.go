package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type GrantRoleStmt struct {
	AdminOpt     bool
	Behavior     DropBehavior
	GrantedRoles *ast.List
	GranteeRoles *ast.List
	Grantor      *RoleSpec
	IsGrant      bool
}

func (n *GrantRoleStmt) Pos() int {
	return 0
}
