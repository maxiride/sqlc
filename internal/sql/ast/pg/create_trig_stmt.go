package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateTrigStmt struct {
	Args           *ast.List
	Columns        *ast.List
	Constrrel      *RangeVar
	Deferrable     bool
	Events         int16
	Funcname       *ast.List
	Initdeferred   bool
	Isconstraint   bool
	Relation       *RangeVar
	Row            bool
	Timing         int16
	TransitionRels *ast.List
	Trigname       *string
	WhenClause     ast.Node
}

func (n *CreateTrigStmt) Pos() int {
	return 0
}
