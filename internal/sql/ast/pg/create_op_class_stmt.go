package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateOpClassStmt struct {
	Amname       *string
	Datatype     *TypeName
	IsDefault    bool
	Items        *ast.List
	Opclassname  *ast.List
	Opfamilyname *ast.List
}

func (n *CreateOpClassStmt) Pos() int {
	return 0
}
