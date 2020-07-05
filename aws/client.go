// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/macie"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/simpledb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

type Client struct {
	Accountid                     string
	Region                        string
	Accessanalyzerconn            *accessanalyzer.Client
	Acmconn                       *acm.Client
	Acmpcaconn                    *acmpca.Client
	Apigatewayconn                *apigateway.Client
	Apigatewayv2conn              *apigatewayv2.Client
	Applicationautoscalingconn    *applicationautoscaling.Client
	Appmeshconn                   *appmesh.Client
	Appsyncconn                   *appsync.Client
	Athenaconn                    *athena.Client
	Autoscalingconn               *autoscaling.Client
	Backupconn                    *backup.Client
	Batchconn                     *batch.Client
	Budgetsconn                   *budgets.Client
	Cloud9conn                    *cloud9.Client
	Cloudformationconn            *cloudformation.Client
	Cloudfrontconn                *cloudfront.Client
	Cloudhsmv2conn                *cloudhsmv2.Client
	Cloudtrailconn                *cloudtrail.Client
	Cloudwatchconn                *cloudwatch.Client
	Cloudwatcheventsconn          *cloudwatchevents.Client
	Cloudwatchlogsconn            *cloudwatchlogs.Client
	Codebuildconn                 *codebuild.Client
	Codecommitconn                *codecommit.Client
	Codedeployconn                *codedeploy.Client
	Codepipelineconn              *codepipeline.Client
	Codestarnotificationsconn     *codestarnotifications.Client
	Cognitoidentityconn           *cognitoidentity.Client
	Cognitoidentityproviderconn   *cognitoidentityprovider.Client
	Configserviceconn             *configservice.Client
	Costandusagereportserviceconn *costandusagereportservice.Client
	Databasemigrationserviceconn  *databasemigrationservice.Client
	Datapipelineconn              *datapipeline.Client
	Datasyncconn                  *datasync.Client
	Daxconn                       *dax.Client
	Devicefarmconn                *devicefarm.Client
	Directconnectconn             *directconnect.Client
	Directoryserviceconn          *directoryservice.Client
	Dlmconn                       *dlm.Client
	Docdbconn                     *docdb.Client
	Dynamodbconn                  *dynamodb.Client
	Ec2conn                       *ec2.Client
	Ecrconn                       *ecr.Client
	Ecsconn                       *ecs.Client
	Efsconn                       *efs.Client
	Eksconn                       *eks.Client
	Elasticacheconn               *elasticache.Client
	Elasticbeanstalkconn          *elasticbeanstalk.Client
	Elasticloadbalancingconn      *elasticloadbalancing.Client
	Elasticloadbalancingv2conn    *elasticloadbalancingv2.Client
	Elasticsearchserviceconn      *elasticsearchservice.Client
	Elastictranscoderconn         *elastictranscoder.Client
	Emrconn                       *emr.Client
	Firehoseconn                  *firehose.Client
	Fmsconn                       *fms.Client
	Fsxconn                       *fsx.Client
	Gameliftconn                  *gamelift.Client
	Glacierconn                   *glacier.Client
	Globalacceleratorconn         *globalaccelerator.Client
	Glueconn                      *glue.Client
	Guarddutyconn                 *guardduty.Client
	Iamconn                       *iam.Client
	Inspectorconn                 *inspector.Client
	Iotconn                       *iot.Client
	Kafkaconn                     *kafka.Client
	Kinesisconn                   *kinesis.Client
	Kinesisanalyticsconn          *kinesisanalytics.Client
	Kinesisvideoconn              *kinesisvideo.Client
	Kmsconn                       *kms.Client
	Lambdaconn                    *lambda.Client
	Licensemanagerconn            *licensemanager.Client
	Lightsailconn                 *lightsail.Client
	Macieconn                     *macie.Client
	Mediaconvertconn              *mediaconvert.Client
	Mediapackageconn              *mediapackage.Client
	Mediastoreconn                *mediastore.Client
	Mqconn                        *mq.Client
	Neptuneconn                   *neptune.Client
	Opsworksconn                  *opsworks.Client
	Organizationsconn             *organizations.Client
	Pinpointconn                  *pinpoint.Client
	Qldbconn                      *qldb.Client
	Quicksightconn                *quicksight.Client
	Ramconn                       *ram.Client
	Rdsconn                       *rds.Client
	Redshiftconn                  *redshift.Client
	Resourcegroupsconn            *resourcegroups.Client
	Route53conn                   *route53.Client
	Route53resolverconn           *route53resolver.Client
	S3conn                        *s3.Client
	S3controlconn                 *s3control.Client
	Sagemakerconn                 *sagemaker.Client
	Secretsmanagerconn            *secretsmanager.Client
	Securityhubconn               *securityhub.Client
	Servicecatalogconn            *servicecatalog.Client
	Servicediscoveryconn          *servicediscovery.Client
	Servicequotasconn             *servicequotas.Client
	Sesconn                       *ses.Client
	Sfnconn                       *sfn.Client
	Shieldconn                    *shield.Client
	Simpledbconn                  *simpledb.Client
	Snsconn                       *sns.Client
	Sqsconn                       *sqs.Client
	Ssmconn                       *ssm.Client
	Storagegatewayconn            *storagegateway.Client
	Swfconn                       *swf.Client
	Transferconn                  *transfer.Client
	Wafconn                       *waf.Client
	Wafregionalconn               *wafregional.Client
	Wafv2conn                     *wafv2.Client
	Worklinkconn                  *worklink.Client
	Workspacesconn                *workspaces.Client
	Xrayconn                      *xray.Client
}

