package pg

type SortGroupClause struct {
	Eqop            Oid
	Hashable        bool
	NullsFirst      bool
	Sortop          Oid
	TleSortGroupRef Index
}

func (n *SortGroupClause) Pos() int {
	return 0
}
