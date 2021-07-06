// Code generated from SQLParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // SQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSQLParserListener is a complete listener for a parse tree produced by SQLParser.
type BaseSQLParserListener struct{}

var _ SQLParserListener = &BaseSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSql is called when production sql is entered.
func (s *BaseSQLParserListener) EnterSql(ctx *SqlContext) {}

// ExitSql is called when production sql is exited.
func (s *BaseSQLParserListener) ExitSql(ctx *SqlContext) {}

// EnterQname_parser is called when production qname_parser is entered.
func (s *BaseSQLParserListener) EnterQname_parser(ctx *Qname_parserContext) {}

// ExitQname_parser is called when production qname_parser is exited.
func (s *BaseSQLParserListener) ExitQname_parser(ctx *Qname_parserContext) {}

// EnterFunction_args_parser is called when production function_args_parser is entered.
func (s *BaseSQLParserListener) EnterFunction_args_parser(ctx *Function_args_parserContext) {}

// ExitFunction_args_parser is called when production function_args_parser is exited.
func (s *BaseSQLParserListener) ExitFunction_args_parser(ctx *Function_args_parserContext) {}

// EnterVex_eof is called when production vex_eof is entered.
func (s *BaseSQLParserListener) EnterVex_eof(ctx *Vex_eofContext) {}

// ExitVex_eof is called when production vex_eof is exited.
func (s *BaseSQLParserListener) ExitVex_eof(ctx *Vex_eofContext) {}

// EnterPlpgsql_function is called when production plpgsql_function is entered.
func (s *BaseSQLParserListener) EnterPlpgsql_function(ctx *Plpgsql_functionContext) {}

// ExitPlpgsql_function is called when production plpgsql_function is exited.
func (s *BaseSQLParserListener) ExitPlpgsql_function(ctx *Plpgsql_functionContext) {}

// EnterPlpgsql_function_test_list is called when production plpgsql_function_test_list is entered.
func (s *BaseSQLParserListener) EnterPlpgsql_function_test_list(ctx *Plpgsql_function_test_listContext) {
}

// ExitPlpgsql_function_test_list is called when production plpgsql_function_test_list is exited.
func (s *BaseSQLParserListener) ExitPlpgsql_function_test_list(ctx *Plpgsql_function_test_listContext) {
}

// EnterStatement is called when production statement is entered.
func (s *BaseSQLParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSQLParserListener) ExitStatement(ctx *StatementContext) {}

// EnterData_statement is called when production data_statement is entered.
func (s *BaseSQLParserListener) EnterData_statement(ctx *Data_statementContext) {}

// ExitData_statement is called when production data_statement is exited.
func (s *BaseSQLParserListener) ExitData_statement(ctx *Data_statementContext) {}

// EnterScript_statement is called when production script_statement is entered.
func (s *BaseSQLParserListener) EnterScript_statement(ctx *Script_statementContext) {}

// ExitScript_statement is called when production script_statement is exited.
func (s *BaseSQLParserListener) ExitScript_statement(ctx *Script_statementContext) {}

// EnterScript_transaction is called when production script_transaction is entered.
func (s *BaseSQLParserListener) EnterScript_transaction(ctx *Script_transactionContext) {}

// ExitScript_transaction is called when production script_transaction is exited.
func (s *BaseSQLParserListener) ExitScript_transaction(ctx *Script_transactionContext) {}

// EnterTransaction_mode is called when production transaction_mode is entered.
func (s *BaseSQLParserListener) EnterTransaction_mode(ctx *Transaction_modeContext) {}

// ExitTransaction_mode is called when production transaction_mode is exited.
func (s *BaseSQLParserListener) ExitTransaction_mode(ctx *Transaction_modeContext) {}

// EnterLock_table is called when production lock_table is entered.
func (s *BaseSQLParserListener) EnterLock_table(ctx *Lock_tableContext) {}

// ExitLock_table is called when production lock_table is exited.
func (s *BaseSQLParserListener) ExitLock_table(ctx *Lock_tableContext) {}

// EnterLock_mode is called when production lock_mode is entered.
func (s *BaseSQLParserListener) EnterLock_mode(ctx *Lock_modeContext) {}

// ExitLock_mode is called when production lock_mode is exited.
func (s *BaseSQLParserListener) ExitLock_mode(ctx *Lock_modeContext) {}

// EnterScript_additional is called when production script_additional is entered.
func (s *BaseSQLParserListener) EnterScript_additional(ctx *Script_additionalContext) {}

// ExitScript_additional is called when production script_additional is exited.
func (s *BaseSQLParserListener) ExitScript_additional(ctx *Script_additionalContext) {}

// EnterAdditional_statement is called when production additional_statement is entered.
func (s *BaseSQLParserListener) EnterAdditional_statement(ctx *Additional_statementContext) {}

// ExitAdditional_statement is called when production additional_statement is exited.
func (s *BaseSQLParserListener) ExitAdditional_statement(ctx *Additional_statementContext) {}

// EnterExplain_statement is called when production explain_statement is entered.
func (s *BaseSQLParserListener) EnterExplain_statement(ctx *Explain_statementContext) {}

// ExitExplain_statement is called when production explain_statement is exited.
func (s *BaseSQLParserListener) ExitExplain_statement(ctx *Explain_statementContext) {}

// EnterExplain_query is called when production explain_query is entered.
func (s *BaseSQLParserListener) EnterExplain_query(ctx *Explain_queryContext) {}

// ExitExplain_query is called when production explain_query is exited.
func (s *BaseSQLParserListener) ExitExplain_query(ctx *Explain_queryContext) {}

// EnterExecute_statement is called when production execute_statement is entered.
func (s *BaseSQLParserListener) EnterExecute_statement(ctx *Execute_statementContext) {}

// ExitExecute_statement is called when production execute_statement is exited.
func (s *BaseSQLParserListener) ExitExecute_statement(ctx *Execute_statementContext) {}

// EnterDeclare_statement is called when production declare_statement is entered.
func (s *BaseSQLParserListener) EnterDeclare_statement(ctx *Declare_statementContext) {}

// ExitDeclare_statement is called when production declare_statement is exited.
func (s *BaseSQLParserListener) ExitDeclare_statement(ctx *Declare_statementContext) {}

// EnterShow_statement is called when production show_statement is entered.
func (s *BaseSQLParserListener) EnterShow_statement(ctx *Show_statementContext) {}

// ExitShow_statement is called when production show_statement is exited.
func (s *BaseSQLParserListener) ExitShow_statement(ctx *Show_statementContext) {}

// EnterExplain_option is called when production explain_option is entered.
func (s *BaseSQLParserListener) EnterExplain_option(ctx *Explain_optionContext) {}

// ExitExplain_option is called when production explain_option is exited.
func (s *BaseSQLParserListener) ExitExplain_option(ctx *Explain_optionContext) {}

// EnterUser_name is called when production user_name is entered.
func (s *BaseSQLParserListener) EnterUser_name(ctx *User_nameContext) {}

// ExitUser_name is called when production user_name is exited.
func (s *BaseSQLParserListener) ExitUser_name(ctx *User_nameContext) {}

// EnterTable_cols_list is called when production table_cols_list is entered.
func (s *BaseSQLParserListener) EnterTable_cols_list(ctx *Table_cols_listContext) {}

// ExitTable_cols_list is called when production table_cols_list is exited.
func (s *BaseSQLParserListener) ExitTable_cols_list(ctx *Table_cols_listContext) {}

// EnterTable_cols is called when production table_cols is entered.
func (s *BaseSQLParserListener) EnterTable_cols(ctx *Table_colsContext) {}

// ExitTable_cols is called when production table_cols is exited.
func (s *BaseSQLParserListener) ExitTable_cols(ctx *Table_colsContext) {}

// EnterVacuum_mode is called when production vacuum_mode is entered.
func (s *BaseSQLParserListener) EnterVacuum_mode(ctx *Vacuum_modeContext) {}

// ExitVacuum_mode is called when production vacuum_mode is exited.
func (s *BaseSQLParserListener) ExitVacuum_mode(ctx *Vacuum_modeContext) {}

// EnterVacuum_option is called when production vacuum_option is entered.
func (s *BaseSQLParserListener) EnterVacuum_option(ctx *Vacuum_optionContext) {}

// ExitVacuum_option is called when production vacuum_option is exited.
func (s *BaseSQLParserListener) ExitVacuum_option(ctx *Vacuum_optionContext) {}

// EnterAnalyze_mode is called when production analyze_mode is entered.
func (s *BaseSQLParserListener) EnterAnalyze_mode(ctx *Analyze_modeContext) {}

// ExitAnalyze_mode is called when production analyze_mode is exited.
func (s *BaseSQLParserListener) ExitAnalyze_mode(ctx *Analyze_modeContext) {}

// EnterBoolean_value is called when production boolean_value is entered.
func (s *BaseSQLParserListener) EnterBoolean_value(ctx *Boolean_valueContext) {}

// ExitBoolean_value is called when production boolean_value is exited.
func (s *BaseSQLParserListener) ExitBoolean_value(ctx *Boolean_valueContext) {}

// EnterFetch_move_direction is called when production fetch_move_direction is entered.
func (s *BaseSQLParserListener) EnterFetch_move_direction(ctx *Fetch_move_directionContext) {}

// ExitFetch_move_direction is called when production fetch_move_direction is exited.
func (s *BaseSQLParserListener) ExitFetch_move_direction(ctx *Fetch_move_directionContext) {}

// EnterSchema_statement is called when production schema_statement is entered.
func (s *BaseSQLParserListener) EnterSchema_statement(ctx *Schema_statementContext) {}

// ExitSchema_statement is called when production schema_statement is exited.
func (s *BaseSQLParserListener) ExitSchema_statement(ctx *Schema_statementContext) {}

// EnterSchema_create is called when production schema_create is entered.
func (s *BaseSQLParserListener) EnterSchema_create(ctx *Schema_createContext) {}

// ExitSchema_create is called when production schema_create is exited.
func (s *BaseSQLParserListener) ExitSchema_create(ctx *Schema_createContext) {}

// EnterSchema_alter is called when production schema_alter is entered.
func (s *BaseSQLParserListener) EnterSchema_alter(ctx *Schema_alterContext) {}

// ExitSchema_alter is called when production schema_alter is exited.
func (s *BaseSQLParserListener) ExitSchema_alter(ctx *Schema_alterContext) {}

// EnterSchema_drop is called when production schema_drop is entered.
func (s *BaseSQLParserListener) EnterSchema_drop(ctx *Schema_dropContext) {}

// ExitSchema_drop is called when production schema_drop is exited.
func (s *BaseSQLParserListener) ExitSchema_drop(ctx *Schema_dropContext) {}

// EnterSchema_import is called when production schema_import is entered.
func (s *BaseSQLParserListener) EnterSchema_import(ctx *Schema_importContext) {}

// ExitSchema_import is called when production schema_import is exited.
func (s *BaseSQLParserListener) ExitSchema_import(ctx *Schema_importContext) {}

// EnterAlter_function_statement is called when production alter_function_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_function_statement(ctx *Alter_function_statementContext) {}

// ExitAlter_function_statement is called when production alter_function_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_function_statement(ctx *Alter_function_statementContext) {}

// EnterAlter_aggregate_statement is called when production alter_aggregate_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_aggregate_statement(ctx *Alter_aggregate_statementContext) {
}

// ExitAlter_aggregate_statement is called when production alter_aggregate_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_aggregate_statement(ctx *Alter_aggregate_statementContext) {
}

// EnterAlter_extension_statement is called when production alter_extension_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_extension_statement(ctx *Alter_extension_statementContext) {
}

// ExitAlter_extension_statement is called when production alter_extension_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_extension_statement(ctx *Alter_extension_statementContext) {
}

// EnterAlter_extension_action is called when production alter_extension_action is entered.
func (s *BaseSQLParserListener) EnterAlter_extension_action(ctx *Alter_extension_actionContext) {}

// ExitAlter_extension_action is called when production alter_extension_action is exited.
func (s *BaseSQLParserListener) ExitAlter_extension_action(ctx *Alter_extension_actionContext) {}

// EnterExtension_member_object is called when production extension_member_object is entered.
func (s *BaseSQLParserListener) EnterExtension_member_object(ctx *Extension_member_objectContext) {}

// ExitExtension_member_object is called when production extension_member_object is exited.
func (s *BaseSQLParserListener) ExitExtension_member_object(ctx *Extension_member_objectContext) {}

// EnterAlter_schema_statement is called when production alter_schema_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_schema_statement(ctx *Alter_schema_statementContext) {}

// ExitAlter_schema_statement is called when production alter_schema_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_schema_statement(ctx *Alter_schema_statementContext) {}

// EnterAlter_language_statement is called when production alter_language_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_language_statement(ctx *Alter_language_statementContext) {}

// ExitAlter_language_statement is called when production alter_language_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_language_statement(ctx *Alter_language_statementContext) {}

// EnterAlter_table_statement is called when production alter_table_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_table_statement(ctx *Alter_table_statementContext) {}

// ExitAlter_table_statement is called when production alter_table_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_table_statement(ctx *Alter_table_statementContext) {}

// EnterTable_action is called when production table_action is entered.
func (s *BaseSQLParserListener) EnterTable_action(ctx *Table_actionContext) {}

// ExitTable_action is called when production table_action is exited.
func (s *BaseSQLParserListener) ExitTable_action(ctx *Table_actionContext) {}

// EnterColumn_action is called when production column_action is entered.
func (s *BaseSQLParserListener) EnterColumn_action(ctx *Column_actionContext) {}

// ExitColumn_action is called when production column_action is exited.
func (s *BaseSQLParserListener) ExitColumn_action(ctx *Column_actionContext) {}

// EnterIdentity_body is called when production identity_body is entered.
func (s *BaseSQLParserListener) EnterIdentity_body(ctx *Identity_bodyContext) {}

// ExitIdentity_body is called when production identity_body is exited.
func (s *BaseSQLParserListener) ExitIdentity_body(ctx *Identity_bodyContext) {}

// EnterAlter_identity is called when production alter_identity is entered.
func (s *BaseSQLParserListener) EnterAlter_identity(ctx *Alter_identityContext) {}

// ExitAlter_identity is called when production alter_identity is exited.
func (s *BaseSQLParserListener) ExitAlter_identity(ctx *Alter_identityContext) {}

// EnterStorage_option is called when production storage_option is entered.
func (s *BaseSQLParserListener) EnterStorage_option(ctx *Storage_optionContext) {}

// ExitStorage_option is called when production storage_option is exited.
func (s *BaseSQLParserListener) ExitStorage_option(ctx *Storage_optionContext) {}

// EnterValidate_constraint is called when production validate_constraint is entered.
func (s *BaseSQLParserListener) EnterValidate_constraint(ctx *Validate_constraintContext) {}

// ExitValidate_constraint is called when production validate_constraint is exited.
func (s *BaseSQLParserListener) ExitValidate_constraint(ctx *Validate_constraintContext) {}

// EnterDrop_constraint is called when production drop_constraint is entered.
func (s *BaseSQLParserListener) EnterDrop_constraint(ctx *Drop_constraintContext) {}

// ExitDrop_constraint is called when production drop_constraint is exited.
func (s *BaseSQLParserListener) ExitDrop_constraint(ctx *Drop_constraintContext) {}

