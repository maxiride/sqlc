package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ObjectWithArgs struct {
	ArgsUnspecified bool
	Objargs         *ast.List
	Objname         *ast.List
}

func (n *ObjectWithArgs) Pos() int {
	return 0
}
