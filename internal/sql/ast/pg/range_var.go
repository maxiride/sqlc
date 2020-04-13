package pg

type RangeVar struct {
	Alias          *Alias
	Catalogname    *string
	Inh            bool
	Location       int
	Relname        *string
	Relpersistence byte
	Schemaname     *string
}

func (n *RangeVar) Pos() int {
	return n.Location
}
