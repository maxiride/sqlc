package postgresql

import (
	nodes "github.com/lfittl/pg_query_go/nodes"

	"github.com/kyleconroy/sqlc/cmd/sqlc-gen/pg"
	"github.com/kyleconroy/sqlc/internal/sql/ast"
)

func convertList(l nodes.List) *ast.List {
	out := &ast.List{}
	for _, item := range l.Items {
		out.Items = append(out.Items, convertNode(item))
	}
	return out
}

func convertValuesList(l [][]nodes.Node) [][]ast.Node {
	out := [][]ast.Node{}
	for _, outer := range l {
		o := []ast.Node{}
		for _, inner := range outer {
			o = append(o, convertNode(inner))
		}
		out = append(out, o)
	}
	return out
}

func convert(node nodes.Node) (ast.Node, error) {
	return convertNode(node), nil
}

func convertUnlistenStmt(n *nodes.UnlistenStmt) *pg.UnlistenStmt {
	return &pg.UnlistenStmt{
		Conditionname: n.Conditionname,
	}
}

func convertDeleteStmt(n *nodes.DeleteStmt) *pg.DeleteStmt {
	return &pg.DeleteStmt{
		Relation:      convertRangeVar(n.Relation),
		UsingClause:   convertList(n.UsingClause),
		WhereClause:   convertNode(n.WhereClause),
		ReturningList: convertList(n.ReturningList),
		WithClause:    convertWithClause(n.WithClause),
	}
}

func convertRelabelType(n *nodes.RelabelType) *pg.RelabelType {
	return &pg.RelabelType{
		Xpr:           convertNode(n.Xpr),
		Arg:           convertNode(n.Arg),
		Resulttype:    pg.Oid(n.Resulttype),
		Resulttypmod:  n.Resulttypmod,
		Resultcollid:  pg.Oid(n.Resultcollid),
		Relabelformat: pg.CoercionForm(n.Relabelformat),
		Location:      n.Location,
	}
}

func convertCreateUserMappingStmt(n *nodes.CreateUserMappingStmt) *pg.CreateUserMappingStmt {
	return &pg.CreateUserMappingStmt{
		User:        convertRoleSpec(n.User),
		Servername:  n.Servername,
		IfNotExists: n.IfNotExists,
		Options:     convertList(n.Options),
	}
}

func convertCreateAmStmt(n *nodes.CreateAmStmt) *pg.CreateAmStmt {
	return &pg.CreateAmStmt{
		Amname:      n.Amname,
		HandlerName: convertList(n.HandlerName),
		Amtype:      n.Amtype,
	}
}

func convertGrantRoleStmt(n *nodes.GrantRoleStmt) *pg.GrantRoleStmt {
	return &pg.GrantRoleStmt{
		GrantedRoles: convertList(n.GrantedRoles),
		GranteeRoles: convertList(n.GranteeRoles),
		IsGrant:      n.IsGrant,
		AdminOpt:     n.AdminOpt,
		Grantor:      convertRoleSpec(n.Grantor),
		Behavior:     pg.DropBehavior(n.Behavior),
	}
}

func convertFuncExpr(n *nodes.FuncExpr) *pg.FuncExpr {
	return &pg.FuncExpr{
		Xpr:            convertNode(n.Xpr),
		Funcid:         pg.Oid(n.Funcid),
		Funcresulttype: pg.Oid(n.Funcresulttype),
		Funcretset:     n.Funcretset,
		Funcvariadic:   n.Funcvariadic,
		Funcformat:     pg.CoercionForm(n.Funcformat),
		Funccollid:     pg.Oid(n.Funccollid),
		Inputcollid:    pg.Oid(n.Inputcollid),
		Args:           convertList(n.Args),
		Location:       n.Location,
	}
}

func convertCaseTestExpr(n *nodes.CaseTestExpr) *pg.CaseTestExpr {
	return &pg.CaseTestExpr{
		Xpr:       convertNode(n.Xpr),
		TypeId:    pg.Oid(n.TypeId),
		TypeMod:   n.TypeMod,
		Collation: pg.Oid(n.Collation),
	}
}

func convertCaseWhen(n *nodes.CaseWhen) *pg.CaseWhen {
	return &pg.CaseWhen{
		Xpr:      convertNode(n.Xpr),
		Expr:     convertNode(n.Expr),
		Result:   convertNode(n.Result),
		Location: n.Location,
	}
}

func convertResTarget(n *nodes.ResTarget) *pg.ResTarget {
	return &pg.ResTarget{
		Name:        n.Name,
		Indirection: convertList(n.Indirection),
		Val:         convertNode(n.Val),
		Location:    n.Location,
	}
}

func convertAlterExtensionContentsStmt(n *nodes.AlterExtensionContentsStmt) *pg.AlterExtensionContentsStmt {
	return &pg.AlterExtensionContentsStmt{
		Extname: n.Extname,
		Action:  n.Action,
		Objtype: pg.ObjectType(n.Objtype),
		Object:  convertNode(n.Object),
	}
}

func convertReindexStmt(n *nodes.ReindexStmt) *pg.ReindexStmt {
	return &pg.ReindexStmt{
		Kind:     pg.ReindexObjectType(n.Kind),
		Relation: convertRangeVar(n.Relation),
		Name:     n.Name,
		Options:  n.Options,
	}
}

func convertWindowDef(n *nodes.WindowDef) *pg.WindowDef {
	return &pg.WindowDef{
		Name:            n.Name,
		Refname:         n.Refname,
		PartitionClause: convertList(n.PartitionClause),
		OrderClause:     convertList(n.OrderClause),
		FrameOptions:    n.FrameOptions,
		StartOffset:     convertNode(n.StartOffset),
		EndOffset:       convertNode(n.EndOffset),
		Location:        n.Location,
	}
}

func convertDiscardStmt(n *nodes.DiscardStmt) *pg.DiscardStmt {
	return &pg.DiscardStmt{
		Target: pg.DiscardMode(n.Target),
	}
}

func convertParamListInfoData(n *nodes.ParamListInfoData) *pg.ParamListInfoData {
	return &pg.ParamListInfoData{
		ParamFetchArg:  &ast.TODO{},
		ParserSetupArg: &ast.TODO{},
		NumParams:      n.NumParams,
		ParamMask:      n.ParamMask,
	}
}

func convertCollateClause(n *nodes.CollateClause) *pg.CollateClause {
	return &pg.CollateClause{
		Arg:      convertNode(n.Arg),
		Collname: convertList(n.Collname),
		Location: n.Location,
	}
}

func convertCreatedbStmt(n *nodes.CreatedbStmt) *pg.CreatedbStmt {
	return &pg.CreatedbStmt{
		Dbname:  n.Dbname,
		Options: convertList(n.Options),
	}
}

func convertNull(n *nodes.Null) *pg.Null {
	return &pg.Null{}
}

func convertRefreshMatViewStmt(n *nodes.RefreshMatViewStmt) *pg.RefreshMatViewStmt {
	return &pg.RefreshMatViewStmt{
		Concurrent: n.Concurrent,
		SkipData:   n.SkipData,
		Relation:   convertRangeVar(n.Relation),
	}
}

func convertListenStmt(n *nodes.ListenStmt) *pg.ListenStmt {
	return &pg.ListenStmt{
		Conditionname: n.Conditionname,
	}
}

func convertExpr(n *nodes.Expr) *pg.Expr {
	return &pg.Expr{}
}

func convertSortGroupClause(n *nodes.SortGroupClause) *pg.SortGroupClause {
	return &pg.SortGroupClause{
		TleSortGroupRef: pg.Index(n.TleSortGroupRef),
		Eqop:            pg.Oid(n.Eqop),
		Sortop:          pg.Oid(n.Sortop),
		NullsFirst:      n.NullsFirst,
		Hashable:        n.Hashable,
	}
}

func convertCreateExtensionStmt(n *nodes.CreateExtensionStmt) *pg.CreateExtensionStmt {
	return &pg.CreateExtensionStmt{
		Extname:     n.Extname,
		IfNotExists: n.IfNotExists,
		Options:     convertList(n.Options),
	}
}

func convertCreateEnumStmt(n *nodes.CreateEnumStmt) *pg.CreateEnumStmt {
	return &pg.CreateEnumStmt{
		TypeName: convertList(n.TypeName),
		Vals:     convertList(n.Vals),
	}
}

func convertDoStmt(n *nodes.DoStmt) *pg.DoStmt {
	return &pg.DoStmt{
		Args: convertList(n.Args),
	}
}

func convertXmlSerialize(n *nodes.XmlSerialize) *pg.XmlSerialize {
	return &pg.XmlSerialize{
		Xmloption: pg.XmlOptionType(n.Xmloption),
		Expr:      convertNode(n.Expr),
		TypeName:  convertTypeName(n.TypeName),
		Location:  n.Location,
	}
}

func convertPartitionCmd(n *nodes.PartitionCmd) *pg.PartitionCmd {
	return &pg.PartitionCmd{
		Name:  convertRangeVar(n.Name),
		Bound: convertPartitionBoundSpec(n.Bound),
	}
}

func convertWithCheckOption(n *nodes.WithCheckOption) *pg.WithCheckOption {
	return &pg.WithCheckOption{
		Kind:     pg.WCOKind(n.Kind),
		Relname:  n.Relname,
		Polname:  n.Polname,
		Qual:     convertNode(n.Qual),
		Cascaded: n.Cascaded,
	}
}

func convertCreateRoleStmt(n *nodes.CreateRoleStmt) *pg.CreateRoleStmt {
	return &pg.CreateRoleStmt{
		StmtType: pg.RoleStmtType(n.StmtType),
		Role:     n.Role,
		Options:  convertList(n.Options),
	}
}

func convertClosePortalStmt(n *nodes.ClosePortalStmt) *pg.ClosePortalStmt {
	return &pg.ClosePortalStmt{
		Portalname: n.Portalname,
	}
}

func convertRangeTableFunc(n *nodes.RangeTableFunc) *pg.RangeTableFunc {
	return &pg.RangeTableFunc{
		Lateral:    n.Lateral,
		Docexpr:    convertNode(n.Docexpr),
		Rowexpr:    convertNode(n.Rowexpr),
		Namespaces: convertList(n.Namespaces),
		Columns:    convertList(n.Columns),
		Alias:      convertAlias(n.Alias),
		Location:   n.Location,
	}
}

func convertDropdbStmt(n *nodes.DropdbStmt) *pg.DropdbStmt {
	return &pg.DropdbStmt{
		Dbname:    n.Dbname,
		MissingOk: n.MissingOk,
	}
}

func convertRangeSubselect(n *nodes.RangeSubselect) *pg.RangeSubselect {
	return &pg.RangeSubselect{
		Lateral:  n.Lateral,
		Subquery: convertNode(n.Subquery),
		Alias:    convertAlias(n.Alias),
	}
}

func convertAlterSystemStmt(n *nodes.AlterSystemStmt) *pg.AlterSystemStmt {
	return &pg.AlterSystemStmt{
		Setstmt: convertVariableSetStmt(n.Setstmt),
	}
}

func convertArrayCoerceExpr(n *nodes.ArrayCoerceExpr) *pg.ArrayCoerceExpr {
	return &pg.ArrayCoerceExpr{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Elemfuncid:   pg.Oid(n.Elemfuncid),
		Resulttype:   pg.Oid(n.Resulttype),
		Resulttypmod: n.Resulttypmod,
		Resultcollid: pg.Oid(n.Resultcollid),
		IsExplicit:   n.IsExplicit,
		Coerceformat: pg.CoercionForm(n.Coerceformat),
		Location:     n.Location,
	}
}

func convertAlterDatabaseStmt(n *nodes.AlterDatabaseStmt) *pg.AlterDatabaseStmt {
	return &pg.AlterDatabaseStmt{
		Dbname:  n.Dbname,
		Options: convertList(n.Options),
	}
}

func convertDefElem(n *nodes.DefElem) *pg.DefElem {
	return &pg.DefElem{
		Defnamespace: n.Defnamespace,
		Defname:      n.Defname,
		Arg:          convertNode(n.Arg),
		Defaction:    pg.DefElemAction(n.Defaction),
		Location:     n.Location,
	}
}

func convertTypeCast(n *nodes.TypeCast) *pg.TypeCast {
	return &pg.TypeCast{
		Arg:      convertNode(n.Arg),
		TypeName: convertTypeName(n.TypeName),
		Location: n.Location,
	}
}

func convertClusterStmt(n *nodes.ClusterStmt) *pg.ClusterStmt {
	return &pg.ClusterStmt{
		Relation:  convertRangeVar(n.Relation),
		Indexname: n.Indexname,
		Verbose:   n.Verbose,
	}
}

func convertInferenceElem(n *nodes.InferenceElem) *pg.InferenceElem {
	return &pg.InferenceElem{
		Xpr:          convertNode(n.Xpr),
		Expr:         convertNode(n.Expr),
		Infercollid:  pg.Oid(n.Infercollid),
		Inferopclass: pg.Oid(n.Inferopclass),
	}
}

