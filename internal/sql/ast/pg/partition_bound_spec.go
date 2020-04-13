package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type PartitionBoundSpec struct {
	Listdatums  *ast.List
	Location    int
	Lowerdatums *ast.List
	Strategy    byte
	Upperdatums *ast.List
}

func (n *PartitionBoundSpec) Pos() int {
	return n.Location
}
