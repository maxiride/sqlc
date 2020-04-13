package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ImportForeignSchemaStmt struct {
	ListType     ImportForeignSchemaType
	LocalSchema  *string
	Options      *ast.List
	RemoteSchema *string
	ServerName   *string
	TableList    *ast.List
}

func (n *ImportForeignSchemaStmt) Pos() int {
	return 0
}
