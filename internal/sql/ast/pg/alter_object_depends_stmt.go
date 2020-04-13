package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterObjectDependsStmt struct {
	Extname    ast.Node
	Object     ast.Node
	ObjectType ObjectType
	Relation   *RangeVar
}

func (n *AlterObjectDependsStmt) Pos() int {
	return 0
}
