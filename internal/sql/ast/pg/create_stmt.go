package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateStmt struct {
	Constraints    *ast.List
	IfNotExists    bool
	InhRelations   *ast.List
	OfTypename     *TypeName
	Oncommit       OnCommitAction
	Options        *ast.List
	Partbound      *PartitionBoundSpec
	Partspec       *PartitionSpec
	Relation       *RangeVar
	TableElts      *ast.List
	Tablespacename *string
}

func (n *CreateStmt) Pos() int {
	return 0
}
