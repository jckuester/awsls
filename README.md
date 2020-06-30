# awsls

[![Release](https://img.shields.io/github/release/jckuester/awsls.svg?style=for-the-badge)](https://github.com/jckuester/awsls/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE.md)
[![Travis](https://img.shields.io/travis/jckuester/awsls/master.svg?style=for-the-badge)](https://travis-ci.org/jckuester/awsls)

A list command for AWS resources. It supports listing of [over 200 types of resources](#supported-resources)
across 76 different AWS services.
 
The goal is to support every AWS resource that is also covered by Terraform (currently over 500) without adding
much code but rather generating it. If you are interested, the code of the list function generator is [here](./gen).

If you run into issues with any resources, please open an issue or ping me on [Twitter](https://twitter.com/jckuester).

Happy listing!

## Example

(**Note:** In the examples below, credentials and region for the AWS account you want to list resources in have been set
via the usual [environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-envvars.html), e.g.,
`AWS_PROFILE=<myaccount>` and `AWS_DEFAULT_REGION=<myregion>`, but can also be provided via `-profile` and `-region` flag.)

Terraform resource names are used to tell awsls which resources to list. For example, `awsls aws_instance` shows
all EC2 instances. In addition to the default attributes `TYPE`, `ID`, `REGION`, and `CREATED` timestamp, additional attributes
can be displayed via the `-attributes <comma-separated list>` flag. Every attribute in the Terraform documentation 
([here](https://www.terraform.io/docs/providers/aws/r/instance.html) are the attributes for `aws_instance`) is a valid one:

![](img/instance.gif)

To list multiple resource types at once, use glob patterns. For example, `awsls "aws_iam_*` lists all IAM resources:

![](img/iam.gif)

## Usage

	awsls [flags] <resource_type glob pattern>

To see options available run `awsls --help`.

## Installation

It's recommended to install a specific version of awsls available on the
[releases page](https://github.com/jckuester/awsls/releases).

Here is the recommended way to install awsls v0.1.2:

```bash
# install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/jckuester/awsls/master/install.sh | sh -s v0.1.2
```

## Supported resources

Currently, all types of resources in the table below can be listed with awsls. The `Tags` column shows if a resource
supports displaying tags, the `Creation Time` column if a resource has a creation timestamp, and the `Owner` column if
resources are pre-filtered by account ID.

| Service / Type | Tags | Creation Time | Owner
| :------------- | :--: | :-----------: | :---:
| **accessanalyzer** |
| aws_accessanalyzer_analyzer |  x  |  |
| **acm** |
| aws_acm_certificate |  x  |  |
| **apigateway** |
| aws_api_gateway_api_key |  x  |  |
| aws_api_gateway_client_certificate |  x  |  |
| aws_api_gateway_domain_name |  x  |  |
| aws_api_gateway_rest_api |  x  |  |
| aws_api_gateway_usage_plan |  x  |  |
| aws_api_gateway_vpc_link |  x  |  |
| **apigatewayv2** |
| aws_apigatewayv2_api |  x  |  |
| aws_apigatewayv2_domain_name |  x  |  |
| aws_apigatewayv2_vpc_link |  x  |  |
| **appmesh** |
| aws_appmesh_mesh |  x  |  |
| **appsync** |
| aws_appsync_graphql_api |  x  |  |
| **athena** |
| aws_athena_workgroup |  x  |  x  |
| **autoscaling** |
| aws_autoscaling_group |  x  |  x  |
| aws_launch_configuration |  |  x  |
| **backup** |
| aws_backup_plan |  x  |  x  |
| aws_backup_vault |  x  |  x  |
| **batch** |
| aws_batch_compute_environment |  |  |
| aws_batch_job_definition |  |  |
| aws_batch_job_queue |  |  |
| **cloudformation** |
| aws_cloudformation_stack |  x  |  x  |
| aws_cloudformation_stack_set |  x  |  |
| **cloudhsmv2** |
| aws_cloudhsm_v2_cluster |  x  |  |
| **cloudwatch** |
| aws_cloudwatch_dashboard |  |  |
| **cloudwatchevents** |
| aws_cloudwatch_event_rule |  x  |  |
| **cloudwatchlogs** |
| aws_cloudwatch_log_destination |  |  x  |
| aws_cloudwatch_log_group |  x  |  x  |
| aws_cloudwatch_log_resource_policy |  |  |
| **codebuild** |
| aws_codebuild_source_credential |  |  |
| **codecommit** |
| aws_codecommit_repository |  x  |  |
| **codepipeline** |
| aws_codepipeline_webhook |  x  |  |
| **codestarnotifications** |
| aws_codestarnotifications_notification_rule |  x  |  |
| **configservice** |
| aws_config_config_rule |  x  |  |
| aws_config_configuration_recorder |  |  |
| aws_config_delivery_channel |  |  |
| **costandusagereportservice** |
| aws_cur_report_definition |  |  |
| **databasemigrationservice** |
| aws_dms_certificate |  |  |
| aws_dms_endpoint |  x  |  |
| aws_dms_replication_subnet_group |  x  |  |
| aws_dms_replication_task |  x  |  |
| **datasync** |
| aws_datasync_agent |  x  |  |
| aws_datasync_task |  x  |  |
| **dax** |
| aws_dax_parameter_group |  |  |
| aws_dax_subnet_group |  |  |
| **devicefarm** |
| aws_devicefarm_project |  |  |
| **directconnect** |
| aws_dx_connection |  x  |  |
| aws_dx_hosted_private_virtual_interface |  |  |
| aws_dx_hosted_public_virtual_interface |  |  |
| aws_dx_hosted_transit_virtual_interface |  |  |
| aws_dx_lag |  x  |  |
| aws_dx_private_virtual_interface |  x  |  |
| aws_dx_public_virtual_interface |  x  |  |
| aws_dx_transit_virtual_interface |  x  |  |
| **dlm** |
| aws_dlm_lifecycle_policy |  x  |  |
| **dynamodb** |
| aws_dynamodb_global_table |  |  |
| **ec2** |
| aws_ami |  x  |  x  | x |
| aws_ebs_snapshot |  x  |  x  | x |
| aws_ebs_volume |  x  |  x  |
| aws_ec2_capacity_reservation |  x  |  x  | x |
| aws_ec2_client_vpn_endpoint |  x  |  x  |
| aws_ec2_fleet |  x  |  x  |
| aws_ec2_local_gateway_route_table_vpc_association |  x  |  |
| aws_ec2_traffic_mirror_filter |  x  |  |
| aws_ec2_traffic_mirror_session |  x  |  | x |
| aws_ec2_traffic_mirror_target |  x  |  | x |
| aws_ec2_transit_gateway |  x  |  x  | x |
| aws_ec2_transit_gateway_peering_attachment |  x  |  x  |
| aws_ec2_transit_gateway_route_table |  x  |  x  |
| aws_ec2_transit_gateway_vpc_attachment |  x  |  x  |
| aws_egress_only_internet_gateway |  x  |  |
| aws_eip |  x  |  |
| aws_instance |  x  | x | x |
| aws_internet_gateway |  x  |  | x |
| aws_key_pair |  x  |  |
| aws_launch_template |  x  |  x  |
| aws_nat_gateway |  x  |  x  |
| aws_network_acl |  x  |  | x |
| aws_network_interface |  x  |  | x |
| aws_placement_group |  x  |  |
| aws_route_table |  x  |  | x |
| aws_security_group |  x  |  | x |
| aws_spot_fleet_request |  x  |  x  |
| aws_spot_instance_request |  x  |  x  |
| aws_subnet |  x  |  | x |
| aws_vpc |  x  |  | x |
| aws_vpc_endpoint |  x  |  x  | x |
| aws_vpc_endpoint_connection_notification |  |  |
| aws_vpc_endpoint_service |  x  |  |
| aws_vpc_peering_connection |  x  |  |
| aws_vpn_gateway |  x  |  |
| **ecr** |
| aws_ecr_repository |  x  |  |
| **ecs** |
| aws_ecs_cluster |  x  |  |
| **efs** |
| aws_efs_access_point |  x  |  | x |
| aws_efs_file_system |  x  |  x  | x |
| **elasticache** |
| aws_elasticache_replication_group |  x  |  |
| **elasticbeanstalk** |
| aws_elastic_beanstalk_application |  x  |  |
| aws_elastic_beanstalk_application_version |  x  |  |
| aws_elastic_beanstalk_environment |  x  |  |
| **elasticloadbalancing** |
| aws_elb |  x  |  x  |
| **elasticloadbalancingv2** |
| aws_alb_target_group |  x  |  |
| aws_lb_target_group |  x  |  |
| **elastictranscoder** |
| aws_elastictranscoder_pipeline |  |  |
| aws_elastictranscoder_preset |  |  |
| **emr** |
| aws_emr_security_configuration |  |  |
| **fsx** |
| aws_fsx_lustre_file_system |  x  |  x  | x |
| aws_fsx_windows_file_system |  x  |  x  | x |
| **gamelift** |
| aws_gamelift_alias |  x  |  x  |
| aws_gamelift_build |  x  |  x  |
| aws_gamelift_game_session_queue |  x  |  |
| **globalaccelerator** |
| aws_globalaccelerator_accelerator |  x  |  x  |
| **glue** |
| aws_glue_crawler |  x  |  x  |
| aws_glue_job |  x  |  |
| aws_glue_security_configuration |  |  |
| aws_glue_trigger |  x  |  |
| **iam** |
| aws_iam_access_key |  |  x  |
| aws_iam_group |  |  x  |
| aws_iam_instance_profile |  |  x  |
| aws_iam_policy |  |  x  |
| aws_iam_role |  x  |  x  |
| aws_iam_server_certificate |  |  |
| aws_iam_service_linked_role |  |  x  |
| aws_iam_user |  x  |  x  |
| **iot** |
| aws_iot_certificate |  |  x  |
| aws_iot_policy |  |  |
| aws_iot_thing |  |  |
| aws_iot_thing_type |  |  |
| aws_iot_topic_rule |  x  |  |
| **kafka** |
| aws_msk_cluster |  x  |  x  |
| aws_msk_configuration |  |  x  |
| **kinesisanalytics** |
| aws_kinesis_analytics_application |  x  |  |
| **kms** |
| aws_kms_external_key |  x  |  |
| aws_kms_key |  x  |  |
| **lambda** |
| aws_lambda_event_source_mapping |  |  |
| aws_lambda_function |  x  |  |
| **licensemanager** |
| aws_licensemanager_license_configuration |  x  |  |
| **lightsail** |
| aws_lightsail_domain |  |  |
| aws_lightsail_instance |  x  |  |
| aws_lightsail_key_pair |  |  |
| aws_lightsail_static_ip |  |  |
| **mediaconvert** |
| aws_media_convert_queue |  x  |  |
| **mediapackage** |
| aws_media_package_channel |  x  |  |
| **mediastore** |
| aws_media_store_container |  x  |  x  |
| **mq** |
| aws_mq_broker |  x  |  |
| aws_mq_configuration |  x  |  |
| **neptune** |
| aws_neptune_event_subscription |  x  |  |
| **opsworks** |
| aws_opsworks_stack |  x  |  |
| aws_opsworks_user_profile |  |  |
| **qldb** |
| aws_qldb_ledger |  x  |  |
| **rds** |
| aws_db_event_subscription |  x  |  |
| aws_db_instance |  x  |  x  |
| aws_db_parameter_group |  x  |  |
| aws_db_security_group |  x  |  | x |
| aws_db_snapshot |  x  |  x  |
| aws_db_subnet_group |  x  |  |
| aws_rds_global_cluster |  |  |
| **redshift** |
| aws_redshift_cluster |  x  |  |
| aws_redshift_event_subscription |  x  |  |
| aws_redshift_snapshot_copy_grant |  x  |  |
| aws_redshift_snapshot_schedule |  x  |  |
| **route53** |
| aws_route53_health_check |  x  |  |
| aws_route53_zone |  x  |  |
| **route53resolver** |
| aws_route53_resolver_endpoint |  x  |  x  |
| aws_route53_resolver_rule |  x  |  | x |
| aws_route53_resolver_rule_association |  |  |
| **s3** |
| aws_s3_bucket |  x  |  x  |
| **sagemaker** |
| aws_sagemaker_endpoint |  x  |  x  |
| aws_sagemaker_model |  x  |  x  |
| **secretsmanager** |
| aws_secretsmanager_secret |  x  |  |
| **servicecatalog** |
| aws_servicecatalog_portfolio |  x  |  x  |
| **servicediscovery** |
| aws_service_discovery_service |  x  |  x  |
| **ses** |
| aws_ses_active_receipt_rule_set |  |  |
| aws_ses_configuration_set |  |  |
| aws_ses_receipt_filter |  |  |
| aws_ses_receipt_rule_set |  |  |
| aws_ses_template |  |  |
| **sfn** |
| aws_sfn_activity |  x  |  x  |
| aws_sfn_state_machine |  x  |  x  |
| **sns** |
| aws_sns_platform_application |  |  |
| aws_sns_topic |  x  |  |
| aws_sns_topic_subscription |  |  |
| **ssm** |
| aws_ssm_activation |  x  |  |
| aws_ssm_association |  |  |
| aws_ssm_document |  x  |  |
| aws_ssm_maintenance_window |  x  |  |
| aws_ssm_patch_baseline |  x  |  |
| aws_ssm_patch_group |  |  |
| **storagegateway** |
| aws_storagegateway_gateway |  x  |  |
| **transfer** |
| aws_transfer_server |  x  |  |
| **waf** |
| aws_waf_byte_match_set |  |  |
| aws_waf_geo_match_set |  |  |
| aws_waf_ipset |  |  |
| aws_waf_rate_based_rule |  x  |  |
| aws_waf_regex_match_set |  |  |
| aws_waf_regex_pattern_set |  |  |
| aws_waf_rule |  x  |  |
| aws_waf_rule_group |  x  |  |
| aws_waf_size_constraint_set |  |  |
| aws_waf_sql_injection_match_set |  |  |
| aws_waf_web_acl |  x  |  |
| aws_waf_xss_match_set |  |  |
| **wafregional** |
| aws_wafregional_byte_match_set |  |  |
| aws_wafregional_geo_match_set |  |  |
| aws_wafregional_ipset |  |  |
| aws_wafregional_rate_based_rule |  x  |  |
| aws_wafregional_regex_match_set |  |  |
| aws_wafregional_regex_pattern_set |  |  |
| aws_wafregional_rule |  x  |  |
| aws_wafregional_rule_group |  x  |  |
| aws_wafregional_size_constraint_set |  |  |
| aws_wafregional_sql_injection_match_set |  |  |
| aws_wafregional_web_acl |  x  |  |
| aws_wafregional_xss_match_set |  |  |
| **wafv2** |
| aws_wafv2_web_acl_logging_configuration |  |  |
| **worklink** |
| aws_worklink_fleet |  |  x  |
| **workspaces** |
| aws_workspaces_ip_group |  x  |  |