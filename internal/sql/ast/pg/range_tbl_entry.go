package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RangeTblEntry struct {
	Alias           *Alias
	CheckAsUser     Oid
	Colcollations   *ast.List
	Coltypes        *ast.List
	Coltypmods      *ast.List
	Ctelevelsup     Index
	Ctename         *string
	Enrname         *string
	Enrtuples       float64
	Eref            *Alias
	Funcordinality  bool
	Functions       *ast.List
	InFromCl        bool
	Inh             bool
	InsertedCols    []uint32
	Joinaliasvars   *ast.List
	Jointype        JoinType
	Lateral         bool
	Relid           Oid
	Relkind         byte
	RequiredPerms   AclMode
	Rtekind         RTEKind
	SecurityBarrier bool
	SecurityQuals   *ast.List
	SelectedCols    []uint32
	SelfReference   bool
	Subquery        *Query
	Tablefunc       *TableFunc
	Tablesample     *TableSampleClause
	UpdatedCols     []uint32
	ValuesLists     *ast.List
}

func (n *RangeTblEntry) Pos() int {
	return 0
}
