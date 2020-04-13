package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type TransactionStmt struct {
	Gid     *string
	Kind    TransactionStmtKind
	Options *ast.List
}

func (n *TransactionStmt) Pos() int {
	return 0
}
