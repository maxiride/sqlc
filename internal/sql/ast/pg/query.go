package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type Query struct {
	CanSetTag        bool
	CommandType      CmdType
	ConstraintDeps   *ast.List
	CteList          *ast.List
	DistinctClause   *ast.List
	GroupClause      *ast.List
	GroupingSets     *ast.List
	HasAggs          bool
	HasDistinctOn    bool
	HasForUpdate     bool
	HasModifyingCte  bool
	HasRecursive     bool
	HasRowSecurity   bool
	HasSubLinks      bool
	HasTargetSrfs    bool
	HasWindowFuncs   bool
	HavingQual       ast.Node
	Jointree         *FromExpr
	LimitCount       ast.Node
	LimitOffset      ast.Node
	OnConflict       *OnConflictExpr
	Override         OverridingKind
	QueryId          uint32
	QuerySource      QuerySource
	ResultRelation   int
	ReturningList    *ast.List
	RowMarks         *ast.List
	Rtable           *ast.List
	SetOperations    ast.Node
	SortClause       *ast.List
	StmtLen          int
	StmtLocation     int
	TargetList       *ast.List
	UtilityStmt      ast.Node
	WindowClause     *ast.List
	WithCheckOptions *ast.List
}

func (n *Query) Pos() int {
	return 0
}
