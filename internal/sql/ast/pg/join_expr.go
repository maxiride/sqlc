package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type JoinExpr struct {
	Alias       *Alias
	IsNatural   bool
	Jointype    JoinType
	Larg        ast.Node
	Quals       ast.Node
	Rarg        ast.Node
	Rtindex     int
	UsingClause *ast.List
}

func (n *JoinExpr) Pos() int {
	return 0
}
