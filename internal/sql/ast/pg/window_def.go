package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type WindowDef struct {
	EndOffset       ast.Node
	FrameOptions    int
	Location        int
	Name            *string
	OrderClause     *ast.List
	PartitionClause *ast.List
	Refname         *string
	StartOffset     ast.Node
}

func (n *WindowDef) Pos() int {
	return n.Location
}
