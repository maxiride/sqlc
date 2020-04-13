package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type GroupingFunc struct {
	Agglevelsup Index
	Args        *ast.List
	Cols        *ast.List
	Location    int
	Refs        *ast.List
	Xpr         ast.Node
}

func (n *GroupingFunc) Pos() int {
	return n.Location
}