func convertExplainStmt(n *nodes.ExplainStmt) *pg.ExplainStmt {
	return &pg.ExplainStmt{
		Query:   convertNode(n.Query),
		Options: convertList(n.Options),
	}
}

func convertNamedArgExpr(n *nodes.NamedArgExpr) *pg.NamedArgExpr {
	return &pg.NamedArgExpr{
		Xpr:       convertNode(n.Xpr),
		Arg:       convertNode(n.Arg),
		Name:      n.Name,
		Argnumber: n.Argnumber,
		Location:  n.Location,
	}
}

func convertDefineStmt(n *nodes.DefineStmt) *pg.DefineStmt {
	return &pg.DefineStmt{
		Kind:        pg.ObjectType(n.Kind),
		Oldstyle:    n.Oldstyle,
		Defnames:    convertList(n.Defnames),
		Args:        convertList(n.Args),
		Definition:  convertList(n.Definition),
		IfNotExists: n.IfNotExists,
	}
}

func convertSubLink(n *nodes.SubLink) *pg.SubLink {
	return &pg.SubLink{
		Xpr:         convertNode(n.Xpr),
		SubLinkType: pg.SubLinkType(n.SubLinkType),
		SubLinkId:   n.SubLinkId,
		Testexpr:    convertNode(n.Testexpr),
		OperName:    convertList(n.OperName),
		Subselect:   convertNode(n.Subselect),
		Location:    n.Location,
	}
}

func convertWindowFunc(n *nodes.WindowFunc) *pg.WindowFunc {
	return &pg.WindowFunc{
		Xpr:         convertNode(n.Xpr),
		Winfnoid:    pg.Oid(n.Winfnoid),
		Wintype:     pg.Oid(n.Wintype),
		Wincollid:   pg.Oid(n.Wincollid),
		Inputcollid: pg.Oid(n.Inputcollid),
		Args:        convertList(n.Args),
		Aggfilter:   convertNode(n.Aggfilter),
		Winref:      pg.Index(n.Winref),
		Winstar:     n.Winstar,
		Winagg:      n.Winagg,
		Location:    n.Location,
	}
}

func convertString(n *nodes.String) *pg.String {
	return &pg.String{
		Str: n.Str,
	}
}

func convertObjectWithArgs(n *nodes.ObjectWithArgs) *pg.ObjectWithArgs {
	return &pg.ObjectWithArgs{
		Objname:         convertList(n.Objname),
		Objargs:         convertList(n.Objargs),
		ArgsUnspecified: n.ArgsUnspecified,
	}
}

func convertNullTest(n *nodes.NullTest) *pg.NullTest {
	return &pg.NullTest{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Nulltesttype: pg.NullTestType(n.Nulltesttype),
		Argisrow:     n.Argisrow,
		Location:     n.Location,
	}
}

func convertIndexElem(n *nodes.IndexElem) *pg.IndexElem {
	return &pg.IndexElem{
		Name:          n.Name,
		Expr:          convertNode(n.Expr),
		Indexcolname:  n.Indexcolname,
		Collation:     convertList(n.Collation),
		Opclass:       convertList(n.Opclass),
		Ordering:      pg.SortByDir(n.Ordering),
		NullsOrdering: pg.SortByNulls(n.NullsOrdering),
	}
}

func convertTableFunc(n *nodes.TableFunc) *pg.TableFunc {
	return &pg.TableFunc{
		NsUris:        convertList(n.NsUris),
		NsNames:       convertList(n.NsNames),
		Docexpr:       convertNode(n.Docexpr),
		Rowexpr:       convertNode(n.Rowexpr),
		Colnames:      convertList(n.Colnames),
		Coltypes:      convertList(n.Coltypes),
		Coltypmods:    convertList(n.Coltypmods),
		Colcollations: convertList(n.Colcollations),
		Colexprs:      convertList(n.Colexprs),
		Coldefexprs:   convertList(n.Coldefexprs),
		Notnulls:      n.Notnulls,
		Ordinalitycol: n.Ordinalitycol,
		Location:      n.Location,
	}
}

func convertCreateSeqStmt(n *nodes.CreateSeqStmt) *pg.CreateSeqStmt {
	return &pg.CreateSeqStmt{
		Sequence:    convertRangeVar(n.Sequence),
		Options:     convertList(n.Options),
		OwnerId:     pg.Oid(n.OwnerId),
		ForIdentity: n.ForIdentity,
		IfNotExists: n.IfNotExists,
	}
}

func convertCreateStatsStmt(n *nodes.CreateStatsStmt) *pg.CreateStatsStmt {
	return &pg.CreateStatsStmt{
		Defnames:    convertList(n.Defnames),
		StatTypes:   convertList(n.StatTypes),
		Exprs:       convertList(n.Exprs),
		Relations:   convertList(n.Relations),
		IfNotExists: n.IfNotExists,
	}
}

func convertCreateRangeStmt(n *nodes.CreateRangeStmt) *pg.CreateRangeStmt {
	return &pg.CreateRangeStmt{
		TypeName: convertList(n.TypeName),
		Params:   convertList(n.Params),
	}
}

func convertDropStmt(n *nodes.DropStmt) *pg.DropStmt {
	return &pg.DropStmt{
		Objects:    convertList(n.Objects),
		RemoveType: pg.ObjectType(n.RemoveType),
		Behavior:   pg.DropBehavior(n.Behavior),
		MissingOk:  n.MissingOk,
		Concurrent: n.Concurrent,
	}
}

func convertCommentStmt(n *nodes.CommentStmt) *pg.CommentStmt {
	return &pg.CommentStmt{
		Objtype: pg.ObjectType(n.Objtype),
		Object:  convertNode(n.Object),
		Comment: n.Comment,
	}
}

func convertCreateSubscriptionStmt(n *nodes.CreateSubscriptionStmt) *pg.CreateSubscriptionStmt {
	return &pg.CreateSubscriptionStmt{
		Subname:     n.Subname,
		Conninfo:    n.Conninfo,
		Publication: convertList(n.Publication),
		Options:     convertList(n.Options),
	}
}

func convertCreatePolicyStmt(n *nodes.CreatePolicyStmt) *pg.CreatePolicyStmt {
	return &pg.CreatePolicyStmt{
		PolicyName: n.PolicyName,
		Table:      convertRangeVar(n.Table),
		CmdName:    n.CmdName,
		Permissive: n.Permissive,
		Roles:      convertList(n.Roles),
		Qual:       convertNode(n.Qual),
		WithCheck:  convertNode(n.WithCheck),
	}
}

func convertAlterTableMoveAllStmt(n *nodes.AlterTableMoveAllStmt) *pg.AlterTableMoveAllStmt {
	return &pg.AlterTableMoveAllStmt{
		OrigTablespacename: n.OrigTablespacename,
		Objtype:            pg.ObjectType(n.Objtype),
		Roles:              convertList(n.Roles),
		NewTablespacename:  n.NewTablespacename,
		Nowait:             n.Nowait,
	}
}

func convertAlterObjectDependsStmt(n *nodes.AlterObjectDependsStmt) *pg.AlterObjectDependsStmt {
	return &pg.AlterObjectDependsStmt{
		ObjectType: pg.ObjectType(n.ObjectType),
		Relation:   convertRangeVar(n.Relation),
		Object:     convertNode(n.Object),
		Extname:    convertNode(n.Extname),
	}
}

func convertWindowClause(n *nodes.WindowClause) *pg.WindowClause {
	return &pg.WindowClause{
		Name:            n.Name,
		Refname:         n.Refname,
		PartitionClause: convertList(n.PartitionClause),
		OrderClause:     convertList(n.OrderClause),
		FrameOptions:    n.FrameOptions,
		StartOffset:     convertNode(n.StartOffset),
		EndOffset:       convertNode(n.EndOffset),
		Winref:          pg.Index(n.Winref),
		CopiedOrder:     n.CopiedOrder,
	}
}

func convertColumnDef(n *nodes.ColumnDef) *pg.ColumnDef {
	return &pg.ColumnDef{
		Colname:       n.Colname,
		TypeName:      convertTypeName(n.TypeName),
		Inhcount:      n.Inhcount,
		IsLocal:       n.IsLocal,
		IsNotNull:     n.IsNotNull,
		IsFromType:    n.IsFromType,
		IsFromParent:  n.IsFromParent,
		Storage:       n.Storage,
		RawDefault:    convertNode(n.RawDefault),
		CookedDefault: convertNode(n.CookedDefault),
		Identity:      n.Identity,
		CollClause:    convertCollateClause(n.CollClause),
		CollOid:       pg.Oid(n.CollOid),
		Constraints:   convertList(n.Constraints),
		Fdwoptions:    convertList(n.Fdwoptions),
		Location:      n.Location,
	}
}

func convertSelectStmt(n *nodes.SelectStmt) *pg.SelectStmt {
	return &pg.SelectStmt{
		DistinctClause: convertList(n.DistinctClause),
		IntoClause:     convertIntoClause(n.IntoClause),
		TargetList:     convertList(n.TargetList),
		FromClause:     convertList(n.FromClause),
		WhereClause:    convertNode(n.WhereClause),
		GroupClause:    convertList(n.GroupClause),
		HavingClause:   convertNode(n.HavingClause),
		WindowClause:   convertList(n.WindowClause),
		ValuesLists:    convertValuesList(n.ValuesLists),
		SortClause:     convertList(n.SortClause),
		LimitOffset:    convertNode(n.LimitOffset),
		LimitCount:     convertNode(n.LimitCount),
		LockingClause:  convertList(n.LockingClause),
		WithClause:     convertWithClause(n.WithClause),
		Op:             pg.SetOperation(n.Op),
		All:            n.All,
		Larg:           convertSelectStmt(n.Larg),
		Rarg:           convertSelectStmt(n.Rarg),
	}
}

func convertLockStmt(n *nodes.LockStmt) *pg.LockStmt {
	return &pg.LockStmt{
		Relations: convertList(n.Relations),
		Mode:      n.Mode,
		Nowait:    n.Nowait,
	}
}

func convertFieldStore(n *nodes.FieldStore) *pg.FieldStore {
	return &pg.FieldStore{
		Xpr:        convertNode(n.Xpr),
		Arg:        convertNode(n.Arg),
		Newvals:    convertList(n.Newvals),
		Fieldnums:  convertList(n.Fieldnums),
		Resulttype: pg.Oid(n.Resulttype),
	}
}

func convertAlterSubscriptionStmt(n *nodes.AlterSubscriptionStmt) *pg.AlterSubscriptionStmt {
	return &pg.AlterSubscriptionStmt{
		Kind:        pg.AlterSubscriptionType(n.Kind),
		Subname:     n.Subname,
		Conninfo:    n.Conninfo,
		Publication: convertList(n.Publication),
		Options:     convertList(n.Options),
	}
}

func convertAlterFdwStmt(n *nodes.AlterFdwStmt) *pg.AlterFdwStmt {
	return &pg.AlterFdwStmt{
		Fdwname:     n.Fdwname,
		FuncOptions: convertList(n.FuncOptions),
		Options:     convertList(n.Options),
	}
}

func convertInlineCodeBlock(n *nodes.InlineCodeBlock) *pg.InlineCodeBlock {
	return &pg.InlineCodeBlock{
		SourceText:    n.SourceText,
		LangOid:       pg.Oid(n.LangOid),
		LangIsTrusted: n.LangIsTrusted,
	}
}

func convertA_Indices(n *nodes.A_Indices) *pg.A_Indices {
	return &pg.A_Indices{
		IsSlice: n.IsSlice,
		Lidx:    convertNode(n.Lidx),
		Uidx:    convertNode(n.Uidx),
	}
}

func convertAlterTableSpaceOptionsStmt(n *nodes.AlterTableSpaceOptionsStmt) *pg.AlterTableSpaceOptionsStmt {
	return &pg.AlterTableSpaceOptionsStmt{
		Tablespacename: n.Tablespacename,
		Options:        convertList(n.Options),
		IsReset:        n.IsReset,
	}
}

func convertCoerceToDomainValue(n *nodes.CoerceToDomainValue) *pg.CoerceToDomainValue {
	return &pg.CoerceToDomainValue{
		Xpr:       convertNode(n.Xpr),
		TypeId:    pg.Oid(n.TypeId),
		TypeMod:   n.TypeMod,
		Collation: pg.Oid(n.Collation),
		Location:  n.Location,
	}
}

func convertBooleanTest(n *nodes.BooleanTest) *pg.BooleanTest {
	return &pg.BooleanTest{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Booltesttype: pg.BoolTestType(n.Booltesttype),
		Location:     n.Location,
	}
}

func convertVariableSetStmt(n *nodes.VariableSetStmt) *pg.VariableSetStmt {
	return &pg.VariableSetStmt{
		Kind:    pg.VariableSetKind(n.Kind),
		Name:    n.Name,
		Args:    convertList(n.Args),
		IsLocal: n.IsLocal,
	}
}

