package pg

type FetchStmt struct {
	Direction  FetchDirection
	HowMany    int64
	Ismove     bool
	Portalname *string
}

func (n *FetchStmt) Pos() int {
	return 0
}
