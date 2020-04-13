package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type PartitionSpec struct {
	Location   int
	PartParams *ast.List
	Strategy   *string
}

func (n *PartitionSpec) Pos() int {
	return n.Location
}
