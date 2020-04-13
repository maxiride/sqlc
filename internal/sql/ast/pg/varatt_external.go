package pg

type varatt_external struct {
	VaExtsize    int32
	VaRawsize    int32
	VaToastrelid Oid
	VaValueid    Oid
}

func (n *varatt_external) Pos() int {
	return 0
}
