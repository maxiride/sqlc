package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ArrayRef struct {
	Refarraytype    Oid
	Refassgnexpr    ast.Node
	Refcollid       Oid
	Refelemtype     Oid
	Refexpr         ast.Node
	Reflowerindexpr *ast.List
	Reftypmod       int32
	Refupperindexpr *ast.List
	Xpr             ast.Node
}

func (n *ArrayRef) Pos() int {
	return 0
}
