package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type FieldStore struct {
	Arg        ast.Node
	Fieldnums  *ast.List
	Newvals    *ast.List
	Resulttype Oid
	Xpr        ast.Node
}

func (n *FieldStore) Pos() int {
	return 0
}
