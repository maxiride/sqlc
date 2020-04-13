package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateTableAsStmt struct {
	IfNotExists  bool
	Into         *IntoClause
	IsSelectInto bool
	Query        ast.Node
	Relkind      ObjectType
}

func (n *CreateTableAsStmt) Pos() int {
	return 0
}
