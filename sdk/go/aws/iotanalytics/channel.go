// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoTAnalytics::Channel
//
// Deprecated: Channel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Channel struct {
	pulumi.CustomResourceState

	ChannelName     pulumi.StringOutput             `pulumi:"channelName"`
	ChannelStorage  ChannelStoragePtrOutput         `pulumi:"channelStorage"`
	RetentionPeriod ChannelRetentionPeriodPtrOutput `pulumi:"retentionPeriod"`
	Tags            ChannelTagArrayOutput           `pulumi:"tags"`
}

// NewChannel registers a new resource with the given unique name, arguments, and options.
func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		args = &ChannelArgs{}
	}

	var resource Channel
	err := ctx.RegisterResource("aws-native:iotanalytics:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannel gets an existing Channel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("aws-native:iotanalytics:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Channel resources.
type channelState struct {
}

type ChannelState struct {
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	ChannelName     *string                 `pulumi:"channelName"`
	ChannelStorage  *ChannelStorage         `pulumi:"channelStorage"`
	RetentionPeriod *ChannelRetentionPeriod `pulumi:"retentionPeriod"`
	Tags            []ChannelTag            `pulumi:"tags"`
}

// The set of arguments for constructing a Channel resource.
type ChannelArgs struct {
	ChannelName     pulumi.StringPtrInput
	ChannelStorage  ChannelStoragePtrInput
	RetentionPeriod ChannelRetentionPeriodPtrInput
	Tags            ChannelTagArrayInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((*Channel)(nil))
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Channel)(nil))
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelInput)(nil)).Elem(), &Channel{})
	pulumi.RegisterOutputType(ChannelOutput{})
}
