package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type InferClause struct {
	Conname     *string
	IndexElems  *ast.List
	Location    int
	WhereClause ast.Node
}

func (n *InferClause) Pos() int {
	return n.Location
}
