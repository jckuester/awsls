// +build codegen

package aws

// some resource types are excluded as they need be handled slightly differently
var ExcludedResourceTypes = map[string]bool{
	// is not a resource
	"aws_acm_certificate_validation": true,
	// ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_alb_listener": true,
	//ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_alb_listener_rule": true,
	// A normal subnet with "DefaultForAz" attribute set to true
	"aws_default_subnet": true,
	// BadRequest: Must specify exactly one mutually exclusive parameter.
	"aws_acmpca_certificate_authority":  true,
	"aws_ami_from_instance":             true,
	"aws_appmesh_route":                 true,
	"aws_appmesh_virtual_node":          true,
	"aws_appmesh_virtual_router":        true,
	"aws_appmesh_virtual_service":       true,
	"aws_datapipeline_pipeline":         true,
	"aws_efs_mount_target":              true,
	"aws_emr_cluster":                   true,
	"aws_glue_classifier":               true,
	"aws_iam_openid_connect_provider":   true,
	"aws_iam_saml_provider":             true,
	"aws_inspector_assessment_target":   true,
	"aws_inspector_assessment_template": true,
	"aws_inspector_resource_group":      true,
	"aws_instance":                      true,
	"aws_resourcegroups_group":          true,
	"aws_shield_protection":             true,
	"aws_vpn_connection":                true,
	"aws_xray_sampling_rule":            true,
	// ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_lb_listener": true,
	// ValidationError: You must specify either either listener ARNs or a load balancer ARN
	"aws_lb_listener_rule": true,
	// ValidationException: Please provide either one or more instance ID or one stack ID or one layer ID
	"aws_opsworks_instance":      true,
	"aws_default_route_table":    true,
	"aws_default_security_group": true,
	// A normal VPC with "IsDefault" attribute set to true
	"aws_default_vpc": true,
	// Error: failed to read current state of resource: rpc error: code = Canceled desc = context canceled
	"aws_kms_alias": true,
}

// manualMatchedListOps are list operations that could not be matched automatically
var ManualMatchedListOps = map[string]string{
	// The AWS API calls it images not AMI
	"aws_ami":                  "DescribeImages",
	"aws_ecs_service":          "DescribeServices",
	"aws_ecs_cluster":          "DescribeClusters",
	"aws_eip":                  "DescribeAddresses",
	"aws_elb":                  "DescribeLoadBalancers",
	"aws_route53_zone":         "ListHostedZones",
	"aws_cloudformation_stack": "DescribeStacks",
}

var ManualMatchedOutputFields = map[string]string{
	"aws_ecs_service": "ServicePkgNames",
	"aws_ecs_cluster": "Clusters",
}

var ManualMatchedResourceID = map[string]string{
	"aws_autoscaling_group":    "AutoScalingGroupName",
	"aws_launch_configuration": "LaunchConfigurationName",
	"aws_s3_bucket":            "Name",
	"aws_elb":                  "LoadBalancerName",
	"aws_db_instance":          "DBInstanceIdentifier",
	"aws_route53_zone":         "Id",
}
