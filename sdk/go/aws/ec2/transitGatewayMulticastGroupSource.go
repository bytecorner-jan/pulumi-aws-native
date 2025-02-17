// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::EC2::TransitGatewayMulticastGroupSource registers and deregisters members and sources (network interfaces) with the transit gateway multicast group
type TransitGatewayMulticastGroupSource struct {
	pulumi.CustomResourceState

	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress pulumi.StringOutput `pulumi:"groupIpAddress"`
	// Indicates that the resource is a transit gateway multicast group member.
	GroupMember pulumi.BoolOutput `pulumi:"groupMember"`
	// Indicates that the resource is a transit gateway multicast group member.
	GroupSource pulumi.BoolOutput `pulumi:"groupSource"`
	// The member type (for example, static).
	MemberType pulumi.StringOutput `pulumi:"memberType"`
	// The ID of the transit gateway attachment.
	NetworkInterfaceId pulumi.StringOutput `pulumi:"networkInterfaceId"`
	// The ID of the resource.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// The type of resource, for example a VPC attachment.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// The source type.
	SourceType pulumi.StringOutput `pulumi:"sourceType"`
	// The ID of the subnet.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId pulumi.StringOutput `pulumi:"transitGatewayMulticastDomainId"`
}

// NewTransitGatewayMulticastGroupSource registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayMulticastGroupSource(ctx *pulumi.Context,
	name string, args *TransitGatewayMulticastGroupSourceArgs, opts ...pulumi.ResourceOption) (*TransitGatewayMulticastGroupSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupIpAddress == nil {
		return nil, errors.New("invalid value for required argument 'GroupIpAddress'")
	}
	if args.NetworkInterfaceId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkInterfaceId'")
	}
	if args.TransitGatewayMulticastDomainId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayMulticastDomainId'")
	}
	var resource TransitGatewayMulticastGroupSource
	err := ctx.RegisterResource("aws-native:ec2:TransitGatewayMulticastGroupSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayMulticastGroupSource gets an existing TransitGatewayMulticastGroupSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayMulticastGroupSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayMulticastGroupSourceState, opts ...pulumi.ResourceOption) (*TransitGatewayMulticastGroupSource, error) {
	var resource TransitGatewayMulticastGroupSource
	err := ctx.ReadResource("aws-native:ec2:TransitGatewayMulticastGroupSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayMulticastGroupSource resources.
type transitGatewayMulticastGroupSourceState struct {
}

type TransitGatewayMulticastGroupSourceState struct {
}

func (TransitGatewayMulticastGroupSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayMulticastGroupSourceState)(nil)).Elem()
}

type transitGatewayMulticastGroupSourceArgs struct {
	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress string `pulumi:"groupIpAddress"`
	// The ID of the transit gateway attachment.
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId string `pulumi:"transitGatewayMulticastDomainId"`
}

// The set of arguments for constructing a TransitGatewayMulticastGroupSource resource.
type TransitGatewayMulticastGroupSourceArgs struct {
	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress pulumi.StringInput
	// The ID of the transit gateway attachment.
	NetworkInterfaceId pulumi.StringInput
	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId pulumi.StringInput
}

func (TransitGatewayMulticastGroupSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayMulticastGroupSourceArgs)(nil)).Elem()
}

type TransitGatewayMulticastGroupSourceInput interface {
	pulumi.Input

	ToTransitGatewayMulticastGroupSourceOutput() TransitGatewayMulticastGroupSourceOutput
	ToTransitGatewayMulticastGroupSourceOutputWithContext(ctx context.Context) TransitGatewayMulticastGroupSourceOutput
}

func (*TransitGatewayMulticastGroupSource) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayMulticastGroupSource)(nil))
}

func (i *TransitGatewayMulticastGroupSource) ToTransitGatewayMulticastGroupSourceOutput() TransitGatewayMulticastGroupSourceOutput {
	return i.ToTransitGatewayMulticastGroupSourceOutputWithContext(context.Background())
}

func (i *TransitGatewayMulticastGroupSource) ToTransitGatewayMulticastGroupSourceOutputWithContext(ctx context.Context) TransitGatewayMulticastGroupSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayMulticastGroupSourceOutput)
}

type TransitGatewayMulticastGroupSourceOutput struct{ *pulumi.OutputState }

func (TransitGatewayMulticastGroupSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayMulticastGroupSource)(nil))
}

func (o TransitGatewayMulticastGroupSourceOutput) ToTransitGatewayMulticastGroupSourceOutput() TransitGatewayMulticastGroupSourceOutput {
	return o
}

func (o TransitGatewayMulticastGroupSourceOutput) ToTransitGatewayMulticastGroupSourceOutputWithContext(ctx context.Context) TransitGatewayMulticastGroupSourceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayMulticastGroupSourceInput)(nil)).Elem(), &TransitGatewayMulticastGroupSource{})
	pulumi.RegisterOutputType(TransitGatewayMulticastGroupSourceOutput{})
}
