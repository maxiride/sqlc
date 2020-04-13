package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RangeTableFuncCol struct {
	Coldefexpr    ast.Node
	Colexpr       ast.Node
	Colname       *string
	ForOrdinality bool
	IsNotNull     bool
	Location      int
	TypeName      *TypeName
}

func (n *RangeTableFuncCol) Pos() int {
	return n.Location
}
