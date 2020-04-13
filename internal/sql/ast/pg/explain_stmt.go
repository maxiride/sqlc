package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ExplainStmt struct {
	Options *ast.List
	Query   ast.Node
}

func (n *ExplainStmt) Pos() int {
	return 0
}