func convertAlterDatabaseSetStmt(n *nodes.AlterDatabaseSetStmt) *pg.AlterDatabaseSetStmt {
	return &pg.AlterDatabaseSetStmt{
		Dbname:  n.Dbname,
		Setstmt: convertVariableSetStmt(n.Setstmt),
	}
}

func convertUpdateStmt(n *nodes.UpdateStmt) *pg.UpdateStmt {
	return &pg.UpdateStmt{
		Relation:      convertRangeVar(n.Relation),
		TargetList:    convertList(n.TargetList),
		WhereClause:   convertNode(n.WhereClause),
		FromClause:    convertList(n.FromClause),
		ReturningList: convertList(n.ReturningList),
		WithClause:    convertWithClause(n.WithClause),
	}
}

func convertA_Const(n *nodes.A_Const) *pg.A_Const {
	return &pg.A_Const{
		Val:      convertNode(n.Val),
		Location: n.Location,
	}
}

func convertDropSubscriptionStmt(n *nodes.DropSubscriptionStmt) *pg.DropSubscriptionStmt {
	return &pg.DropSubscriptionStmt{
		Subname:   n.Subname,
		MissingOk: n.MissingOk,
		Behavior:  pg.DropBehavior(n.Behavior),
	}
}

func convertA_ArrayExpr(n *nodes.A_ArrayExpr) *pg.A_ArrayExpr {
	return &pg.A_ArrayExpr{
		Elements: convertList(n.Elements),
		Location: n.Location,
	}
}

func convertBoolExpr(n *nodes.BoolExpr) *pg.BoolExpr {
	return &pg.BoolExpr{
		Xpr:      convertNode(n.Xpr),
		Boolop:   pg.BoolExprType(n.Boolop),
		Args:     convertList(n.Args),
		Location: n.Location,
	}
}

func convertA_Star(n *nodes.A_Star) *pg.A_Star {
	return &pg.A_Star{}
}

func convertAlterTSDictionaryStmt(n *nodes.AlterTSDictionaryStmt) *pg.AlterTSDictionaryStmt {
	return &pg.AlterTSDictionaryStmt{
		Dictname: convertList(n.Dictname),
		Options:  convertList(n.Options),
	}
}

func convertConst(n *nodes.Const) *pg.Const {
	return &pg.Const{
		Xpr:         convertNode(n.Xpr),
		Consttype:   pg.Oid(n.Consttype),
		Consttypmod: n.Consttypmod,
		Constcollid: pg.Oid(n.Constcollid),
		Constlen:    n.Constlen,
		Constvalue:  pg.Datum(n.Constvalue),
		Constisnull: n.Constisnull,
		Constbyval:  n.Constbyval,
		Location:    n.Location,
	}
}

func convertCommonTableExpr(n *nodes.CommonTableExpr) *pg.CommonTableExpr {
	return &pg.CommonTableExpr{
		Ctename:          n.Ctename,
		Aliascolnames:    convertList(n.Aliascolnames),
		Ctequery:         convertNode(n.Ctequery),
		Location:         n.Location,
		Cterecursive:     n.Cterecursive,
		Cterefcount:      n.Cterefcount,
		Ctecolnames:      convertList(n.Ctecolnames),
		Ctecoltypes:      convertList(n.Ctecoltypes),
		Ctecoltypmods:    convertList(n.Ctecoltypmods),
		Ctecolcollations: convertList(n.Ctecolcollations),
	}
}

func convertRenameStmt(n *nodes.RenameStmt) *pg.RenameStmt {
	return &pg.RenameStmt{
		RenameType:   pg.ObjectType(n.RenameType),
		RelationType: pg.ObjectType(n.RelationType),
		Relation:     convertRangeVar(n.Relation),
		Object:       convertNode(n.Object),
		Subname:      n.Subname,
		Newname:      n.Newname,
		Behavior:     pg.DropBehavior(n.Behavior),
		MissingOk:    n.MissingOk,
	}
}

func convertCurrentOfExpr(n *nodes.CurrentOfExpr) *pg.CurrentOfExpr {
	return &pg.CurrentOfExpr{
		Xpr:         convertNode(n.Xpr),
		Cvarno:      pg.Index(n.Cvarno),
		CursorName:  n.CursorName,
		CursorParam: n.CursorParam,
	}
}

func convertA_Expr(n *nodes.A_Expr) *pg.A_Expr {
	return &pg.A_Expr{
		Kind:     pg.A_Expr_Kind(n.Kind),
		Name:     convertList(n.Name),
		Lexpr:    convertNode(n.Lexpr),
		Rexpr:    convertNode(n.Rexpr),
		Location: n.Location,
	}
}

func convertIndexStmt(n *nodes.IndexStmt) *pg.IndexStmt {
	return &pg.IndexStmt{
		Idxname:        n.Idxname,
		Relation:       convertRangeVar(n.Relation),
		AccessMethod:   n.AccessMethod,
		TableSpace:     n.TableSpace,
		IndexParams:    convertList(n.IndexParams),
		Options:        convertList(n.Options),
		WhereClause:    convertNode(n.WhereClause),
		ExcludeOpNames: convertList(n.ExcludeOpNames),
		Idxcomment:     n.Idxcomment,
		IndexOid:       pg.Oid(n.IndexOid),
		OldNode:        pg.Oid(n.OldNode),
		Unique:         n.Unique,
		Primary:        n.Primary,
		Isconstraint:   n.Isconstraint,
		Deferrable:     n.Deferrable,
		Initdeferred:   n.Initdeferred,
		Transformed:    n.Transformed,
		Concurrent:     n.Concurrent,
		IfNotExists:    n.IfNotExists,
	}
}

func convertPartitionElem(n *nodes.PartitionElem) *pg.PartitionElem {
	return &pg.PartitionElem{
		Name:      n.Name,
		Expr:      convertNode(n.Expr),
		Collation: convertList(n.Collation),
		Opclass:   convertList(n.Opclass),
		Location:  n.Location,
	}
}

func convertCreateTableAsStmt(n *nodes.CreateTableAsStmt) *pg.CreateTableAsStmt {
	return &pg.CreateTableAsStmt{
		Query:        convertNode(n.Query),
		Into:         convertIntoClause(n.Into),
		Relkind:      pg.ObjectType(n.Relkind),
		IsSelectInto: n.IsSelectInto,
		IfNotExists:  n.IfNotExists,
	}
}

func convertRowMarkClause(n *nodes.RowMarkClause) *pg.RowMarkClause {
	return &pg.RowMarkClause{
		Rti:        pg.Index(n.Rti),
		Strength:   pg.LockClauseStrength(n.Strength),
		WaitPolicy: pg.LockWaitPolicy(n.WaitPolicy),
		PushedDown: n.PushedDown,
	}
}

func convertAlterObjectSchemaStmt(n *nodes.AlterObjectSchemaStmt) *pg.AlterObjectSchemaStmt {
	return &pg.AlterObjectSchemaStmt{
		ObjectType: pg.ObjectType(n.ObjectType),
		Relation:   convertRangeVar(n.Relation),
		Object:     convertNode(n.Object),
		Newschema:  n.Newschema,
		MissingOk:  n.MissingOk,
	}
}

func convertSQLValueFunction(n *nodes.SQLValueFunction) *pg.SQLValueFunction {
	return &pg.SQLValueFunction{
		Xpr:      convertNode(n.Xpr),
		Op:       pg.SQLValueFunctionOp(n.Op),
		Type:     pg.Oid(n.Type),
		Typmod:   n.Typmod,
		Location: n.Location,
	}
}

func convertVacuumStmt(n *nodes.VacuumStmt) *pg.VacuumStmt {
	return &pg.VacuumStmt{
		Options:  n.Options,
		Relation: convertRangeVar(n.Relation),
		VaCols:   convertList(n.VaCols),
	}
}

func convertCollateExpr(n *nodes.CollateExpr) *pg.CollateExpr {
	return &pg.CollateExpr{
		Xpr:      convertNode(n.Xpr),
		Arg:      convertNode(n.Arg),
		CollOid:  pg.Oid(n.CollOid),
		Location: n.Location,
	}
}

func convertConstraintsSetStmt(n *nodes.ConstraintsSetStmt) *pg.ConstraintsSetStmt {
	return &pg.ConstraintsSetStmt{
		Constraints: convertList(n.Constraints),
		Deferred:    n.Deferred,
	}
}

func convertNextValueExpr(n *nodes.NextValueExpr) *pg.NextValueExpr {
	return &pg.NextValueExpr{
		Xpr:    convertNode(n.Xpr),
		Seqid:  pg.Oid(n.Seqid),
		TypeId: pg.Oid(n.TypeId),
	}
}

func convertGroupingSet(n *nodes.GroupingSet) *pg.GroupingSet {
	return &pg.GroupingSet{
		Kind:     pg.GroupingSetKind(n.Kind),
		Content:  convertList(n.Content),
		Location: n.Location,
	}
}

func convertRangeVar(n *nodes.RangeVar) *pg.RangeVar {
	return &pg.RangeVar{
		Catalogname:    n.Catalogname,
		Schemaname:     n.Schemaname,
		Relname:        n.Relname,
		Inh:            n.Inh,
		Relpersistence: n.Relpersistence,
		Alias:          convertAlias(n.Alias),
		Location:       n.Location,
	}
}

func convertMinMaxExpr(n *nodes.MinMaxExpr) *pg.MinMaxExpr {
	return &pg.MinMaxExpr{
		Xpr:          convertNode(n.Xpr),
		Minmaxtype:   pg.Oid(n.Minmaxtype),
		Minmaxcollid: pg.Oid(n.Minmaxcollid),
		Inputcollid:  pg.Oid(n.Inputcollid),
		Op:           pg.MinMaxOp(n.Op),
		Args:         convertList(n.Args),
		Location:     n.Location,
	}
}

func convertArrayExpr(n *nodes.ArrayExpr) *pg.ArrayExpr {
	return &pg.ArrayExpr{
		Xpr:           convertNode(n.Xpr),
		ArrayTypeid:   pg.Oid(n.ArrayTypeid),
		ArrayCollid:   pg.Oid(n.ArrayCollid),
		ElementTypeid: pg.Oid(n.ElementTypeid),
		Elements:      convertList(n.Elements),
		Multidims:     n.Multidims,
		Location:      n.Location,
	}
}

func convertFieldSelect(n *nodes.FieldSelect) *pg.FieldSelect {
	return &pg.FieldSelect{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Fieldnum:     pg.AttrNumber(n.Fieldnum),
		Resulttype:   pg.Oid(n.Resulttype),
		Resulttypmod: n.Resulttypmod,
		Resultcollid: pg.Oid(n.Resultcollid),
	}
}

func convertQuery(n *nodes.Query) *pg.Query {
	return &pg.Query{
		CommandType:      pg.CmdType(n.CommandType),
		QuerySource:      pg.QuerySource(n.QuerySource),
		QueryId:          n.QueryId,
		CanSetTag:        n.CanSetTag,
		UtilityStmt:      convertNode(n.UtilityStmt),
		ResultRelation:   n.ResultRelation,
		HasAggs:          n.HasAggs,
		HasWindowFuncs:   n.HasWindowFuncs,
		HasTargetSrfs:    n.HasTargetSrfs,
		HasSubLinks:      n.HasSubLinks,
		HasDistinctOn:    n.HasDistinctOn,
		HasRecursive:     n.HasRecursive,
		HasModifyingCte:  n.HasModifyingCte,
		HasForUpdate:     n.HasForUpdate,
		HasRowSecurity:   n.HasRowSecurity,
		CteList:          convertList(n.CteList),
		Rtable:           convertList(n.Rtable),
		Jointree:         convertFromExpr(n.Jointree),
		TargetList:       convertList(n.TargetList),
		Override:         pg.OverridingKind(n.Override),
		OnConflict:       convertOnConflictExpr(n.OnConflict),
		ReturningList:    convertList(n.ReturningList),
		GroupClause:      convertList(n.GroupClause),
		GroupingSets:     convertList(n.GroupingSets),
		HavingQual:       convertNode(n.HavingQual),
		WindowClause:     convertList(n.WindowClause),
		DistinctClause:   convertList(n.DistinctClause),
		SortClause:       convertList(n.SortClause),
		LimitOffset:      convertNode(n.LimitOffset),
		LimitCount:       convertNode(n.LimitCount),
		RowMarks:         convertList(n.RowMarks),
		SetOperations:    convertNode(n.SetOperations),
		ConstraintDeps:   convertList(n.ConstraintDeps),
		WithCheckOptions: convertList(n.WithCheckOptions),
		StmtLocation:     n.StmtLocation,
		StmtLen:          n.StmtLen,
	}
}

func convertNotifyStmt(n *nodes.NotifyStmt) *pg.NotifyStmt {
	return &pg.NotifyStmt{
		Conditionname: n.Conditionname,
		Payload:       n.Payload,
	}
}

