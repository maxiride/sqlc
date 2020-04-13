package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CommentStmt struct {
	Comment *string
	Object  ast.Node
	Objtype ObjectType
}

func (n *CommentStmt) Pos() int {
	return 0
}
