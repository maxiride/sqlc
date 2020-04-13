package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreatePublicationStmt struct {
	ForAllTables bool
	Options      *ast.List
	Pubname      *string
	Tables       *ast.List
}

func (n *CreatePublicationStmt) Pos() int {
	return 0
}
