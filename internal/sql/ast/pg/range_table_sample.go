package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RangeTableSample struct {
	Args       *ast.List
	Location   int
	Method     *ast.List
	Relation   ast.Node
	Repeatable ast.Node
}

func (n *RangeTableSample) Pos() int {
	return n.Location
}
