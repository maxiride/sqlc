package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterFunctionStmt struct {
	Actions *ast.List
	Func    *ObjectWithArgs
}

func (n *AlterFunctionStmt) Pos() int {
	return 0
}
