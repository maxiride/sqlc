package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type IntoClause struct {
	ColNames       *ast.List
	OnCommit       OnCommitAction
	Options        *ast.List
	Rel            *RangeVar
	SkipData       bool
	TableSpaceName *string
	ViewQuery      ast.Node
}

func (n *IntoClause) Pos() int {
	return 0
}
