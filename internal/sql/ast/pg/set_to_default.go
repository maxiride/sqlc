package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SetToDefault struct {
	Collation Oid
	Location  int
	TypeId    Oid
	TypeMod   int32
	Xpr       ast.Node
}

func (n *SetToDefault) Pos() int {
	return n.Location
}
