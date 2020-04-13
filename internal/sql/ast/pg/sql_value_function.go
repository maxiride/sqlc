package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SQLValueFunction struct {
	Location int
	Op       SQLValueFunctionOp
	Type     Oid
	Typmod   int32
	Xpr      ast.Node
}

func (n *SQLValueFunction) Pos() int {
	return n.Location
}
