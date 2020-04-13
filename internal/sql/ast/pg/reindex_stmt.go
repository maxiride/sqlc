package pg

type ReindexStmt struct {
	Kind     ReindexObjectType
	Name     *string
	Options  int
	Relation *RangeVar
}

func (n *ReindexStmt) Pos() int {
	return 0
}
