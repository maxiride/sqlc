package pg

type InlineCodeBlock struct {
	LangIsTrusted bool
	LangOid       Oid
	SourceText    *string
}

func (n *InlineCodeBlock) Pos() int {
	return 0
}
