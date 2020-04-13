package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type CreateConversionStmt struct {
	ConversionName  *ast.List
	Def             bool
	ForEncodingName *string
	FuncName        *ast.List
	ToEncodingName  *string
}

func (n *CreateConversionStmt) Pos() int {
	return 0
}