func convertFuncCall(n *nodes.FuncCall) *pg.FuncCall {
	return &pg.FuncCall{
		Funcname:       convertList(n.Funcname),
		Args:           convertList(n.Args),
		AggOrder:       convertList(n.AggOrder),
		AggFilter:      convertNode(n.AggFilter),
		AggWithinGroup: n.AggWithinGroup,
		AggStar:        n.AggStar,
		AggDistinct:    n.AggDistinct,
		FuncVariadic:   n.FuncVariadic,
		Over:           convertWindowDef(n.Over),
		Location:       n.Location,
	}
}

func convertCoalesceExpr(n *nodes.CoalesceExpr) *pg.CoalesceExpr {
	return &pg.CoalesceExpr{
		Xpr:            convertNode(n.Xpr),
		Coalescetype:   pg.Oid(n.Coalescetype),
		Coalescecollid: pg.Oid(n.Coalescecollid),
		Args:           convertList(n.Args),
		Location:       n.Location,
	}
}

func convertExecuteStmt(n *nodes.ExecuteStmt) *pg.ExecuteStmt {
	return &pg.ExecuteStmt{
		Name:   n.Name,
		Params: convertList(n.Params),
	}
}

func convertCreateForeignTableStmt(n *nodes.CreateForeignTableStmt) *pg.CreateForeignTableStmt {
	return &pg.CreateForeignTableStmt{
		Base:       convertCreateStmt(&n.Base),
		Servername: n.Servername,
		Options:    convertList(n.Options),
	}
}

func convertCreateStmt(n *nodes.CreateStmt) *pg.CreateStmt {
	return &pg.CreateStmt{
		Relation:       convertRangeVar(n.Relation),
		TableElts:      convertList(n.TableElts),
		InhRelations:   convertList(n.InhRelations),
		Partbound:      convertPartitionBoundSpec(n.Partbound),
		Partspec:       convertPartitionSpec(n.Partspec),
		OfTypename:     convertTypeName(n.OfTypename),
		Constraints:    convertList(n.Constraints),
		Options:        convertList(n.Options),
		Oncommit:       pg.OnCommitAction(n.Oncommit),
		Tablespacename: n.Tablespacename,
		IfNotExists:    n.IfNotExists,
	}
}

func convertCreateOpClassStmt(n *nodes.CreateOpClassStmt) *pg.CreateOpClassStmt {
	return &pg.CreateOpClassStmt{
		Opclassname:  convertList(n.Opclassname),
		Opfamilyname: convertList(n.Opfamilyname),
		Amname:       n.Amname,
		Datatype:     convertTypeName(n.Datatype),
		Items:        convertList(n.Items),
		IsDefault:    n.IsDefault,
	}
}

func convertFromExpr(n *nodes.FromExpr) *pg.FromExpr {
	return &pg.FromExpr{
		Fromlist: convertList(n.Fromlist),
		Quals:    convertNode(n.Quals),
	}
}

func convertInsertStmt(n *nodes.InsertStmt) *pg.InsertStmt {
	return &pg.InsertStmt{
		Relation:         convertRangeVar(n.Relation),
		Cols:             convertList(n.Cols),
		SelectStmt:       convertNode(n.SelectStmt),
		OnConflictClause: convertOnConflictClause(n.OnConflictClause),
		ReturningList:    convertList(n.ReturningList),
		WithClause:       convertWithClause(n.WithClause),
		Override:         pg.OverridingKind(n.Override),
	}
}

func convertColumnRef(n *nodes.ColumnRef) *pg.ColumnRef {
	return &pg.ColumnRef{
		Fields:   convertList(n.Fields),
		Location: n.Location,
	}
}

func convertCreateCastStmt(n *nodes.CreateCastStmt) *pg.CreateCastStmt {
	return &pg.CreateCastStmt{
		Sourcetype: convertTypeName(n.Sourcetype),
		Targettype: convertTypeName(n.Targettype),
		Func:       convertObjectWithArgs(n.Func),
		Context:    pg.CoercionContext(n.Context),
		Inout:      n.Inout,
	}
}

func convertCoerceToDomain(n *nodes.CoerceToDomain) *pg.CoerceToDomain {
	return &pg.CoerceToDomain{
		Xpr:            convertNode(n.Xpr),
		Arg:            convertNode(n.Arg),
		Resulttype:     pg.Oid(n.Resulttype),
		Resulttypmod:   n.Resulttypmod,
		Resultcollid:   pg.Oid(n.Resultcollid),
		Coercionformat: pg.CoercionForm(n.Coercionformat),
		Location:       n.Location,
	}
}

func convertAlterEnumStmt(n *nodes.AlterEnumStmt) *pg.AlterEnumStmt {
	return &pg.AlterEnumStmt{
		TypeName:           convertList(n.TypeName),
		OldVal:             n.OldVal,
		NewVal:             n.NewVal,
		NewValNeighbor:     n.NewValNeighbor,
		NewValIsAfter:      n.NewValIsAfter,
		SkipIfNewValExists: n.SkipIfNewValExists,
	}
}

func convertRawStmt(n *nodes.RawStmt) *pg.RawStmt {
	return &pg.RawStmt{
		Stmt:         convertNode(n.Stmt),
		StmtLocation: n.StmtLocation,
		StmtLen:      n.StmtLen,
	}
}

func convertCreateDomainStmt(n *nodes.CreateDomainStmt) *pg.CreateDomainStmt {
	return &pg.CreateDomainStmt{
		Domainname:  convertList(n.Domainname),
		TypeName:    convertTypeName(n.TypeName),
		CollClause:  convertCollateClause(n.CollClause),
		Constraints: convertList(n.Constraints),
	}
}

func convertParamRef(n *nodes.ParamRef) *pg.ParamRef {
	return &pg.ParamRef{
		Number:   n.Number,
		Location: n.Location,
	}
}

func convertFloat(n *nodes.Float) *pg.Float {
	return &pg.Float{
		Str: n.Str,
	}
}

func convertMultiAssignRef(n *nodes.MultiAssignRef) *pg.MultiAssignRef {
	return &pg.MultiAssignRef{
		Source:   convertNode(n.Source),
		Colno:    n.Colno,
		Ncolumns: n.Ncolumns,
	}
}

func convertTriggerTransition(n *nodes.TriggerTransition) *pg.TriggerTransition {
	return &pg.TriggerTransition{
		Name:    n.Name,
		IsNew:   n.IsNew,
		IsTable: n.IsTable,
	}
}

func convertArrayRef(n *nodes.ArrayRef) *pg.ArrayRef {
	return &pg.ArrayRef{
		Xpr:             convertNode(n.Xpr),
		Refarraytype:    pg.Oid(n.Refarraytype),
		Refelemtype:     pg.Oid(n.Refelemtype),
		Reftypmod:       n.Reftypmod,
		Refcollid:       pg.Oid(n.Refcollid),
		Refupperindexpr: convertList(n.Refupperindexpr),
		Reflowerindexpr: convertList(n.Reflowerindexpr),
		Refexpr:         convertNode(n.Refexpr),
		Refassgnexpr:    convertNode(n.Refassgnexpr),
	}
}

func convertImportForeignSchemaStmt(n *nodes.ImportForeignSchemaStmt) *pg.ImportForeignSchemaStmt {
	return &pg.ImportForeignSchemaStmt{
		ServerName:   n.ServerName,
		RemoteSchema: n.RemoteSchema,
		LocalSchema:  n.LocalSchema,
		ListType:     pg.ImportForeignSchemaType(n.ListType),
		TableList:    convertList(n.TableList),
		Options:      convertList(n.Options),
	}
}

func convertAlterOpFamilyStmt(n *nodes.AlterOpFamilyStmt) *pg.AlterOpFamilyStmt {
	return &pg.AlterOpFamilyStmt{
		Opfamilyname: convertList(n.Opfamilyname),
		Amname:       n.Amname,
		IsDrop:       n.IsDrop,
		Items:        convertList(n.Items),
	}
}

func convertAlterOwnerStmt(n *nodes.AlterOwnerStmt) *pg.AlterOwnerStmt {
	return &pg.AlterOwnerStmt{
		ObjectType: pg.ObjectType(n.ObjectType),
		Relation:   convertRangeVar(n.Relation),
		Object:     convertNode(n.Object),
		Newowner:   convertRoleSpec(n.Newowner),
	}
}

func convertOpExpr(n *nodes.OpExpr) *pg.OpExpr {
	return &pg.OpExpr{
		Xpr:          convertNode(n.Xpr),
		Opno:         pg.Oid(n.Opno),
		Opfuncid:     pg.Oid(n.Opfuncid),
		Opresulttype: pg.Oid(n.Opresulttype),
		Opretset:     n.Opretset,
		Opcollid:     pg.Oid(n.Opcollid),
		Inputcollid:  pg.Oid(n.Inputcollid),
		Args:         convertList(n.Args),
		Location:     n.Location,
	}
}

func convertAlterRoleStmt(n *nodes.AlterRoleStmt) *pg.AlterRoleStmt {
	return &pg.AlterRoleStmt{
		Role:    convertRoleSpec(n.Role),
		Options: convertList(n.Options),
		Action:  n.Action,
	}
}

func convertRangeTblRef(n *nodes.RangeTblRef) *pg.RangeTblRef {
	return &pg.RangeTblRef{
		Rtindex: n.Rtindex,
	}
}

func convertCreateTransformStmt(n *nodes.CreateTransformStmt) *pg.CreateTransformStmt {
	return &pg.CreateTransformStmt{
		Replace:  n.Replace,
		TypeName: convertTypeName(n.TypeName),
		Lang:     n.Lang,
		Fromsql:  convertObjectWithArgs(n.Fromsql),
		Tosql:    convertObjectWithArgs(n.Tosql),
	}
}

func convertConvertRowtypeExpr(n *nodes.ConvertRowtypeExpr) *pg.ConvertRowtypeExpr {
	return &pg.ConvertRowtypeExpr{
		Xpr:           convertNode(n.Xpr),
		Arg:           convertNode(n.Arg),
		Resulttype:    pg.Oid(n.Resulttype),
		Convertformat: pg.CoercionForm(n.Convertformat),
		Location:      n.Location,
	}
}

func convertCompositeTypeStmt(n *nodes.CompositeTypeStmt) *pg.CompositeTypeStmt {
	return &pg.CompositeTypeStmt{
		Typevar:    convertRangeVar(n.Typevar),
		Coldeflist: convertList(n.Coldeflist),
	}
}

func convertRangeFunction(n *nodes.RangeFunction) *pg.RangeFunction {
	return &pg.RangeFunction{
		Lateral:    n.Lateral,
		Ordinality: n.Ordinality,
		IsRowsfrom: n.IsRowsfrom,
		Functions:  convertList(n.Functions),
		Alias:      convertAlias(n.Alias),
		Coldeflist: convertList(n.Coldeflist),
	}
}

func convertCreateForeignServerStmt(n *nodes.CreateForeignServerStmt) *pg.CreateForeignServerStmt {
	return &pg.CreateForeignServerStmt{
		Servername:  n.Servername,
		Servertype:  n.Servertype,
		Version:     n.Version,
		Fdwname:     n.Fdwname,
		IfNotExists: n.IfNotExists,
		Options:     convertList(n.Options),
	}
}

func convertAlterCollationStmt(n *nodes.AlterCollationStmt) *pg.AlterCollationStmt {
	return &pg.AlterCollationStmt{
		Collname: convertList(n.Collname),
	}
}

func convertAlterEventTrigStmt(n *nodes.AlterEventTrigStmt) *pg.AlterEventTrigStmt {
	return &pg.AlterEventTrigStmt{
		Trigname:  n.Trigname,
		Tgenabled: n.Tgenabled,
	}
}

func convertPrepareStmt(n *nodes.PrepareStmt) *pg.PrepareStmt {
	return &pg.PrepareStmt{
		Name:     n.Name,
		Argtypes: convertList(n.Argtypes),
		Query:    convertNode(n.Query),
	}
}

func convertCreateConversionStmt(n *nodes.CreateConversionStmt) *pg.CreateConversionStmt {
	return &pg.CreateConversionStmt{
		ConversionName:  convertList(n.ConversionName),
		ForEncodingName: n.ForEncodingName,
		ToEncodingName:  n.ToEncodingName,
		FuncName:        convertList(n.FuncName),
		Def:             n.Def,
	}
}

func convertAlterTableStmt(n *nodes.AlterTableStmt) *pg.AlterTableStmt {
	return &pg.AlterTableStmt{
		Relation:  convertRangeVar(n.Relation),
		Cmds:      convertList(n.Cmds),
		Relkind:   pg.ObjectType(n.Relkind),
		MissingOk: n.MissingOk,
	}
}

func convertCreatePublicationStmt(n *nodes.CreatePublicationStmt) *pg.CreatePublicationStmt {
	return &pg.CreatePublicationStmt{
		Pubname:      n.Pubname,
		Options:      convertList(n.Options),
		Tables:       convertList(n.Tables),
		ForAllTables: n.ForAllTables,
	}
}

