package pg

type DropUserMappingStmt struct {
	MissingOk  bool
	Servername *string
	User       *RoleSpec
}

func (n *DropUserMappingStmt) Pos() int {
	return 0
}
