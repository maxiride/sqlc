package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type TypeName struct {
	ArrayBounds *ast.List
	Location    int
	Names       *ast.List
	PctType     bool
	Setof       bool
	TypeOid     Oid
	Typemod     int32
	Typmods     *ast.List
}

func (n *TypeName) Pos() int {
	return n.Location
}
