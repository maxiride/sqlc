package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateForeignServerStmt struct {
	Fdwname     *string
	IfNotExists bool
	Options     *ast.List
	Servername  *string
	Servertype  *string
	Version     *string
}

func (n *CreateForeignServerStmt) Pos() int {
	return 0
}
