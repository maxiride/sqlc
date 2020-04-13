package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DeclareCursorStmt struct {
	Options    int
	Portalname *string
	Query      ast.Node
}

func (n *DeclareCursorStmt) Pos() int {
	return 0
}
