# awsls

**Note: this tool is still WIP**

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

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_accessanalyzer_analyzer |  x  |  | 

### acm 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_acm_certificate |  |  | 

### apigateway 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_api_gateway_api_key |  x  |  | 
| aws_api_gateway_rest_api |  x  |  | 
| aws_api_gateway_domain_name |  x  |  | 
| aws_api_gateway_usage_plan |  x  |  | 
| aws_api_gateway_vpc_link |  x  |  | 
| aws_api_gateway_client_certificate |  x  |  | 

### apigatewayv2 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_apigatewayv2_api |  x  |  | 

### appmesh 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_appmesh_mesh |  |  | 

### appsync 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_appsync_graphql_api |  x  |  | 

### athena 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_athena_workgroup |  |  x  |  x 

### autoscaling 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_launch_configuration |  |  x  |  x 
| aws_autoscaling_group |  x  |  x  |  x 

### backup 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_backup_plan |  |  x  |  x 
| aws_backup_vault |  |  x  |  x 

### batch 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_batch_job_queue |  |  | 
| aws_batch_compute_environment |  |  | 
| aws_batch_job_definition |  |  | 

### cloudformation 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_cloudformation_stack |  x  |  x  |  x 
| aws_cloudformation_stack_set |  |  | 

### cloudhsmv2 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_cloudhsm_v2_cluster |  |  | 

### cloudwatch 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_cloudwatch_dashboard |  |  | 

### cloudwatchevents 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_cloudwatch_event_rule |  |  | 

### cloudwatchlogs 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_cloudwatch_log_destination |  |  x  |  x 
| aws_cloudwatch_log_group |  |  x  |  x 
| aws_cloudwatch_log_resource_policy |  |  | 

### codebuild 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_codebuild_source_credential |  |  | 

### codecommit 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_codecommit_repository |  |  | 

### codepipeline 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_codepipeline_webhook |  x  |  | 

### codestarnotifications 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_codestarnotifications_notification_rule |  |  | 

### configservice 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_config_delivery_channel |  |  | 
| aws_config_configuration_recorder |  |  | 
| aws_config_config_rule |  |  | 

### costandusagereportservice 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_cur_report_definition |  |  | 

### databasemigrationservice 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_dms_certificate |  |  | 
| aws_dms_endpoint |  |  | 
| aws_dms_replication_subnet_group |  |  | 
| aws_dms_replication_task |  |  | 

### datasync 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_datasync_agent |  |  | 
| aws_datasync_task |  |  | 

### dax 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_dax_subnet_group |  |  | 
| aws_dax_parameter_group |  |  | 

### devicefarm 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_devicefarm_project |  |  | 

### directconnect 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_dx_connection |  x  |  | 
| aws_dx_public_virtual_interface |  x  |  | 
| aws_dx_hosted_public_virtual_interface |  x  |  | 
| aws_dx_transit_virtual_interface |  x  |  | 
| aws_dx_hosted_private_virtual_interface |  x  |  | 
| aws_dx_lag |  x  |  | 
| aws_dx_private_virtual_interface |  x  |  | 
| aws_dx_hosted_transit_virtual_interface |  x  |  | 

### dlm 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_dlm_lifecycle_policy |  x  |  | 

### dynamodb 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_dynamodb_global_table |  |  | 

