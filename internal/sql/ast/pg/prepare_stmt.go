package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type PrepareStmt struct {
	Argtypes *ast.List
	Name     *string
	Query    ast.Node
}

func (n *PrepareStmt) Pos() int {
	return 0
}
