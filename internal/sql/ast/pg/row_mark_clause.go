package pg

type RowMarkClause struct {
	PushedDown bool
	Rti        Index
	Strength   LockClauseStrength
	WaitPolicy LockWaitPolicy
}

func (n *RowMarkClause) Pos() int {
	return 0
}