// EnterTable_deferrable is called when production table_deferrable is entered.
func (s *BaseSQLParserListener) EnterTable_deferrable(ctx *Table_deferrableContext) {}

// ExitTable_deferrable is called when production table_deferrable is exited.
func (s *BaseSQLParserListener) ExitTable_deferrable(ctx *Table_deferrableContext) {}

// EnterTable_initialy_immed is called when production table_initialy_immed is entered.
func (s *BaseSQLParserListener) EnterTable_initialy_immed(ctx *Table_initialy_immedContext) {}

// ExitTable_initialy_immed is called when production table_initialy_immed is exited.
func (s *BaseSQLParserListener) ExitTable_initialy_immed(ctx *Table_initialy_immedContext) {}

// EnterFunction_actions_common is called when production function_actions_common is entered.
func (s *BaseSQLParserListener) EnterFunction_actions_common(ctx *Function_actions_commonContext) {}

// ExitFunction_actions_common is called when production function_actions_common is exited.
func (s *BaseSQLParserListener) ExitFunction_actions_common(ctx *Function_actions_commonContext) {}

// EnterFunction_def is called when production function_def is entered.
func (s *BaseSQLParserListener) EnterFunction_def(ctx *Function_defContext) {}

// ExitFunction_def is called when production function_def is exited.
func (s *BaseSQLParserListener) ExitFunction_def(ctx *Function_defContext) {}

// EnterAlter_index_statement is called when production alter_index_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_index_statement(ctx *Alter_index_statementContext) {}

// ExitAlter_index_statement is called when production alter_index_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_index_statement(ctx *Alter_index_statementContext) {}

// EnterIndex_def_action is called when production index_def_action is entered.
func (s *BaseSQLParserListener) EnterIndex_def_action(ctx *Index_def_actionContext) {}

// ExitIndex_def_action is called when production index_def_action is exited.
func (s *BaseSQLParserListener) ExitIndex_def_action(ctx *Index_def_actionContext) {}

// EnterAlter_default_privileges_statement is called when production alter_default_privileges_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_default_privileges_statement(ctx *Alter_default_privileges_statementContext) {
}

// ExitAlter_default_privileges_statement is called when production alter_default_privileges_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_default_privileges_statement(ctx *Alter_default_privileges_statementContext) {
}

// EnterAbbreviated_grant_or_revoke is called when production abbreviated_grant_or_revoke is entered.
func (s *BaseSQLParserListener) EnterAbbreviated_grant_or_revoke(ctx *Abbreviated_grant_or_revokeContext) {
}

// ExitAbbreviated_grant_or_revoke is called when production abbreviated_grant_or_revoke is exited.
func (s *BaseSQLParserListener) ExitAbbreviated_grant_or_revoke(ctx *Abbreviated_grant_or_revokeContext) {
}

// EnterGrant_option_for is called when production grant_option_for is entered.
func (s *BaseSQLParserListener) EnterGrant_option_for(ctx *Grant_option_forContext) {}

// ExitGrant_option_for is called when production grant_option_for is exited.
func (s *BaseSQLParserListener) ExitGrant_option_for(ctx *Grant_option_forContext) {}

// EnterAlter_sequence_statement is called when production alter_sequence_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_sequence_statement(ctx *Alter_sequence_statementContext) {}

// ExitAlter_sequence_statement is called when production alter_sequence_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_sequence_statement(ctx *Alter_sequence_statementContext) {}

// EnterAlter_view_statement is called when production alter_view_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_view_statement(ctx *Alter_view_statementContext) {}

// ExitAlter_view_statement is called when production alter_view_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_view_statement(ctx *Alter_view_statementContext) {}

// EnterAlter_view_action is called when production alter_view_action is entered.
func (s *BaseSQLParserListener) EnterAlter_view_action(ctx *Alter_view_actionContext) {}

// ExitAlter_view_action is called when production alter_view_action is exited.
func (s *BaseSQLParserListener) ExitAlter_view_action(ctx *Alter_view_actionContext) {}

// EnterAlter_materialized_view_statement is called when production alter_materialized_view_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_materialized_view_statement(ctx *Alter_materialized_view_statementContext) {
}

// ExitAlter_materialized_view_statement is called when production alter_materialized_view_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_materialized_view_statement(ctx *Alter_materialized_view_statementContext) {
}

// EnterAlter_materialized_view_action is called when production alter_materialized_view_action is entered.
func (s *BaseSQLParserListener) EnterAlter_materialized_view_action(ctx *Alter_materialized_view_actionContext) {
}

// ExitAlter_materialized_view_action is called when production alter_materialized_view_action is exited.
func (s *BaseSQLParserListener) ExitAlter_materialized_view_action(ctx *Alter_materialized_view_actionContext) {
}

// EnterMaterialized_view_action is called when production materialized_view_action is entered.
func (s *BaseSQLParserListener) EnterMaterialized_view_action(ctx *Materialized_view_actionContext) {}

// ExitMaterialized_view_action is called when production materialized_view_action is exited.
func (s *BaseSQLParserListener) ExitMaterialized_view_action(ctx *Materialized_view_actionContext) {}

// EnterAlter_event_trigger_statement is called when production alter_event_trigger_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_event_trigger_statement(ctx *Alter_event_trigger_statementContext) {
}

// ExitAlter_event_trigger_statement is called when production alter_event_trigger_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_event_trigger_statement(ctx *Alter_event_trigger_statementContext) {
}

// EnterAlter_event_trigger_action is called when production alter_event_trigger_action is entered.
func (s *BaseSQLParserListener) EnterAlter_event_trigger_action(ctx *Alter_event_trigger_actionContext) {
}

// ExitAlter_event_trigger_action is called when production alter_event_trigger_action is exited.
func (s *BaseSQLParserListener) ExitAlter_event_trigger_action(ctx *Alter_event_trigger_actionContext) {
}

// EnterAlter_type_statement is called when production alter_type_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_type_statement(ctx *Alter_type_statementContext) {}

// ExitAlter_type_statement is called when production alter_type_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_type_statement(ctx *Alter_type_statementContext) {}

// EnterAlter_domain_statement is called when production alter_domain_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_domain_statement(ctx *Alter_domain_statementContext) {}

// ExitAlter_domain_statement is called when production alter_domain_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_domain_statement(ctx *Alter_domain_statementContext) {}

// EnterAlter_server_statement is called when production alter_server_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_server_statement(ctx *Alter_server_statementContext) {}

// ExitAlter_server_statement is called when production alter_server_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_server_statement(ctx *Alter_server_statementContext) {}

// EnterAlter_server_action is called when production alter_server_action is entered.
func (s *BaseSQLParserListener) EnterAlter_server_action(ctx *Alter_server_actionContext) {}

// ExitAlter_server_action is called when production alter_server_action is exited.
func (s *BaseSQLParserListener) ExitAlter_server_action(ctx *Alter_server_actionContext) {}

// EnterAlter_fts_statement is called when production alter_fts_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_fts_statement(ctx *Alter_fts_statementContext) {}

// ExitAlter_fts_statement is called when production alter_fts_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_fts_statement(ctx *Alter_fts_statementContext) {}

// EnterAlter_fts_configuration is called when production alter_fts_configuration is entered.
func (s *BaseSQLParserListener) EnterAlter_fts_configuration(ctx *Alter_fts_configurationContext) {}

// ExitAlter_fts_configuration is called when production alter_fts_configuration is exited.
func (s *BaseSQLParserListener) ExitAlter_fts_configuration(ctx *Alter_fts_configurationContext) {}

// EnterType_action is called when production type_action is entered.
func (s *BaseSQLParserListener) EnterType_action(ctx *Type_actionContext) {}

// ExitType_action is called when production type_action is exited.
func (s *BaseSQLParserListener) ExitType_action(ctx *Type_actionContext) {}

// EnterType_property is called when production type_property is entered.
func (s *BaseSQLParserListener) EnterType_property(ctx *Type_propertyContext) {}

// ExitType_property is called when production type_property is exited.
func (s *BaseSQLParserListener) ExitType_property(ctx *Type_propertyContext) {}

// EnterSet_def_column is called when production set_def_column is entered.
func (s *BaseSQLParserListener) EnterSet_def_column(ctx *Set_def_columnContext) {}

// ExitSet_def_column is called when production set_def_column is exited.
func (s *BaseSQLParserListener) ExitSet_def_column(ctx *Set_def_columnContext) {}

// EnterDrop_def is called when production drop_def is entered.
func (s *BaseSQLParserListener) EnterDrop_def(ctx *Drop_defContext) {}

// ExitDrop_def is called when production drop_def is exited.
func (s *BaseSQLParserListener) ExitDrop_def(ctx *Drop_defContext) {}

// EnterCreate_index_statement is called when production create_index_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_index_statement(ctx *Create_index_statementContext) {}

// ExitCreate_index_statement is called when production create_index_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_index_statement(ctx *Create_index_statementContext) {}

// EnterIndex_rest is called when production index_rest is entered.
func (s *BaseSQLParserListener) EnterIndex_rest(ctx *Index_restContext) {}

// ExitIndex_rest is called when production index_rest is exited.
func (s *BaseSQLParserListener) ExitIndex_rest(ctx *Index_restContext) {}

// EnterIndex_sort is called when production index_sort is entered.
func (s *BaseSQLParserListener) EnterIndex_sort(ctx *Index_sortContext) {}

// ExitIndex_sort is called when production index_sort is exited.
func (s *BaseSQLParserListener) ExitIndex_sort(ctx *Index_sortContext) {}

// EnterIndex_column is called when production index_column is entered.
func (s *BaseSQLParserListener) EnterIndex_column(ctx *Index_columnContext) {}

// ExitIndex_column is called when production index_column is exited.
func (s *BaseSQLParserListener) ExitIndex_column(ctx *Index_columnContext) {}

// EnterIncluding_index is called when production including_index is entered.
func (s *BaseSQLParserListener) EnterIncluding_index(ctx *Including_indexContext) {}

// ExitIncluding_index is called when production including_index is exited.
func (s *BaseSQLParserListener) ExitIncluding_index(ctx *Including_indexContext) {}

// EnterIndex_where is called when production index_where is entered.
func (s *BaseSQLParserListener) EnterIndex_where(ctx *Index_whereContext) {}

// ExitIndex_where is called when production index_where is exited.
func (s *BaseSQLParserListener) ExitIndex_where(ctx *Index_whereContext) {}

// EnterCreate_extension_statement is called when production create_extension_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_extension_statement(ctx *Create_extension_statementContext) {
}

// ExitCreate_extension_statement is called when production create_extension_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_extension_statement(ctx *Create_extension_statementContext) {
}

// EnterCreate_language_statement is called when production create_language_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_language_statement(ctx *Create_language_statementContext) {
}

// ExitCreate_language_statement is called when production create_language_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_language_statement(ctx *Create_language_statementContext) {
}

// EnterCreate_event_trigger_statement is called when production create_event_trigger_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_event_trigger_statement(ctx *Create_event_trigger_statementContext) {
}

// ExitCreate_event_trigger_statement is called when production create_event_trigger_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_event_trigger_statement(ctx *Create_event_trigger_statementContext) {
}

// EnterCreate_type_statement is called when production create_type_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_type_statement(ctx *Create_type_statementContext) {}

// ExitCreate_type_statement is called when production create_type_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_type_statement(ctx *Create_type_statementContext) {}

// EnterCreate_domain_statement is called when production create_domain_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_domain_statement(ctx *Create_domain_statementContext) {}

// ExitCreate_domain_statement is called when production create_domain_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_domain_statement(ctx *Create_domain_statementContext) {}

// EnterCreate_server_statement is called when production create_server_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_server_statement(ctx *Create_server_statementContext) {}

// ExitCreate_server_statement is called when production create_server_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_server_statement(ctx *Create_server_statementContext) {}

// EnterCreate_fts_dictionary_statement is called when production create_fts_dictionary_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_fts_dictionary_statement(ctx *Create_fts_dictionary_statementContext) {
}

// ExitCreate_fts_dictionary_statement is called when production create_fts_dictionary_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_fts_dictionary_statement(ctx *Create_fts_dictionary_statementContext) {
}

// EnterOption_with_value is called when production option_with_value is entered.
func (s *BaseSQLParserListener) EnterOption_with_value(ctx *Option_with_valueContext) {}

// ExitOption_with_value is called when production option_with_value is exited.
func (s *BaseSQLParserListener) ExitOption_with_value(ctx *Option_with_valueContext) {}

// EnterCreate_fts_configuration_statement is called when production create_fts_configuration_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_fts_configuration_statement(ctx *Create_fts_configuration_statementContext) {
}

// ExitCreate_fts_configuration_statement is called when production create_fts_configuration_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_fts_configuration_statement(ctx *Create_fts_configuration_statementContext) {
}

// EnterCreate_fts_template_statement is called when production create_fts_template_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_fts_template_statement(ctx *Create_fts_template_statementContext) {
}

// ExitCreate_fts_template_statement is called when production create_fts_template_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_fts_template_statement(ctx *Create_fts_template_statementContext) {
}

// EnterCreate_fts_parser_statement is called when production create_fts_parser_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_fts_parser_statement(ctx *Create_fts_parser_statementContext) {
}

// ExitCreate_fts_parser_statement is called when production create_fts_parser_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_fts_parser_statement(ctx *Create_fts_parser_statementContext) {
}

// EnterCreate_collation_statement is called when production create_collation_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_collation_statement(ctx *Create_collation_statementContext) {
}

// ExitCreate_collation_statement is called when production create_collation_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_collation_statement(ctx *Create_collation_statementContext) {
}

// EnterAlter_collation_statement is called when production alter_collation_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_collation_statement(ctx *Alter_collation_statementContext) {
}

// ExitAlter_collation_statement is called when production alter_collation_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_collation_statement(ctx *Alter_collation_statementContext) {
}

// EnterCollation_option is called when production collation_option is entered.
func (s *BaseSQLParserListener) EnterCollation_option(ctx *Collation_optionContext) {}

// ExitCollation_option is called when production collation_option is exited.
func (s *BaseSQLParserListener) ExitCollation_option(ctx *Collation_optionContext) {}

// EnterCreate_user_mapping_statement is called when production create_user_mapping_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_user_mapping_statement(ctx *Create_user_mapping_statementContext) {
}

// ExitCreate_user_mapping_statement is called when production create_user_mapping_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_user_mapping_statement(ctx *Create_user_mapping_statementContext) {
}

// EnterAlter_user_mapping_statement is called when production alter_user_mapping_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_user_mapping_statement(ctx *Alter_user_mapping_statementContext) {
}

// ExitAlter_user_mapping_statement is called when production alter_user_mapping_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_user_mapping_statement(ctx *Alter_user_mapping_statementContext) {
}

// EnterAlter_user_or_role_statement is called when production alter_user_or_role_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_user_or_role_statement(ctx *Alter_user_or_role_statementContext) {
}

// ExitAlter_user_or_role_statement is called when production alter_user_or_role_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_user_or_role_statement(ctx *Alter_user_or_role_statementContext) {
}

// EnterAlter_user_or_role_set_reset is called when production alter_user_or_role_set_reset is entered.
func (s *BaseSQLParserListener) EnterAlter_user_or_role_set_reset(ctx *Alter_user_or_role_set_resetContext) {
}

// ExitAlter_user_or_role_set_reset is called when production alter_user_or_role_set_reset is exited.
func (s *BaseSQLParserListener) ExitAlter_user_or_role_set_reset(ctx *Alter_user_or_role_set_resetContext) {
}

