package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterPolicyStmt struct {
	PolicyName *string
	Qual       ast.Node
	Roles      *ast.List
	Table      *RangeVar
	WithCheck  ast.Node
}

func (n *AlterPolicyStmt) Pos() int {
	return 0
}
