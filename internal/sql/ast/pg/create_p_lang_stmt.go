package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreatePLangStmt struct {
	Plhandler   *ast.List
	Plinline    *ast.List
	Plname      *string
	Pltrusted   bool
	Plvalidator *ast.List
	Replace     bool
}

func (n *CreatePLangStmt) Pos() int {
	return 0
}
