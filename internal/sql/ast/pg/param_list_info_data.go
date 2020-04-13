package pg

type ParamListInfoData struct {
	NumParams      int
	ParamFetchArg  interface{}
	ParamMask      []uint32
	ParserSetupArg interface{}
}

func (n *ParamListInfoData) Pos() int {
	return 0
}
