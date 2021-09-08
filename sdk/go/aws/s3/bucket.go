// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
type Bucket struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accelerateconfiguration
	AccelerateConfiguration BucketAccelerateConfigurationPtrOutput `pulumi:"accelerateConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol
	AccessControl pulumi.StringPtrOutput `pulumi:"accessControl"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-analyticsconfigurations
	AnalyticsConfigurations BucketAnalyticsConfigurationArrayOutput `pulumi:"analyticsConfigurations"`
	Arn                     pulumi.StringOutput                     `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-bucketencryption
	BucketEncryption BucketBucketEncryptionPtrOutput `pulumi:"bucketEncryption"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name
	BucketName pulumi.StringPtrOutput `pulumi:"bucketName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig
	CorsConfiguration   BucketCorsConfigurationPtrOutput `pulumi:"corsConfiguration"`
	DomainName          pulumi.StringOutput              `pulumi:"domainName"`
	DualStackDomainName pulumi.StringOutput              `pulumi:"dualStackDomainName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-intelligenttieringconfigurations
	IntelligentTieringConfigurations BucketIntelligentTieringConfigurationArrayOutput `pulumi:"intelligentTieringConfigurations"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-inventoryconfigurations
	InventoryConfigurations BucketInventoryConfigurationArrayOutput `pulumi:"inventoryConfigurations"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig
	LifecycleConfiguration BucketLifecycleConfigurationPtrOutput `pulumi:"lifecycleConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig
	LoggingConfiguration BucketLoggingConfigurationPtrOutput `pulumi:"loggingConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-metricsconfigurations
	MetricsConfigurations BucketMetricsConfigurationArrayOutput `pulumi:"metricsConfigurations"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification
	NotificationConfiguration BucketNotificationConfigurationPtrOutput `pulumi:"notificationConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockconfiguration
	ObjectLockConfiguration BucketObjectLockConfigurationPtrOutput `pulumi:"objectLockConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockenabled
	ObjectLockEnabled pulumi.BoolPtrOutput `pulumi:"objectLockEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-ownershipcontrols
	OwnershipControls BucketOwnershipControlsPtrOutput `pulumi:"ownershipControls"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-publicaccessblockconfiguration
	PublicAccessBlockConfiguration BucketPublicAccessBlockConfigurationPtrOutput `pulumi:"publicAccessBlockConfiguration"`
	RegionalDomainName             pulumi.StringOutput                           `pulumi:"regionalDomainName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration
	ReplicationConfiguration BucketReplicationConfigurationPtrOutput `pulumi:"replicationConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning
	VersioningConfiguration BucketVersioningConfigurationPtrOutput `pulumi:"versioningConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration
	WebsiteConfiguration BucketWebsiteConfigurationPtrOutput `pulumi:"websiteConfiguration"`
	WebsiteURL           pulumi.StringOutput                 `pulumi:"websiteURL"`
}

