package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterTableMoveAllStmt struct {
	NewTablespacename  *string
	Nowait             bool
	Objtype            ObjectType
	OrigTablespacename *string
	Roles              *ast.List
}

func (n *AlterTableMoveAllStmt) Pos() int {
	return 0
}