func convertDeallocateStmt(n *nodes.DeallocateStmt) *pg.DeallocateStmt {
	return &pg.DeallocateStmt{
		Name: n.Name,
	}
}

func convertPartitionSpec(n *nodes.PartitionSpec) *pg.PartitionSpec {
	return &pg.PartitionSpec{
		Strategy:   n.Strategy,
		PartParams: convertList(n.PartParams),
		Location:   n.Location,
	}
}

func convertOnConflictClause(n *nodes.OnConflictClause) *pg.OnConflictClause {
	return &pg.OnConflictClause{
		Action:      pg.OnConflictAction(n.Action),
		Infer:       convertInferClause(n.Infer),
		TargetList:  convertList(n.TargetList),
		WhereClause: convertNode(n.WhereClause),
		Location:    n.Location,
	}
}

func convertDropOwnedStmt(n *nodes.DropOwnedStmt) *pg.DropOwnedStmt {
	return &pg.DropOwnedStmt{
		Roles:    convertList(n.Roles),
		Behavior: pg.DropBehavior(n.Behavior),
	}
}

func convertSetToDefault(n *nodes.SetToDefault) *pg.SetToDefault {
	return &pg.SetToDefault{
		Xpr:       convertNode(n.Xpr),
		TypeId:    pg.Oid(n.TypeId),
		TypeMod:   n.TypeMod,
		Collation: pg.Oid(n.Collation),
		Location:  n.Location,
	}
}

func convertSortBy(n *nodes.SortBy) *pg.SortBy {
	return &pg.SortBy{
		Node:        convertNode(n.Node),
		SortbyDir:   pg.SortByDir(n.SortbyDir),
		SortbyNulls: pg.SortByNulls(n.SortbyNulls),
		UseOp:       convertList(n.UseOp),
		Location:    n.Location,
	}
}

func convertAlternativeSubPlan(n *nodes.AlternativeSubPlan) *pg.AlternativeSubPlan {
	return &pg.AlternativeSubPlan{
		Xpr:      convertNode(n.Xpr),
		Subplans: convertList(n.Subplans),
	}
}

func convertVariableShowStmt(n *nodes.VariableShowStmt) *pg.VariableShowStmt {
	return &pg.VariableShowStmt{
		Name: n.Name,
	}
}

func convertInferClause(n *nodes.InferClause) *pg.InferClause {
	return &pg.InferClause{
		IndexElems:  convertList(n.IndexElems),
		WhereClause: convertNode(n.WhereClause),
		Conname:     n.Conname,
		Location:    n.Location,
	}
}

func convertA_Indirection(n *nodes.A_Indirection) *pg.A_Indirection {
	return &pg.A_Indirection{
		Arg:         convertNode(n.Arg),
		Indirection: convertList(n.Indirection),
	}
}

func convertAlterOperatorStmt(n *nodes.AlterOperatorStmt) *pg.AlterOperatorStmt {
	return &pg.AlterOperatorStmt{
		Opername: convertObjectWithArgs(n.Opername),
		Options:  convertList(n.Options),
	}
}

func convertAlterDefaultPrivilegesStmt(n *nodes.AlterDefaultPrivilegesStmt) *pg.AlterDefaultPrivilegesStmt {
	return &pg.AlterDefaultPrivilegesStmt{
		Options: convertList(n.Options),
		Action:  convertGrantStmt(n.Action),
	}
}

func convertFetchStmt(n *nodes.FetchStmt) *pg.FetchStmt {
	return &pg.FetchStmt{
		Direction:  pg.FetchDirection(n.Direction),
		HowMany:    n.HowMany,
		Portalname: n.Portalname,
		Ismove:     n.Ismove,
	}
}

func convertConstraint(n *nodes.Constraint) *pg.Constraint {
	return &pg.Constraint{
		Contype:        pg.ConstrType(n.Contype),
		Conname:        n.Conname,
		Deferrable:     n.Deferrable,
		Initdeferred:   n.Initdeferred,
		Location:       n.Location,
		IsNoInherit:    n.IsNoInherit,
		RawExpr:        convertNode(n.RawExpr),
		CookedExpr:     n.CookedExpr,
		GeneratedWhen:  n.GeneratedWhen,
		Keys:           convertList(n.Keys),
		Exclusions:     convertList(n.Exclusions),
		Options:        convertList(n.Options),
		Indexname:      n.Indexname,
		Indexspace:     n.Indexspace,
		AccessMethod:   n.AccessMethod,
		WhereClause:    convertNode(n.WhereClause),
		Pktable:        convertRangeVar(n.Pktable),
		FkAttrs:        convertList(n.FkAttrs),
		PkAttrs:        convertList(n.PkAttrs),
		FkMatchtype:    n.FkMatchtype,
		FkUpdAction:    n.FkUpdAction,
		FkDelAction:    n.FkDelAction,
		OldConpfeqop:   convertList(n.OldConpfeqop),
		OldPktableOid:  pg.Oid(n.OldPktableOid),
		SkipValidation: n.SkipValidation,
		InitiallyValid: n.InitiallyValid,
	}
}

func convertCopyStmt(n *nodes.CopyStmt) *pg.CopyStmt {
	return &pg.CopyStmt{
		Relation:  convertRangeVar(n.Relation),
		Query:     convertNode(n.Query),
		Attlist:   convertList(n.Attlist),
		IsFrom:    n.IsFrom,
		IsProgram: n.IsProgram,
		Filename:  n.Filename,
		Options:   convertList(n.Options),
	}
}

func convertCreateOpClassItem(n *nodes.CreateOpClassItem) *pg.CreateOpClassItem {
	return &pg.CreateOpClassItem{
		Itemtype:    n.Itemtype,
		Name:        convertObjectWithArgs(n.Name),
		Number:      n.Number,
		OrderFamily: convertList(n.OrderFamily),
		ClassArgs:   convertList(n.ClassArgs),
		Storedtype:  convertTypeName(n.Storedtype),
	}
}

func convertCreateSchemaStmt(n *nodes.CreateSchemaStmt) *pg.CreateSchemaStmt {
	return &pg.CreateSchemaStmt{
		Schemaname:  n.Schemaname,
		Authrole:    convertRoleSpec(n.Authrole),
		SchemaElts:  convertList(n.SchemaElts),
		IfNotExists: n.IfNotExists,
	}
}

func convertWithClause(n *nodes.WithClause) *pg.WithClause {
	return &pg.WithClause{
		Ctes:      convertList(n.Ctes),
		Recursive: n.Recursive,
		Location:  n.Location,
	}
}

func convertCreateFunctionStmt(n *nodes.CreateFunctionStmt) *pg.CreateFunctionStmt {
	return &pg.CreateFunctionStmt{
		Replace:    n.Replace,
		Funcname:   convertList(n.Funcname),
		Parameters: convertList(n.Parameters),
		ReturnType: convertTypeName(n.ReturnType),
		Options:    convertList(n.Options),
		WithClause: convertList(n.WithClause),
	}
}

func convertCheckPointStmt(n *nodes.CheckPointStmt) *pg.CheckPointStmt {
	return &pg.CheckPointStmt{}
}

func convertAlterDomainStmt(n *nodes.AlterDomainStmt) *pg.AlterDomainStmt {
	return &pg.AlterDomainStmt{
		Subtype:   n.Subtype,
		TypeName:  convertList(n.TypeName),
		Name:      n.Name,
		Def:       convertNode(n.Def),
		Behavior:  pg.DropBehavior(n.Behavior),
		MissingOk: n.MissingOk,
	}
}

func convertAlterFunctionStmt(n *nodes.AlterFunctionStmt) *pg.AlterFunctionStmt {
	return &pg.AlterFunctionStmt{
		Func:    convertObjectWithArgs(n.Func),
		Actions: convertList(n.Actions),
	}
}

func convertScalarArrayOpExpr(n *nodes.ScalarArrayOpExpr) *pg.ScalarArrayOpExpr {
	return &pg.ScalarArrayOpExpr{
		Xpr:         convertNode(n.Xpr),
		Opno:        pg.Oid(n.Opno),
		Opfuncid:    pg.Oid(n.Opfuncid),
		UseOr:       n.UseOr,
		Inputcollid: pg.Oid(n.Inputcollid),
		Args:        convertList(n.Args),
		Location:    n.Location,
	}
}

func convertXmlExpr(n *nodes.XmlExpr) *pg.XmlExpr {
	return &pg.XmlExpr{
		Xpr:       convertNode(n.Xpr),
		Op:        pg.XmlExprOp(n.Op),
		Name:      n.Name,
		NamedArgs: convertList(n.NamedArgs),
		ArgNames:  convertList(n.ArgNames),
		Args:      convertList(n.Args),
		Xmloption: pg.XmlOptionType(n.Xmloption),
		Type:      pg.Oid(n.Type),
		Typmod:    n.Typmod,
		Location:  n.Location,
	}
}

func convertParam(n *nodes.Param) *pg.Param {
	return &pg.Param{
		Xpr:         convertNode(n.Xpr),
		Paramkind:   pg.ParamKind(n.Paramkind),
		Paramid:     n.Paramid,
		Paramtype:   pg.Oid(n.Paramtype),
		Paramtypmod: n.Paramtypmod,
		Paramcollid: pg.Oid(n.Paramcollid),
		Location:    n.Location,
	}
}

func convertTransactionStmt(n *nodes.TransactionStmt) *pg.TransactionStmt {
	return &pg.TransactionStmt{
		Kind:    pg.TransactionStmtKind(n.Kind),
		Options: convertList(n.Options),
		Gid:     n.Gid,
	}
}

func convertFunctionParameter(n *nodes.FunctionParameter) *pg.FunctionParameter {
	return &pg.FunctionParameter{
		Name:    n.Name,
		ArgType: convertTypeName(n.ArgType),
		Mode:    pg.FunctionParameterMode(n.Mode),
		Defexpr: convertNode(n.Defexpr),
	}
}

func convertTableSampleClause(n *nodes.TableSampleClause) *pg.TableSampleClause {
	return &pg.TableSampleClause{
		Tsmhandler: pg.Oid(n.Tsmhandler),
		Args:       convertList(n.Args),
		Repeatable: convertNode(n.Repeatable),
	}
}

func convertAlterRoleSetStmt(n *nodes.AlterRoleSetStmt) *pg.AlterRoleSetStmt {
	return &pg.AlterRoleSetStmt{
		Role:     convertRoleSpec(n.Role),
		Database: n.Database,
		Setstmt:  convertVariableSetStmt(n.Setstmt),
	}
}

func convertRoleSpec(n *nodes.RoleSpec) *pg.RoleSpec {
	return &pg.RoleSpec{
		Roletype: pg.RoleSpecType(n.Roletype),
		Rolename: n.Rolename,
		Location: n.Location,
	}
}

func convertRangeTblFunction(n *nodes.RangeTblFunction) *pg.RangeTblFunction {
	return &pg.RangeTblFunction{
		Funcexpr:          convertNode(n.Funcexpr),
		Funccolcount:      n.Funccolcount,
		Funccolnames:      convertList(n.Funccolnames),
		Funccoltypes:      convertList(n.Funccoltypes),
		Funccoltypmods:    convertList(n.Funccoltypmods),
		Funccolcollations: convertList(n.Funccolcollations),
		Funcparams:        n.Funcparams,
	}
}

func convertAlterPolicyStmt(n *nodes.AlterPolicyStmt) *pg.AlterPolicyStmt {
	return &pg.AlterPolicyStmt{
		PolicyName: n.PolicyName,
		Table:      convertRangeVar(n.Table),
		Roles:      convertList(n.Roles),
		Qual:       convertNode(n.Qual),
		WithCheck:  convertNode(n.WithCheck),
	}
}

func convertCreatePLangStmt(n *nodes.CreatePLangStmt) *pg.CreatePLangStmt {
	return &pg.CreatePLangStmt{
		Replace:     n.Replace,
		Plname:      n.Plname,
		Plhandler:   convertList(n.Plhandler),
		Plinline:    convertList(n.Plinline),
		Plvalidator: convertList(n.Plvalidator),
		Pltrusted:   n.Pltrusted,
	}
}

func convertDropTableSpaceStmt(n *nodes.DropTableSpaceStmt) *pg.DropTableSpaceStmt {
	return &pg.DropTableSpaceStmt{
		Tablespacename: n.Tablespacename,
		MissingOk:      n.MissingOk,
	}
}

func convertAlterTableCmd(n *nodes.AlterTableCmd) *pg.AlterTableCmd {
	return &pg.AlterTableCmd{
		Subtype:   pg.AlterTableType(n.Subtype),
		Name:      n.Name,
		Newowner:  convertRoleSpec(n.Newowner),
		Def:       convertNode(n.Def),
		Behavior:  pg.DropBehavior(n.Behavior),
		MissingOk: n.MissingOk,
	}
}

