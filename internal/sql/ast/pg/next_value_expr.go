package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type NextValueExpr struct {
	Seqid  Oid
	TypeId Oid
	Xpr    ast.Node
}

func (n *NextValueExpr) Pos() int {
	return 0
}
