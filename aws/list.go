// Code is generated. DO NOT EDIT.

package aws

import (
	"fmt"
	"time"

	terradozer "github.com/jckuester/terradozer/pkg/resource"
)

type Resource struct {
	Type      string
	ID        string
	Region    string
	Tags      map[string]string
	CreatedAt *time.Time
	terradozer.UpdatableResource
}

func ListResourcesByType(client *Client, resourceType string) ([]Resource, error) {
	switch resourceType {
	case "aws_accessanalyzer_analyzer":
		return ListAccessanalyzerAnalyzer(client)
	case "aws_acm_certificate":
		return ListAcmCertificate(client)
	case "aws_alb_target_group":
		return ListAlbTargetGroup(client)
	case "aws_ami":
		return ListAmi(client)
	case "aws_api_gateway_api_key":
		return ListApiGatewayApiKey(client)
	case "aws_api_gateway_client_certificate":
		return ListApiGatewayClientCertificate(client)
	case "aws_api_gateway_domain_name":
		return ListApiGatewayDomainName(client)
	case "aws_api_gateway_rest_api":
		return ListApiGatewayRestApi(client)
	case "aws_api_gateway_usage_plan":
		return ListApiGatewayUsagePlan(client)
	case "aws_api_gateway_vpc_link":
		return ListApiGatewayVpcLink(client)
	case "aws_apigatewayv2_api":
		return ListApigatewayv2Api(client)
	case "aws_apigatewayv2_domain_name":
		return ListApigatewayv2DomainName(client)
	case "aws_apigatewayv2_vpc_link":
		return ListApigatewayv2VpcLink(client)
	case "aws_appmesh_mesh":
		return ListAppmeshMesh(client)
	case "aws_appsync_graphql_api":
		return ListAppsyncGraphqlApi(client)
	case "aws_athena_workgroup":
		return ListAthenaWorkgroup(client)
	case "aws_autoscaling_group":
		return ListAutoscalingGroup(client)
	case "aws_backup_plan":
		return ListBackupPlan(client)
	case "aws_backup_vault":
		return ListBackupVault(client)
	case "aws_batch_compute_environment":
		return ListBatchComputeEnvironment(client)
	case "aws_batch_job_definition":
		return ListBatchJobDefinition(client)
	case "aws_batch_job_queue":
		return ListBatchJobQueue(client)
	case "aws_cloudformation_stack":
		return ListCloudformationStack(client)
	case "aws_cloudformation_stack_set":
		return ListCloudformationStackSet(client)
	case "aws_cloudhsm_v2_cluster":
		return ListCloudhsmV2Cluster(client)
	case "aws_cloudwatch_dashboard":
		return ListCloudwatchDashboard(client)
	case "aws_cloudwatch_event_rule":
		return ListCloudwatchEventRule(client)
	case "aws_cloudwatch_log_destination":
		return ListCloudwatchLogDestination(client)
	case "aws_cloudwatch_log_group":
		return ListCloudwatchLogGroup(client)
	case "aws_cloudwatch_log_resource_policy":
		return ListCloudwatchLogResourcePolicy(client)
	case "aws_codebuild_source_credential":
		return ListCodebuildSourceCredential(client)
	case "aws_codecommit_repository":
		return ListCodecommitRepository(client)
	case "aws_codepipeline_webhook":
		return ListCodepipelineWebhook(client)
	case "aws_codestarnotifications_notification_rule":
		return ListCodestarnotificationsNotificationRule(client)
	case "aws_config_config_rule":
		return ListConfigConfigRule(client)
	case "aws_config_configuration_recorder":
		return ListConfigConfigurationRecorder(client)
	case "aws_config_delivery_channel":
		return ListConfigDeliveryChannel(client)
	case "aws_cur_report_definition":
		return ListCurReportDefinition(client)
	case "aws_datasync_agent":
		return ListDatasyncAgent(client)
	case "aws_datasync_task":
		return ListDatasyncTask(client)
	case "aws_dax_parameter_group":
		return ListDaxParameterGroup(client)
	case "aws_dax_subnet_group":
		return ListDaxSubnetGroup(client)
	case "aws_db_event_subscription":
		return ListDbEventSubscription(client)
	case "aws_db_instance":
		return ListDbInstance(client)
	case "aws_db_parameter_group":
		return ListDbParameterGroup(client)
	case "aws_db_security_group":
		return ListDbSecurityGroup(client)
	case "aws_db_snapshot":
		return ListDbSnapshot(client)
	case "aws_db_subnet_group":
		return ListDbSubnetGroup(client)
	case "aws_devicefarm_project":
		return ListDevicefarmProject(client)
	case "aws_dlm_lifecycle_policy":
		return ListDlmLifecyclePolicy(client)
	case "aws_dms_certificate":
		return ListDmsCertificate(client)
	case "aws_dms_endpoint":
		return ListDmsEndpoint(client)
	case "aws_dms_replication_subnet_group":
		return ListDmsReplicationSubnetGroup(client)
	case "aws_dms_replication_task":
		return ListDmsReplicationTask(client)
	case "aws_dx_connection":
		return ListDxConnection(client)
	case "aws_dx_hosted_private_virtual_interface":
		return ListDxHostedPrivateVirtualInterface(client)
	case "aws_dx_hosted_public_virtual_interface":
		return ListDxHostedPublicVirtualInterface(client)
	case "aws_dx_hosted_transit_virtual_interface":
		return ListDxHostedTransitVirtualInterface(client)
	case "aws_dx_lag":
		return ListDxLag(client)
	case "aws_dx_private_virtual_interface":
		return ListDxPrivateVirtualInterface(client)
	case "aws_dx_public_virtual_interface":
		return ListDxPublicVirtualInterface(client)
	case "aws_dx_transit_virtual_interface":
		return ListDxTransitVirtualInterface(client)
	case "aws_dynamodb_global_table":
		return ListDynamodbGlobalTable(client)
	case "aws_ebs_snapshot":
		return ListEbsSnapshot(client)
	case "aws_ebs_volume":
		return ListEbsVolume(client)
	case "aws_ec2_capacity_reservation":
		return ListEc2CapacityReservation(client)
	case "aws_ec2_client_vpn_endpoint":
		return ListEc2ClientVpnEndpoint(client)
	case "aws_ec2_fleet":
		return ListEc2Fleet(client)
	case "aws_ec2_local_gateway_route_table_vpc_association":
		return ListEc2LocalGatewayRouteTableVpcAssociation(client)
	case "aws_ec2_traffic_mirror_filter":
		return ListEc2TrafficMirrorFilter(client)
	case "aws_ec2_traffic_mirror_session":
		return ListEc2TrafficMirrorSession(client)
	case "aws_ec2_traffic_mirror_target":
		return ListEc2TrafficMirrorTarget(client)
	case "aws_ec2_transit_gateway":
		return ListEc2TransitGateway(client)
	case "aws_ec2_transit_gateway_peering_attachment":
		return ListEc2TransitGatewayPeeringAttachment(client)
	case "aws_ec2_transit_gateway_route_table":
		return ListEc2TransitGatewayRouteTable(client)
	case "aws_ec2_transit_gateway_vpc_attachment":
		return ListEc2TransitGatewayVpcAttachment(client)
	case "aws_ecr_repository":
		return ListEcrRepository(client)
	case "aws_ecs_cluster":
		return ListEcsCluster(client)
	case "aws_efs_access_point":
		return ListEfsAccessPoint(client)
	case "aws_efs_file_system":
		return ListEfsFileSystem(client)
	case "aws_egress_only_internet_gateway":
		return ListEgressOnlyInternetGateway(client)
	case "aws_eip":
		return ListEip(client)
	case "aws_elastic_beanstalk_application":
		return ListElasticBeanstalkApplication(client)
	case "aws_elastic_beanstalk_application_version":
		return ListElasticBeanstalkApplicationVersion(client)
	case "aws_elastic_beanstalk_environment":
		return ListElasticBeanstalkEnvironment(client)
	case "aws_elasticache_replication_group":
		return ListElasticacheReplicationGroup(client)
	case "aws_elastictranscoder_pipeline":
		return ListElastictranscoderPipeline(client)
	case "aws_elastictranscoder_preset":
		return ListElastictranscoderPreset(client)
	case "aws_elb":
		return ListElb(client)
	case "aws_emr_security_configuration":
		return ListEmrSecurityConfiguration(client)
	case "aws_fsx_lustre_file_system":
		return ListFsxLustreFileSystem(client)
	case "aws_fsx_windows_file_system":
		return ListFsxWindowsFileSystem(client)
	case "aws_gamelift_alias":
		return ListGameliftAlias(client)
	case "aws_gamelift_build":
		return ListGameliftBuild(client)
	case "aws_gamelift_game_session_queue":
		return ListGameliftGameSessionQueue(client)
	case "aws_globalaccelerator_accelerator":
		return ListGlobalacceleratorAccelerator(client)
	case "aws_glue_crawler":
		return ListGlueCrawler(client)
	case "aws_glue_job":
		return ListGlueJob(client)
	case "aws_glue_security_configuration":
		return ListGlueSecurityConfiguration(client)
	case "aws_glue_trigger":
		return ListGlueTrigger(client)
	case "aws_iam_access_key":
		return ListIamAccessKey(client)
	case "aws_iam_group":
		return ListIamGroup(client)
	case "aws_iam_instance_profile":
		return ListIamInstanceProfile(client)
	case "aws_iam_policy":
		return ListIamPolicy(client)
	case "aws_iam_role":
		return ListIamRole(client)
	case "aws_iam_server_certificate":
		return ListIamServerCertificate(client)
	case "aws_iam_service_linked_role":
		return ListIamServiceLinkedRole(client)
	case "aws_iam_user":
		return ListIamUser(client)
	case "aws_instance":
		return ListInstance(client)
	case "aws_internet_gateway":
		return ListInternetGateway(client)
	case "aws_iot_certificate":
		return ListIotCertificate(client)
	case "aws_iot_policy":
		return ListIotPolicy(client)
	case "aws_iot_thing":
		return ListIotThing(client)
	case "aws_iot_thing_type":
		return ListIotThingType(client)
	case "aws_iot_topic_rule":
		return ListIotTopicRule(client)
	case "aws_key_pair":
		return ListKeyPair(client)
	case "aws_kinesis_analytics_application":
		return ListKinesisAnalyticsApplication(client)
	case "aws_kms_external_key":
		return ListKmsExternalKey(client)
	case "aws_kms_key":
		return ListKmsKey(client)
	case "aws_lambda_event_source_mapping":
		return ListLambdaEventSourceMapping(client)
	case "aws_lambda_function":
		return ListLambdaFunction(client)
	case "aws_launch_configuration":
		return ListLaunchConfiguration(client)
	case "aws_launch_template":
		return ListLaunchTemplate(client)
	case "aws_lb_target_group":
		return ListLbTargetGroup(client)
	case "aws_licensemanager_license_configuration":
		return ListLicensemanagerLicenseConfiguration(client)
	case "aws_lightsail_domain":
		return ListLightsailDomain(client)
	case "aws_lightsail_instance":
		return ListLightsailInstance(client)
	case "aws_lightsail_key_pair":
		return ListLightsailKeyPair(client)
	case "aws_lightsail_static_ip":
		return ListLightsailStaticIp(client)
	case "aws_media_convert_queue":
		return ListMediaConvertQueue(client)
	case "aws_media_package_channel":
		return ListMediaPackageChannel(client)
	case "aws_media_store_container":
		return ListMediaStoreContainer(client)
	case "aws_mq_broker":
		return ListMqBroker(client)
	case "aws_mq_configuration":
		return ListMqConfiguration(client)
	case "aws_msk_cluster":
		return ListMskCluster(client)
	case "aws_msk_configuration":
		return ListMskConfiguration(client)
	case "aws_nat_gateway":
		return ListNatGateway(client)
	case "aws_neptune_event_subscription":
		return ListNeptuneEventSubscription(client)
	case "aws_network_acl":
		return ListNetworkAcl(client)
	case "aws_network_interface":
		return ListNetworkInterface(client)
	case "aws_opsworks_stack":
		return ListOpsworksStack(client)
	case "aws_opsworks_user_profile":
		return ListOpsworksUserProfile(client)
	case "aws_placement_group":
		return ListPlacementGroup(client)
	case "aws_qldb_ledger":
		return ListQldbLedger(client)
	case "aws_rds_global_cluster":
		return ListRdsGlobalCluster(client)
	case "aws_redshift_cluster":
		return ListRedshiftCluster(client)
	case "aws_redshift_event_subscription":
		return ListRedshiftEventSubscription(client)
	case "aws_redshift_snapshot_copy_grant":
		return ListRedshiftSnapshotCopyGrant(client)
	case "aws_redshift_snapshot_schedule":
		return ListRedshiftSnapshotSchedule(client)
	case "aws_route53_health_check":
		return ListRoute53HealthCheck(client)
	case "aws_route53_resolver_endpoint":
		return ListRoute53ResolverEndpoint(client)
	case "aws_route53_resolver_rule":
		return ListRoute53ResolverRule(client)
	case "aws_route53_resolver_rule_association":
		return ListRoute53ResolverRuleAssociation(client)
	case "aws_route53_zone":
		return ListRoute53Zone(client)
	case "aws_route_table":
		return ListRouteTable(client)
	case "aws_s3_bucket":
		return ListS3Bucket(client)
	case "aws_sagemaker_endpoint":
		return ListSagemakerEndpoint(client)
	case "aws_sagemaker_model":
		return ListSagemakerModel(client)
	case "aws_secretsmanager_secret":
		return ListSecretsmanagerSecret(client)
	case "aws_security_group":
		return ListSecurityGroup(client)
	case "aws_service_discovery_service":
		return ListServiceDiscoveryService(client)
	case "aws_servicecatalog_portfolio":
		return ListServicecatalogPortfolio(client)
	case "aws_ses_active_receipt_rule_set":
		return ListSesActiveReceiptRuleSet(client)
	case "aws_ses_configuration_set":
		return ListSesConfigurationSet(client)
	case "aws_ses_receipt_filter":
		return ListSesReceiptFilter(client)
	case "aws_ses_receipt_rule_set":
		return ListSesReceiptRuleSet(client)
	case "aws_ses_template":
		return ListSesTemplate(client)
	case "aws_sfn_activity":
		return ListSfnActivity(client)
	case "aws_sfn_state_machine":
		return ListSfnStateMachine(client)
	case "aws_sns_platform_application":
		return ListSnsPlatformApplication(client)
	case "aws_sns_topic":
		return ListSnsTopic(client)
	case "aws_sns_topic_subscription":
		return ListSnsTopicSubscription(client)
	case "aws_spot_fleet_request":
		return ListSpotFleetRequest(client)
	case "aws_spot_instance_request":
		return ListSpotInstanceRequest(client)
	case "aws_ssm_activation":
		return ListSsmActivation(client)
	case "aws_ssm_association":
		return ListSsmAssociation(client)
	case "aws_ssm_document":
		return ListSsmDocument(client)
	case "aws_ssm_maintenance_window":
		return ListSsmMaintenanceWindow(client)
	case "aws_ssm_parameter":
		return ListSsmParameter(client)
	case "aws_ssm_patch_baseline":
		return ListSsmPatchBaseline(client)
	case "aws_ssm_patch_group":
		return ListSsmPatchGroup(client)
	case "aws_ssm_resource_data_sync":
		return ListSsmResourceDataSync(client)
	case "aws_storagegateway_gateway":
		return ListStoragegatewayGateway(client)
	case "aws_subnet":
		return ListSubnet(client)
	case "aws_transfer_server":
		return ListTransferServer(client)
	case "aws_vpc":
		return ListVpc(client)
	case "aws_vpc_endpoint":
		return ListVpcEndpoint(client)
	case "aws_vpc_endpoint_connection_notification":
		return ListVpcEndpointConnectionNotification(client)
	case "aws_vpc_endpoint_service":
		return ListVpcEndpointService(client)
	case "aws_vpc_peering_connection":
		return ListVpcPeeringConnection(client)
	case "aws_vpn_gateway":
		return ListVpnGateway(client)
	case "aws_waf_byte_match_set":
		return ListWafByteMatchSet(client)
	case "aws_waf_geo_match_set":
		return ListWafGeoMatchSet(client)
	case "aws_waf_ipset":
		return ListWafIpset(client)
	case "aws_waf_rate_based_rule":
		return ListWafRateBasedRule(client)
	case "aws_waf_regex_match_set":
		return ListWafRegexMatchSet(client)
	case "aws_waf_regex_pattern_set":
		return ListWafRegexPatternSet(client)
	case "aws_waf_rule":
		return ListWafRule(client)
	case "aws_waf_rule_group":
		return ListWafRuleGroup(client)
	case "aws_waf_size_constraint_set":
		return ListWafSizeConstraintSet(client)
	case "aws_waf_sql_injection_match_set":
		return ListWafSqlInjectionMatchSet(client)
	case "aws_waf_web_acl":
		return ListWafWebAcl(client)
	case "aws_waf_xss_match_set":
		return ListWafXssMatchSet(client)
	case "aws_wafregional_byte_match_set":
		return ListWafregionalByteMatchSet(client)
	case "aws_wafregional_geo_match_set":
		return ListWafregionalGeoMatchSet(client)
	case "aws_wafregional_ipset":
		return ListWafregionalIpset(client)
	case "aws_wafregional_rate_based_rule":
		return ListWafregionalRateBasedRule(client)
	case "aws_wafregional_regex_match_set":
		return ListWafregionalRegexMatchSet(client)
	case "aws_wafregional_regex_pattern_set":
		return ListWafregionalRegexPatternSet(client)
	case "aws_wafregional_rule":
		return ListWafregionalRule(client)
	case "aws_wafregional_rule_group":
		return ListWafregionalRuleGroup(client)
	case "aws_wafregional_size_constraint_set":
		return ListWafregionalSizeConstraintSet(client)
	case "aws_wafregional_sql_injection_match_set":
		return ListWafregionalSqlInjectionMatchSet(client)
	case "aws_wafregional_web_acl":
		return ListWafregionalWebAcl(client)
	case "aws_wafregional_xss_match_set":
		return ListWafregionalXssMatchSet(client)
	case "aws_wafv2_web_acl_logging_configuration":
		return ListWafv2WebAclLoggingConfiguration(client)
	case "aws_worklink_fleet":
		return ListWorklinkFleet(client)
	case "aws_workspaces_ip_group":
		return ListWorkspacesIpGroup(client)
	default:
		return nil, fmt.Errorf("resource type is not (yet) supported: %s", resourceType)
	}
}