func convertAggref(n *nodes.Aggref) *pg.Aggref {
	return &pg.Aggref{
		Xpr:           convertNode(n.Xpr),
		Aggfnoid:      pg.Oid(n.Aggfnoid),
		Aggtype:       pg.Oid(n.Aggtype),
		Aggcollid:     pg.Oid(n.Aggcollid),
		Inputcollid:   pg.Oid(n.Inputcollid),
		Aggtranstype:  pg.Oid(n.Aggtranstype),
		Aggargtypes:   convertList(n.Aggargtypes),
		Aggdirectargs: convertList(n.Aggdirectargs),
		Args:          convertList(n.Args),
		Aggorder:      convertList(n.Aggorder),
		Aggdistinct:   convertList(n.Aggdistinct),
		Aggfilter:     convertNode(n.Aggfilter),
		Aggstar:       n.Aggstar,
		Aggvariadic:   n.Aggvariadic,
		Aggkind:       n.Aggkind,
		Agglevelsup:   pg.Index(n.Agglevelsup),
		Aggsplit:      pg.AggSplit(n.Aggsplit),
		Location:      n.Location,
	}
}

func convertReassignOwnedStmt(n *nodes.ReassignOwnedStmt) *pg.ReassignOwnedStmt {
	return &pg.ReassignOwnedStmt{
		Roles:   convertList(n.Roles),
		Newrole: convertRoleSpec(n.Newrole),
	}
}

func convertRangeTblEntry(n *nodes.RangeTblEntry) *pg.RangeTblEntry {
	return &pg.RangeTblEntry{
		Rtekind:         pg.RTEKind(n.Rtekind),
		Relid:           pg.Oid(n.Relid),
		Relkind:         n.Relkind,
		Tablesample:     convertTableSampleClause(n.Tablesample),
		Subquery:        convertQuery(n.Subquery),
		SecurityBarrier: n.SecurityBarrier,
		Jointype:        pg.JoinType(n.Jointype),
		Joinaliasvars:   convertList(n.Joinaliasvars),
		Functions:       convertList(n.Functions),
		Funcordinality:  n.Funcordinality,
		Tablefunc:       convertTableFunc(n.Tablefunc),
		ValuesLists:     convertList(n.ValuesLists),
		Ctename:         n.Ctename,
		Ctelevelsup:     pg.Index(n.Ctelevelsup),
		SelfReference:   n.SelfReference,
		Coltypes:        convertList(n.Coltypes),
		Coltypmods:      convertList(n.Coltypmods),
		Colcollations:   convertList(n.Colcollations),
		Enrname:         n.Enrname,
		Enrtuples:       n.Enrtuples,
		Alias:           convertAlias(n.Alias),
		Eref:            convertAlias(n.Eref),
		Lateral:         n.Lateral,
		Inh:             n.Inh,
		InFromCl:        n.InFromCl,
		RequiredPerms:   pg.AclMode(n.RequiredPerms),
		CheckAsUser:     pg.Oid(n.CheckAsUser),
		SelectedCols:    n.SelectedCols,
		InsertedCols:    n.InsertedCols,
		UpdatedCols:     n.UpdatedCols,
		SecurityQuals:   convertList(n.SecurityQuals),
	}
}

func convertAccessPriv(n *nodes.AccessPriv) *pg.AccessPriv {
	return &pg.AccessPriv{
		PrivName: n.PrivName,
		Cols:     convertList(n.Cols),
	}
}

func convertTypeName(n *nodes.TypeName) *pg.TypeName {
	return &pg.TypeName{
		Names:       convertList(n.Names),
		TypeOid:     pg.Oid(n.TypeOid),
		Setof:       n.Setof,
		PctType:     n.PctType,
		Typmods:     convertList(n.Typmods),
		Typemod:     n.Typemod,
		ArrayBounds: convertList(n.ArrayBounds),
		Location:    n.Location,
	}
}

func convertSubPlan(n *nodes.SubPlan) *pg.SubPlan {
	return &pg.SubPlan{
		Xpr:               convertNode(n.Xpr),
		SubLinkType:       pg.SubLinkType(n.SubLinkType),
		Testexpr:          convertNode(n.Testexpr),
		ParamIds:          convertList(n.ParamIds),
		PlanId:            n.PlanId,
		PlanName:          n.PlanName,
		FirstColType:      pg.Oid(n.FirstColType),
		FirstColTypmod:    n.FirstColTypmod,
		FirstColCollation: pg.Oid(n.FirstColCollation),
		UseHashTable:      n.UseHashTable,
		UnknownEqFalse:    n.UnknownEqFalse,
		ParallelSafe:      n.ParallelSafe,
		SetParam:          convertList(n.SetParam),
		ParParam:          convertList(n.ParParam),
		Args:              convertList(n.Args),
		StartupCost:       pg.Cost(n.StartupCost),
		PerCallCost:       pg.Cost(n.PerCallCost),
	}
}

func convertCreateTrigStmt(n *nodes.CreateTrigStmt) *pg.CreateTrigStmt {
	return &pg.CreateTrigStmt{
		Trigname:       n.Trigname,
		Relation:       convertRangeVar(n.Relation),
		Funcname:       convertList(n.Funcname),
		Args:           convertList(n.Args),
		Row:            n.Row,
		Timing:         n.Timing,
		Events:         n.Events,
		Columns:        convertList(n.Columns),
		WhenClause:     convertNode(n.WhenClause),
		Isconstraint:   n.Isconstraint,
		TransitionRels: convertList(n.TransitionRels),
		Deferrable:     n.Deferrable,
		Initdeferred:   n.Initdeferred,
		Constrrel:      convertRangeVar(n.Constrrel),
	}
}

func convertCoerceViaIO(n *nodes.CoerceViaIO) *pg.CoerceViaIO {
	return &pg.CoerceViaIO{
		Xpr:          convertNode(n.Xpr),
		Arg:          convertNode(n.Arg),
		Resulttype:   pg.Oid(n.Resulttype),
		Resultcollid: pg.Oid(n.Resultcollid),
		Coerceformat: pg.CoercionForm(n.Coerceformat),
		Location:     n.Location,
	}
}

func convertViewStmt(n *nodes.ViewStmt) *pg.ViewStmt {
	return &pg.ViewStmt{
		View:            convertRangeVar(n.View),
		Aliases:         convertList(n.Aliases),
		Query:           convertNode(n.Query),
		Replace:         n.Replace,
		Options:         convertList(n.Options),
		WithCheckOption: pg.ViewCheckOption(n.WithCheckOption),
	}
}

func convertCaseExpr(n *nodes.CaseExpr) *pg.CaseExpr {
	return &pg.CaseExpr{
		Xpr:        convertNode(n.Xpr),
		Casetype:   pg.Oid(n.Casetype),
		Casecollid: pg.Oid(n.Casecollid),
		Arg:        convertNode(n.Arg),
		Args:       convertList(n.Args),
		Defresult:  convertNode(n.Defresult),
		Location:   n.Location,
	}
}

func convertRowExpr(n *nodes.RowExpr) *pg.RowExpr {
	return &pg.RowExpr{
		Xpr:       convertNode(n.Xpr),
		Args:      convertList(n.Args),
		RowTypeid: pg.Oid(n.RowTypeid),
		RowFormat: pg.CoercionForm(n.RowFormat),
		Colnames:  convertList(n.Colnames),
		Location:  n.Location,
	}
}

func convertSetOperationStmt(n *nodes.SetOperationStmt) *pg.SetOperationStmt {
	return &pg.SetOperationStmt{
		Op:            pg.SetOperation(n.Op),
		All:           n.All,
		Larg:          convertNode(n.Larg),
		Rarg:          convertNode(n.Rarg),
		ColTypes:      convertList(n.ColTypes),
		ColTypmods:    convertList(n.ColTypmods),
		ColCollations: convertList(n.ColCollations),
		GroupClauses:  convertList(n.GroupClauses),
	}
}

func convertAlterUserMappingStmt(n *nodes.AlterUserMappingStmt) *pg.AlterUserMappingStmt {
	return &pg.AlterUserMappingStmt{
		User:       convertRoleSpec(n.User),
		Servername: n.Servername,
		Options:    convertList(n.Options),
	}
}

func convertLockingClause(n *nodes.LockingClause) *pg.LockingClause {
	return &pg.LockingClause{
		LockedRels: convertList(n.LockedRels),
		Strength:   pg.LockClauseStrength(n.Strength),
		WaitPolicy: pg.LockWaitPolicy(n.WaitPolicy),
	}
}

func convertRangeTableFuncCol(n *nodes.RangeTableFuncCol) *pg.RangeTableFuncCol {
	return &pg.RangeTableFuncCol{
		Colname:       n.Colname,
		TypeName:      convertTypeName(n.TypeName),
		ForOrdinality: n.ForOrdinality,
		IsNotNull:     n.IsNotNull,
		Colexpr:       convertNode(n.Colexpr),
		Coldefexpr:    convertNode(n.Coldefexpr),
		Location:      n.Location,
	}
}

func convertAlterExtensionStmt(n *nodes.AlterExtensionStmt) *pg.AlterExtensionStmt {
	return &pg.AlterExtensionStmt{
		Extname: n.Extname,
		Options: convertList(n.Options),
	}
}

func convertParamExternData(n *nodes.ParamExternData) *pg.ParamExternData {
	return &pg.ParamExternData{
		Value:  pg.Datum(n.Value),
		Isnull: n.Isnull,
		Pflags: n.Pflags,
		Ptype:  pg.Oid(n.Ptype),
	}
}

func convertCreateFdwStmt(n *nodes.CreateFdwStmt) *pg.CreateFdwStmt {
	return &pg.CreateFdwStmt{
		Fdwname:     n.Fdwname,
		FuncOptions: convertList(n.FuncOptions),
		Options:     convertList(n.Options),
	}
}

func convertCreateOpFamilyStmt(n *nodes.CreateOpFamilyStmt) *pg.CreateOpFamilyStmt {
	return &pg.CreateOpFamilyStmt{
		Opfamilyname: convertList(n.Opfamilyname),
		Amname:       n.Amname,
	}
}

func convertDeclareCursorStmt(n *nodes.DeclareCursorStmt) *pg.DeclareCursorStmt {
	return &pg.DeclareCursorStmt{
		Portalname: n.Portalname,
		Options:    n.Options,
		Query:      convertNode(n.Query),
	}
}

func convertPartitionRangeDatum(n *nodes.PartitionRangeDatum) *pg.PartitionRangeDatum {
	return &pg.PartitionRangeDatum{
		Kind:     pg.PartitionRangeDatumKind(n.Kind),
		Value:    convertNode(n.Value),
		Location: n.Location,
	}
}

func convertAlterPublicationStmt(n *nodes.AlterPublicationStmt) *pg.AlterPublicationStmt {
	return &pg.AlterPublicationStmt{
		Pubname:      n.Pubname,
		Options:      convertList(n.Options),
		Tables:       convertList(n.Tables),
		ForAllTables: n.ForAllTables,
		TableAction:  pg.DefElemAction(n.TableAction),
	}
}

func convertDropRoleStmt(n *nodes.DropRoleStmt) *pg.DropRoleStmt {
	return &pg.DropRoleStmt{
		Roles:     convertList(n.Roles),
		MissingOk: n.MissingOk,
	}
}

func convertGroupingFunc(n *nodes.GroupingFunc) *pg.GroupingFunc {
	return &pg.GroupingFunc{
		Xpr:         convertNode(n.Xpr),
		Args:        convertList(n.Args),
		Refs:        convertList(n.Refs),
		Cols:        convertList(n.Cols),
		Agglevelsup: pg.Index(n.Agglevelsup),
		Location:    n.Location,
	}
}

func convertRowCompareExpr(n *nodes.RowCompareExpr) *pg.RowCompareExpr {
	return &pg.RowCompareExpr{
		Xpr:          convertNode(n.Xpr),
		Rctype:       pg.RowCompareType(n.Rctype),
		Opnos:        convertList(n.Opnos),
		Opfamilies:   convertList(n.Opfamilies),
		Inputcollids: convertList(n.Inputcollids),
		Largs:        convertList(n.Largs),
		Rargs:        convertList(n.Rargs),
	}
}

func convertReplicaIdentityStmt(n *nodes.ReplicaIdentityStmt) *pg.ReplicaIdentityStmt {
	return &pg.ReplicaIdentityStmt{
		IdentityType: n.IdentityType,
		Name:         n.Name,
	}
}

func convertAlterSeqStmt(n *nodes.AlterSeqStmt) *pg.AlterSeqStmt {
	return &pg.AlterSeqStmt{
		Sequence:    convertRangeVar(n.Sequence),
		Options:     convertList(n.Options),
		ForIdentity: n.ForIdentity,
		MissingOk:   n.MissingOk,
	}
}

