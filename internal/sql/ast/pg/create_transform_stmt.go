package pg

type CreateTransformStmt struct {
	Fromsql  *ObjectWithArgs
	Lang     *string
	Replace  bool
	Tosql    *ObjectWithArgs
	TypeName *TypeName
}

func (n *CreateTransformStmt) Pos() int {
	return 0
}