// EnterSet_reset_parameter is called when production set_reset_parameter is entered.
func (s *BaseSQLParserListener) EnterSet_reset_parameter(ctx *Set_reset_parameterContext) {}

// ExitSet_reset_parameter is called when production set_reset_parameter is exited.
func (s *BaseSQLParserListener) ExitSet_reset_parameter(ctx *Set_reset_parameterContext) {}

// EnterAlter_group_statement is called when production alter_group_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_group_statement(ctx *Alter_group_statementContext) {}

// ExitAlter_group_statement is called when production alter_group_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_group_statement(ctx *Alter_group_statementContext) {}

// EnterAlter_group_action is called when production alter_group_action is entered.
func (s *BaseSQLParserListener) EnterAlter_group_action(ctx *Alter_group_actionContext) {}

// ExitAlter_group_action is called when production alter_group_action is exited.
func (s *BaseSQLParserListener) ExitAlter_group_action(ctx *Alter_group_actionContext) {}

// EnterAlter_tablespace_statement is called when production alter_tablespace_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_tablespace_statement(ctx *Alter_tablespace_statementContext) {
}

// ExitAlter_tablespace_statement is called when production alter_tablespace_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_tablespace_statement(ctx *Alter_tablespace_statementContext) {
}

// EnterAlter_owner_statement is called when production alter_owner_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_owner_statement(ctx *Alter_owner_statementContext) {}

// ExitAlter_owner_statement is called when production alter_owner_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_owner_statement(ctx *Alter_owner_statementContext) {}

// EnterAlter_tablespace_action is called when production alter_tablespace_action is entered.
func (s *BaseSQLParserListener) EnterAlter_tablespace_action(ctx *Alter_tablespace_actionContext) {}

// ExitAlter_tablespace_action is called when production alter_tablespace_action is exited.
func (s *BaseSQLParserListener) ExitAlter_tablespace_action(ctx *Alter_tablespace_actionContext) {}

// EnterAlter_statistics_statement is called when production alter_statistics_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_statistics_statement(ctx *Alter_statistics_statementContext) {
}

// ExitAlter_statistics_statement is called when production alter_statistics_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_statistics_statement(ctx *Alter_statistics_statementContext) {
}

// EnterSet_statistics is called when production set_statistics is entered.
func (s *BaseSQLParserListener) EnterSet_statistics(ctx *Set_statisticsContext) {}

// ExitSet_statistics is called when production set_statistics is exited.
func (s *BaseSQLParserListener) ExitSet_statistics(ctx *Set_statisticsContext) {}

// EnterAlter_foreign_data_wrapper is called when production alter_foreign_data_wrapper is entered.
func (s *BaseSQLParserListener) EnterAlter_foreign_data_wrapper(ctx *Alter_foreign_data_wrapperContext) {
}

// ExitAlter_foreign_data_wrapper is called when production alter_foreign_data_wrapper is exited.
func (s *BaseSQLParserListener) ExitAlter_foreign_data_wrapper(ctx *Alter_foreign_data_wrapperContext) {
}

// EnterAlter_foreign_data_wrapper_action is called when production alter_foreign_data_wrapper_action is entered.
func (s *BaseSQLParserListener) EnterAlter_foreign_data_wrapper_action(ctx *Alter_foreign_data_wrapper_actionContext) {
}

// ExitAlter_foreign_data_wrapper_action is called when production alter_foreign_data_wrapper_action is exited.
func (s *BaseSQLParserListener) ExitAlter_foreign_data_wrapper_action(ctx *Alter_foreign_data_wrapper_actionContext) {
}

// EnterAlter_operator_statement is called when production alter_operator_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_operator_statement(ctx *Alter_operator_statementContext) {}

// ExitAlter_operator_statement is called when production alter_operator_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_operator_statement(ctx *Alter_operator_statementContext) {}

// EnterAlter_operator_action is called when production alter_operator_action is entered.
func (s *BaseSQLParserListener) EnterAlter_operator_action(ctx *Alter_operator_actionContext) {}

// ExitAlter_operator_action is called when production alter_operator_action is exited.
func (s *BaseSQLParserListener) ExitAlter_operator_action(ctx *Alter_operator_actionContext) {}

// EnterOperator_set_restrict_join is called when production operator_set_restrict_join is entered.
func (s *BaseSQLParserListener) EnterOperator_set_restrict_join(ctx *Operator_set_restrict_joinContext) {
}

// ExitOperator_set_restrict_join is called when production operator_set_restrict_join is exited.
func (s *BaseSQLParserListener) ExitOperator_set_restrict_join(ctx *Operator_set_restrict_joinContext) {
}

// EnterDrop_user_mapping_statement is called when production drop_user_mapping_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_user_mapping_statement(ctx *Drop_user_mapping_statementContext) {
}

// ExitDrop_user_mapping_statement is called when production drop_user_mapping_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_user_mapping_statement(ctx *Drop_user_mapping_statementContext) {
}

// EnterDrop_owned_statement is called when production drop_owned_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_owned_statement(ctx *Drop_owned_statementContext) {}

// ExitDrop_owned_statement is called when production drop_owned_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_owned_statement(ctx *Drop_owned_statementContext) {}

// EnterDrop_operator_statement is called when production drop_operator_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_operator_statement(ctx *Drop_operator_statementContext) {}

// ExitDrop_operator_statement is called when production drop_operator_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_operator_statement(ctx *Drop_operator_statementContext) {}

// EnterTarget_operator is called when production target_operator is entered.
func (s *BaseSQLParserListener) EnterTarget_operator(ctx *Target_operatorContext) {}

// ExitTarget_operator is called when production target_operator is exited.
func (s *BaseSQLParserListener) ExitTarget_operator(ctx *Target_operatorContext) {}

// EnterDomain_constraint is called when production domain_constraint is entered.
func (s *BaseSQLParserListener) EnterDomain_constraint(ctx *Domain_constraintContext) {}

// ExitDomain_constraint is called when production domain_constraint is exited.
func (s *BaseSQLParserListener) ExitDomain_constraint(ctx *Domain_constraintContext) {}

// EnterCreate_transform_statement is called when production create_transform_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_transform_statement(ctx *Create_transform_statementContext) {
}

// ExitCreate_transform_statement is called when production create_transform_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_transform_statement(ctx *Create_transform_statementContext) {
}

// EnterCreate_access_method_statement is called when production create_access_method_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_access_method_statement(ctx *Create_access_method_statementContext) {
}

// ExitCreate_access_method_statement is called when production create_access_method_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_access_method_statement(ctx *Create_access_method_statementContext) {
}

// EnterCreate_user_or_role_statement is called when production create_user_or_role_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_user_or_role_statement(ctx *Create_user_or_role_statementContext) {
}

// ExitCreate_user_or_role_statement is called when production create_user_or_role_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_user_or_role_statement(ctx *Create_user_or_role_statementContext) {
}

// EnterUser_or_role_option is called when production user_or_role_option is entered.
func (s *BaseSQLParserListener) EnterUser_or_role_option(ctx *User_or_role_optionContext) {}

// ExitUser_or_role_option is called when production user_or_role_option is exited.
func (s *BaseSQLParserListener) ExitUser_or_role_option(ctx *User_or_role_optionContext) {}

// EnterUser_or_role_option_for_alter is called when production user_or_role_option_for_alter is entered.
func (s *BaseSQLParserListener) EnterUser_or_role_option_for_alter(ctx *User_or_role_option_for_alterContext) {
}

// ExitUser_or_role_option_for_alter is called when production user_or_role_option_for_alter is exited.
func (s *BaseSQLParserListener) ExitUser_or_role_option_for_alter(ctx *User_or_role_option_for_alterContext) {
}

// EnterUser_or_role_or_group_common_option is called when production user_or_role_or_group_common_option is entered.
func (s *BaseSQLParserListener) EnterUser_or_role_or_group_common_option(ctx *User_or_role_or_group_common_optionContext) {
}

// ExitUser_or_role_or_group_common_option is called when production user_or_role_or_group_common_option is exited.
func (s *BaseSQLParserListener) ExitUser_or_role_or_group_common_option(ctx *User_or_role_or_group_common_optionContext) {
}

// EnterUser_or_role_common_option is called when production user_or_role_common_option is entered.
func (s *BaseSQLParserListener) EnterUser_or_role_common_option(ctx *User_or_role_common_optionContext) {
}

// ExitUser_or_role_common_option is called when production user_or_role_common_option is exited.
func (s *BaseSQLParserListener) ExitUser_or_role_common_option(ctx *User_or_role_common_optionContext) {
}

// EnterUser_or_role_or_group_option_for_create is called when production user_or_role_or_group_option_for_create is entered.
func (s *BaseSQLParserListener) EnterUser_or_role_or_group_option_for_create(ctx *User_or_role_or_group_option_for_createContext) {
}

// ExitUser_or_role_or_group_option_for_create is called when production user_or_role_or_group_option_for_create is exited.
func (s *BaseSQLParserListener) ExitUser_or_role_or_group_option_for_create(ctx *User_or_role_or_group_option_for_createContext) {
}

// EnterCreate_group_statement is called when production create_group_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_group_statement(ctx *Create_group_statementContext) {}

// ExitCreate_group_statement is called when production create_group_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_group_statement(ctx *Create_group_statementContext) {}

// EnterGroup_option is called when production group_option is entered.
func (s *BaseSQLParserListener) EnterGroup_option(ctx *Group_optionContext) {}

// ExitGroup_option is called when production group_option is exited.
func (s *BaseSQLParserListener) ExitGroup_option(ctx *Group_optionContext) {}

// EnterCreate_tablespace_statement is called when production create_tablespace_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_tablespace_statement(ctx *Create_tablespace_statementContext) {
}

// ExitCreate_tablespace_statement is called when production create_tablespace_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_tablespace_statement(ctx *Create_tablespace_statementContext) {
}

// EnterCreate_statistics_statement is called when production create_statistics_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_statistics_statement(ctx *Create_statistics_statementContext) {
}

// ExitCreate_statistics_statement is called when production create_statistics_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_statistics_statement(ctx *Create_statistics_statementContext) {
}

// EnterCreate_foreign_data_wrapper_statement is called when production create_foreign_data_wrapper_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_foreign_data_wrapper_statement(ctx *Create_foreign_data_wrapper_statementContext) {
}

// ExitCreate_foreign_data_wrapper_statement is called when production create_foreign_data_wrapper_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_foreign_data_wrapper_statement(ctx *Create_foreign_data_wrapper_statementContext) {
}

// EnterOption_without_equal is called when production option_without_equal is entered.
func (s *BaseSQLParserListener) EnterOption_without_equal(ctx *Option_without_equalContext) {}

// ExitOption_without_equal is called when production option_without_equal is exited.
func (s *BaseSQLParserListener) ExitOption_without_equal(ctx *Option_without_equalContext) {}

// EnterCreate_operator_statement is called when production create_operator_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_operator_statement(ctx *Create_operator_statementContext) {
}

// ExitCreate_operator_statement is called when production create_operator_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_operator_statement(ctx *Create_operator_statementContext) {
}

// EnterOperator_name is called when production operator_name is entered.
func (s *BaseSQLParserListener) EnterOperator_name(ctx *Operator_nameContext) {}

// ExitOperator_name is called when production operator_name is exited.
func (s *BaseSQLParserListener) ExitOperator_name(ctx *Operator_nameContext) {}

// EnterOperator_option is called when production operator_option is entered.
func (s *BaseSQLParserListener) EnterOperator_option(ctx *Operator_optionContext) {}

// ExitOperator_option is called when production operator_option is exited.
func (s *BaseSQLParserListener) ExitOperator_option(ctx *Operator_optionContext) {}

// EnterCreate_aggregate_statement is called when production create_aggregate_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_aggregate_statement(ctx *Create_aggregate_statementContext) {
}

// ExitCreate_aggregate_statement is called when production create_aggregate_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_aggregate_statement(ctx *Create_aggregate_statementContext) {
}

// EnterAggregate_param is called when production aggregate_param is entered.
func (s *BaseSQLParserListener) EnterAggregate_param(ctx *Aggregate_paramContext) {}

// ExitAggregate_param is called when production aggregate_param is exited.
func (s *BaseSQLParserListener) ExitAggregate_param(ctx *Aggregate_paramContext) {}

// EnterSet_statement is called when production set_statement is entered.
func (s *BaseSQLParserListener) EnterSet_statement(ctx *Set_statementContext) {}

// ExitSet_statement is called when production set_statement is exited.
func (s *BaseSQLParserListener) ExitSet_statement(ctx *Set_statementContext) {}

// EnterSet_action is called when production set_action is entered.
func (s *BaseSQLParserListener) EnterSet_action(ctx *Set_actionContext) {}

// ExitSet_action is called when production set_action is exited.
func (s *BaseSQLParserListener) ExitSet_action(ctx *Set_actionContext) {}

// EnterSession_local_option is called when production session_local_option is entered.
func (s *BaseSQLParserListener) EnterSession_local_option(ctx *Session_local_optionContext) {}

// ExitSession_local_option is called when production session_local_option is exited.
func (s *BaseSQLParserListener) ExitSession_local_option(ctx *Session_local_optionContext) {}

// EnterSet_statement_value is called when production set_statement_value is entered.
func (s *BaseSQLParserListener) EnterSet_statement_value(ctx *Set_statement_valueContext) {}

// ExitSet_statement_value is called when production set_statement_value is exited.
func (s *BaseSQLParserListener) ExitSet_statement_value(ctx *Set_statement_valueContext) {}

// EnterCreate_rewrite_statement is called when production create_rewrite_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_rewrite_statement(ctx *Create_rewrite_statementContext) {}

// ExitCreate_rewrite_statement is called when production create_rewrite_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_rewrite_statement(ctx *Create_rewrite_statementContext) {}

// EnterRewrite_command is called when production rewrite_command is entered.
func (s *BaseSQLParserListener) EnterRewrite_command(ctx *Rewrite_commandContext) {}

// ExitRewrite_command is called when production rewrite_command is exited.
func (s *BaseSQLParserListener) ExitRewrite_command(ctx *Rewrite_commandContext) {}

// EnterCreate_trigger_statement is called when production create_trigger_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_trigger_statement(ctx *Create_trigger_statementContext) {}

// ExitCreate_trigger_statement is called when production create_trigger_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_trigger_statement(ctx *Create_trigger_statementContext) {}

// EnterTrigger_referencing is called when production trigger_referencing is entered.
func (s *BaseSQLParserListener) EnterTrigger_referencing(ctx *Trigger_referencingContext) {}

// ExitTrigger_referencing is called when production trigger_referencing is exited.
func (s *BaseSQLParserListener) ExitTrigger_referencing(ctx *Trigger_referencingContext) {}

// EnterWhen_trigger is called when production when_trigger is entered.
func (s *BaseSQLParserListener) EnterWhen_trigger(ctx *When_triggerContext) {}

// ExitWhen_trigger is called when production when_trigger is exited.
func (s *BaseSQLParserListener) ExitWhen_trigger(ctx *When_triggerContext) {}

// EnterRule_common is called when production rule_common is entered.
func (s *BaseSQLParserListener) EnterRule_common(ctx *Rule_commonContext) {}

