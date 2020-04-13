package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CurrentOfExpr struct {
	CursorName  *string
	CursorParam int
	Cvarno      Index
	Xpr         ast.Node
}

func (n *CurrentOfExpr) Pos() int {
	return 0
}
