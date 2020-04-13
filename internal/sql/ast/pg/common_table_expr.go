package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CommonTableExpr struct {
	Aliascolnames    *ast.List
	Ctecolcollations *ast.List
	Ctecolnames      *ast.List
	Ctecoltypes      *ast.List
	Ctecoltypmods    *ast.List
	Ctename          *string
	Ctequery         ast.Node
	Cterecursive     bool
	Cterefcount      int
	Location         int
}

func (n *CommonTableExpr) Pos() int {
	return n.Location
}
