package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateForeignTableStmt struct {
	Base       *CreateStmt
	Options    *ast.List
	Servername *string
}

func (n *CreateForeignTableStmt) Pos() int {
	return 0
}
