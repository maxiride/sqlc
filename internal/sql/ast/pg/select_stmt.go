package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type SelectStmt struct {
	All            bool
	DistinctClause *ast.List
	FromClause     *ast.List
	GroupClause    *ast.List
	HavingClause   ast.Node
	IntoClause     *IntoClause
	Larg           *SelectStmt
	LimitCount     ast.Node
	LimitOffset    ast.Node
	LockingClause  *ast.List
	Op             SetOperation
	Rarg           *SelectStmt
	SortClause     *ast.List
	TargetList     *ast.List
	ValuesLists    [][]ast.Node
	WhereClause    ast.Node
	WindowClause   *ast.List
	WithClause     *WithClause
}

func (n *SelectStmt) Pos() int {
	return 0
}
