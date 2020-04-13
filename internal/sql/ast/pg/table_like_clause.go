package pg

type TableLikeClause struct {
	Options  uint32
	Relation *RangeVar
}

func (n *TableLikeClause) Pos() int {
	return 0
}
