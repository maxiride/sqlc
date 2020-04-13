package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CoerceToDomain struct {
	Arg            ast.Node
	Coercionformat CoercionForm
	Location       int
	Resultcollid   Oid
	Resulttype     Oid
	Resulttypmod   int32
	Xpr            ast.Node
}

func (n *CoerceToDomain) Pos() int {
	return n.Location
}
