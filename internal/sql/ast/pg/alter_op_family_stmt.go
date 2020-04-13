package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterOpFamilyStmt struct {
	Amname       *string
	IsDrop       bool
	Items        *ast.List
	Opfamilyname *ast.List
}

func (n *AlterOpFamilyStmt) Pos() int {
	return 0
}
