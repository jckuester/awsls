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

## Supported Resource Types

| Service / Type | Tags | Creation Time | Owner
| :------------- |:----:|:------------: | :---:
| **accessanalyzer** |
| aws_accessanalyzer_analyzer |  x  |  | 
| **acm** |
| aws_acm_certificate |  |  | 
| **apigateway** |
| aws_api_gateway_api_key |  x  |  | 
| aws_api_gateway_client_certificate |  x  |  | 
| aws_api_gateway_domain_name |  x  |  | 
| aws_api_gateway_rest_api |  x  |  | 
| aws_api_gateway_usage_plan |  x  |  | 
| aws_api_gateway_vpc_link |  x  |  | 
| **apigatewayv2** |
| aws_apigatewayv2_api |  x  |  | 
| **appmesh** |
| aws_appmesh_mesh |  |  | 
| **appsync** |
| aws_appsync_graphql_api |  x  |  | 
| **athena** |
| aws_athena_workgroup |  |  x  |  x 
| **autoscaling** |
| aws_autoscaling_group |  x  |  x  |  x 
| aws_launch_configuration |  |  x  |  x 
| **backup** |
| aws_backup_plan |  |  x  |  x 
| aws_backup_vault |  |  x  |  x 
| **batch** |
| aws_batch_compute_environment |  |  | 
| aws_batch_job_definition |  |  | 
| aws_batch_job_queue |  |  | 
| **cloudformation** |
| aws_cloudformation_stack |  x  |  x  |  x 
| aws_cloudformation_stack_set |  |  | 
| **cloudhsmv2** |
| aws_cloudhsm_v2_cluster |  |  | 
| **cloudwatch** |
| aws_cloudwatch_dashboard |  |  | 
| **cloudwatchevents** |
| aws_cloudwatch_event_rule |  |  | 
| **cloudwatchlogs** |
| aws_cloudwatch_log_destination |  |  x  |  x 
| aws_cloudwatch_log_group |  |  x  |  x 
| aws_cloudwatch_log_resource_policy |  |  | 
| **codebuild** |
| aws_codebuild_source_credential |  |  | 
| **codecommit** |
| aws_codecommit_repository |  |  | 
| **codepipeline** |
| aws_codepipeline_webhook |  x  |  | 
| **codestarnotifications** |
| aws_codestarnotifications_notification_rule |  |  | 
| **configservice** |
| aws_config_config_rule |  |  | 
| aws_config_configuration_recorder |  |  | 
| aws_config_delivery_channel |  |  | 
| **costandusagereportservice** |
| aws_cur_report_definition |  |  | 
| **databasemigrationservice** |
| aws_dms_certificate |  |  | 
| aws_dms_endpoint |  |  | 
| aws_dms_replication_subnet_group |  |  | 
| aws_dms_replication_task |  |  | 
| **datasync** |
| aws_datasync_agent |  |  | 
| aws_datasync_task |  |  | 
| **dax** |
| aws_dax_parameter_group |  |  | 
| aws_dax_subnet_group |  |  | 
| **devicefarm** |
| aws_devicefarm_project |  |  | 
| **directconnect** |
| aws_dx_connection |  x  |  | 
| aws_dx_hosted_private_virtual_interface |  x  |  | 
| aws_dx_hosted_public_virtual_interface |  x  |  | 
| aws_dx_hosted_transit_virtual_interface |  x  |  | 
| aws_dx_lag |  x  |  | 
| aws_dx_private_virtual_interface |  x  |  | 
| aws_dx_public_virtual_interface |  x  |  | 
| aws_dx_transit_virtual_interface |  x  |  | 
| **dlm** |
| aws_dlm_lifecycle_policy |  x  |  | 
| **dynamodb** |
| aws_dynamodb_global_table |  |  | 
| **ec2** |
| aws_ami |  x  |  x  |  x 
| aws_customer_gateway |  x  |  | 
| aws_default_route_table |  x  |  | 
| aws_default_security_group |  x  |  | 
| aws_default_vpc |  x  |  | 
| aws_ebs_snapshot |  x  |  x  |  x 
| aws_ebs_volume |  x  |  x  |  x 
| aws_ec2_capacity_reservation |  x  |  x  |  x 
| aws_ec2_client_vpn_endpoint |  x  |  x  |  x 
| aws_ec2_fleet |  x  |  x  |  x 
| aws_ec2_traffic_mirror_filter |  x  |  | 
| aws_ec2_traffic_mirror_session |  x  |  | 
| aws_ec2_traffic_mirror_target |  x  |  | 
| aws_ec2_transit_gateway |  x  |  x  |  x 
| aws_ec2_transit_gateway_route_table |  x  |  x  |  x 
| aws_ec2_transit_gateway_vpc_attachment |  x  |  x  |  x 
| aws_egress_only_internet_gateway |  x  |  | 
| aws_eip |  x  |  | 
| aws_internet_gateway |  x  |  | 
| aws_key_pair |  x  |  | 
| aws_launch_template |  x  |  x  |  x 
| aws_nat_gateway |  x  |  x  |  x 
| aws_network_acl |  x  |  | 
| aws_network_interface |  |  | 
| aws_placement_group |  x  |  | 
| aws_route_table |  x  |  | 
| aws_security_group |  x  |  | 
| aws_spot_fleet_request |  |  x  |  x 
| aws_spot_instance_request |  x  |  x  |  x 
| aws_subnet |  x  |  | 
| aws_vpc |  x  |  | 
| aws_vpc_endpoint |  x  |  x  |  x 
| aws_vpc_endpoint_connection_notification |  |  | 
| aws_vpc_endpoint_service |  x  |  | 
| aws_vpc_peering_connection |  x  |  | 
| aws_vpn_gateway |  x  |  | 
| **ecr** |
| aws_ecr_repository |  |  | 
| **ecs** |
| aws_ecs_cluster |  x  |  | 
| **efs** |
| aws_efs_file_system |  x  |  x  |  x 
| aws_efs_mount_target |  |  | 
| **elasticache** |
| aws_elasticache_replication_group |  |  | 
| **elasticbeanstalk** |
| aws_elastic_beanstalk_application |  |  | 
| aws_elastic_beanstalk_application_version |  |  | 
| aws_elastic_beanstalk_environment |  |  | 
| **elasticloadbalancing** |
| aws_elb |  |  x  |  x 
| **elasticloadbalancingv2** |
| aws_alb_listener |  |  | 
| aws_alb_listener_rule |  |  | 
| aws_alb_target_group |  |  | 
| aws_lb_listener |  |  | 
| aws_lb_listener_rule |  |  | 
| aws_lb_target_group |  |  | 
| **elastictranscoder** |
| aws_elastictranscoder_pipeline |  |  | 
| aws_elastictranscoder_preset |  |  | 
| **emr** |
| aws_emr_security_configuration |  |  | 
| **fsx** |
| aws_fsx_lustre_file_system |  x  |  x  |  x 
| aws_fsx_windows_file_system |  x  |  x  |  x 
| **gamelift** |
| aws_gamelift_alias |  |  x  |  x 
| aws_gamelift_build |  |  x  |  x 
| aws_gamelift_game_session_queue |  |  | 
| **globalaccelerator** |
| aws_globalaccelerator_accelerator |  |  x  |  x 
| **glue** |
| aws_glue_crawler |  |  x  |  x 
| aws_glue_job |  |  | 
| aws_glue_security_configuration |  |  | 
| aws_glue_trigger |  |  | 
| **iam** |
| aws_iam_access_key |  |  x  |  x 
| aws_iam_group |  |  x  |  x 
| aws_iam_instance_profile |  |  x  |  x 
| aws_iam_policy |  |  x  |  x 
| aws_iam_role |  x  |  x  |  x 
| aws_iam_server_certificate |  |  | 
| aws_iam_service_linked_role |  x  |  x  |  x 
| aws_iam_user |  x  |  x  |  x 
| **iot** |
| aws_iot_certificate |  |  x  |  x 
| aws_iot_policy |  |  | 
| aws_iot_thing |  |  | 
| aws_iot_thing_type |  |  | 
| aws_iot_topic_rule |  |  | 
| **kafka** |
| aws_msk_cluster |  x  |  x  |  x 
| aws_msk_configuration |  |  x  |  x 
| **kinesisanalytics** |
| aws_kinesis_analytics_application |  |  | 
| **kms** |
| aws_kms_alias |  |  | 
| aws_kms_external_key |  |  | 
| aws_kms_key |  |  | 
| **lambda** |
| aws_lambda_event_source_mapping |  |  | 
| aws_lambda_function |  |  | 
| **licensemanager** |
| aws_licensemanager_license_configuration |  |  | 
| **lightsail** |
| aws_lightsail_domain |  x  |  | 
| aws_lightsail_instance |  x  |  | 
| aws_lightsail_key_pair |  x  |  | 
| aws_lightsail_static_ip |  |  | 
| **mediaconvert** |
| aws_media_convert_queue |  |  | 
| **mediapackage** |
| aws_media_package_channel |  x  |  | 
| **mediastore** |
| aws_media_store_container |  |  x  |  x 
| **mq** |
| aws_mq_broker |  |  | 
| aws_mq_configuration |  x  |  | 
| **neptune** |
| aws_neptune_event_subscription |  |  | 
| **opsworks** |
| aws_opsworks_instance |  |  | 
| aws_opsworks_stack |  |  | 
| aws_opsworks_user_profile |  |  | 
| **qldb** |
| aws_qldb_ledger |  |  | 
| **rds** |
| aws_db_event_subscription |  |  | 
| aws_db_instance |  |  x  |  x 
| aws_db_parameter_group |  |  | 
| aws_db_security_group |  |  | 
| aws_db_subnet_group |  |  | 
| aws_rds_global_cluster |  |  | 
| **redshift** |
| aws_redshift_cluster |  x  |  | 
| aws_redshift_event_subscription |  x  |  | 
| aws_redshift_snapshot_copy_grant |  x  |  | 
| aws_redshift_snapshot_schedule |  x  |  | 
| **route53** |
| aws_route53_health_check |  |  | 
| aws_route53_zone |  |  | 
| **route53resolver** |
| aws_route53_resolver_endpoint |  |  x  |  x 
| aws_route53_resolver_rule |  |  | 
| aws_route53_resolver_rule_association |  |  | 
| **s3** |
| aws_s3_bucket |  |  x  |  x 
| **sagemaker** |
| aws_sagemaker_endpoint |  |  x  |  x 
| aws_sagemaker_model |  |  x  |  x 
| **secretsmanager** |
| aws_secretsmanager_secret |  x  |  | 
| **servicecatalog** |
| aws_servicecatalog_portfolio |  |  x  |  x 
| **servicediscovery** |
| aws_service_discovery_service |  |  x  |  x 
| **ses** |
| aws_ses_active_receipt_rule_set |  |  | 
| aws_ses_configuration_set |  |  | 
| aws_ses_receipt_filter |  |  | 
| aws_ses_receipt_rule_set |  |  | 
| aws_ses_template |  |  | 
| **sfn** |
| aws_sfn_activity |  |  x  |  x 
| aws_sfn_state_machine |  |  x  |  x 
| **sns** |
| aws_sns_platform_application |  |  | 
| aws_sns_topic |  |  | 
| aws_sns_topic_subscription |  |  | 
| **ssm** |
| aws_ssm_activation |  x  |  | 
| aws_ssm_association |  |  | 
| aws_ssm_document |  x  |  | 
| aws_ssm_maintenance_window |  |  | 
| aws_ssm_patch_baseline |  |  | 
| aws_ssm_patch_group |  |  | 
| **storagegateway** |
| aws_storagegateway_gateway |  |  | 
| **transfer** |
| aws_transfer_server |  |  | 
| **waf** |
| aws_waf_byte_match_set |  |  | 
| aws_waf_geo_match_set |  |  | 
| aws_waf_ipset |  |  | 
| aws_waf_rate_based_rule |  |  | 
| aws_waf_regex_match_set |  |  | 
| aws_waf_regex_pattern_set |  |  | 
| aws_waf_rule |  |  | 
| aws_waf_rule_group |  |  | 
| aws_waf_size_constraint_set |  |  | 
| aws_waf_sql_injection_match_set |  |  | 
| aws_waf_web_acl |  |  | 
| aws_waf_xss_match_set |  |  | 
| **wafregional** |
| aws_wafregional_byte_match_set |  |  | 
| aws_wafregional_geo_match_set |  |  | 
| aws_wafregional_ipset |  |  | 
| aws_wafregional_rate_based_rule |  |  | 
| aws_wafregional_regex_match_set |  |  | 
| aws_wafregional_regex_pattern_set |  |  | 
| aws_wafregional_rule |  |  | 
| aws_wafregional_rule_group |  |  | 
| aws_wafregional_size_constraint_set |  |  | 
| aws_wafregional_sql_injection_match_set |  |  | 
| aws_wafregional_web_acl |  |  | 
| aws_wafregional_xss_match_set |  |  | 
| **worklink** |
| aws_worklink_fleet |  |  x  |  x 
| **workspaces** |
| aws_workspaces_ip_group |  |  |