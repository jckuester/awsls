# awsls

A list command for AWS.

	awsls <terraform_resource_type>

Run, for example

    awsls aws_vpc

## Supported resource types

| Resource Type                    | Show Tags | Show Creation Time
| :-----------------------------   |:-------------:|:-----------------------:
| aws_accessanalyzer_analyzer         |    x  |  
| aws_acm_certificate         |    |  
| aws_api_gateway_usage_plan         |    x  |  
| aws_api_gateway_client_certificate         |    x  |  
| aws_api_gateway_api_key         |    x  |  
| aws_api_gateway_vpc_link         |    x  |  
| aws_api_gateway_rest_api         |    x  |  
| aws_api_gateway_domain_name         |    x  |  
| aws_apigatewayv2_api         |    x  |  
| aws_appmesh_mesh         |    |  
| aws_appsync_graphql_api         |    x  |  
| aws_athena_workgroup         |    |   x 
| aws_launch_configuration         |    |   x 
| aws_autoscaling_group         |    x  |   x 
| aws_backup_plan         |    |   x 
| aws_backup_vault         |    |   x 
| aws_batch_compute_environment         |    |  
| aws_batch_job_definition         |    |  
| aws_batch_job_queue         |    |  
| aws_cloudformation_stack         |    x  |   x 
| aws_cloudformation_stack_set         |    |  
| aws_cloudhsm_v2_cluster         |    |  
| aws_cloudwatch_dashboard         |    |  
| aws_cloudwatch_event_rule         |    |  
| aws_cloudwatch_log_destination         |    |   x 
| aws_cloudwatch_log_resource_policy         |    |  
| aws_cloudwatch_log_group         |    |   x 
| aws_codebuild_source_credential         |    |  
| aws_codecommit_repository         |    |  
| aws_codepipeline_webhook         |    x  |  
| aws_codestarnotifications_notification_rule         |    |  
| aws_config_config_rule         |    |  
| aws_config_delivery_channel         |    |  
| aws_config_configuration_recorder         |    |  
| aws_cur_report_definition         |    |  
| aws_dms_replication_subnet_group         |    |  
| aws_dms_replication_task         |    |  
| aws_dms_certificate         |    |  
| aws_dms_endpoint         |    |  
| aws_datasync_task         |    |  
| aws_datasync_agent         |    |  
| aws_dax_subnet_group         |    |  
| aws_dax_parameter_group         |    |  
| aws_devicefarm_project         |    |  
| aws_dx_hosted_public_virtual_interface         |    x  |  
| aws_dx_transit_virtual_interface         |    x  |  
| aws_dx_hosted_private_virtual_interface         |    x  |  
| aws_dx_connection         |    x  |  
| aws_dx_lag         |    x  |  
| aws_dx_public_virtual_interface         |    x  |  
| aws_dx_private_virtual_interface         |    x  |  
| aws_dx_hosted_transit_virtual_interface         |    x  |  
| aws_dlm_lifecycle_policy         |    x  |  
| aws_dynamodb_global_table         |    |  
| aws_placement_group         |    x  |  
| aws_ec2_transit_gateway         |    x  |   x 
| aws_default_route_table         |    x  |  
| aws_vpn_gateway         |    x  |  
| aws_ec2_capacity_reservation         |    x  |   x 
| aws_customer_gateway         |    x  |  
| aws_ec2_transit_gateway_vpc_attachment         |    x  |   x 
| aws_vpc_endpoint         |    x  |   x 
| aws_default_security_group         |    x  |  
| aws_vpc         |    x  |  
| aws_default_vpc         |    x  |  
| aws_ec2_fleet         |    x  |   x 
| aws_route_table         |    x  |  
| aws_egress_only_internet_gateway         |    x  |  
| aws_nat_gateway         |    x  |   x 
| aws_network_interface         |    |  
| aws_ec2_transit_gateway_route_table         |    x  |   x 
| aws_ec2_client_vpn_endpoint         |    x  |   x 
| aws_spot_instance_request         |    x  |   x 
| aws_subnet         |    x  |  
| aws_ec2_traffic_mirror_target         |    x  |  
| aws_eip         |    x  |  
| aws_network_acl         |    x  |  
| aws_vpc_peering_connection         |    x  |  
| aws_internet_gateway         |    x  |  
| aws_launch_template         |    x  |   x 
| aws_vpc_endpoint_service         |    x  |  
| aws_security_group         |    x  |  
| aws_ami         |    x  |   x 
| aws_key_pair         |    x  |  
| aws_ebs_snapshot         |    x  |   x 
| aws_ec2_traffic_mirror_session         |    x  |  
| aws_vpc_endpoint_connection_notification         |    |  
| aws_ebs_volume         |    x  |   x 
| aws_ec2_traffic_mirror_filter         |    x  |  
| aws_spot_fleet_request         |    |   x 
| aws_ecr_repository         |    |  
| aws_ecs_cluster         |    x  |  
| aws_efs_file_system         |    x  |   x 
| aws_efs_mount_target         |    |  
| aws_elasticache_replication_group         |    |  
| aws_elastic_beanstalk_application         |    |  
| aws_elastic_beanstalk_environment         |    |  
| aws_elastic_beanstalk_application_version         |    |  
| aws_elb         |    |   x 
| aws_alb_listener         |    |  
| aws_lb_listener_rule         |    |  
| aws_alb_listener_rule         |    |  
| aws_alb_target_group         |    |  
| aws_lb_target_group         |    |  
| aws_lb_listener         |    |  
| aws_elastictranscoder_pipeline         |    |  
| aws_elastictranscoder_preset         |    |  
| aws_emr_security_configuration         |    |  
| aws_fsx_lustre_file_system         |    x  |   x 
| aws_fsx_windows_file_system         |    x  |   x 
| aws_gamelift_alias         |    |   x 
| aws_gamelift_build         |    |   x 
| aws_gamelift_game_session_queue         |    |  
| aws_globalaccelerator_accelerator         |    |   x 
| aws_glue_crawler         |    |   x 
| aws_glue_security_configuration         |    |  
| aws_glue_job         |    |  
| aws_glue_trigger         |    |  
| aws_iam_user         |    x  |   x 
| aws_iam_access_key         |    |   x 
| aws_iam_policy         |    |   x 
| aws_iam_server_certificate         |    |  
| aws_iam_group         |    |   x 
| aws_iam_role         |    x  |   x 
| aws_iam_instance_profile         |    |   x 
| aws_iam_service_linked_role         |    x  |   x 
| aws_iot_certificate         |    |   x 
| aws_iot_thing_type         |    |  
| aws_iot_policy         |    |  
| aws_iot_thing         |    |  
| aws_iot_topic_rule         |    |  
| aws_msk_cluster         |    x  |   x 
| aws_msk_configuration         |    |   x 
| aws_kinesis_analytics_application         |    |  
| aws_kms_key         |    |  
| aws_kms_alias         |    |  
| aws_kms_external_key         |    |  
| aws_lambda_function         |    |  
| aws_lambda_event_source_mapping         |    |  
| aws_licensemanager_license_configuration         |    |  
| aws_lightsail_static_ip         |    |  
| aws_lightsail_domain         |    x  |  
| aws_lightsail_key_pair         |    x  |  
| aws_lightsail_instance         |    x  |  
| aws_media_convert_queue         |    |  
| aws_media_package_channel         |    x  |  
| aws_media_store_container         |    |   x 
| aws_mq_broker         |    |  
| aws_mq_configuration         |    x  |  
| aws_neptune_event_subscription         |    |  
| aws_opsworks_stack         |    |  
| aws_opsworks_user_profile         |    |  
| aws_opsworks_instance         |    |  
| aws_qldb_ledger         |    |  
| aws_db_event_subscription         |    |  
| aws_db_parameter_group         |    |  
| aws_db_subnet_group         |    |  
| aws_db_instance         |    |   x 
| aws_rds_global_cluster         |    |  
| aws_db_security_group         |    |  
| aws_redshift_snapshot_schedule         |    x  |  
| aws_redshift_event_subscription         |    x  |  
| aws_redshift_snapshot_copy_grant         |    x  |  
| aws_redshift_cluster         |    x  |  
| aws_route53_zone         |    |  
| aws_route53_health_check         |    |  
| aws_route53_resolver_rule         |    |  
| aws_route53_resolver_endpoint         |    |   x 
| aws_route53_resolver_rule_association         |    |  
| aws_s3_bucket         |    |   x 
| aws_sagemaker_model         |    |   x 
| aws_sagemaker_endpoint         |    |   x 
| aws_secretsmanager_secret         |    x  |  
| aws_servicecatalog_portfolio         |    |   x 
| aws_service_discovery_service         |    |   x 
| aws_ses_receipt_rule_set         |    |  
| aws_ses_template         |    |  
| aws_ses_configuration_set         |    |  
| aws_ses_active_receipt_rule_set         |    |  
| aws_ses_receipt_filter         |    |  
| aws_sfn_activity         |    |   x 
| aws_sfn_state_machine         |    |   x 
| aws_sns_topic_subscription         |    |  
| aws_sns_topic         |    |  
| aws_sns_platform_application         |    |  
| aws_ssm_document         |    x  |  
| aws_ssm_association         |    |  
| aws_ssm_activation         |    x  |  
| aws_ssm_patch_group         |    |  
| aws_ssm_maintenance_window         |    |  
| aws_ssm_patch_baseline         |    |  
| aws_storagegateway_gateway         |    |  
| aws_transfer_server         |    |  
| aws_waf_rule_group         |    |  
| aws_waf_rule         |    |  
| aws_waf_ipset         |    |  
| aws_waf_geo_match_set         |    |  
| aws_waf_byte_match_set         |    |  
| aws_waf_sql_injection_match_set         |    |  
| aws_waf_web_acl         |    |  
| aws_waf_size_constraint_set         |    |  
| aws_waf_xss_match_set         |    |  
| aws_waf_regex_pattern_set         |    |  
| aws_waf_rate_based_rule         |    |  
| aws_waf_regex_match_set         |    |  
| aws_wafregional_rate_based_rule         |    |  
| aws_wafregional_byte_match_set         |    |  
| aws_wafregional_web_acl         |    |  
| aws_wafregional_rule         |    |  
| aws_wafregional_size_constraint_set         |    |  
| aws_wafregional_regex_pattern_set         |    |  
| aws_wafregional_sql_injection_match_set         |    |  
| aws_wafregional_geo_match_set         |    |  
| aws_wafregional_xss_match_set         |    |  
| aws_wafregional_rule_group         |    |  
| aws_wafregional_regex_match_set         |    |  
| aws_wafregional_ipset         |    |  
| aws_worklink_fleet         |    |   x 
| aws_workspaces_ip_group         |    |