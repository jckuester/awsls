// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"

	"github.com/jckuester/awstools-lib/aws"
	"github.com/jckuester/awstools-lib/terraform"
)

func ListResourcesByType(ctx context.Context, client *aws.Client, resourceType string) ([]terraform.Resource, error) {
	switch resourceType {
	case "aws_accessanalyzer_analyzer":
		return ListAccessanalyzerAnalyzer(ctx, client)
	case "aws_acm_certificate":
		return ListAcmCertificate(ctx, client)
	case "aws_alb_target_group":
		return ListAlbTargetGroup(ctx, client)
	case "aws_ami":
		return ListAmi(ctx, client)
	case "aws_api_gateway_api_key":
		return ListApiGatewayApiKey(ctx, client)
	case "aws_api_gateway_client_certificate":
		return ListApiGatewayClientCertificate(ctx, client)
	case "aws_api_gateway_domain_name":
		return ListApiGatewayDomainName(ctx, client)
	case "aws_api_gateway_rest_api":
		return ListApiGatewayRestApi(ctx, client)
	case "aws_api_gateway_usage_plan":
		return ListApiGatewayUsagePlan(ctx, client)
	case "aws_api_gateway_vpc_link":
		return ListApiGatewayVpcLink(ctx, client)
	case "aws_apigatewayv2_api":
		return ListApigatewayv2Api(ctx, client)
	case "aws_apigatewayv2_domain_name":
		return ListApigatewayv2DomainName(ctx, client)
	case "aws_apigatewayv2_vpc_link":
		return ListApigatewayv2VpcLink(ctx, client)
	case "aws_appmesh_mesh":
		return ListAppmeshMesh(ctx, client)
	case "aws_appsync_graphql_api":
		return ListAppsyncGraphqlApi(ctx, client)
	case "aws_athena_named_query":
		return ListAthenaNamedQuery(ctx, client)
	case "aws_athena_workgroup":
		return ListAthenaWorkgroup(ctx, client)
	case "aws_autoscaling_group":
		return ListAutoscalingGroup(ctx, client)
	case "aws_backup_plan":
		return ListBackupPlan(ctx, client)
	case "aws_backup_vault":
		return ListBackupVault(ctx, client)
	case "aws_batch_compute_environment":
		return ListBatchComputeEnvironment(ctx, client)
	case "aws_batch_job_definition":
		return ListBatchJobDefinition(ctx, client)
	case "aws_batch_job_queue":
		return ListBatchJobQueue(ctx, client)
	case "aws_cloudformation_stack":
		return ListCloudformationStack(ctx, client)
	case "aws_cloudformation_stack_set":
		return ListCloudformationStackSet(ctx, client)
	case "aws_cloudhsm_v2_cluster":
		return ListCloudhsmV2Cluster(ctx, client)
	case "aws_cloudtrail":
		return ListCloudtrail(ctx, client)
	case "aws_cloudwatch_dashboard":
		return ListCloudwatchDashboard(ctx, client)
	case "aws_cloudwatch_event_archive":
		return ListCloudwatchEventArchive(ctx, client)
	case "aws_cloudwatch_event_bus":
		return ListCloudwatchEventBus(ctx, client)
	case "aws_cloudwatch_log_destination":
		return ListCloudwatchLogDestination(ctx, client)
	case "aws_cloudwatch_log_group":
		return ListCloudwatchLogGroup(ctx, client)
	case "aws_cloudwatch_log_resource_policy":
		return ListCloudwatchLogResourcePolicy(ctx, client)
	case "aws_codeartifact_domain":
		return ListCodeartifactDomain(ctx, client)
	case "aws_codeartifact_repository":
		return ListCodeartifactRepository(ctx, client)
	case "aws_codebuild_project":
		return ListCodebuildProject(ctx, client)
	case "aws_codebuild_report_group":
		return ListCodebuildReportGroup(ctx, client)
	case "aws_codebuild_source_credential":
		return ListCodebuildSourceCredential(ctx, client)
	case "aws_codecommit_repository":
		return ListCodecommitRepository(ctx, client)
	case "aws_codedeploy_deployment_config":
		return ListCodedeployDeploymentConfig(ctx, client)
	case "aws_codepipeline_webhook":
		return ListCodepipelineWebhook(ctx, client)
	case "aws_codestarconnections_connection":
		return ListCodestarconnectionsConnection(ctx, client)
	case "aws_codestarnotifications_notification_rule":
		return ListCodestarnotificationsNotificationRule(ctx, client)
	case "aws_config_config_rule":
		return ListConfigConfigRule(ctx, client)
	case "aws_config_configuration_recorder":
		return ListConfigConfigurationRecorder(ctx, client)
	case "aws_config_conformance_pack":
		return ListConfigConformancePack(ctx, client)
	case "aws_config_delivery_channel":
		return ListConfigDeliveryChannel(ctx, client)
	case "aws_cur_report_definition":
		return ListCurReportDefinition(ctx, client)
	case "aws_datasync_agent":
		return ListDatasyncAgent(ctx, client)
	case "aws_datasync_task":
		return ListDatasyncTask(ctx, client)
	case "aws_dax_parameter_group":
		return ListDaxParameterGroup(ctx, client)
	case "aws_dax_subnet_group":
		return ListDaxSubnetGroup(ctx, client)
	case "aws_db_event_subscription":
		return ListDbEventSubscription(ctx, client)
	case "aws_db_instance":
		return ListDbInstance(ctx, client)
	case "aws_db_parameter_group":
		return ListDbParameterGroup(ctx, client)
	case "aws_db_proxy":
		return ListDbProxy(ctx, client)
	case "aws_db_security_group":
		return ListDbSecurityGroup(ctx, client)
	case "aws_db_snapshot":
		return ListDbSnapshot(ctx, client)
	case "aws_db_subnet_group":
		return ListDbSubnetGroup(ctx, client)
	case "aws_devicefarm_project":
		return ListDevicefarmProject(ctx, client)
	case "aws_dlm_lifecycle_policy":
		return ListDlmLifecyclePolicy(ctx, client)
	case "aws_dms_certificate":
		return ListDmsCertificate(ctx, client)
	case "aws_dms_endpoint":
		return ListDmsEndpoint(ctx, client)
	case "aws_dms_replication_subnet_group":
		return ListDmsReplicationSubnetGroup(ctx, client)
	case "aws_dms_replication_task":
		return ListDmsReplicationTask(ctx, client)
	case "aws_dx_connection":
		return ListDxConnection(ctx, client)
	case "aws_dx_hosted_private_virtual_interface":
		return ListDxHostedPrivateVirtualInterface(ctx, client)
	case "aws_dx_hosted_public_virtual_interface":
		return ListDxHostedPublicVirtualInterface(ctx, client)
	case "aws_dx_hosted_transit_virtual_interface":
		return ListDxHostedTransitVirtualInterface(ctx, client)
	case "aws_dx_lag":
		return ListDxLag(ctx, client)
	case "aws_dx_private_virtual_interface":
		return ListDxPrivateVirtualInterface(ctx, client)
	case "aws_dx_public_virtual_interface":
		return ListDxPublicVirtualInterface(ctx, client)
	case "aws_dx_transit_virtual_interface":
		return ListDxTransitVirtualInterface(ctx, client)
	case "aws_dynamodb_global_table":
		return ListDynamodbGlobalTable(ctx, client)
	case "aws_dynamodb_table":
		return ListDynamodbTable(ctx, client)
	case "aws_ebs_snapshot":
		return ListEbsSnapshot(ctx, client)
	case "aws_ebs_volume":
		return ListEbsVolume(ctx, client)
	case "aws_ec2_capacity_reservation":
		return ListEc2CapacityReservation(ctx, client)
	case "aws_ec2_carrier_gateway":
		return ListEc2CarrierGateway(ctx, client)
	case "aws_ec2_client_vpn_endpoint":
		return ListEc2ClientVpnEndpoint(ctx, client)
	case "aws_ec2_fleet":
		return ListEc2Fleet(ctx, client)
	case "aws_ec2_local_gateway_route_table_vpc_association":
		return ListEc2LocalGatewayRouteTableVpcAssociation(ctx, client)
	case "aws_ec2_managed_prefix_list":
		return ListEc2ManagedPrefixList(ctx, client)
	case "aws_ec2_traffic_mirror_filter":
		return ListEc2TrafficMirrorFilter(ctx, client)
	case "aws_ec2_traffic_mirror_session":
		return ListEc2TrafficMirrorSession(ctx, client)
	case "aws_ec2_traffic_mirror_target":
		return ListEc2TrafficMirrorTarget(ctx, client)
	case "aws_ec2_transit_gateway":
		return ListEc2TransitGateway(ctx, client)
	case "aws_ec2_transit_gateway_peering_attachment":
		return ListEc2TransitGatewayPeeringAttachment(ctx, client)
	case "aws_ec2_transit_gateway_route_table":
		return ListEc2TransitGatewayRouteTable(ctx, client)
	case "aws_ec2_transit_gateway_vpc_attachment":
		return ListEc2TransitGatewayVpcAttachment(ctx, client)
	case "aws_ecr_repository":
		return ListEcrRepository(ctx, client)
	case "aws_ecrpublic_repository":
		return ListEcrpublicRepository(ctx, client)
	case "aws_ecs_cluster":
		return ListEcsCluster(ctx, client)
	case "aws_ecs_task_definition":
		return ListEcsTaskDefinition(ctx, client)
	case "aws_efs_access_point":
		return ListEfsAccessPoint(ctx, client)
	case "aws_efs_file_system":
		return ListEfsFileSystem(ctx, client)
	case "aws_egress_only_internet_gateway":
		return ListEgressOnlyInternetGateway(ctx, client)
	case "aws_eip":
		return ListEip(ctx, client)
	case "aws_eks_cluster":
		return ListEksCluster(ctx, client)
	case "aws_elastic_beanstalk_application":
		return ListElasticBeanstalkApplication(ctx, client)
	case "aws_elastic_beanstalk_application_version":
		return ListElasticBeanstalkApplicationVersion(ctx, client)
	case "aws_elastic_beanstalk_environment":
		return ListElasticBeanstalkEnvironment(ctx, client)
	case "aws_elasticache_global_replication_group":
		return ListElasticacheGlobalReplicationGroup(ctx, client)
	case "aws_elasticache_replication_group":
		return ListElasticacheReplicationGroup(ctx, client)
	case "aws_elastictranscoder_pipeline":
		return ListElastictranscoderPipeline(ctx, client)
	case "aws_elastictranscoder_preset":
		return ListElastictranscoderPreset(ctx, client)
	case "aws_elb":
		return ListElb(ctx, client)
	case "aws_emr_security_configuration":
		return ListEmrSecurityConfiguration(ctx, client)
	case "aws_fms_policy":
		return ListFmsPolicy(ctx, client)
	case "aws_fsx_lustre_file_system":
		return ListFsxLustreFileSystem(ctx, client)
	case "aws_fsx_windows_file_system":
		return ListFsxWindowsFileSystem(ctx, client)
	case "aws_gamelift_alias":
		return ListGameliftAlias(ctx, client)
	case "aws_gamelift_build":
		return ListGameliftBuild(ctx, client)
	case "aws_gamelift_fleet":
		return ListGameliftFleet(ctx, client)
	case "aws_gamelift_game_session_queue":
		return ListGameliftGameSessionQueue(ctx, client)
	case "aws_globalaccelerator_accelerator":
		return ListGlobalacceleratorAccelerator(ctx, client)
	case "aws_glue_crawler":
		return ListGlueCrawler(ctx, client)
	case "aws_glue_dev_endpoint":
		return ListGlueDevEndpoint(ctx, client)
	case "aws_glue_job":
		return ListGlueJob(ctx, client)
	case "aws_glue_ml_transform":
		return ListGlueMlTransform(ctx, client)
	case "aws_glue_registry":
		return ListGlueRegistry(ctx, client)
	case "aws_glue_schema":
		return ListGlueSchema(ctx, client)
	case "aws_glue_security_configuration":
		return ListGlueSecurityConfiguration(ctx, client)
	case "aws_glue_trigger":
		return ListGlueTrigger(ctx, client)
	case "aws_glue_workflow":
		return ListGlueWorkflow(ctx, client)
	case "aws_guardduty_detector":
		return ListGuarddutyDetector(ctx, client)
	case "aws_iam_access_key":
		return ListIamAccessKey(ctx, client)
	case "aws_iam_account_alias":
		return ListIamAccountAlias(ctx, client)
	case "aws_iam_group":
		return ListIamGroup(ctx, client)
	case "aws_iam_instance_profile":
		return ListIamInstanceProfile(ctx, client)
	case "aws_iam_policy":
		return ListIamPolicy(ctx, client)
	case "aws_iam_role":
		return ListIamRole(ctx, client)
	case "aws_iam_server_certificate":
		return ListIamServerCertificate(ctx, client)
	case "aws_iam_service_linked_role":
		return ListIamServiceLinkedRole(ctx, client)
	case "aws_iam_user":
		return ListIamUser(ctx, client)
	case "aws_imagebuilder_component":
		return ListImagebuilderComponent(ctx, client)
	case "aws_imagebuilder_distribution_configuration":
		return ListImagebuilderDistributionConfiguration(ctx, client)
	case "aws_imagebuilder_image":
		return ListImagebuilderImage(ctx, client)
	case "aws_imagebuilder_image_pipeline":
		return ListImagebuilderImagePipeline(ctx, client)
	case "aws_imagebuilder_image_recipe":
		return ListImagebuilderImageRecipe(ctx, client)
	case "aws_imagebuilder_infrastructure_configuration":
		return ListImagebuilderInfrastructureConfiguration(ctx, client)
	case "aws_instance":
		return ListInstance(ctx, client)
	case "aws_internet_gateway":
		return ListInternetGateway(ctx, client)
	case "aws_iot_certificate":
		return ListIotCertificate(ctx, client)
	case "aws_iot_policy":
		return ListIotPolicy(ctx, client)
	case "aws_iot_role_alias":
		return ListIotRoleAlias(ctx, client)
	case "aws_iot_thing":
		return ListIotThing(ctx, client)
	case "aws_iot_thing_type":
		return ListIotThingType(ctx, client)
	case "aws_iot_topic_rule":
		return ListIotTopicRule(ctx, client)
	case "aws_key_pair":
		return ListKeyPair(ctx, client)
	case "aws_kinesis_analytics_application":
		return ListKinesisAnalyticsApplication(ctx, client)
	case "aws_kinesis_firehose_delivery_stream":
		return ListKinesisFirehoseDeliveryStream(ctx, client)
	case "aws_kinesis_stream":
		return ListKinesisStream(ctx, client)
	case "aws_kinesisanalyticsv2_application":
		return ListKinesisanalyticsv2Application(ctx, client)
	case "aws_kms_external_key":
		return ListKmsExternalKey(ctx, client)
	case "aws_kms_key":
		return ListKmsKey(ctx, client)
	case "aws_lambda_code_signing_config":
		return ListLambdaCodeSigningConfig(ctx, client)
	case "aws_lambda_event_source_mapping":
		return ListLambdaEventSourceMapping(ctx, client)
	case "aws_lambda_function":
		return ListLambdaFunction(ctx, client)
	case "aws_launch_configuration":
		return ListLaunchConfiguration(ctx, client)
	case "aws_launch_template":
		return ListLaunchTemplate(ctx, client)
	case "aws_lb":
		return ListLb(ctx, client)
	case "aws_lb_target_group":
		return ListLbTargetGroup(ctx, client)
	case "aws_lex_bot":
		return ListLexBot(ctx, client)
	case "aws_lex_intent":
		return ListLexIntent(ctx, client)
	case "aws_lex_slot_type":
		return ListLexSlotType(ctx, client)
	case "aws_licensemanager_license_configuration":
		return ListLicensemanagerLicenseConfiguration(ctx, client)
	case "aws_lightsail_domain":
		return ListLightsailDomain(ctx, client)
	case "aws_lightsail_instance":
		return ListLightsailInstance(ctx, client)
	case "aws_lightsail_key_pair":
		return ListLightsailKeyPair(ctx, client)
	case "aws_lightsail_static_ip":
		return ListLightsailStaticIp(ctx, client)
	case "aws_media_convert_queue":
		return ListMediaConvertQueue(ctx, client)
	case "aws_media_package_channel":
		return ListMediaPackageChannel(ctx, client)
	case "aws_media_store_container":
		return ListMediaStoreContainer(ctx, client)
	case "aws_mq_broker":
		return ListMqBroker(ctx, client)
	case "aws_mq_configuration":
		return ListMqConfiguration(ctx, client)
	case "aws_msk_cluster":
		return ListMskCluster(ctx, client)
	case "aws_msk_configuration":
		return ListMskConfiguration(ctx, client)
	case "aws_nat_gateway":
		return ListNatGateway(ctx, client)
	case "aws_neptune_event_subscription":
		return ListNeptuneEventSubscription(ctx, client)
	case "aws_network_acl":
		return ListNetworkAcl(ctx, client)
	case "aws_network_interface":
		return ListNetworkInterface(ctx, client)
	case "aws_networkfirewall_firewall":
		return ListNetworkfirewallFirewall(ctx, client)
	case "aws_networkfirewall_firewall_policy":
		return ListNetworkfirewallFirewallPolicy(ctx, client)
	case "aws_networkfirewall_rule_group":
		return ListNetworkfirewallRuleGroup(ctx, client)
	case "aws_opsworks_stack":
		return ListOpsworksStack(ctx, client)
	case "aws_opsworks_user_profile":
		return ListOpsworksUserProfile(ctx, client)
	case "aws_placement_group":
		return ListPlacementGroup(ctx, client)
	case "aws_qldb_ledger":
		return ListQldbLedger(ctx, client)
	case "aws_rds_cluster":
		return ListRdsCluster(ctx, client)
	case "aws_rds_cluster_endpoint":
		return ListRdsClusterEndpoint(ctx, client)
	case "aws_rds_cluster_parameter_group":
		return ListRdsClusterParameterGroup(ctx, client)
	case "aws_rds_global_cluster":
		return ListRdsGlobalCluster(ctx, client)
	case "aws_redshift_cluster":
		return ListRedshiftCluster(ctx, client)
	case "aws_redshift_event_subscription":
		return ListRedshiftEventSubscription(ctx, client)
	case "aws_redshift_parameter_group":
		return ListRedshiftParameterGroup(ctx, client)
	case "aws_redshift_security_group":
		return ListRedshiftSecurityGroup(ctx, client)
	case "aws_redshift_snapshot_copy_grant":
		return ListRedshiftSnapshotCopyGrant(ctx, client)
	case "aws_redshift_snapshot_schedule":
		return ListRedshiftSnapshotSchedule(ctx, client)
	case "aws_redshift_subnet_group":
		return ListRedshiftSubnetGroup(ctx, client)
	case "aws_route53_health_check":
		return ListRoute53HealthCheck(ctx, client)
	case "aws_route53_resolver_endpoint":
		return ListRoute53ResolverEndpoint(ctx, client)
	case "aws_route53_resolver_query_log_config":
		return ListRoute53ResolverQueryLogConfig(ctx, client)
	case "aws_route53_resolver_query_log_config_association":
		return ListRoute53ResolverQueryLogConfigAssociation(ctx, client)
	case "aws_route53_resolver_rule":
		return ListRoute53ResolverRule(ctx, client)
	case "aws_route53_resolver_rule_association":
		return ListRoute53ResolverRuleAssociation(ctx, client)
	case "aws_route53_zone":
		return ListRoute53Zone(ctx, client)
	case "aws_route_table":
		return ListRouteTable(ctx, client)
	case "aws_s3_bucket":
		return ListS3Bucket(ctx, client)
	case "aws_s3outposts_endpoint":
		return ListS3outpostsEndpoint(ctx, client)
	case "aws_sagemaker_app_image_config":
		return ListSagemakerAppImageConfig(ctx, client)
	case "aws_sagemaker_code_repository":
		return ListSagemakerCodeRepository(ctx, client)
	case "aws_sagemaker_endpoint":
		return ListSagemakerEndpoint(ctx, client)
	case "aws_sagemaker_feature_group":
		return ListSagemakerFeatureGroup(ctx, client)
	case "aws_sagemaker_model":
		return ListSagemakerModel(ctx, client)
	case "aws_sagemaker_model_package_group":
		return ListSagemakerModelPackageGroup(ctx, client)
	case "aws_secretsmanager_secret":
		return ListSecretsmanagerSecret(ctx, client)
	case "aws_security_group":
		return ListSecurityGroup(ctx, client)
	case "aws_securityhub_action_target":
		return ListSecurityhubActionTarget(ctx, client)
	case "aws_service_discovery_service":
		return ListServiceDiscoveryService(ctx, client)
	case "aws_servicecatalog_portfolio":
		return ListServicecatalogPortfolio(ctx, client)
	case "aws_ses_active_receipt_rule_set":
		return ListSesActiveReceiptRuleSet(ctx, client)
	case "aws_ses_configuration_set":
		return ListSesConfigurationSet(ctx, client)
	case "aws_ses_domain_identity":
		return ListSesDomainIdentity(ctx, client)
	case "aws_ses_email_identity":
		return ListSesEmailIdentity(ctx, client)
	case "aws_ses_receipt_filter":
		return ListSesReceiptFilter(ctx, client)
	case "aws_ses_receipt_rule_set":
		return ListSesReceiptRuleSet(ctx, client)
	case "aws_ses_template":
		return ListSesTemplate(ctx, client)
	case "aws_sfn_activity":
		return ListSfnActivity(ctx, client)
	case "aws_sfn_state_machine":
		return ListSfnStateMachine(ctx, client)
	case "aws_sns_platform_application":
		return ListSnsPlatformApplication(ctx, client)
	case "aws_sns_topic":
		return ListSnsTopic(ctx, client)
	case "aws_sns_topic_subscription":
		return ListSnsTopicSubscription(ctx, client)
	case "aws_spot_fleet_request":
		return ListSpotFleetRequest(ctx, client)
	case "aws_spot_instance_request":
		return ListSpotInstanceRequest(ctx, client)
	case "aws_sqs_queue":
		return ListSqsQueue(ctx, client)
	case "aws_ssm_activation":
		return ListSsmActivation(ctx, client)
	case "aws_ssm_association":
		return ListSsmAssociation(ctx, client)
	case "aws_ssm_document":
		return ListSsmDocument(ctx, client)
	case "aws_ssm_maintenance_window":
		return ListSsmMaintenanceWindow(ctx, client)
	case "aws_ssm_parameter":
		return ListSsmParameter(ctx, client)
	case "aws_ssm_patch_baseline":
		return ListSsmPatchBaseline(ctx, client)
	case "aws_ssm_patch_group":
		return ListSsmPatchGroup(ctx, client)
	case "aws_ssm_resource_data_sync":
		return ListSsmResourceDataSync(ctx, client)
	case "aws_storagegateway_gateway":
		return ListStoragegatewayGateway(ctx, client)
	case "aws_storagegateway_tape_pool":
		return ListStoragegatewayTapePool(ctx, client)
	case "aws_subnet":
		return ListSubnet(ctx, client)
	case "aws_synthetics_canary":
		return ListSyntheticsCanary(ctx, client)
	case "aws_transfer_server":
		return ListTransferServer(ctx, client)
	case "aws_vpc":
		return ListVpc(ctx, client)
	case "aws_vpc_endpoint":
		return ListVpcEndpoint(ctx, client)
	case "aws_vpc_endpoint_connection_notification":
		return ListVpcEndpointConnectionNotification(ctx, client)
	case "aws_vpc_endpoint_service":
		return ListVpcEndpointService(ctx, client)
	case "aws_vpc_peering_connection":
		return ListVpcPeeringConnection(ctx, client)
	case "aws_vpn_gateway":
		return ListVpnGateway(ctx, client)
	case "aws_waf_byte_match_set":
		return ListWafByteMatchSet(ctx, client)
	case "aws_waf_geo_match_set":
		return ListWafGeoMatchSet(ctx, client)
	case "aws_waf_ipset":
		return ListWafIpset(ctx, client)
	case "aws_waf_rate_based_rule":
		return ListWafRateBasedRule(ctx, client)
	case "aws_waf_regex_match_set":
		return ListWafRegexMatchSet(ctx, client)
	case "aws_waf_regex_pattern_set":
		return ListWafRegexPatternSet(ctx, client)
	case "aws_waf_rule":
		return ListWafRule(ctx, client)
	case "aws_waf_rule_group":
		return ListWafRuleGroup(ctx, client)
	case "aws_waf_size_constraint_set":
		return ListWafSizeConstraintSet(ctx, client)
	case "aws_waf_sql_injection_match_set":
		return ListWafSqlInjectionMatchSet(ctx, client)
	case "aws_waf_web_acl":
		return ListWafWebAcl(ctx, client)
	case "aws_waf_xss_match_set":
		return ListWafXssMatchSet(ctx, client)
	case "aws_wafv2_web_acl_logging_configuration":
		return ListWafv2WebAclLoggingConfiguration(ctx, client)
	case "aws_worklink_fleet":
		return ListWorklinkFleet(ctx, client)
	case "aws_workspaces_directory":
		return ListWorkspacesDirectory(ctx, client)
	case "aws_workspaces_ip_group":
		return ListWorkspacesIpGroup(ctx, client)
	case "aws_workspaces_workspace":
		return ListWorkspacesWorkspace(ctx, client)
	case "aws_xray_group":
		return ListXrayGroup(ctx, client)
	default:
		return nil, fmt.Errorf("resource type is not (yet) supported: %s", resourceType)
	}
}
