package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type BooleanTest struct {
	Arg          ast.Node
	Booltesttype BoolTestType
	Location     int
	Xpr          ast.Node
}

func (n *BooleanTest) Pos() int {
	return n.Location
}
