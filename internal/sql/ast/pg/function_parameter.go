package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type FunctionParameter struct {
	ArgType *TypeName
	Defexpr ast.Node
	Mode    FunctionParameterMode
	Name    *string
}

func (n *FunctionParameter) Pos() int {
	return 0
}
