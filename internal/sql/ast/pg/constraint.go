package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type Constraint struct {
	AccessMethod   *string
	Conname        *string
	Contype        ConstrType
	CookedExpr     *string
	Deferrable     bool
	Exclusions     *ast.List
	FkAttrs        *ast.List
	FkDelAction    byte
	FkMatchtype    byte
	FkUpdAction    byte
	GeneratedWhen  byte
	Indexname      *string
	Indexspace     *string
	Initdeferred   bool
	InitiallyValid bool
	IsNoInherit    bool
	Keys           *ast.List
	Location       int
	OldConpfeqop   *ast.List
	OldPktableOid  Oid
	Options        *ast.List
	PkAttrs        *ast.List
	Pktable        *RangeVar
	RawExpr        ast.Node
	SkipValidation bool
	WhereClause    ast.Node
}

func (n *Constraint) Pos() int {
	return n.Location
}
