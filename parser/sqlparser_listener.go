// Code generated from SQLParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser // SQLParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SQLParserListener is a complete listener for a parse tree produced by SQLParser.
type SQLParserListener interface {
	antlr.ParseTreeListener

	// EnterSql is called when entering the sql production.
	EnterSql(c *SqlContext)

	// EnterQname_parser is called when entering the qname_parser production.
	EnterQname_parser(c *Qname_parserContext)

	// EnterFunction_args_parser is called when entering the function_args_parser production.
	EnterFunction_args_parser(c *Function_args_parserContext)

	// EnterVex_eof is called when entering the vex_eof production.
	EnterVex_eof(c *Vex_eofContext)

	// EnterPlpgsql_function is called when entering the plpgsql_function production.
	EnterPlpgsql_function(c *Plpgsql_functionContext)

	// EnterPlpgsql_function_test_list is called when entering the plpgsql_function_test_list production.
	EnterPlpgsql_function_test_list(c *Plpgsql_function_test_listContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterData_statement is called when entering the data_statement production.
	EnterData_statement(c *Data_statementContext)

	// EnterScript_statement is called when entering the script_statement production.
	EnterScript_statement(c *Script_statementContext)

	// EnterScript_transaction is called when entering the script_transaction production.
	EnterScript_transaction(c *Script_transactionContext)

	// EnterTransaction_mode is called when entering the transaction_mode production.
	EnterTransaction_mode(c *Transaction_modeContext)

	// EnterLock_table is called when entering the lock_table production.
	EnterLock_table(c *Lock_tableContext)

	// EnterLock_mode is called when entering the lock_mode production.
	EnterLock_mode(c *Lock_modeContext)

	// EnterScript_additional is called when entering the script_additional production.
	EnterScript_additional(c *Script_additionalContext)

	// EnterAdditional_statement is called when entering the additional_statement production.
	EnterAdditional_statement(c *Additional_statementContext)

	// EnterExplain_statement is called when entering the explain_statement production.
	EnterExplain_statement(c *Explain_statementContext)

	// EnterExplain_query is called when entering the explain_query production.
	EnterExplain_query(c *Explain_queryContext)

	// EnterExecute_statement is called when entering the execute_statement production.
	EnterExecute_statement(c *Execute_statementContext)

	// EnterDeclare_statement is called when entering the declare_statement production.
	EnterDeclare_statement(c *Declare_statementContext)

	// EnterShow_statement is called when entering the show_statement production.
	EnterShow_statement(c *Show_statementContext)

	// EnterExplain_option is called when entering the explain_option production.
	EnterExplain_option(c *Explain_optionContext)

	// EnterUser_name is called when entering the user_name production.
	EnterUser_name(c *User_nameContext)

	// EnterTable_cols_list is called when entering the table_cols_list production.
	EnterTable_cols_list(c *Table_cols_listContext)

	// EnterTable_cols is called when entering the table_cols production.
	EnterTable_cols(c *Table_colsContext)

	// EnterVacuum_mode is called when entering the vacuum_mode production.
	EnterVacuum_mode(c *Vacuum_modeContext)

	// EnterVacuum_option is called when entering the vacuum_option production.
	EnterVacuum_option(c *Vacuum_optionContext)

	// EnterAnalyze_mode is called when entering the analyze_mode production.
	EnterAnalyze_mode(c *Analyze_modeContext)

	// EnterBoolean_value is called when entering the boolean_value production.
	EnterBoolean_value(c *Boolean_valueContext)

	// EnterFetch_move_direction is called when entering the fetch_move_direction production.
	EnterFetch_move_direction(c *Fetch_move_directionContext)

	// EnterSchema_statement is called when entering the schema_statement production.
	EnterSchema_statement(c *Schema_statementContext)

	// EnterSchema_create is called when entering the schema_create production.
	EnterSchema_create(c *Schema_createContext)

	// EnterSchema_alter is called when entering the schema_alter production.
	EnterSchema_alter(c *Schema_alterContext)

	// EnterSchema_drop is called when entering the schema_drop production.
	EnterSchema_drop(c *Schema_dropContext)

	// EnterSchema_import is called when entering the schema_import production.
	EnterSchema_import(c *Schema_importContext)

	// EnterAlter_function_statement is called when entering the alter_function_statement production.
	EnterAlter_function_statement(c *Alter_function_statementContext)

	// EnterAlter_aggregate_statement is called when entering the alter_aggregate_statement production.
	EnterAlter_aggregate_statement(c *Alter_aggregate_statementContext)

	// EnterAlter_extension_statement is called when entering the alter_extension_statement production.
	EnterAlter_extension_statement(c *Alter_extension_statementContext)

	// EnterAlter_extension_action is called when entering the alter_extension_action production.
	EnterAlter_extension_action(c *Alter_extension_actionContext)

	// EnterExtension_member_object is called when entering the extension_member_object production.
	EnterExtension_member_object(c *Extension_member_objectContext)

	// EnterAlter_schema_statement is called when entering the alter_schema_statement production.
	EnterAlter_schema_statement(c *Alter_schema_statementContext)

	// EnterAlter_language_statement is called when entering the alter_language_statement production.
	EnterAlter_language_statement(c *Alter_language_statementContext)

	// EnterAlter_table_statement is called when entering the alter_table_statement production.
	EnterAlter_table_statement(c *Alter_table_statementContext)

	// EnterTable_action is called when entering the table_action production.
	EnterTable_action(c *Table_actionContext)

	// EnterColumn_action is called when entering the column_action production.
	EnterColumn_action(c *Column_actionContext)

	// EnterIdentity_body is called when entering the identity_body production.
	EnterIdentity_body(c *Identity_bodyContext)

	// EnterAlter_identity is called when entering the alter_identity production.
	EnterAlter_identity(c *Alter_identityContext)

	// EnterStorage_option is called when entering the storage_option production.
	EnterStorage_option(c *Storage_optionContext)

	// EnterValidate_constraint is called when entering the validate_constraint production.
	EnterValidate_constraint(c *Validate_constraintContext)

	// EnterDrop_constraint is called when entering the drop_constraint production.
	EnterDrop_constraint(c *Drop_constraintContext)

	// EnterTable_deferrable is called when entering the table_deferrable production.
	EnterTable_deferrable(c *Table_deferrableContext)

	// EnterTable_initialy_immed is called when entering the table_initialy_immed production.
	EnterTable_initialy_immed(c *Table_initialy_immedContext)

	// EnterFunction_actions_common is called when entering the function_actions_common production.
	EnterFunction_actions_common(c *Function_actions_commonContext)

	// EnterFunction_def is called when entering the function_def production.
	EnterFunction_def(c *Function_defContext)

	// EnterAlter_index_statement is called when entering the alter_index_statement production.
	EnterAlter_index_statement(c *Alter_index_statementContext)

	// EnterIndex_def_action is called when entering the index_def_action production.
	EnterIndex_def_action(c *Index_def_actionContext)

	// EnterAlter_default_privileges_statement is called when entering the alter_default_privileges_statement production.
	EnterAlter_default_privileges_statement(c *Alter_default_privileges_statementContext)

	// EnterAbbreviated_grant_or_revoke is called when entering the abbreviated_grant_or_revoke production.
	EnterAbbreviated_grant_or_revoke(c *Abbreviated_grant_or_revokeContext)

	// EnterGrant_option_for is called when entering the grant_option_for production.
	EnterGrant_option_for(c *Grant_option_forContext)

	// EnterAlter_sequence_statement is called when entering the alter_sequence_statement production.
	EnterAlter_sequence_statement(c *Alter_sequence_statementContext)

	// EnterAlter_view_statement is called when entering the alter_view_statement production.
	EnterAlter_view_statement(c *Alter_view_statementContext)

	// EnterAlter_view_action is called when entering the alter_view_action production.
	EnterAlter_view_action(c *Alter_view_actionContext)

	// EnterAlter_materialized_view_statement is called when entering the alter_materialized_view_statement production.
	EnterAlter_materialized_view_statement(c *Alter_materialized_view_statementContext)

	// EnterAlter_materialized_view_action is called when entering the alter_materialized_view_action production.
	EnterAlter_materialized_view_action(c *Alter_materialized_view_actionContext)

	// EnterMaterialized_view_action is called when entering the materialized_view_action production.
	EnterMaterialized_view_action(c *Materialized_view_actionContext)

	// EnterAlter_event_trigger_statement is called when entering the alter_event_trigger_statement production.
	EnterAlter_event_trigger_statement(c *Alter_event_trigger_statementContext)

	// EnterAlter_event_trigger_action is called when entering the alter_event_trigger_action production.
	EnterAlter_event_trigger_action(c *Alter_event_trigger_actionContext)

	// EnterAlter_type_statement is called when entering the alter_type_statement production.
	EnterAlter_type_statement(c *Alter_type_statementContext)

	// EnterAlter_domain_statement is called when entering the alter_domain_statement production.
	EnterAlter_domain_statement(c *Alter_domain_statementContext)

	// EnterAlter_server_statement is called when entering the alter_server_statement production.
	EnterAlter_server_statement(c *Alter_server_statementContext)

	// EnterAlter_server_action is called when entering the alter_server_action production.
	EnterAlter_server_action(c *Alter_server_actionContext)

	// EnterAlter_fts_statement is called when entering the alter_fts_statement production.
	EnterAlter_fts_statement(c *Alter_fts_statementContext)

	// EnterAlter_fts_configuration is called when entering the alter_fts_configuration production.
	EnterAlter_fts_configuration(c *Alter_fts_configurationContext)

	// EnterType_action is called when entering the type_action production.
	EnterType_action(c *Type_actionContext)

	// EnterType_property is called when entering the type_property production.
	EnterType_property(c *Type_propertyContext)

	// EnterSet_def_column is called when entering the set_def_column production.
	EnterSet_def_column(c *Set_def_columnContext)

	// EnterDrop_def is called when entering the drop_def production.
	EnterDrop_def(c *Drop_defContext)

	// EnterCreate_index_statement is called when entering the create_index_statement production.
	EnterCreate_index_statement(c *Create_index_statementContext)

	// EnterIndex_rest is called when entering the index_rest production.
	EnterIndex_rest(c *Index_restContext)

	// EnterIndex_sort is called when entering the index_sort production.
	EnterIndex_sort(c *Index_sortContext)

	// EnterIndex_column is called when entering the index_column production.
	EnterIndex_column(c *Index_columnContext)

	// EnterIncluding_index is called when entering the including_index production.
	EnterIncluding_index(c *Including_indexContext)

	// EnterIndex_where is called when entering the index_where production.
	EnterIndex_where(c *Index_whereContext)

	// EnterCreate_extension_statement is called when entering the create_extension_statement production.
	EnterCreate_extension_statement(c *Create_extension_statementContext)

	// EnterCreate_language_statement is called when entering the create_language_statement production.
	EnterCreate_language_statement(c *Create_language_statementContext)

	// EnterCreate_event_trigger_statement is called when entering the create_event_trigger_statement production.
	EnterCreate_event_trigger_statement(c *Create_event_trigger_statementContext)

	// EnterCreate_type_statement is called when entering the create_type_statement production.
	EnterCreate_type_statement(c *Create_type_statementContext)

	// EnterCreate_domain_statement is called when entering the create_domain_statement production.
	EnterCreate_domain_statement(c *Create_domain_statementContext)

	// EnterCreate_server_statement is called when entering the create_server_statement production.
	EnterCreate_server_statement(c *Create_server_statementContext)

	// EnterCreate_fts_dictionary_statement is called when entering the create_fts_dictionary_statement production.
	EnterCreate_fts_dictionary_statement(c *Create_fts_dictionary_statementContext)

	// EnterOption_with_value is called when entering the option_with_value production.
	EnterOption_with_value(c *Option_with_valueContext)

	// EnterCreate_fts_configuration_statement is called when entering the create_fts_configuration_statement production.
	EnterCreate_fts_configuration_statement(c *Create_fts_configuration_statementContext)

	// EnterCreate_fts_template_statement is called when entering the create_fts_template_statement production.
	EnterCreate_fts_template_statement(c *Create_fts_template_statementContext)

	// EnterCreate_fts_parser_statement is called when entering the create_fts_parser_statement production.
	EnterCreate_fts_parser_statement(c *Create_fts_parser_statementContext)

	// EnterCreate_collation_statement is called when entering the create_collation_statement production.
	EnterCreate_collation_statement(c *Create_collation_statementContext)

	// EnterAlter_collation_statement is called when entering the alter_collation_statement production.
	EnterAlter_collation_statement(c *Alter_collation_statementContext)

	// EnterCollation_option is called when entering the collation_option production.
	EnterCollation_option(c *Collation_optionContext)

	// EnterCreate_user_mapping_statement is called when entering the create_user_mapping_statement production.
	EnterCreate_user_mapping_statement(c *Create_user_mapping_statementContext)

	// EnterAlter_user_mapping_statement is called when entering the alter_user_mapping_statement production.
	EnterAlter_user_mapping_statement(c *Alter_user_mapping_statementContext)

	// EnterAlter_user_or_role_statement is called when entering the alter_user_or_role_statement production.
	EnterAlter_user_or_role_statement(c *Alter_user_or_role_statementContext)

	// EnterAlter_user_or_role_set_reset is called when entering the alter_user_or_role_set_reset production.
	EnterAlter_user_or_role_set_reset(c *Alter_user_or_role_set_resetContext)

	// EnterSet_reset_parameter is called when entering the set_reset_parameter production.
	EnterSet_reset_parameter(c *Set_reset_parameterContext)

	// EnterAlter_group_statement is called when entering the alter_group_statement production.
	EnterAlter_group_statement(c *Alter_group_statementContext)

	// EnterAlter_group_action is called when entering the alter_group_action production.
	EnterAlter_group_action(c *Alter_group_actionContext)

	// EnterAlter_tablespace_statement is called when entering the alter_tablespace_statement production.
	EnterAlter_tablespace_statement(c *Alter_tablespace_statementContext)

	// EnterAlter_owner_statement is called when entering the alter_owner_statement production.
	EnterAlter_owner_statement(c *Alter_owner_statementContext)

	// EnterAlter_tablespace_action is called when entering the alter_tablespace_action production.
	EnterAlter_tablespace_action(c *Alter_tablespace_actionContext)

	// EnterAlter_statistics_statement is called when entering the alter_statistics_statement production.
	EnterAlter_statistics_statement(c *Alter_statistics_statementContext)

	// EnterSet_statistics is called when entering the set_statistics production.
	EnterSet_statistics(c *Set_statisticsContext)

	// EnterAlter_foreign_data_wrapper is called when entering the alter_foreign_data_wrapper production.
	EnterAlter_foreign_data_wrapper(c *Alter_foreign_data_wrapperContext)

	// EnterAlter_foreign_data_wrapper_action is called when entering the alter_foreign_data_wrapper_action production.
	EnterAlter_foreign_data_wrapper_action(c *Alter_foreign_data_wrapper_actionContext)

	// EnterAlter_operator_statement is called when entering the alter_operator_statement production.
	EnterAlter_operator_statement(c *Alter_operator_statementContext)

	// EnterAlter_operator_action is called when entering the alter_operator_action production.
	EnterAlter_operator_action(c *Alter_operator_actionContext)

	// EnterOperator_set_restrict_join is called when entering the operator_set_restrict_join production.
	EnterOperator_set_restrict_join(c *Operator_set_restrict_joinContext)

	// EnterDrop_user_mapping_statement is called when entering the drop_user_mapping_statement production.
	EnterDrop_user_mapping_statement(c *Drop_user_mapping_statementContext)

	// EnterDrop_owned_statement is called when entering the drop_owned_statement production.
	EnterDrop_owned_statement(c *Drop_owned_statementContext)

	// EnterDrop_operator_statement is called when entering the drop_operator_statement production.
	EnterDrop_operator_statement(c *Drop_operator_statementContext)

	// EnterTarget_operator is called when entering the target_operator production.
	EnterTarget_operator(c *Target_operatorContext)

	// EnterDomain_constraint is called when entering the domain_constraint production.
	EnterDomain_constraint(c *Domain_constraintContext)

	// EnterCreate_transform_statement is called when entering the create_transform_statement production.
	EnterCreate_transform_statement(c *Create_transform_statementContext)

	// EnterCreate_access_method_statement is called when entering the create_access_method_statement production.
	EnterCreate_access_method_statement(c *Create_access_method_statementContext)

	// EnterCreate_user_or_role_statement is called when entering the create_user_or_role_statement production.
	EnterCreate_user_or_role_statement(c *Create_user_or_role_statementContext)

	// EnterUser_or_role_option is called when entering the user_or_role_option production.
	EnterUser_or_role_option(c *User_or_role_optionContext)

	// EnterUser_or_role_option_for_alter is called when entering the user_or_role_option_for_alter production.
	EnterUser_or_role_option_for_alter(c *User_or_role_option_for_alterContext)

	// EnterUser_or_role_or_group_common_option is called when entering the user_or_role_or_group_common_option production.
	EnterUser_or_role_or_group_common_option(c *User_or_role_or_group_common_optionContext)

	// EnterUser_or_role_common_option is called when entering the user_or_role_common_option production.
	EnterUser_or_role_common_option(c *User_or_role_common_optionContext)

	// EnterUser_or_role_or_group_option_for_create is called when entering the user_or_role_or_group_option_for_create production.
	EnterUser_or_role_or_group_option_for_create(c *User_or_role_or_group_option_for_createContext)

	// EnterCreate_group_statement is called when entering the create_group_statement production.
	EnterCreate_group_statement(c *Create_group_statementContext)

	// EnterGroup_option is called when entering the group_option production.
	EnterGroup_option(c *Group_optionContext)

	// EnterCreate_tablespace_statement is called when entering the create_tablespace_statement production.
	EnterCreate_tablespace_statement(c *Create_tablespace_statementContext)

	// EnterCreate_statistics_statement is called when entering the create_statistics_statement production.
	EnterCreate_statistics_statement(c *Create_statistics_statementContext)

	// EnterCreate_foreign_data_wrapper_statement is called when entering the create_foreign_data_wrapper_statement production.
	EnterCreate_foreign_data_wrapper_statement(c *Create_foreign_data_wrapper_statementContext)

	// EnterOption_without_equal is called when entering the option_without_equal production.
	EnterOption_without_equal(c *Option_without_equalContext)

	// EnterCreate_operator_statement is called when entering the create_operator_statement production.
	EnterCreate_operator_statement(c *Create_operator_statementContext)

	// EnterOperator_name is called when entering the operator_name production.
	EnterOperator_name(c *Operator_nameContext)

	// EnterOperator_option is called when entering the operator_option production.
	EnterOperator_option(c *Operator_optionContext)

	// EnterCreate_aggregate_statement is called when entering the create_aggregate_statement production.
	EnterCreate_aggregate_statement(c *Create_aggregate_statementContext)

	// EnterAggregate_param is called when entering the aggregate_param production.
	EnterAggregate_param(c *Aggregate_paramContext)

	// EnterSet_statement is called when entering the set_statement production.
	EnterSet_statement(c *Set_statementContext)

	// EnterSet_action is called when entering the set_action production.
	EnterSet_action(c *Set_actionContext)

	// EnterSession_local_option is called when entering the session_local_option production.
	EnterSession_local_option(c *Session_local_optionContext)

	// EnterSet_statement_value is called when entering the set_statement_value production.
	EnterSet_statement_value(c *Set_statement_valueContext)

	// EnterCreate_rewrite_statement is called when entering the create_rewrite_statement production.
	EnterCreate_rewrite_statement(c *Create_rewrite_statementContext)

	// EnterRewrite_command is called when entering the rewrite_command production.
	EnterRewrite_command(c *Rewrite_commandContext)

	// EnterCreate_trigger_statement is called when entering the create_trigger_statement production.
	EnterCreate_trigger_statement(c *Create_trigger_statementContext)

	// EnterTrigger_referencing is called when entering the trigger_referencing production.
	EnterTrigger_referencing(c *Trigger_referencingContext)

	// EnterWhen_trigger is called when entering the when_trigger production.
	EnterWhen_trigger(c *When_triggerContext)

	// EnterRule_common is called when entering the rule_common production.
	EnterRule_common(c *Rule_commonContext)

	// EnterRule_member_object is called when entering the rule_member_object production.
	EnterRule_member_object(c *Rule_member_objectContext)

	// EnterColumns_permissions is called when entering the columns_permissions production.
	EnterColumns_permissions(c *Columns_permissionsContext)

	// EnterTable_column_privileges is called when entering the table_column_privileges production.
	EnterTable_column_privileges(c *Table_column_privilegesContext)

	// EnterPermissions is called when entering the permissions production.
	EnterPermissions(c *PermissionsContext)

	// EnterPermission is called when entering the permission production.
	EnterPermission(c *PermissionContext)

	// EnterOther_rules is called when entering the other_rules production.
	EnterOther_rules(c *Other_rulesContext)

	// EnterGrant_to_rule is called when entering the grant_to_rule production.
	EnterGrant_to_rule(c *Grant_to_ruleContext)

	// EnterRevoke_from_cascade_restrict is called when entering the revoke_from_cascade_restrict production.
	EnterRevoke_from_cascade_restrict(c *Revoke_from_cascade_restrictContext)

	// EnterRoles_names is called when entering the roles_names production.
	EnterRoles_names(c *Roles_namesContext)

	// EnterRole_name_with_group is called when entering the role_name_with_group production.
	EnterRole_name_with_group(c *Role_name_with_groupContext)

	// EnterComment_on_statement is called when entering the comment_on_statement production.
	EnterComment_on_statement(c *Comment_on_statementContext)

	// EnterSecurity_label is called when entering the security_label production.
	EnterSecurity_label(c *Security_labelContext)

	// EnterComment_member_object is called when entering the comment_member_object production.
	EnterComment_member_object(c *Comment_member_objectContext)

	// EnterLabel_member_object is called when entering the label_member_object production.
	EnterLabel_member_object(c *Label_member_objectContext)

	// EnterCreate_function_statement is called when entering the create_function_statement production.
	EnterCreate_function_statement(c *Create_function_statementContext)

	// EnterCreate_funct_params is called when entering the create_funct_params production.
	EnterCreate_funct_params(c *Create_funct_paramsContext)

	// EnterTransform_for_type is called when entering the transform_for_type production.
	EnterTransform_for_type(c *Transform_for_typeContext)

	// EnterFunction_ret_table is called when entering the function_ret_table production.
	EnterFunction_ret_table(c *Function_ret_tableContext)

	// EnterFunction_column_name_type is called when entering the function_column_name_type production.
	EnterFunction_column_name_type(c *Function_column_name_typeContext)

	// EnterFunction_parameters is called when entering the function_parameters production.
	EnterFunction_parameters(c *Function_parametersContext)

	// EnterFunction_args is called when entering the function_args production.
	EnterFunction_args(c *Function_argsContext)

	// EnterAgg_order is called when entering the agg_order production.
	EnterAgg_order(c *Agg_orderContext)

	// EnterCharacter_string is called when entering the character_string production.
	EnterCharacter_string(c *Character_stringContext)

	// EnterFunction_arguments is called when entering the function_arguments production.
	EnterFunction_arguments(c *Function_argumentsContext)

	// EnterArgmode is called when entering the argmode production.
	EnterArgmode(c *ArgmodeContext)

	// EnterCreate_sequence_statement is called when entering the create_sequence_statement production.
	EnterCreate_sequence_statement(c *Create_sequence_statementContext)

	// EnterSequence_body is called when entering the sequence_body production.
	EnterSequence_body(c *Sequence_bodyContext)

	// EnterSigned_number_literal is called when entering the signed_number_literal production.
	EnterSigned_number_literal(c *Signed_number_literalContext)

	// EnterSigned_numerical_literal is called when entering the signed_numerical_literal production.
	EnterSigned_numerical_literal(c *Signed_numerical_literalContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterCreate_schema_statement is called when entering the create_schema_statement production.
	EnterCreate_schema_statement(c *Create_schema_statementContext)

	// EnterCreate_policy_statement is called when entering the create_policy_statement production.
	EnterCreate_policy_statement(c *Create_policy_statementContext)

	// EnterAlter_policy_statement is called when entering the alter_policy_statement production.
	EnterAlter_policy_statement(c *Alter_policy_statementContext)

	// EnterDrop_policy_statement is called when entering the drop_policy_statement production.
	EnterDrop_policy_statement(c *Drop_policy_statementContext)

	// EnterCreate_subscription_statement is called when entering the create_subscription_statement production.
	EnterCreate_subscription_statement(c *Create_subscription_statementContext)

	// EnterAlter_subscription_statement is called when entering the alter_subscription_statement production.
	EnterAlter_subscription_statement(c *Alter_subscription_statementContext)

	// EnterAlter_subscription_action is called when entering the alter_subscription_action production.
	EnterAlter_subscription_action(c *Alter_subscription_actionContext)

	// EnterCreate_cast_statement is called when entering the create_cast_statement production.
	EnterCreate_cast_statement(c *Create_cast_statementContext)

	// EnterDrop_cast_statement is called when entering the drop_cast_statement production.
	EnterDrop_cast_statement(c *Drop_cast_statementContext)

	// EnterCreate_operator_family_statement is called when entering the create_operator_family_statement production.
	EnterCreate_operator_family_statement(c *Create_operator_family_statementContext)

	// EnterAlter_operator_family_statement is called when entering the alter_operator_family_statement production.
	EnterAlter_operator_family_statement(c *Alter_operator_family_statementContext)

	// EnterOperator_family_action is called when entering the operator_family_action production.
	EnterOperator_family_action(c *Operator_family_actionContext)

	// EnterAdd_operator_to_family is called when entering the add_operator_to_family production.
	EnterAdd_operator_to_family(c *Add_operator_to_familyContext)

	// EnterDrop_operator_from_family is called when entering the drop_operator_from_family production.
	EnterDrop_operator_from_family(c *Drop_operator_from_familyContext)

	// EnterDrop_operator_family_statement is called when entering the drop_operator_family_statement production.
	EnterDrop_operator_family_statement(c *Drop_operator_family_statementContext)

	// EnterCreate_operator_class_statement is called when entering the create_operator_class_statement production.
	EnterCreate_operator_class_statement(c *Create_operator_class_statementContext)

	// EnterCreate_operator_class_option is called when entering the create_operator_class_option production.
	EnterCreate_operator_class_option(c *Create_operator_class_optionContext)

	// EnterAlter_operator_class_statement is called when entering the alter_operator_class_statement production.
	EnterAlter_operator_class_statement(c *Alter_operator_class_statementContext)

	// EnterDrop_operator_class_statement is called when entering the drop_operator_class_statement production.
	EnterDrop_operator_class_statement(c *Drop_operator_class_statementContext)

	// EnterCreate_conversion_statement is called when entering the create_conversion_statement production.
	EnterCreate_conversion_statement(c *Create_conversion_statementContext)

	// EnterAlter_conversion_statement is called when entering the alter_conversion_statement production.
	EnterAlter_conversion_statement(c *Alter_conversion_statementContext)

	// EnterCreate_publication_statement is called when entering the create_publication_statement production.
	EnterCreate_publication_statement(c *Create_publication_statementContext)

	// EnterAlter_publication_statement is called when entering the alter_publication_statement production.
	EnterAlter_publication_statement(c *Alter_publication_statementContext)

	// EnterAlter_publication_action is called when entering the alter_publication_action production.
	EnterAlter_publication_action(c *Alter_publication_actionContext)

	// EnterOnly_table_multiply is called when entering the only_table_multiply production.
	EnterOnly_table_multiply(c *Only_table_multiplyContext)

	// EnterAlter_trigger_statement is called when entering the alter_trigger_statement production.
	EnterAlter_trigger_statement(c *Alter_trigger_statementContext)

	// EnterAlter_rule_statement is called when entering the alter_rule_statement production.
	EnterAlter_rule_statement(c *Alter_rule_statementContext)

	// EnterCopy_statement is called when entering the copy_statement production.
	EnterCopy_statement(c *Copy_statementContext)

	// EnterCopy_from_statement is called when entering the copy_from_statement production.
	EnterCopy_from_statement(c *Copy_from_statementContext)

	// EnterCopy_to_statement is called when entering the copy_to_statement production.
	EnterCopy_to_statement(c *Copy_to_statementContext)

	// EnterCopy_option_list is called when entering the copy_option_list production.
	EnterCopy_option_list(c *Copy_option_listContext)

	// EnterCopy_option is called when entering the copy_option production.
	EnterCopy_option(c *Copy_optionContext)

	// EnterCreate_view_statement is called when entering the create_view_statement production.
	EnterCreate_view_statement(c *Create_view_statementContext)

	// EnterIf_exists is called when entering the if_exists production.
	EnterIf_exists(c *If_existsContext)

	// EnterIf_not_exists is called when entering the if_not_exists production.
	EnterIf_not_exists(c *If_not_existsContext)

	// EnterView_columns is called when entering the view_columns production.
	EnterView_columns(c *View_columnsContext)

	// EnterWith_check_option is called when entering the with_check_option production.
	EnterWith_check_option(c *With_check_optionContext)

	// EnterCreate_database_statement is called when entering the create_database_statement production.
	EnterCreate_database_statement(c *Create_database_statementContext)

	// EnterCreate_database_option is called when entering the create_database_option production.
	EnterCreate_database_option(c *Create_database_optionContext)

	// EnterAlter_database_statement is called when entering the alter_database_statement production.
	EnterAlter_database_statement(c *Alter_database_statementContext)

	// EnterAlter_database_action is called when entering the alter_database_action production.
	EnterAlter_database_action(c *Alter_database_actionContext)

	// EnterAlter_database_option is called when entering the alter_database_option production.
	EnterAlter_database_option(c *Alter_database_optionContext)

	// EnterCreate_table_statement is called when entering the create_table_statement production.
	EnterCreate_table_statement(c *Create_table_statementContext)

	// EnterCreate_table_as_statement is called when entering the create_table_as_statement production.
	EnterCreate_table_as_statement(c *Create_table_as_statementContext)

	// EnterCreate_foreign_table_statement is called when entering the create_foreign_table_statement production.
	EnterCreate_foreign_table_statement(c *Create_foreign_table_statementContext)

	// EnterDefine_table is called when entering the define_table production.
	EnterDefine_table(c *Define_tableContext)

	// EnterDefine_partition is called when entering the define_partition production.
	EnterDefine_partition(c *Define_partitionContext)

	// EnterFor_values_bound is called when entering the for_values_bound production.
	EnterFor_values_bound(c *For_values_boundContext)

	// EnterPartition_bound_spec is called when entering the partition_bound_spec production.
	EnterPartition_bound_spec(c *Partition_bound_specContext)

	// EnterPartition_bound_part is called when entering the partition_bound_part production.
	EnterPartition_bound_part(c *Partition_bound_partContext)

	// EnterDefine_columns is called when entering the define_columns production.
	EnterDefine_columns(c *Define_columnsContext)

	// EnterDefine_type is called when entering the define_type production.
	EnterDefine_type(c *Define_typeContext)

	// EnterPartition_by is called when entering the partition_by production.
	EnterPartition_by(c *Partition_byContext)

	// EnterPartition_method is called when entering the partition_method production.
	EnterPartition_method(c *Partition_methodContext)

	// EnterPartition_column is called when entering the partition_column production.
	EnterPartition_column(c *Partition_columnContext)

	// EnterDefine_server is called when entering the define_server production.
	EnterDefine_server(c *Define_serverContext)

	// EnterDefine_foreign_options is called when entering the define_foreign_options production.
	EnterDefine_foreign_options(c *Define_foreign_optionsContext)

	// EnterForeign_option is called when entering the foreign_option production.
	EnterForeign_option(c *Foreign_optionContext)

	// EnterForeign_option_name is called when entering the foreign_option_name production.
	EnterForeign_option_name(c *Foreign_option_nameContext)

	// EnterList_of_type_column_def is called when entering the list_of_type_column_def production.
	EnterList_of_type_column_def(c *List_of_type_column_defContext)

	// EnterTable_column_def is called when entering the table_column_def production.
	EnterTable_column_def(c *Table_column_defContext)

	// EnterTable_of_type_column_def is called when entering the table_of_type_column_def production.
	EnterTable_of_type_column_def(c *Table_of_type_column_defContext)

	// EnterTable_column_definition is called when entering the table_column_definition production.
	EnterTable_column_definition(c *Table_column_definitionContext)

	// EnterLike_option is called when entering the like_option production.
	EnterLike_option(c *Like_optionContext)

	// EnterConstraint_common is called when entering the constraint_common production.
	EnterConstraint_common(c *Constraint_commonContext)

	// EnterConstr_body is called when entering the constr_body production.
	EnterConstr_body(c *Constr_bodyContext)

	// EnterAll_op is called when entering the all_op production.
	EnterAll_op(c *All_opContext)

	// EnterAll_simple_op is called when entering the all_simple_op production.
	EnterAll_simple_op(c *All_simple_opContext)

	// EnterOp_chars is called when entering the op_chars production.
	EnterOp_chars(c *Op_charsContext)

	// EnterIndex_parameters is called when entering the index_parameters production.
	EnterIndex_parameters(c *Index_parametersContext)

	// EnterNames_in_parens is called when entering the names_in_parens production.
	EnterNames_in_parens(c *Names_in_parensContext)

	// EnterNames_references is called when entering the names_references production.
	EnterNames_references(c *Names_referencesContext)

	// EnterStorage_parameter is called when entering the storage_parameter production.
	EnterStorage_parameter(c *Storage_parameterContext)

	// EnterStorage_parameter_option is called when entering the storage_parameter_option production.
	EnterStorage_parameter_option(c *Storage_parameter_optionContext)

	// EnterStorage_parameter_name is called when entering the storage_parameter_name production.
	EnterStorage_parameter_name(c *Storage_parameter_nameContext)

	// EnterWith_storage_parameter is called when entering the with_storage_parameter production.
	EnterWith_storage_parameter(c *With_storage_parameterContext)

	// EnterStorage_parameter_oid is called when entering the storage_parameter_oid production.
	EnterStorage_parameter_oid(c *Storage_parameter_oidContext)

	// EnterOn_commit is called when entering the on_commit production.
	EnterOn_commit(c *On_commitContext)

	// EnterTable_space is called when entering the table_space production.
	EnterTable_space(c *Table_spaceContext)

	// EnterSet_tablespace is called when entering the set_tablespace production.
	EnterSet_tablespace(c *Set_tablespaceContext)

	// EnterX_action is called when entering the x_action production.
	EnterX_action(c *X_actionContext)

	// EnterOwner_to is called when entering the owner_to production.
	EnterOwner_to(c *Owner_toContext)

	// EnterRename_to is called when entering the rename_to production.
	EnterRename_to(c *Rename_toContext)

	// EnterSet_schema is called when entering the set_schema production.
	EnterSet_schema(c *Set_schemaContext)

	// EnterTable_column_privilege is called when entering the table_column_privilege production.
	EnterTable_column_privilege(c *Table_column_privilegeContext)

	// EnterUsage_select_update is called when entering the usage_select_update production.
	EnterUsage_select_update(c *Usage_select_updateContext)

	// EnterPartition_by_columns is called when entering the partition_by_columns production.
	EnterPartition_by_columns(c *Partition_by_columnsContext)

	// EnterCascade_restrict is called when entering the cascade_restrict production.
	EnterCascade_restrict(c *Cascade_restrictContext)

	// EnterCollate_identifier is called when entering the collate_identifier production.
	EnterCollate_identifier(c *Collate_identifierContext)

	// EnterIndirection_var is called when entering the indirection_var production.
	EnterIndirection_var(c *Indirection_varContext)

	// EnterDollar_number is called when entering the dollar_number production.
	EnterDollar_number(c *Dollar_numberContext)

	// EnterIndirection_list is called when entering the indirection_list production.
	EnterIndirection_list(c *Indirection_listContext)

	// EnterIndirection is called when entering the indirection production.
	EnterIndirection(c *IndirectionContext)

	// EnterDrop_database_statement is called when entering the drop_database_statement production.
	EnterDrop_database_statement(c *Drop_database_statementContext)

	// EnterDrop_function_statement is called when entering the drop_function_statement production.
	EnterDrop_function_statement(c *Drop_function_statementContext)

	// EnterDrop_trigger_statement is called when entering the drop_trigger_statement production.
	EnterDrop_trigger_statement(c *Drop_trigger_statementContext)

	// EnterDrop_rule_statement is called when entering the drop_rule_statement production.
	EnterDrop_rule_statement(c *Drop_rule_statementContext)

	// EnterDrop_statements is called when entering the drop_statements production.
	EnterDrop_statements(c *Drop_statementsContext)

	// EnterIf_exist_names_restrict_cascade is called when entering the if_exist_names_restrict_cascade production.
	EnterIf_exist_names_restrict_cascade(c *If_exist_names_restrict_cascadeContext)

	// EnterId_token is called when entering the id_token production.
	EnterId_token(c *Id_tokenContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifier_nontype is called when entering the identifier_nontype production.
	EnterIdentifier_nontype(c *Identifier_nontypeContext)

	// EnterCol_label is called when entering the col_label production.
	EnterCol_label(c *Col_labelContext)

	// EnterTokens_nonreserved is called when entering the tokens_nonreserved production.
	EnterTokens_nonreserved(c *Tokens_nonreservedContext)

	// EnterTokens_nonreserved_except_function_type is called when entering the tokens_nonreserved_except_function_type production.
	EnterTokens_nonreserved_except_function_type(c *Tokens_nonreserved_except_function_typeContext)

	// EnterTokens_reserved_except_function_type is called when entering the tokens_reserved_except_function_type production.
	EnterTokens_reserved_except_function_type(c *Tokens_reserved_except_function_typeContext)

	// EnterTokens_reserved is called when entering the tokens_reserved production.
	EnterTokens_reserved(c *Tokens_reservedContext)

	// EnterTokens_nonkeyword is called when entering the tokens_nonkeyword production.
	EnterTokens_nonkeyword(c *Tokens_nonkeywordContext)

	// EnterSchema_qualified_name_nontype is called when entering the schema_qualified_name_nontype production.
	EnterSchema_qualified_name_nontype(c *Schema_qualified_name_nontypeContext)

	// EnterType_list is called when entering the type_list production.
	EnterType_list(c *Type_listContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterArray_type is called when entering the array_type production.
	EnterArray_type(c *Array_typeContext)

	// EnterPredefined_type is called when entering the predefined_type production.
	EnterPredefined_type(c *Predefined_typeContext)

	// EnterInterval_field is called when entering the interval_field production.
	EnterInterval_field(c *Interval_fieldContext)

	// EnterType_length is called when entering the type_length production.
	EnterType_length(c *Type_lengthContext)

	// EnterPrecision_param is called when entering the precision_param production.
	EnterPrecision_param(c *Precision_paramContext)

	// EnterVex is called when entering the vex production.
	EnterVex(c *VexContext)

	// EnterVex_b is called when entering the vex_b production.
	EnterVex_b(c *Vex_bContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterAll_op_ref is called when entering the all_op_ref production.
	EnterAll_op_ref(c *All_op_refContext)

	// EnterDatetime_overlaps is called when entering the datetime_overlaps production.
	EnterDatetime_overlaps(c *Datetime_overlapsContext)

	// EnterValue_expression_primary is called when entering the value_expression_primary production.
	EnterValue_expression_primary(c *Value_expression_primaryContext)

	// EnterUnsigned_value_specification is called when entering the unsigned_value_specification production.
	EnterUnsigned_value_specification(c *Unsigned_value_specificationContext)

	// EnterUnsigned_numeric_literal is called when entering the unsigned_numeric_literal production.
	EnterUnsigned_numeric_literal(c *Unsigned_numeric_literalContext)

	// EnterTruth_value is called when entering the truth_value production.
	EnterTruth_value(c *Truth_valueContext)

	// EnterCase_expression is called when entering the case_expression production.
	EnterCase_expression(c *Case_expressionContext)

	// EnterCast_specification is called when entering the cast_specification production.
	EnterCast_specification(c *Cast_specificationContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterVex_or_named_notation is called when entering the vex_or_named_notation production.
	EnterVex_or_named_notation(c *Vex_or_named_notationContext)

	// EnterPointer is called when entering the pointer production.
	EnterPointer(c *PointerContext)

	// EnterFunction_construct is called when entering the function_construct production.
	EnterFunction_construct(c *Function_constructContext)

	// EnterExtract_function is called when entering the extract_function production.
	EnterExtract_function(c *Extract_functionContext)

	// EnterSystem_function is called when entering the system_function production.
	EnterSystem_function(c *System_functionContext)

	// EnterDate_time_function is called when entering the date_time_function production.
	EnterDate_time_function(c *Date_time_functionContext)

	// EnterString_value_function is called when entering the string_value_function production.
	EnterString_value_function(c *String_value_functionContext)

	// EnterXml_function is called when entering the xml_function production.
	EnterXml_function(c *Xml_functionContext)

	// EnterXml_table_column is called when entering the xml_table_column production.
	EnterXml_table_column(c *Xml_table_columnContext)

	// EnterComparison_mod is called when entering the comparison_mod production.
	EnterComparison_mod(c *Comparison_modContext)

	// EnterFilter_clause is called when entering the filter_clause production.
	EnterFilter_clause(c *Filter_clauseContext)

	// EnterWindow_definition is called when entering the window_definition production.
	EnterWindow_definition(c *Window_definitionContext)

	// EnterFrame_clause is called when entering the frame_clause production.
	EnterFrame_clause(c *Frame_clauseContext)

	// EnterFrame_bound is called when entering the frame_bound production.
	EnterFrame_bound(c *Frame_boundContext)

	// EnterArray_expression is called when entering the array_expression production.
	EnterArray_expression(c *Array_expressionContext)

	// EnterArray_elements is called when entering the array_elements production.
	EnterArray_elements(c *Array_elementsContext)

	// EnterType_coercion is called when entering the type_coercion production.
	EnterType_coercion(c *Type_coercionContext)

	// EnterSchema_qualified_name is called when entering the schema_qualified_name production.
	EnterSchema_qualified_name(c *Schema_qualified_nameContext)

	// EnterSet_qualifier is called when entering the set_qualifier production.
	EnterSet_qualifier(c *Set_qualifierContext)

	// EnterTable_subquery is called when entering the table_subquery production.
	EnterTable_subquery(c *Table_subqueryContext)

	// EnterSelect_stmt is called when entering the select_stmt production.
	EnterSelect_stmt(c *Select_stmtContext)

	// EnterAfter_ops is called when entering the after_ops production.
	EnterAfter_ops(c *After_opsContext)

	// EnterSelect_stmt_no_parens is called when entering the select_stmt_no_parens production.
	EnterSelect_stmt_no_parens(c *Select_stmt_no_parensContext)

	// EnterWith_clause is called when entering the with_clause production.
	EnterWith_clause(c *With_clauseContext)

	// EnterWith_query is called when entering the with_query production.
	EnterWith_query(c *With_queryContext)

	// EnterSelect_ops is called when entering the select_ops production.
	EnterSelect_ops(c *Select_opsContext)

	// EnterSelect_ops_no_parens is called when entering the select_ops_no_parens production.
	EnterSelect_ops_no_parens(c *Select_ops_no_parensContext)

	// EnterSelect_primary is called when entering the select_primary production.
	EnterSelect_primary(c *Select_primaryContext)

	// EnterSelect_list is called when entering the select_list production.
	EnterSelect_list(c *Select_listContext)

	// EnterSelect_sublist is called when entering the select_sublist production.
	EnterSelect_sublist(c *Select_sublistContext)

	// EnterInto_table is called when entering the into_table production.
	EnterInto_table(c *Into_tableContext)

	// EnterFrom_item is called when entering the from_item production.
	EnterFrom_item(c *From_itemContext)

	// EnterFrom_primary is called when entering the from_primary production.
	EnterFrom_primary(c *From_primaryContext)

	// EnterAlias_clause is called when entering the alias_clause production.
	EnterAlias_clause(c *Alias_clauseContext)

	// EnterFrom_function_column_def is called when entering the from_function_column_def production.
	EnterFrom_function_column_def(c *From_function_column_defContext)

	// EnterGroupby_clause is called when entering the groupby_clause production.
	EnterGroupby_clause(c *Groupby_clauseContext)

	// EnterGrouping_element_list is called when entering the grouping_element_list production.
	EnterGrouping_element_list(c *Grouping_element_listContext)

	// EnterGrouping_element is called when entering the grouping_element production.
	EnterGrouping_element(c *Grouping_elementContext)

	// EnterValues_stmt is called when entering the values_stmt production.
	EnterValues_stmt(c *Values_stmtContext)

	// EnterValues_values is called when entering the values_values production.
	EnterValues_values(c *Values_valuesContext)

	// EnterOrderby_clause is called when entering the orderby_clause production.
	EnterOrderby_clause(c *Orderby_clauseContext)

	// EnterSort_specifier is called when entering the sort_specifier production.
	EnterSort_specifier(c *Sort_specifierContext)

	// EnterOrder_specification is called when entering the order_specification production.
	EnterOrder_specification(c *Order_specificationContext)

	// EnterNull_ordering is called when entering the null_ordering production.
	EnterNull_ordering(c *Null_orderingContext)

	// EnterInsert_stmt_for_psql is called when entering the insert_stmt_for_psql production.
	EnterInsert_stmt_for_psql(c *Insert_stmt_for_psqlContext)

	// EnterInsert_columns is called when entering the insert_columns production.
	EnterInsert_columns(c *Insert_columnsContext)

	// EnterIndirection_identifier is called when entering the indirection_identifier production.
	EnterIndirection_identifier(c *Indirection_identifierContext)

	// EnterConflict_object is called when entering the conflict_object production.
	EnterConflict_object(c *Conflict_objectContext)

	// EnterConflict_action is called when entering the conflict_action production.
	EnterConflict_action(c *Conflict_actionContext)

	// EnterDelete_stmt_for_psql is called when entering the delete_stmt_for_psql production.
	EnterDelete_stmt_for_psql(c *Delete_stmt_for_psqlContext)

	// EnterUpdate_stmt_for_psql is called when entering the update_stmt_for_psql production.
	EnterUpdate_stmt_for_psql(c *Update_stmt_for_psqlContext)

	// EnterUpdate_set is called when entering the update_set production.
	EnterUpdate_set(c *Update_setContext)

	// EnterNotify_stmt is called when entering the notify_stmt production.
	EnterNotify_stmt(c *Notify_stmtContext)

	// EnterTruncate_stmt is called when entering the truncate_stmt production.
	EnterTruncate_stmt(c *Truncate_stmtContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterAnonymous_block is called when entering the anonymous_block production.
	EnterAnonymous_block(c *Anonymous_blockContext)

	// EnterComp_options is called when entering the comp_options production.
	EnterComp_options(c *Comp_optionsContext)

	// EnterFunction_block is called when entering the function_block production.
	EnterFunction_block(c *Function_blockContext)

	// EnterStart_label is called when entering the start_label production.
	EnterStart_label(c *Start_labelContext)

	// EnterDeclarations is called when entering the declarations production.
	EnterDeclarations(c *DeclarationsContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterType_declaration is called when entering the type_declaration production.
	EnterType_declaration(c *Type_declarationContext)

	// EnterArguments_list is called when entering the arguments_list production.
	EnterArguments_list(c *Arguments_listContext)

	// EnterData_type_dec is called when entering the data_type_dec production.
	EnterData_type_dec(c *Data_type_decContext)

	// EnterException_statement is called when entering the exception_statement production.
	EnterException_statement(c *Exception_statementContext)

	// EnterFunction_statements is called when entering the function_statements production.
	EnterFunction_statements(c *Function_statementsContext)

	// EnterFunction_statement is called when entering the function_statement production.
	EnterFunction_statement(c *Function_statementContext)

	// EnterBase_statement is called when entering the base_statement production.
	EnterBase_statement(c *Base_statementContext)

	// EnterX_var is called when entering the x_var production.
	EnterX_var(c *X_varContext)

	// EnterDiagnostic_option is called when entering the diagnostic_option production.
	EnterDiagnostic_option(c *Diagnostic_optionContext)

	// EnterPerform_stmt is called when entering the perform_stmt production.
	EnterPerform_stmt(c *Perform_stmtContext)

	// EnterAssign_stmt is called when entering the assign_stmt production.
	EnterAssign_stmt(c *Assign_stmtContext)

	// EnterExecute_stmt is called when entering the execute_stmt production.
	EnterExecute_stmt(c *Execute_stmtContext)

	// EnterControl_statement is called when entering the control_statement production.
	EnterControl_statement(c *Control_statementContext)

	// EnterCursor_statement is called when entering the cursor_statement production.
	EnterCursor_statement(c *Cursor_statementContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterTransaction_statement is called when entering the transaction_statement production.
	EnterTransaction_statement(c *Transaction_statementContext)

	// EnterMessage_statement is called when entering the message_statement production.
	EnterMessage_statement(c *Message_statementContext)

	// EnterLog_level is called when entering the log_level production.
	EnterLog_level(c *Log_levelContext)

	// EnterRaise_using is called when entering the raise_using production.
	EnterRaise_using(c *Raise_usingContext)

	// EnterRaise_param is called when entering the raise_param production.
	EnterRaise_param(c *Raise_paramContext)

	// EnterReturn_stmt is called when entering the return_stmt production.
	EnterReturn_stmt(c *Return_stmtContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterLoop_start is called when entering the loop_start production.
	EnterLoop_start(c *Loop_startContext)

	// EnterUsing_vex is called when entering the using_vex production.
	EnterUsing_vex(c *Using_vexContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterPlpgsql_query is called when entering the plpgsql_query production.
	EnterPlpgsql_query(c *Plpgsql_queryContext)

	// ExitSql is called when exiting the sql production.
	ExitSql(c *SqlContext)

	// ExitQname_parser is called when exiting the qname_parser production.
	ExitQname_parser(c *Qname_parserContext)

	// ExitFunction_args_parser is called when exiting the function_args_parser production.
	ExitFunction_args_parser(c *Function_args_parserContext)

	// ExitVex_eof is called when exiting the vex_eof production.
	ExitVex_eof(c *Vex_eofContext)

	// ExitPlpgsql_function is called when exiting the plpgsql_function production.
	ExitPlpgsql_function(c *Plpgsql_functionContext)

	// ExitPlpgsql_function_test_list is called when exiting the plpgsql_function_test_list production.
	ExitPlpgsql_function_test_list(c *Plpgsql_function_test_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitData_statement is called when exiting the data_statement production.
	ExitData_statement(c *Data_statementContext)

	// ExitScript_statement is called when exiting the script_statement production.
	ExitScript_statement(c *Script_statementContext)

	// ExitScript_transaction is called when exiting the script_transaction production.
	ExitScript_transaction(c *Script_transactionContext)

	// ExitTransaction_mode is called when exiting the transaction_mode production.
	ExitTransaction_mode(c *Transaction_modeContext)

	// ExitLock_table is called when exiting the lock_table production.
	ExitLock_table(c *Lock_tableContext)

	// ExitLock_mode is called when exiting the lock_mode production.
	ExitLock_mode(c *Lock_modeContext)

	// ExitScript_additional is called when exiting the script_additional production.
	ExitScript_additional(c *Script_additionalContext)

	// ExitAdditional_statement is called when exiting the additional_statement production.
	ExitAdditional_statement(c *Additional_statementContext)

	// ExitExplain_statement is called when exiting the explain_statement production.
	ExitExplain_statement(c *Explain_statementContext)

	// ExitExplain_query is called when exiting the explain_query production.
	ExitExplain_query(c *Explain_queryContext)

	// ExitExecute_statement is called when exiting the execute_statement production.
	ExitExecute_statement(c *Execute_statementContext)

	// ExitDeclare_statement is called when exiting the declare_statement production.
	ExitDeclare_statement(c *Declare_statementContext)

	// ExitShow_statement is called when exiting the show_statement production.
	ExitShow_statement(c *Show_statementContext)

	// ExitExplain_option is called when exiting the explain_option production.
	ExitExplain_option(c *Explain_optionContext)

	// ExitUser_name is called when exiting the user_name production.
	ExitUser_name(c *User_nameContext)

	// ExitTable_cols_list is called when exiting the table_cols_list production.
	ExitTable_cols_list(c *Table_cols_listContext)

	// ExitTable_cols is called when exiting the table_cols production.
	ExitTable_cols(c *Table_colsContext)

	// ExitVacuum_mode is called when exiting the vacuum_mode production.
	ExitVacuum_mode(c *Vacuum_modeContext)

	// ExitVacuum_option is called when exiting the vacuum_option production.
	ExitVacuum_option(c *Vacuum_optionContext)

	// ExitAnalyze_mode is called when exiting the analyze_mode production.
	ExitAnalyze_mode(c *Analyze_modeContext)

	// ExitBoolean_value is called when exiting the boolean_value production.
	ExitBoolean_value(c *Boolean_valueContext)

	// ExitFetch_move_direction is called when exiting the fetch_move_direction production.
	ExitFetch_move_direction(c *Fetch_move_directionContext)

	// ExitSchema_statement is called when exiting the schema_statement production.
	ExitSchema_statement(c *Schema_statementContext)

	// ExitSchema_create is called when exiting the schema_create production.
	ExitSchema_create(c *Schema_createContext)

	// ExitSchema_alter is called when exiting the schema_alter production.
	ExitSchema_alter(c *Schema_alterContext)

	// ExitSchema_drop is called when exiting the schema_drop production.
	ExitSchema_drop(c *Schema_dropContext)

	// ExitSchema_import is called when exiting the schema_import production.
	ExitSchema_import(c *Schema_importContext)

	// ExitAlter_function_statement is called when exiting the alter_function_statement production.
	ExitAlter_function_statement(c *Alter_function_statementContext)

	// ExitAlter_aggregate_statement is called when exiting the alter_aggregate_statement production.
	ExitAlter_aggregate_statement(c *Alter_aggregate_statementContext)

	// ExitAlter_extension_statement is called when exiting the alter_extension_statement production.
	ExitAlter_extension_statement(c *Alter_extension_statementContext)

	// ExitAlter_extension_action is called when exiting the alter_extension_action production.
	ExitAlter_extension_action(c *Alter_extension_actionContext)

	// ExitExtension_member_object is called when exiting the extension_member_object production.
	ExitExtension_member_object(c *Extension_member_objectContext)

	// ExitAlter_schema_statement is called when exiting the alter_schema_statement production.
	ExitAlter_schema_statement(c *Alter_schema_statementContext)

	// ExitAlter_language_statement is called when exiting the alter_language_statement production.
	ExitAlter_language_statement(c *Alter_language_statementContext)

	// ExitAlter_table_statement is called when exiting the alter_table_statement production.
	ExitAlter_table_statement(c *Alter_table_statementContext)

	// ExitTable_action is called when exiting the table_action production.
	ExitTable_action(c *Table_actionContext)

	// ExitColumn_action is called when exiting the column_action production.
	ExitColumn_action(c *Column_actionContext)

	// ExitIdentity_body is called when exiting the identity_body production.
	ExitIdentity_body(c *Identity_bodyContext)

	// ExitAlter_identity is called when exiting the alter_identity production.
	ExitAlter_identity(c *Alter_identityContext)

	// ExitStorage_option is called when exiting the storage_option production.
	ExitStorage_option(c *Storage_optionContext)

	// ExitValidate_constraint is called when exiting the validate_constraint production.
	ExitValidate_constraint(c *Validate_constraintContext)

	// ExitDrop_constraint is called when exiting the drop_constraint production.
	ExitDrop_constraint(c *Drop_constraintContext)

	// ExitTable_deferrable is called when exiting the table_deferrable production.
	ExitTable_deferrable(c *Table_deferrableContext)

	// ExitTable_initialy_immed is called when exiting the table_initialy_immed production.
	ExitTable_initialy_immed(c *Table_initialy_immedContext)

	// ExitFunction_actions_common is called when exiting the function_actions_common production.
	ExitFunction_actions_common(c *Function_actions_commonContext)

	// ExitFunction_def is called when exiting the function_def production.
	ExitFunction_def(c *Function_defContext)

	// ExitAlter_index_statement is called when exiting the alter_index_statement production.
	ExitAlter_index_statement(c *Alter_index_statementContext)

	// ExitIndex_def_action is called when exiting the index_def_action production.
	ExitIndex_def_action(c *Index_def_actionContext)

	// ExitAlter_default_privileges_statement is called when exiting the alter_default_privileges_statement production.
	ExitAlter_default_privileges_statement(c *Alter_default_privileges_statementContext)

	// ExitAbbreviated_grant_or_revoke is called when exiting the abbreviated_grant_or_revoke production.
	ExitAbbreviated_grant_or_revoke(c *Abbreviated_grant_or_revokeContext)

	// ExitGrant_option_for is called when exiting the grant_option_for production.
	ExitGrant_option_for(c *Grant_option_forContext)

	// ExitAlter_sequence_statement is called when exiting the alter_sequence_statement production.
	ExitAlter_sequence_statement(c *Alter_sequence_statementContext)

	// ExitAlter_view_statement is called when exiting the alter_view_statement production.
	ExitAlter_view_statement(c *Alter_view_statementContext)

	// ExitAlter_view_action is called when exiting the alter_view_action production.
	ExitAlter_view_action(c *Alter_view_actionContext)

	// ExitAlter_materialized_view_statement is called when exiting the alter_materialized_view_statement production.
	ExitAlter_materialized_view_statement(c *Alter_materialized_view_statementContext)

	// ExitAlter_materialized_view_action is called when exiting the alter_materialized_view_action production.
	ExitAlter_materialized_view_action(c *Alter_materialized_view_actionContext)

	// ExitMaterialized_view_action is called when exiting the materialized_view_action production.
	ExitMaterialized_view_action(c *Materialized_view_actionContext)

	// ExitAlter_event_trigger_statement is called when exiting the alter_event_trigger_statement production.
	ExitAlter_event_trigger_statement(c *Alter_event_trigger_statementContext)

	// ExitAlter_event_trigger_action is called when exiting the alter_event_trigger_action production.
	ExitAlter_event_trigger_action(c *Alter_event_trigger_actionContext)

	// ExitAlter_type_statement is called when exiting the alter_type_statement production.
	ExitAlter_type_statement(c *Alter_type_statementContext)

	// ExitAlter_domain_statement is called when exiting the alter_domain_statement production.
	ExitAlter_domain_statement(c *Alter_domain_statementContext)

	// ExitAlter_server_statement is called when exiting the alter_server_statement production.
	ExitAlter_server_statement(c *Alter_server_statementContext)

	// ExitAlter_server_action is called when exiting the alter_server_action production.
	ExitAlter_server_action(c *Alter_server_actionContext)

	// ExitAlter_fts_statement is called when exiting the alter_fts_statement production.
	ExitAlter_fts_statement(c *Alter_fts_statementContext)

	// ExitAlter_fts_configuration is called when exiting the alter_fts_configuration production.
	ExitAlter_fts_configuration(c *Alter_fts_configurationContext)

	// ExitType_action is called when exiting the type_action production.
	ExitType_action(c *Type_actionContext)

	// ExitType_property is called when exiting the type_property production.
	ExitType_property(c *Type_propertyContext)

	// ExitSet_def_column is called when exiting the set_def_column production.
	ExitSet_def_column(c *Set_def_columnContext)

	// ExitDrop_def is called when exiting the drop_def production.
	ExitDrop_def(c *Drop_defContext)

	// ExitCreate_index_statement is called when exiting the create_index_statement production.
	ExitCreate_index_statement(c *Create_index_statementContext)

	// ExitIndex_rest is called when exiting the index_rest production.
	ExitIndex_rest(c *Index_restContext)

	// ExitIndex_sort is called when exiting the index_sort production.
	ExitIndex_sort(c *Index_sortContext)

	// ExitIndex_column is called when exiting the index_column production.
	ExitIndex_column(c *Index_columnContext)

	// ExitIncluding_index is called when exiting the including_index production.
	ExitIncluding_index(c *Including_indexContext)

	// ExitIndex_where is called when exiting the index_where production.
	ExitIndex_where(c *Index_whereContext)

	// ExitCreate_extension_statement is called when exiting the create_extension_statement production.
	ExitCreate_extension_statement(c *Create_extension_statementContext)

	// ExitCreate_language_statement is called when exiting the create_language_statement production.
	ExitCreate_language_statement(c *Create_language_statementContext)

	// ExitCreate_event_trigger_statement is called when exiting the create_event_trigger_statement production.
	ExitCreate_event_trigger_statement(c *Create_event_trigger_statementContext)

	// ExitCreate_type_statement is called when exiting the create_type_statement production.
	ExitCreate_type_statement(c *Create_type_statementContext)

	// ExitCreate_domain_statement is called when exiting the create_domain_statement production.
	ExitCreate_domain_statement(c *Create_domain_statementContext)

	// ExitCreate_server_statement is called when exiting the create_server_statement production.
	ExitCreate_server_statement(c *Create_server_statementContext)

	// ExitCreate_fts_dictionary_statement is called when exiting the create_fts_dictionary_statement production.
	ExitCreate_fts_dictionary_statement(c *Create_fts_dictionary_statementContext)

	// ExitOption_with_value is called when exiting the option_with_value production.
	ExitOption_with_value(c *Option_with_valueContext)

	// ExitCreate_fts_configuration_statement is called when exiting the create_fts_configuration_statement production.
	ExitCreate_fts_configuration_statement(c *Create_fts_configuration_statementContext)

	// ExitCreate_fts_template_statement is called when exiting the create_fts_template_statement production.
	ExitCreate_fts_template_statement(c *Create_fts_template_statementContext)

	// ExitCreate_fts_parser_statement is called when exiting the create_fts_parser_statement production.
	ExitCreate_fts_parser_statement(c *Create_fts_parser_statementContext)

	// ExitCreate_collation_statement is called when exiting the create_collation_statement production.
	ExitCreate_collation_statement(c *Create_collation_statementContext)

	// ExitAlter_collation_statement is called when exiting the alter_collation_statement production.
	ExitAlter_collation_statement(c *Alter_collation_statementContext)

	// ExitCollation_option is called when exiting the collation_option production.
	ExitCollation_option(c *Collation_optionContext)

	// ExitCreate_user_mapping_statement is called when exiting the create_user_mapping_statement production.
	ExitCreate_user_mapping_statement(c *Create_user_mapping_statementContext)

	// ExitAlter_user_mapping_statement is called when exiting the alter_user_mapping_statement production.
	ExitAlter_user_mapping_statement(c *Alter_user_mapping_statementContext)

	// ExitAlter_user_or_role_statement is called when exiting the alter_user_or_role_statement production.
	ExitAlter_user_or_role_statement(c *Alter_user_or_role_statementContext)

	// ExitAlter_user_or_role_set_reset is called when exiting the alter_user_or_role_set_reset production.
	ExitAlter_user_or_role_set_reset(c *Alter_user_or_role_set_resetContext)

	// ExitSet_reset_parameter is called when exiting the set_reset_parameter production.
	ExitSet_reset_parameter(c *Set_reset_parameterContext)

	// ExitAlter_group_statement is called when exiting the alter_group_statement production.
	ExitAlter_group_statement(c *Alter_group_statementContext)

	// ExitAlter_group_action is called when exiting the alter_group_action production.
	ExitAlter_group_action(c *Alter_group_actionContext)

	// ExitAlter_tablespace_statement is called when exiting the alter_tablespace_statement production.
	ExitAlter_tablespace_statement(c *Alter_tablespace_statementContext)

	// ExitAlter_owner_statement is called when exiting the alter_owner_statement production.
	ExitAlter_owner_statement(c *Alter_owner_statementContext)

	// ExitAlter_tablespace_action is called when exiting the alter_tablespace_action production.
	ExitAlter_tablespace_action(c *Alter_tablespace_actionContext)

	// ExitAlter_statistics_statement is called when exiting the alter_statistics_statement production.
	ExitAlter_statistics_statement(c *Alter_statistics_statementContext)

	// ExitSet_statistics is called when exiting the set_statistics production.
	ExitSet_statistics(c *Set_statisticsContext)

	// ExitAlter_foreign_data_wrapper is called when exiting the alter_foreign_data_wrapper production.
	ExitAlter_foreign_data_wrapper(c *Alter_foreign_data_wrapperContext)

	// ExitAlter_foreign_data_wrapper_action is called when exiting the alter_foreign_data_wrapper_action production.
	ExitAlter_foreign_data_wrapper_action(c *Alter_foreign_data_wrapper_actionContext)

	// ExitAlter_operator_statement is called when exiting the alter_operator_statement production.
	ExitAlter_operator_statement(c *Alter_operator_statementContext)

	// ExitAlter_operator_action is called when exiting the alter_operator_action production.
	ExitAlter_operator_action(c *Alter_operator_actionContext)

	// ExitOperator_set_restrict_join is called when exiting the operator_set_restrict_join production.
	ExitOperator_set_restrict_join(c *Operator_set_restrict_joinContext)

	// ExitDrop_user_mapping_statement is called when exiting the drop_user_mapping_statement production.
	ExitDrop_user_mapping_statement(c *Drop_user_mapping_statementContext)

	// ExitDrop_owned_statement is called when exiting the drop_owned_statement production.
	ExitDrop_owned_statement(c *Drop_owned_statementContext)

	// ExitDrop_operator_statement is called when exiting the drop_operator_statement production.
	ExitDrop_operator_statement(c *Drop_operator_statementContext)

	// ExitTarget_operator is called when exiting the target_operator production.
	ExitTarget_operator(c *Target_operatorContext)

	// ExitDomain_constraint is called when exiting the domain_constraint production.
	ExitDomain_constraint(c *Domain_constraintContext)

	// ExitCreate_transform_statement is called when exiting the create_transform_statement production.
	ExitCreate_transform_statement(c *Create_transform_statementContext)

	// ExitCreate_access_method_statement is called when exiting the create_access_method_statement production.
	ExitCreate_access_method_statement(c *Create_access_method_statementContext)

	// ExitCreate_user_or_role_statement is called when exiting the create_user_or_role_statement production.
	ExitCreate_user_or_role_statement(c *Create_user_or_role_statementContext)

	// ExitUser_or_role_option is called when exiting the user_or_role_option production.
	ExitUser_or_role_option(c *User_or_role_optionContext)

	// ExitUser_or_role_option_for_alter is called when exiting the user_or_role_option_for_alter production.
	ExitUser_or_role_option_for_alter(c *User_or_role_option_for_alterContext)

	// ExitUser_or_role_or_group_common_option is called when exiting the user_or_role_or_group_common_option production.
	ExitUser_or_role_or_group_common_option(c *User_or_role_or_group_common_optionContext)

	// ExitUser_or_role_common_option is called when exiting the user_or_role_common_option production.
	ExitUser_or_role_common_option(c *User_or_role_common_optionContext)

	// ExitUser_or_role_or_group_option_for_create is called when exiting the user_or_role_or_group_option_for_create production.
	ExitUser_or_role_or_group_option_for_create(c *User_or_role_or_group_option_for_createContext)

	// ExitCreate_group_statement is called when exiting the create_group_statement production.
	ExitCreate_group_statement(c *Create_group_statementContext)

	// ExitGroup_option is called when exiting the group_option production.
	ExitGroup_option(c *Group_optionContext)

	// ExitCreate_tablespace_statement is called when exiting the create_tablespace_statement production.
	ExitCreate_tablespace_statement(c *Create_tablespace_statementContext)

	// ExitCreate_statistics_statement is called when exiting the create_statistics_statement production.
	ExitCreate_statistics_statement(c *Create_statistics_statementContext)

	// ExitCreate_foreign_data_wrapper_statement is called when exiting the create_foreign_data_wrapper_statement production.
	ExitCreate_foreign_data_wrapper_statement(c *Create_foreign_data_wrapper_statementContext)

	// ExitOption_without_equal is called when exiting the option_without_equal production.
	ExitOption_without_equal(c *Option_without_equalContext)

	// ExitCreate_operator_statement is called when exiting the create_operator_statement production.
	ExitCreate_operator_statement(c *Create_operator_statementContext)

	// ExitOperator_name is called when exiting the operator_name production.
	ExitOperator_name(c *Operator_nameContext)

	// ExitOperator_option is called when exiting the operator_option production.
	ExitOperator_option(c *Operator_optionContext)

	// ExitCreate_aggregate_statement is called when exiting the create_aggregate_statement production.
	ExitCreate_aggregate_statement(c *Create_aggregate_statementContext)

	// ExitAggregate_param is called when exiting the aggregate_param production.
	ExitAggregate_param(c *Aggregate_paramContext)

	// ExitSet_statement is called when exiting the set_statement production.
	ExitSet_statement(c *Set_statementContext)

	// ExitSet_action is called when exiting the set_action production.
	ExitSet_action(c *Set_actionContext)

	// ExitSession_local_option is called when exiting the session_local_option production.
	ExitSession_local_option(c *Session_local_optionContext)

	// ExitSet_statement_value is called when exiting the set_statement_value production.
	ExitSet_statement_value(c *Set_statement_valueContext)

	// ExitCreate_rewrite_statement is called when exiting the create_rewrite_statement production.
	ExitCreate_rewrite_statement(c *Create_rewrite_statementContext)

	// ExitRewrite_command is called when exiting the rewrite_command production.
	ExitRewrite_command(c *Rewrite_commandContext)

	// ExitCreate_trigger_statement is called when exiting the create_trigger_statement production.
	ExitCreate_trigger_statement(c *Create_trigger_statementContext)

	// ExitTrigger_referencing is called when exiting the trigger_referencing production.
	ExitTrigger_referencing(c *Trigger_referencingContext)

	// ExitWhen_trigger is called when exiting the when_trigger production.
	ExitWhen_trigger(c *When_triggerContext)

	// ExitRule_common is called when exiting the rule_common production.
	ExitRule_common(c *Rule_commonContext)

	// ExitRule_member_object is called when exiting the rule_member_object production.
	ExitRule_member_object(c *Rule_member_objectContext)

	// ExitColumns_permissions is called when exiting the columns_permissions production.
	ExitColumns_permissions(c *Columns_permissionsContext)

	// ExitTable_column_privileges is called when exiting the table_column_privileges production.
	ExitTable_column_privileges(c *Table_column_privilegesContext)

	// ExitPermissions is called when exiting the permissions production.
	ExitPermissions(c *PermissionsContext)

	// ExitPermission is called when exiting the permission production.
	ExitPermission(c *PermissionContext)

	// ExitOther_rules is called when exiting the other_rules production.
	ExitOther_rules(c *Other_rulesContext)

	// ExitGrant_to_rule is called when exiting the grant_to_rule production.
	ExitGrant_to_rule(c *Grant_to_ruleContext)

	// ExitRevoke_from_cascade_restrict is called when exiting the revoke_from_cascade_restrict production.
	ExitRevoke_from_cascade_restrict(c *Revoke_from_cascade_restrictContext)

	// ExitRoles_names is called when exiting the roles_names production.
	ExitRoles_names(c *Roles_namesContext)

	// ExitRole_name_with_group is called when exiting the role_name_with_group production.
	ExitRole_name_with_group(c *Role_name_with_groupContext)

	// ExitComment_on_statement is called when exiting the comment_on_statement production.
	ExitComment_on_statement(c *Comment_on_statementContext)

	// ExitSecurity_label is called when exiting the security_label production.
	ExitSecurity_label(c *Security_labelContext)

	// ExitComment_member_object is called when exiting the comment_member_object production.
	ExitComment_member_object(c *Comment_member_objectContext)

	// ExitLabel_member_object is called when exiting the label_member_object production.
	ExitLabel_member_object(c *Label_member_objectContext)

	// ExitCreate_function_statement is called when exiting the create_function_statement production.
	ExitCreate_function_statement(c *Create_function_statementContext)

	// ExitCreate_funct_params is called when exiting the create_funct_params production.
	ExitCreate_funct_params(c *Create_funct_paramsContext)

	// ExitTransform_for_type is called when exiting the transform_for_type production.
	ExitTransform_for_type(c *Transform_for_typeContext)

	// ExitFunction_ret_table is called when exiting the function_ret_table production.
	ExitFunction_ret_table(c *Function_ret_tableContext)

	// ExitFunction_column_name_type is called when exiting the function_column_name_type production.
	ExitFunction_column_name_type(c *Function_column_name_typeContext)

	// ExitFunction_parameters is called when exiting the function_parameters production.
	ExitFunction_parameters(c *Function_parametersContext)

	// ExitFunction_args is called when exiting the function_args production.
	ExitFunction_args(c *Function_argsContext)

	// ExitAgg_order is called when exiting the agg_order production.
	ExitAgg_order(c *Agg_orderContext)

	// ExitCharacter_string is called when exiting the character_string production.
	ExitCharacter_string(c *Character_stringContext)

	// ExitFunction_arguments is called when exiting the function_arguments production.
	ExitFunction_arguments(c *Function_argumentsContext)

	// ExitArgmode is called when exiting the argmode production.
	ExitArgmode(c *ArgmodeContext)

	// ExitCreate_sequence_statement is called when exiting the create_sequence_statement production.
	ExitCreate_sequence_statement(c *Create_sequence_statementContext)

	// ExitSequence_body is called when exiting the sequence_body production.
	ExitSequence_body(c *Sequence_bodyContext)

	// ExitSigned_number_literal is called when exiting the signed_number_literal production.
	ExitSigned_number_literal(c *Signed_number_literalContext)

	// ExitSigned_numerical_literal is called when exiting the signed_numerical_literal production.
	ExitSigned_numerical_literal(c *Signed_numerical_literalContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitCreate_schema_statement is called when exiting the create_schema_statement production.
	ExitCreate_schema_statement(c *Create_schema_statementContext)

	// ExitCreate_policy_statement is called when exiting the create_policy_statement production.
	ExitCreate_policy_statement(c *Create_policy_statementContext)

	// ExitAlter_policy_statement is called when exiting the alter_policy_statement production.
	ExitAlter_policy_statement(c *Alter_policy_statementContext)

	// ExitDrop_policy_statement is called when exiting the drop_policy_statement production.
	ExitDrop_policy_statement(c *Drop_policy_statementContext)

	// ExitCreate_subscription_statement is called when exiting the create_subscription_statement production.
	ExitCreate_subscription_statement(c *Create_subscription_statementContext)

	// ExitAlter_subscription_statement is called when exiting the alter_subscription_statement production.
	ExitAlter_subscription_statement(c *Alter_subscription_statementContext)

	// ExitAlter_subscription_action is called when exiting the alter_subscription_action production.
	ExitAlter_subscription_action(c *Alter_subscription_actionContext)

	// ExitCreate_cast_statement is called when exiting the create_cast_statement production.
	ExitCreate_cast_statement(c *Create_cast_statementContext)

	// ExitDrop_cast_statement is called when exiting the drop_cast_statement production.
	ExitDrop_cast_statement(c *Drop_cast_statementContext)

	// ExitCreate_operator_family_statement is called when exiting the create_operator_family_statement production.
	ExitCreate_operator_family_statement(c *Create_operator_family_statementContext)

	// ExitAlter_operator_family_statement is called when exiting the alter_operator_family_statement production.
	ExitAlter_operator_family_statement(c *Alter_operator_family_statementContext)

	// ExitOperator_family_action is called when exiting the operator_family_action production.
	ExitOperator_family_action(c *Operator_family_actionContext)

	// ExitAdd_operator_to_family is called when exiting the add_operator_to_family production.
	ExitAdd_operator_to_family(c *Add_operator_to_familyContext)

	// ExitDrop_operator_from_family is called when exiting the drop_operator_from_family production.
	ExitDrop_operator_from_family(c *Drop_operator_from_familyContext)

	// ExitDrop_operator_family_statement is called when exiting the drop_operator_family_statement production.
	ExitDrop_operator_family_statement(c *Drop_operator_family_statementContext)

	// ExitCreate_operator_class_statement is called when exiting the create_operator_class_statement production.
	ExitCreate_operator_class_statement(c *Create_operator_class_statementContext)

	// ExitCreate_operator_class_option is called when exiting the create_operator_class_option production.
	ExitCreate_operator_class_option(c *Create_operator_class_optionContext)

	// ExitAlter_operator_class_statement is called when exiting the alter_operator_class_statement production.
	ExitAlter_operator_class_statement(c *Alter_operator_class_statementContext)

	// ExitDrop_operator_class_statement is called when exiting the drop_operator_class_statement production.
	ExitDrop_operator_class_statement(c *Drop_operator_class_statementContext)

	// ExitCreate_conversion_statement is called when exiting the create_conversion_statement production.
	ExitCreate_conversion_statement(c *Create_conversion_statementContext)

	// ExitAlter_conversion_statement is called when exiting the alter_conversion_statement production.
	ExitAlter_conversion_statement(c *Alter_conversion_statementContext)

	// ExitCreate_publication_statement is called when exiting the create_publication_statement production.
	ExitCreate_publication_statement(c *Create_publication_statementContext)

	// ExitAlter_publication_statement is called when exiting the alter_publication_statement production.
	ExitAlter_publication_statement(c *Alter_publication_statementContext)

	// ExitAlter_publication_action is called when exiting the alter_publication_action production.
	ExitAlter_publication_action(c *Alter_publication_actionContext)

	// ExitOnly_table_multiply is called when exiting the only_table_multiply production.
	ExitOnly_table_multiply(c *Only_table_multiplyContext)

	// ExitAlter_trigger_statement is called when exiting the alter_trigger_statement production.
	ExitAlter_trigger_statement(c *Alter_trigger_statementContext)

	// ExitAlter_rule_statement is called when exiting the alter_rule_statement production.
	ExitAlter_rule_statement(c *Alter_rule_statementContext)

	// ExitCopy_statement is called when exiting the copy_statement production.
	ExitCopy_statement(c *Copy_statementContext)

	// ExitCopy_from_statement is called when exiting the copy_from_statement production.
	ExitCopy_from_statement(c *Copy_from_statementContext)

	// ExitCopy_to_statement is called when exiting the copy_to_statement production.
	ExitCopy_to_statement(c *Copy_to_statementContext)

	// ExitCopy_option_list is called when exiting the copy_option_list production.
	ExitCopy_option_list(c *Copy_option_listContext)

	// ExitCopy_option is called when exiting the copy_option production.
	ExitCopy_option(c *Copy_optionContext)

	// ExitCreate_view_statement is called when exiting the create_view_statement production.
	ExitCreate_view_statement(c *Create_view_statementContext)

	// ExitIf_exists is called when exiting the if_exists production.
	ExitIf_exists(c *If_existsContext)

	// ExitIf_not_exists is called when exiting the if_not_exists production.
	ExitIf_not_exists(c *If_not_existsContext)

	// ExitView_columns is called when exiting the view_columns production.
	ExitView_columns(c *View_columnsContext)

	// ExitWith_check_option is called when exiting the with_check_option production.
	ExitWith_check_option(c *With_check_optionContext)

	// ExitCreate_database_statement is called when exiting the create_database_statement production.
	ExitCreate_database_statement(c *Create_database_statementContext)

	// ExitCreate_database_option is called when exiting the create_database_option production.
	ExitCreate_database_option(c *Create_database_optionContext)

	// ExitAlter_database_statement is called when exiting the alter_database_statement production.
	ExitAlter_database_statement(c *Alter_database_statementContext)

	// ExitAlter_database_action is called when exiting the alter_database_action production.
	ExitAlter_database_action(c *Alter_database_actionContext)

	// ExitAlter_database_option is called when exiting the alter_database_option production.
	ExitAlter_database_option(c *Alter_database_optionContext)

	// ExitCreate_table_statement is called when exiting the create_table_statement production.
	ExitCreate_table_statement(c *Create_table_statementContext)

	// ExitCreate_table_as_statement is called when exiting the create_table_as_statement production.
	ExitCreate_table_as_statement(c *Create_table_as_statementContext)

	// ExitCreate_foreign_table_statement is called when exiting the create_foreign_table_statement production.
	ExitCreate_foreign_table_statement(c *Create_foreign_table_statementContext)

	// ExitDefine_table is called when exiting the define_table production.
	ExitDefine_table(c *Define_tableContext)

	// ExitDefine_partition is called when exiting the define_partition production.
	ExitDefine_partition(c *Define_partitionContext)

	// ExitFor_values_bound is called when exiting the for_values_bound production.
	ExitFor_values_bound(c *For_values_boundContext)

	// ExitPartition_bound_spec is called when exiting the partition_bound_spec production.
	ExitPartition_bound_spec(c *Partition_bound_specContext)

	// ExitPartition_bound_part is called when exiting the partition_bound_part production.
	ExitPartition_bound_part(c *Partition_bound_partContext)

	// ExitDefine_columns is called when exiting the define_columns production.
	ExitDefine_columns(c *Define_columnsContext)

	// ExitDefine_type is called when exiting the define_type production.
	ExitDefine_type(c *Define_typeContext)

	// ExitPartition_by is called when exiting the partition_by production.
	ExitPartition_by(c *Partition_byContext)

	// ExitPartition_method is called when exiting the partition_method production.
	ExitPartition_method(c *Partition_methodContext)

	// ExitPartition_column is called when exiting the partition_column production.
	ExitPartition_column(c *Partition_columnContext)

	// ExitDefine_server is called when exiting the define_server production.
	ExitDefine_server(c *Define_serverContext)

	// ExitDefine_foreign_options is called when exiting the define_foreign_options production.
	ExitDefine_foreign_options(c *Define_foreign_optionsContext)

	// ExitForeign_option is called when exiting the foreign_option production.
	ExitForeign_option(c *Foreign_optionContext)

	// ExitForeign_option_name is called when exiting the foreign_option_name production.
	ExitForeign_option_name(c *Foreign_option_nameContext)

	// ExitList_of_type_column_def is called when exiting the list_of_type_column_def production.
	ExitList_of_type_column_def(c *List_of_type_column_defContext)

	// ExitTable_column_def is called when exiting the table_column_def production.
	ExitTable_column_def(c *Table_column_defContext)

	// ExitTable_of_type_column_def is called when exiting the table_of_type_column_def production.
	ExitTable_of_type_column_def(c *Table_of_type_column_defContext)

	// ExitTable_column_definition is called when exiting the table_column_definition production.
	ExitTable_column_definition(c *Table_column_definitionContext)

	// ExitLike_option is called when exiting the like_option production.
	ExitLike_option(c *Like_optionContext)

	// ExitConstraint_common is called when exiting the constraint_common production.
	ExitConstraint_common(c *Constraint_commonContext)

	// ExitConstr_body is called when exiting the constr_body production.
	ExitConstr_body(c *Constr_bodyContext)

	// ExitAll_op is called when exiting the all_op production.
	ExitAll_op(c *All_opContext)

	// ExitAll_simple_op is called when exiting the all_simple_op production.
	ExitAll_simple_op(c *All_simple_opContext)

	// ExitOp_chars is called when exiting the op_chars production.
	ExitOp_chars(c *Op_charsContext)

	// ExitIndex_parameters is called when exiting the index_parameters production.
	ExitIndex_parameters(c *Index_parametersContext)

	// ExitNames_in_parens is called when exiting the names_in_parens production.
	ExitNames_in_parens(c *Names_in_parensContext)

	// ExitNames_references is called when exiting the names_references production.
	ExitNames_references(c *Names_referencesContext)

	// ExitStorage_parameter is called when exiting the storage_parameter production.
	ExitStorage_parameter(c *Storage_parameterContext)

	// ExitStorage_parameter_option is called when exiting the storage_parameter_option production.
	ExitStorage_parameter_option(c *Storage_parameter_optionContext)

	// ExitStorage_parameter_name is called when exiting the storage_parameter_name production.
	ExitStorage_parameter_name(c *Storage_parameter_nameContext)

	// ExitWith_storage_parameter is called when exiting the with_storage_parameter production.
	ExitWith_storage_parameter(c *With_storage_parameterContext)

	// ExitStorage_parameter_oid is called when exiting the storage_parameter_oid production.
	ExitStorage_parameter_oid(c *Storage_parameter_oidContext)

	// ExitOn_commit is called when exiting the on_commit production.
	ExitOn_commit(c *On_commitContext)

	// ExitTable_space is called when exiting the table_space production.
	ExitTable_space(c *Table_spaceContext)

	// ExitSet_tablespace is called when exiting the set_tablespace production.
	ExitSet_tablespace(c *Set_tablespaceContext)

	// ExitX_action is called when exiting the x_action production.
	ExitX_action(c *X_actionContext)

	// ExitOwner_to is called when exiting the owner_to production.
	ExitOwner_to(c *Owner_toContext)

	// ExitRename_to is called when exiting the rename_to production.
	ExitRename_to(c *Rename_toContext)

	// ExitSet_schema is called when exiting the set_schema production.
	ExitSet_schema(c *Set_schemaContext)

	// ExitTable_column_privilege is called when exiting the table_column_privilege production.
	ExitTable_column_privilege(c *Table_column_privilegeContext)

	// ExitUsage_select_update is called when exiting the usage_select_update production.
	ExitUsage_select_update(c *Usage_select_updateContext)

	// ExitPartition_by_columns is called when exiting the partition_by_columns production.
	ExitPartition_by_columns(c *Partition_by_columnsContext)

	// ExitCascade_restrict is called when exiting the cascade_restrict production.
	ExitCascade_restrict(c *Cascade_restrictContext)

	// ExitCollate_identifier is called when exiting the collate_identifier production.
	ExitCollate_identifier(c *Collate_identifierContext)

	// ExitIndirection_var is called when exiting the indirection_var production.
	ExitIndirection_var(c *Indirection_varContext)

	// ExitDollar_number is called when exiting the dollar_number production.
	ExitDollar_number(c *Dollar_numberContext)

	// ExitIndirection_list is called when exiting the indirection_list production.
	ExitIndirection_list(c *Indirection_listContext)

	// ExitIndirection is called when exiting the indirection production.
	ExitIndirection(c *IndirectionContext)

	// ExitDrop_database_statement is called when exiting the drop_database_statement production.
	ExitDrop_database_statement(c *Drop_database_statementContext)

	// ExitDrop_function_statement is called when exiting the drop_function_statement production.
	ExitDrop_function_statement(c *Drop_function_statementContext)

	// ExitDrop_trigger_statement is called when exiting the drop_trigger_statement production.
	ExitDrop_trigger_statement(c *Drop_trigger_statementContext)

	// ExitDrop_rule_statement is called when exiting the drop_rule_statement production.
	ExitDrop_rule_statement(c *Drop_rule_statementContext)

	// ExitDrop_statements is called when exiting the drop_statements production.
	ExitDrop_statements(c *Drop_statementsContext)

	// ExitIf_exist_names_restrict_cascade is called when exiting the if_exist_names_restrict_cascade production.
	ExitIf_exist_names_restrict_cascade(c *If_exist_names_restrict_cascadeContext)

	// ExitId_token is called when exiting the id_token production.
	ExitId_token(c *Id_tokenContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifier_nontype is called when exiting the identifier_nontype production.
	ExitIdentifier_nontype(c *Identifier_nontypeContext)

	// ExitCol_label is called when exiting the col_label production.
	ExitCol_label(c *Col_labelContext)

	// ExitTokens_nonreserved is called when exiting the tokens_nonreserved production.
	ExitTokens_nonreserved(c *Tokens_nonreservedContext)

	// ExitTokens_nonreserved_except_function_type is called when exiting the tokens_nonreserved_except_function_type production.
	ExitTokens_nonreserved_except_function_type(c *Tokens_nonreserved_except_function_typeContext)

	// ExitTokens_reserved_except_function_type is called when exiting the tokens_reserved_except_function_type production.
	ExitTokens_reserved_except_function_type(c *Tokens_reserved_except_function_typeContext)

	// ExitTokens_reserved is called when exiting the tokens_reserved production.
	ExitTokens_reserved(c *Tokens_reservedContext)

	// ExitTokens_nonkeyword is called when exiting the tokens_nonkeyword production.
	ExitTokens_nonkeyword(c *Tokens_nonkeywordContext)

	// ExitSchema_qualified_name_nontype is called when exiting the schema_qualified_name_nontype production.
	ExitSchema_qualified_name_nontype(c *Schema_qualified_name_nontypeContext)

	// ExitType_list is called when exiting the type_list production.
	ExitType_list(c *Type_listContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitArray_type is called when exiting the array_type production.
	ExitArray_type(c *Array_typeContext)

	// ExitPredefined_type is called when exiting the predefined_type production.
	ExitPredefined_type(c *Predefined_typeContext)

	// ExitInterval_field is called when exiting the interval_field production.
	ExitInterval_field(c *Interval_fieldContext)

	// ExitType_length is called when exiting the type_length production.
	ExitType_length(c *Type_lengthContext)

	// ExitPrecision_param is called when exiting the precision_param production.
	ExitPrecision_param(c *Precision_paramContext)

	// ExitVex is called when exiting the vex production.
	ExitVex(c *VexContext)

	// ExitVex_b is called when exiting the vex_b production.
	ExitVex_b(c *Vex_bContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitAll_op_ref is called when exiting the all_op_ref production.
	ExitAll_op_ref(c *All_op_refContext)

	// ExitDatetime_overlaps is called when exiting the datetime_overlaps production.
	ExitDatetime_overlaps(c *Datetime_overlapsContext)

	// ExitValue_expression_primary is called when exiting the value_expression_primary production.
	ExitValue_expression_primary(c *Value_expression_primaryContext)

	// ExitUnsigned_value_specification is called when exiting the unsigned_value_specification production.
	ExitUnsigned_value_specification(c *Unsigned_value_specificationContext)

	// ExitUnsigned_numeric_literal is called when exiting the unsigned_numeric_literal production.
	ExitUnsigned_numeric_literal(c *Unsigned_numeric_literalContext)

	// ExitTruth_value is called when exiting the truth_value production.
	ExitTruth_value(c *Truth_valueContext)

	// ExitCase_expression is called when exiting the case_expression production.
	ExitCase_expression(c *Case_expressionContext)

	// ExitCast_specification is called when exiting the cast_specification production.
	ExitCast_specification(c *Cast_specificationContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitVex_or_named_notation is called when exiting the vex_or_named_notation production.
	ExitVex_or_named_notation(c *Vex_or_named_notationContext)

	// ExitPointer is called when exiting the pointer production.
	ExitPointer(c *PointerContext)

	// ExitFunction_construct is called when exiting the function_construct production.
	ExitFunction_construct(c *Function_constructContext)

	// ExitExtract_function is called when exiting the extract_function production.
	ExitExtract_function(c *Extract_functionContext)

	// ExitSystem_function is called when exiting the system_function production.
	ExitSystem_function(c *System_functionContext)

	// ExitDate_time_function is called when exiting the date_time_function production.
	ExitDate_time_function(c *Date_time_functionContext)

	// ExitString_value_function is called when exiting the string_value_function production.
	ExitString_value_function(c *String_value_functionContext)

	// ExitXml_function is called when exiting the xml_function production.
	ExitXml_function(c *Xml_functionContext)

	// ExitXml_table_column is called when exiting the xml_table_column production.
	ExitXml_table_column(c *Xml_table_columnContext)

	// ExitComparison_mod is called when exiting the comparison_mod production.
	ExitComparison_mod(c *Comparison_modContext)

	// ExitFilter_clause is called when exiting the filter_clause production.
	ExitFilter_clause(c *Filter_clauseContext)

	// ExitWindow_definition is called when exiting the window_definition production.
	ExitWindow_definition(c *Window_definitionContext)

	// ExitFrame_clause is called when exiting the frame_clause production.
	ExitFrame_clause(c *Frame_clauseContext)

	// ExitFrame_bound is called when exiting the frame_bound production.
	ExitFrame_bound(c *Frame_boundContext)

	// ExitArray_expression is called when exiting the array_expression production.
	ExitArray_expression(c *Array_expressionContext)

	// ExitArray_elements is called when exiting the array_elements production.
	ExitArray_elements(c *Array_elementsContext)

	// ExitType_coercion is called when exiting the type_coercion production.
	ExitType_coercion(c *Type_coercionContext)

	// ExitSchema_qualified_name is called when exiting the schema_qualified_name production.
	ExitSchema_qualified_name(c *Schema_qualified_nameContext)

	// ExitSet_qualifier is called when exiting the set_qualifier production.
	ExitSet_qualifier(c *Set_qualifierContext)

	// ExitTable_subquery is called when exiting the table_subquery production.
	ExitTable_subquery(c *Table_subqueryContext)

	// ExitSelect_stmt is called when exiting the select_stmt production.
	ExitSelect_stmt(c *Select_stmtContext)

	// ExitAfter_ops is called when exiting the after_ops production.
	ExitAfter_ops(c *After_opsContext)

	// ExitSelect_stmt_no_parens is called when exiting the select_stmt_no_parens production.
	ExitSelect_stmt_no_parens(c *Select_stmt_no_parensContext)

	// ExitWith_clause is called when exiting the with_clause production.
	ExitWith_clause(c *With_clauseContext)

	// ExitWith_query is called when exiting the with_query production.
	ExitWith_query(c *With_queryContext)

	// ExitSelect_ops is called when exiting the select_ops production.
	ExitSelect_ops(c *Select_opsContext)

	// ExitSelect_ops_no_parens is called when exiting the select_ops_no_parens production.
	ExitSelect_ops_no_parens(c *Select_ops_no_parensContext)

	// ExitSelect_primary is called when exiting the select_primary production.
	ExitSelect_primary(c *Select_primaryContext)

	// ExitSelect_list is called when exiting the select_list production.
	ExitSelect_list(c *Select_listContext)

	// ExitSelect_sublist is called when exiting the select_sublist production.
	ExitSelect_sublist(c *Select_sublistContext)

	// ExitInto_table is called when exiting the into_table production.
	ExitInto_table(c *Into_tableContext)

	// ExitFrom_item is called when exiting the from_item production.
	ExitFrom_item(c *From_itemContext)

	// ExitFrom_primary is called when exiting the from_primary production.
	ExitFrom_primary(c *From_primaryContext)

	// ExitAlias_clause is called when exiting the alias_clause production.
	ExitAlias_clause(c *Alias_clauseContext)

	// ExitFrom_function_column_def is called when exiting the from_function_column_def production.
	ExitFrom_function_column_def(c *From_function_column_defContext)

	// ExitGroupby_clause is called when exiting the groupby_clause production.
	ExitGroupby_clause(c *Groupby_clauseContext)

	// ExitGrouping_element_list is called when exiting the grouping_element_list production.
	ExitGrouping_element_list(c *Grouping_element_listContext)

	// ExitGrouping_element is called when exiting the grouping_element production.
	ExitGrouping_element(c *Grouping_elementContext)

	// ExitValues_stmt is called when exiting the values_stmt production.
	ExitValues_stmt(c *Values_stmtContext)

	// ExitValues_values is called when exiting the values_values production.
	ExitValues_values(c *Values_valuesContext)

	// ExitOrderby_clause is called when exiting the orderby_clause production.
	ExitOrderby_clause(c *Orderby_clauseContext)

	// ExitSort_specifier is called when exiting the sort_specifier production.
	ExitSort_specifier(c *Sort_specifierContext)

	// ExitOrder_specification is called when exiting the order_specification production.
	ExitOrder_specification(c *Order_specificationContext)

	// ExitNull_ordering is called when exiting the null_ordering production.
	ExitNull_ordering(c *Null_orderingContext)

	// ExitInsert_stmt_for_psql is called when exiting the insert_stmt_for_psql production.
	ExitInsert_stmt_for_psql(c *Insert_stmt_for_psqlContext)

	// ExitInsert_columns is called when exiting the insert_columns production.
	ExitInsert_columns(c *Insert_columnsContext)

	// ExitIndirection_identifier is called when exiting the indirection_identifier production.
	ExitIndirection_identifier(c *Indirection_identifierContext)

	// ExitConflict_object is called when exiting the conflict_object production.
	ExitConflict_object(c *Conflict_objectContext)

	// ExitConflict_action is called when exiting the conflict_action production.
	ExitConflict_action(c *Conflict_actionContext)

	// ExitDelete_stmt_for_psql is called when exiting the delete_stmt_for_psql production.
	ExitDelete_stmt_for_psql(c *Delete_stmt_for_psqlContext)

	// ExitUpdate_stmt_for_psql is called when exiting the update_stmt_for_psql production.
	ExitUpdate_stmt_for_psql(c *Update_stmt_for_psqlContext)

	// ExitUpdate_set is called when exiting the update_set production.
	ExitUpdate_set(c *Update_setContext)

	// ExitNotify_stmt is called when exiting the notify_stmt production.
	ExitNotify_stmt(c *Notify_stmtContext)

	// ExitTruncate_stmt is called when exiting the truncate_stmt production.
	ExitTruncate_stmt(c *Truncate_stmtContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitAnonymous_block is called when exiting the anonymous_block production.
	ExitAnonymous_block(c *Anonymous_blockContext)

	// ExitComp_options is called when exiting the comp_options production.
	ExitComp_options(c *Comp_optionsContext)

	// ExitFunction_block is called when exiting the function_block production.
	ExitFunction_block(c *Function_blockContext)

	// ExitStart_label is called when exiting the start_label production.
	ExitStart_label(c *Start_labelContext)

	// ExitDeclarations is called when exiting the declarations production.
	ExitDeclarations(c *DeclarationsContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitType_declaration is called when exiting the type_declaration production.
	ExitType_declaration(c *Type_declarationContext)

	// ExitArguments_list is called when exiting the arguments_list production.
	ExitArguments_list(c *Arguments_listContext)

	// ExitData_type_dec is called when exiting the data_type_dec production.
	ExitData_type_dec(c *Data_type_decContext)

	// ExitException_statement is called when exiting the exception_statement production.
	ExitException_statement(c *Exception_statementContext)

	// ExitFunction_statements is called when exiting the function_statements production.
	ExitFunction_statements(c *Function_statementsContext)

	// ExitFunction_statement is called when exiting the function_statement production.
	ExitFunction_statement(c *Function_statementContext)

	// ExitBase_statement is called when exiting the base_statement production.
	ExitBase_statement(c *Base_statementContext)

	// ExitX_var is called when exiting the x_var production.
	ExitX_var(c *X_varContext)

	// ExitDiagnostic_option is called when exiting the diagnostic_option production.
	ExitDiagnostic_option(c *Diagnostic_optionContext)

	// ExitPerform_stmt is called when exiting the perform_stmt production.
	ExitPerform_stmt(c *Perform_stmtContext)

	// ExitAssign_stmt is called when exiting the assign_stmt production.
	ExitAssign_stmt(c *Assign_stmtContext)

	// ExitExecute_stmt is called when exiting the execute_stmt production.
	ExitExecute_stmt(c *Execute_stmtContext)

	// ExitControl_statement is called when exiting the control_statement production.
	ExitControl_statement(c *Control_statementContext)

	// ExitCursor_statement is called when exiting the cursor_statement production.
	ExitCursor_statement(c *Cursor_statementContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitTransaction_statement is called when exiting the transaction_statement production.
	ExitTransaction_statement(c *Transaction_statementContext)

	// ExitMessage_statement is called when exiting the message_statement production.
	ExitMessage_statement(c *Message_statementContext)

	// ExitLog_level is called when exiting the log_level production.
	ExitLog_level(c *Log_levelContext)

	// ExitRaise_using is called when exiting the raise_using production.
	ExitRaise_using(c *Raise_usingContext)

	// ExitRaise_param is called when exiting the raise_param production.
	ExitRaise_param(c *Raise_paramContext)

	// ExitReturn_stmt is called when exiting the return_stmt production.
	ExitReturn_stmt(c *Return_stmtContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitLoop_start is called when exiting the loop_start production.
	ExitLoop_start(c *Loop_startContext)

	// ExitUsing_vex is called when exiting the using_vex production.
	ExitUsing_vex(c *Using_vexContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitPlpgsql_query is called when exiting the plpgsql_query production.
	ExitPlpgsql_query(c *Plpgsql_queryContext)
}
