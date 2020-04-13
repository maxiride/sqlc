package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterEnumStmt struct {
	NewVal             *string
	NewValIsAfter      bool
	NewValNeighbor     *string
	OldVal             *string
	SkipIfNewValExists bool
	TypeName           *ast.List
}

func (n *AlterEnumStmt) Pos() int {
	return 0
}
