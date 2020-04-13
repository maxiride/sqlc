package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type DefineStmt struct {
	Args        *ast.List
	Definition  *ast.List
	Defnames    *ast.List
	IfNotExists bool
	Kind        ObjectType
	Oldstyle    bool
}

func (n *DefineStmt) Pos() int {
	return 0
}
