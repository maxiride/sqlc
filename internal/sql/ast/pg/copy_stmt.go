package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CopyStmt struct {
	Attlist   *ast.List
	Filename  *string
	IsFrom    bool
	IsProgram bool
	Options   *ast.List
	Query     ast.Node
	Relation  *RangeVar
}

func (n *CopyStmt) Pos() int {
	return 0
}
