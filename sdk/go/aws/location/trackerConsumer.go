// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html
type TrackerConsumer struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-consumerarn
	ConsumerArn pulumi.StringOutput `pulumi:"consumerArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-trackername
	TrackerName pulumi.StringOutput `pulumi:"trackerName"`
}

// NewTrackerConsumer registers a new resource with the given unique name, arguments, and options.
func NewTrackerConsumer(ctx *pulumi.Context,
	name string, args *TrackerConsumerArgs, opts ...pulumi.ResourceOption) (*TrackerConsumer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerArn == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerArn'")
	}
	if args.TrackerName == nil {
		return nil, errors.New("invalid value for required argument 'TrackerName'")
	}
	var resource TrackerConsumer
	err := ctx.RegisterResource("aws-native:Location:TrackerConsumer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrackerConsumer gets an existing TrackerConsumer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrackerConsumer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrackerConsumerState, opts ...pulumi.ResourceOption) (*TrackerConsumer, error) {
	var resource TrackerConsumer
	err := ctx.ReadResource("aws-native:Location:TrackerConsumer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TrackerConsumer resources.
type trackerConsumerState struct {
}

type TrackerConsumerState struct {
}

func (TrackerConsumerState) ElementType() reflect.Type {
	return reflect.TypeOf((*trackerConsumerState)(nil)).Elem()
}

type trackerConsumerArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-consumerarn
	ConsumerArn string `pulumi:"consumerArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-trackername
	TrackerName string `pulumi:"trackerName"`
}

// The set of arguments for constructing a TrackerConsumer resource.
type TrackerConsumerArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-consumerarn
	ConsumerArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-location-trackerconsumer.html#cfn-location-trackerconsumer-trackername
	TrackerName pulumi.StringInput
}

func (TrackerConsumerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trackerConsumerArgs)(nil)).Elem()
}

type TrackerConsumerInput interface {
	pulumi.Input

	ToTrackerConsumerOutput() TrackerConsumerOutput
	ToTrackerConsumerOutputWithContext(ctx context.Context) TrackerConsumerOutput
}

func (*TrackerConsumer) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackerConsumer)(nil))
}

func (i *TrackerConsumer) ToTrackerConsumerOutput() TrackerConsumerOutput {
	return i.ToTrackerConsumerOutputWithContext(context.Background())
}

func (i *TrackerConsumer) ToTrackerConsumerOutputWithContext(ctx context.Context) TrackerConsumerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrackerConsumerOutput)
}

type TrackerConsumerOutput struct{ *pulumi.OutputState }

func (TrackerConsumerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrackerConsumer)(nil))
}

func (o TrackerConsumerOutput) ToTrackerConsumerOutput() TrackerConsumerOutput {
	return o
}

func (o TrackerConsumerOutput) ToTrackerConsumerOutputWithContext(ctx context.Context) TrackerConsumerOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TrackerConsumerOutput{})
}
