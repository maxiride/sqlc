package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type A_Const struct {
	Location int
	Val      ast.Node
}

func (n *A_Const) Pos() int {
	return n.Location
}
