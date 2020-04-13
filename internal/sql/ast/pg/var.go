package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type Var struct {
	Location    int
	Varattno    AttrNumber
	Varcollid   Oid
	Varlevelsup Index
	Varno       Index
	Varnoold    Index
	Varoattno   AttrNumber
	Vartype     Oid
	Vartypmod   int32
	Xpr         ast.Node
}

func (n *Var) Pos() int {
	return n.Location
}
