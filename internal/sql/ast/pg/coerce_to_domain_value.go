package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CoerceToDomainValue struct {
	Collation Oid
	Location  int
	TypeId    Oid
	TypeMod   int32
	Xpr       ast.Node
}

func (n *CoerceToDomainValue) Pos() int {
	return n.Location
}
