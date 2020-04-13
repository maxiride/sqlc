package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateSeqStmt struct {
	ForIdentity bool
	IfNotExists bool
	Options     *ast.List
	OwnerId     Oid
	Sequence    *RangeVar
}

func (n *CreateSeqStmt) Pos() int {
	return 0
}
