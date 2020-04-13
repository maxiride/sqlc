package pg

type CreateCastStmt struct {
	Context    CoercionContext
	Func       *ObjectWithArgs
	Inout      bool
	Sourcetype *TypeName
	Targettype *TypeName
}

func (n *CreateCastStmt) Pos() int {
	return 0
}
