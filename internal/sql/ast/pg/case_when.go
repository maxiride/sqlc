package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CaseWhen struct {
	Expr     ast.Node
	Location int
	Result   ast.Node
	Xpr      ast.Node
}

func (n *CaseWhen) Pos() int {
	return n.Location
}
