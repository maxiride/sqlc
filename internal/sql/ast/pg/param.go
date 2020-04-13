package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type Param struct {
	Location    int
	Paramcollid Oid
	Paramid     int
	Paramkind   ParamKind
	Paramtype   Oid
	Paramtypmod int32
	Xpr         ast.Node
}

func (n *Param) Pos() int {
	return n.Location
}
