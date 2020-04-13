package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SecLabelStmt struct {
	Label    *string
	Object   ast.Node
	Objtype  ObjectType
	Provider *string
}

func (n *SecLabelStmt) Pos() int {
	return 0
}
