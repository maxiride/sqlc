package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SubPlan struct {
	Args              *ast.List
	FirstColCollation Oid
	FirstColType      Oid
	FirstColTypmod    int32
	ParParam          *ast.List
	ParallelSafe      bool
	ParamIds          *ast.List
	PerCallCost       Cost
	PlanId            int
	PlanName          *string
	SetParam          *ast.List
	StartupCost       Cost
	SubLinkType       SubLinkType
	Testexpr          ast.Node
	UnknownEqFalse    bool
	UseHashTable      bool
	Xpr               ast.Node
}

func (n *SubPlan) Pos() int {
	return 0
}
