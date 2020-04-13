package postgresql

import (
	nodes "github.com/lfittl/pg_query_go/nodes"

	"github.com/kyleconroy/sqlc/internal/sql/ast"
	"github.com/kyleconroy/sqlc/internal/sql/ast/pg"
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


func convertA_ArrayExpr(n *nodes.A_ArrayExpr) *pg.A_ArrayExpr {
	if n == nil {
		return nil
	}
	return &pg.A_ArrayExpr{
		  Elements: convertList(n.Elements),
		  Location: n.Location,
	}
}

func convertA_Const(n *nodes.A_Const) *pg.A_Const {
	if n == nil {
		return nil
	}
	return &pg.A_Const{
		  Location: n.Location,
		  Val: convertNode(n.Val),
	}
}

func convertA_Expr(n *nodes.A_Expr) *pg.A_Expr {
	if n == nil {
		return nil
	}
	return &pg.A_Expr{
		  Kind: pg.A_Expr_Kind(n.Kind),
		  Lexpr: convertNode(n.Lexpr),
		  Location: n.Location,
		  Name: convertList(n.Name),
		  Rexpr: convertNode(n.Rexpr),
	}
}

func convertA_Indices(n *nodes.A_Indices) *pg.A_Indices {
	if n == nil {
		return nil
	}
	return &pg.A_Indices{
		  IsSlice: n.IsSlice,
		  Lidx: convertNode(n.Lidx),
		  Uidx: convertNode(n.Uidx),
	}
}

func convertA_Indirection(n *nodes.A_Indirection) *pg.A_Indirection {
	if n == nil {
		return nil
	}
	return &pg.A_Indirection{
		  Arg: convertNode(n.Arg),
		  Indirection: convertList(n.Indirection),
	}
}

func convertA_Star(n *nodes.A_Star) *pg.A_Star {
	if n == nil {
		return nil
	}
	return &pg.A_Star{
	}
}

func convertAccessPriv(n *nodes.AccessPriv) *pg.AccessPriv {
	if n == nil {
		return nil
	}
	return &pg.AccessPriv{
		  Cols: convertList(n.Cols),
		  PrivName: n.PrivName,
	}
}

func convertAggref(n *nodes.Aggref) *pg.Aggref {
	if n == nil {
		return nil
	}
	return &pg.Aggref{
		  Aggargtypes: convertList(n.Aggargtypes),
		  Aggcollid: pg.Oid(n.Aggcollid),
		  Aggdirectargs: convertList(n.Aggdirectargs),
		  Aggdistinct: convertList(n.Aggdistinct),
		  Aggfilter: convertNode(n.Aggfilter),
		  Aggfnoid: pg.Oid(n.Aggfnoid),
		  Aggkind: n.Aggkind,
		  Agglevelsup: pg.Index(n.Agglevelsup),
		  Aggorder: convertList(n.Aggorder),
		  Aggsplit: pg.AggSplit(n.Aggsplit),
		  Aggstar: n.Aggstar,
		  Aggtranstype: pg.Oid(n.Aggtranstype),
		  Aggtype: pg.Oid(n.Aggtype),
		  Aggvariadic: n.Aggvariadic,
		  Args: convertList(n.Args),
		  Inputcollid: pg.Oid(n.Inputcollid),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertAlias(n *nodes.Alias) *pg.Alias {
	if n == nil {
		return nil
	}
	return &pg.Alias{
		  Aliasname: n.Aliasname,
		  Colnames: convertList(n.Colnames),
	}
}

func convertAlterCollationStmt(n *nodes.AlterCollationStmt) *pg.AlterCollationStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterCollationStmt{
		  Collname: convertList(n.Collname),
	}
}

func convertAlterDatabaseSetStmt(n *nodes.AlterDatabaseSetStmt) *pg.AlterDatabaseSetStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterDatabaseSetStmt{
		  Dbname: n.Dbname,
		  Setstmt: convertVariableSetStmt(n.Setstmt),
	}
}

func convertAlterDatabaseStmt(n *nodes.AlterDatabaseStmt) *pg.AlterDatabaseStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterDatabaseStmt{
		  Dbname: n.Dbname,
		  Options: convertList(n.Options),
	}
}

func convertAlterDefaultPrivilegesStmt(n *nodes.AlterDefaultPrivilegesStmt) *pg.AlterDefaultPrivilegesStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterDefaultPrivilegesStmt{
		  Action: convertGrantStmt(n.Action),
		  Options: convertList(n.Options),
	}
}

func convertAlterDomainStmt(n *nodes.AlterDomainStmt) *pg.AlterDomainStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterDomainStmt{
		  Behavior: pg.DropBehavior(n.Behavior),
		  Def: convertNode(n.Def),
		  MissingOk: n.MissingOk,
		  Name: n.Name,
		  Subtype: n.Subtype,
		  TypeName: convertList(n.TypeName),
	}
}

func convertAlterEnumStmt(n *nodes.AlterEnumStmt) *pg.AlterEnumStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterEnumStmt{
		  NewVal: n.NewVal,
		  NewValIsAfter: n.NewValIsAfter,
		  NewValNeighbor: n.NewValNeighbor,
		  OldVal: n.OldVal,
		  SkipIfNewValExists: n.SkipIfNewValExists,
		  TypeName: convertList(n.TypeName),
	}
}

func convertAlterEventTrigStmt(n *nodes.AlterEventTrigStmt) *pg.AlterEventTrigStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterEventTrigStmt{
		  Tgenabled: n.Tgenabled,
		  Trigname: n.Trigname,
	}
}

func convertAlterExtensionContentsStmt(n *nodes.AlterExtensionContentsStmt) *pg.AlterExtensionContentsStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterExtensionContentsStmt{
		  Action: n.Action,
		  Extname: n.Extname,
		  Object: convertNode(n.Object),
		  Objtype: pg.ObjectType(n.Objtype),
	}
}

func convertAlterExtensionStmt(n *nodes.AlterExtensionStmt) *pg.AlterExtensionStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterExtensionStmt{
		  Extname: n.Extname,
		  Options: convertList(n.Options),
	}
}

func convertAlterFdwStmt(n *nodes.AlterFdwStmt) *pg.AlterFdwStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterFdwStmt{
		  Fdwname: n.Fdwname,
		  FuncOptions: convertList(n.FuncOptions),
		  Options: convertList(n.Options),
	}
}

func convertAlterForeignServerStmt(n *nodes.AlterForeignServerStmt) *pg.AlterForeignServerStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterForeignServerStmt{
		  HasVersion: n.HasVersion,
		  Options: convertList(n.Options),
		  Servername: n.Servername,
		  Version: n.Version,
	}
}

func convertAlterFunctionStmt(n *nodes.AlterFunctionStmt) *pg.AlterFunctionStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterFunctionStmt{
		  Actions: convertList(n.Actions),
		  Func: convertObjectWithArgs(n.Func),
	}
}

func convertAlterObjectDependsStmt(n *nodes.AlterObjectDependsStmt) *pg.AlterObjectDependsStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterObjectDependsStmt{
		  Extname: convertNode(n.Extname),
		  Object: convertNode(n.Object),
		  ObjectType: pg.ObjectType(n.ObjectType),
		  Relation: convertRangeVar(n.Relation),
	}
}

func convertAlterObjectSchemaStmt(n *nodes.AlterObjectSchemaStmt) *pg.AlterObjectSchemaStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterObjectSchemaStmt{
		  MissingOk: n.MissingOk,
		  Newschema: n.Newschema,
		  Object: convertNode(n.Object),
		  ObjectType: pg.ObjectType(n.ObjectType),
		  Relation: convertRangeVar(n.Relation),
	}
}

func convertAlterOpFamilyStmt(n *nodes.AlterOpFamilyStmt) *pg.AlterOpFamilyStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterOpFamilyStmt{
		  Amname: n.Amname,
		  IsDrop: n.IsDrop,
		  Items: convertList(n.Items),
		  Opfamilyname: convertList(n.Opfamilyname),
	}
}

func convertAlterOperatorStmt(n *nodes.AlterOperatorStmt) *pg.AlterOperatorStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterOperatorStmt{
		  Opername: convertObjectWithArgs(n.Opername),
		  Options: convertList(n.Options),
	}
}

func convertAlterOwnerStmt(n *nodes.AlterOwnerStmt) *pg.AlterOwnerStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterOwnerStmt{
		  Newowner: convertRoleSpec(n.Newowner),
		  Object: convertNode(n.Object),
		  ObjectType: pg.ObjectType(n.ObjectType),
		  Relation: convertRangeVar(n.Relation),
	}
}

func convertAlterPolicyStmt(n *nodes.AlterPolicyStmt) *pg.AlterPolicyStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterPolicyStmt{
		  PolicyName: n.PolicyName,
		  Qual: convertNode(n.Qual),
		  Roles: convertList(n.Roles),
		  Table: convertRangeVar(n.Table),
		  WithCheck: convertNode(n.WithCheck),
	}
}

func convertAlterPublicationStmt(n *nodes.AlterPublicationStmt) *pg.AlterPublicationStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterPublicationStmt{
		  ForAllTables: n.ForAllTables,
		  Options: convertList(n.Options),
		  Pubname: n.Pubname,
		  TableAction: pg.DefElemAction(n.TableAction),
		  Tables: convertList(n.Tables),
	}
}

func convertAlterRoleSetStmt(n *nodes.AlterRoleSetStmt) *pg.AlterRoleSetStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterRoleSetStmt{
		  Database: n.Database,
		  Role: convertRoleSpec(n.Role),
		  Setstmt: convertVariableSetStmt(n.Setstmt),
	}
}

func convertAlterRoleStmt(n *nodes.AlterRoleStmt) *pg.AlterRoleStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterRoleStmt{
		  Action: n.Action,
		  Options: convertList(n.Options),
		  Role: convertRoleSpec(n.Role),
	}
}

func convertAlterSeqStmt(n *nodes.AlterSeqStmt) *pg.AlterSeqStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterSeqStmt{
		  ForIdentity: n.ForIdentity,
		  MissingOk: n.MissingOk,
		  Options: convertList(n.Options),
		  Sequence: convertRangeVar(n.Sequence),
	}
}

func convertAlterSubscriptionStmt(n *nodes.AlterSubscriptionStmt) *pg.AlterSubscriptionStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterSubscriptionStmt{
		  Conninfo: n.Conninfo,
		  Kind: pg.AlterSubscriptionType(n.Kind),
		  Options: convertList(n.Options),
		  Publication: convertList(n.Publication),
		  Subname: n.Subname,
	}
}

func convertAlterSystemStmt(n *nodes.AlterSystemStmt) *pg.AlterSystemStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterSystemStmt{
		  Setstmt: convertVariableSetStmt(n.Setstmt),
	}
}

func convertAlterTSConfigurationStmt(n *nodes.AlterTSConfigurationStmt) *pg.AlterTSConfigurationStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterTSConfigurationStmt{
		  Cfgname: convertList(n.Cfgname),
		  Dicts: convertList(n.Dicts),
		  Kind: pg.AlterTSConfigType(n.Kind),
		  MissingOk: n.MissingOk,
		  Override: n.Override,
		  Replace: n.Replace,
		  Tokentype: convertList(n.Tokentype),
	}
}

func convertAlterTSDictionaryStmt(n *nodes.AlterTSDictionaryStmt) *pg.AlterTSDictionaryStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterTSDictionaryStmt{
		  Dictname: convertList(n.Dictname),
		  Options: convertList(n.Options),
	}
}

func convertAlterTableCmd(n *nodes.AlterTableCmd) *pg.AlterTableCmd {
	if n == nil {
		return nil
	}
	return &pg.AlterTableCmd{
		  Behavior: pg.DropBehavior(n.Behavior),
		  Def: convertNode(n.Def),
		  MissingOk: n.MissingOk,
		  Name: n.Name,
		  Newowner: convertRoleSpec(n.Newowner),
		  Subtype: pg.AlterTableType(n.Subtype),
	}
}

func convertAlterTableMoveAllStmt(n *nodes.AlterTableMoveAllStmt) *pg.AlterTableMoveAllStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterTableMoveAllStmt{
		  NewTablespacename: n.NewTablespacename,
		  Nowait: n.Nowait,
		  Objtype: pg.ObjectType(n.Objtype),
		  OrigTablespacename: n.OrigTablespacename,
		  Roles: convertList(n.Roles),
	}
}

func convertAlterTableSpaceOptionsStmt(n *nodes.AlterTableSpaceOptionsStmt) *pg.AlterTableSpaceOptionsStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterTableSpaceOptionsStmt{
		  IsReset: n.IsReset,
		  Options: convertList(n.Options),
		  Tablespacename: n.Tablespacename,
	}
}

func convertAlterTableStmt(n *nodes.AlterTableStmt) *pg.AlterTableStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterTableStmt{
		  Cmds: convertList(n.Cmds),
		  MissingOk: n.MissingOk,
		  Relation: convertRangeVar(n.Relation),
		  Relkind: pg.ObjectType(n.Relkind),
	}
}

func convertAlterUserMappingStmt(n *nodes.AlterUserMappingStmt) *pg.AlterUserMappingStmt {
	if n == nil {
		return nil
	}
	return &pg.AlterUserMappingStmt{
		  Options: convertList(n.Options),
		  Servername: n.Servername,
		  User: convertRoleSpec(n.User),
	}
}

