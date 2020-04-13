package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateDomainStmt struct {
	CollClause  *CollateClause
	Constraints *ast.List
	Domainname  *ast.List
	TypeName    *TypeName
}

func (n *CreateDomainStmt) Pos() int {
	return 0
}
