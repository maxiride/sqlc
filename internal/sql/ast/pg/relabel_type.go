package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RelabelType struct {
	Arg           ast.Node
	Location      int
	Relabelformat CoercionForm
	Resultcollid  Oid
	Resulttype    Oid
	Resulttypmod  int32
	Xpr           ast.Node
}

func (n *RelabelType) Pos() int {
	return n.Location
}
