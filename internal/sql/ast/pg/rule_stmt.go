package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RuleStmt struct {
	Actions     *ast.List
	Event       CmdType
	Instead     bool
	Relation    *RangeVar
	Replace     bool
	Rulename    *string
	WhereClause ast.Node
}

func (n *RuleStmt) Pos() int {
	return 0
}
