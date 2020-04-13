package pg

type PartitionCmd struct {
	Bound *PartitionBoundSpec
	Name  *RangeVar
}

func (n *PartitionCmd) Pos() int {
	return 0
}