func convertAlternativeSubPlan(n *nodes.AlternativeSubPlan) *pg.AlternativeSubPlan {
	if n == nil {
		return nil
	}
	return &pg.AlternativeSubPlan{
		  Subplans: convertList(n.Subplans),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertArrayCoerceExpr(n *nodes.ArrayCoerceExpr) *pg.ArrayCoerceExpr {
	if n == nil {
		return nil
	}
	return &pg.ArrayCoerceExpr{
		  Arg: convertNode(n.Arg),
		  Coerceformat: pg.CoercionForm(n.Coerceformat),
		  Elemfuncid: pg.Oid(n.Elemfuncid),
		  IsExplicit: n.IsExplicit,
		  Location: n.Location,
		  Resultcollid: pg.Oid(n.Resultcollid),
		  Resulttype: pg.Oid(n.Resulttype),
		  Resulttypmod: n.Resulttypmod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertArrayExpr(n *nodes.ArrayExpr) *pg.ArrayExpr {
	if n == nil {
		return nil
	}
	return &pg.ArrayExpr{
		  ArrayCollid: pg.Oid(n.ArrayCollid),
		  ArrayTypeid: pg.Oid(n.ArrayTypeid),
		  ElementTypeid: pg.Oid(n.ElementTypeid),
		  Elements: convertList(n.Elements),
		  Location: n.Location,
		  Multidims: n.Multidims,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertArrayRef(n *nodes.ArrayRef) *pg.ArrayRef {
	if n == nil {
		return nil
	}
	return &pg.ArrayRef{
		  Refarraytype: pg.Oid(n.Refarraytype),
		  Refassgnexpr: convertNode(n.Refassgnexpr),
		  Refcollid: pg.Oid(n.Refcollid),
		  Refelemtype: pg.Oid(n.Refelemtype),
		  Refexpr: convertNode(n.Refexpr),
		  Reflowerindexpr: convertList(n.Reflowerindexpr),
		  Reftypmod: n.Reftypmod,
		  Refupperindexpr: convertList(n.Refupperindexpr),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertBitString(n *nodes.BitString) *pg.BitString {
	if n == nil {
		return nil
	}
	return &pg.BitString{
		  Str: n.Str,
	}
}

func convertBlockIdData(n *nodes.BlockIdData) *pg.BlockIdData {
	if n == nil {
		return nil
	}
	return &pg.BlockIdData{
		  BiHi: n.BiHi,
		  BiLo: n.BiLo,
	}
}

func convertBoolExpr(n *nodes.BoolExpr) *pg.BoolExpr {
	if n == nil {
		return nil
	}
	return &pg.BoolExpr{
		  Args: convertList(n.Args),
		  Boolop: pg.BoolExprType(n.Boolop),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertBooleanTest(n *nodes.BooleanTest) *pg.BooleanTest {
	if n == nil {
		return nil
	}
	return &pg.BooleanTest{
		  Arg: convertNode(n.Arg),
		  Booltesttype: pg.BoolTestType(n.Booltesttype),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCaseExpr(n *nodes.CaseExpr) *pg.CaseExpr {
	if n == nil {
		return nil
	}
	return &pg.CaseExpr{
		  Arg: convertNode(n.Arg),
		  Args: convertList(n.Args),
		  Casecollid: pg.Oid(n.Casecollid),
		  Casetype: pg.Oid(n.Casetype),
		  Defresult: convertNode(n.Defresult),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCaseTestExpr(n *nodes.CaseTestExpr) *pg.CaseTestExpr {
	if n == nil {
		return nil
	}
	return &pg.CaseTestExpr{
		  Collation: pg.Oid(n.Collation),
		  TypeId: pg.Oid(n.TypeId),
		  TypeMod: n.TypeMod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCaseWhen(n *nodes.CaseWhen) *pg.CaseWhen {
	if n == nil {
		return nil
	}
	return &pg.CaseWhen{
		  Expr: convertNode(n.Expr),
		  Location: n.Location,
		  Result: convertNode(n.Result),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCheckPointStmt(n *nodes.CheckPointStmt) *pg.CheckPointStmt {
	if n == nil {
		return nil
	}
	return &pg.CheckPointStmt{
	}
}

func convertClosePortalStmt(n *nodes.ClosePortalStmt) *pg.ClosePortalStmt {
	if n == nil {
		return nil
	}
	return &pg.ClosePortalStmt{
		  Portalname: n.Portalname,
	}
}

func convertClusterStmt(n *nodes.ClusterStmt) *pg.ClusterStmt {
	if n == nil {
		return nil
	}
	return &pg.ClusterStmt{
		  Indexname: n.Indexname,
		  Relation: convertRangeVar(n.Relation),
		  Verbose: n.Verbose,
	}
}

func convertCoalesceExpr(n *nodes.CoalesceExpr) *pg.CoalesceExpr {
	if n == nil {
		return nil
	}
	return &pg.CoalesceExpr{
		  Args: convertList(n.Args),
		  Coalescecollid: pg.Oid(n.Coalescecollid),
		  Coalescetype: pg.Oid(n.Coalescetype),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCoerceToDomain(n *nodes.CoerceToDomain) *pg.CoerceToDomain {
	if n == nil {
		return nil
	}
	return &pg.CoerceToDomain{
		  Arg: convertNode(n.Arg),
		  Coercionformat: pg.CoercionForm(n.Coercionformat),
		  Location: n.Location,
		  Resultcollid: pg.Oid(n.Resultcollid),
		  Resulttype: pg.Oid(n.Resulttype),
		  Resulttypmod: n.Resulttypmod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCoerceToDomainValue(n *nodes.CoerceToDomainValue) *pg.CoerceToDomainValue {
	if n == nil {
		return nil
	}
	return &pg.CoerceToDomainValue{
		  Collation: pg.Oid(n.Collation),
		  Location: n.Location,
		  TypeId: pg.Oid(n.TypeId),
		  TypeMod: n.TypeMod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCoerceViaIO(n *nodes.CoerceViaIO) *pg.CoerceViaIO {
	if n == nil {
		return nil
	}
	return &pg.CoerceViaIO{
		  Arg: convertNode(n.Arg),
		  Coerceformat: pg.CoercionForm(n.Coerceformat),
		  Location: n.Location,
		  Resultcollid: pg.Oid(n.Resultcollid),
		  Resulttype: pg.Oid(n.Resulttype),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCollateClause(n *nodes.CollateClause) *pg.CollateClause {
	if n == nil {
		return nil
	}
	return &pg.CollateClause{
		  Arg: convertNode(n.Arg),
		  Collname: convertList(n.Collname),
		  Location: n.Location,
	}
}

func convertCollateExpr(n *nodes.CollateExpr) *pg.CollateExpr {
	if n == nil {
		return nil
	}
	return &pg.CollateExpr{
		  Arg: convertNode(n.Arg),
		  CollOid: pg.Oid(n.CollOid),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertColumnDef(n *nodes.ColumnDef) *pg.ColumnDef {
	if n == nil {
		return nil
	}
	return &pg.ColumnDef{
		  CollClause: convertCollateClause(n.CollClause),
		  CollOid: pg.Oid(n.CollOid),
		  Colname: n.Colname,
		  Constraints: convertList(n.Constraints),
		  CookedDefault: convertNode(n.CookedDefault),
		  Fdwoptions: convertList(n.Fdwoptions),
		  Identity: n.Identity,
		  Inhcount: n.Inhcount,
		  IsFromParent: n.IsFromParent,
		  IsFromType: n.IsFromType,
		  IsLocal: n.IsLocal,
		  IsNotNull: n.IsNotNull,
		  Location: n.Location,
		  RawDefault: convertNode(n.RawDefault),
		  Storage: n.Storage,
		  TypeName: convertTypeName(n.TypeName),
	}
}

func convertColumnRef(n *nodes.ColumnRef) *pg.ColumnRef {
	if n == nil {
		return nil
	}
	return &pg.ColumnRef{
		  Fields: convertList(n.Fields),
		  Location: n.Location,
	}
}

func convertCommentStmt(n *nodes.CommentStmt) *pg.CommentStmt {
	if n == nil {
		return nil
	}
	return &pg.CommentStmt{
		  Comment: n.Comment,
		  Object: convertNode(n.Object),
		  Objtype: pg.ObjectType(n.Objtype),
	}
}

func convertCommonTableExpr(n *nodes.CommonTableExpr) *pg.CommonTableExpr {
	if n == nil {
		return nil
	}
	return &pg.CommonTableExpr{
		  Aliascolnames: convertList(n.Aliascolnames),
		  Ctecolcollations: convertList(n.Ctecolcollations),
		  Ctecolnames: convertList(n.Ctecolnames),
		  Ctecoltypes: convertList(n.Ctecoltypes),
		  Ctecoltypmods: convertList(n.Ctecoltypmods),
		  Ctename: n.Ctename,
		  Ctequery: convertNode(n.Ctequery),
		  Cterecursive: n.Cterecursive,
		  Cterefcount: n.Cterefcount,
		  Location: n.Location,
	}
}

func convertCompositeTypeStmt(n *nodes.CompositeTypeStmt) *pg.CompositeTypeStmt {
	if n == nil {
		return nil
	}
	return &pg.CompositeTypeStmt{
		  Coldeflist: convertList(n.Coldeflist),
		  Typevar: convertRangeVar(n.Typevar),
	}
}

func convertConst(n *nodes.Const) *pg.Const {
	if n == nil {
		return nil
	}
	return &pg.Const{
		  Constbyval: n.Constbyval,
		  Constcollid: pg.Oid(n.Constcollid),
		  Constisnull: n.Constisnull,
		  Constlen: n.Constlen,
		  Consttype: pg.Oid(n.Consttype),
		  Consttypmod: n.Consttypmod,
		  Constvalue: pg.Datum(n.Constvalue),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertConstraint(n *nodes.Constraint) *pg.Constraint {
	if n == nil {
		return nil
	}
	return &pg.Constraint{
		  AccessMethod: n.AccessMethod,
		  Conname: n.Conname,
		  Contype: pg.ConstrType(n.Contype),
		  CookedExpr: n.CookedExpr,
		  Deferrable: n.Deferrable,
		  Exclusions: convertList(n.Exclusions),
		  FkAttrs: convertList(n.FkAttrs),
		  FkDelAction: n.FkDelAction,
		  FkMatchtype: n.FkMatchtype,
		  FkUpdAction: n.FkUpdAction,
		  GeneratedWhen: n.GeneratedWhen,
		  Indexname: n.Indexname,
		  Indexspace: n.Indexspace,
		  Initdeferred: n.Initdeferred,
		  InitiallyValid: n.InitiallyValid,
		  IsNoInherit: n.IsNoInherit,
		  Keys: convertList(n.Keys),
		  Location: n.Location,
		  OldConpfeqop: convertList(n.OldConpfeqop),
		  OldPktableOid: pg.Oid(n.OldPktableOid),
		  Options: convertList(n.Options),
		  PkAttrs: convertList(n.PkAttrs),
		  Pktable: convertRangeVar(n.Pktable),
		  RawExpr: convertNode(n.RawExpr),
		  SkipValidation: n.SkipValidation,
		  WhereClause: convertNode(n.WhereClause),
	}
}

func convertConstraintsSetStmt(n *nodes.ConstraintsSetStmt) *pg.ConstraintsSetStmt {
	if n == nil {
		return nil
	}
	return &pg.ConstraintsSetStmt{
		  Constraints: convertList(n.Constraints),
		  Deferred: n.Deferred,
	}
}

func convertConvertRowtypeExpr(n *nodes.ConvertRowtypeExpr) *pg.ConvertRowtypeExpr {
	if n == nil {
		return nil
	}
	return &pg.ConvertRowtypeExpr{
		  Arg: convertNode(n.Arg),
		  Convertformat: pg.CoercionForm(n.Convertformat),
		  Location: n.Location,
		  Resulttype: pg.Oid(n.Resulttype),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertCopyStmt(n *nodes.CopyStmt) *pg.CopyStmt {
	if n == nil {
		return nil
	}
	return &pg.CopyStmt{
		  Attlist: convertList(n.Attlist),
		  Filename: n.Filename,
		  IsFrom: n.IsFrom,
		  IsProgram: n.IsProgram,
		  Options: convertList(n.Options),
		  Query: convertNode(n.Query),
		  Relation: convertRangeVar(n.Relation),
	}
}

func convertCreateAmStmt(n *nodes.CreateAmStmt) *pg.CreateAmStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateAmStmt{
		  Amname: n.Amname,
		  Amtype: n.Amtype,
		  HandlerName: convertList(n.HandlerName),
	}
}

func convertCreateCastStmt(n *nodes.CreateCastStmt) *pg.CreateCastStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateCastStmt{
		  Context: pg.CoercionContext(n.Context),
		  Func: convertObjectWithArgs(n.Func),
		  Inout: n.Inout,
		  Sourcetype: convertTypeName(n.Sourcetype),
		  Targettype: convertTypeName(n.Targettype),
	}
}

func convertCreateConversionStmt(n *nodes.CreateConversionStmt) *pg.CreateConversionStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateConversionStmt{
		  ConversionName: convertList(n.ConversionName),
		  Def: n.Def,
		  ForEncodingName: n.ForEncodingName,
		  FuncName: convertList(n.FuncName),
		  ToEncodingName: n.ToEncodingName,
	}
}

func convertCreateDomainStmt(n *nodes.CreateDomainStmt) *pg.CreateDomainStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateDomainStmt{
		  CollClause: convertCollateClause(n.CollClause),
		  Constraints: convertList(n.Constraints),
		  Domainname: convertList(n.Domainname),
		  TypeName: convertTypeName(n.TypeName),
	}
}

func convertCreateEnumStmt(n *nodes.CreateEnumStmt) *pg.CreateEnumStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateEnumStmt{
		  TypeName: convertList(n.TypeName),
		  Vals: convertList(n.Vals),
	}
}

func convertCreateEventTrigStmt(n *nodes.CreateEventTrigStmt) *pg.CreateEventTrigStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateEventTrigStmt{
		  Eventname: n.Eventname,
		  Funcname: convertList(n.Funcname),
		  Trigname: n.Trigname,
		  Whenclause: convertList(n.Whenclause),
	}
}

func convertCreateExtensionStmt(n *nodes.CreateExtensionStmt) *pg.CreateExtensionStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateExtensionStmt{
		  Extname: n.Extname,
		  IfNotExists: n.IfNotExists,
		  Options: convertList(n.Options),
	}
}

func convertCreateFdwStmt(n *nodes.CreateFdwStmt) *pg.CreateFdwStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateFdwStmt{
		  Fdwname: n.Fdwname,
		  FuncOptions: convertList(n.FuncOptions),
		  Options: convertList(n.Options),
	}
}

func convertCreateForeignServerStmt(n *nodes.CreateForeignServerStmt) *pg.CreateForeignServerStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateForeignServerStmt{
		  Fdwname: n.Fdwname,
		  IfNotExists: n.IfNotExists,
		  Options: convertList(n.Options),
		  Servername: n.Servername,
		  Servertype: n.Servertype,
		  Version: n.Version,
	}
}

func convertCreateForeignTableStmt(n *nodes.CreateForeignTableStmt) *pg.CreateForeignTableStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateForeignTableStmt{
		  Base: convertCreateStmt(&n.Base),
		  Options: convertList(n.Options),
		  Servername: n.Servername,
	}
}

func convertCreateFunctionStmt(n *nodes.CreateFunctionStmt) *pg.CreateFunctionStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateFunctionStmt{
		  Funcname: convertList(n.Funcname),
		  Options: convertList(n.Options),
		  Parameters: convertList(n.Parameters),
		  Replace: n.Replace,
		  ReturnType: convertTypeName(n.ReturnType),
		  WithClause: convertList(n.WithClause),
	}
}

func convertCreateOpClassItem(n *nodes.CreateOpClassItem) *pg.CreateOpClassItem {
	if n == nil {
		return nil
	}
	return &pg.CreateOpClassItem{
		  ClassArgs: convertList(n.ClassArgs),
		  Itemtype: n.Itemtype,
		  Name: convertObjectWithArgs(n.Name),
		  Number: n.Number,
		  OrderFamily: convertList(n.OrderFamily),
		  Storedtype: convertTypeName(n.Storedtype),
	}
}

func convertCreateOpClassStmt(n *nodes.CreateOpClassStmt) *pg.CreateOpClassStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateOpClassStmt{
		  Amname: n.Amname,
		  Datatype: convertTypeName(n.Datatype),
		  IsDefault: n.IsDefault,
		  Items: convertList(n.Items),
		  Opclassname: convertList(n.Opclassname),
		  Opfamilyname: convertList(n.Opfamilyname),
	}
}

func convertCreateOpFamilyStmt(n *nodes.CreateOpFamilyStmt) *pg.CreateOpFamilyStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateOpFamilyStmt{
		  Amname: n.Amname,
		  Opfamilyname: convertList(n.Opfamilyname),
	}
}

func convertCreatePLangStmt(n *nodes.CreatePLangStmt) *pg.CreatePLangStmt {
	if n == nil {
		return nil
	}
	return &pg.CreatePLangStmt{
		  Plhandler: convertList(n.Plhandler),
		  Plinline: convertList(n.Plinline),
		  Plname: n.Plname,
		  Pltrusted: n.Pltrusted,
		  Plvalidator: convertList(n.Plvalidator),
		  Replace: n.Replace,
	}
}

func convertCreatePolicyStmt(n *nodes.CreatePolicyStmt) *pg.CreatePolicyStmt {
	if n == nil {
		return nil
	}
	return &pg.CreatePolicyStmt{
		  CmdName: n.CmdName,
		  Permissive: n.Permissive,
		  PolicyName: n.PolicyName,
		  Qual: convertNode(n.Qual),
		  Roles: convertList(n.Roles),
		  Table: convertRangeVar(n.Table),
		  WithCheck: convertNode(n.WithCheck),
	}
}

func convertCreatePublicationStmt(n *nodes.CreatePublicationStmt) *pg.CreatePublicationStmt {
	if n == nil {
		return nil
	}
	return &pg.CreatePublicationStmt{
		  ForAllTables: n.ForAllTables,
		  Options: convertList(n.Options),
		  Pubname: n.Pubname,
		  Tables: convertList(n.Tables),
	}
}

func convertCreateRangeStmt(n *nodes.CreateRangeStmt) *pg.CreateRangeStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateRangeStmt{
		  Params: convertList(n.Params),
		  TypeName: convertList(n.TypeName),
	}
}

func convertCreateRoleStmt(n *nodes.CreateRoleStmt) *pg.CreateRoleStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateRoleStmt{
		  Options: convertList(n.Options),
		  Role: n.Role,
		  StmtType: pg.RoleStmtType(n.StmtType),
	}
}

func convertCreateSchemaStmt(n *nodes.CreateSchemaStmt) *pg.CreateSchemaStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateSchemaStmt{
		  Authrole: convertRoleSpec(n.Authrole),
		  IfNotExists: n.IfNotExists,
		  SchemaElts: convertList(n.SchemaElts),
		  Schemaname: n.Schemaname,
	}
}

func convertCreateSeqStmt(n *nodes.CreateSeqStmt) *pg.CreateSeqStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateSeqStmt{
		  ForIdentity: n.ForIdentity,
		  IfNotExists: n.IfNotExists,
		  Options: convertList(n.Options),
		  OwnerId: pg.Oid(n.OwnerId),
		  Sequence: convertRangeVar(n.Sequence),
	}
}

func convertCreateStatsStmt(n *nodes.CreateStatsStmt) *pg.CreateStatsStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateStatsStmt{
		  Defnames: convertList(n.Defnames),
		  Exprs: convertList(n.Exprs),
		  IfNotExists: n.IfNotExists,
		  Relations: convertList(n.Relations),
		  StatTypes: convertList(n.StatTypes),
	}
}

func convertCreateStmt(n *nodes.CreateStmt) *pg.CreateStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateStmt{
		  Constraints: convertList(n.Constraints),
		  IfNotExists: n.IfNotExists,
		  InhRelations: convertList(n.InhRelations),
		  OfTypename: convertTypeName(n.OfTypename),
		  Oncommit: pg.OnCommitAction(n.Oncommit),
		  Options: convertList(n.Options),
		  Partbound: convertPartitionBoundSpec(n.Partbound),
		  Partspec: convertPartitionSpec(n.Partspec),
		  Relation: convertRangeVar(n.Relation),
		  TableElts: convertList(n.TableElts),
		  Tablespacename: n.Tablespacename,
	}
}

func convertCreateSubscriptionStmt(n *nodes.CreateSubscriptionStmt) *pg.CreateSubscriptionStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateSubscriptionStmt{
		  Conninfo: n.Conninfo,
		  Options: convertList(n.Options),
		  Publication: convertList(n.Publication),
		  Subname: n.Subname,
	}
}

func convertCreateTableAsStmt(n *nodes.CreateTableAsStmt) *pg.CreateTableAsStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateTableAsStmt{
		  IfNotExists: n.IfNotExists,
		  Into: convertIntoClause(n.Into),
		  IsSelectInto: n.IsSelectInto,
		  Query: convertNode(n.Query),
		  Relkind: pg.ObjectType(n.Relkind),
	}
}

func convertCreateTableSpaceStmt(n *nodes.CreateTableSpaceStmt) *pg.CreateTableSpaceStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateTableSpaceStmt{
		  Location: n.Location,
		  Options: convertList(n.Options),
		  Owner: convertRoleSpec(n.Owner),
		  Tablespacename: n.Tablespacename,
	}
}

func convertCreateTransformStmt(n *nodes.CreateTransformStmt) *pg.CreateTransformStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateTransformStmt{
		  Fromsql: convertObjectWithArgs(n.Fromsql),
		  Lang: n.Lang,
		  Replace: n.Replace,
		  Tosql: convertObjectWithArgs(n.Tosql),
		  TypeName: convertTypeName(n.TypeName),
	}
}

func convertCreateTrigStmt(n *nodes.CreateTrigStmt) *pg.CreateTrigStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateTrigStmt{
		  Args: convertList(n.Args),
		  Columns: convertList(n.Columns),
		  Constrrel: convertRangeVar(n.Constrrel),
		  Deferrable: n.Deferrable,
		  Events: n.Events,
		  Funcname: convertList(n.Funcname),
		  Initdeferred: n.Initdeferred,
		  Isconstraint: n.Isconstraint,
		  Relation: convertRangeVar(n.Relation),
		  Row: n.Row,
		  Timing: n.Timing,
		  TransitionRels: convertList(n.TransitionRels),
		  Trigname: n.Trigname,
		  WhenClause: convertNode(n.WhenClause),
	}
}

func convertCreateUserMappingStmt(n *nodes.CreateUserMappingStmt) *pg.CreateUserMappingStmt {
	if n == nil {
		return nil
	}
	return &pg.CreateUserMappingStmt{
		  IfNotExists: n.IfNotExists,
		  Options: convertList(n.Options),
		  Servername: n.Servername,
		  User: convertRoleSpec(n.User),
	}
}

func convertCreatedbStmt(n *nodes.CreatedbStmt) *pg.CreatedbStmt {
	if n == nil {
		return nil
	}
	return &pg.CreatedbStmt{
		  Dbname: n.Dbname,
		  Options: convertList(n.Options),
	}
}

func convertCurrentOfExpr(n *nodes.CurrentOfExpr) *pg.CurrentOfExpr {
	if n == nil {
		return nil
	}
	return &pg.CurrentOfExpr{
		  CursorName: n.CursorName,
		  CursorParam: n.CursorParam,
		  Cvarno: pg.Index(n.Cvarno),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertDeallocateStmt(n *nodes.DeallocateStmt) *pg.DeallocateStmt {
	if n == nil {
		return nil
	}
	return &pg.DeallocateStmt{
		  Name: n.Name,
	}
}

func convertDeclareCursorStmt(n *nodes.DeclareCursorStmt) *pg.DeclareCursorStmt {
	if n == nil {
		return nil
	}
	return &pg.DeclareCursorStmt{
		  Options: n.Options,
		  Portalname: n.Portalname,
		  Query: convertNode(n.Query),
	}
}

func convertDefElem(n *nodes.DefElem) *pg.DefElem {
	if n == nil {
		return nil
	}
	return &pg.DefElem{
		  Arg: convertNode(n.Arg),
		  Defaction: pg.DefElemAction(n.Defaction),
		  Defname: n.Defname,
		  Defnamespace: n.Defnamespace,
		  Location: n.Location,
	}
}

func convertDefineStmt(n *nodes.DefineStmt) *pg.DefineStmt {
	if n == nil {
		return nil
	}
	return &pg.DefineStmt{
		  Args: convertList(n.Args),
		  Definition: convertList(n.Definition),
		  Defnames: convertList(n.Defnames),
		  IfNotExists: n.IfNotExists,
		  Kind: pg.ObjectType(n.Kind),
		  Oldstyle: n.Oldstyle,
	}
}

func convertDeleteStmt(n *nodes.DeleteStmt) *pg.DeleteStmt {
	if n == nil {
		return nil
	}
	return &pg.DeleteStmt{
		  Relation: convertRangeVar(n.Relation),
		  ReturningList: convertList(n.ReturningList),
		  UsingClause: convertList(n.UsingClause),
		  WhereClause: convertNode(n.WhereClause),
		  WithClause: convertWithClause(n.WithClause),
	}
}

func convertDiscardStmt(n *nodes.DiscardStmt) *pg.DiscardStmt {
	if n == nil {
		return nil
	}
	return &pg.DiscardStmt{
		  Target: pg.DiscardMode(n.Target),
	}
}

func convertDoStmt(n *nodes.DoStmt) *pg.DoStmt {
	if n == nil {
		return nil
	}
	return &pg.DoStmt{
		  Args: convertList(n.Args),
	}
}

func convertDropOwnedStmt(n *nodes.DropOwnedStmt) *pg.DropOwnedStmt {
	if n == nil {
		return nil
	}
	return &pg.DropOwnedStmt{
		  Behavior: pg.DropBehavior(n.Behavior),
		  Roles: convertList(n.Roles),
	}
}

func convertDropRoleStmt(n *nodes.DropRoleStmt) *pg.DropRoleStmt {
	if n == nil {
		return nil
	}
	return &pg.DropRoleStmt{
		  MissingOk: n.MissingOk,
		  Roles: convertList(n.Roles),
	}
}

func convertDropStmt(n *nodes.DropStmt) *pg.DropStmt {
	if n == nil {
		return nil
	}
	return &pg.DropStmt{
		  Behavior: pg.DropBehavior(n.Behavior),
		  Concurrent: n.Concurrent,
		  MissingOk: n.MissingOk,
		  Objects: convertList(n.Objects),
		  RemoveType: pg.ObjectType(n.RemoveType),
	}
}

func convertDropSubscriptionStmt(n *nodes.DropSubscriptionStmt) *pg.DropSubscriptionStmt {
	if n == nil {
		return nil
	}
	return &pg.DropSubscriptionStmt{
		  Behavior: pg.DropBehavior(n.Behavior),
		  MissingOk: n.MissingOk,
		  Subname: n.Subname,
	}
}

func convertDropTableSpaceStmt(n *nodes.DropTableSpaceStmt) *pg.DropTableSpaceStmt {
	if n == nil {
		return nil
	}
	return &pg.DropTableSpaceStmt{
		  MissingOk: n.MissingOk,
		  Tablespacename: n.Tablespacename,
	}
}

func convertDropUserMappingStmt(n *nodes.DropUserMappingStmt) *pg.DropUserMappingStmt {
	if n == nil {
		return nil
	}
	return &pg.DropUserMappingStmt{
		  MissingOk: n.MissingOk,
		  Servername: n.Servername,
		  User: convertRoleSpec(n.User),
	}
}

func convertDropdbStmt(n *nodes.DropdbStmt) *pg.DropdbStmt {
	if n == nil {
		return nil
	}
	return &pg.DropdbStmt{
		  Dbname: n.Dbname,
		  MissingOk: n.MissingOk,
	}
}

func convertExecuteStmt(n *nodes.ExecuteStmt) *pg.ExecuteStmt {
	if n == nil {
		return nil
	}
	return &pg.ExecuteStmt{
		  Name: n.Name,
		  Params: convertList(n.Params),
	}
}

func convertExplainStmt(n *nodes.ExplainStmt) *pg.ExplainStmt {
	if n == nil {
		return nil
	}
	return &pg.ExplainStmt{
		  Options: convertList(n.Options),
		  Query: convertNode(n.Query),
	}
}

func convertExpr(n *nodes.Expr) *pg.Expr {
	if n == nil {
		return nil
	}
	return &pg.Expr{
	}
}

func convertFetchStmt(n *nodes.FetchStmt) *pg.FetchStmt {
	if n == nil {
		return nil
	}
	return &pg.FetchStmt{
		  Direction: pg.FetchDirection(n.Direction),
		  HowMany: n.HowMany,
		  Ismove: n.Ismove,
		  Portalname: n.Portalname,
	}
}

func convertFieldSelect(n *nodes.FieldSelect) *pg.FieldSelect {
	if n == nil {
		return nil
	}
	return &pg.FieldSelect{
		  Arg: convertNode(n.Arg),
		  Fieldnum: pg.AttrNumber(n.Fieldnum),
		  Resultcollid: pg.Oid(n.Resultcollid),
		  Resulttype: pg.Oid(n.Resulttype),
		  Resulttypmod: n.Resulttypmod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertFieldStore(n *nodes.FieldStore) *pg.FieldStore {
	if n == nil {
		return nil
	}
	return &pg.FieldStore{
		  Arg: convertNode(n.Arg),
		  Fieldnums: convertList(n.Fieldnums),
		  Newvals: convertList(n.Newvals),
		  Resulttype: pg.Oid(n.Resulttype),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertFloat(n *nodes.Float) *pg.Float {
	if n == nil {
		return nil
	}
	return &pg.Float{
		  Str: n.Str,
	}
}

func convertFromExpr(n *nodes.FromExpr) *pg.FromExpr {
	if n == nil {
		return nil
	}
	return &pg.FromExpr{
		  Fromlist: convertList(n.Fromlist),
		  Quals: convertNode(n.Quals),
	}
}

func convertFuncCall(n *nodes.FuncCall) *pg.FuncCall {
	if n == nil {
		return nil
	}
	return &pg.FuncCall{
		  AggDistinct: n.AggDistinct,
		  AggFilter: convertNode(n.AggFilter),
		  AggOrder: convertList(n.AggOrder),
		  AggStar: n.AggStar,
		  AggWithinGroup: n.AggWithinGroup,
		  Args: convertList(n.Args),
		  FuncVariadic: n.FuncVariadic,
		  Funcname: convertList(n.Funcname),
		  Location: n.Location,
		  Over: convertWindowDef(n.Over),
	}
}

func convertFuncExpr(n *nodes.FuncExpr) *pg.FuncExpr {
	if n == nil {
		return nil
	}
	return &pg.FuncExpr{
		  Args: convertList(n.Args),
		  Funccollid: pg.Oid(n.Funccollid),
		  Funcformat: pg.CoercionForm(n.Funcformat),
		  Funcid: pg.Oid(n.Funcid),
		  Funcresulttype: pg.Oid(n.Funcresulttype),
		  Funcretset: n.Funcretset,
		  Funcvariadic: n.Funcvariadic,
		  Inputcollid: pg.Oid(n.Inputcollid),
		  Location: n.Location,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertFunctionParameter(n *nodes.FunctionParameter) *pg.FunctionParameter {
	if n == nil {
		return nil
	}
	return &pg.FunctionParameter{
		  ArgType: convertTypeName(n.ArgType),
		  Defexpr: convertNode(n.Defexpr),
		  Mode: pg.FunctionParameterMode(n.Mode),
		  Name: n.Name,
	}
}

func convertGrantRoleStmt(n *nodes.GrantRoleStmt) *pg.GrantRoleStmt {
	if n == nil {
		return nil
	}
	return &pg.GrantRoleStmt{
		  AdminOpt: n.AdminOpt,
		  Behavior: pg.DropBehavior(n.Behavior),
		  GrantedRoles: convertList(n.GrantedRoles),
		  GranteeRoles: convertList(n.GranteeRoles),
		  Grantor: convertRoleSpec(n.Grantor),
		  IsGrant: n.IsGrant,
	}
}

func convertGrantStmt(n *nodes.GrantStmt) *pg.GrantStmt {
	if n == nil {
		return nil
	}
	return &pg.GrantStmt{
		  Behavior: pg.DropBehavior(n.Behavior),
		  GrantOption: n.GrantOption,
		  Grantees: convertList(n.Grantees),
		  IsGrant: n.IsGrant,
		  Objects: convertList(n.Objects),
		  Objtype: pg.GrantObjectType(n.Objtype),
		  Privileges: convertList(n.Privileges),
		  Targtype: pg.GrantTargetType(n.Targtype),
	}
}

func convertGroupingFunc(n *nodes.GroupingFunc) *pg.GroupingFunc {
	if n == nil {
		return nil
	}
	return &pg.GroupingFunc{
		  Agglevelsup: pg.Index(n.Agglevelsup),
		  Args: convertList(n.Args),
		  Cols: convertList(n.Cols),
		  Location: n.Location,
		  Refs: convertList(n.Refs),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertGroupingSet(n *nodes.GroupingSet) *pg.GroupingSet {
	if n == nil {
		return nil
	}
	return &pg.GroupingSet{
		  Content: convertList(n.Content),
		  Kind: pg.GroupingSetKind(n.Kind),
		  Location: n.Location,
	}
}

func convertImportForeignSchemaStmt(n *nodes.ImportForeignSchemaStmt) *pg.ImportForeignSchemaStmt {
	if n == nil {
		return nil
	}
	return &pg.ImportForeignSchemaStmt{
		  ListType: pg.ImportForeignSchemaType(n.ListType),
		  LocalSchema: n.LocalSchema,
		  Options: convertList(n.Options),
		  RemoteSchema: n.RemoteSchema,
		  ServerName: n.ServerName,
		  TableList: convertList(n.TableList),
	}
}

func convertIndexElem(n *nodes.IndexElem) *pg.IndexElem {
	if n == nil {
		return nil
	}
	return &pg.IndexElem{
		  Collation: convertList(n.Collation),
		  Expr: convertNode(n.Expr),
		  Indexcolname: n.Indexcolname,
		  Name: n.Name,
		  NullsOrdering: pg.SortByNulls(n.NullsOrdering),
		  Opclass: convertList(n.Opclass),
		  Ordering: pg.SortByDir(n.Ordering),
	}
}

func convertIndexStmt(n *nodes.IndexStmt) *pg.IndexStmt {
	if n == nil {
		return nil
	}
	return &pg.IndexStmt{
		  AccessMethod: n.AccessMethod,
		  Concurrent: n.Concurrent,
		  Deferrable: n.Deferrable,
		  ExcludeOpNames: convertList(n.ExcludeOpNames),
		  Idxcomment: n.Idxcomment,
		  Idxname: n.Idxname,
		  IfNotExists: n.IfNotExists,
		  IndexOid: pg.Oid(n.IndexOid),
		  IndexParams: convertList(n.IndexParams),
		  Initdeferred: n.Initdeferred,
		  Isconstraint: n.Isconstraint,
		  OldNode: pg.Oid(n.OldNode),
		  Options: convertList(n.Options),
		  Primary: n.Primary,
		  Relation: convertRangeVar(n.Relation),
		  TableSpace: n.TableSpace,
		  Transformed: n.Transformed,
		  Unique: n.Unique,
		  WhereClause: convertNode(n.WhereClause),
	}
}

func convertInferClause(n *nodes.InferClause) *pg.InferClause {
	if n == nil {
		return nil
	}
	return &pg.InferClause{
		  Conname: n.Conname,
		  IndexElems: convertList(n.IndexElems),
		  Location: n.Location,
		  WhereClause: convertNode(n.WhereClause),
	}
}

func convertInferenceElem(n *nodes.InferenceElem) *pg.InferenceElem {
	if n == nil {
		return nil
	}
	return &pg.InferenceElem{
		  Expr: convertNode(n.Expr),
		  Infercollid: pg.Oid(n.Infercollid),
		  Inferopclass: pg.Oid(n.Inferopclass),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertInlineCodeBlock(n *nodes.InlineCodeBlock) *pg.InlineCodeBlock {
	if n == nil {
		return nil
	}
	return &pg.InlineCodeBlock{
		  LangIsTrusted: n.LangIsTrusted,
		  LangOid: pg.Oid(n.LangOid),
		  SourceText: n.SourceText,
	}
}

func convertInsertStmt(n *nodes.InsertStmt) *pg.InsertStmt {
	if n == nil {
		return nil
	}
	return &pg.InsertStmt{
		  Cols: convertList(n.Cols),
		  OnConflictClause: convertOnConflictClause(n.OnConflictClause),
		  Override: pg.OverridingKind(n.Override),
		  Relation: convertRangeVar(n.Relation),
		  ReturningList: convertList(n.ReturningList),
		  SelectStmt: convertNode(n.SelectStmt),
		  WithClause: convertWithClause(n.WithClause),
	}
}

func convertInteger(n *nodes.Integer) *pg.Integer {
	if n == nil {
		return nil
	}
	return &pg.Integer{
		  Ival: n.Ival,
	}
}

func convertIntoClause(n *nodes.IntoClause) *pg.IntoClause {
	if n == nil {
		return nil
	}
	return &pg.IntoClause{
		  ColNames: convertList(n.ColNames),
		  OnCommit: pg.OnCommitAction(n.OnCommit),
		  Options: convertList(n.Options),
		  Rel: convertRangeVar(n.Rel),
		  SkipData: n.SkipData,
		  TableSpaceName: n.TableSpaceName,
		  ViewQuery: convertNode(n.ViewQuery),
	}
}

func convertJoinExpr(n *nodes.JoinExpr) *pg.JoinExpr {
	if n == nil {
		return nil
	}
	return &pg.JoinExpr{
		  Alias: convertAlias(n.Alias),
		  IsNatural: n.IsNatural,
		  Jointype: pg.JoinType(n.Jointype),
		  Larg: convertNode(n.Larg),
		  Quals: convertNode(n.Quals),
		  Rarg: convertNode(n.Rarg),
		  Rtindex: n.Rtindex,
		  UsingClause: convertList(n.UsingClause),
	}
}

func convertListenStmt(n *nodes.ListenStmt) *pg.ListenStmt {
	if n == nil {
		return nil
	}
	return &pg.ListenStmt{
		  Conditionname: n.Conditionname,
	}
}

func convertLoadStmt(n *nodes.LoadStmt) *pg.LoadStmt {
	if n == nil {
		return nil
	}
	return &pg.LoadStmt{
		  Filename: n.Filename,
	}
}

func convertLockStmt(n *nodes.LockStmt) *pg.LockStmt {
	if n == nil {
		return nil
	}
	return &pg.LockStmt{
		  Mode: n.Mode,
		  Nowait: n.Nowait,
		  Relations: convertList(n.Relations),
	}
}

func convertLockingClause(n *nodes.LockingClause) *pg.LockingClause {
	if n == nil {
		return nil
	}
	return &pg.LockingClause{
		  LockedRels: convertList(n.LockedRels),
		  Strength: pg.LockClauseStrength(n.Strength),
		  WaitPolicy: pg.LockWaitPolicy(n.WaitPolicy),
	}
}

func convertMinMaxExpr(n *nodes.MinMaxExpr) *pg.MinMaxExpr {
	if n == nil {
		return nil
	}
	return &pg.MinMaxExpr{
		  Args: convertList(n.Args),
		  Inputcollid: pg.Oid(n.Inputcollid),
		  Location: n.Location,
		  Minmaxcollid: pg.Oid(n.Minmaxcollid),
		  Minmaxtype: pg.Oid(n.Minmaxtype),
		  Op: pg.MinMaxOp(n.Op),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertMultiAssignRef(n *nodes.MultiAssignRef) *pg.MultiAssignRef {
	if n == nil {
		return nil
	}
	return &pg.MultiAssignRef{
		  Colno: n.Colno,
		  Ncolumns: n.Ncolumns,
		  Source: convertNode(n.Source),
	}
}

func convertNamedArgExpr(n *nodes.NamedArgExpr) *pg.NamedArgExpr {
	if n == nil {
		return nil
	}
	return &pg.NamedArgExpr{
		  Arg: convertNode(n.Arg),
		  Argnumber: n.Argnumber,
		  Location: n.Location,
		  Name: n.Name,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertNextValueExpr(n *nodes.NextValueExpr) *pg.NextValueExpr {
	if n == nil {
		return nil
	}
	return &pg.NextValueExpr{
		  Seqid: pg.Oid(n.Seqid),
		  TypeId: pg.Oid(n.TypeId),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertNotifyStmt(n *nodes.NotifyStmt) *pg.NotifyStmt {
	if n == nil {
		return nil
	}
	return &pg.NotifyStmt{
		  Conditionname: n.Conditionname,
		  Payload: n.Payload,
	}
}

func convertNull(n *nodes.Null) *pg.Null {
	if n == nil {
		return nil
	}
	return &pg.Null{
	}
}

func convertNullTest(n *nodes.NullTest) *pg.NullTest {
	if n == nil {
		return nil
	}
	return &pg.NullTest{
		  Arg: convertNode(n.Arg),
		  Argisrow: n.Argisrow,
		  Location: n.Location,
		  Nulltesttype: pg.NullTestType(n.Nulltesttype),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertObjectWithArgs(n *nodes.ObjectWithArgs) *pg.ObjectWithArgs {
	if n == nil {
		return nil
	}
	return &pg.ObjectWithArgs{
		  ArgsUnspecified: n.ArgsUnspecified,
		  Objargs: convertList(n.Objargs),
		  Objname: convertList(n.Objname),
	}
}

func convertOnConflictClause(n *nodes.OnConflictClause) *pg.OnConflictClause {
	if n == nil {
		return nil
	}
	return &pg.OnConflictClause{
		  Action: pg.OnConflictAction(n.Action),
		  Infer: convertInferClause(n.Infer),
		  Location: n.Location,
		  TargetList: convertList(n.TargetList),
		  WhereClause: convertNode(n.WhereClause),
	}
}

func convertOnConflictExpr(n *nodes.OnConflictExpr) *pg.OnConflictExpr {
	if n == nil {
		return nil
	}
	return &pg.OnConflictExpr{
		  Action: pg.OnConflictAction(n.Action),
		  ArbiterElems: convertList(n.ArbiterElems),
		  ArbiterWhere: convertNode(n.ArbiterWhere),
		  Constraint: pg.Oid(n.Constraint),
		  ExclRelIndex: n.ExclRelIndex,
		  ExclRelTlist: convertList(n.ExclRelTlist),
		  OnConflictSet: convertList(n.OnConflictSet),
		  OnConflictWhere: convertNode(n.OnConflictWhere),
	}
}

func convertOpExpr(n *nodes.OpExpr) *pg.OpExpr {
	if n == nil {
		return nil
	}
	return &pg.OpExpr{
		  Args: convertList(n.Args),
		  Inputcollid: pg.Oid(n.Inputcollid),
		  Location: n.Location,
		  Opcollid: pg.Oid(n.Opcollid),
		  Opfuncid: pg.Oid(n.Opfuncid),
		  Opno: pg.Oid(n.Opno),
		  Opresulttype: pg.Oid(n.Opresulttype),
		  Opretset: n.Opretset,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertParam(n *nodes.Param) *pg.Param {
	if n == nil {
		return nil
	}
	return &pg.Param{
		  Location: n.Location,
		  Paramcollid: pg.Oid(n.Paramcollid),
		  Paramid: n.Paramid,
		  Paramkind: pg.ParamKind(n.Paramkind),
		  Paramtype: pg.Oid(n.Paramtype),
		  Paramtypmod: n.Paramtypmod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertParamExecData(n *nodes.ParamExecData) *pg.ParamExecData {
	if n == nil {
		return nil
	}
	return &pg.ParamExecData{
		  ExecPlan: &ast.TODO{},
		  Isnull: n.Isnull,
		  Value: pg.Datum(n.Value),
	}
}

func convertParamExternData(n *nodes.ParamExternData) *pg.ParamExternData {
	if n == nil {
		return nil
	}
	return &pg.ParamExternData{
		  Isnull: n.Isnull,
		  Pflags: n.Pflags,
		  Ptype: pg.Oid(n.Ptype),
		  Value: pg.Datum(n.Value),
	}
}

func convertParamListInfoData(n *nodes.ParamListInfoData) *pg.ParamListInfoData {
	if n == nil {
		return nil
	}
	return &pg.ParamListInfoData{
		  NumParams: n.NumParams,
		  ParamFetchArg: &ast.TODO{},
		  ParamMask: n.ParamMask,
		  ParserSetupArg: &ast.TODO{},
	}
}

func convertParamRef(n *nodes.ParamRef) *pg.ParamRef {
	if n == nil {
		return nil
	}
	return &pg.ParamRef{
		  Location: n.Location,
		  Number: n.Number,
	}
}

func convertPartitionBoundSpec(n *nodes.PartitionBoundSpec) *pg.PartitionBoundSpec {
	if n == nil {
		return nil
	}
	return &pg.PartitionBoundSpec{
		  Listdatums: convertList(n.Listdatums),
		  Location: n.Location,
		  Lowerdatums: convertList(n.Lowerdatums),
		  Strategy: n.Strategy,
		  Upperdatums: convertList(n.Upperdatums),
	}
}

func convertPartitionCmd(n *nodes.PartitionCmd) *pg.PartitionCmd {
	if n == nil {
		return nil
	}
	return &pg.PartitionCmd{
		  Bound: convertPartitionBoundSpec(n.Bound),
		  Name: convertRangeVar(n.Name),
	}
}

func convertPartitionElem(n *nodes.PartitionElem) *pg.PartitionElem {
	if n == nil {
		return nil
	}
	return &pg.PartitionElem{
		  Collation: convertList(n.Collation),
		  Expr: convertNode(n.Expr),
		  Location: n.Location,
		  Name: n.Name,
		  Opclass: convertList(n.Opclass),
	}
}

func convertPartitionRangeDatum(n *nodes.PartitionRangeDatum) *pg.PartitionRangeDatum {
	if n == nil {
		return nil
	}
	return &pg.PartitionRangeDatum{
		  Kind: pg.PartitionRangeDatumKind(n.Kind),
		  Location: n.Location,
		  Value: convertNode(n.Value),
	}
}

func convertPartitionSpec(n *nodes.PartitionSpec) *pg.PartitionSpec {
	if n == nil {
		return nil
	}
	return &pg.PartitionSpec{
		  Location: n.Location,
		  PartParams: convertList(n.PartParams),
		  Strategy: n.Strategy,
	}
}

func convertPrepareStmt(n *nodes.PrepareStmt) *pg.PrepareStmt {
	if n == nil {
		return nil
	}
	return &pg.PrepareStmt{
		  Argtypes: convertList(n.Argtypes),
		  Name: n.Name,
		  Query: convertNode(n.Query),
	}
}

func convertQuery(n *nodes.Query) *pg.Query {
	if n == nil {
		return nil
	}
	return &pg.Query{
		  CanSetTag: n.CanSetTag,
		  CommandType: pg.CmdType(n.CommandType),
		  ConstraintDeps: convertList(n.ConstraintDeps),
		  CteList: convertList(n.CteList),
		  DistinctClause: convertList(n.DistinctClause),
		  GroupClause: convertList(n.GroupClause),
		  GroupingSets: convertList(n.GroupingSets),
		  HasAggs: n.HasAggs,
		  HasDistinctOn: n.HasDistinctOn,
		  HasForUpdate: n.HasForUpdate,
		  HasModifyingCte: n.HasModifyingCte,
		  HasRecursive: n.HasRecursive,
		  HasRowSecurity: n.HasRowSecurity,
		  HasSubLinks: n.HasSubLinks,
		  HasTargetSrfs: n.HasTargetSrfs,
		  HasWindowFuncs: n.HasWindowFuncs,
		  HavingQual: convertNode(n.HavingQual),
		  Jointree: convertFromExpr(n.Jointree),
		  LimitCount: convertNode(n.LimitCount),
		  LimitOffset: convertNode(n.LimitOffset),
		  OnConflict: convertOnConflictExpr(n.OnConflict),
		  Override: pg.OverridingKind(n.Override),
		  QueryId: n.QueryId,
		  QuerySource: pg.QuerySource(n.QuerySource),
		  ResultRelation: n.ResultRelation,
		  ReturningList: convertList(n.ReturningList),
		  RowMarks: convertList(n.RowMarks),
		  Rtable: convertList(n.Rtable),
		  SetOperations: convertNode(n.SetOperations),
		  SortClause: convertList(n.SortClause),
		  StmtLen: n.StmtLen,
		  StmtLocation: n.StmtLocation,
		  TargetList: convertList(n.TargetList),
		  UtilityStmt: convertNode(n.UtilityStmt),
		  WindowClause: convertList(n.WindowClause),
		  WithCheckOptions: convertList(n.WithCheckOptions),
	}
}

func convertRangeFunction(n *nodes.RangeFunction) *pg.RangeFunction {
	if n == nil {
		return nil
	}
	return &pg.RangeFunction{
		  Alias: convertAlias(n.Alias),
		  Coldeflist: convertList(n.Coldeflist),
		  Functions: convertList(n.Functions),
		  IsRowsfrom: n.IsRowsfrom,
		  Lateral: n.Lateral,
		  Ordinality: n.Ordinality,
	}
}

func convertRangeSubselect(n *nodes.RangeSubselect) *pg.RangeSubselect {
	if n == nil {
		return nil
	}
	return &pg.RangeSubselect{
		  Alias: convertAlias(n.Alias),
		  Lateral: n.Lateral,
		  Subquery: convertNode(n.Subquery),
	}
}

func convertRangeTableFunc(n *nodes.RangeTableFunc) *pg.RangeTableFunc {
	if n == nil {
		return nil
	}
	return &pg.RangeTableFunc{
		  Alias: convertAlias(n.Alias),
		  Columns: convertList(n.Columns),
		  Docexpr: convertNode(n.Docexpr),
		  Lateral: n.Lateral,
		  Location: n.Location,
		  Namespaces: convertList(n.Namespaces),
		  Rowexpr: convertNode(n.Rowexpr),
	}
}

func convertRangeTableFuncCol(n *nodes.RangeTableFuncCol) *pg.RangeTableFuncCol {
	if n == nil {
		return nil
	}
	return &pg.RangeTableFuncCol{
		  Coldefexpr: convertNode(n.Coldefexpr),
		  Colexpr: convertNode(n.Colexpr),
		  Colname: n.Colname,
		  ForOrdinality: n.ForOrdinality,
		  IsNotNull: n.IsNotNull,
		  Location: n.Location,
		  TypeName: convertTypeName(n.TypeName),
	}
}

func convertRangeTableSample(n *nodes.RangeTableSample) *pg.RangeTableSample {
	if n == nil {
		return nil
	}
	return &pg.RangeTableSample{
		  Args: convertList(n.Args),
		  Location: n.Location,
		  Method: convertList(n.Method),
		  Relation: convertNode(n.Relation),
		  Repeatable: convertNode(n.Repeatable),
	}
}

func convertRangeTblEntry(n *nodes.RangeTblEntry) *pg.RangeTblEntry {
	if n == nil {
		return nil
	}
	return &pg.RangeTblEntry{
		  Alias: convertAlias(n.Alias),
		  CheckAsUser: pg.Oid(n.CheckAsUser),
		  Colcollations: convertList(n.Colcollations),
		  Coltypes: convertList(n.Coltypes),
		  Coltypmods: convertList(n.Coltypmods),
		  Ctelevelsup: pg.Index(n.Ctelevelsup),
		  Ctename: n.Ctename,
		  Enrname: n.Enrname,
		  Enrtuples: n.Enrtuples,
		  Eref: convertAlias(n.Eref),
		  Funcordinality: n.Funcordinality,
		  Functions: convertList(n.Functions),
		  InFromCl: n.InFromCl,
		  Inh: n.Inh,
		  InsertedCols: n.InsertedCols,
		  Joinaliasvars: convertList(n.Joinaliasvars),
		  Jointype: pg.JoinType(n.Jointype),
		  Lateral: n.Lateral,
		  Relid: pg.Oid(n.Relid),
		  Relkind: n.Relkind,
		  RequiredPerms: pg.AclMode(n.RequiredPerms),
		  Rtekind: pg.RTEKind(n.Rtekind),
		  SecurityBarrier: n.SecurityBarrier,
		  SecurityQuals: convertList(n.SecurityQuals),
		  SelectedCols: n.SelectedCols,
		  SelfReference: n.SelfReference,
		  Subquery: convertQuery(n.Subquery),
		  Tablefunc: convertTableFunc(n.Tablefunc),
		  Tablesample: convertTableSampleClause(n.Tablesample),
		  UpdatedCols: n.UpdatedCols,
		  ValuesLists: convertList(n.ValuesLists),
	}
}

func convertRangeTblFunction(n *nodes.RangeTblFunction) *pg.RangeTblFunction {
	if n == nil {
		return nil
	}
	return &pg.RangeTblFunction{
		  Funccolcollations: convertList(n.Funccolcollations),
		  Funccolcount: n.Funccolcount,
		  Funccolnames: convertList(n.Funccolnames),
		  Funccoltypes: convertList(n.Funccoltypes),
		  Funccoltypmods: convertList(n.Funccoltypmods),
		  Funcexpr: convertNode(n.Funcexpr),
		  Funcparams: n.Funcparams,
	}
}

func convertRangeTblRef(n *nodes.RangeTblRef) *pg.RangeTblRef {
	if n == nil {
		return nil
	}
	return &pg.RangeTblRef{
		  Rtindex: n.Rtindex,
	}
}

func convertRangeVar(n *nodes.RangeVar) *pg.RangeVar {
	if n == nil {
		return nil
	}
	return &pg.RangeVar{
		  Alias: convertAlias(n.Alias),
		  Catalogname: n.Catalogname,
		  Inh: n.Inh,
		  Location: n.Location,
		  Relname: n.Relname,
		  Relpersistence: n.Relpersistence,
		  Schemaname: n.Schemaname,
	}
}

func convertRawStmt(n *nodes.RawStmt) *pg.RawStmt {
	if n == nil {
		return nil
	}
	return &pg.RawStmt{
		  Stmt: convertNode(n.Stmt),
		  StmtLen: n.StmtLen,
		  StmtLocation: n.StmtLocation,
	}
}

func convertReassignOwnedStmt(n *nodes.ReassignOwnedStmt) *pg.ReassignOwnedStmt {
	if n == nil {
		return nil
	}
	return &pg.ReassignOwnedStmt{
		  Newrole: convertRoleSpec(n.Newrole),
		  Roles: convertList(n.Roles),
	}
}

func convertRefreshMatViewStmt(n *nodes.RefreshMatViewStmt) *pg.RefreshMatViewStmt {
	if n == nil {
		return nil
	}
	return &pg.RefreshMatViewStmt{
		  Concurrent: n.Concurrent,
		  Relation: convertRangeVar(n.Relation),
		  SkipData: n.SkipData,
	}
}

func convertReindexStmt(n *nodes.ReindexStmt) *pg.ReindexStmt {
	if n == nil {
		return nil
	}
	return &pg.ReindexStmt{
		  Kind: pg.ReindexObjectType(n.Kind),
		  Name: n.Name,
		  Options: n.Options,
		  Relation: convertRangeVar(n.Relation),
	}
}

func convertRelabelType(n *nodes.RelabelType) *pg.RelabelType {
	if n == nil {
		return nil
	}
	return &pg.RelabelType{
		  Arg: convertNode(n.Arg),
		  Location: n.Location,
		  Relabelformat: pg.CoercionForm(n.Relabelformat),
		  Resultcollid: pg.Oid(n.Resultcollid),
		  Resulttype: pg.Oid(n.Resulttype),
		  Resulttypmod: n.Resulttypmod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertRenameStmt(n *nodes.RenameStmt) *pg.RenameStmt {
	if n == nil {
		return nil
	}
	return &pg.RenameStmt{
		  Behavior: pg.DropBehavior(n.Behavior),
		  MissingOk: n.MissingOk,
		  Newname: n.Newname,
		  Object: convertNode(n.Object),
		  Relation: convertRangeVar(n.Relation),
		  RelationType: pg.ObjectType(n.RelationType),
		  RenameType: pg.ObjectType(n.RenameType),
		  Subname: n.Subname,
	}
}

func convertReplicaIdentityStmt(n *nodes.ReplicaIdentityStmt) *pg.ReplicaIdentityStmt {
	if n == nil {
		return nil
	}
	return &pg.ReplicaIdentityStmt{
		  IdentityType: n.IdentityType,
		  Name: n.Name,
	}
}

func convertResTarget(n *nodes.ResTarget) *pg.ResTarget {
	if n == nil {
		return nil
	}
	return &pg.ResTarget{
		  Indirection: convertList(n.Indirection),
		  Location: n.Location,
		  Name: n.Name,
		  Val: convertNode(n.Val),
	}
}

func convertRoleSpec(n *nodes.RoleSpec) *pg.RoleSpec {
	if n == nil {
		return nil
	}
	return &pg.RoleSpec{
		  Location: n.Location,
		  Rolename: n.Rolename,
		  Roletype: pg.RoleSpecType(n.Roletype),
	}
}

func convertRowCompareExpr(n *nodes.RowCompareExpr) *pg.RowCompareExpr {
	if n == nil {
		return nil
	}
	return &pg.RowCompareExpr{
		  Inputcollids: convertList(n.Inputcollids),
		  Largs: convertList(n.Largs),
		  Opfamilies: convertList(n.Opfamilies),
		  Opnos: convertList(n.Opnos),
		  Rargs: convertList(n.Rargs),
		  Rctype: pg.RowCompareType(n.Rctype),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertRowExpr(n *nodes.RowExpr) *pg.RowExpr {
	if n == nil {
		return nil
	}
	return &pg.RowExpr{
		  Args: convertList(n.Args),
		  Colnames: convertList(n.Colnames),
		  Location: n.Location,
		  RowFormat: pg.CoercionForm(n.RowFormat),
		  RowTypeid: pg.Oid(n.RowTypeid),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertRowMarkClause(n *nodes.RowMarkClause) *pg.RowMarkClause {
	if n == nil {
		return nil
	}
	return &pg.RowMarkClause{
		  PushedDown: n.PushedDown,
		  Rti: pg.Index(n.Rti),
		  Strength: pg.LockClauseStrength(n.Strength),
		  WaitPolicy: pg.LockWaitPolicy(n.WaitPolicy),
	}
}

func convertRuleStmt(n *nodes.RuleStmt) *pg.RuleStmt {
	if n == nil {
		return nil
	}
	return &pg.RuleStmt{
		  Actions: convertList(n.Actions),
		  Event: pg.CmdType(n.Event),
		  Instead: n.Instead,
		  Relation: convertRangeVar(n.Relation),
		  Replace: n.Replace,
		  Rulename: n.Rulename,
		  WhereClause: convertNode(n.WhereClause),
	}
}

func convertSQLValueFunction(n *nodes.SQLValueFunction) *pg.SQLValueFunction {
	if n == nil {
		return nil
	}
	return &pg.SQLValueFunction{
		  Location: n.Location,
		  Op: pg.SQLValueFunctionOp(n.Op),
		  Type: pg.Oid(n.Type),
		  Typmod: n.Typmod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertScalarArrayOpExpr(n *nodes.ScalarArrayOpExpr) *pg.ScalarArrayOpExpr {
	if n == nil {
		return nil
	}
	return &pg.ScalarArrayOpExpr{
		  Args: convertList(n.Args),
		  Inputcollid: pg.Oid(n.Inputcollid),
		  Location: n.Location,
		  Opfuncid: pg.Oid(n.Opfuncid),
		  Opno: pg.Oid(n.Opno),
		  UseOr: n.UseOr,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertSecLabelStmt(n *nodes.SecLabelStmt) *pg.SecLabelStmt {
	if n == nil {
		return nil
	}
	return &pg.SecLabelStmt{
		  Label: n.Label,
		  Object: convertNode(n.Object),
		  Objtype: pg.ObjectType(n.Objtype),
		  Provider: n.Provider,
	}
}

func convertSelectStmt(n *nodes.SelectStmt) *pg.SelectStmt {
	if n == nil {
		return nil
	}
	return &pg.SelectStmt{
		  All: n.All,
		  DistinctClause: convertList(n.DistinctClause),
		  FromClause: convertList(n.FromClause),
		  GroupClause: convertList(n.GroupClause),
		  HavingClause: convertNode(n.HavingClause),
		  IntoClause: convertIntoClause(n.IntoClause),
		  Larg: convertSelectStmt(n.Larg),
		  LimitCount: convertNode(n.LimitCount),
		  LimitOffset: convertNode(n.LimitOffset),
		  LockingClause: convertList(n.LockingClause),
		  Op: pg.SetOperation(n.Op),
		  Rarg: convertSelectStmt(n.Rarg),
		  SortClause: convertList(n.SortClause),
		  TargetList: convertList(n.TargetList),
		  ValuesLists: convertValuesList(n.ValuesLists),
		  WhereClause: convertNode(n.WhereClause),
		  WindowClause: convertList(n.WindowClause),
		  WithClause: convertWithClause(n.WithClause),
	}
}

func convertSetOperationStmt(n *nodes.SetOperationStmt) *pg.SetOperationStmt {
	if n == nil {
		return nil
	}
	return &pg.SetOperationStmt{
		  All: n.All,
		  ColCollations: convertList(n.ColCollations),
		  ColTypes: convertList(n.ColTypes),
		  ColTypmods: convertList(n.ColTypmods),
		  GroupClauses: convertList(n.GroupClauses),
		  Larg: convertNode(n.Larg),
		  Op: pg.SetOperation(n.Op),
		  Rarg: convertNode(n.Rarg),
	}
}

func convertSetToDefault(n *nodes.SetToDefault) *pg.SetToDefault {
	if n == nil {
		return nil
	}
	return &pg.SetToDefault{
		  Collation: pg.Oid(n.Collation),
		  Location: n.Location,
		  TypeId: pg.Oid(n.TypeId),
		  TypeMod: n.TypeMod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertSortBy(n *nodes.SortBy) *pg.SortBy {
	if n == nil {
		return nil
	}
	return &pg.SortBy{
		  Location: n.Location,
		  Node: convertNode(n.Node),
		  SortbyDir: pg.SortByDir(n.SortbyDir),
		  SortbyNulls: pg.SortByNulls(n.SortbyNulls),
		  UseOp: convertList(n.UseOp),
	}
}

func convertSortGroupClause(n *nodes.SortGroupClause) *pg.SortGroupClause {
	if n == nil {
		return nil
	}
	return &pg.SortGroupClause{
		  Eqop: pg.Oid(n.Eqop),
		  Hashable: n.Hashable,
		  NullsFirst: n.NullsFirst,
		  Sortop: pg.Oid(n.Sortop),
		  TleSortGroupRef: pg.Index(n.TleSortGroupRef),
	}
}

func convertString(n *nodes.String) *pg.String {
	if n == nil {
		return nil
	}
	return &pg.String{
		  Str: n.Str,
	}
}

func convertSubLink(n *nodes.SubLink) *pg.SubLink {
	if n == nil {
		return nil
	}
	return &pg.SubLink{
		  Location: n.Location,
		  OperName: convertList(n.OperName),
		  SubLinkId: n.SubLinkId,
		  SubLinkType: pg.SubLinkType(n.SubLinkType),
		  Subselect: convertNode(n.Subselect),
		  Testexpr: convertNode(n.Testexpr),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertSubPlan(n *nodes.SubPlan) *pg.SubPlan {
	if n == nil {
		return nil
	}
	return &pg.SubPlan{
		  Args: convertList(n.Args),
		  FirstColCollation: pg.Oid(n.FirstColCollation),
		  FirstColType: pg.Oid(n.FirstColType),
		  FirstColTypmod: n.FirstColTypmod,
		  ParParam: convertList(n.ParParam),
		  ParallelSafe: n.ParallelSafe,
		  ParamIds: convertList(n.ParamIds),
		  PerCallCost: pg.Cost(n.PerCallCost),
		  PlanId: n.PlanId,
		  PlanName: n.PlanName,
		  SetParam: convertList(n.SetParam),
		  StartupCost: pg.Cost(n.StartupCost),
		  SubLinkType: pg.SubLinkType(n.SubLinkType),
		  Testexpr: convertNode(n.Testexpr),
		  UnknownEqFalse: n.UnknownEqFalse,
		  UseHashTable: n.UseHashTable,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertTableFunc(n *nodes.TableFunc) *pg.TableFunc {
	if n == nil {
		return nil
	}
	return &pg.TableFunc{
		  Colcollations: convertList(n.Colcollations),
		  Coldefexprs: convertList(n.Coldefexprs),
		  Colexprs: convertList(n.Colexprs),
		  Colnames: convertList(n.Colnames),
		  Coltypes: convertList(n.Coltypes),
		  Coltypmods: convertList(n.Coltypmods),
		  Docexpr: convertNode(n.Docexpr),
		  Location: n.Location,
		  Notnulls: n.Notnulls,
		  NsNames: convertList(n.NsNames),
		  NsUris: convertList(n.NsUris),
		  Ordinalitycol: n.Ordinalitycol,
		  Rowexpr: convertNode(n.Rowexpr),
	}
}

func convertTableLikeClause(n *nodes.TableLikeClause) *pg.TableLikeClause {
	if n == nil {
		return nil
	}
	return &pg.TableLikeClause{
		  Options: n.Options,
		  Relation: convertRangeVar(n.Relation),
	}
}

func convertTableSampleClause(n *nodes.TableSampleClause) *pg.TableSampleClause {
	if n == nil {
		return nil
	}
	return &pg.TableSampleClause{
		  Args: convertList(n.Args),
		  Repeatable: convertNode(n.Repeatable),
		  Tsmhandler: pg.Oid(n.Tsmhandler),
	}
}

func convertTargetEntry(n *nodes.TargetEntry) *pg.TargetEntry {
	if n == nil {
		return nil
	}
	return &pg.TargetEntry{
		  Expr: convertNode(n.Expr),
		  Resjunk: n.Resjunk,
		  Resname: n.Resname,
		  Resno: pg.AttrNumber(n.Resno),
		  Resorigcol: pg.AttrNumber(n.Resorigcol),
		  Resorigtbl: pg.Oid(n.Resorigtbl),
		  Ressortgroupref: pg.Index(n.Ressortgroupref),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertTransactionStmt(n *nodes.TransactionStmt) *pg.TransactionStmt {
	if n == nil {
		return nil
	}
	return &pg.TransactionStmt{
		  Gid: n.Gid,
		  Kind: pg.TransactionStmtKind(n.Kind),
		  Options: convertList(n.Options),
	}
}

func convertTriggerTransition(n *nodes.TriggerTransition) *pg.TriggerTransition {
	if n == nil {
		return nil
	}
	return &pg.TriggerTransition{
		  IsNew: n.IsNew,
		  IsTable: n.IsTable,
		  Name: n.Name,
	}
}

func convertTruncateStmt(n *nodes.TruncateStmt) *pg.TruncateStmt {
	if n == nil {
		return nil
	}
	return &pg.TruncateStmt{
		  Behavior: pg.DropBehavior(n.Behavior),
		  Relations: convertList(n.Relations),
		  RestartSeqs: n.RestartSeqs,
	}
}

func convertTypeCast(n *nodes.TypeCast) *pg.TypeCast {
	if n == nil {
		return nil
	}
	return &pg.TypeCast{
		  Arg: convertNode(n.Arg),
		  Location: n.Location,
		  TypeName: convertTypeName(n.TypeName),
	}
}

func convertTypeName(n *nodes.TypeName) *pg.TypeName {
	if n == nil {
		return nil
	}
	return &pg.TypeName{
		  ArrayBounds: convertList(n.ArrayBounds),
		  Location: n.Location,
		  Names: convertList(n.Names),
		  PctType: n.PctType,
		  Setof: n.Setof,
		  TypeOid: pg.Oid(n.TypeOid),
		  Typemod: n.Typemod,
		  Typmods: convertList(n.Typmods),
	}
}

func convertUnlistenStmt(n *nodes.UnlistenStmt) *pg.UnlistenStmt {
	if n == nil {
		return nil
	}
	return &pg.UnlistenStmt{
		  Conditionname: n.Conditionname,
	}
}

func convertUpdateStmt(n *nodes.UpdateStmt) *pg.UpdateStmt {
	if n == nil {
		return nil
	}
	return &pg.UpdateStmt{
		  FromClause: convertList(n.FromClause),
		  Relation: convertRangeVar(n.Relation),
		  ReturningList: convertList(n.ReturningList),
		  TargetList: convertList(n.TargetList),
		  WhereClause: convertNode(n.WhereClause),
		  WithClause: convertWithClause(n.WithClause),
	}
}

func convertVacuumStmt(n *nodes.VacuumStmt) *pg.VacuumStmt {
	if n == nil {
		return nil
	}
	return &pg.VacuumStmt{
		  Options: n.Options,
		  Relation: convertRangeVar(n.Relation),
		  VaCols: convertList(n.VaCols),
	}
}

func convertVar(n *nodes.Var) *pg.Var {
	if n == nil {
		return nil
	}
	return &pg.Var{
		  Location: n.Location,
		  Varattno: pg.AttrNumber(n.Varattno),
		  Varcollid: pg.Oid(n.Varcollid),
		  Varlevelsup: pg.Index(n.Varlevelsup),
		  Varno: pg.Index(n.Varno),
		  Varnoold: pg.Index(n.Varnoold),
		  Varoattno: pg.AttrNumber(n.Varoattno),
		  Vartype: pg.Oid(n.Vartype),
		  Vartypmod: n.Vartypmod,
		  Xpr: convertNode(n.Xpr),
	}
}

func convertVariableSetStmt(n *nodes.VariableSetStmt) *pg.VariableSetStmt {
	if n == nil {
		return nil
	}
	return &pg.VariableSetStmt{
		  Args: convertList(n.Args),
		  IsLocal: n.IsLocal,
		  Kind: pg.VariableSetKind(n.Kind),
		  Name: n.Name,
	}
}

func convertVariableShowStmt(n *nodes.VariableShowStmt) *pg.VariableShowStmt {
	if n == nil {
		return nil
	}
	return &pg.VariableShowStmt{
		  Name: n.Name,
	}
}

func convertViewStmt(n *nodes.ViewStmt) *pg.ViewStmt {
	if n == nil {
		return nil
	}
	return &pg.ViewStmt{
		  Aliases: convertList(n.Aliases),
		  Options: convertList(n.Options),
		  Query: convertNode(n.Query),
		  Replace: n.Replace,
		  View: convertRangeVar(n.View),
		  WithCheckOption: pg.ViewCheckOption(n.WithCheckOption),
	}
}

func convertWindowClause(n *nodes.WindowClause) *pg.WindowClause {
	if n == nil {
		return nil
	}
	return &pg.WindowClause{
		  CopiedOrder: n.CopiedOrder,
		  EndOffset: convertNode(n.EndOffset),
		  FrameOptions: n.FrameOptions,
		  Name: n.Name,
		  OrderClause: convertList(n.OrderClause),
		  PartitionClause: convertList(n.PartitionClause),
		  Refname: n.Refname,
		  StartOffset: convertNode(n.StartOffset),
		  Winref: pg.Index(n.Winref),
	}
}

func convertWindowDef(n *nodes.WindowDef) *pg.WindowDef {
	if n == nil {
		return nil
	}
	return &pg.WindowDef{
		  EndOffset: convertNode(n.EndOffset),
		  FrameOptions: n.FrameOptions,
		  Location: n.Location,
		  Name: n.Name,
		  OrderClause: convertList(n.OrderClause),
		  PartitionClause: convertList(n.PartitionClause),
		  Refname: n.Refname,
		  StartOffset: convertNode(n.StartOffset),
	}
}

func convertWindowFunc(n *nodes.WindowFunc) *pg.WindowFunc {
	if n == nil {
		return nil
	}
	return &pg.WindowFunc{
		  Aggfilter: convertNode(n.Aggfilter),
		  Args: convertList(n.Args),
		  Inputcollid: pg.Oid(n.Inputcollid),
		  Location: n.Location,
		  Winagg: n.Winagg,
		  Wincollid: pg.Oid(n.Wincollid),
		  Winfnoid: pg.Oid(n.Winfnoid),
		  Winref: pg.Index(n.Winref),
		  Winstar: n.Winstar,
		  Wintype: pg.Oid(n.Wintype),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertWithCheckOption(n *nodes.WithCheckOption) *pg.WithCheckOption {
	if n == nil {
		return nil
	}
	return &pg.WithCheckOption{
		  Cascaded: n.Cascaded,
		  Kind: pg.WCOKind(n.Kind),
		  Polname: n.Polname,
		  Qual: convertNode(n.Qual),
		  Relname: n.Relname,
	}
}

func convertWithClause(n *nodes.WithClause) *pg.WithClause {
	if n == nil {
		return nil
	}
	return &pg.WithClause{
		  Ctes: convertList(n.Ctes),
		  Location: n.Location,
		  Recursive: n.Recursive,
	}
}

func convertXmlExpr(n *nodes.XmlExpr) *pg.XmlExpr {
	if n == nil {
		return nil
	}
	return &pg.XmlExpr{
		  ArgNames: convertList(n.ArgNames),
		  Args: convertList(n.Args),
		  Location: n.Location,
		  Name: n.Name,
		  NamedArgs: convertList(n.NamedArgs),
		  Op: pg.XmlExprOp(n.Op),
		  Type: pg.Oid(n.Type),
		  Typmod: n.Typmod,
		  Xmloption: pg.XmlOptionType(n.Xmloption),
		  Xpr: convertNode(n.Xpr),
	}
}

func convertXmlSerialize(n *nodes.XmlSerialize) *pg.XmlSerialize {
	if n == nil {
		return nil
	}
	return &pg.XmlSerialize{
		  Expr: convertNode(n.Expr),
		  Location: n.Location,
		  TypeName: convertTypeName(n.TypeName),
		  Xmloption: pg.XmlOptionType(n.Xmloption),
	}
}



func convertNode(node nodes.Node) ast.Node {
	switch n := node.(type) {
	
	case nodes.A_ArrayExpr:
		return convertA_ArrayExpr(&n)
	
	case nodes.A_Const:
		return convertA_Const(&n)
	
	case nodes.A_Expr:
		return convertA_Expr(&n)
	
	case nodes.A_Indices:
		return convertA_Indices(&n)
	
	case nodes.A_Indirection:
		return convertA_Indirection(&n)
	
	case nodes.A_Star:
		return convertA_Star(&n)
	
	case nodes.AccessPriv:
		return convertAccessPriv(&n)
	
	case nodes.Aggref:
		return convertAggref(&n)
	
	case nodes.Alias:
		return convertAlias(&n)
	
	case nodes.AlterCollationStmt:
		return convertAlterCollationStmt(&n)
	
	case nodes.AlterDatabaseSetStmt:
		return convertAlterDatabaseSetStmt(&n)
	
	case nodes.AlterDatabaseStmt:
		return convertAlterDatabaseStmt(&n)
	
	case nodes.AlterDefaultPrivilegesStmt:
		return convertAlterDefaultPrivilegesStmt(&n)
	
	case nodes.AlterDomainStmt:
		return convertAlterDomainStmt(&n)
	
	case nodes.AlterEnumStmt:
		return convertAlterEnumStmt(&n)
	
	case nodes.AlterEventTrigStmt:
		return convertAlterEventTrigStmt(&n)
	
	case nodes.AlterExtensionContentsStmt:
		return convertAlterExtensionContentsStmt(&n)
	
	case nodes.AlterExtensionStmt:
		return convertAlterExtensionStmt(&n)
	
	case nodes.AlterFdwStmt:
		return convertAlterFdwStmt(&n)
	
	case nodes.AlterForeignServerStmt:
		return convertAlterForeignServerStmt(&n)
	
	case nodes.AlterFunctionStmt:
		return convertAlterFunctionStmt(&n)
	
	case nodes.AlterObjectDependsStmt:
		return convertAlterObjectDependsStmt(&n)
	
	case nodes.AlterObjectSchemaStmt:
		return convertAlterObjectSchemaStmt(&n)
	
	case nodes.AlterOpFamilyStmt:
		return convertAlterOpFamilyStmt(&n)
	
	case nodes.AlterOperatorStmt:
		return convertAlterOperatorStmt(&n)
	
	case nodes.AlterOwnerStmt:
		return convertAlterOwnerStmt(&n)
	
	case nodes.AlterPolicyStmt:
		return convertAlterPolicyStmt(&n)
	
	case nodes.AlterPublicationStmt:
		return convertAlterPublicationStmt(&n)
	
	case nodes.AlterRoleSetStmt:
		return convertAlterRoleSetStmt(&n)
	
	case nodes.AlterRoleStmt:
		return convertAlterRoleStmt(&n)
	
	case nodes.AlterSeqStmt:
		return convertAlterSeqStmt(&n)
	
	case nodes.AlterSubscriptionStmt:
		return convertAlterSubscriptionStmt(&n)
	
	case nodes.AlterSystemStmt:
		return convertAlterSystemStmt(&n)
	
	case nodes.AlterTSConfigurationStmt:
		return convertAlterTSConfigurationStmt(&n)
	
	case nodes.AlterTSDictionaryStmt:
		return convertAlterTSDictionaryStmt(&n)
	
	case nodes.AlterTableCmd:
		return convertAlterTableCmd(&n)
	
	case nodes.AlterTableMoveAllStmt:
		return convertAlterTableMoveAllStmt(&n)
	
	case nodes.AlterTableSpaceOptionsStmt:
		return convertAlterTableSpaceOptionsStmt(&n)
	
	case nodes.AlterTableStmt:
		return convertAlterTableStmt(&n)
	
	case nodes.AlterUserMappingStmt:
		return convertAlterUserMappingStmt(&n)
	
	case nodes.AlternativeSubPlan:
		return convertAlternativeSubPlan(&n)
	
	case nodes.ArrayCoerceExpr:
		return convertArrayCoerceExpr(&n)
	
	case nodes.ArrayExpr:
		return convertArrayExpr(&n)
	
	case nodes.ArrayRef:
		return convertArrayRef(&n)
	
	case nodes.BitString:
		return convertBitString(&n)
	
	case nodes.BlockIdData:
		return convertBlockIdData(&n)
	
	case nodes.BoolExpr:
		return convertBoolExpr(&n)
	
	case nodes.BooleanTest:
		return convertBooleanTest(&n)
	
	case nodes.CaseExpr:
		return convertCaseExpr(&n)
	
	case nodes.CaseTestExpr:
		return convertCaseTestExpr(&n)
	
	case nodes.CaseWhen:
		return convertCaseWhen(&n)
	
	case nodes.CheckPointStmt:
		return convertCheckPointStmt(&n)
	
	case nodes.ClosePortalStmt:
		return convertClosePortalStmt(&n)
	
	case nodes.ClusterStmt:
		return convertClusterStmt(&n)
	
	case nodes.CoalesceExpr:
		return convertCoalesceExpr(&n)
	
	case nodes.CoerceToDomain:
		return convertCoerceToDomain(&n)
	
	case nodes.CoerceToDomainValue:
		return convertCoerceToDomainValue(&n)
	
	case nodes.CoerceViaIO:
		return convertCoerceViaIO(&n)
	
	case nodes.CollateClause:
		return convertCollateClause(&n)
	
	case nodes.CollateExpr:
		return convertCollateExpr(&n)
	
	case nodes.ColumnDef:
		return convertColumnDef(&n)
	
	case nodes.ColumnRef:
		return convertColumnRef(&n)
	
	case nodes.CommentStmt:
		return convertCommentStmt(&n)
	
	case nodes.CommonTableExpr:
		return convertCommonTableExpr(&n)
	
	case nodes.CompositeTypeStmt:
		return convertCompositeTypeStmt(&n)
	
	case nodes.Const:
		return convertConst(&n)
	
	case nodes.Constraint:
		return convertConstraint(&n)
	
	case nodes.ConstraintsSetStmt:
		return convertConstraintsSetStmt(&n)
	
	case nodes.ConvertRowtypeExpr:
		return convertConvertRowtypeExpr(&n)
	
	case nodes.CopyStmt:
		return convertCopyStmt(&n)
	
	case nodes.CreateAmStmt:
		return convertCreateAmStmt(&n)
	
	case nodes.CreateCastStmt:
		return convertCreateCastStmt(&n)
	
	case nodes.CreateConversionStmt:
		return convertCreateConversionStmt(&n)
	
	case nodes.CreateDomainStmt:
		return convertCreateDomainStmt(&n)
	
	case nodes.CreateEnumStmt:
		return convertCreateEnumStmt(&n)
	
	case nodes.CreateEventTrigStmt:
		return convertCreateEventTrigStmt(&n)
	
	case nodes.CreateExtensionStmt:
		return convertCreateExtensionStmt(&n)
	
	case nodes.CreateFdwStmt:
		return convertCreateFdwStmt(&n)
	
	case nodes.CreateForeignServerStmt:
		return convertCreateForeignServerStmt(&n)
	
	case nodes.CreateForeignTableStmt:
		return convertCreateForeignTableStmt(&n)
	
	case nodes.CreateFunctionStmt:
		return convertCreateFunctionStmt(&n)
	
	case nodes.CreateOpClassItem:
		return convertCreateOpClassItem(&n)
	
	case nodes.CreateOpClassStmt:
		return convertCreateOpClassStmt(&n)
	
	case nodes.CreateOpFamilyStmt:
		return convertCreateOpFamilyStmt(&n)
	
	case nodes.CreatePLangStmt:
		return convertCreatePLangStmt(&n)
	
	case nodes.CreatePolicyStmt:
		return convertCreatePolicyStmt(&n)
	
	case nodes.CreatePublicationStmt:
		return convertCreatePublicationStmt(&n)
	
	case nodes.CreateRangeStmt:
		return convertCreateRangeStmt(&n)
	
	case nodes.CreateRoleStmt:
		return convertCreateRoleStmt(&n)
	
	case nodes.CreateSchemaStmt:
		return convertCreateSchemaStmt(&n)
	
	case nodes.CreateSeqStmt:
		return convertCreateSeqStmt(&n)
	
	case nodes.CreateStatsStmt:
		return convertCreateStatsStmt(&n)
	
	case nodes.CreateStmt:
		return convertCreateStmt(&n)
	
	case nodes.CreateSubscriptionStmt:
		return convertCreateSubscriptionStmt(&n)
	
	case nodes.CreateTableAsStmt:
		return convertCreateTableAsStmt(&n)
	
	case nodes.CreateTableSpaceStmt:
		return convertCreateTableSpaceStmt(&n)
	
	case nodes.CreateTransformStmt:
		return convertCreateTransformStmt(&n)
	
	case nodes.CreateTrigStmt:
		return convertCreateTrigStmt(&n)
	
	case nodes.CreateUserMappingStmt:
		return convertCreateUserMappingStmt(&n)
	
	case nodes.CreatedbStmt:
		return convertCreatedbStmt(&n)
	
	case nodes.CurrentOfExpr:
		return convertCurrentOfExpr(&n)
	
	case nodes.DeallocateStmt:
		return convertDeallocateStmt(&n)
	
	case nodes.DeclareCursorStmt:
		return convertDeclareCursorStmt(&n)
	
	case nodes.DefElem:
		return convertDefElem(&n)
	
	case nodes.DefineStmt:
		return convertDefineStmt(&n)
	
	case nodes.DeleteStmt:
		return convertDeleteStmt(&n)
	
	case nodes.DiscardStmt:
		return convertDiscardStmt(&n)
	
	case nodes.DoStmt:
		return convertDoStmt(&n)
	
	case nodes.DropOwnedStmt:
		return convertDropOwnedStmt(&n)
	
	case nodes.DropRoleStmt:
		return convertDropRoleStmt(&n)
	
	case nodes.DropStmt:
		return convertDropStmt(&n)
	
	case nodes.DropSubscriptionStmt:
		return convertDropSubscriptionStmt(&n)
	
	case nodes.DropTableSpaceStmt:
		return convertDropTableSpaceStmt(&n)
	
	case nodes.DropUserMappingStmt:
		return convertDropUserMappingStmt(&n)
	
	case nodes.DropdbStmt:
		return convertDropdbStmt(&n)
	
	case nodes.ExecuteStmt:
		return convertExecuteStmt(&n)
	
	case nodes.ExplainStmt:
		return convertExplainStmt(&n)
	
	case nodes.Expr:
		return convertExpr(&n)
	
	case nodes.FetchStmt:
		return convertFetchStmt(&n)
	
	case nodes.FieldSelect:
		return convertFieldSelect(&n)
	
	case nodes.FieldStore:
		return convertFieldStore(&n)
	
	case nodes.Float:
		return convertFloat(&n)
	
	case nodes.FromExpr:
		return convertFromExpr(&n)
	
	case nodes.FuncCall:
		return convertFuncCall(&n)
	
	case nodes.FuncExpr:
		return convertFuncExpr(&n)
	
	case nodes.FunctionParameter:
		return convertFunctionParameter(&n)
	
	case nodes.GrantRoleStmt:
		return convertGrantRoleStmt(&n)
	
	case nodes.GrantStmt:
		return convertGrantStmt(&n)
	
	case nodes.GroupingFunc:
		return convertGroupingFunc(&n)
	
	case nodes.GroupingSet:
		return convertGroupingSet(&n)
	
	case nodes.ImportForeignSchemaStmt:
		return convertImportForeignSchemaStmt(&n)
	
	case nodes.IndexElem:
		return convertIndexElem(&n)
	
	case nodes.IndexStmt:
		return convertIndexStmt(&n)
	
	case nodes.InferClause:
		return convertInferClause(&n)
	
	case nodes.InferenceElem:
		return convertInferenceElem(&n)
	
	case nodes.InlineCodeBlock:
		return convertInlineCodeBlock(&n)
	
	case nodes.InsertStmt:
		return convertInsertStmt(&n)
	
	case nodes.Integer:
		return convertInteger(&n)
	
	case nodes.IntoClause:
		return convertIntoClause(&n)
	
	case nodes.JoinExpr:
		return convertJoinExpr(&n)
	
	case nodes.ListenStmt:
		return convertListenStmt(&n)
	
	case nodes.LoadStmt:
		return convertLoadStmt(&n)
	
	case nodes.LockStmt:
		return convertLockStmt(&n)
	
	case nodes.LockingClause:
		return convertLockingClause(&n)
	
	case nodes.MinMaxExpr:
		return convertMinMaxExpr(&n)
	
	case nodes.MultiAssignRef:
		return convertMultiAssignRef(&n)
	
	case nodes.NamedArgExpr:
		return convertNamedArgExpr(&n)
	
	case nodes.NextValueExpr:
		return convertNextValueExpr(&n)
	
	case nodes.NotifyStmt:
		return convertNotifyStmt(&n)
	
	case nodes.Null:
		return convertNull(&n)
	
	case nodes.NullTest:
		return convertNullTest(&n)
	
	case nodes.ObjectWithArgs:
		return convertObjectWithArgs(&n)
	
	case nodes.OnConflictClause:
		return convertOnConflictClause(&n)
	
	case nodes.OnConflictExpr:
		return convertOnConflictExpr(&n)
	
	case nodes.OpExpr:
		return convertOpExpr(&n)
	
	case nodes.Param:
		return convertParam(&n)
	
	case nodes.ParamExecData:
		return convertParamExecData(&n)
	
	case nodes.ParamExternData:
		return convertParamExternData(&n)
	
	case nodes.ParamListInfoData:
		return convertParamListInfoData(&n)
	
	case nodes.ParamRef:
		return convertParamRef(&n)
	
	case nodes.PartitionBoundSpec:
		return convertPartitionBoundSpec(&n)
	
	case nodes.PartitionCmd:
		return convertPartitionCmd(&n)
	
	case nodes.PartitionElem:
		return convertPartitionElem(&n)
	
	case nodes.PartitionRangeDatum:
		return convertPartitionRangeDatum(&n)
	
	case nodes.PartitionSpec:
		return convertPartitionSpec(&n)
	
	case nodes.PrepareStmt:
		return convertPrepareStmt(&n)
	
	case nodes.Query:
		return convertQuery(&n)
	
	case nodes.RangeFunction:
		return convertRangeFunction(&n)
	
	case nodes.RangeSubselect:
		return convertRangeSubselect(&n)
	
	case nodes.RangeTableFunc:
		return convertRangeTableFunc(&n)
	
	case nodes.RangeTableFuncCol:
		return convertRangeTableFuncCol(&n)
	
	case nodes.RangeTableSample:
		return convertRangeTableSample(&n)
	
	case nodes.RangeTblEntry:
		return convertRangeTblEntry(&n)
	
	case nodes.RangeTblFunction:
		return convertRangeTblFunction(&n)
	
	case nodes.RangeTblRef:
		return convertRangeTblRef(&n)
	
	case nodes.RangeVar:
		return convertRangeVar(&n)
	
	case nodes.RawStmt:
		return convertRawStmt(&n)
	
	case nodes.ReassignOwnedStmt:
		return convertReassignOwnedStmt(&n)
	
	case nodes.RefreshMatViewStmt:
		return convertRefreshMatViewStmt(&n)
	
	case nodes.ReindexStmt:
		return convertReindexStmt(&n)
	
	case nodes.RelabelType:
		return convertRelabelType(&n)
	
	case nodes.RenameStmt:
		return convertRenameStmt(&n)
	
	case nodes.ReplicaIdentityStmt:
		return convertReplicaIdentityStmt(&n)
	
	case nodes.ResTarget:
		return convertResTarget(&n)
	
	case nodes.RoleSpec:
		return convertRoleSpec(&n)
	
	case nodes.RowCompareExpr:
		return convertRowCompareExpr(&n)
	
	case nodes.RowExpr:
		return convertRowExpr(&n)
	
	case nodes.RowMarkClause:
		return convertRowMarkClause(&n)
	
	case nodes.RuleStmt:
		return convertRuleStmt(&n)
	
	case nodes.SQLValueFunction:
		return convertSQLValueFunction(&n)
	
	case nodes.ScalarArrayOpExpr:
		return convertScalarArrayOpExpr(&n)
	
	case nodes.SecLabelStmt:
		return convertSecLabelStmt(&n)
	
	case nodes.SelectStmt:
		return convertSelectStmt(&n)
	
	case nodes.SetOperationStmt:
		return convertSetOperationStmt(&n)
	
	case nodes.SetToDefault:
		return convertSetToDefault(&n)
	
	case nodes.SortBy:
		return convertSortBy(&n)
	
	case nodes.SortGroupClause:
		return convertSortGroupClause(&n)
	
	case nodes.String:
		return convertString(&n)
	
	case nodes.SubLink:
		return convertSubLink(&n)
	
	case nodes.SubPlan:
		return convertSubPlan(&n)
	
	case nodes.TableFunc:
		return convertTableFunc(&n)
	
	case nodes.TableLikeClause:
		return convertTableLikeClause(&n)
	
	case nodes.TableSampleClause:
		return convertTableSampleClause(&n)
	
	case nodes.TargetEntry:
		return convertTargetEntry(&n)
	
	case nodes.TransactionStmt:
		return convertTransactionStmt(&n)
	
	case nodes.TriggerTransition:
		return convertTriggerTransition(&n)
	
	case nodes.TruncateStmt:
		return convertTruncateStmt(&n)
	
	case nodes.TypeCast:
		return convertTypeCast(&n)
	
	case nodes.TypeName:
		return convertTypeName(&n)
	
	case nodes.UnlistenStmt:
		return convertUnlistenStmt(&n)
	
	case nodes.UpdateStmt:
		return convertUpdateStmt(&n)
	
	case nodes.VacuumStmt:
		return convertVacuumStmt(&n)
	
	case nodes.Var:
		return convertVar(&n)
	
	case nodes.VariableSetStmt:
		return convertVariableSetStmt(&n)
	
	case nodes.VariableShowStmt:
		return convertVariableShowStmt(&n)
	
	case nodes.ViewStmt:
		return convertViewStmt(&n)
	
	case nodes.WindowClause:
		return convertWindowClause(&n)
	
	case nodes.WindowDef:
		return convertWindowDef(&n)
	
	case nodes.WindowFunc:
		return convertWindowFunc(&n)
	
	case nodes.WithCheckOption:
		return convertWithCheckOption(&n)
	
	case nodes.WithClause:
		return convertWithClause(&n)
	
	case nodes.XmlExpr:
		return convertXmlExpr(&n)
	
	case nodes.XmlSerialize:
		return convertXmlSerialize(&n)
	
	default:
		return &ast.TODO{}
	}
}
