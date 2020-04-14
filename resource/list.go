// Code is generated. DO NOT EDIT.

package resource

import (
	"fmt"
	"github.com/jckuester/awsls/aws"
)

func ListResources(client *aws.Client) error {
	aws.ListAccessanalyzerAnalyzer(client)
	aws.ListAcmCertificate(client)
	aws.ListAlbListener(client)
	aws.ListAlbListenerRule(client)
	aws.ListAlbTargetGroup(client)
	aws.ListAmi(client)
	aws.ListApiGatewayApiKey(client)
	aws.ListApiGatewayClientCertificate(client)
	aws.ListApiGatewayDomainName(client)
	aws.ListApiGatewayRestApi(client)
	aws.ListApiGatewayUsagePlan(client)
	aws.ListApiGatewayVpcLink(client)
	aws.ListApigatewayv2Api(client)
	aws.ListAppmeshMesh(client)
	aws.ListAppsyncGraphqlApi(client)
	aws.ListAthenaWorkgroup(client)
	aws.ListAutoscalingGroup(client)
	aws.ListBackupPlan(client)
	aws.ListBackupVault(client)
	aws.ListBatchComputeEnvironment(client)
	aws.ListBatchJobDefinition(client)
	aws.ListBatchJobQueue(client)
	aws.ListCloudformationStack(client)
	aws.ListCloudformationStackSet(client)
	aws.ListCloudhsmV2Cluster(client)
	aws.ListCloudwatchDashboard(client)
	aws.ListCloudwatchEventRule(client)
	aws.ListCloudwatchLogDestination(client)
	aws.ListCloudwatchLogGroup(client)
	aws.ListCloudwatchLogResourcePolicy(client)
	aws.ListCodebuildSourceCredential(client)
	aws.ListCodecommitRepository(client)
	aws.ListCodepipelineWebhook(client)
	aws.ListCodestarnotificationsNotificationRule(client)
	aws.ListConfigConfigRule(client)
	aws.ListConfigConfigurationRecorder(client)
	aws.ListConfigDeliveryChannel(client)
	aws.ListCurReportDefinition(client)
	aws.ListCustomerGateway(client)
	aws.ListDatasyncAgent(client)
	aws.ListDatasyncTask(client)
	aws.ListDaxParameterGroup(client)
	aws.ListDaxSubnetGroup(client)
	aws.ListDbEventSubscription(client)
	aws.ListDbInstance(client)
	aws.ListDbParameterGroup(client)
	aws.ListDbSecurityGroup(client)
	aws.ListDbSubnetGroup(client)
	aws.ListDefaultRouteTable(client)
	aws.ListDefaultSecurityGroup(client)
	aws.ListDefaultVpc(client)
	aws.ListDevicefarmProject(client)
	aws.ListDlmLifecyclePolicy(client)
	aws.ListDmsCertificate(client)
	aws.ListDmsEndpoint(client)
	aws.ListDmsReplicationSubnetGroup(client)
	aws.ListDmsReplicationTask(client)
	aws.ListDxConnection(client)
	aws.ListDxHostedPrivateVirtualInterface(client)
	aws.ListDxHostedPublicVirtualInterface(client)
	aws.ListDxHostedTransitVirtualInterface(client)
	aws.ListDxLag(client)
	aws.ListDxPrivateVirtualInterface(client)
	aws.ListDxPublicVirtualInterface(client)
	aws.ListDxTransitVirtualInterface(client)
	aws.ListDynamodbGlobalTable(client)
	aws.ListEbsSnapshot(client)
	aws.ListEbsVolume(client)
	aws.ListEc2CapacityReservation(client)
	aws.ListEc2ClientVpnEndpoint(client)
	aws.ListEc2Fleet(client)
	aws.ListEc2TrafficMirrorFilter(client)
	aws.ListEc2TrafficMirrorSession(client)
	aws.ListEc2TrafficMirrorTarget(client)
	aws.ListEc2TransitGateway(client)
	aws.ListEc2TransitGatewayRouteTable(client)
	aws.ListEc2TransitGatewayVpcAttachment(client)
	aws.ListEcrRepository(client)
	aws.ListEcsCluster(client)
	aws.ListEfsFileSystem(client)
	aws.ListEfsMountTarget(client)
	aws.ListEgressOnlyInternetGateway(client)
	aws.ListEip(client)
	aws.ListElasticBeanstalkApplication(client)
	aws.ListElasticBeanstalkApplicationVersion(client)
	aws.ListElasticBeanstalkEnvironment(client)
	aws.ListElasticacheReplicationGroup(client)
	aws.ListElastictranscoderPipeline(client)
	aws.ListElastictranscoderPreset(client)
	aws.ListElb(client)
	aws.ListEmrSecurityConfiguration(client)
	aws.ListFsxLustreFileSystem(client)
	aws.ListFsxWindowsFileSystem(client)
	aws.ListGameliftAlias(client)
	aws.ListGameliftBuild(client)
	aws.ListGameliftGameSessionQueue(client)
	aws.ListGlobalacceleratorAccelerator(client)
	aws.ListGlueCrawler(client)
	aws.ListGlueJob(client)
	aws.ListGlueSecurityConfiguration(client)
	aws.ListGlueTrigger(client)
	aws.ListIamAccessKey(client)
	aws.ListIamGroup(client)
	aws.ListIamInstanceProfile(client)
	aws.ListIamPolicy(client)
	aws.ListIamRole(client)
	aws.ListIamServerCertificate(client)
	aws.ListIamServiceLinkedRole(client)
	aws.ListIamUser(client)
	aws.ListInternetGateway(client)
	aws.ListIotCertificate(client)
	aws.ListIotPolicy(client)
	aws.ListIotThing(client)
	aws.ListIotThingType(client)
	aws.ListIotTopicRule(client)
	aws.ListKeyPair(client)
	aws.ListKinesisAnalyticsApplication(client)
	aws.ListKmsAlias(client)
	aws.ListKmsExternalKey(client)
	aws.ListKmsKey(client)
	aws.ListLambdaEventSourceMapping(client)
	aws.ListLambdaFunction(client)
	aws.ListLaunchConfiguration(client)
	aws.ListLaunchTemplate(client)
	aws.ListLbListener(client)
	aws.ListLbListenerRule(client)
	aws.ListLbTargetGroup(client)
	aws.ListLicensemanagerLicenseConfiguration(client)
	aws.ListLightsailDomain(client)
	aws.ListLightsailInstance(client)
	aws.ListLightsailKeyPair(client)
	aws.ListLightsailStaticIp(client)
	aws.ListMediaConvertQueue(client)
	aws.ListMediaPackageChannel(client)
	aws.ListMediaStoreContainer(client)
	aws.ListMqBroker(client)
	aws.ListMqConfiguration(client)
	aws.ListMskCluster(client)
	aws.ListMskConfiguration(client)
	aws.ListNatGateway(client)
	aws.ListNeptuneEventSubscription(client)
	aws.ListNetworkAcl(client)
	aws.ListNetworkInterface(client)
	aws.ListOpsworksInstance(client)
	aws.ListOpsworksStack(client)
	aws.ListOpsworksUserProfile(client)
	aws.ListPlacementGroup(client)
	aws.ListQldbLedger(client)
	aws.ListRdsGlobalCluster(client)
	aws.ListRedshiftCluster(client)
	aws.ListRedshiftEventSubscription(client)
	aws.ListRedshiftSnapshotCopyGrant(client)
	aws.ListRedshiftSnapshotSchedule(client)
	aws.ListRoute53HealthCheck(client)
	aws.ListRoute53ResolverEndpoint(client)
	aws.ListRoute53ResolverRule(client)
	aws.ListRoute53ResolverRuleAssociation(client)
	aws.ListRoute53Zone(client)
	aws.ListRouteTable(client)
	aws.ListS3Bucket(client)
	aws.ListSagemakerEndpoint(client)
	aws.ListSagemakerModel(client)
	aws.ListSecretsmanagerSecret(client)
	aws.ListSecurityGroup(client)
	aws.ListServiceDiscoveryService(client)
	aws.ListServicecatalogPortfolio(client)
	aws.ListSesActiveReceiptRuleSet(client)
	aws.ListSesConfigurationSet(client)
	aws.ListSesReceiptFilter(client)
	aws.ListSesReceiptRuleSet(client)
	aws.ListSesTemplate(client)
	aws.ListSfnActivity(client)
	aws.ListSfnStateMachine(client)
	aws.ListSnsPlatformApplication(client)
	aws.ListSnsTopic(client)
	aws.ListSnsTopicSubscription(client)
	aws.ListSpotFleetRequest(client)
	aws.ListSpotInstanceRequest(client)
	aws.ListSsmActivation(client)
	aws.ListSsmAssociation(client)
	aws.ListSsmDocument(client)
	aws.ListSsmMaintenanceWindow(client)
	aws.ListSsmPatchBaseline(client)
	aws.ListSsmPatchGroup(client)
	aws.ListStoragegatewayGateway(client)
	aws.ListSubnet(client)
	aws.ListTransferServer(client)
	aws.ListVpc(client)
	aws.ListVpcEndpoint(client)
	aws.ListVpcEndpointConnectionNotification(client)
	aws.ListVpcEndpointService(client)
	aws.ListVpcPeeringConnection(client)
	aws.ListVpnGateway(client)
	aws.ListWafByteMatchSet(client)
	aws.ListWafGeoMatchSet(client)
	aws.ListWafIpset(client)
	aws.ListWafRateBasedRule(client)
	aws.ListWafRegexMatchSet(client)
	aws.ListWafRegexPatternSet(client)
	aws.ListWafRule(client)
	aws.ListWafRuleGroup(client)
	aws.ListWafSizeConstraintSet(client)
	aws.ListWafSqlInjectionMatchSet(client)
	aws.ListWafWebAcl(client)
	aws.ListWafXssMatchSet(client)
	aws.ListWafregionalByteMatchSet(client)
	aws.ListWafregionalGeoMatchSet(client)
	aws.ListWafregionalIpset(client)
	aws.ListWafregionalRateBasedRule(client)
	aws.ListWafregionalRegexMatchSet(client)
	aws.ListWafregionalRegexPatternSet(client)
	aws.ListWafregionalRule(client)
	aws.ListWafregionalRuleGroup(client)
	aws.ListWafregionalSizeConstraintSet(client)
	aws.ListWafregionalSqlInjectionMatchSet(client)
	aws.ListWafregionalWebAcl(client)
	aws.ListWafregionalXssMatchSet(client)
	aws.ListWorklinkFleet(client)
	aws.ListWorkspacesIpGroup(client)

	return nil
}

