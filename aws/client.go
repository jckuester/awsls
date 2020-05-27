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
	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

type Client struct {
	accountid                     string
	accessanalyzerconn            *accessanalyzer.Client
	acmconn                       *acm.Client
	acmpcaconn                    *acmpca.Client
	apigatewayconn                *apigateway.Client
	apigatewayv2conn              *apigatewayv2.Client
	applicationautoscalingconn    *applicationautoscaling.Client
	appmeshconn                   *appmesh.Client
	appsyncconn                   *appsync.Client
	athenaconn                    *athena.Client
	autoscalingconn               *autoscaling.Client
	backupconn                    *backup.Client
	batchconn                     *batch.Client
	budgetsconn                   *budgets.Client
	cloud9conn                    *cloud9.Client
	cloudformationconn            *cloudformation.Client
	cloudfrontconn                *cloudfront.Client
	cloudhsmv2conn                *cloudhsmv2.Client
	cloudtrailconn                *cloudtrail.Client
	cloudwatchconn                *cloudwatch.Client
	cloudwatcheventsconn          *cloudwatchevents.Client
	cloudwatchlogsconn            *cloudwatchlogs.Client
	codebuildconn                 *codebuild.Client
	codecommitconn                *codecommit.Client
	codedeployconn                *codedeploy.Client
	codepipelineconn              *codepipeline.Client
	codestarnotificationsconn     *codestarnotifications.Client
	cognitoidentityconn           *cognitoidentity.Client
	cognitoidentityproviderconn   *cognitoidentityprovider.Client
	configserviceconn             *configservice.Client
	costandusagereportserviceconn *costandusagereportservice.Client
	databasemigrationserviceconn  *databasemigrationservice.Client
	datapipelineconn              *datapipeline.Client
	datasyncconn                  *datasync.Client
	daxconn                       *dax.Client
	devicefarmconn                *devicefarm.Client
	directconnectconn             *directconnect.Client
	directoryserviceconn          *directoryservice.Client
	dlmconn                       *dlm.Client
	docdbconn                     *docdb.Client
	dynamodbconn                  *dynamodb.Client
	ec2conn                       *ec2.Client
	ecrconn                       *ecr.Client
	ecsconn                       *ecs.Client
	efsconn                       *efs.Client
	eksconn                       *eks.Client
	elasticacheconn               *elasticache.Client
	elasticbeanstalkconn          *elasticbeanstalk.Client
	elasticloadbalancingconn      *elasticloadbalancing.Client
	elasticloadbalancingv2conn    *elasticloadbalancingv2.Client
	elasticsearchserviceconn      *elasticsearchservice.Client
	elastictranscoderconn         *elastictranscoder.Client
	emrconn                       *emr.Client
	firehoseconn                  *firehose.Client
	fmsconn                       *fms.Client
	fsxconn                       *fsx.Client
	gameliftconn                  *gamelift.Client
	glacierconn                   *glacier.Client
	globalacceleratorconn         *globalaccelerator.Client
	glueconn                      *glue.Client
	guarddutyconn                 *guardduty.Client
	iamconn                       *iam.Client
	inspectorconn                 *inspector.Client
	iotconn                       *iot.Client
	kafkaconn                     *kafka.Client
	kinesisconn                   *kinesis.Client
	kinesisanalyticsconn          *kinesisanalytics.Client
	kinesisvideoconn              *kinesisvideo.Client
	kmsconn                       *kms.Client
	lambdaconn                    *lambda.Client
	licensemanagerconn            *licensemanager.Client
	lightsailconn                 *lightsail.Client
	macieconn                     *macie.Client
	mediaconvertconn              *mediaconvert.Client
	mediapackageconn              *mediapackage.Client
	mediastoreconn                *mediastore.Client
	mqconn                        *mq.Client
	neptuneconn                   *neptune.Client
	opsworksconn                  *opsworks.Client
	organizationsconn             *organizations.Client
	pinpointconn                  *pinpoint.Client
	qldbconn                      *qldb.Client
	quicksightconn                *quicksight.Client
	ramconn                       *ram.Client
	rdsconn                       *rds.Client
	redshiftconn                  *redshift.Client
	resourcegroupsconn            *resourcegroups.Client
	route53conn                   *route53.Client
	route53resolverconn           *route53resolver.Client
	s3conn                        *s3.Client
	s3controlconn                 *s3control.Client
	sagemakerconn                 *sagemaker.Client
	secretsmanagerconn            *secretsmanager.Client
	securityhubconn               *securityhub.Client
	servicecatalogconn            *servicecatalog.Client
	servicediscoveryconn          *servicediscovery.Client
	servicequotasconn             *servicequotas.Client
	sesconn                       *ses.Client
	sfnconn                       *sfn.Client
	shieldconn                    *shield.Client
	simpledbconn                  *simpledb.Client
	snsconn                       *sns.Client
	sqsconn                       *sqs.Client
	ssmconn                       *ssm.Client
	storagegatewayconn            *storagegateway.Client
	swfconn                       *swf.Client
	transferconn                  *transfer.Client
	wafconn                       *waf.Client
	wafregionalconn               *wafregional.Client
	worklinkconn                  *worklink.Client
	workspacesconn                *workspaces.Client
	xrayconn                      *xray.Client
}

