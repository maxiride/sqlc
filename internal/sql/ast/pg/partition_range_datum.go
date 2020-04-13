package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type PartitionRangeDatum struct {
	Kind     PartitionRangeDatumKind
	Location int
	Value    ast.Node
}

func (n *PartitionRangeDatum) Pos() int {
	return n.Location
}
