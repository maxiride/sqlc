package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterPublicationStmt struct {
	ForAllTables bool
	Options      *ast.List
	Pubname      *string
	TableAction  DefElemAction
	Tables       *ast.List
}

func (n *AlterPublicationStmt) Pos() int {
	return 0
}
