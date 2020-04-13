package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlternativeSubPlan struct {
	Subplans *ast.List
	Xpr      ast.Node
}

func (n *AlternativeSubPlan) Pos() int {
	return 0
}
