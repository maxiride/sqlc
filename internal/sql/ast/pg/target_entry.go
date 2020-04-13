package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type TargetEntry struct {
	Expr            ast.Node
	Resjunk         bool
	Resname         *string
	Resno           AttrNumber
	Resorigcol      AttrNumber
	Resorigtbl      Oid
	Ressortgroupref Index
	Xpr             ast.Node
}

func (n *TargetEntry) Pos() int {
	return 0
}
