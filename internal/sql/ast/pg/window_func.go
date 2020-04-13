package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type WindowFunc struct {
	Aggfilter   ast.Node
	Args        *ast.List
	Inputcollid Oid
	Location    int
	Winagg      bool
	Wincollid   Oid
	Winfnoid    Oid
	Winref      Index
	Winstar     bool
	Wintype     Oid
	Xpr         ast.Node
}

func (n *WindowFunc) Pos() int {
	return n.Location
}
