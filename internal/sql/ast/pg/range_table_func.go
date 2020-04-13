package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RangeTableFunc struct {
	Alias      *Alias
	Columns    *ast.List
	Docexpr    ast.Node
	Lateral    bool
	Location   int
	Namespaces *ast.List
	Rowexpr    ast.Node
}

func (n *RangeTableFunc) Pos() int {
	return n.Location
}
