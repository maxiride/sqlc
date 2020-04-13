package pg

type TriggerTransition struct {
	IsNew   bool
	IsTable bool
	Name    *string
}

func (n *TriggerTransition) Pos() int {
	return 0
}
