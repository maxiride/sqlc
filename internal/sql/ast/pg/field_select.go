package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type FieldSelect struct {
	Arg          ast.Node
	Fieldnum     AttrNumber
	Resultcollid Oid
	Resulttype   Oid
	Resulttypmod int32
	Xpr          ast.Node
}

func (n *FieldSelect) Pos() int {
	return 0
}
