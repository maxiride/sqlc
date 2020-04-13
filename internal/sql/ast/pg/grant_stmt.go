package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type GrantStmt struct {
	Behavior    DropBehavior
	GrantOption bool
	Grantees    *ast.List
	IsGrant     bool
	Objects     *ast.List
	Objtype     GrantObjectType
	Privileges  *ast.List
	Targtype    GrantTargetType
}

func (n *GrantStmt) Pos() int {
	return 0
}
