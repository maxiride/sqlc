package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type NamedArgExpr struct {
	Arg       ast.Node
	Argnumber int
	Location  int
	Name      *string
	Xpr       ast.Node
}

func (n *NamedArgExpr) Pos() int {
	return n.Location
}
