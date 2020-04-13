package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type InferenceElem struct {
	Expr         ast.Node
	Infercollid  Oid
	Inferopclass Oid
	Xpr          ast.Node
}

func (n *InferenceElem) Pos() int {
	return 0
}
