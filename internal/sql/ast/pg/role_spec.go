package pg

type RoleSpec struct {
	Location int
	Rolename *string
	Roletype RoleSpecType
}

func (n *RoleSpec) Pos() int {
	return n.Location
}
