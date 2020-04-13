package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SortBy struct {
	Location    int
	Node        ast.Node
	SortbyDir   SortByDir
	SortbyNulls SortByNulls
	UseOp       *ast.List
}

func (n *SortBy) Pos() int {
	return n.Location
}
