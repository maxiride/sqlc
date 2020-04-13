package pg

type DropTableSpaceStmt struct {
	MissingOk      bool
	Tablespacename *string
}

func (n *DropTableSpaceStmt) Pos() int {
	return 0
}