// ExitRule_common is called when production rule_common is exited.
func (s *BaseSQLParserListener) ExitRule_common(ctx *Rule_commonContext) {}

// EnterRule_member_object is called when production rule_member_object is entered.
func (s *BaseSQLParserListener) EnterRule_member_object(ctx *Rule_member_objectContext) {}

// ExitRule_member_object is called when production rule_member_object is exited.
func (s *BaseSQLParserListener) ExitRule_member_object(ctx *Rule_member_objectContext) {}

// EnterColumns_permissions is called when production columns_permissions is entered.
func (s *BaseSQLParserListener) EnterColumns_permissions(ctx *Columns_permissionsContext) {}

// ExitColumns_permissions is called when production columns_permissions is exited.
func (s *BaseSQLParserListener) ExitColumns_permissions(ctx *Columns_permissionsContext) {}

// EnterTable_column_privileges is called when production table_column_privileges is entered.
func (s *BaseSQLParserListener) EnterTable_column_privileges(ctx *Table_column_privilegesContext) {}

// ExitTable_column_privileges is called when production table_column_privileges is exited.
func (s *BaseSQLParserListener) ExitTable_column_privileges(ctx *Table_column_privilegesContext) {}

// EnterPermissions is called when production permissions is entered.
func (s *BaseSQLParserListener) EnterPermissions(ctx *PermissionsContext) {}

// ExitPermissions is called when production permissions is exited.
func (s *BaseSQLParserListener) ExitPermissions(ctx *PermissionsContext) {}

// EnterPermission is called when production permission is entered.
func (s *BaseSQLParserListener) EnterPermission(ctx *PermissionContext) {}

// ExitPermission is called when production permission is exited.
func (s *BaseSQLParserListener) ExitPermission(ctx *PermissionContext) {}

// EnterOther_rules is called when production other_rules is entered.
func (s *BaseSQLParserListener) EnterOther_rules(ctx *Other_rulesContext) {}

// ExitOther_rules is called when production other_rules is exited.
func (s *BaseSQLParserListener) ExitOther_rules(ctx *Other_rulesContext) {}

// EnterGrant_to_rule is called when production grant_to_rule is entered.
func (s *BaseSQLParserListener) EnterGrant_to_rule(ctx *Grant_to_ruleContext) {}

// ExitGrant_to_rule is called when production grant_to_rule is exited.
func (s *BaseSQLParserListener) ExitGrant_to_rule(ctx *Grant_to_ruleContext) {}

// EnterRevoke_from_cascade_restrict is called when production revoke_from_cascade_restrict is entered.
func (s *BaseSQLParserListener) EnterRevoke_from_cascade_restrict(ctx *Revoke_from_cascade_restrictContext) {
}

// ExitRevoke_from_cascade_restrict is called when production revoke_from_cascade_restrict is exited.
func (s *BaseSQLParserListener) ExitRevoke_from_cascade_restrict(ctx *Revoke_from_cascade_restrictContext) {
}

// EnterRoles_names is called when production roles_names is entered.
func (s *BaseSQLParserListener) EnterRoles_names(ctx *Roles_namesContext) {}

// ExitRoles_names is called when production roles_names is exited.
func (s *BaseSQLParserListener) ExitRoles_names(ctx *Roles_namesContext) {}

// EnterRole_name_with_group is called when production role_name_with_group is entered.
func (s *BaseSQLParserListener) EnterRole_name_with_group(ctx *Role_name_with_groupContext) {}

// ExitRole_name_with_group is called when production role_name_with_group is exited.
func (s *BaseSQLParserListener) ExitRole_name_with_group(ctx *Role_name_with_groupContext) {}

// EnterComment_on_statement is called when production comment_on_statement is entered.
func (s *BaseSQLParserListener) EnterComment_on_statement(ctx *Comment_on_statementContext) {}

// ExitComment_on_statement is called when production comment_on_statement is exited.
func (s *BaseSQLParserListener) ExitComment_on_statement(ctx *Comment_on_statementContext) {}

// EnterSecurity_label is called when production security_label is entered.
func (s *BaseSQLParserListener) EnterSecurity_label(ctx *Security_labelContext) {}

// ExitSecurity_label is called when production security_label is exited.
func (s *BaseSQLParserListener) ExitSecurity_label(ctx *Security_labelContext) {}

// EnterComment_member_object is called when production comment_member_object is entered.
func (s *BaseSQLParserListener) EnterComment_member_object(ctx *Comment_member_objectContext) {}

// ExitComment_member_object is called when production comment_member_object is exited.
func (s *BaseSQLParserListener) ExitComment_member_object(ctx *Comment_member_objectContext) {}

// EnterLabel_member_object is called when production label_member_object is entered.
func (s *BaseSQLParserListener) EnterLabel_member_object(ctx *Label_member_objectContext) {}

// ExitLabel_member_object is called when production label_member_object is exited.
func (s *BaseSQLParserListener) ExitLabel_member_object(ctx *Label_member_objectContext) {}

// EnterCreate_function_statement is called when production create_function_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_function_statement(ctx *Create_function_statementContext) {
}

// ExitCreate_function_statement is called when production create_function_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_function_statement(ctx *Create_function_statementContext) {
}

// EnterCreate_funct_params is called when production create_funct_params is entered.
func (s *BaseSQLParserListener) EnterCreate_funct_params(ctx *Create_funct_paramsContext) {}

// ExitCreate_funct_params is called when production create_funct_params is exited.
func (s *BaseSQLParserListener) ExitCreate_funct_params(ctx *Create_funct_paramsContext) {}

// EnterTransform_for_type is called when production transform_for_type is entered.
func (s *BaseSQLParserListener) EnterTransform_for_type(ctx *Transform_for_typeContext) {}

// ExitTransform_for_type is called when production transform_for_type is exited.
func (s *BaseSQLParserListener) ExitTransform_for_type(ctx *Transform_for_typeContext) {}

// EnterFunction_ret_table is called when production function_ret_table is entered.
func (s *BaseSQLParserListener) EnterFunction_ret_table(ctx *Function_ret_tableContext) {}

// ExitFunction_ret_table is called when production function_ret_table is exited.
func (s *BaseSQLParserListener) ExitFunction_ret_table(ctx *Function_ret_tableContext) {}

// EnterFunction_column_name_type is called when production function_column_name_type is entered.
func (s *BaseSQLParserListener) EnterFunction_column_name_type(ctx *Function_column_name_typeContext) {
}

// ExitFunction_column_name_type is called when production function_column_name_type is exited.
func (s *BaseSQLParserListener) ExitFunction_column_name_type(ctx *Function_column_name_typeContext) {
}

// EnterFunction_parameters is called when production function_parameters is entered.
func (s *BaseSQLParserListener) EnterFunction_parameters(ctx *Function_parametersContext) {}

// ExitFunction_parameters is called when production function_parameters is exited.
func (s *BaseSQLParserListener) ExitFunction_parameters(ctx *Function_parametersContext) {}

// EnterFunction_args is called when production function_args is entered.
func (s *BaseSQLParserListener) EnterFunction_args(ctx *Function_argsContext) {}

// ExitFunction_args is called when production function_args is exited.
func (s *BaseSQLParserListener) ExitFunction_args(ctx *Function_argsContext) {}

// EnterAgg_order is called when production agg_order is entered.
func (s *BaseSQLParserListener) EnterAgg_order(ctx *Agg_orderContext) {}

// ExitAgg_order is called when production agg_order is exited.
func (s *BaseSQLParserListener) ExitAgg_order(ctx *Agg_orderContext) {}

// EnterCharacter_string is called when production character_string is entered.
func (s *BaseSQLParserListener) EnterCharacter_string(ctx *Character_stringContext) {}

// ExitCharacter_string is called when production character_string is exited.
func (s *BaseSQLParserListener) ExitCharacter_string(ctx *Character_stringContext) {}

// EnterFunction_arguments is called when production function_arguments is entered.
func (s *BaseSQLParserListener) EnterFunction_arguments(ctx *Function_argumentsContext) {}

// ExitFunction_arguments is called when production function_arguments is exited.
func (s *BaseSQLParserListener) ExitFunction_arguments(ctx *Function_argumentsContext) {}

// EnterArgmode is called when production argmode is entered.
func (s *BaseSQLParserListener) EnterArgmode(ctx *ArgmodeContext) {}

// ExitArgmode is called when production argmode is exited.
func (s *BaseSQLParserListener) ExitArgmode(ctx *ArgmodeContext) {}

// EnterCreate_sequence_statement is called when production create_sequence_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_sequence_statement(ctx *Create_sequence_statementContext) {
}

// ExitCreate_sequence_statement is called when production create_sequence_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_sequence_statement(ctx *Create_sequence_statementContext) {
}

// EnterSequence_body is called when production sequence_body is entered.
func (s *BaseSQLParserListener) EnterSequence_body(ctx *Sequence_bodyContext) {}

// ExitSequence_body is called when production sequence_body is exited.
func (s *BaseSQLParserListener) ExitSequence_body(ctx *Sequence_bodyContext) {}

// EnterSigned_number_literal is called when production signed_number_literal is entered.
func (s *BaseSQLParserListener) EnterSigned_number_literal(ctx *Signed_number_literalContext) {}

// ExitSigned_number_literal is called when production signed_number_literal is exited.
func (s *BaseSQLParserListener) ExitSigned_number_literal(ctx *Signed_number_literalContext) {}

// EnterSigned_numerical_literal is called when production signed_numerical_literal is entered.
func (s *BaseSQLParserListener) EnterSigned_numerical_literal(ctx *Signed_numerical_literalContext) {}

// ExitSigned_numerical_literal is called when production signed_numerical_literal is exited.
func (s *BaseSQLParserListener) ExitSigned_numerical_literal(ctx *Signed_numerical_literalContext) {}

// EnterSign is called when production sign is entered.
func (s *BaseSQLParserListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BaseSQLParserListener) ExitSign(ctx *SignContext) {}

// EnterCreate_schema_statement is called when production create_schema_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_schema_statement(ctx *Create_schema_statementContext) {}

// ExitCreate_schema_statement is called when production create_schema_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_schema_statement(ctx *Create_schema_statementContext) {}

// EnterCreate_policy_statement is called when production create_policy_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_policy_statement(ctx *Create_policy_statementContext) {}

// ExitCreate_policy_statement is called when production create_policy_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_policy_statement(ctx *Create_policy_statementContext) {}

// EnterAlter_policy_statement is called when production alter_policy_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_policy_statement(ctx *Alter_policy_statementContext) {}

// ExitAlter_policy_statement is called when production alter_policy_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_policy_statement(ctx *Alter_policy_statementContext) {}

// EnterDrop_policy_statement is called when production drop_policy_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_policy_statement(ctx *Drop_policy_statementContext) {}

// ExitDrop_policy_statement is called when production drop_policy_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_policy_statement(ctx *Drop_policy_statementContext) {}

// EnterCreate_subscription_statement is called when production create_subscription_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_subscription_statement(ctx *Create_subscription_statementContext) {
}

// ExitCreate_subscription_statement is called when production create_subscription_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_subscription_statement(ctx *Create_subscription_statementContext) {
}

// EnterAlter_subscription_statement is called when production alter_subscription_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_subscription_statement(ctx *Alter_subscription_statementContext) {
}

// ExitAlter_subscription_statement is called when production alter_subscription_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_subscription_statement(ctx *Alter_subscription_statementContext) {
}

// EnterAlter_subscription_action is called when production alter_subscription_action is entered.
func (s *BaseSQLParserListener) EnterAlter_subscription_action(ctx *Alter_subscription_actionContext) {
}

// ExitAlter_subscription_action is called when production alter_subscription_action is exited.
func (s *BaseSQLParserListener) ExitAlter_subscription_action(ctx *Alter_subscription_actionContext) {
}

// EnterCreate_cast_statement is called when production create_cast_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_cast_statement(ctx *Create_cast_statementContext) {}

// ExitCreate_cast_statement is called when production create_cast_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_cast_statement(ctx *Create_cast_statementContext) {}

// EnterDrop_cast_statement is called when production drop_cast_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_cast_statement(ctx *Drop_cast_statementContext) {}

// ExitDrop_cast_statement is called when production drop_cast_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_cast_statement(ctx *Drop_cast_statementContext) {}

// EnterCreate_operator_family_statement is called when production create_operator_family_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_operator_family_statement(ctx *Create_operator_family_statementContext) {
}

// ExitCreate_operator_family_statement is called when production create_operator_family_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_operator_family_statement(ctx *Create_operator_family_statementContext) {
}

// EnterAlter_operator_family_statement is called when production alter_operator_family_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_operator_family_statement(ctx *Alter_operator_family_statementContext) {
}

// ExitAlter_operator_family_statement is called when production alter_operator_family_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_operator_family_statement(ctx *Alter_operator_family_statementContext) {
}

// EnterOperator_family_action is called when production operator_family_action is entered.
func (s *BaseSQLParserListener) EnterOperator_family_action(ctx *Operator_family_actionContext) {}

// ExitOperator_family_action is called when production operator_family_action is exited.
func (s *BaseSQLParserListener) ExitOperator_family_action(ctx *Operator_family_actionContext) {}

// EnterAdd_operator_to_family is called when production add_operator_to_family is entered.
func (s *BaseSQLParserListener) EnterAdd_operator_to_family(ctx *Add_operator_to_familyContext) {}

// ExitAdd_operator_to_family is called when production add_operator_to_family is exited.
func (s *BaseSQLParserListener) ExitAdd_operator_to_family(ctx *Add_operator_to_familyContext) {}

// EnterDrop_operator_from_family is called when production drop_operator_from_family is entered.
func (s *BaseSQLParserListener) EnterDrop_operator_from_family(ctx *Drop_operator_from_familyContext) {
}

// ExitDrop_operator_from_family is called when production drop_operator_from_family is exited.
func (s *BaseSQLParserListener) ExitDrop_operator_from_family(ctx *Drop_operator_from_familyContext) {
}

// EnterDrop_operator_family_statement is called when production drop_operator_family_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_operator_family_statement(ctx *Drop_operator_family_statementContext) {
}

// ExitDrop_operator_family_statement is called when production drop_operator_family_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_operator_family_statement(ctx *Drop_operator_family_statementContext) {
}

// EnterCreate_operator_class_statement is called when production create_operator_class_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_operator_class_statement(ctx *Create_operator_class_statementContext) {
}

// ExitCreate_operator_class_statement is called when production create_operator_class_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_operator_class_statement(ctx *Create_operator_class_statementContext) {
}

// EnterCreate_operator_class_option is called when production create_operator_class_option is entered.
func (s *BaseSQLParserListener) EnterCreate_operator_class_option(ctx *Create_operator_class_optionContext) {
}

// ExitCreate_operator_class_option is called when production create_operator_class_option is exited.
func (s *BaseSQLParserListener) ExitCreate_operator_class_option(ctx *Create_operator_class_optionContext) {
}

// EnterAlter_operator_class_statement is called when production alter_operator_class_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_operator_class_statement(ctx *Alter_operator_class_statementContext) {
}

// ExitAlter_operator_class_statement is called when production alter_operator_class_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_operator_class_statement(ctx *Alter_operator_class_statementContext) {
}

// EnterDrop_operator_class_statement is called when production drop_operator_class_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_operator_class_statement(ctx *Drop_operator_class_statementContext) {
}

