package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CoerceViaIO struct {
	Arg          ast.Node
	Coerceformat CoercionForm
	Location     int
	Resultcollid Oid
	Resulttype   Oid
	Xpr          ast.Node
}

func (n *CoerceViaIO) Pos() int {
	return n.Location
}
