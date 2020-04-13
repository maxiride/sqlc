package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterTableCmd struct {
	Behavior  DropBehavior
	Def       ast.Node
	MissingOk bool
	Name      *string
	Newowner  *RoleSpec
	Subtype   AlterTableType
}

func (n *AlterTableCmd) Pos() int {
	return 0
}
