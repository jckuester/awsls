// +build codegen

package terraform

// ManualResourceServiceMap manually matches some resources to its AWS service since the
// service couldn't be automatically discovered.
var ManualResourceServiceMap = map[string]string{
	"aws_alb":                                 "elbv2",
	"aws_autoscaling_group":                   "autoscaling",
	"aws_efs_mount_target":                    "efs",
	"aws_elastic_beanstalk_environment":       "elasticbeanstalk",
	"aws_elb":                                 "elb",
	"aws_iam_server_certificate":              "iam",
	"aws_lambda_event_source_mapping":         "lambda",
	"aws_launch_configuration":                "autoscaling",
	"aws_lb":                                  "elbv2",
	"aws_s3_bucket":                           "s3",
	"aws_s3_bucket_object":                    "s3",
	"aws_s3_bucket_public_access_block":       "s3",
	"aws_wafregional_byte_match_set":          "wafregional",
	"aws_wafregional_geo_match_set":           "wafregional",
	"aws_wafregional_ipset":                   "wafregional",
	"aws_wafregional_rate_based_rule":         "wafregional",
	"aws_wafregional_regex_match_set":         "wafregional",
	"aws_wafregional_regex_pattern_set":       "wafregional",
	"aws_wafregional_rule":                    "wafregional",
	"aws_wafregional_rule_group":              "wafregional",
	"aws_wafregional_size_constraint_set":     "wafregional",
	"aws_wafregional_sql_injection_match_set": "wafregional",
	"aws_wafregional_web_acl":                 "wafregional",
	"aws_wafregional_xss_match_set":           "wafregional",
}
