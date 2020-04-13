package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateSubscriptionStmt struct {
	Conninfo    *string
	Options     *ast.List
	Publication *ast.List
	Subname     *string
}

func (n *CreateSubscriptionStmt) Pos() int {
	return 0
}
