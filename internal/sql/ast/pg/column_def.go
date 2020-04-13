package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type ColumnDef struct {
	CollClause    *CollateClause
	CollOid       Oid
	Colname       *string
	Constraints   *ast.List
	CookedDefault ast.Node
	Fdwoptions    *ast.List
	Identity      byte
	Inhcount      int
	IsFromParent  bool
	IsFromType    bool
	IsLocal       bool
	IsNotNull     bool
	Location      int
	RawDefault    ast.Node
	Storage       byte
	TypeName      *TypeName
}

func (n *ColumnDef) Pos() int {
	return n.Location
}
