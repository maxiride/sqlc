package pg

type RefreshMatViewStmt struct {
	Concurrent bool
	Relation   *RangeVar
	SkipData   bool
}

func (n *RefreshMatViewStmt) Pos() int {
	return 0
}
