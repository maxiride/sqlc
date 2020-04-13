package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterDomainStmt struct {
	Behavior  DropBehavior
	Def       ast.Node
	MissingOk bool
	Name      *string
	Subtype   byte
	TypeName  *ast.List
}

func (n *AlterDomainStmt) Pos() int {
	return 0
}
