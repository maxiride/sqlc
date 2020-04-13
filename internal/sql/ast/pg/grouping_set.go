package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type GroupingSet struct {
	Content  *ast.List
	Kind     GroupingSetKind
	Location int
}

func (n *GroupingSet) Pos() int {
	return n.Location
}
