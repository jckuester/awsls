// Code is generated. DO NOT EDIT.

package aws

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acmpca"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go-v2/service/amplify"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/aws/aws-sdk-go-v2/service/appconfig"
	"github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights"
	"github.com/aws/aws-sdk-go-v2/service/appmesh"
	"github.com/aws/aws-sdk-go-v2/service/appstream"
	"github.com/aws/aws-sdk-go-v2/service/appsync"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/batch"
	"github.com/aws/aws-sdk-go-v2/service/budgets"
	"github.com/aws/aws-sdk-go-v2/service/chime"
	"github.com/aws/aws-sdk-go-v2/service/cloud9"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsm"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearch"
	"github.com/aws/aws-sdk-go-v2/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact"
	"github.com/aws/aws-sdk-go-v2/service/codebuild"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler"
	"github.com/aws/aws-sdk-go-v2/service/codegurureviewer"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline"
	"github.com/aws/aws-sdk-go-v2/service/codestar"
	"github.com/aws/aws-sdk-go-v2/service/codestarconnections"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/connect"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline"
	"github.com/aws/aws-sdk-go-v2/service/datasync"
	"github.com/aws/aws-sdk-go-v2/service/dax"
	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm"
	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice"
	"github.com/aws/aws-sdk-go-v2/service/dlm"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go-v2/service/ebs"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go-v2/service/elasticinference"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/fms"
	"github.com/aws/aws-sdk-go-v2/service/forecast"
	"github.com/aws/aws-sdk-go-v2/service/forecastquery"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
	"github.com/aws/aws-sdk-go-v2/service/gamelift"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/greengrass"
	"github.com/aws/aws-sdk-go-v2/service/groundstation"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	"github.com/aws/aws-sdk-go-v2/service/health"
	"github.com/aws/aws-sdk-go-v2/service/honeycode"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/aws/aws-sdk-go-v2/service/iotevents"
	"github.com/aws/aws-sdk-go-v2/service/ioteventsdata"
	"github.com/aws/aws-sdk-go-v2/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go-v2/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go-v2/service/ivs"
	"github.com/aws/aws-sdk-go-v2/service/kafka"
	"github.com/aws/aws-sdk-go-v2/service/kendra"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalytics"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideoarchivedmedia"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideomedia"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideosignaling"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning"
	"github.com/aws/aws-sdk-go-v2/service/macie"
	"github.com/aws/aws-sdk-go-v2/service/macie2"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go-v2/service/marketplacecommerceanalytics"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert"
	"github.com/aws/aws-sdk-go-v2/service/medialive"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage"
	"github.com/aws/aws-sdk-go-v2/service/mediapackagevod"
	"github.com/aws/aws-sdk-go-v2/service/mediastore"
	"github.com/aws/aws-sdk-go-v2/service/mediastoredata"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub"
	"github.com/aws/aws-sdk-go-v2/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go-v2/service/mobile"
	"github.com/aws/aws-sdk-go-v2/service/mobileanalytics"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mturk"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager"
	"github.com/aws/aws-sdk-go-v2/service/opsworks"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/outposts"
	"github.com/aws/aws-sdk-go-v2/service/personalize"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents"
	"github.com/aws/aws-sdk-go-v2/service/personalizeruntime"
	"github.com/aws/aws-sdk-go-v2/service/pi"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go-v2/service/polly"
	"github.com/aws/aws-sdk-go-v2/service/pricing"
	"github.com/aws/aws-sdk-go-v2/service/qldb"
	"github.com/aws/aws-sdk-go-v2/service/qldbsession"
	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	"github.com/aws/aws-sdk-go-v2/service/robomaker"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53domains"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker"
	"github.com/aws/aws-sdk-go-v2/service/sagemakera2iruntime"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
	"github.com/aws/aws-sdk-go-v2/service/savingsplans"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/securityhub"
	"github.com/aws/aws-sdk-go-v2/service/serverlessapplicationrepository"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/aws/aws-sdk-go-v2/service/signer"
	"github.com/aws/aws-sdk-go-v2/service/simpledb"
	"github.com/aws/aws-sdk-go-v2/service/sms"
	"github.com/aws/aws-sdk-go-v2/service/snowball"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/sso"
	"github.com/aws/aws-sdk-go-v2/service/ssooidc"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/support"
	"github.com/aws/aws-sdk-go-v2/service/swf"
	"github.com/aws/aws-sdk-go-v2/service/synthetics"
	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/transcribe"
	"github.com/aws/aws-sdk-go-v2/service/transfer"
	"github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/workdocs"
	"github.com/aws/aws-sdk-go-v2/service/worklink"
	"github.com/aws/aws-sdk-go-v2/service/workmail"
	"github.com/aws/aws-sdk-go-v2/service/workmailmessageflow"
	"github.com/aws/aws-sdk-go-v2/service/workspaces"
	"github.com/aws/aws-sdk-go-v2/service/xray"
)

