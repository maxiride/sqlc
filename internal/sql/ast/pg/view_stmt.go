package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ViewStmt struct {
	Aliases         *ast.List
	Options         *ast.List
	Query           ast.Node
	Replace         bool
	View            *RangeVar
	WithCheckOption ViewCheckOption
}

func (n *ViewStmt) Pos() int {
	return 0
}
