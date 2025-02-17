// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::VPNGateway
//
// Deprecated: VPNGateway is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type VPNGateway struct {
	pulumi.CustomResourceState

	AmazonSideAsn pulumi.IntPtrOutput      `pulumi:"amazonSideAsn"`
	Tags          VPNGatewayTagArrayOutput `pulumi:"tags"`
	Type          pulumi.StringOutput      `pulumi:"type"`
}

// NewVPNGateway registers a new resource with the given unique name, arguments, and options.
func NewVPNGateway(ctx *pulumi.Context,
	name string, args *VPNGatewayArgs, opts ...pulumi.ResourceOption) (*VPNGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource VPNGateway
	err := ctx.RegisterResource("aws-native:ec2:VPNGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPNGateway gets an existing VPNGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPNGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPNGatewayState, opts ...pulumi.ResourceOption) (*VPNGateway, error) {
	var resource VPNGateway
	err := ctx.ReadResource("aws-native:ec2:VPNGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPNGateway resources.
type vpngatewayState struct {
}

type VPNGatewayState struct {
}

func (VPNGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpngatewayState)(nil)).Elem()
}

type vpngatewayArgs struct {
	AmazonSideAsn *int            `pulumi:"amazonSideAsn"`
	Tags          []VPNGatewayTag `pulumi:"tags"`
	Type          string          `pulumi:"type"`
}

// The set of arguments for constructing a VPNGateway resource.
type VPNGatewayArgs struct {
	AmazonSideAsn pulumi.IntPtrInput
	Tags          VPNGatewayTagArrayInput
	Type          pulumi.StringInput
}

func (VPNGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpngatewayArgs)(nil)).Elem()
}

type VPNGatewayInput interface {
	pulumi.Input

	ToVPNGatewayOutput() VPNGatewayOutput
	ToVPNGatewayOutputWithContext(ctx context.Context) VPNGatewayOutput
}

func (*VPNGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*VPNGateway)(nil))
}

func (i *VPNGateway) ToVPNGatewayOutput() VPNGatewayOutput {
	return i.ToVPNGatewayOutputWithContext(context.Background())
}

func (i *VPNGateway) ToVPNGatewayOutputWithContext(ctx context.Context) VPNGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPNGatewayOutput)
}

type VPNGatewayOutput struct{ *pulumi.OutputState }

func (VPNGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VPNGateway)(nil))
}

func (o VPNGatewayOutput) ToVPNGatewayOutput() VPNGatewayOutput {
	return o
}

func (o VPNGatewayOutput) ToVPNGatewayOutputWithContext(ctx context.Context) VPNGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VPNGatewayInput)(nil)).Elem(), &VPNGateway{})
	pulumi.RegisterOutputType(VPNGatewayOutput{})
}
