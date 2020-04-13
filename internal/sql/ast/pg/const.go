package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type Const struct {
	Constbyval  bool
	Constcollid Oid
	Constisnull bool
	Constlen    int
	Consttype   Oid
	Consttypmod int32
	Constvalue  Datum
	Location    int
	Xpr         ast.Node
}

func (n *Const) Pos() int {
	return n.Location
}
