package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterTableSpaceOptionsStmt struct {
	IsReset        bool
	Options        *ast.List
	Tablespacename *string
}

func (n *AlterTableSpaceOptionsStmt) Pos() int {
	return 0
}
