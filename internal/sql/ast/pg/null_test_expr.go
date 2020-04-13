package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type NullTest struct {
	Arg          ast.Node
	Argisrow     bool
	Location     int
	Nulltesttype NullTestType
	Xpr          ast.Node
}

func (n *NullTest) Pos() int {
	return n.Location
}