func NewClient(region string) (*Client, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	if region != "" {
		cfg.Region = region
	}

	client := &Client{
		Accessanalyzerconn:            accessanalyzer.New(cfg),
		Acmconn:                       acm.New(cfg),
		Acmpcaconn:                    acmpca.New(cfg),
		Apigatewayconn:                apigateway.New(cfg),
		Apigatewayv2conn:              apigatewayv2.New(cfg),
		Applicationautoscalingconn:    applicationautoscaling.New(cfg),
		Appmeshconn:                   appmesh.New(cfg),
		Appsyncconn:                   appsync.New(cfg),
		Athenaconn:                    athena.New(cfg),
		Autoscalingconn:               autoscaling.New(cfg),
		Backupconn:                    backup.New(cfg),
		Batchconn:                     batch.New(cfg),
		Budgetsconn:                   budgets.New(cfg),
		Cloud9conn:                    cloud9.New(cfg),
		Cloudformationconn:            cloudformation.New(cfg),
		Cloudfrontconn:                cloudfront.New(cfg),
		Cloudhsmv2conn:                cloudhsmv2.New(cfg),
		Cloudtrailconn:                cloudtrail.New(cfg),
		Cloudwatchconn:                cloudwatch.New(cfg),
		Cloudwatcheventsconn:          cloudwatchevents.New(cfg),
		Cloudwatchlogsconn:            cloudwatchlogs.New(cfg),
		Codebuildconn:                 codebuild.New(cfg),
		Codecommitconn:                codecommit.New(cfg),
		Codedeployconn:                codedeploy.New(cfg),
		Codepipelineconn:              codepipeline.New(cfg),
		Codestarnotificationsconn:     codestarnotifications.New(cfg),
		Cognitoidentityconn:           cognitoidentity.New(cfg),
		Cognitoidentityproviderconn:   cognitoidentityprovider.New(cfg),
		Configserviceconn:             configservice.New(cfg),
		Costandusagereportserviceconn: costandusagereportservice.New(cfg),
		Databasemigrationserviceconn:  databasemigrationservice.New(cfg),
		Datapipelineconn:              datapipeline.New(cfg),
		Datasyncconn:                  datasync.New(cfg),
		Daxconn:                       dax.New(cfg),
		Devicefarmconn:                devicefarm.New(cfg),
		Directconnectconn:             directconnect.New(cfg),
		Directoryserviceconn:          directoryservice.New(cfg),
		Dlmconn:                       dlm.New(cfg),
		Docdbconn:                     docdb.New(cfg),
		Dynamodbconn:                  dynamodb.New(cfg),
		Ec2conn:                       ec2.New(cfg),
		Ecrconn:                       ecr.New(cfg),
		Ecsconn:                       ecs.New(cfg),
		Efsconn:                       efs.New(cfg),
		Eksconn:                       eks.New(cfg),
		Elasticacheconn:               elasticache.New(cfg),
		Elasticbeanstalkconn:          elasticbeanstalk.New(cfg),
		Elasticloadbalancingconn:      elasticloadbalancing.New(cfg),
		Elasticloadbalancingv2conn:    elasticloadbalancingv2.New(cfg),
		Elasticsearchserviceconn:      elasticsearchservice.New(cfg),
		Elastictranscoderconn:         elastictranscoder.New(cfg),
		Emrconn:                       emr.New(cfg),
		Firehoseconn:                  firehose.New(cfg),
		Fmsconn:                       fms.New(cfg),
		Fsxconn:                       fsx.New(cfg),
		Gameliftconn:                  gamelift.New(cfg),
		Glacierconn:                   glacier.New(cfg),
		Globalacceleratorconn:         globalaccelerator.New(cfg),
		Glueconn:                      glue.New(cfg),
		Guarddutyconn:                 guardduty.New(cfg),
		Iamconn:                       iam.New(cfg),
		Inspectorconn:                 inspector.New(cfg),
		Iotconn:                       iot.New(cfg),
		Kafkaconn:                     kafka.New(cfg),
		Kinesisconn:                   kinesis.New(cfg),
		Kinesisanalyticsconn:          kinesisanalytics.New(cfg),
		Kinesisvideoconn:              kinesisvideo.New(cfg),
		Kmsconn:                       kms.New(cfg),
		Lambdaconn:                    lambda.New(cfg),
		Licensemanagerconn:            licensemanager.New(cfg),
		Lightsailconn:                 lightsail.New(cfg),
		Macieconn:                     macie.New(cfg),
		Mediaconvertconn:              mediaconvert.New(cfg),
		Mediapackageconn:              mediapackage.New(cfg),
		Mediastoreconn:                mediastore.New(cfg),
		Mqconn:                        mq.New(cfg),
		Neptuneconn:                   neptune.New(cfg),
		Opsworksconn:                  opsworks.New(cfg),
		Organizationsconn:             organizations.New(cfg),
		Pinpointconn:                  pinpoint.New(cfg),
		Qldbconn:                      qldb.New(cfg),
		Quicksightconn:                quicksight.New(cfg),
		Ramconn:                       ram.New(cfg),
		Rdsconn:                       rds.New(cfg),
		Redshiftconn:                  redshift.New(cfg),
		Resourcegroupsconn:            resourcegroups.New(cfg),
		Route53conn:                   route53.New(cfg),
		Route53resolverconn:           route53resolver.New(cfg),
		S3conn:                        s3.New(cfg),
		S3controlconn:                 s3control.New(cfg),
		Sagemakerconn:                 sagemaker.New(cfg),
		Secretsmanagerconn:            secretsmanager.New(cfg),
		Securityhubconn:               securityhub.New(cfg),
		Servicecatalogconn:            servicecatalog.New(cfg),
		Servicediscoveryconn:          servicediscovery.New(cfg),
		Servicequotasconn:             servicequotas.New(cfg),
		Sesconn:                       ses.New(cfg),
		Sfnconn:                       sfn.New(cfg),
		Shieldconn:                    shield.New(cfg),
		Simpledbconn:                  simpledb.New(cfg),
		Snsconn:                       sns.New(cfg),
		Sqsconn:                       sqs.New(cfg),
		Ssmconn:                       ssm.New(cfg),
		Storagegatewayconn:            storagegateway.New(cfg),
		Swfconn:                       swf.New(cfg),
		Transferconn:                  transfer.New(cfg),
		Wafconn:                       waf.New(cfg),
		Wafregionalconn:               wafregional.New(cfg),
		Wafv2conn:                     wafv2.New(cfg),
		Worklinkconn:                  worklink.New(cfg),
		Workspacesconn:                workspaces.New(cfg),
		Xrayconn:                      xray.New(cfg),
	}

	stsconn := sts.New(cfg)
	req := stsconn.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}

	client.Accountid = *resp.Account
	client.Region = cfg.Region

	return client, nil
}
