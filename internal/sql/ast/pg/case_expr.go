package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CaseExpr struct {
	Arg        ast.Node
	Args       *ast.List
	Casecollid Oid
	Casetype   Oid
	Defresult  ast.Node
	Location   int
	Xpr        ast.Node
}

func (n *CaseExpr) Pos() int {
	return n.Location
}
