package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type XmlExpr struct {
	ArgNames  *ast.List
	Args      *ast.List
	Location  int
	Name      *string
	NamedArgs *ast.List
	Op        XmlExprOp
	Type      Oid
	Typmod    int32
	Xmloption XmlOptionType
	Xpr       ast.Node
}

func (n *XmlExpr) Pos() int {
	return n.Location
}
