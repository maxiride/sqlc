package pg

type ParamExternData struct {
	Isnull bool
	Pflags uint16
	Ptype  Oid
	Value  Datum
}

func (n *ParamExternData) Pos() int {
	return 0
}
