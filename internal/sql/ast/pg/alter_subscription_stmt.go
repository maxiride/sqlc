package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterSubscriptionStmt struct {
	Conninfo    *string
	Kind        AlterSubscriptionType
	Options     *ast.List
	Publication *ast.List
	Subname     *string
}

func (n *AlterSubscriptionStmt) Pos() int {
	return 0
}
