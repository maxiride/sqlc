package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterOwnerStmt struct {
	Newowner   *RoleSpec
	Object     ast.Node
	ObjectType ObjectType
	Relation   *RangeVar
}

func (n *AlterOwnerStmt) Pos() int {
	return 0
}
