package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type WindowClause struct {
	CopiedOrder     bool
	EndOffset       ast.Node
	FrameOptions    int
	Name            *string
	OrderClause     *ast.List
	PartitionClause *ast.List
	Refname         *string
	StartOffset     ast.Node
	Winref          Index
}

func (n *WindowClause) Pos() int {
	return 0
}
