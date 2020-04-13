package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type WithClause struct {
	Ctes      *ast.List
	Location  int
	Recursive bool
}

func (n *WithClause) Pos() int {
	return n.Location
}
