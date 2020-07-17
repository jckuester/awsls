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
	Eksconn                             *eks.Client
	Clouddirectoryconn                  *clouddirectory.Client
	Swfconn                             *swf.Client
	Rdsconn                             *rds.Client
	Marketplacemeteringconn             *marketplacemetering.Client
	Marketplaceentitlementserviceconn   *marketplaceentitlementservice.Client
	Codepipelineconn                    *codepipeline.Client
	Mediastoredataconn                  *mediastoredata.Client
	Docdbconn                           *docdb.Client
	Cloudsearchdomainconn               *cloudsearchdomain.Client
	Batchconn                           *batch.Client
	Licensemanagerconn                  *licensemanager.Client
	Kinesisvideoarchivedmediaconn       *kinesisvideoarchivedmedia.Client
	Iot1clickprojectsconn               *iot1clickprojects.Client
	Elasticloadbalancingv2conn          *elasticloadbalancingv2.Client
	Ecsconn                             *ecs.Client
	Elasticinferenceconn                *elasticinference.Client
	Sesv2conn                           *sesv2.Client
	Personalizeruntimeconn              *personalizeruntime.Client
	Codestarconn                        *codestar.Client
	Detectiveconn                       *detective.Client
	Sfnconn                             *sfn.Client
	Ssoconn                             *sso.Client
	Servicediscoveryconn                *servicediscovery.Client
	Lexruntimeserviceconn               *lexruntimeservice.Client
	Healthconn                          *health.Client
	Guarddutyconn                       *guardduty.Client
	Directoryserviceconn                *directoryservice.Client
	Athenaconn                          *athena.Client
	Mediapackageconn                    *mediapackage.Client
	Kinesisconn                         *kinesis.Client
	Kafkaconn                           *kafka.Client
	Ramconn                             *ram.Client
	Dynamodbconn                        *dynamodb.Client
	Appmeshconn                         *appmesh.Client
	Alexaforbusinessconn                *alexaforbusiness.Client
	Stsconn                             *sts.Client
	Resourcegroupstaggingapiconn        *resourcegroupstaggingapi.Client
	Pollyconn                           *polly.Client
	Macieconn                           *macie.Client
	Configserviceconn                   *configservice.Client
	Codebuildconn                       *codebuild.Client
	Kinesisvideomediaconn               *kinesisvideomedia.Client
	Applicationinsightsconn             *applicationinsights.Client
	Route53conn                         *route53.Client
	Resourcegroupsconn                  *resourcegroups.Client
	Qldbconn                            *qldb.Client
	Mediatailorconn                     *mediatailor.Client
	Lexmodelbuildingserviceconn         *lexmodelbuildingservice.Client
	Applicationdiscoveryserviceconn     *applicationdiscoveryservice.Client
	Simpledbconn                        *simpledb.Client
	Pricingconn                         *pricing.Client
	Kinesisvideoconn                    *kinesisvideo.Client
	Forecastconn                        *forecast.Client
	Securityhubconn                     *securityhub.Client
	Redshiftconn                        *redshift.Client
	Mturkconn                           *mturk.Client
	Datasyncconn                        *datasync.Client
	Migrationhubconfigconn              *migrationhubconfig.Client
	Ec2conn                             *ec2.Client
	Directconnectconn                   *directconnect.Client
	Pinpointsmsvoiceconn                *pinpointsmsvoice.Client
	Eventbridgeconn                     *eventbridge.Client
	Backupconn                          *backup.Client
	Servicecatalogconn                  *servicecatalog.Client
	Cloudwatchlogsconn                  *cloudwatchlogs.Client
	Iotjobsdataplaneconn                *iotjobsdataplane.Client
	Costandusagereportserviceconn       *costandusagereportservice.Client
	Budgetsconn                         *budgets.Client
	Wafregionalconn                     *wafregional.Client
	Qldbsessionconn                     *qldbsession.Client
	Syntheticsconn                      *synthetics.Client
	Ssooidcconn                         *ssooidc.Client
	Personalizeconn                     *personalize.Client
	Sesconn                             *ses.Client
	Cloudhsmconn                        *cloudhsm.Client
	Xrayconn                            *xray.Client
	S3conn                              *s3.Client
	Medialiveconn                       *medialive.Client
	Iotsitewiseconn                     *iotsitewise.Client
	Connectparticipantconn              *connectparticipant.Client
	Route53domainsconn                  *route53domains.Client
	Ebsconn                             *ebs.Client
	Chimeconn                           *chime.Client
	Comprehendmedicalconn               *comprehendmedical.Client
	Workmailconn                        *workmail.Client
	Wafconn                             *waf.Client
	Storagegatewayconn                  *storagegateway.Client
	Servicequotasconn                   *servicequotas.Client
	Macie2conn                          *macie2.Client
	Machinelearningconn                 *machinelearning.Client
	Kmsconn                             *kms.Client
	Cloud9conn                          *cloud9.Client
	Workdocsconn                        *workdocs.Client
	Appstreamconn                       *appstream.Client
	Route53resolverconn                 *route53resolver.Client
	Iotconn                             *iot.Client
	Iotdataplaneconn                    *iotdataplane.Client
	Datapipelineconn                    *datapipeline.Client
	Workspacesconn                      *workspaces.Client
	Sagemakerruntimeconn                *sagemakerruntime.Client
	Emrconn                             *emr.Client
	Dataexchangeconn                    *dataexchange.Client
	Codestarnotificationsconn           *codestarnotifications.Client
	Wafv2conn                           *wafv2.Client
	Shieldconn                          *shield.Client
	Pinpointconn                        *pinpoint.Client
	Serverlessapplicationrepositoryconn *serverlessapplicationrepository.Client
	Lambdaconn                          *lambda.Client
	Fsxconn                             *fsx.Client
	Firehoseconn                        *firehose.Client
	Codecommitconn                      *codecommit.Client
	Translateconn                       *translate.Client
	Kinesisanalyticsconn                *kinesisanalytics.Client
	Lakeformationconn                   *lakeformation.Client
	Iamconn                             *iam.Client
	Gameliftconn                        *gamelift.Client
	Schemasconn                         *schemas.Client
	Marketplacecatalogconn              *marketplacecatalog.Client
	Sqsconn                             *sqs.Client
	Frauddetectorconn                   *frauddetector.Client
	Elasticloadbalancingconn            *elasticloadbalancing.Client
	Cloudformationconn                  *cloudformation.Client
	Apigatewayv2conn                    *apigatewayv2.Client
	Apigatewayconn                      *apigateway.Client
	Mediapackagevodconn                 *mediapackagevod.Client
	Cloudwatcheventsconn                *cloudwatchevents.Client
	Elasticbeanstalkconn                *elasticbeanstalk.Client
	Codedeployconn                      *codedeploy.Client
	Apigatewaymanagementapiconn         *apigatewaymanagementapi.Client
	S3controlconn                       *s3control.Client
	Personalizeeventsconn               *personalizeevents.Client
	Organizationsconn                   *organizations.Client
	Opsworksconn                        *opsworks.Client
	Mobileconn                          *mobile.Client
	Devicefarmconn                      *devicefarm.Client
	Signerconn                          *signer.Client
	Marketplacecommerceanalyticsconn    *marketplacecommerceanalytics.Client
	Greengrassconn                      *greengrass.Client
	Cognitosyncconn                     *cognitosync.Client
	Textractconn                        *textract.Client
	Imagebuilderconn                    *imagebuilder.Client
	Computeoptimizerconn                *computeoptimizer.Client
	Opsworkscmconn                      *opsworkscm.Client
	Globalacceleratorconn               *globalaccelerator.Client
	Acmpcaconn                          *acmpca.Client
	Groundstationconn                   *groundstation.Client
	Codeguruprofilerconn                *codeguruprofiler.Client
	Cloudtrailconn                      *cloudtrail.Client
	Amplifyconn                         *amplify.Client
	Robomakerconn                       *robomaker.Client
	Piconn                              *pi.Client
	Mediastoreconn                      *mediastore.Client
	Glueconn                            *glue.Client
	Codestarconnectionsconn             *codestarconnections.Client
	Outpostsconn                        *outposts.Client
	Mediaconnectconn                    *mediaconnect.Client
	Elasticsearchserviceconn            *elasticsearchservice.Client
	Elastictranscoderconn               *elastictranscoder.Client
	Ec2instanceconnectconn              *ec2instanceconnect.Client
	Comprehendconn                      *comprehend.Client
	Acmconn                             *acm.Client
	Supportconn                         *support.Client
	Sagemakerconn                       *sagemaker.Client
	Codegurureviewerconn                *codegurureviewer.Client
	Ssmconn                             *ssm.Client
	Sagemakera2iruntimeconn             *sagemakera2iruntime.Client
	Worklinkconn                        *worklink.Client
	Snowballconn                        *snowball.Client
	Kinesisanalyticsv2conn              *kinesisanalyticsv2.Client
	Cloudhsmv2conn                      *cloudhsmv2.Client
	Smsconn                             *sms.Client
	Mobileanalyticsconn                 *mobileanalytics.Client
	Managedblockchainconn               *managedblockchain.Client
	Connectconn                         *connect.Client
	Kendraconn                          *kendra.Client
	Iotsecuretunnelingconn              *iotsecuretunneling.Client
	Cognitoidentityconn                 *cognitoidentity.Client
	Cloudsearchconn                     *cloudsearch.Client
	Migrationhubconn                    *migrationhub.Client
	Appconfigconn                       *appconfig.Client
	Accessanalyzerconn                  *accessanalyzer.Client
	Savingsplansconn                    *savingsplans.Client
	Neptuneconn                         *neptune.Client
	Forecastqueryconn                   *forecastquery.Client
	Applicationautoscalingconn          *applicationautoscaling.Client
	Pinpointemailconn                   *pinpointemail.Client
	Daxconn                             *dax.Client
	Transcribeconn                      *transcribe.Client
	Fmsconn                             *fms.Client
	Mqconn                              *mq.Client
	Iotanalyticsconn                    *iotanalytics.Client
	Glacierconn                         *glacier.Client
	Autoscalingplansconn                *autoscalingplans.Client
	Cognitoidentityproviderconn         *cognitoidentityprovider.Client
	Networkmanagerconn                  *networkmanager.Client
	Ioteventsdataconn                   *ioteventsdata.Client
	Inspectorconn                       *inspector.Client
	Ecrconn                             *ecr.Client
	Databasemigrationserviceconn        *databasemigrationservice.Client
	Workmailmessageflowconn             *workmailmessageflow.Client
	Secretsmanagerconn                  *secretsmanager.Client
	Quicksightconn                      *quicksight.Client
	Ioteventsconn                       *iotevents.Client
	Iot1clickdevicesserviceconn         *iot1clickdevicesservice.Client
	Appsyncconn                         *appsync.Client
	Rekognitionconn                     *rekognition.Client
	Efsconn                             *efs.Client
	Snsconn                             *sns.Client
	Rdsdataconn                         *rdsdata.Client
	Cloudfrontconn                      *cloudfront.Client
	Dynamodbstreamsconn                 *dynamodbstreams.Client
	Cloudwatchconn                      *cloudwatch.Client
	Mediaconvertconn                    *mediaconvert.Client
	Elasticacheconn                     *elasticache.Client
	Dlmconn                             *dlm.Client
	Costexplorerconn                    *costexplorer.Client
	Transferconn                        *transfer.Client
	Lightsailconn                       *lightsail.Client
	Kinesisvideosignalingconn           *kinesisvideosignaling.Client
	Iotthingsgraphconn                  *iotthingsgraph.Client
	Autoscalingconn                     *autoscaling.Client
}

