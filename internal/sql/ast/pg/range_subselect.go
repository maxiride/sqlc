package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RangeSubselect struct {
	Alias    *Alias
	Lateral  bool
	Subquery ast.Node
}

func (n *RangeSubselect) Pos() int {
	return 0
}
