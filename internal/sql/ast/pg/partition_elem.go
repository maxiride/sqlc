package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type PartitionElem struct {
	Collation *ast.List
	Expr      ast.Node
	Location  int
	Name      *string
	Opclass   *ast.List
}

func (n *PartitionElem) Pos() int {
	return n.Location
}
