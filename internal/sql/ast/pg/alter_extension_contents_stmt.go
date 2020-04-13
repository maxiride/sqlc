package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterExtensionContentsStmt struct {
	Action  int
	Extname *string
	Object  ast.Node
	Objtype ObjectType
}

func (n *AlterExtensionContentsStmt) Pos() int {
	return 0
}
