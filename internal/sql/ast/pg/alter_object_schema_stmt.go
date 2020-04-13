package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterObjectSchemaStmt struct {
	MissingOk  bool
	Newschema  *string
	Object     ast.Node
	ObjectType ObjectType
	Relation   *RangeVar
}

func (n *AlterObjectSchemaStmt) Pos() int {
	return 0
}
