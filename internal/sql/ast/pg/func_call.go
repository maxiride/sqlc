package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type FuncCall struct {
	AggDistinct    bool
	AggFilter      ast.Node
	AggOrder       *ast.List
	AggStar        bool
	AggWithinGroup bool
	Args           *ast.List
	FuncVariadic   bool
	Funcname       *ast.List
	Location       int
	Over           *WindowDef
}

func (n *FuncCall) Pos() int {
	return n.Location
}
