// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// Export sub-modules:
import * as amplify from "./amplify";
import * as apigateway from "./apigateway";
import * as appflow from "./appflow";
import * as applicationinsights from "./applicationinsights";
import * as apprunner from "./apprunner";
import * as athena from "./athena";
import * as auditmanager from "./auditmanager";
import * as budgets from "./budgets";
import * as cassandra from "./cassandra";
import * as ce from "./ce";
import * as cloudformation from "./cloudformation";
import * as cloudtrail from "./cloudtrail";
import * as codeguruprofiler from "./codeguruprofiler";
import * as codegurureviewer from "./codegurureviewer";
import * as codestarnotifications from "./codestarnotifications";
import * as connect from "./connect";
import * as customerprofiles from "./customerprofiles";
import * as databrew from "./databrew";
import * as datasync from "./datasync";
import * as devopsguru from "./devopsguru";
import * as ec2 from "./ec2";
import * as ecr from "./ecr";
import * as ecs from "./ecs";
import * as eks from "./eks";
import * as elasticache from "./elasticache";
import * as emr from "./emr";
import * as events from "./events";
import * as finspace from "./finspace";
import * as frauddetector from "./frauddetector";
import * as gamelift from "./gamelift";
import * as globalaccelerator from "./globalaccelerator";
import * as glue from "./glue";
import * as greengrassv2 from "./greengrassv2";
import * as groundstation from "./groundstation";
import * as healthlake from "./healthlake";
import * as imagebuilder from "./imagebuilder";
import * as iot from "./iot";
import * as iotevents from "./iotevents";
import * as iotsitewise from "./iotsitewise";
import * as iotwireless from "./iotwireless";
import * as ivs from "./ivs";
import * as kendra from "./kendra";
import * as kinesis from "./kinesis";
import * as kinesisfirehose from "./kinesisfirehose";
import * as kms from "./kms";
import * as lambda from "./lambda";
import * as location from "./location";
import * as lookoutmetrics from "./lookoutmetrics";
import * as macie from "./macie";
import * as mediaconnect from "./mediaconnect";
import * as mediapackage from "./mediapackage";
import * as mwaa from "./mwaa";
import * as networkfirewall from "./networkfirewall";
import * as quicksight from "./quicksight";
import * as rds from "./rds";
import * as resourcegroups from "./resourcegroups";
import * as robomaker from "./robomaker";
import * as route53 from "./route53";
import * as route53recoverycontrol from "./route53recoverycontrol";
import * as route53resolver from "./route53resolver";
import * as s3 from "./s3";
import * as s3outposts from "./s3outposts";
import * as sagemaker from "./sagemaker";
import * as servicecatalog from "./servicecatalog";
import * as servicecatalogappregistry from "./servicecatalogappregistry";
import * as signer from "./signer";
import * as ssm from "./ssm";
import * as ssmcontacts from "./ssmcontacts";
import * as ssmincidents from "./ssmincidents";
import * as sso from "./sso";
import * as stepfunctions from "./stepfunctions";
import * as wafv2 from "./wafv2";
import * as workspaces from "./workspaces";

export {
    amplify,
    apigateway,
    appflow,
    applicationinsights,
    apprunner,
    athena,
    auditmanager,
    budgets,
    cassandra,
    ce,
    cloudformation,
    cloudtrail,
    codeguruprofiler,
    codegurureviewer,
    codestarnotifications,
    connect,
    customerprofiles,
    databrew,
    datasync,
    devopsguru,
    ec2,
    ecr,
    ecs,
    eks,
    elasticache,
    emr,
    events,
    finspace,
    frauddetector,
    gamelift,
    globalaccelerator,
    glue,
    greengrassv2,
    groundstation,
    healthlake,
    imagebuilder,
    iot,
    iotevents,
    iotsitewise,
    iotwireless,
    ivs,
    kendra,
    kinesis,
    kinesisfirehose,
    kms,
    lambda,
    location,
    lookoutmetrics,
    macie,
    mediaconnect,
    mediapackage,
    mwaa,
    networkfirewall,
    quicksight,
    rds,
    resourcegroups,
    robomaker,
    route53,
    route53recoverycontrol,
    route53resolver,
    s3,
    s3outposts,
    sagemaker,
    servicecatalog,
    servicecatalogappregistry,
    signer,
    ssm,
    ssmcontacts,
    ssmincidents,
    sso,
    stepfunctions,
    wafv2,
    workspaces,
};

export const Region = {
    /**
     * Africa (Cape Town)
     */
    AFSouth1: "af-south-1",
    /**
     * Asia Pacific (Hong Kong)
     */
    APEast1: "ap-east-1",
    /**
     * Asia Pacific (Tokyo)
     */
    APNortheast1: "ap-northeast-1",
    /**
     * Asia Pacific (Seoul)
     */
    APNortheast2: "ap-northeast-2",
    /**
     * Asia Pacific (Osaka)
     */
    APNortheast3: "ap-northeast-3",
    /**
     * Asia Pacific (Mumbai)
     */
    APSouth1: "ap-south-1",
    /**
     * Asia Pacific (Singapore)
     */
    APSoutheast1: "ap-southeast-1",
    /**
     * Asia Pacific (Sydney)
     */
    APSoutheast2: "ap-southeast-2",
    /**
     * Canada (Central)
     */
    CACentral: "ca-central-1",
    /**
     * China (Beijing)
     */
    CNNorth1: "cn-north-1",
    /**
     * China (Ningxia)
     */
    CNNorthwest1: "cn-northwest-1",
    /**
     * Europe (Frankfurt)
     */
    EUCentral1: "eu-central-1",
    /**
     * Europe (Stockholm)
     */
    EUNorth1: "eu-north-1",
    /**
     * Europe (Ireland)
     */
    EUWest1: "eu-west-1",
    /**
     * Europe (London)
     */
    EUWest2: "eu-west-2",
    /**
     * Europe (Paris)
     */
    EUWest3: "eu-west-3",
    /**
     * Europe (Milan)
     */
    EUSouth1: "eu-south-1",
    /**
     * Middle East (Bahrain)
     */
    MESouth1: "me-south-1",
    /**
     * South America (São Paulo)
     */
    SAEast1: "sa-east-1",
    /**
     * AWS GovCloud (US-East)
     */
    USGovEast1: "us-gov-east-1",
    /**
     * AWS GovCloud (US-West)
     */
    USGovWest1: "us-gov-west-1",
    /**
     * US East (N. Virginia)
     */
    USEast1: "us-east-1",
    /**
     * US East (Ohio)
     */
    USEast2: "us-east-2",
    /**
     * US West (N. California)
     */
    USWest1: "us-west-1",
    /**
     * US West (Oregon)
     */
    USWest2: "us-west-2",
} as const;

/**
 * A Region represents any valid Amazon region that may be targeted with deployments.
 */
export type Region = (typeof Region)[keyof typeof Region];
