package pg

import (
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

type RawStmt struct {
	Stmt         ast.Node
	StmtLen      int
	StmtLocation int
}

func (n *RawStmt) Pos() int {
	return 0
}
