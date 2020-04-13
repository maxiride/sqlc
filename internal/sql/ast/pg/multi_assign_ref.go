package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type MultiAssignRef struct {
	Colno    int
	Ncolumns int
	Source   ast.Node
}

func (n *MultiAssignRef) Pos() int {
	return 0
}