// ExitDrop_operator_class_statement is called when production drop_operator_class_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_operator_class_statement(ctx *Drop_operator_class_statementContext) {
}

// EnterCreate_conversion_statement is called when production create_conversion_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_conversion_statement(ctx *Create_conversion_statementContext) {
}

// ExitCreate_conversion_statement is called when production create_conversion_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_conversion_statement(ctx *Create_conversion_statementContext) {
}

// EnterAlter_conversion_statement is called when production alter_conversion_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_conversion_statement(ctx *Alter_conversion_statementContext) {
}

// ExitAlter_conversion_statement is called when production alter_conversion_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_conversion_statement(ctx *Alter_conversion_statementContext) {
}

// EnterCreate_publication_statement is called when production create_publication_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_publication_statement(ctx *Create_publication_statementContext) {
}

// ExitCreate_publication_statement is called when production create_publication_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_publication_statement(ctx *Create_publication_statementContext) {
}

// EnterAlter_publication_statement is called when production alter_publication_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_publication_statement(ctx *Alter_publication_statementContext) {
}

// ExitAlter_publication_statement is called when production alter_publication_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_publication_statement(ctx *Alter_publication_statementContext) {
}

// EnterAlter_publication_action is called when production alter_publication_action is entered.
func (s *BaseSQLParserListener) EnterAlter_publication_action(ctx *Alter_publication_actionContext) {}

// ExitAlter_publication_action is called when production alter_publication_action is exited.
func (s *BaseSQLParserListener) ExitAlter_publication_action(ctx *Alter_publication_actionContext) {}

// EnterOnly_table_multiply is called when production only_table_multiply is entered.
func (s *BaseSQLParserListener) EnterOnly_table_multiply(ctx *Only_table_multiplyContext) {}

// ExitOnly_table_multiply is called when production only_table_multiply is exited.
func (s *BaseSQLParserListener) ExitOnly_table_multiply(ctx *Only_table_multiplyContext) {}

// EnterAlter_trigger_statement is called when production alter_trigger_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_trigger_statement(ctx *Alter_trigger_statementContext) {}

// ExitAlter_trigger_statement is called when production alter_trigger_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_trigger_statement(ctx *Alter_trigger_statementContext) {}

// EnterAlter_rule_statement is called when production alter_rule_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_rule_statement(ctx *Alter_rule_statementContext) {}

// ExitAlter_rule_statement is called when production alter_rule_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_rule_statement(ctx *Alter_rule_statementContext) {}

// EnterCopy_statement is called when production copy_statement is entered.
func (s *BaseSQLParserListener) EnterCopy_statement(ctx *Copy_statementContext) {}

// ExitCopy_statement is called when production copy_statement is exited.
func (s *BaseSQLParserListener) ExitCopy_statement(ctx *Copy_statementContext) {}

// EnterCopy_from_statement is called when production copy_from_statement is entered.
func (s *BaseSQLParserListener) EnterCopy_from_statement(ctx *Copy_from_statementContext) {}

// ExitCopy_from_statement is called when production copy_from_statement is exited.
func (s *BaseSQLParserListener) ExitCopy_from_statement(ctx *Copy_from_statementContext) {}

// EnterCopy_to_statement is called when production copy_to_statement is entered.
func (s *BaseSQLParserListener) EnterCopy_to_statement(ctx *Copy_to_statementContext) {}

// ExitCopy_to_statement is called when production copy_to_statement is exited.
func (s *BaseSQLParserListener) ExitCopy_to_statement(ctx *Copy_to_statementContext) {}

// EnterCopy_option_list is called when production copy_option_list is entered.
func (s *BaseSQLParserListener) EnterCopy_option_list(ctx *Copy_option_listContext) {}

// ExitCopy_option_list is called when production copy_option_list is exited.
func (s *BaseSQLParserListener) ExitCopy_option_list(ctx *Copy_option_listContext) {}

// EnterCopy_option is called when production copy_option is entered.
func (s *BaseSQLParserListener) EnterCopy_option(ctx *Copy_optionContext) {}

// ExitCopy_option is called when production copy_option is exited.
func (s *BaseSQLParserListener) ExitCopy_option(ctx *Copy_optionContext) {}

// EnterCreate_view_statement is called when production create_view_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_view_statement(ctx *Create_view_statementContext) {}

// ExitCreate_view_statement is called when production create_view_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_view_statement(ctx *Create_view_statementContext) {}

// EnterIf_exists is called when production if_exists is entered.
func (s *BaseSQLParserListener) EnterIf_exists(ctx *If_existsContext) {}

// ExitIf_exists is called when production if_exists is exited.
func (s *BaseSQLParserListener) ExitIf_exists(ctx *If_existsContext) {}

// EnterIf_not_exists is called when production if_not_exists is entered.
func (s *BaseSQLParserListener) EnterIf_not_exists(ctx *If_not_existsContext) {}

// ExitIf_not_exists is called when production if_not_exists is exited.
func (s *BaseSQLParserListener) ExitIf_not_exists(ctx *If_not_existsContext) {}

// EnterView_columns is called when production view_columns is entered.
func (s *BaseSQLParserListener) EnterView_columns(ctx *View_columnsContext) {}

// ExitView_columns is called when production view_columns is exited.
func (s *BaseSQLParserListener) ExitView_columns(ctx *View_columnsContext) {}

// EnterWith_check_option is called when production with_check_option is entered.
func (s *BaseSQLParserListener) EnterWith_check_option(ctx *With_check_optionContext) {}

// ExitWith_check_option is called when production with_check_option is exited.
func (s *BaseSQLParserListener) ExitWith_check_option(ctx *With_check_optionContext) {}

// EnterCreate_database_statement is called when production create_database_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_database_statement(ctx *Create_database_statementContext) {
}

// ExitCreate_database_statement is called when production create_database_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_database_statement(ctx *Create_database_statementContext) {
}

// EnterCreate_database_option is called when production create_database_option is entered.
func (s *BaseSQLParserListener) EnterCreate_database_option(ctx *Create_database_optionContext) {}

// ExitCreate_database_option is called when production create_database_option is exited.
func (s *BaseSQLParserListener) ExitCreate_database_option(ctx *Create_database_optionContext) {}

// EnterAlter_database_statement is called when production alter_database_statement is entered.
func (s *BaseSQLParserListener) EnterAlter_database_statement(ctx *Alter_database_statementContext) {}

// ExitAlter_database_statement is called when production alter_database_statement is exited.
func (s *BaseSQLParserListener) ExitAlter_database_statement(ctx *Alter_database_statementContext) {}

// EnterAlter_database_action is called when production alter_database_action is entered.
func (s *BaseSQLParserListener) EnterAlter_database_action(ctx *Alter_database_actionContext) {}

// ExitAlter_database_action is called when production alter_database_action is exited.
func (s *BaseSQLParserListener) ExitAlter_database_action(ctx *Alter_database_actionContext) {}

// EnterAlter_database_option is called when production alter_database_option is entered.
func (s *BaseSQLParserListener) EnterAlter_database_option(ctx *Alter_database_optionContext) {}

// ExitAlter_database_option is called when production alter_database_option is exited.
func (s *BaseSQLParserListener) ExitAlter_database_option(ctx *Alter_database_optionContext) {}

// EnterCreate_table_statement is called when production create_table_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_table_statement(ctx *Create_table_statementContext) {}

// ExitCreate_table_statement is called when production create_table_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_table_statement(ctx *Create_table_statementContext) {}

// EnterCreate_table_as_statement is called when production create_table_as_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_table_as_statement(ctx *Create_table_as_statementContext) {
}

// ExitCreate_table_as_statement is called when production create_table_as_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_table_as_statement(ctx *Create_table_as_statementContext) {
}

// EnterCreate_foreign_table_statement is called when production create_foreign_table_statement is entered.
func (s *BaseSQLParserListener) EnterCreate_foreign_table_statement(ctx *Create_foreign_table_statementContext) {
}

// ExitCreate_foreign_table_statement is called when production create_foreign_table_statement is exited.
func (s *BaseSQLParserListener) ExitCreate_foreign_table_statement(ctx *Create_foreign_table_statementContext) {
}

// EnterDefine_table is called when production define_table is entered.
func (s *BaseSQLParserListener) EnterDefine_table(ctx *Define_tableContext) {}

// ExitDefine_table is called when production define_table is exited.
func (s *BaseSQLParserListener) ExitDefine_table(ctx *Define_tableContext) {}

// EnterDefine_partition is called when production define_partition is entered.
func (s *BaseSQLParserListener) EnterDefine_partition(ctx *Define_partitionContext) {}

// ExitDefine_partition is called when production define_partition is exited.
func (s *BaseSQLParserListener) ExitDefine_partition(ctx *Define_partitionContext) {}

// EnterFor_values_bound is called when production for_values_bound is entered.
func (s *BaseSQLParserListener) EnterFor_values_bound(ctx *For_values_boundContext) {}

// ExitFor_values_bound is called when production for_values_bound is exited.
func (s *BaseSQLParserListener) ExitFor_values_bound(ctx *For_values_boundContext) {}

// EnterPartition_bound_spec is called when production partition_bound_spec is entered.
func (s *BaseSQLParserListener) EnterPartition_bound_spec(ctx *Partition_bound_specContext) {}

// ExitPartition_bound_spec is called when production partition_bound_spec is exited.
func (s *BaseSQLParserListener) ExitPartition_bound_spec(ctx *Partition_bound_specContext) {}

// EnterPartition_bound_part is called when production partition_bound_part is entered.
func (s *BaseSQLParserListener) EnterPartition_bound_part(ctx *Partition_bound_partContext) {}

// ExitPartition_bound_part is called when production partition_bound_part is exited.
func (s *BaseSQLParserListener) ExitPartition_bound_part(ctx *Partition_bound_partContext) {}

// EnterDefine_columns is called when production define_columns is entered.
func (s *BaseSQLParserListener) EnterDefine_columns(ctx *Define_columnsContext) {}

// ExitDefine_columns is called when production define_columns is exited.
func (s *BaseSQLParserListener) ExitDefine_columns(ctx *Define_columnsContext) {}

// EnterDefine_type is called when production define_type is entered.
func (s *BaseSQLParserListener) EnterDefine_type(ctx *Define_typeContext) {}

// ExitDefine_type is called when production define_type is exited.
func (s *BaseSQLParserListener) ExitDefine_type(ctx *Define_typeContext) {}

// EnterPartition_by is called when production partition_by is entered.
func (s *BaseSQLParserListener) EnterPartition_by(ctx *Partition_byContext) {}

// ExitPartition_by is called when production partition_by is exited.
func (s *BaseSQLParserListener) ExitPartition_by(ctx *Partition_byContext) {}

// EnterPartition_method is called when production partition_method is entered.
func (s *BaseSQLParserListener) EnterPartition_method(ctx *Partition_methodContext) {}

// ExitPartition_method is called when production partition_method is exited.
func (s *BaseSQLParserListener) ExitPartition_method(ctx *Partition_methodContext) {}

// EnterPartition_column is called when production partition_column is entered.
func (s *BaseSQLParserListener) EnterPartition_column(ctx *Partition_columnContext) {}

// ExitPartition_column is called when production partition_column is exited.
func (s *BaseSQLParserListener) ExitPartition_column(ctx *Partition_columnContext) {}

// EnterDefine_server is called when production define_server is entered.
func (s *BaseSQLParserListener) EnterDefine_server(ctx *Define_serverContext) {}

// ExitDefine_server is called when production define_server is exited.
func (s *BaseSQLParserListener) ExitDefine_server(ctx *Define_serverContext) {}

// EnterDefine_foreign_options is called when production define_foreign_options is entered.
func (s *BaseSQLParserListener) EnterDefine_foreign_options(ctx *Define_foreign_optionsContext) {}

// ExitDefine_foreign_options is called when production define_foreign_options is exited.
func (s *BaseSQLParserListener) ExitDefine_foreign_options(ctx *Define_foreign_optionsContext) {}

// EnterForeign_option is called when production foreign_option is entered.
func (s *BaseSQLParserListener) EnterForeign_option(ctx *Foreign_optionContext) {}

// ExitForeign_option is called when production foreign_option is exited.
func (s *BaseSQLParserListener) ExitForeign_option(ctx *Foreign_optionContext) {}

// EnterForeign_option_name is called when production foreign_option_name is entered.
func (s *BaseSQLParserListener) EnterForeign_option_name(ctx *Foreign_option_nameContext) {}

// ExitForeign_option_name is called when production foreign_option_name is exited.
func (s *BaseSQLParserListener) ExitForeign_option_name(ctx *Foreign_option_nameContext) {}

// EnterList_of_type_column_def is called when production list_of_type_column_def is entered.
func (s *BaseSQLParserListener) EnterList_of_type_column_def(ctx *List_of_type_column_defContext) {}

// ExitList_of_type_column_def is called when production list_of_type_column_def is exited.
func (s *BaseSQLParserListener) ExitList_of_type_column_def(ctx *List_of_type_column_defContext) {}

// EnterTable_column_def is called when production table_column_def is entered.
func (s *BaseSQLParserListener) EnterTable_column_def(ctx *Table_column_defContext) {}

// ExitTable_column_def is called when production table_column_def is exited.
func (s *BaseSQLParserListener) ExitTable_column_def(ctx *Table_column_defContext) {}

// EnterTable_of_type_column_def is called when production table_of_type_column_def is entered.
func (s *BaseSQLParserListener) EnterTable_of_type_column_def(ctx *Table_of_type_column_defContext) {}

// ExitTable_of_type_column_def is called when production table_of_type_column_def is exited.
func (s *BaseSQLParserListener) ExitTable_of_type_column_def(ctx *Table_of_type_column_defContext) {}

// EnterTable_column_definition is called when production table_column_definition is entered.
func (s *BaseSQLParserListener) EnterTable_column_definition(ctx *Table_column_definitionContext) {}

// ExitTable_column_definition is called when production table_column_definition is exited.
func (s *BaseSQLParserListener) ExitTable_column_definition(ctx *Table_column_definitionContext) {}

// EnterLike_option is called when production like_option is entered.
func (s *BaseSQLParserListener) EnterLike_option(ctx *Like_optionContext) {}

// ExitLike_option is called when production like_option is exited.
func (s *BaseSQLParserListener) ExitLike_option(ctx *Like_optionContext) {}

// EnterConstraint_common is called when production constraint_common is entered.
func (s *BaseSQLParserListener) EnterConstraint_common(ctx *Constraint_commonContext) {}

// ExitConstraint_common is called when production constraint_common is exited.
func (s *BaseSQLParserListener) ExitConstraint_common(ctx *Constraint_commonContext) {}

// EnterConstr_body is called when production constr_body is entered.
func (s *BaseSQLParserListener) EnterConstr_body(ctx *Constr_bodyContext) {}

// ExitConstr_body is called when production constr_body is exited.
func (s *BaseSQLParserListener) ExitConstr_body(ctx *Constr_bodyContext) {}

// EnterAll_op is called when production all_op is entered.
func (s *BaseSQLParserListener) EnterAll_op(ctx *All_opContext) {}

// ExitAll_op is called when production all_op is exited.
func (s *BaseSQLParserListener) ExitAll_op(ctx *All_opContext) {}

// EnterAll_simple_op is called when production all_simple_op is entered.
func (s *BaseSQLParserListener) EnterAll_simple_op(ctx *All_simple_opContext) {}

