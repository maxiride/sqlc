package pg

type ClusterStmt struct {
	Indexname *string
	Relation  *RangeVar
	Verbose   bool
}

func (n *ClusterStmt) Pos() int {
	return 0
}