func NewClient() (*Client, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
		accessanalyzerconn:            accessanalyzer.New(cfg),
		acmconn:                       acm.New(cfg),
		acmpcaconn:                    acmpca.New(cfg),
		apigatewayconn:                apigateway.New(cfg),
		apigatewayv2conn:              apigatewayv2.New(cfg),
		applicationautoscalingconn:    applicationautoscaling.New(cfg),
		appmeshconn:                   appmesh.New(cfg),
		appsyncconn:                   appsync.New(cfg),
		athenaconn:                    athena.New(cfg),
		autoscalingconn:               autoscaling.New(cfg),
		backupconn:                    backup.New(cfg),
		batchconn:                     batch.New(cfg),
		budgetsconn:                   budgets.New(cfg),
		cloud9conn:                    cloud9.New(cfg),
		cloudformationconn:            cloudformation.New(cfg),
		cloudfrontconn:                cloudfront.New(cfg),
		cloudhsmv2conn:                cloudhsmv2.New(cfg),
		cloudtrailconn:                cloudtrail.New(cfg),
		cloudwatchconn:                cloudwatch.New(cfg),
		cloudwatcheventsconn:          cloudwatchevents.New(cfg),
		cloudwatchlogsconn:            cloudwatchlogs.New(cfg),
		codebuildconn:                 codebuild.New(cfg),
		codecommitconn:                codecommit.New(cfg),
		codedeployconn:                codedeploy.New(cfg),
		codepipelineconn:              codepipeline.New(cfg),
		codestarnotificationsconn:     codestarnotifications.New(cfg),
		cognitoidentityconn:           cognitoidentity.New(cfg),
		cognitoidentityproviderconn:   cognitoidentityprovider.New(cfg),
		configserviceconn:             configservice.New(cfg),
		costandusagereportserviceconn: costandusagereportservice.New(cfg),
		databasemigrationserviceconn:  databasemigrationservice.New(cfg),
		datapipelineconn:              datapipeline.New(cfg),
		datasyncconn:                  datasync.New(cfg),
		daxconn:                       dax.New(cfg),
		devicefarmconn:                devicefarm.New(cfg),
		directconnectconn:             directconnect.New(cfg),
		directoryserviceconn:          directoryservice.New(cfg),
		dlmconn:                       dlm.New(cfg),
		docdbconn:                     docdb.New(cfg),
		dynamodbconn:                  dynamodb.New(cfg),
		ec2conn:                       ec2.New(cfg),
		ecrconn:                       ecr.New(cfg),
		ecsconn:                       ecs.New(cfg),
		efsconn:                       efs.New(cfg),
		eksconn:                       eks.New(cfg),
		elasticacheconn:               elasticache.New(cfg),
		elasticbeanstalkconn:          elasticbeanstalk.New(cfg),
		elasticloadbalancingconn:      elasticloadbalancing.New(cfg),
		elasticloadbalancingv2conn:    elasticloadbalancingv2.New(cfg),
		elasticsearchserviceconn:      elasticsearchservice.New(cfg),
		elastictranscoderconn:         elastictranscoder.New(cfg),
		emrconn:                       emr.New(cfg),
		firehoseconn:                  firehose.New(cfg),
		fmsconn:                       fms.New(cfg),
		fsxconn:                       fsx.New(cfg),
		gameliftconn:                  gamelift.New(cfg),
		glacierconn:                   glacier.New(cfg),
		globalacceleratorconn:         globalaccelerator.New(cfg),
		glueconn:                      glue.New(cfg),
		guarddutyconn:                 guardduty.New(cfg),
		iamconn:                       iam.New(cfg),
		inspectorconn:                 inspector.New(cfg),
		iotconn:                       iot.New(cfg),
		kafkaconn:                     kafka.New(cfg),
		kinesisconn:                   kinesis.New(cfg),
		kinesisanalyticsconn:          kinesisanalytics.New(cfg),
		kinesisvideoconn:              kinesisvideo.New(cfg),
		kmsconn:                       kms.New(cfg),
		lambdaconn:                    lambda.New(cfg),
		licensemanagerconn:            licensemanager.New(cfg),
		lightsailconn:                 lightsail.New(cfg),
		macieconn:                     macie.New(cfg),
		mediaconvertconn:              mediaconvert.New(cfg),
		mediapackageconn:              mediapackage.New(cfg),
		mediastoreconn:                mediastore.New(cfg),
		mqconn:                        mq.New(cfg),
		neptuneconn:                   neptune.New(cfg),
		opsworksconn:                  opsworks.New(cfg),
		organizationsconn:             organizations.New(cfg),
		pinpointconn:                  pinpoint.New(cfg),
		qldbconn:                      qldb.New(cfg),
		quicksightconn:                quicksight.New(cfg),
		ramconn:                       ram.New(cfg),
		rdsconn:                       rds.New(cfg),
		redshiftconn:                  redshift.New(cfg),
		resourcegroupsconn:            resourcegroups.New(cfg),
		route53conn:                   route53.New(cfg),
		route53resolverconn:           route53resolver.New(cfg),
		s3conn:                        s3.New(cfg),
		s3controlconn:                 s3control.New(cfg),
		sagemakerconn:                 sagemaker.New(cfg),
		secretsmanagerconn:            secretsmanager.New(cfg),
		securityhubconn:               securityhub.New(cfg),
		servicecatalogconn:            servicecatalog.New(cfg),
		servicediscoveryconn:          servicediscovery.New(cfg),
		servicequotasconn:             servicequotas.New(cfg),
		sesconn:                       ses.New(cfg),
		sfnconn:                       sfn.New(cfg),
		shieldconn:                    shield.New(cfg),
		simpledbconn:                  simpledb.New(cfg),
		snsconn:                       sns.New(cfg),
		sqsconn:                       sqs.New(cfg),
		ssmconn:                       ssm.New(cfg),
		storagegatewayconn:            storagegateway.New(cfg),
		swfconn:                       swf.New(cfg),
		transferconn:                  transfer.New(cfg),
		wafconn:                       waf.New(cfg),
		wafregionalconn:               wafregional.New(cfg),
		worklinkconn:                  worklink.New(cfg),
		workspacesconn:                workspaces.New(cfg),
		xrayconn:                      xray.New(cfg),
	}

	stsconn := sts.New(cfg)
	req := stsconn.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	resp, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	client.accountid = *resp.Account

	return client, nil
}
