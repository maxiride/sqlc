package pg

type ParamExecData struct {
	ExecPlan interface{}
	Isnull   bool
	Value    Datum
}

func (n *ParamExecData) Pos() int {
	return 0
}