func NewClient(configs ...external.Config) (*Client, error) {
	cfg, err := external.LoadDefaultAWSConfig(configs...)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %s", err)
	}

	client := &Client{
		Eksconn:                             eks.New(cfg),
		Clouddirectoryconn:                  clouddirectory.New(cfg),
		Swfconn:                             swf.New(cfg),
		Rdsconn:                             rds.New(cfg),
		Marketplacemeteringconn:             marketplacemetering.New(cfg),
		Marketplaceentitlementserviceconn:   marketplaceentitlementservice.New(cfg),
		Codepipelineconn:                    codepipeline.New(cfg),
		Mediastoredataconn:                  mediastoredata.New(cfg),
		Docdbconn:                           docdb.New(cfg),
		Cloudsearchdomainconn:               cloudsearchdomain.New(cfg),
		Batchconn:                           batch.New(cfg),
		Licensemanagerconn:                  licensemanager.New(cfg),
		Kinesisvideoarchivedmediaconn:       kinesisvideoarchivedmedia.New(cfg),
		Iot1clickprojectsconn:               iot1clickprojects.New(cfg),
		Elasticloadbalancingv2conn:          elasticloadbalancingv2.New(cfg),
		Ecsconn:                             ecs.New(cfg),
		Elasticinferenceconn:                elasticinference.New(cfg),
		Sesv2conn:                           sesv2.New(cfg),
		Personalizeruntimeconn:              personalizeruntime.New(cfg),
		Codestarconn:                        codestar.New(cfg),
		Detectiveconn:                       detective.New(cfg),
		Sfnconn:                             sfn.New(cfg),
		Ssoconn:                             sso.New(cfg),
		Servicediscoveryconn:                servicediscovery.New(cfg),
		Lexruntimeserviceconn:               lexruntimeservice.New(cfg),
		Healthconn:                          health.New(cfg),
		Guarddutyconn:                       guardduty.New(cfg),
		Directoryserviceconn:                directoryservice.New(cfg),
		Athenaconn:                          athena.New(cfg),
		Mediapackageconn:                    mediapackage.New(cfg),
		Kinesisconn:                         kinesis.New(cfg),
		Kafkaconn:                           kafka.New(cfg),
		Ramconn:                             ram.New(cfg),
		Dynamodbconn:                        dynamodb.New(cfg),
		Appmeshconn:                         appmesh.New(cfg),
		Alexaforbusinessconn:                alexaforbusiness.New(cfg),
		Stsconn:                             sts.New(cfg),
		Resourcegroupstaggingapiconn:        resourcegroupstaggingapi.New(cfg),
		Pollyconn:                           polly.New(cfg),
		Macieconn:                           macie.New(cfg),
		Configserviceconn:                   configservice.New(cfg),
		Codebuildconn:                       codebuild.New(cfg),
		Kinesisvideomediaconn:               kinesisvideomedia.New(cfg),
		Applicationinsightsconn:             applicationinsights.New(cfg),
		Route53conn:                         route53.New(cfg),
		Resourcegroupsconn:                  resourcegroups.New(cfg),
		Qldbconn:                            qldb.New(cfg),
		Mediatailorconn:                     mediatailor.New(cfg),
		Lexmodelbuildingserviceconn:         lexmodelbuildingservice.New(cfg),
		Applicationdiscoveryserviceconn:     applicationdiscoveryservice.New(cfg),
		Simpledbconn:                        simpledb.New(cfg),
		Pricingconn:                         pricing.New(cfg),
		Kinesisvideoconn:                    kinesisvideo.New(cfg),
		Forecastconn:                        forecast.New(cfg),
		Securityhubconn:                     securityhub.New(cfg),
		Redshiftconn:                        redshift.New(cfg),
		Mturkconn:                           mturk.New(cfg),
		Datasyncconn:                        datasync.New(cfg),
		Migrationhubconfigconn:              migrationhubconfig.New(cfg),
		Ec2conn:                             ec2.New(cfg),
		Directconnectconn:                   directconnect.New(cfg),
		Pinpointsmsvoiceconn:                pinpointsmsvoice.New(cfg),
		Eventbridgeconn:                     eventbridge.New(cfg),
		Backupconn:                          backup.New(cfg),
		Servicecatalogconn:                  servicecatalog.New(cfg),
		Cloudwatchlogsconn:                  cloudwatchlogs.New(cfg),
		Iotjobsdataplaneconn:                iotjobsdataplane.New(cfg),
		Costandusagereportserviceconn:       costandusagereportservice.New(cfg),
		Budgetsconn:                         budgets.New(cfg),
		Wafregionalconn:                     wafregional.New(cfg),
		Qldbsessionconn:                     qldbsession.New(cfg),
		Syntheticsconn:                      synthetics.New(cfg),
		Ssooidcconn:                         ssooidc.New(cfg),
		Personalizeconn:                     personalize.New(cfg),
		Sesconn:                             ses.New(cfg),
		Cloudhsmconn:                        cloudhsm.New(cfg),
		Xrayconn:                            xray.New(cfg),
		S3conn:                              s3.New(cfg),
		Medialiveconn:                       medialive.New(cfg),
		Iotsitewiseconn:                     iotsitewise.New(cfg),
		Connectparticipantconn:              connectparticipant.New(cfg),
		Route53domainsconn:                  route53domains.New(cfg),
		Ebsconn:                             ebs.New(cfg),
		Chimeconn:                           chime.New(cfg),
		Comprehendmedicalconn:               comprehendmedical.New(cfg),
		Workmailconn:                        workmail.New(cfg),
		Wafconn:                             waf.New(cfg),
		Storagegatewayconn:                  storagegateway.New(cfg),
		Servicequotasconn:                   servicequotas.New(cfg),
		Macie2conn:                          macie2.New(cfg),
		Machinelearningconn:                 machinelearning.New(cfg),
		Kmsconn:                             kms.New(cfg),
		Cloud9conn:                          cloud9.New(cfg),
		Workdocsconn:                        workdocs.New(cfg),
		Appstreamconn:                       appstream.New(cfg),
		Route53resolverconn:                 route53resolver.New(cfg),
		Iotconn:                             iot.New(cfg),
		Iotdataplaneconn:                    iotdataplane.New(cfg),
		Datapipelineconn:                    datapipeline.New(cfg),
		Workspacesconn:                      workspaces.New(cfg),
		Sagemakerruntimeconn:                sagemakerruntime.New(cfg),
		Emrconn:                             emr.New(cfg),
		Dataexchangeconn:                    dataexchange.New(cfg),
		Codestarnotificationsconn:           codestarnotifications.New(cfg),
		Wafv2conn:                           wafv2.New(cfg),
		Shieldconn:                          shield.New(cfg),
		Pinpointconn:                        pinpoint.New(cfg),
		Serverlessapplicationrepositoryconn: serverlessapplicationrepository.New(cfg),
		Lambdaconn:                          lambda.New(cfg),
		Fsxconn:                             fsx.New(cfg),
		Firehoseconn:                        firehose.New(cfg),
		Codecommitconn:                      codecommit.New(cfg),
		Translateconn:                       translate.New(cfg),
		Kinesisanalyticsconn:                kinesisanalytics.New(cfg),
		Lakeformationconn:                   lakeformation.New(cfg),
		Iamconn:                             iam.New(cfg),
		Gameliftconn:                        gamelift.New(cfg),
		Schemasconn:                         schemas.New(cfg),
		Marketplacecatalogconn:              marketplacecatalog.New(cfg),
		Sqsconn:                             sqs.New(cfg),
		Frauddetectorconn:                   frauddetector.New(cfg),
		Elasticloadbalancingconn:            elasticloadbalancing.New(cfg),
		Cloudformationconn:                  cloudformation.New(cfg),
		Apigatewayv2conn:                    apigatewayv2.New(cfg),
		Apigatewayconn:                      apigateway.New(cfg),
		Mediapackagevodconn:                 mediapackagevod.New(cfg),
		Cloudwatcheventsconn:                cloudwatchevents.New(cfg),
		Elasticbeanstalkconn:                elasticbeanstalk.New(cfg),
		Codedeployconn:                      codedeploy.New(cfg),
		Apigatewaymanagementapiconn:         apigatewaymanagementapi.New(cfg),
		S3controlconn:                       s3control.New(cfg),
		Personalizeeventsconn:               personalizeevents.New(cfg),
		Organizationsconn:                   organizations.New(cfg),
		Opsworksconn:                        opsworks.New(cfg),
		Mobileconn:                          mobile.New(cfg),
		Devicefarmconn:                      devicefarm.New(cfg),
		Signerconn:                          signer.New(cfg),
		Marketplacecommerceanalyticsconn:    marketplacecommerceanalytics.New(cfg),
		Greengrassconn:                      greengrass.New(cfg),
		Cognitosyncconn:                     cognitosync.New(cfg),
		Textractconn:                        textract.New(cfg),
		Imagebuilderconn:                    imagebuilder.New(cfg),
		Computeoptimizerconn:                computeoptimizer.New(cfg),
		Opsworkscmconn:                      opsworkscm.New(cfg),
		Globalacceleratorconn:               globalaccelerator.New(cfg),
		Acmpcaconn:                          acmpca.New(cfg),
		Groundstationconn:                   groundstation.New(cfg),
		Codeguruprofilerconn:                codeguruprofiler.New(cfg),
		Cloudtrailconn:                      cloudtrail.New(cfg),
		Amplifyconn:                         amplify.New(cfg),
		Robomakerconn:                       robomaker.New(cfg),
		Piconn:                              pi.New(cfg),
		Mediastoreconn:                      mediastore.New(cfg),
		Glueconn:                            glue.New(cfg),
		Codestarconnectionsconn:             codestarconnections.New(cfg),
		Outpostsconn:                        outposts.New(cfg),
		Mediaconnectconn:                    mediaconnect.New(cfg),
		Elasticsearchserviceconn:            elasticsearchservice.New(cfg),
		Elastictranscoderconn:               elastictranscoder.New(cfg),
		Ec2instanceconnectconn:              ec2instanceconnect.New(cfg),
		Comprehendconn:                      comprehend.New(cfg),
		Acmconn:                             acm.New(cfg),
		Supportconn:                         support.New(cfg),
		Sagemakerconn:                       sagemaker.New(cfg),
		Codegurureviewerconn:                codegurureviewer.New(cfg),
		Ssmconn:                             ssm.New(cfg),
		Sagemakera2iruntimeconn:             sagemakera2iruntime.New(cfg),
		Worklinkconn:                        worklink.New(cfg),
		Snowballconn:                        snowball.New(cfg),
		Kinesisanalyticsv2conn:              kinesisanalyticsv2.New(cfg),
		Cloudhsmv2conn:                      cloudhsmv2.New(cfg),
		Smsconn:                             sms.New(cfg),
		Mobileanalyticsconn:                 mobileanalytics.New(cfg),
		Managedblockchainconn:               managedblockchain.New(cfg),
		Connectconn:                         connect.New(cfg),
		Kendraconn:                          kendra.New(cfg),
		Iotsecuretunnelingconn:              iotsecuretunneling.New(cfg),
		Cognitoidentityconn:                 cognitoidentity.New(cfg),
		Cloudsearchconn:                     cloudsearch.New(cfg),
		Migrationhubconn:                    migrationhub.New(cfg),
		Appconfigconn:                       appconfig.New(cfg),
		Accessanalyzerconn:                  accessanalyzer.New(cfg),
		Savingsplansconn:                    savingsplans.New(cfg),
		Neptuneconn:                         neptune.New(cfg),
		Forecastqueryconn:                   forecastquery.New(cfg),
		Applicationautoscalingconn:          applicationautoscaling.New(cfg),
		Pinpointemailconn:                   pinpointemail.New(cfg),
		Daxconn:                             dax.New(cfg),
		Transcribeconn:                      transcribe.New(cfg),
		Fmsconn:                             fms.New(cfg),
		Mqconn:                              mq.New(cfg),
		Iotanalyticsconn:                    iotanalytics.New(cfg),
		Glacierconn:                         glacier.New(cfg),
		Autoscalingplansconn:                autoscalingplans.New(cfg),
		Cognitoidentityproviderconn:         cognitoidentityprovider.New(cfg),
		Networkmanagerconn:                  networkmanager.New(cfg),
		Ioteventsdataconn:                   ioteventsdata.New(cfg),
		Inspectorconn:                       inspector.New(cfg),
		Ecrconn:                             ecr.New(cfg),
		Databasemigrationserviceconn:        databasemigrationservice.New(cfg),
		Workmailmessageflowconn:             workmailmessageflow.New(cfg),
		Secretsmanagerconn:                  secretsmanager.New(cfg),
		Quicksightconn:                      quicksight.New(cfg),
		Ioteventsconn:                       iotevents.New(cfg),
		Iot1clickdevicesserviceconn:         iot1clickdevicesservice.New(cfg),
		Appsyncconn:                         appsync.New(cfg),
		Rekognitionconn:                     rekognition.New(cfg),
		Efsconn:                             efs.New(cfg),
		Snsconn:                             sns.New(cfg),
		Rdsdataconn:                         rdsdata.New(cfg),
		Cloudfrontconn:                      cloudfront.New(cfg),
		Dynamodbstreamsconn:                 dynamodbstreams.New(cfg),
		Cloudwatchconn:                      cloudwatch.New(cfg),
		Mediaconvertconn:                    mediaconvert.New(cfg),
		Elasticacheconn:                     elasticache.New(cfg),
		Dlmconn:                             dlm.New(cfg),
		Costexplorerconn:                    costexplorer.New(cfg),
		Transferconn:                        transfer.New(cfg),
		Lightsailconn:                       lightsail.New(cfg),
		Kinesisvideosignalingconn:           kinesisvideosignaling.New(cfg),
		Iotthingsgraphconn:                  iotthingsgraph.New(cfg),
		Autoscalingconn:                     autoscaling.New(cfg),
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
