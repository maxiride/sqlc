package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateOpFamilyStmt struct {
	Amname       *string
	Opfamilyname *ast.List
}

func (n *CreateOpFamilyStmt) Pos() int {
	return 0
}
