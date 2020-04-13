# awsls

A list command for AWS. Supports listing of 219 resource types across
76 different services.

How can awsls cover so many resource types? The answer is that its code is generated; here is the 
[code of the generator](./gen). Feel free to fork it and generate something else.

## Usage

	awsls <resource_type>

Run, for example

    awsls aws_vpc

## Supported resource types

### accessanalyzer 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_accessanalyzer_analyzer |  x  | 

### acm 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_acm_certificate |  | 

### apigateway 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_api_gateway_client_certificate |  x  | 
| aws_api_gateway_vpc_link |  x  | 
| aws_api_gateway_api_key |  x  | 
| aws_api_gateway_usage_plan |  x  | 
| aws_api_gateway_domain_name |  x  | 
| aws_api_gateway_rest_api |  x  | 

### apigatewayv2 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_apigatewayv2_api |  x  | 

### appmesh 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_appmesh_mesh |  | 

### appsync 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_appsync_graphql_api |  x  | 

### athena 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_athena_workgroup |  |  x 

### autoscaling 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_autoscaling_group |  x  |  x 
| aws_launch_configuration |  |  x 

### backup 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_backup_vault |  |  x 
| aws_backup_plan |  |  x 

### batch 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_batch_job_queue |  | 
| aws_batch_job_definition |  | 
| aws_batch_compute_environment |  | 

### cloudformation 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_cloudformation_stack_set |  | 
| aws_cloudformation_stack |  x  |  x 

### cloudhsmv2 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_cloudhsm_v2_cluster |  | 

### cloudwatch 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_cloudwatch_dashboard |  | 

### cloudwatchevents 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_cloudwatch_event_rule |  | 

### cloudwatchlogs 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_cloudwatch_log_destination |  |  x 
| aws_cloudwatch_log_group |  |  x 
| aws_cloudwatch_log_resource_policy |  | 

### codebuild 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_codebuild_source_credential |  | 

### codecommit 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_codecommit_repository |  | 

### codepipeline 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_codepipeline_webhook |  x  | 

### codestarnotifications 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_codestarnotifications_notification_rule |  | 

### configservice 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_config_config_rule |  | 
| aws_config_delivery_channel |  | 
| aws_config_configuration_recorder |  | 

### costandusagereportservice 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_cur_report_definition |  | 

### databasemigrationservice 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_dms_replication_subnet_group |  | 
| aws_dms_endpoint |  | 
| aws_dms_certificate |  | 
| aws_dms_replication_task |  | 

### datasync 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_datasync_agent |  | 
| aws_datasync_task |  | 

### dax 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_dax_subnet_group |  | 
| aws_dax_parameter_group |  | 

### devicefarm 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_devicefarm_project |  | 

### directconnect 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_dx_hosted_transit_virtual_interface |  x  | 
| aws_dx_transit_virtual_interface |  x  | 
| aws_dx_connection |  x  | 
| aws_dx_hosted_public_virtual_interface |  x  | 
| aws_dx_hosted_private_virtual_interface |  x  | 
| aws_dx_public_virtual_interface |  x  | 
| aws_dx_lag |  x  | 
| aws_dx_private_virtual_interface |  x  | 

### dlm 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_dlm_lifecycle_policy |  x  | 

### dynamodb 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_dynamodb_global_table |  | 

### ec2 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_internet_gateway |  x  | 
| aws_vpc_endpoint_service |  x  | 
| aws_security_group |  x  | 
| aws_vpc_endpoint |  x  |  x 
| aws_ec2_transit_gateway_vpc_attachment |  x  |  x 
| aws_ec2_capacity_reservation |  x  |  x 
| aws_vpn_gateway |  x  | 
| aws_ami |  x  |  x 
| aws_spot_fleet_request |  |  x 
| aws_ebs_volume |  x  |  x 
| aws_spot_instance_request |  x  |  x 
| aws_ec2_traffic_mirror_target |  x  | 
| aws_eip |  x  | 
| aws_ec2_traffic_mirror_session |  x  | 
| aws_default_security_group |  x  | 
| aws_ec2_fleet |  x  |  x 
| aws_vpc_peering_connection |  x  | 
| aws_nat_gateway |  x  |  x 
| aws_key_pair |  x  | 
| aws_default_route_table |  x  | 
| aws_customer_gateway |  x  | 
| aws_vpc |  x  | 
| aws_subnet |  x  | 
| aws_ec2_transit_gateway |  x  |  x 
| aws_egress_only_internet_gateway |  x  | 
| aws_network_acl |  x  | 
| aws_route_table |  x  | 
| aws_network_interface |  | 
| aws_ec2_traffic_mirror_filter |  x  | 
| aws_ec2_transit_gateway_route_table |  x  |  x 
| aws_placement_group |  x  | 
| aws_vpc_endpoint_connection_notification |  | 
| aws_default_vpc |  x  | 
| aws_ec2_client_vpn_endpoint |  x  |  x 
| aws_ebs_snapshot |  x  |  x 
| aws_launch_template |  x  |  x 

### ecr 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_ecr_repository |  | 

### ecs 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_ecs_cluster |  x  | 

### efs 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_efs_mount_target |  | 
| aws_efs_file_system |  x  |  x 

### elasticache 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_elasticache_replication_group |  | 

### elasticbeanstalk 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_elastic_beanstalk_application_version |  | 
| aws_elastic_beanstalk_application |  | 
| aws_elastic_beanstalk_environment |  | 

### elasticloadbalancing 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_elb |  |  x 

