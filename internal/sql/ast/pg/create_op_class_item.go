package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateOpClassItem struct {
	ClassArgs   *ast.List
	Itemtype    int
	Name        *ObjectWithArgs
	Number      int
	OrderFamily *ast.List
	Storedtype  *TypeName
}

func (n *CreateOpClassItem) Pos() int {
	return 0
}
