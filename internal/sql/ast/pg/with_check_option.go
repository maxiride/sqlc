package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type WithCheckOption struct {
	Cascaded bool
	Kind     WCOKind
	Polname  *string
	Qual     ast.Node
	Relname  *string
}

func (n *WithCheckOption) Pos() int {
	return 0
}
