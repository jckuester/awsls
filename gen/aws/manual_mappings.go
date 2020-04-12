// +build codegen

package aws

// some resource types are excluded as they need be handled slightly differently
var ExcludedResourceTypes = map[string]bool{
	// A normal subnet with "DefaultForAz" attribute set to true
	"aws_default_subnet":                true,
	"aws_ami_from_instance":             true,
	"aws_inspector_assessment_target":   true,
	"aws_inspector_assessment_template": true,
	"aws_instance":                      true,
	"aws_shield_protection":             true,
	"aws_iam_saml_provider":             true,
	"aws_iam_openid_connect_provider":   true,
	"aws_emr_cluster":                   true,
	"aws_acmpca_certificate_authority":  true,
	"aws_appmesh_route":                 true,
	"aws_appmesh_virtual_node":          true,
	"aws_appmesh_virtual_router":        true,
	"aws_appmesh_virtual_service":       true,
	"aws_inspector_resource_group":      true,
	"aws_vpn_connection":                true,
	"aws_xray_sampling_rule":            true,
	"aws_resourcegroups_group":          true,
	"aws_datapipeline_pipeline":         true,
	"aws_glue_classifier":               true,
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
