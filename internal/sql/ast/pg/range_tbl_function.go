package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RangeTblFunction struct {
	Funccolcollations *ast.List
	Funccolcount      int
	Funccolnames      *ast.List
	Funccoltypes      *ast.List
	Funccoltypmods    *ast.List
	Funcexpr          ast.Node
	Funcparams        []uint32
}

func (n *RangeTblFunction) Pos() int {
	return 0
}