// ExitAll_simple_op is called when production all_simple_op is exited.
func (s *BaseSQLParserListener) ExitAll_simple_op(ctx *All_simple_opContext) {}

// EnterOp_chars is called when production op_chars is entered.
func (s *BaseSQLParserListener) EnterOp_chars(ctx *Op_charsContext) {}

// ExitOp_chars is called when production op_chars is exited.
func (s *BaseSQLParserListener) ExitOp_chars(ctx *Op_charsContext) {}

// EnterIndex_parameters is called when production index_parameters is entered.
func (s *BaseSQLParserListener) EnterIndex_parameters(ctx *Index_parametersContext) {}

// ExitIndex_parameters is called when production index_parameters is exited.
func (s *BaseSQLParserListener) ExitIndex_parameters(ctx *Index_parametersContext) {}

// EnterNames_in_parens is called when production names_in_parens is entered.
func (s *BaseSQLParserListener) EnterNames_in_parens(ctx *Names_in_parensContext) {}

// ExitNames_in_parens is called when production names_in_parens is exited.
func (s *BaseSQLParserListener) ExitNames_in_parens(ctx *Names_in_parensContext) {}

// EnterNames_references is called when production names_references is entered.
func (s *BaseSQLParserListener) EnterNames_references(ctx *Names_referencesContext) {}

// ExitNames_references is called when production names_references is exited.
func (s *BaseSQLParserListener) ExitNames_references(ctx *Names_referencesContext) {}

// EnterStorage_parameter is called when production storage_parameter is entered.
func (s *BaseSQLParserListener) EnterStorage_parameter(ctx *Storage_parameterContext) {}

// ExitStorage_parameter is called when production storage_parameter is exited.
func (s *BaseSQLParserListener) ExitStorage_parameter(ctx *Storage_parameterContext) {}

// EnterStorage_parameter_option is called when production storage_parameter_option is entered.
func (s *BaseSQLParserListener) EnterStorage_parameter_option(ctx *Storage_parameter_optionContext) {}

// ExitStorage_parameter_option is called when production storage_parameter_option is exited.
func (s *BaseSQLParserListener) ExitStorage_parameter_option(ctx *Storage_parameter_optionContext) {}

// EnterStorage_parameter_name is called when production storage_parameter_name is entered.
func (s *BaseSQLParserListener) EnterStorage_parameter_name(ctx *Storage_parameter_nameContext) {}

// ExitStorage_parameter_name is called when production storage_parameter_name is exited.
func (s *BaseSQLParserListener) ExitStorage_parameter_name(ctx *Storage_parameter_nameContext) {}

// EnterWith_storage_parameter is called when production with_storage_parameter is entered.
func (s *BaseSQLParserListener) EnterWith_storage_parameter(ctx *With_storage_parameterContext) {}

// ExitWith_storage_parameter is called when production with_storage_parameter is exited.
func (s *BaseSQLParserListener) ExitWith_storage_parameter(ctx *With_storage_parameterContext) {}

// EnterStorage_parameter_oid is called when production storage_parameter_oid is entered.
func (s *BaseSQLParserListener) EnterStorage_parameter_oid(ctx *Storage_parameter_oidContext) {}

// ExitStorage_parameter_oid is called when production storage_parameter_oid is exited.
func (s *BaseSQLParserListener) ExitStorage_parameter_oid(ctx *Storage_parameter_oidContext) {}

// EnterOn_commit is called when production on_commit is entered.
func (s *BaseSQLParserListener) EnterOn_commit(ctx *On_commitContext) {}

// ExitOn_commit is called when production on_commit is exited.
func (s *BaseSQLParserListener) ExitOn_commit(ctx *On_commitContext) {}

// EnterTable_space is called when production table_space is entered.
func (s *BaseSQLParserListener) EnterTable_space(ctx *Table_spaceContext) {}

// ExitTable_space is called when production table_space is exited.
func (s *BaseSQLParserListener) ExitTable_space(ctx *Table_spaceContext) {}

// EnterSet_tablespace is called when production set_tablespace is entered.
func (s *BaseSQLParserListener) EnterSet_tablespace(ctx *Set_tablespaceContext) {}

// ExitSet_tablespace is called when production set_tablespace is exited.
func (s *BaseSQLParserListener) ExitSet_tablespace(ctx *Set_tablespaceContext) {}

// EnterX_action is called when production x_action is entered.
func (s *BaseSQLParserListener) EnterX_action(ctx *X_actionContext) {}

// ExitX_action is called when production x_action is exited.
func (s *BaseSQLParserListener) ExitX_action(ctx *X_actionContext) {}

// EnterOwner_to is called when production owner_to is entered.
func (s *BaseSQLParserListener) EnterOwner_to(ctx *Owner_toContext) {}

// ExitOwner_to is called when production owner_to is exited.
func (s *BaseSQLParserListener) ExitOwner_to(ctx *Owner_toContext) {}

// EnterRename_to is called when production rename_to is entered.
func (s *BaseSQLParserListener) EnterRename_to(ctx *Rename_toContext) {}

// ExitRename_to is called when production rename_to is exited.
func (s *BaseSQLParserListener) ExitRename_to(ctx *Rename_toContext) {}

// EnterSet_schema is called when production set_schema is entered.
func (s *BaseSQLParserListener) EnterSet_schema(ctx *Set_schemaContext) {}

// ExitSet_schema is called when production set_schema is exited.
func (s *BaseSQLParserListener) ExitSet_schema(ctx *Set_schemaContext) {}

// EnterTable_column_privilege is called when production table_column_privilege is entered.
func (s *BaseSQLParserListener) EnterTable_column_privilege(ctx *Table_column_privilegeContext) {}

// ExitTable_column_privilege is called when production table_column_privilege is exited.
func (s *BaseSQLParserListener) ExitTable_column_privilege(ctx *Table_column_privilegeContext) {}

// EnterUsage_select_update is called when production usage_select_update is entered.
func (s *BaseSQLParserListener) EnterUsage_select_update(ctx *Usage_select_updateContext) {}

// ExitUsage_select_update is called when production usage_select_update is exited.
func (s *BaseSQLParserListener) ExitUsage_select_update(ctx *Usage_select_updateContext) {}

// EnterPartition_by_columns is called when production partition_by_columns is entered.
func (s *BaseSQLParserListener) EnterPartition_by_columns(ctx *Partition_by_columnsContext) {}

// ExitPartition_by_columns is called when production partition_by_columns is exited.
func (s *BaseSQLParserListener) ExitPartition_by_columns(ctx *Partition_by_columnsContext) {}

// EnterCascade_restrict is called when production cascade_restrict is entered.
func (s *BaseSQLParserListener) EnterCascade_restrict(ctx *Cascade_restrictContext) {}

// ExitCascade_restrict is called when production cascade_restrict is exited.
func (s *BaseSQLParserListener) ExitCascade_restrict(ctx *Cascade_restrictContext) {}

// EnterCollate_identifier is called when production collate_identifier is entered.
func (s *BaseSQLParserListener) EnterCollate_identifier(ctx *Collate_identifierContext) {}

// ExitCollate_identifier is called when production collate_identifier is exited.
func (s *BaseSQLParserListener) ExitCollate_identifier(ctx *Collate_identifierContext) {}

// EnterIndirection_var is called when production indirection_var is entered.
func (s *BaseSQLParserListener) EnterIndirection_var(ctx *Indirection_varContext) {}

// ExitIndirection_var is called when production indirection_var is exited.
func (s *BaseSQLParserListener) ExitIndirection_var(ctx *Indirection_varContext) {}

// EnterDollar_number is called when production dollar_number is entered.
func (s *BaseSQLParserListener) EnterDollar_number(ctx *Dollar_numberContext) {}

// ExitDollar_number is called when production dollar_number is exited.
func (s *BaseSQLParserListener) ExitDollar_number(ctx *Dollar_numberContext) {}

// EnterIndirection_list is called when production indirection_list is entered.
func (s *BaseSQLParserListener) EnterIndirection_list(ctx *Indirection_listContext) {}

// ExitIndirection_list is called when production indirection_list is exited.
func (s *BaseSQLParserListener) ExitIndirection_list(ctx *Indirection_listContext) {}

// EnterIndirection is called when production indirection is entered.
func (s *BaseSQLParserListener) EnterIndirection(ctx *IndirectionContext) {}

// ExitIndirection is called when production indirection is exited.
func (s *BaseSQLParserListener) ExitIndirection(ctx *IndirectionContext) {}

// EnterDrop_database_statement is called when production drop_database_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_database_statement(ctx *Drop_database_statementContext) {}

// ExitDrop_database_statement is called when production drop_database_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_database_statement(ctx *Drop_database_statementContext) {}

// EnterDrop_function_statement is called when production drop_function_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_function_statement(ctx *Drop_function_statementContext) {}

// ExitDrop_function_statement is called when production drop_function_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_function_statement(ctx *Drop_function_statementContext) {}

// EnterDrop_trigger_statement is called when production drop_trigger_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_trigger_statement(ctx *Drop_trigger_statementContext) {}

// ExitDrop_trigger_statement is called when production drop_trigger_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_trigger_statement(ctx *Drop_trigger_statementContext) {}

// EnterDrop_rule_statement is called when production drop_rule_statement is entered.
func (s *BaseSQLParserListener) EnterDrop_rule_statement(ctx *Drop_rule_statementContext) {}

// ExitDrop_rule_statement is called when production drop_rule_statement is exited.
func (s *BaseSQLParserListener) ExitDrop_rule_statement(ctx *Drop_rule_statementContext) {}

// EnterDrop_statements is called when production drop_statements is entered.
func (s *BaseSQLParserListener) EnterDrop_statements(ctx *Drop_statementsContext) {}

// ExitDrop_statements is called when production drop_statements is exited.
func (s *BaseSQLParserListener) ExitDrop_statements(ctx *Drop_statementsContext) {}

// EnterIf_exist_names_restrict_cascade is called when production if_exist_names_restrict_cascade is entered.
func (s *BaseSQLParserListener) EnterIf_exist_names_restrict_cascade(ctx *If_exist_names_restrict_cascadeContext) {
}

// ExitIf_exist_names_restrict_cascade is called when production if_exist_names_restrict_cascade is exited.
func (s *BaseSQLParserListener) ExitIf_exist_names_restrict_cascade(ctx *If_exist_names_restrict_cascadeContext) {
}

// EnterId_token is called when production id_token is entered.
func (s *BaseSQLParserListener) EnterId_token(ctx *Id_tokenContext) {}