type Client struct {
	AccountID                           string
	Region                              string
	Profile                             string
	Accessanalyzerconn                  *accessanalyzer.Client
	Acmconn                             *acm.Client
	Acmpcaconn                          *acmpca.Client
	Alexaforbusinessconn                *alexaforbusiness.Client
	Amplifyconn                         *amplify.Client
	Apigatewayconn                      *apigateway.Client
	Apigatewaymanagementapiconn         *apigatewaymanagementapi.Client
	Apigatewayv2conn                    *apigatewayv2.Client
	Appconfigconn                       *appconfig.Client
	Applicationautoscalingconn          *applicationautoscaling.Client
	Applicationdiscoveryserviceconn     *applicationdiscoveryservice.Client
	Applicationinsightsconn             *applicationinsights.Client
	Appmeshconn                         *appmesh.Client
	Appstreamconn                       *appstream.Client
	Appsyncconn                         *appsync.Client
	Athenaconn                          *athena.Client
	Autoscalingconn                     *autoscaling.Client
	Autoscalingplansconn                *autoscalingplans.Client
	Backupconn                          *backup.Client
	Batchconn                           *batch.Client
	Budgetsconn                         *budgets.Client
	Chimeconn                           *chime.Client
	Cloud9conn                          *cloud9.Client
	Clouddirectoryconn                  *clouddirectory.Client
	Cloudformationconn                  *cloudformation.Client
	Cloudfrontconn                      *cloudfront.Client
	Cloudhsmconn                        *cloudhsm.Client
	Cloudhsmv2conn                      *cloudhsmv2.Client
	Cloudsearchconn                     *cloudsearch.Client
	Cloudsearchdomainconn               *cloudsearchdomain.Client
	Cloudtrailconn                      *cloudtrail.Client
	Cloudwatchconn                      *cloudwatch.Client
	Cloudwatcheventsconn                *cloudwatchevents.Client
	Cloudwatchlogsconn                  *cloudwatchlogs.Client
	Codeartifactconn                    *codeartifact.Client
	Codebuildconn                       *codebuild.Client
	Codecommitconn                      *codecommit.Client
	Codedeployconn                      *codedeploy.Client
	Codeguruprofilerconn                *codeguruprofiler.Client
	Codegurureviewerconn                *codegurureviewer.Client
	Codepipelineconn                    *codepipeline.Client
	Codestarconn                        *codestar.Client
	Codestarconnectionsconn             *codestarconnections.Client
	Codestarnotificationsconn           *codestarnotifications.Client
	Cognitoidentityconn                 *cognitoidentity.Client
	Cognitoidentityproviderconn         *cognitoidentityprovider.Client
	Cognitosyncconn                     *cognitosync.Client
	Comprehendconn                      *comprehend.Client
	Comprehendmedicalconn               *comprehendmedical.Client
	Computeoptimizerconn                *computeoptimizer.Client
	Configserviceconn                   *configservice.Client
	Connectconn                         *connect.Client
	Connectparticipantconn              *connectparticipant.Client
	Costandusagereportserviceconn       *costandusagereportservice.Client
	Costexplorerconn                    *costexplorer.Client
	Databasemigrationserviceconn        *databasemigrationservice.Client
	Dataexchangeconn                    *dataexchange.Client
	Datapipelineconn                    *datapipeline.Client
	Datasyncconn                        *datasync.Client
	Daxconn                             *dax.Client
	Detectiveconn                       *detective.Client
	Devicefarmconn                      *devicefarm.Client
	Directconnectconn                   *directconnect.Client
	Directoryserviceconn                *directoryservice.Client
	Dlmconn                             *dlm.Client
	Docdbconn                           *docdb.Client
	Dynamodbconn                        *dynamodb.Client
	Dynamodbstreamsconn                 *dynamodbstreams.Client
	Ebsconn                             *ebs.Client
	Ec2conn                             *ec2.Client
	Ec2instanceconnectconn              *ec2instanceconnect.Client
	Ecrconn                             *ecr.Client
	Ecsconn                             *ecs.Client
	Efsconn                             *efs.Client
	Eksconn                             *eks.Client
	Elasticacheconn                     *elasticache.Client
	Elasticbeanstalkconn                *elasticbeanstalk.Client
	Elasticinferenceconn                *elasticinference.Client
	Elasticloadbalancingconn            *elasticloadbalancing.Client
	Elasticloadbalancingv2conn          *elasticloadbalancingv2.Client
	Elasticsearchserviceconn            *elasticsearchservice.Client
	Elastictranscoderconn               *elastictranscoder.Client
	Emrconn                             *emr.Client
	Eventbridgeconn                     *eventbridge.Client
	Firehoseconn                        *firehose.Client
	Fmsconn                             *fms.Client
	Forecastconn                        *forecast.Client
	Forecastqueryconn                   *forecastquery.Client
	Frauddetectorconn                   *frauddetector.Client
	Fsxconn                             *fsx.Client
	Gameliftconn                        *gamelift.Client
	Glacierconn                         *glacier.Client
	Globalacceleratorconn               *globalaccelerator.Client
	Glueconn                            *glue.Client
	Greengrassconn                      *greengrass.Client
	Groundstationconn                   *groundstation.Client
	Guarddutyconn                       *guardduty.Client
	Healthconn                          *health.Client
	Honeycodeconn                       *honeycode.Client
	Iamconn                             *iam.Client
	Imagebuilderconn                    *imagebuilder.Client
	Inspectorconn                       *inspector.Client
	Iotconn                             *iot.Client
	Iot1clickdevicesserviceconn         *iot1clickdevicesservice.Client
	Iot1clickprojectsconn               *iot1clickprojects.Client
	Iotanalyticsconn                    *iotanalytics.Client
	Iotdataplaneconn                    *iotdataplane.Client
	Ioteventsconn                       *iotevents.Client
	Ioteventsdataconn                   *ioteventsdata.Client
	Iotjobsdataplaneconn                *iotjobsdataplane.Client
	Iotsecuretunnelingconn              *iotsecuretunneling.Client
	Iotsitewiseconn                     *iotsitewise.Client
	Iotthingsgraphconn                  *iotthingsgraph.Client
	Ivsconn                             *ivs.Client
	Kafkaconn                           *kafka.Client
	Kendraconn                          *kendra.Client
	Kinesisconn                         *kinesis.Client
	Kinesisanalyticsconn                *kinesisanalytics.Client
	Kinesisanalyticsv2conn              *kinesisanalyticsv2.Client
	Kinesisvideoconn                    *kinesisvideo.Client
	Kinesisvideoarchivedmediaconn       *kinesisvideoarchivedmedia.Client
	Kinesisvideomediaconn               *kinesisvideomedia.Client
	Kinesisvideosignalingconn           *kinesisvideosignaling.Client
	Kmsconn                             *kms.Client
	Lakeformationconn                   *lakeformation.Client
	Lambdaconn                          *lambda.Client
	Lexmodelbuildingserviceconn         *lexmodelbuildingservice.Client
	Lexruntimeserviceconn               *lexruntimeservice.Client
	Licensemanagerconn                  *licensemanager.Client
	Lightsailconn                       *lightsail.Client
	Machinelearningconn                 *machinelearning.Client
	Macieconn                           *macie.Client
	Macie2conn                          *macie2.Client
	Managedblockchainconn               *managedblockchain.Client
	Marketplacecatalogconn              *marketplacecatalog.Client
	Marketplacecommerceanalyticsconn    *marketplacecommerceanalytics.Client
	Marketplaceentitlementserviceconn   *marketplaceentitlementservice.Client
	Marketplacemeteringconn             *marketplacemetering.Client
	Mediaconnectconn                    *mediaconnect.Client
	Mediaconvertconn                    *mediaconvert.Client
	Medialiveconn                       *medialive.Client
	Mediapackageconn                    *mediapackage.Client
	Mediapackagevodconn                 *mediapackagevod.Client
	Mediastoreconn                      *mediastore.Client
	Mediastoredataconn                  *mediastoredata.Client
	Mediatailorconn                     *mediatailor.Client
	Migrationhubconn                    *migrationhub.Client
	Migrationhubconfigconn              *migrationhubconfig.Client
	Mobileconn                          *mobile.Client
	Mobileanalyticsconn                 *mobileanalytics.Client
	Mqconn                              *mq.Client
	Mturkconn                           *mturk.Client
	Neptuneconn                         *neptune.Client
	Networkmanagerconn                  *networkmanager.Client
	Opsworksconn                        *opsworks.Client
	Opsworkscmconn                      *opsworkscm.Client
	Organizationsconn                   *organizations.Client
	Outpostsconn                        *outposts.Client
	Personalizeconn                     *personalize.Client
	Personalizeeventsconn               *personalizeevents.Client
	Personalizeruntimeconn              *personalizeruntime.Client
	Piconn                              *pi.Client
	Pinpointconn                        *pinpoint.Client
	Pinpointemailconn                   *pinpointemail.Client
	Pinpointsmsvoiceconn                *pinpointsmsvoice.Client
	Pollyconn                           *polly.Client
	Pricingconn                         *pricing.Client
	Qldbconn                            *qldb.Client
	Qldbsessionconn                     *qldbsession.Client
	Quicksightconn                      *quicksight.Client
	Ramconn                             *ram.Client
	Rdsconn                             *rds.Client
	Rdsdataconn                         *rdsdata.Client
	Redshiftconn                        *redshift.Client
	Rekognitionconn                     *rekognition.Client
	Resourcegroupsconn                  *resourcegroups.Client
	Resourcegroupstaggingapiconn        *resourcegroupstaggingapi.Client
	Robomakerconn                       *robomaker.Client
	Route53conn                         *route53.Client
	Route53domainsconn                  *route53domains.Client
	Route53resolverconn                 *route53resolver.Client
	S3conn                              *s3.Client
	S3controlconn                       *s3control.Client
	Sagemakerconn                       *sagemaker.Client
	Sagemakera2iruntimeconn             *sagemakera2iruntime.Client
	Sagemakerruntimeconn                *sagemakerruntime.Client
	Savingsplansconn                    *savingsplans.Client
	Schemasconn                         *schemas.Client
	Secretsmanagerconn                  *secretsmanager.Client
	Securityhubconn                     *securityhub.Client
	Serverlessapplicationrepositoryconn *serverlessapplicationrepository.Client
	Servicecatalogconn                  *servicecatalog.Client
	Servicediscoveryconn                *servicediscovery.Client
	Servicequotasconn                   *servicequotas.Client
	Sesconn                             *ses.Client
	Sesv2conn                           *sesv2.Client
	Sfnconn                             *sfn.Client
	Shieldconn                          *shield.Client
	Signerconn                          *signer.Client
	Simpledbconn                        *simpledb.Client
	Smsconn                             *sms.Client
	Snowballconn                        *snowball.Client
	Snsconn                             *sns.Client
	Sqsconn                             *sqs.Client
	Ssmconn                             *ssm.Client
	Ssoconn                             *sso.Client
	Ssooidcconn                         *ssooidc.Client
	Storagegatewayconn                  *storagegateway.Client
	Stsconn                             *sts.Client
	Supportconn                         *support.Client
	Swfconn                             *swf.Client
	Syntheticsconn                      *synthetics.Client
	Textractconn                        *textract.Client
	Transcribeconn                      *transcribe.Client
	Transferconn                        *transfer.Client
	Translateconn                       *translate.Client
	Wafconn                             *waf.Client
	Wafregionalconn                     *wafregional.Client
	Wafv2conn                           *wafv2.Client
	Workdocsconn                        *workdocs.Client
	Worklinkconn                        *worklink.Client
	Workmailconn                        *workmail.Client
	Workmailmessageflowconn             *workmailmessageflow.Client
	Workspacesconn                      *workspaces.Client
	Xrayconn                            *xray.Client
}

