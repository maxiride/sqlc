package pg

type ParamRef struct {
	Location int
	Number   int
}

func (n *ParamRef) Pos() int {
	return n.Location
}