### elasticloadbalancingv2 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_alb_listener_rule |  | 
| aws_alb_listener |  | 
| aws_lb_listener |  | 
| aws_alb_target_group |  | 
| aws_lb_listener_rule |  | 
| aws_lb_target_group |  | 

### elastictranscoder 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_elastictranscoder_preset |  | 
| aws_elastictranscoder_pipeline |  | 

### emr 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_emr_security_configuration |  | 

### fsx 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_fsx_windows_file_system |  x  |  x 
| aws_fsx_lustre_file_system |  x  |  x 

### gamelift 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_gamelift_build |  |  x 
| aws_gamelift_game_session_queue |  | 
| aws_gamelift_alias |  |  x 

### globalaccelerator 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_globalaccelerator_accelerator |  |  x 

### glue 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_glue_crawler |  |  x 
| aws_glue_job |  | 
| aws_glue_trigger |  | 
| aws_glue_security_configuration |  | 

### iam 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_iam_role |  x  |  x 
| aws_iam_policy |  |  x 
| aws_iam_user |  x  |  x 
| aws_iam_server_certificate |  | 
| aws_iam_access_key |  |  x 
| aws_iam_instance_profile |  |  x 
| aws_iam_group |  |  x 
| aws_iam_service_linked_role |  x  |  x 

### iot 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_iot_certificate |  |  x 
| aws_iot_thing |  | 
| aws_iot_topic_rule |  | 
| aws_iot_thing_type |  | 
| aws_iot_policy |  | 

### kafka 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_msk_cluster |  x  |  x 
| aws_msk_configuration |  |  x 

### kinesisanalytics 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_kinesis_analytics_application |  | 

### kms 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_kms_external_key |  | 
| aws_kms_alias |  | 
| aws_kms_key |  | 

### lambda 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_lambda_function |  | 
| aws_lambda_event_source_mapping |  | 

### licensemanager 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_licensemanager_license_configuration |  | 

### lightsail 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_lightsail_instance |  x  | 
| aws_lightsail_key_pair |  x  | 
| aws_lightsail_domain |  x  | 
| aws_lightsail_static_ip |  | 

### mediaconvert 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_media_convert_queue |  | 

### mediapackage 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_media_package_channel |  x  | 

### mediastore 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_media_store_container |  |  x 

### mq 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_mq_broker |  | 
| aws_mq_configuration |  x  | 

### neptune 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_neptune_event_subscription |  | 

### opsworks 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_opsworks_stack |  | 
| aws_opsworks_instance |  | 
| aws_opsworks_user_profile |  | 

### qldb 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_qldb_ledger |  | 

### rds 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_db_parameter_group |  | 
| aws_db_subnet_group |  | 
| aws_db_instance |  |  x 
| aws_db_event_subscription |  | 
| aws_rds_global_cluster |  | 
| aws_db_security_group |  | 

### redshift 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_redshift_snapshot_schedule |  x  | 
| aws_redshift_event_subscription |  x  | 
| aws_redshift_snapshot_copy_grant |  x  | 
| aws_redshift_cluster |  x  | 

### route53 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_route53_health_check |  | 
| aws_route53_zone |  | 

### route53resolver 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_route53_resolver_rule_association |  | 
| aws_route53_resolver_endpoint |  |  x 
| aws_route53_resolver_rule |  | 

### s3 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_s3_bucket |  |  x 

### sagemaker 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_sagemaker_model |  |  x 
| aws_sagemaker_endpoint |  |  x 

### secretsmanager 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_secretsmanager_secret |  x  | 

### servicecatalog 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_servicecatalog_portfolio |  |  x 

### servicediscovery 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_service_discovery_service |  |  x 

### ses 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_ses_template |  | 
| aws_ses_configuration_set |  | 
| aws_ses_receipt_rule_set |  | 
| aws_ses_active_receipt_rule_set |  | 
| aws_ses_receipt_filter |  | 

### sfn 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_sfn_state_machine |  |  x 
| aws_sfn_activity |  |  x 

### sns 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_sns_topic_subscription |  | 
| aws_sns_platform_application |  | 
| aws_sns_topic |  | 

### ssm 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_ssm_maintenance_window |  | 
| aws_ssm_activation |  x  | 
| aws_ssm_association |  | 
| aws_ssm_document |  x  | 
| aws_ssm_patch_group |  | 
| aws_ssm_patch_baseline |  | 

### storagegateway 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_storagegateway_gateway |  | 

### transfer 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_transfer_server |  | 

### waf 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_waf_byte_match_set |  | 
| aws_waf_ipset |  | 
| aws_waf_rule_group |  | 
| aws_waf_rule |  | 
| aws_waf_web_acl |  | 
| aws_waf_xss_match_set |  | 
| aws_waf_regex_match_set |  | 
| aws_waf_geo_match_set |  | 
| aws_waf_sql_injection_match_set |  | 
| aws_waf_regex_pattern_set |  | 
| aws_waf_rate_based_rule |  | 
| aws_waf_size_constraint_set |  | 

### wafregional 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_wafregional_geo_match_set |  | 
| aws_wafregional_xss_match_set |  | 
| aws_wafregional_regex_match_set |  | 
| aws_wafregional_rate_based_rule |  | 
| aws_wafregional_size_constraint_set |  | 
| aws_wafregional_web_acl |  | 
| aws_wafregional_ipset |  | 
| aws_wafregional_rule |  | 
| aws_wafregional_byte_match_set |  | 
| aws_wafregional_rule_group |  | 
| aws_wafregional_regex_pattern_set |  | 
| aws_wafregional_sql_injection_match_set |  | 

### worklink 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_worklink_fleet |  |  x 

### workspaces 

| Resource Type                    | Tags | Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_workspaces_ip_group |  |