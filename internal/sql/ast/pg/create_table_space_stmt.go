package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateTableSpaceStmt struct {
	Location       *string
	Options        *ast.List
	Owner          *RoleSpec
	Tablespacename *string
}

func (n *CreateTableSpaceStmt) Pos() int {
	return 0
}
