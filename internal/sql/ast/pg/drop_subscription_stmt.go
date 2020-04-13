package pg

type DropSubscriptionStmt struct {
	Behavior  DropBehavior
	MissingOk bool
	Subname   *string
}

func (n *DropSubscriptionStmt) Pos() int {
	return 0
}
