// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html
type TransitGatewayConnect struct {
	pulumi.CustomResourceState

	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-options
	Options TransitGatewayConnectTransitGatewayConnectOptionsOutput `pulumi:"options"`
	State   pulumi.StringOutput                                     `pulumi:"state"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-tags
	Tags                       aws.TagArrayOutput  `pulumi:"tags"`
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`
	TransitGatewayId           pulumi.StringOutput `pulumi:"transitGatewayId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-transporttransitgatewayattachmentid
	TransportTransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transportTransitGatewayAttachmentId"`
}

// NewTransitGatewayConnect registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayConnect(ctx *pulumi.Context,
	name string, args *TransitGatewayConnectArgs, opts ...pulumi.ResourceOption) (*TransitGatewayConnect, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.TransportTransitGatewayAttachmentId == nil {
		return nil, errors.New("invalid value for required argument 'TransportTransitGatewayAttachmentId'")
	}
	var resource TransitGatewayConnect
	err := ctx.RegisterResource("aws-native:EC2:TransitGatewayConnect", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayConnect gets an existing TransitGatewayConnect resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayConnect(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayConnectState, opts ...pulumi.ResourceOption) (*TransitGatewayConnect, error) {
	var resource TransitGatewayConnect
	err := ctx.ReadResource("aws-native:EC2:TransitGatewayConnect", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayConnect resources.
type transitGatewayConnectState struct {
}

type TransitGatewayConnectState struct {
}

func (TransitGatewayConnectState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayConnectState)(nil)).Elem()
}

type transitGatewayConnectArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-options
	Options TransitGatewayConnectTransitGatewayConnectOptions `pulumi:"options"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-transporttransitgatewayattachmentid
	TransportTransitGatewayAttachmentId string `pulumi:"transportTransitGatewayAttachmentId"`
}

// The set of arguments for constructing a TransitGatewayConnect resource.
type TransitGatewayConnectArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-options
	Options TransitGatewayConnectTransitGatewayConnectOptionsInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnect.html#cfn-ec2-transitgatewayconnect-transporttransitgatewayattachmentid
	TransportTransitGatewayAttachmentId pulumi.StringInput
}

func (TransitGatewayConnectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayConnectArgs)(nil)).Elem()
}

type TransitGatewayConnectInput interface {
	pulumi.Input

	ToTransitGatewayConnectOutput() TransitGatewayConnectOutput
	ToTransitGatewayConnectOutputWithContext(ctx context.Context) TransitGatewayConnectOutput
}

func (*TransitGatewayConnect) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayConnect)(nil))
}

func (i *TransitGatewayConnect) ToTransitGatewayConnectOutput() TransitGatewayConnectOutput {
	return i.ToTransitGatewayConnectOutputWithContext(context.Background())
}

func (i *TransitGatewayConnect) ToTransitGatewayConnectOutputWithContext(ctx context.Context) TransitGatewayConnectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayConnectOutput)
}

type TransitGatewayConnectOutput struct{ *pulumi.OutputState }

func (TransitGatewayConnectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayConnect)(nil))
}

func (o TransitGatewayConnectOutput) ToTransitGatewayConnectOutput() TransitGatewayConnectOutput {
	return o
}

func (o TransitGatewayConnectOutput) ToTransitGatewayConnectOutputWithContext(ctx context.Context) TransitGatewayConnectOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TransitGatewayConnectOutput{})
}
