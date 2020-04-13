package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type AlterTSConfigurationStmt struct {
	Cfgname   *ast.List
	Dicts     *ast.List
	Kind      AlterTSConfigType
	MissingOk bool
	Override  bool
	Replace   bool
	Tokentype *ast.List
}

func (n *AlterTSConfigurationStmt) Pos() int {
	return 0
}