func convertVar(n *nodes.Var) *pg.Var {
	return &pg.Var{
		Xpr:         convertNode(n.Xpr),
		Varno:       pg.Index(n.Varno),
		Varattno:    pg.AttrNumber(n.Varattno),
		Vartype:     pg.Oid(n.Vartype),
		Vartypmod:   n.Vartypmod,
		Varcollid:   pg.Oid(n.Varcollid),
		Varlevelsup: pg.Index(n.Varlevelsup),
		Varnoold:    pg.Index(n.Varnoold),
		Varoattno:   pg.AttrNumber(n.Varoattno),
		Location:    n.Location,
	}
}

func convertPartitionBoundSpec(n *nodes.PartitionBoundSpec) *pg.PartitionBoundSpec {
	return &pg.PartitionBoundSpec{
		Strategy:    n.Strategy,
		Listdatums:  convertList(n.Listdatums),
		Lowerdatums: convertList(n.Lowerdatums),
		Upperdatums: convertList(n.Upperdatums),
		Location:    n.Location,
	}
}

func convertTruncateStmt(n *nodes.TruncateStmt) *pg.TruncateStmt {
	return &pg.TruncateStmt{
		Relations:   convertList(n.Relations),
		RestartSeqs: n.RestartSeqs,
		Behavior:    pg.DropBehavior(n.Behavior),
	}
}

func convertDropUserMappingStmt(n *nodes.DropUserMappingStmt) *pg.DropUserMappingStmt {
	return &pg.DropUserMappingStmt{
		User:       convertRoleSpec(n.User),
		Servername: n.Servername,
		MissingOk:  n.MissingOk,
	}
}

func convertTableLikeClause(n *nodes.TableLikeClause) *pg.TableLikeClause {
	return &pg.TableLikeClause{
		Relation: convertRangeVar(n.Relation),
		Options:  n.Options,
	}
}

func convertJoinExpr(n *nodes.JoinExpr) *pg.JoinExpr {
	return &pg.JoinExpr{
		Jointype:    pg.JoinType(n.Jointype),
		IsNatural:   n.IsNatural,
		Larg:        convertNode(n.Larg),
		Rarg:        convertNode(n.Rarg),
		UsingClause: convertList(n.UsingClause),
		Quals:       convertNode(n.Quals),
		Alias:       convertAlias(n.Alias),
		Rtindex:     n.Rtindex,
	}
}

func convertCreateTableSpaceStmt(n *nodes.CreateTableSpaceStmt) *pg.CreateTableSpaceStmt {
	return &pg.CreateTableSpaceStmt{
		Tablespacename: n.Tablespacename,
		Owner:          convertRoleSpec(n.Owner),
		Location:       n.Location,
		Options:        convertList(n.Options),
	}
}

func convertOnConflictExpr(n *nodes.OnConflictExpr) *pg.OnConflictExpr {
	return &pg.OnConflictExpr{
		Action:          pg.OnConflictAction(n.Action),
		ArbiterElems:    convertList(n.ArbiterElems),
		ArbiterWhere:    convertNode(n.ArbiterWhere),
		Constraint:      pg.Oid(n.Constraint),
		OnConflictSet:   convertList(n.OnConflictSet),
		OnConflictWhere: convertNode(n.OnConflictWhere),
		ExclRelIndex:    n.ExclRelIndex,
		ExclRelTlist:    convertList(n.ExclRelTlist),
	}
}

func convertParamExecData(n *nodes.ParamExecData) *pg.ParamExecData {
	return &pg.ParamExecData{
		ExecPlan: &ast.TODO{},
		Value:    pg.Datum(n.Value),
		Isnull:   n.Isnull,
	}
}

func convertSecLabelStmt(n *nodes.SecLabelStmt) *pg.SecLabelStmt {
	return &pg.SecLabelStmt{
		Objtype:  pg.ObjectType(n.Objtype),
		Object:   convertNode(n.Object),
		Provider: n.Provider,
		Label:    n.Label,
	}
}

func convertBlockIdData(n *nodes.BlockIdData) *pg.BlockIdData {
	return &pg.BlockIdData{
		BiHi: n.BiHi,
		BiLo: n.BiLo,
	}
}

func convertGrantStmt(n *nodes.GrantStmt) *pg.GrantStmt {
	return &pg.GrantStmt{
		IsGrant:     n.IsGrant,
		Targtype:    pg.GrantTargetType(n.Targtype),
		Objtype:     pg.GrantObjectType(n.Objtype),
		Objects:     convertList(n.Objects),
		Privileges:  convertList(n.Privileges),
		Grantees:    convertList(n.Grantees),
		GrantOption: n.GrantOption,
		Behavior:    pg.DropBehavior(n.Behavior),
	}
}

func convertTargetEntry(n *nodes.TargetEntry) *pg.TargetEntry {
	return &pg.TargetEntry{
		Xpr:             convertNode(n.Xpr),
		Expr:            convertNode(n.Expr),
		Resno:           pg.AttrNumber(n.Resno),
		Resname:         n.Resname,
		Ressortgroupref: pg.Index(n.Ressortgroupref),
		Resorigtbl:      pg.Oid(n.Resorigtbl),
		Resorigcol:      pg.AttrNumber(n.Resorigcol),
		Resjunk:         n.Resjunk,
	}
}

func convertRangeTableSample(n *nodes.RangeTableSample) *pg.RangeTableSample {
	return &pg.RangeTableSample{
		Relation:   convertNode(n.Relation),
		Method:     convertList(n.Method),
		Args:       convertList(n.Args),
		Repeatable: convertNode(n.Repeatable),
		Location:   n.Location,
	}
}

func convertBitString(n *nodes.BitString) *pg.BitString {
	return &pg.BitString{
		Str: n.Str,
	}
}

func convertAlias(n *nodes.Alias) *pg.Alias {
	return &pg.Alias{
		Aliasname: n.Aliasname,
		Colnames:  convertList(n.Colnames),
	}
}

func convertAlterTSConfigurationStmt(n *nodes.AlterTSConfigurationStmt) *pg.AlterTSConfigurationStmt {
	return &pg.AlterTSConfigurationStmt{
		Kind:      pg.AlterTSConfigType(n.Kind),
		Cfgname:   convertList(n.Cfgname),
		Tokentype: convertList(n.Tokentype),
		Dicts:     convertList(n.Dicts),
		Override:  n.Override,
		Replace:   n.Replace,
		MissingOk: n.MissingOk,
	}
}

func convertLoadStmt(n *nodes.LoadStmt) *pg.LoadStmt {
	return &pg.LoadStmt{
		Filename: n.Filename,
	}
}

func convertIntoClause(n *nodes.IntoClause) *pg.IntoClause {
	return &pg.IntoClause{
		Rel:            convertRangeVar(n.Rel),
		ColNames:       convertList(n.ColNames),
		Options:        convertList(n.Options),
		OnCommit:       pg.OnCommitAction(n.OnCommit),
		TableSpaceName: n.TableSpaceName,
		ViewQuery:      convertNode(n.ViewQuery),
		SkipData:       n.SkipData,
	}
}

func convertInteger(n *nodes.Integer) *pg.Integer {
	return &pg.Integer{
		Ival: n.Ival,
	}
}

func convertAlterForeignServerStmt(n *nodes.AlterForeignServerStmt) *pg.AlterForeignServerStmt {
	return &pg.AlterForeignServerStmt{
		Servername: n.Servername,
		Version:    n.Version,
		Options:    convertList(n.Options),
		HasVersion: n.HasVersion,
	}
}

func convertCreateEventTrigStmt(n *nodes.CreateEventTrigStmt) *pg.CreateEventTrigStmt {
	return &pg.CreateEventTrigStmt{
		Trigname:   n.Trigname,
		Eventname:  n.Eventname,
		Whenclause: convertList(n.Whenclause),
		Funcname:   convertList(n.Funcname),
	}
}

func convertRuleStmt(n *nodes.RuleStmt) *pg.RuleStmt {
	return &pg.RuleStmt{
		Relation:    convertRangeVar(n.Relation),
		Rulename:    n.Rulename,
		WhereClause: convertNode(n.WhereClause),
		Event:       pg.CmdType(n.Event),
		Instead:     n.Instead,
		Actions:     convertList(n.Actions),
		Replace:     n.Replace,
	}
}

