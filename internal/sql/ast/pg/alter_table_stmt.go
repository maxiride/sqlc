package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterTableStmt struct {
	Cmds      *ast.List
	MissingOk bool
	Relation  *RangeVar
	Relkind   ObjectType
}

func (n *AlterTableStmt) Pos() int {
	return 0
}
