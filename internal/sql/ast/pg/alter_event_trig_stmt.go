package pg

type AlterEventTrigStmt struct {
	Tgenabled byte
	Trigname  *string
}

func (n *AlterEventTrigStmt) Pos() int {
	return 0
}