func ListResourcesByType(client *aws.Client, resourceType string) error {
	switch resourceType {
	case "aws_accessanalyzer_analyzer":
		return aws.ListAccessanalyzerAnalyzer(client)
	case "aws_acm_certificate":
		return aws.ListAcmCertificate(client)
	case "aws_alb_listener":
		return aws.ListAlbListener(client)
	case "aws_alb_listener_rule":
		return aws.ListAlbListenerRule(client)
	case "aws_alb_target_group":
		return aws.ListAlbTargetGroup(client)
	case "aws_ami":
		return aws.ListAmi(client)
	case "aws_api_gateway_api_key":
		return aws.ListApiGatewayApiKey(client)
	case "aws_api_gateway_client_certificate":
		return aws.ListApiGatewayClientCertificate(client)
	case "aws_api_gateway_domain_name":
		return aws.ListApiGatewayDomainName(client)
	case "aws_api_gateway_rest_api":
		return aws.ListApiGatewayRestApi(client)
	case "aws_api_gateway_usage_plan":
		return aws.ListApiGatewayUsagePlan(client)
	case "aws_api_gateway_vpc_link":
		return aws.ListApiGatewayVpcLink(client)
	case "aws_apigatewayv2_api":
		return aws.ListApigatewayv2Api(client)
	case "aws_appmesh_mesh":
		return aws.ListAppmeshMesh(client)
	case "aws_appsync_graphql_api":
		return aws.ListAppsyncGraphqlApi(client)
	case "aws_athena_workgroup":
		return aws.ListAthenaWorkgroup(client)
	case "aws_autoscaling_group":
		return aws.ListAutoscalingGroup(client)
	case "aws_backup_plan":
		return aws.ListBackupPlan(client)
	case "aws_backup_vault":
		return aws.ListBackupVault(client)
	case "aws_batch_compute_environment":
		return aws.ListBatchComputeEnvironment(client)
	case "aws_batch_job_definition":
		return aws.ListBatchJobDefinition(client)
	case "aws_batch_job_queue":
		return aws.ListBatchJobQueue(client)
	case "aws_cloudformation_stack":
		return aws.ListCloudformationStack(client)
	case "aws_cloudformation_stack_set":
		return aws.ListCloudformationStackSet(client)
	case "aws_cloudhsm_v2_cluster":
		return aws.ListCloudhsmV2Cluster(client)
	case "aws_cloudwatch_dashboard":
		return aws.ListCloudwatchDashboard(client)
	case "aws_cloudwatch_event_rule":
		return aws.ListCloudwatchEventRule(client)
	case "aws_cloudwatch_log_destination":
		return aws.ListCloudwatchLogDestination(client)
	case "aws_cloudwatch_log_group":
		return aws.ListCloudwatchLogGroup(client)
	case "aws_cloudwatch_log_resource_policy":
		return aws.ListCloudwatchLogResourcePolicy(client)
	case "aws_codebuild_source_credential":
		return aws.ListCodebuildSourceCredential(client)
	case "aws_codecommit_repository":
		return aws.ListCodecommitRepository(client)
	case "aws_codepipeline_webhook":
		return aws.ListCodepipelineWebhook(client)
	case "aws_codestarnotifications_notification_rule":
		return aws.ListCodestarnotificationsNotificationRule(client)
	case "aws_config_config_rule":
		return aws.ListConfigConfigRule(client)
	case "aws_config_configuration_recorder":
		return aws.ListConfigConfigurationRecorder(client)
	case "aws_config_delivery_channel":
		return aws.ListConfigDeliveryChannel(client)
	case "aws_cur_report_definition":
		return aws.ListCurReportDefinition(client)
	case "aws_customer_gateway":
		return aws.ListCustomerGateway(client)
	case "aws_datasync_agent":
		return aws.ListDatasyncAgent(client)
	case "aws_datasync_task":
		return aws.ListDatasyncTask(client)
	case "aws_dax_parameter_group":
		return aws.ListDaxParameterGroup(client)
	case "aws_dax_subnet_group":
		return aws.ListDaxSubnetGroup(client)
	case "aws_db_event_subscription":
		return aws.ListDbEventSubscription(client)
	case "aws_db_instance":
		return aws.ListDbInstance(client)
	case "aws_db_parameter_group":
		return aws.ListDbParameterGroup(client)
	case "aws_db_security_group":
		return aws.ListDbSecurityGroup(client)
	case "aws_db_subnet_group":
		return aws.ListDbSubnetGroup(client)
	case "aws_default_route_table":
		return aws.ListDefaultRouteTable(client)
	case "aws_default_security_group":
		return aws.ListDefaultSecurityGroup(client)
	case "aws_default_vpc":
		return aws.ListDefaultVpc(client)
	case "aws_devicefarm_project":
		return aws.ListDevicefarmProject(client)
	case "aws_dlm_lifecycle_policy":
		return aws.ListDlmLifecyclePolicy(client)
	case "aws_dms_certificate":
		return aws.ListDmsCertificate(client)
	case "aws_dms_endpoint":
		return aws.ListDmsEndpoint(client)
	case "aws_dms_replication_subnet_group":
		return aws.ListDmsReplicationSubnetGroup(client)
	case "aws_dms_replication_task":
		return aws.ListDmsReplicationTask(client)
	case "aws_dx_connection":
		return aws.ListDxConnection(client)
	case "aws_dx_hosted_private_virtual_interface":
		return aws.ListDxHostedPrivateVirtualInterface(client)
	case "aws_dx_hosted_public_virtual_interface":
		return aws.ListDxHostedPublicVirtualInterface(client)
	case "aws_dx_hosted_transit_virtual_interface":
		return aws.ListDxHostedTransitVirtualInterface(client)
	case "aws_dx_lag":
		return aws.ListDxLag(client)
	case "aws_dx_private_virtual_interface":
		return aws.ListDxPrivateVirtualInterface(client)
	case "aws_dx_public_virtual_interface":
		return aws.ListDxPublicVirtualInterface(client)
	case "aws_dx_transit_virtual_interface":
		return aws.ListDxTransitVirtualInterface(client)
	case "aws_dynamodb_global_table":
		return aws.ListDynamodbGlobalTable(client)
	case "aws_ebs_snapshot":
		return aws.ListEbsSnapshot(client)
	case "aws_ebs_volume":
		return aws.ListEbsVolume(client)
	case "aws_ec2_capacity_reservation":
		return aws.ListEc2CapacityReservation(client)
	case "aws_ec2_client_vpn_endpoint":
		return aws.ListEc2ClientVpnEndpoint(client)
	case "aws_ec2_fleet":
		return aws.ListEc2Fleet(client)
	case "aws_ec2_traffic_mirror_filter":
		return aws.ListEc2TrafficMirrorFilter(client)
	case "aws_ec2_traffic_mirror_session":
		return aws.ListEc2TrafficMirrorSession(client)
	case "aws_ec2_traffic_mirror_target":
		return aws.ListEc2TrafficMirrorTarget(client)
	case "aws_ec2_transit_gateway":
		return aws.ListEc2TransitGateway(client)
	case "aws_ec2_transit_gateway_route_table":
		return aws.ListEc2TransitGatewayRouteTable(client)
	case "aws_ec2_transit_gateway_vpc_attachment":
		return aws.ListEc2TransitGatewayVpcAttachment(client)
	case "aws_ecr_repository":
		return aws.ListEcrRepository(client)
	case "aws_ecs_cluster":
		return aws.ListEcsCluster(client)
	case "aws_efs_file_system":
		return aws.ListEfsFileSystem(client)
	case "aws_efs_mount_target":
		return aws.ListEfsMountTarget(client)
	case "aws_egress_only_internet_gateway":
		return aws.ListEgressOnlyInternetGateway(client)
	case "aws_eip":
		return aws.ListEip(client)
	case "aws_elastic_beanstalk_application":
		return aws.ListElasticBeanstalkApplication(client)
	case "aws_elastic_beanstalk_application_version":
		return aws.ListElasticBeanstalkApplicationVersion(client)
	case "aws_elastic_beanstalk_environment":
		return aws.ListElasticBeanstalkEnvironment(client)
	case "aws_elasticache_replication_group":
		return aws.ListElasticacheReplicationGroup(client)
	case "aws_elastictranscoder_pipeline":
		return aws.ListElastictranscoderPipeline(client)
	case "aws_elastictranscoder_preset":
		return aws.ListElastictranscoderPreset(client)
	case "aws_elb":
		return aws.ListElb(client)
	case "aws_emr_security_configuration":
		return aws.ListEmrSecurityConfiguration(client)
	case "aws_fsx_lustre_file_system":
		return aws.ListFsxLustreFileSystem(client)
	case "aws_fsx_windows_file_system":
		return aws.ListFsxWindowsFileSystem(client)
	case "aws_gamelift_alias":
		return aws.ListGameliftAlias(client)
	case "aws_gamelift_build":
		return aws.ListGameliftBuild(client)
	case "aws_gamelift_game_session_queue":
		return aws.ListGameliftGameSessionQueue(client)
	case "aws_globalaccelerator_accelerator":
		return aws.ListGlobalacceleratorAccelerator(client)
	case "aws_glue_crawler":
		return aws.ListGlueCrawler(client)
	case "aws_glue_job":
		return aws.ListGlueJob(client)
	case "aws_glue_security_configuration":
		return aws.ListGlueSecurityConfiguration(client)
	case "aws_glue_trigger":
		return aws.ListGlueTrigger(client)
	case "aws_iam_access_key":
		return aws.ListIamAccessKey(client)
	case "aws_iam_group":
		return aws.ListIamGroup(client)
	case "aws_iam_instance_profile":
		return aws.ListIamInstanceProfile(client)
	case "aws_iam_policy":
		return aws.ListIamPolicy(client)
	case "aws_iam_role":
		return aws.ListIamRole(client)
	case "aws_iam_server_certificate":
		return aws.ListIamServerCertificate(client)
	case "aws_iam_service_linked_role":
		return aws.ListIamServiceLinkedRole(client)
	case "aws_iam_user":
		return aws.ListIamUser(client)
	case "aws_internet_gateway":
		return aws.ListInternetGateway(client)
	case "aws_iot_certificate":
		return aws.ListIotCertificate(client)
	case "aws_iot_policy":
		return aws.ListIotPolicy(client)
	case "aws_iot_thing":
		return aws.ListIotThing(client)
	case "aws_iot_thing_type":
		return aws.ListIotThingType(client)
	case "aws_iot_topic_rule":
		return aws.ListIotTopicRule(client)
	case "aws_key_pair":
		return aws.ListKeyPair(client)
	case "aws_kinesis_analytics_application":
		return aws.ListKinesisAnalyticsApplication(client)
	case "aws_kms_alias":
		return aws.ListKmsAlias(client)
	case "aws_kms_external_key":
		return aws.ListKmsExternalKey(client)
	case "aws_kms_key":
		return aws.ListKmsKey(client)
	case "aws_lambda_event_source_mapping":
		return aws.ListLambdaEventSourceMapping(client)
	case "aws_lambda_function":
		return aws.ListLambdaFunction(client)
	case "aws_launch_configuration":
		return aws.ListLaunchConfiguration(client)
	case "aws_launch_template":
		return aws.ListLaunchTemplate(client)
	case "aws_lb_listener":
		return aws.ListLbListener(client)
	case "aws_lb_listener_rule":
		return aws.ListLbListenerRule(client)
	case "aws_lb_target_group":
		return aws.ListLbTargetGroup(client)
	case "aws_licensemanager_license_configuration":
		return aws.ListLicensemanagerLicenseConfiguration(client)
	case "aws_lightsail_domain":
		return aws.ListLightsailDomain(client)
	case "aws_lightsail_instance":
		return aws.ListLightsailInstance(client)
	case "aws_lightsail_key_pair":
		return aws.ListLightsailKeyPair(client)
	case "aws_lightsail_static_ip":
		return aws.ListLightsailStaticIp(client)
	case "aws_media_convert_queue":
		return aws.ListMediaConvertQueue(client)
	case "aws_media_package_channel":
		return aws.ListMediaPackageChannel(client)
	case "aws_media_store_container":
		return aws.ListMediaStoreContainer(client)
	case "aws_mq_broker":
		return aws.ListMqBroker(client)
	case "aws_mq_configuration":
		return aws.ListMqConfiguration(client)
	case "aws_msk_cluster":
		return aws.ListMskCluster(client)
	case "aws_msk_configuration":
		return aws.ListMskConfiguration(client)
	case "aws_nat_gateway":
		return aws.ListNatGateway(client)
	case "aws_neptune_event_subscription":
		return aws.ListNeptuneEventSubscription(client)
	case "aws_network_acl":
		return aws.ListNetworkAcl(client)
	case "aws_network_interface":
		return aws.ListNetworkInterface(client)
	case "aws_opsworks_instance":
		return aws.ListOpsworksInstance(client)
	case "aws_opsworks_stack":
		return aws.ListOpsworksStack(client)
	case "aws_opsworks_user_profile":
		return aws.ListOpsworksUserProfile(client)
	case "aws_placement_group":
		return aws.ListPlacementGroup(client)
	case "aws_qldb_ledger":
		return aws.ListQldbLedger(client)
	case "aws_rds_global_cluster":
		return aws.ListRdsGlobalCluster(client)
	case "aws_redshift_cluster":
		return aws.ListRedshiftCluster(client)
	case "aws_redshift_event_subscription":
		return aws.ListRedshiftEventSubscription(client)
	case "aws_redshift_snapshot_copy_grant":
		return aws.ListRedshiftSnapshotCopyGrant(client)
	case "aws_redshift_snapshot_schedule":
		return aws.ListRedshiftSnapshotSchedule(client)
	case "aws_route53_health_check":
		return aws.ListRoute53HealthCheck(client)
	case "aws_route53_resolver_endpoint":
		return aws.ListRoute53ResolverEndpoint(client)
	case "aws_route53_resolver_rule":
		return aws.ListRoute53ResolverRule(client)
	case "aws_route53_resolver_rule_association":
		return aws.ListRoute53ResolverRuleAssociation(client)
	case "aws_route53_zone":
		return aws.ListRoute53Zone(client)
	case "aws_route_table":
		return aws.ListRouteTable(client)
	case "aws_s3_bucket":
		return aws.ListS3Bucket(client)
	case "aws_sagemaker_endpoint":
		return aws.ListSagemakerEndpoint(client)
	case "aws_sagemaker_model":
		return aws.ListSagemakerModel(client)
	case "aws_secretsmanager_secret":
		return aws.ListSecretsmanagerSecret(client)
	case "aws_security_group":
		return aws.ListSecurityGroup(client)
	case "aws_service_discovery_service":
		return aws.ListServiceDiscoveryService(client)
	case "aws_servicecatalog_portfolio":
		return aws.ListServicecatalogPortfolio(client)
	case "aws_ses_active_receipt_rule_set":
		return aws.ListSesActiveReceiptRuleSet(client)
	case "aws_ses_configuration_set":
		return aws.ListSesConfigurationSet(client)
	case "aws_ses_receipt_filter":
		return aws.ListSesReceiptFilter(client)
	case "aws_ses_receipt_rule_set":
		return aws.ListSesReceiptRuleSet(client)
	case "aws_ses_template":
		return aws.ListSesTemplate(client)
	case "aws_sfn_activity":
		return aws.ListSfnActivity(client)
	case "aws_sfn_state_machine":
		return aws.ListSfnStateMachine(client)
	case "aws_sns_platform_application":
		return aws.ListSnsPlatformApplication(client)
	case "aws_sns_topic":
		return aws.ListSnsTopic(client)
	case "aws_sns_topic_subscription":
		return aws.ListSnsTopicSubscription(client)
	case "aws_spot_fleet_request":
		return aws.ListSpotFleetRequest(client)
	case "aws_spot_instance_request":
		return aws.ListSpotInstanceRequest(client)
	case "aws_ssm_activation":
		return aws.ListSsmActivation(client)
	case "aws_ssm_association":
		return aws.ListSsmAssociation(client)
	case "aws_ssm_document":
		return aws.ListSsmDocument(client)
	case "aws_ssm_maintenance_window":
		return aws.ListSsmMaintenanceWindow(client)
	case "aws_ssm_patch_baseline":
		return aws.ListSsmPatchBaseline(client)
	case "aws_ssm_patch_group":
		return aws.ListSsmPatchGroup(client)
	case "aws_storagegateway_gateway":
		return aws.ListStoragegatewayGateway(client)
	case "aws_subnet":
		return aws.ListSubnet(client)
	case "aws_transfer_server":
		return aws.ListTransferServer(client)
	case "aws_vpc":
		return aws.ListVpc(client)
	case "aws_vpc_endpoint":
		return aws.ListVpcEndpoint(client)
	case "aws_vpc_endpoint_connection_notification":
		return aws.ListVpcEndpointConnectionNotification(client)
	case "aws_vpc_endpoint_service":
		return aws.ListVpcEndpointService(client)
	case "aws_vpc_peering_connection":
		return aws.ListVpcPeeringConnection(client)
	case "aws_vpn_gateway":
		return aws.ListVpnGateway(client)
	case "aws_waf_byte_match_set":
		return aws.ListWafByteMatchSet(client)
	case "aws_waf_geo_match_set":
		return aws.ListWafGeoMatchSet(client)
	case "aws_waf_ipset":
		return aws.ListWafIpset(client)
	case "aws_waf_rate_based_rule":
		return aws.ListWafRateBasedRule(client)
	case "aws_waf_regex_match_set":
		return aws.ListWafRegexMatchSet(client)
	case "aws_waf_regex_pattern_set":
		return aws.ListWafRegexPatternSet(client)
	case "aws_waf_rule":
		return aws.ListWafRule(client)
	case "aws_waf_rule_group":
		return aws.ListWafRuleGroup(client)
	case "aws_waf_size_constraint_set":
		return aws.ListWafSizeConstraintSet(client)
	case "aws_waf_sql_injection_match_set":
		return aws.ListWafSqlInjectionMatchSet(client)
	case "aws_waf_web_acl":
		return aws.ListWafWebAcl(client)
	case "aws_waf_xss_match_set":
		return aws.ListWafXssMatchSet(client)
	case "aws_wafregional_byte_match_set":
		return aws.ListWafregionalByteMatchSet(client)
	case "aws_wafregional_geo_match_set":
		return aws.ListWafregionalGeoMatchSet(client)
	case "aws_wafregional_ipset":
		return aws.ListWafregionalIpset(client)
	case "aws_wafregional_rate_based_rule":
		return aws.ListWafregionalRateBasedRule(client)
	case "aws_wafregional_regex_match_set":
		return aws.ListWafregionalRegexMatchSet(client)
	case "aws_wafregional_regex_pattern_set":
		return aws.ListWafregionalRegexPatternSet(client)
	case "aws_wafregional_rule":
		return aws.ListWafregionalRule(client)
	case "aws_wafregional_rule_group":
		return aws.ListWafregionalRuleGroup(client)
	case "aws_wafregional_size_constraint_set":
		return aws.ListWafregionalSizeConstraintSet(client)
	case "aws_wafregional_sql_injection_match_set":
		return aws.ListWafregionalSqlInjectionMatchSet(client)
	case "aws_wafregional_web_acl":
		return aws.ListWafregionalWebAcl(client)
	case "aws_wafregional_xss_match_set":
		return aws.ListWafregionalXssMatchSet(client)
	case "aws_worklink_fleet":
		return aws.ListWorklinkFleet(client)
	case "aws_workspaces_ip_group":
		return aws.ListWorkspacesIpGroup(client)
	default:
		return fmt.Errorf("resource type is not (yet) supported: %s", resourceType)
	}
}
