// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::DeviceFarm::NetworkProfile creates a new DF Network Profile
type NetworkProfile struct {
	pulumi.CustomResourceState

	Arn                   pulumi.StringOutput          `pulumi:"arn"`
	Description           pulumi.StringPtrOutput       `pulumi:"description"`
	DownlinkBandwidthBits pulumi.IntPtrOutput          `pulumi:"downlinkBandwidthBits"`
	DownlinkDelayMs       pulumi.IntPtrOutput          `pulumi:"downlinkDelayMs"`
	DownlinkJitterMs      pulumi.IntPtrOutput          `pulumi:"downlinkJitterMs"`
	DownlinkLossPercent   pulumi.IntPtrOutput          `pulumi:"downlinkLossPercent"`
	Name                  pulumi.StringOutput          `pulumi:"name"`
	ProjectArn            pulumi.StringOutput          `pulumi:"projectArn"`
	Tags                  NetworkProfileTagArrayOutput `pulumi:"tags"`
	UplinkBandwidthBits   pulumi.IntPtrOutput          `pulumi:"uplinkBandwidthBits"`
	UplinkDelayMs         pulumi.IntPtrOutput          `pulumi:"uplinkDelayMs"`
	UplinkJitterMs        pulumi.IntPtrOutput          `pulumi:"uplinkJitterMs"`
	UplinkLossPercent     pulumi.IntPtrOutput          `pulumi:"uplinkLossPercent"`
}

// NewNetworkProfile registers a new resource with the given unique name, arguments, and options.
func NewNetworkProfile(ctx *pulumi.Context,
	name string, args *NetworkProfileArgs, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ProjectArn == nil {
		return nil, errors.New("invalid value for required argument 'ProjectArn'")
	}
	var resource NetworkProfile
	err := ctx.RegisterResource("aws-native:devicefarm:NetworkProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkProfile gets an existing NetworkProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkProfileState, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	var resource NetworkProfile
	err := ctx.ReadResource("aws-native:devicefarm:NetworkProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkProfile resources.
type networkProfileState struct {
}

type NetworkProfileState struct {
}

func (NetworkProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileState)(nil)).Elem()
}

type networkProfileArgs struct {
	Description           *string             `pulumi:"description"`
	DownlinkBandwidthBits *int                `pulumi:"downlinkBandwidthBits"`
	DownlinkDelayMs       *int                `pulumi:"downlinkDelayMs"`
	DownlinkJitterMs      *int                `pulumi:"downlinkJitterMs"`
	DownlinkLossPercent   *int                `pulumi:"downlinkLossPercent"`
	Name                  string              `pulumi:"name"`
	ProjectArn            string              `pulumi:"projectArn"`
	Tags                  []NetworkProfileTag `pulumi:"tags"`
	UplinkBandwidthBits   *int                `pulumi:"uplinkBandwidthBits"`
	UplinkDelayMs         *int                `pulumi:"uplinkDelayMs"`
	UplinkJitterMs        *int                `pulumi:"uplinkJitterMs"`
	UplinkLossPercent     *int                `pulumi:"uplinkLossPercent"`
}

// The set of arguments for constructing a NetworkProfile resource.
type NetworkProfileArgs struct {
	Description           pulumi.StringPtrInput
	DownlinkBandwidthBits pulumi.IntPtrInput
	DownlinkDelayMs       pulumi.IntPtrInput
	DownlinkJitterMs      pulumi.IntPtrInput
	DownlinkLossPercent   pulumi.IntPtrInput
	Name                  pulumi.StringInput
	ProjectArn            pulumi.StringInput
	Tags                  NetworkProfileTagArrayInput
	UplinkBandwidthBits   pulumi.IntPtrInput
	UplinkDelayMs         pulumi.IntPtrInput
	UplinkJitterMs        pulumi.IntPtrInput
	UplinkLossPercent     pulumi.IntPtrInput
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileArgs)(nil)).Elem()
}

type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput
}

func (*NetworkProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil))
}

func (i *NetworkProfile) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i *NetworkProfile) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil))
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NetworkProfileOutput{})
}
