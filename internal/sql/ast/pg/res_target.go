package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ResTarget struct {
	Indirection *ast.List
	Location    int
	Name        *string
	Val         ast.Node
}

func (n *ResTarget) Pos() int {
	return n.Location
}
