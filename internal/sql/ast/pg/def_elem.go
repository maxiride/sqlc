package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DefElem struct {
	Arg          ast.Node
	Defaction    DefElemAction
	Defname      *string
	Defnamespace *string
	Location     int
}

func (n *DefElem) Pos() int {
	return n.Location
}
