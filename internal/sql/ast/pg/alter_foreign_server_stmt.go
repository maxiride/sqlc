package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterForeignServerStmt struct {
	HasVersion bool
	Options    *ast.List
	Servername *string
	Version    *string
}

func (n *AlterForeignServerStmt) Pos() int {
	return 0
}
