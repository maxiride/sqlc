package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateAmStmt struct {
	Amname      *string
	Amtype      byte
	HandlerName *ast.List
}

func (n *CreateAmStmt) Pos() int {
	return 0
}