func NewClient(configs ...external.Config) (*Client, error) {
	cfg, err := external.LoadDefaultAWSConfig(configs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
		Accessanalyzerconn:                  accessanalyzer.New(cfg),
		Acmconn:                             acm.New(cfg),
		Acmpcaconn:                          acmpca.New(cfg),
		Alexaforbusinessconn:                alexaforbusiness.New(cfg),
		Amplifyconn:                         amplify.New(cfg),
		Apigatewayconn:                      apigateway.New(cfg),
		Apigatewaymanagementapiconn:         apigatewaymanagementapi.New(cfg),
		Apigatewayv2conn:                    apigatewayv2.New(cfg),
		Appconfigconn:                       appconfig.New(cfg),
		Applicationautoscalingconn:          applicationautoscaling.New(cfg),
		Applicationdiscoveryserviceconn:     applicationdiscoveryservice.New(cfg),
		Applicationinsightsconn:             applicationinsights.New(cfg),
		Appmeshconn:                         appmesh.New(cfg),
		Appstreamconn:                       appstream.New(cfg),
		Appsyncconn:                         appsync.New(cfg),
		Athenaconn:                          athena.New(cfg),
		Autoscalingconn:                     autoscaling.New(cfg),
		Autoscalingplansconn:                autoscalingplans.New(cfg),
		Backupconn:                          backup.New(cfg),
		Batchconn:                           batch.New(cfg),
		Budgetsconn:                         budgets.New(cfg),
		Chimeconn:                           chime.New(cfg),
		Cloud9conn:                          cloud9.New(cfg),
		Clouddirectoryconn:                  clouddirectory.New(cfg),
		Cloudformationconn:                  cloudformation.New(cfg),
		Cloudfrontconn:                      cloudfront.New(cfg),
		Cloudhsmconn:                        cloudhsm.New(cfg),
		Cloudhsmv2conn:                      cloudhsmv2.New(cfg),
		Cloudsearchconn:                     cloudsearch.New(cfg),
		Cloudsearchdomainconn:               cloudsearchdomain.New(cfg),
		Cloudtrailconn:                      cloudtrail.New(cfg),
		Cloudwatchconn:                      cloudwatch.New(cfg),
		Cloudwatcheventsconn:                cloudwatchevents.New(cfg),
		Cloudwatchlogsconn:                  cloudwatchlogs.New(cfg),
		Codeartifactconn:                    codeartifact.New(cfg),
		Codebuildconn:                       codebuild.New(cfg),
		Codecommitconn:                      codecommit.New(cfg),
		Codedeployconn:                      codedeploy.New(cfg),
		Codeguruprofilerconn:                codeguruprofiler.New(cfg),
		Codegurureviewerconn:                codegurureviewer.New(cfg),
		Codepipelineconn:                    codepipeline.New(cfg),
		Codestarconn:                        codestar.New(cfg),
		Codestarconnectionsconn:             codestarconnections.New(cfg),
		Codestarnotificationsconn:           codestarnotifications.New(cfg),
		Cognitoidentityconn:                 cognitoidentity.New(cfg),
		Cognitoidentityproviderconn:         cognitoidentityprovider.New(cfg),
		Cognitosyncconn:                     cognitosync.New(cfg),
		Comprehendconn:                      comprehend.New(cfg),
		Comprehendmedicalconn:               comprehendmedical.New(cfg),
		Computeoptimizerconn:                computeoptimizer.New(cfg),
		Configserviceconn:                   configservice.New(cfg),
		Connectconn:                         connect.New(cfg),
		Connectparticipantconn:              connectparticipant.New(cfg),
		Costandusagereportserviceconn:       costandusagereportservice.New(cfg),
		Costexplorerconn:                    costexplorer.New(cfg),
		Databasemigrationserviceconn:        databasemigrationservice.New(cfg),
		Dataexchangeconn:                    dataexchange.New(cfg),
		Datapipelineconn:                    datapipeline.New(cfg),
		Datasyncconn:                        datasync.New(cfg),
		Daxconn:                             dax.New(cfg),
		Detectiveconn:                       detective.New(cfg),
		Devicefarmconn:                      devicefarm.New(cfg),
		Directconnectconn:                   directconnect.New(cfg),
		Directoryserviceconn:                directoryservice.New(cfg),
		Dlmconn:                             dlm.New(cfg),
		Docdbconn:                           docdb.New(cfg),
		Dynamodbconn:                        dynamodb.New(cfg),
		Dynamodbstreamsconn:                 dynamodbstreams.New(cfg),
		Ebsconn:                             ebs.New(cfg),
		Ec2conn:                             ec2.New(cfg),
		Ec2instanceconnectconn:              ec2instanceconnect.New(cfg),
		Ecrconn:                             ecr.New(cfg),
		Ecsconn:                             ecs.New(cfg),
		Efsconn:                             efs.New(cfg),
		Eksconn:                             eks.New(cfg),
		Elasticacheconn:                     elasticache.New(cfg),
		Elasticbeanstalkconn:                elasticbeanstalk.New(cfg),
		Elasticinferenceconn:                elasticinference.New(cfg),
		Elasticloadbalancingconn:            elasticloadbalancing.New(cfg),
		Elasticloadbalancingv2conn:          elasticloadbalancingv2.New(cfg),
		Elasticsearchserviceconn:            elasticsearchservice.New(cfg),
		Elastictranscoderconn:               elastictranscoder.New(cfg),
		Emrconn:                             emr.New(cfg),
		Eventbridgeconn:                     eventbridge.New(cfg),
		Firehoseconn:                        firehose.New(cfg),
		Fmsconn:                             fms.New(cfg),
		Forecastconn:                        forecast.New(cfg),
		Forecastqueryconn:                   forecastquery.New(cfg),
		Frauddetectorconn:                   frauddetector.New(cfg),
		Fsxconn:                             fsx.New(cfg),
		Gameliftconn:                        gamelift.New(cfg),
		Glacierconn:                         glacier.New(cfg),
		Globalacceleratorconn:               globalaccelerator.New(cfg),
		Glueconn:                            glue.New(cfg),
		Greengrassconn:                      greengrass.New(cfg),
		Groundstationconn:                   groundstation.New(cfg),
		Guarddutyconn:                       guardduty.New(cfg),
		Healthconn:                          health.New(cfg),
		Honeycodeconn:                       honeycode.New(cfg),
		Iamconn:                             iam.New(cfg),
		Imagebuilderconn:                    imagebuilder.New(cfg),
		Inspectorconn:                       inspector.New(cfg),
		Iotconn:                             iot.New(cfg),
		Iot1clickdevicesserviceconn:         iot1clickdevicesservice.New(cfg),
		Iot1clickprojectsconn:               iot1clickprojects.New(cfg),
		Iotanalyticsconn:                    iotanalytics.New(cfg),
		Iotdataplaneconn:                    iotdataplane.New(cfg),
		Ioteventsconn:                       iotevents.New(cfg),
		Ioteventsdataconn:                   ioteventsdata.New(cfg),
		Iotjobsdataplaneconn:                iotjobsdataplane.New(cfg),
		Iotsecuretunnelingconn:              iotsecuretunneling.New(cfg),
		Iotsitewiseconn:                     iotsitewise.New(cfg),
		Iotthingsgraphconn:                  iotthingsgraph.New(cfg),
		Ivsconn:                             ivs.New(cfg),
		Kafkaconn:                           kafka.New(cfg),
		Kendraconn:                          kendra.New(cfg),
		Kinesisconn:                         kinesis.New(cfg),
		Kinesisanalyticsconn:                kinesisanalytics.New(cfg),
		Kinesisanalyticsv2conn:              kinesisanalyticsv2.New(cfg),
		Kinesisvideoconn:                    kinesisvideo.New(cfg),
		Kinesisvideoarchivedmediaconn:       kinesisvideoarchivedmedia.New(cfg),
		Kinesisvideomediaconn:               kinesisvideomedia.New(cfg),
		Kinesisvideosignalingconn:           kinesisvideosignaling.New(cfg),
		Kmsconn:                             kms.New(cfg),
		Lakeformationconn:                   lakeformation.New(cfg),
		Lambdaconn:                          lambda.New(cfg),
		Lexmodelbuildingserviceconn:         lexmodelbuildingservice.New(cfg),
		Lexruntimeserviceconn:               lexruntimeservice.New(cfg),
		Licensemanagerconn:                  licensemanager.New(cfg),
		Lightsailconn:                       lightsail.New(cfg),
		Machinelearningconn:                 machinelearning.New(cfg),
		Macieconn:                           macie.New(cfg),
		Macie2conn:                          macie2.New(cfg),
		Managedblockchainconn:               managedblockchain.New(cfg),
		Marketplacecatalogconn:              marketplacecatalog.New(cfg),
		Marketplacecommerceanalyticsconn:    marketplacecommerceanalytics.New(cfg),
		Marketplaceentitlementserviceconn:   marketplaceentitlementservice.New(cfg),
		Marketplacemeteringconn:             marketplacemetering.New(cfg),
		Mediaconnectconn:                    mediaconnect.New(cfg),
		Mediaconvertconn:                    mediaconvert.New(cfg),
		Medialiveconn:                       medialive.New(cfg),
		Mediapackageconn:                    mediapackage.New(cfg),
		Mediapackagevodconn:                 mediapackagevod.New(cfg),
		Mediastoreconn:                      mediastore.New(cfg),
		Mediastoredataconn:                  mediastoredata.New(cfg),
		Mediatailorconn:                     mediatailor.New(cfg),
		Migrationhubconn:                    migrationhub.New(cfg),
		Migrationhubconfigconn:              migrationhubconfig.New(cfg),
		Mobileconn:                          mobile.New(cfg),
		Mobileanalyticsconn:                 mobileanalytics.New(cfg),
		Mqconn:                              mq.New(cfg),
		Mturkconn:                           mturk.New(cfg),
		Neptuneconn:                         neptune.New(cfg),
		Networkmanagerconn:                  networkmanager.New(cfg),
		Opsworksconn:                        opsworks.New(cfg),
		Opsworkscmconn:                      opsworkscm.New(cfg),
		Organizationsconn:                   organizations.New(cfg),
		Outpostsconn:                        outposts.New(cfg),
		Personalizeconn:                     personalize.New(cfg),
		Personalizeeventsconn:               personalizeevents.New(cfg),
		Personalizeruntimeconn:              personalizeruntime.New(cfg),
		Piconn:                              pi.New(cfg),
		Pinpointconn:                        pinpoint.New(cfg),
		Pinpointemailconn:                   pinpointemail.New(cfg),
		Pinpointsmsvoiceconn:                pinpointsmsvoice.New(cfg),
		Pollyconn:                           polly.New(cfg),
		Pricingconn:                         pricing.New(cfg),
		Qldbconn:                            qldb.New(cfg),
		Qldbsessionconn:                     qldbsession.New(cfg),
		Quicksightconn:                      quicksight.New(cfg),
		Ramconn:                             ram.New(cfg),
		Rdsconn:                             rds.New(cfg),
		Rdsdataconn:                         rdsdata.New(cfg),
		Redshiftconn:                        redshift.New(cfg),
		Rekognitionconn:                     rekognition.New(cfg),
		Resourcegroupsconn:                  resourcegroups.New(cfg),
		Resourcegroupstaggingapiconn:        resourcegroupstaggingapi.New(cfg),
		Robomakerconn:                       robomaker.New(cfg),
		Route53conn:                         route53.New(cfg),
		Route53domainsconn:                  route53domains.New(cfg),
		Route53resolverconn:                 route53resolver.New(cfg),
		S3conn:                              s3.New(cfg),
		S3controlconn:                       s3control.New(cfg),
		Sagemakerconn:                       sagemaker.New(cfg),
		Sagemakera2iruntimeconn:             sagemakera2iruntime.New(cfg),
		Sagemakerruntimeconn:                sagemakerruntime.New(cfg),
		Savingsplansconn:                    savingsplans.New(cfg),
		Schemasconn:                         schemas.New(cfg),
		Secretsmanagerconn:                  secretsmanager.New(cfg),
		Securityhubconn:                     securityhub.New(cfg),
		Serverlessapplicationrepositoryconn: serverlessapplicationrepository.New(cfg),
		Servicecatalogconn:                  servicecatalog.New(cfg),
		Servicediscoveryconn:                servicediscovery.New(cfg),
		Servicequotasconn:                   servicequotas.New(cfg),
		Sesconn:                             ses.New(cfg),
		Sesv2conn:                           sesv2.New(cfg),
		Sfnconn:                             sfn.New(cfg),
		Shieldconn:                          shield.New(cfg),
		Signerconn:                          signer.New(cfg),
		Simpledbconn:                        simpledb.New(cfg),
		Smsconn:                             sms.New(cfg),
		Snowballconn:                        snowball.New(cfg),
		Snsconn:                             sns.New(cfg),
		Sqsconn:                             sqs.New(cfg),
		Ssmconn:                             ssm.New(cfg),
		Ssoconn:                             sso.New(cfg),
		Ssooidcconn:                         ssooidc.New(cfg),
		Storagegatewayconn:                  storagegateway.New(cfg),
		Stsconn:                             sts.New(cfg),
		Supportconn:                         support.New(cfg),
		Swfconn:                             swf.New(cfg),
		Syntheticsconn:                      synthetics.New(cfg),
		Textractconn:                        textract.New(cfg),
		Transcribeconn:                      transcribe.New(cfg),
		Transferconn:                        transfer.New(cfg),
		Translateconn:                       translate.New(cfg),
		Wafconn:                             waf.New(cfg),
		Wafregionalconn:                     wafregional.New(cfg),
		Wafv2conn:                           wafv2.New(cfg),
		Workdocsconn:                        workdocs.New(cfg),
		Worklinkconn:                        worklink.New(cfg),
		Workmailconn:                        workmail.New(cfg),
		Workmailmessageflowconn:             workmailmessageflow.New(cfg),
		Workspacesconn:                      workspaces.New(cfg),
		Xrayconn:                            xray.New(cfg),
	}

	profile, _, err := external.GetSharedConfigProfile(configs)
	if err != nil {
		return nil, err
	}

	client.Profile = profile
	client.Region = cfg.Region

	log.WithFields(log.Fields{
		"profile": profile,
		"region":  cfg.Region,
	}).Debugf("created new instance of AWS client")

	return client, nil
}

// SetAccountID populates the AccountID field of the client.
func (client *Client) SetAccountID() error {
	req := client.Stsconn.GetCallerIdentityRequest(&sts.GetCallerIdentityInput{})
	resp, err := req.Send(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get caller identity: %s", err)
	}

	client.AccountID = *resp.Account

	return nil
}
