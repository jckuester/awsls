module github.com/jckuester/awsls

go 1.16

require (
	github.com/apex/log v1.9.0
	github.com/aws/aws-sdk-go v1.38.43
	github.com/aws/aws-sdk-go-v2 v1.6.0
	github.com/aws/aws-sdk-go-v2/internal/ini v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.1.1
	github.com/aws/aws-sdk-go-v2/service/acm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.1.1
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/appmesh v1.1.1
	github.com/aws/aws-sdk-go-v2/service/apprunner v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/appsync v1.1.1
	github.com/aws/aws-sdk-go-v2/service/athena v1.1.1
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.1.1
	github.com/aws/aws-sdk-go-v2/service/backup v1.1.1
	github.com/aws/aws-sdk-go-v2/service/batch v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudhsmv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.1.1
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codeartifact v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codestarconnections v1.1.1
	github.com/aws/aws-sdk-go-v2/service/codestarnotifications v1.1.1
	github.com/aws/aws-sdk-go-v2/service/configservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/costandusagereportservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/datasync v1.1.1
	github.com/aws/aws-sdk-go-v2/service/dax v1.1.1
	github.com/aws/aws-sdk-go-v2/service/devicefarm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/directconnect v1.1.1
	github.com/aws/aws-sdk-go-v2/service/dlm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ecr v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.1.2
	github.com/aws/aws-sdk-go-v2/service/ecs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/efs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/eks v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/elastictranscoder v1.1.1
	github.com/aws/aws-sdk-go-v2/service/emr v1.1.1
	github.com/aws/aws-sdk-go-v2/service/firehose v1.1.1
	github.com/aws/aws-sdk-go-v2/service/fms v1.1.1
	github.com/aws/aws-sdk-go-v2/service/fsx v1.1.1
	github.com/aws/aws-sdk-go-v2/service/gamelift v1.1.1
	github.com/aws/aws-sdk-go-v2/service/globalaccelerator v1.1.1
	github.com/aws/aws-sdk-go-v2/service/glue v1.1.1
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.1.1
	github.com/aws/aws-sdk-go-v2/service/iam v1.1.1
	github.com/aws/aws-sdk-go-v2/service/imagebuilder v1.1.1
	github.com/aws/aws-sdk-go-v2/service/iot v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kafka v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesisanalytics v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/kms v1.1.1
	github.com/aws/aws-sdk-go-v2/service/lambda v1.1.1
	github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice v1.1.1
	github.com/aws/aws-sdk-go-v2/service/licensemanager v1.1.1
	github.com/aws/aws-sdk-go-v2/service/lightsail v1.1.1
	github.com/aws/aws-sdk-go-v2/service/macie2 v1.6.0
	github.com/aws/aws-sdk-go-v2/service/mediaconvert v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mediapackage v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mediastore v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mq v1.1.1
	github.com/aws/aws-sdk-go-v2/service/mwaa v1.1.1
	github.com/aws/aws-sdk-go-v2/service/neptune v1.1.1
	github.com/aws/aws-sdk-go-v2/service/networkfirewall v1.1.2
	github.com/aws/aws-sdk-go-v2/service/opsworks v1.1.1
	github.com/aws/aws-sdk-go-v2/service/qldb v1.1.1
	github.com/aws/aws-sdk-go-v2/service/rds v1.1.1
	github.com/aws/aws-sdk-go-v2/service/redshift v1.1.1
	github.com/aws/aws-sdk-go-v2/service/route53 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/route53resolver v1.1.1
	github.com/aws/aws-sdk-go-v2/service/s3 v1.2.0
	github.com/aws/aws-sdk-go-v2/service/s3outposts v1.1.2
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.1.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.1.1
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.1.1
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v1.1.1
	github.com/aws/aws-sdk-go-v2/service/servicediscovery v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ses v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sfn v1.1.1
	github.com/aws/aws-sdk-go-v2/service/signer v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sns v1.1.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.1.1
	github.com/aws/aws-sdk-go-v2/service/ssm v1.1.1
	github.com/aws/aws-sdk-go-v2/service/storagegateway v1.1.1
	github.com/aws/aws-sdk-go-v2/service/synthetics v1.1.1
	github.com/aws/aws-sdk-go-v2/service/timestreamwrite v1.2.1
	github.com/aws/aws-sdk-go-v2/service/transfer v1.1.1
	github.com/aws/aws-sdk-go-v2/service/waf v1.1.1
	github.com/aws/aws-sdk-go-v2/service/wafregional v1.1.1
	github.com/aws/aws-sdk-go-v2/service/wafv2 v1.1.1
	github.com/aws/aws-sdk-go-v2/service/worklink v1.1.1
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.1.1
	github.com/aws/aws-sdk-go-v2/service/xray v1.1.1
	github.com/disneystreaming/go-ssmhelpers v0.2.1
	github.com/fatih/color v1.10.0
	github.com/gobwas/glob v0.2.3
	github.com/gruntwork-io/terratest v0.23.0
	github.com/jckuester/awstools-lib v0.0.0-20210524191941-23f0e367139d
	github.com/jckuester/terradozer v0.1.4-0.20210524190016-3e6d42479316
	github.com/onsi/gomega v1.9.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.7.0
	github.com/zclconf/go-cty v1.7.1
)
