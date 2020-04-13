package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SubLink struct {
	Location    int
	OperName    *ast.List
	SubLinkId   int
	SubLinkType SubLinkType
	Subselect   ast.Node
	Testexpr    ast.Node
	Xpr         ast.Node
}

func (n *SubLink) Pos() int {
	return n.Location
}
