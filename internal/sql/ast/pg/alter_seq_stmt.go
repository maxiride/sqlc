package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterSeqStmt struct {
	ForIdentity bool
	MissingOk   bool
	Options     *ast.List
	Sequence    *RangeVar
}

func (n *AlterSeqStmt) Pos() int {
	return 0
}
