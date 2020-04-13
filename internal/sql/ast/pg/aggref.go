package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type Aggref struct {
	Aggargtypes   *ast.List
	Aggcollid     Oid
	Aggdirectargs *ast.List
	Aggdistinct   *ast.List
	Aggfilter     ast.Node
	Aggfnoid      Oid
	Aggkind       byte
	Agglevelsup   Index
	Aggorder      *ast.List
	Aggsplit      AggSplit
	Aggstar       bool
	Aggtranstype  Oid
	Aggtype       Oid
	Aggvariadic   bool
	Args          *ast.List
	Inputcollid   Oid
	Location      int
	Xpr           ast.Node
}

func (n *Aggref) Pos() int {
	return n.Location
}