// ExitId_token is called when production id_token is exited.
func (s *BaseSQLParserListener) ExitId_token(ctx *Id_tokenContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSQLParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSQLParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifier_nontype is called when production identifier_nontype is entered.
func (s *BaseSQLParserListener) EnterIdentifier_nontype(ctx *Identifier_nontypeContext) {}

// ExitIdentifier_nontype is called when production identifier_nontype is exited.
func (s *BaseSQLParserListener) ExitIdentifier_nontype(ctx *Identifier_nontypeContext) {}

// EnterCol_label is called when production col_label is entered.
func (s *BaseSQLParserListener) EnterCol_label(ctx *Col_labelContext) {}

// ExitCol_label is called when production col_label is exited.
func (s *BaseSQLParserListener) ExitCol_label(ctx *Col_labelContext) {}

// EnterTokens_nonreserved is called when production tokens_nonreserved is entered.
func (s *BaseSQLParserListener) EnterTokens_nonreserved(ctx *Tokens_nonreservedContext) {}

// ExitTokens_nonreserved is called when production tokens_nonreserved is exited.
func (s *BaseSQLParserListener) ExitTokens_nonreserved(ctx *Tokens_nonreservedContext) {}

// EnterTokens_nonreserved_except_function_type is called when production tokens_nonreserved_except_function_type is entered.
func (s *BaseSQLParserListener) EnterTokens_nonreserved_except_function_type(ctx *Tokens_nonreserved_except_function_typeContext) {
}

// ExitTokens_nonreserved_except_function_type is called when production tokens_nonreserved_except_function_type is exited.
func (s *BaseSQLParserListener) ExitTokens_nonreserved_except_function_type(ctx *Tokens_nonreserved_except_function_typeContext) {
}

// EnterTokens_reserved_except_function_type is called when production tokens_reserved_except_function_type is entered.
func (s *BaseSQLParserListener) EnterTokens_reserved_except_function_type(ctx *Tokens_reserved_except_function_typeContext) {
}

// ExitTokens_reserved_except_function_type is called when production tokens_reserved_except_function_type is exited.
func (s *BaseSQLParserListener) ExitTokens_reserved_except_function_type(ctx *Tokens_reserved_except_function_typeContext) {
}

// EnterTokens_reserved is called when production tokens_reserved is entered.
func (s *BaseSQLParserListener) EnterTokens_reserved(ctx *Tokens_reservedContext) {}

// ExitTokens_reserved is called when production tokens_reserved is exited.
func (s *BaseSQLParserListener) ExitTokens_reserved(ctx *Tokens_reservedContext) {}

// EnterTokens_nonkeyword is called when production tokens_nonkeyword is entered.
func (s *BaseSQLParserListener) EnterTokens_nonkeyword(ctx *Tokens_nonkeywordContext) {}

// ExitTokens_nonkeyword is called when production tokens_nonkeyword is exited.
func (s *BaseSQLParserListener) ExitTokens_nonkeyword(ctx *Tokens_nonkeywordContext) {}

// EnterSchema_qualified_name_nontype is called when production schema_qualified_name_nontype is entered.
func (s *BaseSQLParserListener) EnterSchema_qualified_name_nontype(ctx *Schema_qualified_name_nontypeContext) {
}

// ExitSchema_qualified_name_nontype is called when production schema_qualified_name_nontype is exited.
func (s *BaseSQLParserListener) ExitSchema_qualified_name_nontype(ctx *Schema_qualified_name_nontypeContext) {
}

// EnterType_list is called when production type_list is entered.
func (s *BaseSQLParserListener) EnterType_list(ctx *Type_listContext) {}

// ExitType_list is called when production type_list is exited.
func (s *BaseSQLParserListener) ExitType_list(ctx *Type_listContext) {}

// EnterData_type is called when production data_type is entered.
func (s *BaseSQLParserListener) EnterData_type(ctx *Data_typeContext) {}

// ExitData_type is called when production data_type is exited.
func (s *BaseSQLParserListener) ExitData_type(ctx *Data_typeContext) {}

// EnterArray_type is called when production array_type is entered.
func (s *BaseSQLParserListener) EnterArray_type(ctx *Array_typeContext) {}

// ExitArray_type is called when production array_type is exited.
func (s *BaseSQLParserListener) ExitArray_type(ctx *Array_typeContext) {}

// EnterPredefined_type is called when production predefined_type is entered.
func (s *BaseSQLParserListener) EnterPredefined_type(ctx *Predefined_typeContext) {}

// ExitPredefined_type is called when production predefined_type is exited.
func (s *BaseSQLParserListener) ExitPredefined_type(ctx *Predefined_typeContext) {}

// EnterInterval_field is called when production interval_field is entered.
func (s *BaseSQLParserListener) EnterInterval_field(ctx *Interval_fieldContext) {}

// ExitInterval_field is called when production interval_field is exited.
func (s *BaseSQLParserListener) ExitInterval_field(ctx *Interval_fieldContext) {}

// EnterType_length is called when production type_length is entered.
func (s *BaseSQLParserListener) EnterType_length(ctx *Type_lengthContext) {}

// ExitType_length is called when production type_length is exited.
func (s *BaseSQLParserListener) ExitType_length(ctx *Type_lengthContext) {}

// EnterPrecision_param is called when production precision_param is entered.
func (s *BaseSQLParserListener) EnterPrecision_param(ctx *Precision_paramContext) {}

// ExitPrecision_param is called when production precision_param is exited.
func (s *BaseSQLParserListener) ExitPrecision_param(ctx *Precision_paramContext) {}

// EnterVex is called when production vex is entered.
func (s *BaseSQLParserListener) EnterVex(ctx *VexContext) {}

// ExitVex is called when production vex is exited.
func (s *BaseSQLParserListener) ExitVex(ctx *VexContext) {}

// EnterVex_b is called when production vex_b is entered.
func (s *BaseSQLParserListener) EnterVex_b(ctx *Vex_bContext) {}

// ExitVex_b is called when production vex_b is exited.
func (s *BaseSQLParserListener) ExitVex_b(ctx *Vex_bContext) {}

// EnterOp is called when production op is entered.
func (s *BaseSQLParserListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseSQLParserListener) ExitOp(ctx *OpContext) {}

// EnterAll_op_ref is called when production all_op_ref is entered.
func (s *BaseSQLParserListener) EnterAll_op_ref(ctx *All_op_refContext) {}

// ExitAll_op_ref is called when production all_op_ref is exited.
func (s *BaseSQLParserListener) ExitAll_op_ref(ctx *All_op_refContext) {}

// EnterDatetime_overlaps is called when production datetime_overlaps is entered.
func (s *BaseSQLParserListener) EnterDatetime_overlaps(ctx *Datetime_overlapsContext) {}

// ExitDatetime_overlaps is called when production datetime_overlaps is exited.
func (s *BaseSQLParserListener) ExitDatetime_overlaps(ctx *Datetime_overlapsContext) {}

// EnterValue_expression_primary is called when production value_expression_primary is entered.
func (s *BaseSQLParserListener) EnterValue_expression_primary(ctx *Value_expression_primaryContext) {}

// ExitValue_expression_primary is called when production value_expression_primary is exited.
func (s *BaseSQLParserListener) ExitValue_expression_primary(ctx *Value_expression_primaryContext) {}

// EnterUnsigned_value_specification is called when production unsigned_value_specification is entered.
func (s *BaseSQLParserListener) EnterUnsigned_value_specification(ctx *Unsigned_value_specificationContext) {
}

// ExitUnsigned_value_specification is called when production unsigned_value_specification is exited.
func (s *BaseSQLParserListener) ExitUnsigned_value_specification(ctx *Unsigned_value_specificationContext) {
}

// EnterUnsigned_numeric_literal is called when production unsigned_numeric_literal is entered.
func (s *BaseSQLParserListener) EnterUnsigned_numeric_literal(ctx *Unsigned_numeric_literalContext) {}

// ExitUnsigned_numeric_literal is called when production unsigned_numeric_literal is exited.
func (s *BaseSQLParserListener) ExitUnsigned_numeric_literal(ctx *Unsigned_numeric_literalContext) {}

// EnterTruth_value is called when production truth_value is entered.
func (s *BaseSQLParserListener) EnterTruth_value(ctx *Truth_valueContext) {}

// ExitTruth_value is called when production truth_value is exited.
func (s *BaseSQLParserListener) ExitTruth_value(ctx *Truth_valueContext) {}

// EnterCase_expression is called when production case_expression is entered.
func (s *BaseSQLParserListener) EnterCase_expression(ctx *Case_expressionContext) {}

// ExitCase_expression is called when production case_expression is exited.
func (s *BaseSQLParserListener) ExitCase_expression(ctx *Case_expressionContext) {}

// EnterCast_specification is called when production cast_specification is entered.
func (s *BaseSQLParserListener) EnterCast_specification(ctx *Cast_specificationContext) {}

// ExitCast_specification is called when production cast_specification is exited.
func (s *BaseSQLParserListener) ExitCast_specification(ctx *Cast_specificationContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseSQLParserListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseSQLParserListener) ExitFunction_call(ctx *Function_callContext) {}

// EnterVex_or_named_notation is called when production vex_or_named_notation is entered.
func (s *BaseSQLParserListener) EnterVex_or_named_notation(ctx *Vex_or_named_notationContext) {}

// ExitVex_or_named_notation is called when production vex_or_named_notation is exited.
func (s *BaseSQLParserListener) ExitVex_or_named_notation(ctx *Vex_or_named_notationContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BaseSQLParserListener) EnterPointer(ctx *PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *BaseSQLParserListener) ExitPointer(ctx *PointerContext) {}

// EnterFunction_construct is called when production function_construct is entered.
func (s *BaseSQLParserListener) EnterFunction_construct(ctx *Function_constructContext) {}

// ExitFunction_construct is called when production function_construct is exited.
func (s *BaseSQLParserListener) ExitFunction_construct(ctx *Function_constructContext) {}

// EnterExtract_function is called when production extract_function is entered.
func (s *BaseSQLParserListener) EnterExtract_function(ctx *Extract_functionContext) {}

// ExitExtract_function is called when production extract_function is exited.
func (s *BaseSQLParserListener) ExitExtract_function(ctx *Extract_functionContext) {}

// EnterSystem_function is called when production system_function is entered.
func (s *BaseSQLParserListener) EnterSystem_function(ctx *System_functionContext) {}

// ExitSystem_function is called when production system_function is exited.
func (s *BaseSQLParserListener) ExitSystem_function(ctx *System_functionContext) {}

// EnterDate_time_function is called when production date_time_function is entered.
func (s *BaseSQLParserListener) EnterDate_time_function(ctx *Date_time_functionContext) {}

// ExitDate_time_function is called when production date_time_function is exited.
func (s *BaseSQLParserListener) ExitDate_time_function(ctx *Date_time_functionContext) {}

// EnterString_value_function is called when production string_value_function is entered.
func (s *BaseSQLParserListener) EnterString_value_function(ctx *String_value_functionContext) {}

// ExitString_value_function is called when production string_value_function is exited.
func (s *BaseSQLParserListener) ExitString_value_function(ctx *String_value_functionContext) {}

// EnterXml_function is called when production xml_function is entered.
func (s *BaseSQLParserListener) EnterXml_function(ctx *Xml_functionContext) {}

// ExitXml_function is called when production xml_function is exited.
func (s *BaseSQLParserListener) ExitXml_function(ctx *Xml_functionContext) {}

// EnterXml_table_column is called when production xml_table_column is entered.
func (s *BaseSQLParserListener) EnterXml_table_column(ctx *Xml_table_columnContext) {}

// ExitXml_table_column is called when production xml_table_column is exited.
func (s *BaseSQLParserListener) ExitXml_table_column(ctx *Xml_table_columnContext) {}

// EnterComparison_mod is called when production comparison_mod is entered.
func (s *BaseSQLParserListener) EnterComparison_mod(ctx *Comparison_modContext) {}

// ExitComparison_mod is called when production comparison_mod is exited.
func (s *BaseSQLParserListener) ExitComparison_mod(ctx *Comparison_modContext) {}

// EnterFilter_clause is called when production filter_clause is entered.
func (s *BaseSQLParserListener) EnterFilter_clause(ctx *Filter_clauseContext) {}

// ExitFilter_clause is called when production filter_clause is exited.
func (s *BaseSQLParserListener) ExitFilter_clause(ctx *Filter_clauseContext) {}

// EnterWindow_definition is called when production window_definition is entered.
func (s *BaseSQLParserListener) EnterWindow_definition(ctx *Window_definitionContext) {}

// ExitWindow_definition is called when production window_definition is exited.
func (s *BaseSQLParserListener) ExitWindow_definition(ctx *Window_definitionContext) {}

// EnterFrame_clause is called when production frame_clause is entered.
func (s *BaseSQLParserListener) EnterFrame_clause(ctx *Frame_clauseContext) {}

// ExitFrame_clause is called when production frame_clause is exited.
func (s *BaseSQLParserListener) ExitFrame_clause(ctx *Frame_clauseContext) {}

// EnterFrame_bound is called when production frame_bound is entered.
func (s *BaseSQLParserListener) EnterFrame_bound(ctx *Frame_boundContext) {}

// ExitFrame_bound is called when production frame_bound is exited.
func (s *BaseSQLParserListener) ExitFrame_bound(ctx *Frame_boundContext) {}

// EnterArray_expression is called when production array_expression is entered.
func (s *BaseSQLParserListener) EnterArray_expression(ctx *Array_expressionContext) {}

// ExitArray_expression is called when production array_expression is exited.
func (s *BaseSQLParserListener) ExitArray_expression(ctx *Array_expressionContext) {}

// EnterArray_elements is called when production array_elements is entered.
func (s *BaseSQLParserListener) EnterArray_elements(ctx *Array_elementsContext) {}

// ExitArray_elements is called when production array_elements is exited.
func (s *BaseSQLParserListener) ExitArray_elements(ctx *Array_elementsContext) {}

// EnterType_coercion is called when production type_coercion is entered.
func (s *BaseSQLParserListener) EnterType_coercion(ctx *Type_coercionContext) {}

// ExitType_coercion is called when production type_coercion is exited.
func (s *BaseSQLParserListener) ExitType_coercion(ctx *Type_coercionContext) {}

// EnterSchema_qualified_name is called when production schema_qualified_name is entered.
func (s *BaseSQLParserListener) EnterSchema_qualified_name(ctx *Schema_qualified_nameContext) {}

// ExitSchema_qualified_name is called when production schema_qualified_name is exited.
func (s *BaseSQLParserListener) ExitSchema_qualified_name(ctx *Schema_qualified_nameContext) {}

// EnterSet_qualifier is called when production set_qualifier is entered.
func (s *BaseSQLParserListener) EnterSet_qualifier(ctx *Set_qualifierContext) {}

// ExitSet_qualifier is called when production set_qualifier is exited.
func (s *BaseSQLParserListener) ExitSet_qualifier(ctx *Set_qualifierContext) {}

// EnterTable_subquery is called when production table_subquery is entered.
func (s *BaseSQLParserListener) EnterTable_subquery(ctx *Table_subqueryContext) {}

// ExitTable_subquery is called when production table_subquery is exited.
func (s *BaseSQLParserListener) ExitTable_subquery(ctx *Table_subqueryContext) {}

// EnterSelect_stmt is called when production select_stmt is entered.
func (s *BaseSQLParserListener) EnterSelect_stmt(ctx *Select_stmtContext) {}

// ExitSelect_stmt is called when production select_stmt is exited.
func (s *BaseSQLParserListener) ExitSelect_stmt(ctx *Select_stmtContext) {}

// EnterAfter_ops is called when production after_ops is entered.
func (s *BaseSQLParserListener) EnterAfter_ops(ctx *After_opsContext) {}

// ExitAfter_ops is called when production after_ops is exited.
func (s *BaseSQLParserListener) ExitAfter_ops(ctx *After_opsContext) {}

// EnterSelect_stmt_no_parens is called when production select_stmt_no_parens is entered.
func (s *BaseSQLParserListener) EnterSelect_stmt_no_parens(ctx *Select_stmt_no_parensContext) {}

// ExitSelect_stmt_no_parens is called when production select_stmt_no_parens is exited.
func (s *BaseSQLParserListener) ExitSelect_stmt_no_parens(ctx *Select_stmt_no_parensContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BaseSQLParserListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BaseSQLParserListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterWith_query is called when production with_query is entered.
func (s *BaseSQLParserListener) EnterWith_query(ctx *With_queryContext) {}

// ExitWith_query is called when production with_query is exited.
func (s *BaseSQLParserListener) ExitWith_query(ctx *With_queryContext) {}

// EnterSelect_ops is called when production select_ops is entered.
func (s *BaseSQLParserListener) EnterSelect_ops(ctx *Select_opsContext) {}

// ExitSelect_ops is called when production select_ops is exited.
func (s *BaseSQLParserListener) ExitSelect_ops(ctx *Select_opsContext) {}

// EnterSelect_ops_no_parens is called when production select_ops_no_parens is entered.
func (s *BaseSQLParserListener) EnterSelect_ops_no_parens(ctx *Select_ops_no_parensContext) {}

// ExitSelect_ops_no_parens is called when production select_ops_no_parens is exited.
func (s *BaseSQLParserListener) ExitSelect_ops_no_parens(ctx *Select_ops_no_parensContext) {}

// EnterSelect_primary is called when production select_primary is entered.
func (s *BaseSQLParserListener) EnterSelect_primary(ctx *Select_primaryContext) {}

// ExitSelect_primary is called when production select_primary is exited.
func (s *BaseSQLParserListener) ExitSelect_primary(ctx *Select_primaryContext) {}

// EnterSelect_list is called when production select_list is entered.
func (s *BaseSQLParserListener) EnterSelect_list(ctx *Select_listContext) {}

// ExitSelect_list is called when production select_list is exited.
func (s *BaseSQLParserListener) ExitSelect_list(ctx *Select_listContext) {}

// EnterSelect_sublist is called when production select_sublist is entered.
func (s *BaseSQLParserListener) EnterSelect_sublist(ctx *Select_sublistContext) {}

// ExitSelect_sublist is called when production select_sublist is exited.
func (s *BaseSQLParserListener) ExitSelect_sublist(ctx *Select_sublistContext) {}

// EnterInto_table is called when production into_table is entered.
func (s *BaseSQLParserListener) EnterInto_table(ctx *Into_tableContext) {}

// ExitInto_table is called when production into_table is exited.
func (s *BaseSQLParserListener) ExitInto_table(ctx *Into_tableContext) {}

// EnterFrom_item is called when production from_item is entered.
func (s *BaseSQLParserListener) EnterFrom_item(ctx *From_itemContext) {}

// ExitFrom_item is called when production from_item is exited.
func (s *BaseSQLParserListener) ExitFrom_item(ctx *From_itemContext) {}

// EnterFrom_primary is called when production from_primary is entered.
func (s *BaseSQLParserListener) EnterFrom_primary(ctx *From_primaryContext) {}

// ExitFrom_primary is called when production from_primary is exited.
func (s *BaseSQLParserListener) ExitFrom_primary(ctx *From_primaryContext) {}

// EnterAlias_clause is called when production alias_clause is entered.
func (s *BaseSQLParserListener) EnterAlias_clause(ctx *Alias_clauseContext) {}

// ExitAlias_clause is called when production alias_clause is exited.
func (s *BaseSQLParserListener) ExitAlias_clause(ctx *Alias_clauseContext) {}

// EnterFrom_function_column_def is called when production from_function_column_def is entered.
func (s *BaseSQLParserListener) EnterFrom_function_column_def(ctx *From_function_column_defContext) {}

// ExitFrom_function_column_def is called when production from_function_column_def is exited.
func (s *BaseSQLParserListener) ExitFrom_function_column_def(ctx *From_function_column_defContext) {}

// EnterGroupby_clause is called when production groupby_clause is entered.
func (s *BaseSQLParserListener) EnterGroupby_clause(ctx *Groupby_clauseContext) {}

// ExitGroupby_clause is called when production groupby_clause is exited.
func (s *BaseSQLParserListener) ExitGroupby_clause(ctx *Groupby_clauseContext) {}

// EnterGrouping_element_list is called when production grouping_element_list is entered.
func (s *BaseSQLParserListener) EnterGrouping_element_list(ctx *Grouping_element_listContext) {}

// ExitGrouping_element_list is called when production grouping_element_list is exited.
func (s *BaseSQLParserListener) ExitGrouping_element_list(ctx *Grouping_element_listContext) {}

// EnterGrouping_element is called when production grouping_element is entered.
func (s *BaseSQLParserListener) EnterGrouping_element(ctx *Grouping_elementContext) {}

// ExitGrouping_element is called when production grouping_element is exited.
func (s *BaseSQLParserListener) ExitGrouping_element(ctx *Grouping_elementContext) {}

// EnterValues_stmt is called when production values_stmt is entered.
func (s *BaseSQLParserListener) EnterValues_stmt(ctx *Values_stmtContext) {}

// ExitValues_stmt is called when production values_stmt is exited.
func (s *BaseSQLParserListener) ExitValues_stmt(ctx *Values_stmtContext) {}

// EnterValues_values is called when production values_values is entered.
func (s *BaseSQLParserListener) EnterValues_values(ctx *Values_valuesContext) {}

// ExitValues_values is called when production values_values is exited.
func (s *BaseSQLParserListener) ExitValues_values(ctx *Values_valuesContext) {}

// EnterOrderby_clause is called when production orderby_clause is entered.
func (s *BaseSQLParserListener) EnterOrderby_clause(ctx *Orderby_clauseContext) {}

// ExitOrderby_clause is called when production orderby_clause is exited.
func (s *BaseSQLParserListener) ExitOrderby_clause(ctx *Orderby_clauseContext) {}

// EnterSort_specifier is called when production sort_specifier is entered.
func (s *BaseSQLParserListener) EnterSort_specifier(ctx *Sort_specifierContext) {}

// ExitSort_specifier is called when production sort_specifier is exited.
func (s *BaseSQLParserListener) ExitSort_specifier(ctx *Sort_specifierContext) {}

// EnterOrder_specification is called when production order_specification is entered.
func (s *BaseSQLParserListener) EnterOrder_specification(ctx *Order_specificationContext) {}

// ExitOrder_specification is called when production order_specification is exited.
func (s *BaseSQLParserListener) ExitOrder_specification(ctx *Order_specificationContext) {}

// EnterNull_ordering is called when production null_ordering is entered.
func (s *BaseSQLParserListener) EnterNull_ordering(ctx *Null_orderingContext) {}

// ExitNull_ordering is called when production null_ordering is exited.
func (s *BaseSQLParserListener) ExitNull_ordering(ctx *Null_orderingContext) {}

// EnterInsert_stmt_for_psql is called when production insert_stmt_for_psql is entered.
func (s *BaseSQLParserListener) EnterInsert_stmt_for_psql(ctx *Insert_stmt_for_psqlContext) {}

// ExitInsert_stmt_for_psql is called when production insert_stmt_for_psql is exited.
func (s *BaseSQLParserListener) ExitInsert_stmt_for_psql(ctx *Insert_stmt_for_psqlContext) {}

// EnterInsert_columns is called when production insert_columns is entered.
func (s *BaseSQLParserListener) EnterInsert_columns(ctx *Insert_columnsContext) {}

// ExitInsert_columns is called when production insert_columns is exited.
func (s *BaseSQLParserListener) ExitInsert_columns(ctx *Insert_columnsContext) {}

// EnterIndirection_identifier is called when production indirection_identifier is entered.
func (s *BaseSQLParserListener) EnterIndirection_identifier(ctx *Indirection_identifierContext) {}

// ExitIndirection_identifier is called when production indirection_identifier is exited.
func (s *BaseSQLParserListener) ExitIndirection_identifier(ctx *Indirection_identifierContext) {}

// EnterConflict_object is called when production conflict_object is entered.
func (s *BaseSQLParserListener) EnterConflict_object(ctx *Conflict_objectContext) {}

// ExitConflict_object is called when production conflict_object is exited.
func (s *BaseSQLParserListener) ExitConflict_object(ctx *Conflict_objectContext) {}

// EnterConflict_action is called when production conflict_action is entered.
func (s *BaseSQLParserListener) EnterConflict_action(ctx *Conflict_actionContext) {}

// ExitConflict_action is called when production conflict_action is exited.
func (s *BaseSQLParserListener) ExitConflict_action(ctx *Conflict_actionContext) {}

// EnterDelete_stmt_for_psql is called when production delete_stmt_for_psql is entered.
func (s *BaseSQLParserListener) EnterDelete_stmt_for_psql(ctx *Delete_stmt_for_psqlContext) {}

// ExitDelete_stmt_for_psql is called when production delete_stmt_for_psql is exited.
func (s *BaseSQLParserListener) ExitDelete_stmt_for_psql(ctx *Delete_stmt_for_psqlContext) {}

// EnterUpdate_stmt_for_psql is called when production update_stmt_for_psql is entered.
func (s *BaseSQLParserListener) EnterUpdate_stmt_for_psql(ctx *Update_stmt_for_psqlContext) {}

// ExitUpdate_stmt_for_psql is called when production update_stmt_for_psql is exited.
func (s *BaseSQLParserListener) ExitUpdate_stmt_for_psql(ctx *Update_stmt_for_psqlContext) {}

// EnterUpdate_set is called when production update_set is entered.
func (s *BaseSQLParserListener) EnterUpdate_set(ctx *Update_setContext) {}

// ExitUpdate_set is called when production update_set is exited.
func (s *BaseSQLParserListener) ExitUpdate_set(ctx *Update_setContext) {}

// EnterNotify_stmt is called when production notify_stmt is entered.
func (s *BaseSQLParserListener) EnterNotify_stmt(ctx *Notify_stmtContext) {}

// ExitNotify_stmt is called when production notify_stmt is exited.
func (s *BaseSQLParserListener) ExitNotify_stmt(ctx *Notify_stmtContext) {}

// EnterTruncate_stmt is called when production truncate_stmt is entered.
func (s *BaseSQLParserListener) EnterTruncate_stmt(ctx *Truncate_stmtContext) {}

// ExitTruncate_stmt is called when production truncate_stmt is exited.
func (s *BaseSQLParserListener) ExitTruncate_stmt(ctx *Truncate_stmtContext) {}

// EnterIdentifier_list is called when production identifier_list is entered.
func (s *BaseSQLParserListener) EnterIdentifier_list(ctx *Identifier_listContext) {}

// ExitIdentifier_list is called when production identifier_list is exited.
func (s *BaseSQLParserListener) ExitIdentifier_list(ctx *Identifier_listContext) {}

// EnterAnonymous_block is called when production anonymous_block is entered.
func (s *BaseSQLParserListener) EnterAnonymous_block(ctx *Anonymous_blockContext) {}

// ExitAnonymous_block is called when production anonymous_block is exited.
func (s *BaseSQLParserListener) ExitAnonymous_block(ctx *Anonymous_blockContext) {}

// EnterComp_options is called when production comp_options is entered.
func (s *BaseSQLParserListener) EnterComp_options(ctx *Comp_optionsContext) {}

// ExitComp_options is called when production comp_options is exited.
func (s *BaseSQLParserListener) ExitComp_options(ctx *Comp_optionsContext) {}

// EnterFunction_block is called when production function_block is entered.
func (s *BaseSQLParserListener) EnterFunction_block(ctx *Function_blockContext) {}

// ExitFunction_block is called when production function_block is exited.
func (s *BaseSQLParserListener) ExitFunction_block(ctx *Function_blockContext) {}

// EnterStart_label is called when production start_label is entered.
func (s *BaseSQLParserListener) EnterStart_label(ctx *Start_labelContext) {}

// ExitStart_label is called when production start_label is exited.
func (s *BaseSQLParserListener) ExitStart_label(ctx *Start_labelContext) {}

// EnterDeclarations is called when production declarations is entered.
func (s *BaseSQLParserListener) EnterDeclarations(ctx *DeclarationsContext) {}

// ExitDeclarations is called when production declarations is exited.
func (s *BaseSQLParserListener) ExitDeclarations(ctx *DeclarationsContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseSQLParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseSQLParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterType_declaration is called when production type_declaration is entered.
func (s *BaseSQLParserListener) EnterType_declaration(ctx *Type_declarationContext) {}

// ExitType_declaration is called when production type_declaration is exited.
func (s *BaseSQLParserListener) ExitType_declaration(ctx *Type_declarationContext) {}

// EnterArguments_list is called when production arguments_list is entered.
func (s *BaseSQLParserListener) EnterArguments_list(ctx *Arguments_listContext) {}

// ExitArguments_list is called when production arguments_list is exited.
func (s *BaseSQLParserListener) ExitArguments_list(ctx *Arguments_listContext) {}

// EnterData_type_dec is called when production data_type_dec is entered.
func (s *BaseSQLParserListener) EnterData_type_dec(ctx *Data_type_decContext) {}

// ExitData_type_dec is called when production data_type_dec is exited.
func (s *BaseSQLParserListener) ExitData_type_dec(ctx *Data_type_decContext) {}

// EnterException_statement is called when production exception_statement is entered.
func (s *BaseSQLParserListener) EnterException_statement(ctx *Exception_statementContext) {}

// ExitException_statement is called when production exception_statement is exited.
func (s *BaseSQLParserListener) ExitException_statement(ctx *Exception_statementContext) {}

// EnterFunction_statements is called when production function_statements is entered.
func (s *BaseSQLParserListener) EnterFunction_statements(ctx *Function_statementsContext) {}

// ExitFunction_statements is called when production function_statements is exited.
func (s *BaseSQLParserListener) ExitFunction_statements(ctx *Function_statementsContext) {}

// EnterFunction_statement is called when production function_statement is entered.
func (s *BaseSQLParserListener) EnterFunction_statement(ctx *Function_statementContext) {}

// ExitFunction_statement is called when production function_statement is exited.
func (s *BaseSQLParserListener) ExitFunction_statement(ctx *Function_statementContext) {}

// EnterBase_statement is called when production base_statement is entered.
func (s *BaseSQLParserListener) EnterBase_statement(ctx *Base_statementContext) {}

// ExitBase_statement is called when production base_statement is exited.
func (s *BaseSQLParserListener) ExitBase_statement(ctx *Base_statementContext) {}

// EnterX_var is called when production x_var is entered.
func (s *BaseSQLParserListener) EnterX_var(ctx *X_varContext) {}

// ExitX_var is called when production x_var is exited.
func (s *BaseSQLParserListener) ExitX_var(ctx *X_varContext) {}

// EnterDiagnostic_option is called when production diagnostic_option is entered.
func (s *BaseSQLParserListener) EnterDiagnostic_option(ctx *Diagnostic_optionContext) {}

// ExitDiagnostic_option is called when production diagnostic_option is exited.
func (s *BaseSQLParserListener) ExitDiagnostic_option(ctx *Diagnostic_optionContext) {}

// EnterPerform_stmt is called when production perform_stmt is entered.
func (s *BaseSQLParserListener) EnterPerform_stmt(ctx *Perform_stmtContext) {}

// ExitPerform_stmt is called when production perform_stmt is exited.
func (s *BaseSQLParserListener) ExitPerform_stmt(ctx *Perform_stmtContext) {}

// EnterAssign_stmt is called when production assign_stmt is entered.
func (s *BaseSQLParserListener) EnterAssign_stmt(ctx *Assign_stmtContext) {}

// ExitAssign_stmt is called when production assign_stmt is exited.
func (s *BaseSQLParserListener) ExitAssign_stmt(ctx *Assign_stmtContext) {}

// EnterExecute_stmt is called when production execute_stmt is entered.
func (s *BaseSQLParserListener) EnterExecute_stmt(ctx *Execute_stmtContext) {}

// ExitExecute_stmt is called when production execute_stmt is exited.
func (s *BaseSQLParserListener) ExitExecute_stmt(ctx *Execute_stmtContext) {}

// EnterControl_statement is called when production control_statement is entered.
func (s *BaseSQLParserListener) EnterControl_statement(ctx *Control_statementContext) {}

// ExitControl_statement is called when production control_statement is exited.
func (s *BaseSQLParserListener) ExitControl_statement(ctx *Control_statementContext) {}

// EnterCursor_statement is called when production cursor_statement is entered.
func (s *BaseSQLParserListener) EnterCursor_statement(ctx *Cursor_statementContext) {}

// ExitCursor_statement is called when production cursor_statement is exited.
func (s *BaseSQLParserListener) ExitCursor_statement(ctx *Cursor_statementContext) {}

// EnterOption is called when production option is entered.
func (s *BaseSQLParserListener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseSQLParserListener) ExitOption(ctx *OptionContext) {}

// EnterTransaction_statement is called when production transaction_statement is entered.
func (s *BaseSQLParserListener) EnterTransaction_statement(ctx *Transaction_statementContext) {}

// ExitTransaction_statement is called when production transaction_statement is exited.
func (s *BaseSQLParserListener) ExitTransaction_statement(ctx *Transaction_statementContext) {}

// EnterMessage_statement is called when production message_statement is entered.
func (s *BaseSQLParserListener) EnterMessage_statement(ctx *Message_statementContext) {}

// ExitMessage_statement is called when production message_statement is exited.
func (s *BaseSQLParserListener) ExitMessage_statement(ctx *Message_statementContext) {}

// EnterLog_level is called when production log_level is entered.
func (s *BaseSQLParserListener) EnterLog_level(ctx *Log_levelContext) {}

// ExitLog_level is called when production log_level is exited.
func (s *BaseSQLParserListener) ExitLog_level(ctx *Log_levelContext) {}

// EnterRaise_using is called when production raise_using is entered.
func (s *BaseSQLParserListener) EnterRaise_using(ctx *Raise_usingContext) {}

// ExitRaise_using is called when production raise_using is exited.
func (s *BaseSQLParserListener) ExitRaise_using(ctx *Raise_usingContext) {}

// EnterRaise_param is called when production raise_param is entered.
func (s *BaseSQLParserListener) EnterRaise_param(ctx *Raise_paramContext) {}

// ExitRaise_param is called when production raise_param is exited.
func (s *BaseSQLParserListener) ExitRaise_param(ctx *Raise_paramContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseSQLParserListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseSQLParserListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseSQLParserListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseSQLParserListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterLoop_start is called when production loop_start is entered.
func (s *BaseSQLParserListener) EnterLoop_start(ctx *Loop_startContext) {}

// ExitLoop_start is called when production loop_start is exited.
func (s *BaseSQLParserListener) ExitLoop_start(ctx *Loop_startContext) {}

// EnterUsing_vex is called when production using_vex is entered.
func (s *BaseSQLParserListener) EnterUsing_vex(ctx *Using_vexContext) {}

// ExitUsing_vex is called when production using_vex is exited.
func (s *BaseSQLParserListener) ExitUsing_vex(ctx *Using_vexContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseSQLParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseSQLParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterCase_statement is called when production case_statement is entered.
func (s *BaseSQLParserListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BaseSQLParserListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterPlpgsql_query is called when production plpgsql_query is entered.
func (s *BaseSQLParserListener) EnterPlpgsql_query(ctx *Plpgsql_queryContext) {}

// ExitPlpgsql_query is called when production plpgsql_query is exited.
func (s *BaseSQLParserListener) ExitPlpgsql_query(ctx *Plpgsql_queryContext) {}
