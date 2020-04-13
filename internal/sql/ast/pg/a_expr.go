package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type A_Expr struct {
	Kind     A_Expr_Kind
	Lexpr    ast.Node
	Location int
	Name     *ast.List
	Rexpr    ast.Node
}

func (n *A_Expr) Pos() int {
	return n.Location
}