### ec2 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_default_route_table |  x  |  | 
| aws_security_group |  x  |  | 
| aws_vpn_gateway |  x  |  | 
| aws_ec2_transit_gateway_vpc_attachment |  x  |  x  |  x 
| aws_vpc_endpoint_connection_notification |  |  | 
| aws_nat_gateway |  x  |  x  |  x 
| aws_egress_only_internet_gateway |  x  |  | 
| aws_key_pair |  x  |  | 
| aws_vpc_peering_connection |  x  |  | 
| aws_ec2_traffic_mirror_session |  x  |  | 
| aws_ami |  x  |  x  |  x 
| aws_vpc |  x  |  | 
| aws_customer_gateway |  x  |  | 
| aws_vpc_endpoint |  x  |  x  |  x 
| aws_subnet |  x  |  | 
| aws_placement_group |  x  |  | 
| aws_ec2_client_vpn_endpoint |  x  |  x  |  x 
| aws_ec2_fleet |  x  |  x  |  x 
| aws_ec2_traffic_mirror_target |  x  |  | 
| aws_internet_gateway |  x  |  | 
| aws_default_vpc |  x  |  | 
| aws_launch_template |  x  |  x  |  x 
| aws_spot_fleet_request |  |  x  |  x 
| aws_default_security_group |  x  |  | 
| aws_ec2_capacity_reservation |  x  |  x  |  x 
| aws_eip |  x  |  | 
| aws_spot_instance_request |  x  |  x  |  x 
| aws_ebs_snapshot |  x  |  x  |  x 
| aws_ec2_traffic_mirror_filter |  x  |  | 
| aws_ec2_transit_gateway_route_table |  x  |  x  |  x 
| aws_vpc_endpoint_service |  x  |  | 
| aws_ec2_transit_gateway |  x  |  x  |  x 
| aws_network_interface |  |  | 
| aws_network_acl |  x  |  | 
| aws_ebs_volume |  x  |  x  |  x 
| aws_route_table |  x  |  | 

### ecr 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_ecr_repository |  |  | 

### ecs 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_ecs_cluster |  x  |  | 

### efs 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_efs_mount_target |  |  | 
| aws_efs_file_system |  x  |  x  |  x 

### elasticache 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_elasticache_replication_group |  |  | 

### elasticbeanstalk 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_elastic_beanstalk_application |  |  | 
| aws_elastic_beanstalk_environment |  |  | 
| aws_elastic_beanstalk_application_version |  |  | 

### elasticloadbalancing 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_elb |  |  x  |  x 

### elasticloadbalancingv2 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_alb_target_group |  |  | 
| aws_alb_listener |  |  | 
| aws_lb_target_group |  |  | 
| aws_alb_listener_rule |  |  | 
| aws_lb_listener |  |  | 
| aws_lb_listener_rule |  |  | 

### elastictranscoder 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_elastictranscoder_preset |  |  | 
| aws_elastictranscoder_pipeline |  |  | 

### emr 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_emr_security_configuration |  |  | 

### fsx 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_fsx_windows_file_system |  x  |  x  |  x 
| aws_fsx_lustre_file_system |  x  |  x  |  x 

### gamelift 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_gamelift_build |  |  x  |  x 
| aws_gamelift_alias |  |  x  |  x 
| aws_gamelift_game_session_queue |  |  | 

### globalaccelerator 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_globalaccelerator_accelerator |  |  x  |  x 

### glue 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_glue_trigger |  |  | 
| aws_glue_job |  |  | 
| aws_glue_crawler |  |  x  |  x 
| aws_glue_security_configuration |  |  | 

### iam 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_iam_user |  x  |  x  |  x 
| aws_iam_service_linked_role |  x  |  x  |  x 
| aws_iam_policy |  |  x  |  x 
| aws_iam_access_key |  |  x  |  x 
| aws_iam_role |  x  |  x  |  x 
| aws_iam_server_certificate |  |  | 
| aws_iam_group |  |  x  |  x 
| aws_iam_instance_profile |  |  x  |  x 

### iot 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_iot_thing |  |  | 
| aws_iot_policy |  |  | 
| aws_iot_certificate |  |  x  |  x 
| aws_iot_thing_type |  |  | 
| aws_iot_topic_rule |  |  | 

### kafka 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_msk_cluster |  x  |  x  |  x 
| aws_msk_configuration |  |  x  |  x 

### kinesisanalytics 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_kinesis_analytics_application |  |  | 

### kms 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_kms_alias |  |  | 
| aws_kms_external_key |  |  | 
| aws_kms_key |  |  | 

### lambda 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_lambda_event_source_mapping |  |  | 
| aws_lambda_function |  |  | 

### licensemanager 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_licensemanager_license_configuration |  |  | 

### lightsail 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_lightsail_key_pair |  x  |  | 
| aws_lightsail_domain |  x  |  | 
| aws_lightsail_static_ip |  |  | 
| aws_lightsail_instance |  x  |  | 

