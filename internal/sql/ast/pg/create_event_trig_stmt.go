package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateEventTrigStmt struct {
	Eventname  *string
	Funcname   *ast.List
	Trigname   *string
	Whenclause *ast.List
}

func (n *CreateEventTrigStmt) Pos() int {
	return 0
}
