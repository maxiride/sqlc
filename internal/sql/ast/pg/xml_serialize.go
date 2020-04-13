package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type XmlSerialize struct {
	Expr      ast.Node
	Location  int
	TypeName  *TypeName
	Xmloption XmlOptionType
}

func (n *XmlSerialize) Pos() int {
	return n.Location
}
