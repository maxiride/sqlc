package pg

type AlterRoleSetStmt struct {
	Database *string
	Role     *RoleSpec
	Setstmt  *VariableSetStmt
}

func (n *AlterRoleSetStmt) Pos() int {
	return 0
}
