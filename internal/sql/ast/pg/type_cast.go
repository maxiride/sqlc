package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type TypeCast struct {
	Arg      ast.Node
	Location int
	TypeName *TypeName
}

func (n *TypeCast) Pos() int {
	return n.Location
}
