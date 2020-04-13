package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type TableFunc struct {
	Colcollations *ast.List
	Coldefexprs   *ast.List
	Colexprs      *ast.List
	Colnames      *ast.List
	Coltypes      *ast.List
	Coltypmods    *ast.List
	Docexpr       ast.Node
	Location      int
	Notnulls      []uint32
	NsNames       *ast.List
	NsUris        *ast.List
	Ordinalitycol int
	Rowexpr       ast.Node
}

func (n *TableFunc) Pos() int {
	return n.Location
}