// NewBucket registers a new resource with the given unique name, arguments, and options.
func NewBucket(ctx *pulumi.Context,
	name string, args *BucketArgs, opts ...pulumi.ResourceOption) (*Bucket, error) {
	if args == nil {
		args = &BucketArgs{}
	}

	var resource Bucket
	err := ctx.RegisterResource("aws-native:s3:Bucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucket gets an existing Bucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketState, opts ...pulumi.ResourceOption) (*Bucket, error) {
	var resource Bucket
	err := ctx.ReadResource("aws-native:s3:Bucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bucket resources.
type bucketState struct {
}

type BucketState struct {
}

func (BucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketState)(nil)).Elem()
}

type bucketArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accelerateconfiguration
	AccelerateConfiguration *BucketAccelerateConfiguration `pulumi:"accelerateConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol
	AccessControl *string `pulumi:"accessControl"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-analyticsconfigurations
	AnalyticsConfigurations []BucketAnalyticsConfiguration `pulumi:"analyticsConfigurations"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-bucketencryption
	BucketEncryption *BucketBucketEncryption `pulumi:"bucketEncryption"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name
	BucketName *string `pulumi:"bucketName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig
	CorsConfiguration *BucketCorsConfiguration `pulumi:"corsConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-intelligenttieringconfigurations
	IntelligentTieringConfigurations []BucketIntelligentTieringConfiguration `pulumi:"intelligentTieringConfigurations"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-inventoryconfigurations
	InventoryConfigurations []BucketInventoryConfiguration `pulumi:"inventoryConfigurations"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig
	LifecycleConfiguration *BucketLifecycleConfiguration `pulumi:"lifecycleConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig
	LoggingConfiguration *BucketLoggingConfiguration `pulumi:"loggingConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-metricsconfigurations
	MetricsConfigurations []BucketMetricsConfiguration `pulumi:"metricsConfigurations"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification
	NotificationConfiguration *BucketNotificationConfiguration `pulumi:"notificationConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockconfiguration
	ObjectLockConfiguration *BucketObjectLockConfiguration `pulumi:"objectLockConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockenabled
	ObjectLockEnabled *bool `pulumi:"objectLockEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-ownershipcontrols
	OwnershipControls *BucketOwnershipControls `pulumi:"ownershipControls"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-publicaccessblockconfiguration
	PublicAccessBlockConfiguration *BucketPublicAccessBlockConfiguration `pulumi:"publicAccessBlockConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration
	ReplicationConfiguration *BucketReplicationConfiguration `pulumi:"replicationConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning
	VersioningConfiguration *BucketVersioningConfiguration `pulumi:"versioningConfiguration"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration
	WebsiteConfiguration *BucketWebsiteConfiguration `pulumi:"websiteConfiguration"`
}

// The set of arguments for constructing a Bucket resource.
type BucketArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accelerateconfiguration
	AccelerateConfiguration BucketAccelerateConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-accesscontrol
	AccessControl pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-analyticsconfigurations
	AnalyticsConfigurations BucketAnalyticsConfigurationArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-bucketencryption
	BucketEncryption BucketBucketEncryptionPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-name
	BucketName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-crossoriginconfig
	CorsConfiguration BucketCorsConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-intelligenttieringconfigurations
	IntelligentTieringConfigurations BucketIntelligentTieringConfigurationArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-inventoryconfigurations
	InventoryConfigurations BucketInventoryConfigurationArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-lifecycleconfig
	LifecycleConfiguration BucketLifecycleConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-loggingconfig
	LoggingConfiguration BucketLoggingConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-metricsconfigurations
	MetricsConfigurations BucketMetricsConfigurationArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-notification
	NotificationConfiguration BucketNotificationConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockconfiguration
	ObjectLockConfiguration BucketObjectLockConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-objectlockenabled
	ObjectLockEnabled pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-ownershipcontrols
	OwnershipControls BucketOwnershipControlsPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-publicaccessblockconfiguration
	PublicAccessBlockConfiguration BucketPublicAccessBlockConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-replicationconfiguration
	ReplicationConfiguration BucketReplicationConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-versioning
	VersioningConfiguration BucketVersioningConfigurationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html#cfn-s3-bucket-websiteconfiguration
	WebsiteConfiguration BucketWebsiteConfigurationPtrInput
}

func (BucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketArgs)(nil)).Elem()
}

type BucketInput interface {
	pulumi.Input

	ToBucketOutput() BucketOutput
	ToBucketOutputWithContext(ctx context.Context) BucketOutput
}

func (*Bucket) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (i *Bucket) ToBucketOutput() BucketOutput {
	return i.ToBucketOutputWithContext(context.Background())
}

func (i *Bucket) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketOutput)
}

type BucketOutput struct{ *pulumi.OutputState }

func (BucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Bucket)(nil))
}

func (o BucketOutput) ToBucketOutput() BucketOutput {
	return o
}

func (o BucketOutput) ToBucketOutputWithContext(ctx context.Context) BucketOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketOutput{})
}
