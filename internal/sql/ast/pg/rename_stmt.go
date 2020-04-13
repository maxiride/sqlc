package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RenameStmt struct {
	Behavior     DropBehavior
	MissingOk    bool
	Newname      *string
	Object       ast.Node
	Relation     *RangeVar
	RelationType ObjectType
	RenameType   ObjectType
	Subname      *string
}

func (n *RenameStmt) Pos() int {
	return 0
}
