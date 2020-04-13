package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SetOperationStmt struct {
	All           bool
	ColCollations *ast.List
	ColTypes      *ast.List
	ColTypmods    *ast.List
	GroupClauses  *ast.List
	Larg          ast.Node
	Op            SetOperation
	Rarg          ast.Node
}

func (n *SetOperationStmt) Pos() int {
	return 0
}
