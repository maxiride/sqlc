package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type IndexElem struct {
	Collation     *ast.List
	Expr          ast.Node
	Indexcolname  *string
	Name          *string
	NullsOrdering SortByNulls
	Opclass       *ast.List
	Ordering      SortByDir
}

func (n *IndexElem) Pos() int {
	return 0
}
