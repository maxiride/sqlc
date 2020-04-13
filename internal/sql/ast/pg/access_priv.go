package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AccessPriv struct {
	Cols     *ast.List
	PrivName *string
}

func (n *AccessPriv) Pos() int {
	return 0
}
