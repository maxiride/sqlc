package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateRangeStmt struct {
	Params   *ast.List
	TypeName *ast.List
}

func (n *CreateRangeStmt) Pos() int {
	return 0
}
