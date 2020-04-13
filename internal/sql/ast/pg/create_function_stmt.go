package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateFunctionStmt struct {
	Funcname   *ast.List
	Options    *ast.List
	Parameters *ast.List
	Replace    bool
	ReturnType *TypeName
	WithClause *ast.List
}

func (n *CreateFunctionStmt) Pos() int {
	return 0
}
