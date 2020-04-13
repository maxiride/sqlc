package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateSchemaStmt struct {
	Authrole    *RoleSpec
	IfNotExists bool
	SchemaElts  *ast.List
	Schemaname  *string
}

func (n *CreateSchemaStmt) Pos() int {
	return 0
}
