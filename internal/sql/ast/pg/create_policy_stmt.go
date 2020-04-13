package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreatePolicyStmt struct {
	CmdName    *string
	Permissive bool
	PolicyName *string
	Qual       ast.Node
	Roles      *ast.List
	Table      *RangeVar
	WithCheck  ast.Node
}

func (n *CreatePolicyStmt) Pos() int {
	return 0
}