func convertNode(node nodes.Node) ast.Node {
	switch n := node.(type) {

	case nodes.UnlistenStmt:
		return convertUnlistenStmt(&n)

	case nodes.DeleteStmt:
		return convertDeleteStmt(&n)

	case nodes.RelabelType:
		return convertRelabelType(&n)

	case nodes.CreateUserMappingStmt:
		return convertCreateUserMappingStmt(&n)

	case nodes.CreateAmStmt:
		return convertCreateAmStmt(&n)

	case nodes.GrantRoleStmt:
		return convertGrantRoleStmt(&n)

	case nodes.FuncExpr:
		return convertFuncExpr(&n)

	case nodes.CaseTestExpr:
		return convertCaseTestExpr(&n)

	case nodes.CaseWhen:
		return convertCaseWhen(&n)

	case nodes.ResTarget:
		return convertResTarget(&n)

	case nodes.AlterExtensionContentsStmt:
		return convertAlterExtensionContentsStmt(&n)

	case nodes.ReindexStmt:
		return convertReindexStmt(&n)

	case nodes.WindowDef:
		return convertWindowDef(&n)

	case nodes.DiscardStmt:
		return convertDiscardStmt(&n)

	case nodes.ParamListInfoData:
		return convertParamListInfoData(&n)

	case nodes.CollateClause:
		return convertCollateClause(&n)

	case nodes.CreatedbStmt:
		return convertCreatedbStmt(&n)

	case nodes.Null:
		return convertNull(&n)

	case nodes.RefreshMatViewStmt:
		return convertRefreshMatViewStmt(&n)

	case nodes.ListenStmt:
		return convertListenStmt(&n)

	case nodes.Expr:
		return convertExpr(&n)

	case nodes.SortGroupClause:
		return convertSortGroupClause(&n)

	case nodes.CreateExtensionStmt:
		return convertCreateExtensionStmt(&n)

	case nodes.CreateEnumStmt:
		return convertCreateEnumStmt(&n)

	case nodes.DoStmt:
		return convertDoStmt(&n)

	case nodes.XmlSerialize:
		return convertXmlSerialize(&n)

	case nodes.PartitionCmd:
		return convertPartitionCmd(&n)

	case nodes.WithCheckOption:
		return convertWithCheckOption(&n)

	case nodes.CreateRoleStmt:
		return convertCreateRoleStmt(&n)

	case nodes.ClosePortalStmt:
		return convertClosePortalStmt(&n)

	case nodes.RangeTableFunc:
		return convertRangeTableFunc(&n)

	case nodes.DropdbStmt:
		return convertDropdbStmt(&n)

	case nodes.RangeSubselect:
		return convertRangeSubselect(&n)

	case nodes.AlterSystemStmt:
		return convertAlterSystemStmt(&n)

	case nodes.ArrayCoerceExpr:
		return convertArrayCoerceExpr(&n)

	case nodes.AlterDatabaseStmt:
		return convertAlterDatabaseStmt(&n)

	case nodes.DefElem:
		return convertDefElem(&n)

	case nodes.TypeCast:
		return convertTypeCast(&n)

	case nodes.ClusterStmt:
		return convertClusterStmt(&n)

	case nodes.InferenceElem:
		return convertInferenceElem(&n)

	case nodes.ExplainStmt:
		return convertExplainStmt(&n)

	case nodes.NamedArgExpr:
		return convertNamedArgExpr(&n)

	case nodes.DefineStmt:
		return convertDefineStmt(&n)

	case nodes.SubLink:
		return convertSubLink(&n)

	case nodes.WindowFunc:
		return convertWindowFunc(&n)

	case nodes.String:
		return convertString(&n)

	case nodes.ObjectWithArgs:
		return convertObjectWithArgs(&n)

	case nodes.NullTest:
		return convertNullTest(&n)

	case nodes.IndexElem:
		return convertIndexElem(&n)

	case nodes.TableFunc:
		return convertTableFunc(&n)

	case nodes.CreateSeqStmt:
		return convertCreateSeqStmt(&n)

	case nodes.CreateStatsStmt:
		return convertCreateStatsStmt(&n)

	case nodes.CreateRangeStmt:
		return convertCreateRangeStmt(&n)

	case nodes.DropStmt:
		return convertDropStmt(&n)

	case nodes.CommentStmt:
		return convertCommentStmt(&n)

	case nodes.CreateSubscriptionStmt:
		return convertCreateSubscriptionStmt(&n)

	case nodes.CreatePolicyStmt:
		return convertCreatePolicyStmt(&n)

	case nodes.AlterTableMoveAllStmt:
		return convertAlterTableMoveAllStmt(&n)

	case nodes.AlterObjectDependsStmt:
		return convertAlterObjectDependsStmt(&n)

	case nodes.WindowClause:
		return convertWindowClause(&n)

	case nodes.ColumnDef:
		return convertColumnDef(&n)

	case nodes.SelectStmt:
		return convertSelectStmt(&n)

	case nodes.LockStmt:
		return convertLockStmt(&n)

	case nodes.FieldStore:
		return convertFieldStore(&n)

	case nodes.AlterSubscriptionStmt:
		return convertAlterSubscriptionStmt(&n)

	case nodes.AlterFdwStmt:
		return convertAlterFdwStmt(&n)

	case nodes.InlineCodeBlock:
		return convertInlineCodeBlock(&n)

	case nodes.A_Indices:
		return convertA_Indices(&n)

	case nodes.AlterTableSpaceOptionsStmt:
		return convertAlterTableSpaceOptionsStmt(&n)

	case nodes.CoerceToDomainValue:
		return convertCoerceToDomainValue(&n)

	case nodes.BooleanTest:
		return convertBooleanTest(&n)

	case nodes.VariableSetStmt:
		return convertVariableSetStmt(&n)

	case nodes.AlterDatabaseSetStmt:
		return convertAlterDatabaseSetStmt(&n)

	case nodes.UpdateStmt:
		return convertUpdateStmt(&n)

	case nodes.A_Const:
		return convertA_Const(&n)

	case nodes.DropSubscriptionStmt:
		return convertDropSubscriptionStmt(&n)

	case nodes.A_ArrayExpr:
		return convertA_ArrayExpr(&n)

	case nodes.BoolExpr:
		return convertBoolExpr(&n)

	case nodes.A_Star:
		return convertA_Star(&n)

	case nodes.AlterTSDictionaryStmt:
		return convertAlterTSDictionaryStmt(&n)

	case nodes.Const:
		return convertConst(&n)

	case nodes.CommonTableExpr:
		return convertCommonTableExpr(&n)

	case nodes.RenameStmt:
		return convertRenameStmt(&n)

	case nodes.CurrentOfExpr:
		return convertCurrentOfExpr(&n)

	case nodes.A_Expr:
		return convertA_Expr(&n)

	case nodes.IndexStmt:
		return convertIndexStmt(&n)

	case nodes.PartitionElem:
		return convertPartitionElem(&n)

	case nodes.CreateTableAsStmt:
		return convertCreateTableAsStmt(&n)

	case nodes.RowMarkClause:
		return convertRowMarkClause(&n)

	case nodes.AlterObjectSchemaStmt:
		return convertAlterObjectSchemaStmt(&n)

	case nodes.SQLValueFunction:
		return convertSQLValueFunction(&n)

	case nodes.VacuumStmt:
		return convertVacuumStmt(&n)

	case nodes.CollateExpr:
		return convertCollateExpr(&n)

	case nodes.ConstraintsSetStmt:
		return convertConstraintsSetStmt(&n)

	case nodes.NextValueExpr:
		return convertNextValueExpr(&n)

	case nodes.GroupingSet:
		return convertGroupingSet(&n)

	case nodes.RangeVar:
		return convertRangeVar(&n)

	case nodes.MinMaxExpr:
		return convertMinMaxExpr(&n)

	case nodes.ArrayExpr:
		return convertArrayExpr(&n)

	case nodes.FieldSelect:
		return convertFieldSelect(&n)

	case nodes.Query:
		return convertQuery(&n)

	case nodes.NotifyStmt:
		return convertNotifyStmt(&n)

	case nodes.FuncCall:
		return convertFuncCall(&n)

	case nodes.CoalesceExpr:
		return convertCoalesceExpr(&n)

	case nodes.ExecuteStmt:
		return convertExecuteStmt(&n)

	case nodes.CreateForeignTableStmt:
		return convertCreateForeignTableStmt(&n)

	case nodes.CreateStmt:
		return convertCreateStmt(&n)

	case nodes.CreateOpClassStmt:
		return convertCreateOpClassStmt(&n)

	case nodes.FromExpr:
		return convertFromExpr(&n)

	case nodes.InsertStmt:
		return convertInsertStmt(&n)

	case nodes.ColumnRef:
		return convertColumnRef(&n)

	case nodes.CreateCastStmt:
		return convertCreateCastStmt(&n)

	case nodes.CoerceToDomain:
		return convertCoerceToDomain(&n)

	case nodes.AlterEnumStmt:
		return convertAlterEnumStmt(&n)

	case nodes.RawStmt:
		return convertRawStmt(&n)

	case nodes.CreateDomainStmt:
		return convertCreateDomainStmt(&n)

	case nodes.ParamRef:
		return convertParamRef(&n)

	case nodes.Float:
		return convertFloat(&n)

	case nodes.MultiAssignRef:
		return convertMultiAssignRef(&n)

	case nodes.TriggerTransition:
		return convertTriggerTransition(&n)

	case nodes.ArrayRef:
		return convertArrayRef(&n)

	case nodes.ImportForeignSchemaStmt:
		return convertImportForeignSchemaStmt(&n)

	case nodes.AlterOpFamilyStmt:
		return convertAlterOpFamilyStmt(&n)

	case nodes.AlterOwnerStmt:
		return convertAlterOwnerStmt(&n)

	case nodes.OpExpr:
		return convertOpExpr(&n)

	case nodes.AlterRoleStmt:
		return convertAlterRoleStmt(&n)

	case nodes.RangeTblRef:
		return convertRangeTblRef(&n)

	case nodes.CreateTransformStmt:
		return convertCreateTransformStmt(&n)

	case nodes.ConvertRowtypeExpr:
		return convertConvertRowtypeExpr(&n)

	case nodes.CompositeTypeStmt:
		return convertCompositeTypeStmt(&n)

	case nodes.RangeFunction:
		return convertRangeFunction(&n)

	case nodes.CreateForeignServerStmt:
		return convertCreateForeignServerStmt(&n)

	case nodes.AlterCollationStmt:
		return convertAlterCollationStmt(&n)

	case nodes.AlterEventTrigStmt:
		return convertAlterEventTrigStmt(&n)

	case nodes.PrepareStmt:
		return convertPrepareStmt(&n)

	case nodes.CreateConversionStmt:
		return convertCreateConversionStmt(&n)

	case nodes.AlterTableStmt:
		return convertAlterTableStmt(&n)

	case nodes.CreatePublicationStmt:
		return convertCreatePublicationStmt(&n)

	case nodes.DeallocateStmt:
		return convertDeallocateStmt(&n)

	case nodes.PartitionSpec:
		return convertPartitionSpec(&n)

	case nodes.OnConflictClause:
		return convertOnConflictClause(&n)

	case nodes.DropOwnedStmt:
		return convertDropOwnedStmt(&n)

	case nodes.SetToDefault:
		return convertSetToDefault(&n)

	case nodes.SortBy:
		return convertSortBy(&n)

	case nodes.AlternativeSubPlan:
		return convertAlternativeSubPlan(&n)

	case nodes.VariableShowStmt:
		return convertVariableShowStmt(&n)

	case nodes.InferClause:
		return convertInferClause(&n)

	case nodes.A_Indirection:
		return convertA_Indirection(&n)

	case nodes.AlterOperatorStmt:
		return convertAlterOperatorStmt(&n)

	case nodes.AlterDefaultPrivilegesStmt:
		return convertAlterDefaultPrivilegesStmt(&n)

	case nodes.FetchStmt:
		return convertFetchStmt(&n)

	case nodes.Constraint:
		return convertConstraint(&n)

	case nodes.CopyStmt:
		return convertCopyStmt(&n)

	case nodes.CreateOpClassItem:
		return convertCreateOpClassItem(&n)

	case nodes.CreateSchemaStmt:
		return convertCreateSchemaStmt(&n)

	case nodes.WithClause:
		return convertWithClause(&n)

	case nodes.CreateFunctionStmt:
		return convertCreateFunctionStmt(&n)

	case nodes.CheckPointStmt:
		return convertCheckPointStmt(&n)

	case nodes.AlterDomainStmt:
		return convertAlterDomainStmt(&n)

	case nodes.AlterFunctionStmt:
		return convertAlterFunctionStmt(&n)

	case nodes.ScalarArrayOpExpr:
		return convertScalarArrayOpExpr(&n)

	case nodes.XmlExpr:
		return convertXmlExpr(&n)

	case nodes.Param:
		return convertParam(&n)

	case nodes.TransactionStmt:
		return convertTransactionStmt(&n)

	case nodes.FunctionParameter:
		return convertFunctionParameter(&n)

	case nodes.TableSampleClause:
		return convertTableSampleClause(&n)

	case nodes.AlterRoleSetStmt:
		return convertAlterRoleSetStmt(&n)

	case nodes.RoleSpec:
		return convertRoleSpec(&n)

	case nodes.RangeTblFunction:
		return convertRangeTblFunction(&n)

	case nodes.AlterPolicyStmt:
		return convertAlterPolicyStmt(&n)

	case nodes.CreatePLangStmt:
		return convertCreatePLangStmt(&n)

	case nodes.DropTableSpaceStmt:
		return convertDropTableSpaceStmt(&n)

	case nodes.AlterTableCmd:
		return convertAlterTableCmd(&n)

	case nodes.Aggref:
		return convertAggref(&n)

	case nodes.ReassignOwnedStmt:
		return convertReassignOwnedStmt(&n)

	case nodes.RangeTblEntry:
		return convertRangeTblEntry(&n)

	case nodes.AccessPriv:
		return convertAccessPriv(&n)

	case nodes.TypeName:
		return convertTypeName(&n)

	case nodes.SubPlan:
		return convertSubPlan(&n)

	case nodes.CreateTrigStmt:
		return convertCreateTrigStmt(&n)

	case nodes.CoerceViaIO:
		return convertCoerceViaIO(&n)

	case nodes.ViewStmt:
		return convertViewStmt(&n)

	case nodes.CaseExpr:
		return convertCaseExpr(&n)

	case nodes.RowExpr:
		return convertRowExpr(&n)

	case nodes.SetOperationStmt:
		return convertSetOperationStmt(&n)

	case nodes.AlterUserMappingStmt:
		return convertAlterUserMappingStmt(&n)

	case nodes.LockingClause:
		return convertLockingClause(&n)

	case nodes.RangeTableFuncCol:
		return convertRangeTableFuncCol(&n)

	case nodes.AlterExtensionStmt:
		return convertAlterExtensionStmt(&n)

	case nodes.ParamExternData:
		return convertParamExternData(&n)

	case nodes.CreateFdwStmt:
		return convertCreateFdwStmt(&n)

	case nodes.CreateOpFamilyStmt:
		return convertCreateOpFamilyStmt(&n)

	case nodes.DeclareCursorStmt:
		return convertDeclareCursorStmt(&n)

	case nodes.PartitionRangeDatum:
		return convertPartitionRangeDatum(&n)

	case nodes.AlterPublicationStmt:
		return convertAlterPublicationStmt(&n)

	case nodes.DropRoleStmt:
		return convertDropRoleStmt(&n)

	case nodes.GroupingFunc:
		return convertGroupingFunc(&n)

	case nodes.RowCompareExpr:
		return convertRowCompareExpr(&n)

	case nodes.ReplicaIdentityStmt:
		return convertReplicaIdentityStmt(&n)

	case nodes.AlterSeqStmt:
		return convertAlterSeqStmt(&n)

	case nodes.Var:
		return convertVar(&n)

	case nodes.PartitionBoundSpec:
		return convertPartitionBoundSpec(&n)

	case nodes.TruncateStmt:
		return convertTruncateStmt(&n)

	case nodes.DropUserMappingStmt:
		return convertDropUserMappingStmt(&n)

	case nodes.TableLikeClause:
		return convertTableLikeClause(&n)

	case nodes.JoinExpr:
		return convertJoinExpr(&n)

	case nodes.CreateTableSpaceStmt:
		return convertCreateTableSpaceStmt(&n)

	case nodes.OnConflictExpr:
		return convertOnConflictExpr(&n)

	case nodes.ParamExecData:
		return convertParamExecData(&n)

	case nodes.SecLabelStmt:
		return convertSecLabelStmt(&n)

	case nodes.BlockIdData:
		return convertBlockIdData(&n)

	case nodes.GrantStmt:
		return convertGrantStmt(&n)

	case nodes.TargetEntry:
		return convertTargetEntry(&n)

	case nodes.RangeTableSample:
		return convertRangeTableSample(&n)

	case nodes.BitString:
		return convertBitString(&n)

	case nodes.Alias:
		return convertAlias(&n)

	case nodes.AlterTSConfigurationStmt:
		return convertAlterTSConfigurationStmt(&n)

	case nodes.LoadStmt:
		return convertLoadStmt(&n)

	case nodes.IntoClause:
		return convertIntoClause(&n)

	case nodes.Integer:
		return convertInteger(&n)

	case nodes.AlterForeignServerStmt:
		return convertAlterForeignServerStmt(&n)

	case nodes.CreateEventTrigStmt:
		return convertCreateEventTrigStmt(&n)

	case nodes.RuleStmt:
		return convertRuleStmt(&n)

	default:
		return &ast.TODO{}
	}
}
