package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RangeFunction struct {
	Alias      *Alias
	Coldeflist *ast.List
	Functions  *ast.List
	IsRowsfrom bool
	Lateral    bool
	Ordinality bool
}

func (n *RangeFunction) Pos() int {
	return 0
}