### mediaconvert 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_media_convert_queue |  |  | 

### mediapackage 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_media_package_channel |  x  |  | 

### mediastore 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_media_store_container |  |  x  |  x 

### mq 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_mq_configuration |  x  |  | 
| aws_mq_broker |  |  | 

### neptune 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_neptune_event_subscription |  |  | 

### opsworks 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_opsworks_user_profile |  |  | 
| aws_opsworks_stack |  |  | 
| aws_opsworks_instance |  |  | 

### qldb 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_qldb_ledger |  |  | 

### rds 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_db_event_subscription |  |  | 
| aws_db_subnet_group |  |  | 
| aws_db_instance |  |  x  |  x 
| aws_db_parameter_group |  |  | 
| aws_db_security_group |  |  | 
| aws_rds_global_cluster |  |  | 

### redshift 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_redshift_cluster |  x  |  | 
| aws_redshift_snapshot_copy_grant |  x  |  | 
| aws_redshift_event_subscription |  x  |  | 
| aws_redshift_snapshot_schedule |  x  |  | 

### route53 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_route53_health_check |  |  | 
| aws_route53_zone |  |  | 

### route53resolver 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_route53_resolver_rule_association |  |  | 
| aws_route53_resolver_rule |  |  | 
| aws_route53_resolver_endpoint |  |  x  |  x 

### s3 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_s3_bucket |  |  x  |  x 

### sagemaker 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_sagemaker_model |  |  x  |  x 
| aws_sagemaker_endpoint |  |  x  |  x 

### secretsmanager 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_secretsmanager_secret |  x  |  | 

### servicecatalog 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_servicecatalog_portfolio |  |  x  |  x 

### servicediscovery 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_service_discovery_service |  |  x  |  x 

### ses 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_ses_receipt_filter |  |  | 
| aws_ses_active_receipt_rule_set |  |  | 
| aws_ses_template |  |  | 
| aws_ses_configuration_set |  |  | 
| aws_ses_receipt_rule_set |  |  | 

### sfn 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_sfn_activity |  |  x  |  x 
| aws_sfn_state_machine |  |  x  |  x 

### sns 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_sns_platform_application |  |  | 
| aws_sns_topic_subscription |  |  | 
| aws_sns_topic |  |  | 

### ssm 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_ssm_maintenance_window |  |  | 
| aws_ssm_patch_group |  |  | 
| aws_ssm_document |  x  |  | 
| aws_ssm_patch_baseline |  |  | 
| aws_ssm_association |  |  | 
| aws_ssm_activation |  x  |  | 

### storagegateway 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_storagegateway_gateway |  |  | 

### transfer 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_transfer_server |  |  | 

### waf 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_waf_ipset |  |  | 
| aws_waf_web_acl |  |  | 
| aws_waf_regex_pattern_set |  |  | 
| aws_waf_sql_injection_match_set |  |  | 
| aws_waf_xss_match_set |  |  | 
| aws_waf_rule_group |  |  | 
| aws_waf_byte_match_set |  |  | 
| aws_waf_rate_based_rule |  |  | 
| aws_waf_regex_match_set |  |  | 
| aws_waf_rule |  |  | 
| aws_waf_geo_match_set |  |  | 
| aws_waf_size_constraint_set |  |  | 

### wafregional 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_wafregional_rate_based_rule |  |  | 
| aws_wafregional_size_constraint_set |  |  | 
| aws_wafregional_regex_match_set |  |  | 
| aws_wafregional_xss_match_set |  |  | 
| aws_wafregional_ipset |  |  | 
| aws_wafregional_geo_match_set |  |  | 
| aws_wafregional_regex_pattern_set |  |  | 
| aws_wafregional_rule_group |  |  | 
| aws_wafregional_rule |  |  | 
| aws_wafregional_byte_match_set |  |  | 
| aws_wafregional_web_acl |  |  | 
| aws_wafregional_sql_injection_match_set |  |  | 

### worklink 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_worklink_fleet |  |  x  |  x 

### workspaces 

| Resource Type  | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| aws_workspaces_ip_group |  |  |
