// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::EC2::TransitGatewayPeeringAttachment type
type TransitGatewayPeeringAttachment struct {
	pulumi.CustomResourceState

	// The time the transit gateway peering attachment was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Options for transit gateway peering attachment
	Options TransitGatewayPeeringAttachmentOptionsPtrOutput `pulumi:"options"`
	// The ID of the peer account
	PeerAccountId pulumi.StringOutput `pulumi:"peerAccountId"`
	// Peer Region
	PeerRegion pulumi.StringOutput `pulumi:"peerRegion"`
	// The ID of the peer transit gateway.
	PeerTransitGatewayId pulumi.StringOutput `pulumi:"peerTransitGatewayId"`
	// The state of the transit gateway peering attachment. Note that the initiating state has been deprecated.
	State pulumi.StringOutput `pulumi:"state"`
	// The status of the transit gateway peering attachment.
	Status TransitGatewayPeeringAttachmentPeeringAttachmentStatusOutput `pulumi:"status"`
	// The tags for the transit gateway peering attachment.
	Tags TransitGatewayPeeringAttachmentTagArrayOutput `pulumi:"tags"`
	// The ID of the transit gateway peering attachment.
	TransitGatewayAttachmentId pulumi.StringOutput `pulumi:"transitGatewayAttachmentId"`
	// The ID of the transit gateway.
	TransitGatewayId pulumi.StringOutput `pulumi:"transitGatewayId"`
}

// NewTransitGatewayPeeringAttachment registers a new resource with the given unique name, arguments, and options.
func NewTransitGatewayPeeringAttachment(ctx *pulumi.Context,
	name string, args *TransitGatewayPeeringAttachmentArgs, opts ...pulumi.ResourceOption) (*TransitGatewayPeeringAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeerAccountId == nil {
		return nil, errors.New("invalid value for required argument 'PeerAccountId'")
	}
	if args.PeerRegion == nil {
		return nil, errors.New("invalid value for required argument 'PeerRegion'")
	}
	if args.PeerTransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'PeerTransitGatewayId'")
	}
	if args.TransitGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'TransitGatewayId'")
	}
	var resource TransitGatewayPeeringAttachment
	err := ctx.RegisterResource("aws-native:ec2:TransitGatewayPeeringAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransitGatewayPeeringAttachment gets an existing TransitGatewayPeeringAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransitGatewayPeeringAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransitGatewayPeeringAttachmentState, opts ...pulumi.ResourceOption) (*TransitGatewayPeeringAttachment, error) {
	var resource TransitGatewayPeeringAttachment
	err := ctx.ReadResource("aws-native:ec2:TransitGatewayPeeringAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransitGatewayPeeringAttachment resources.
type transitGatewayPeeringAttachmentState struct {
}

type TransitGatewayPeeringAttachmentState struct {
}

func (TransitGatewayPeeringAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayPeeringAttachmentState)(nil)).Elem()
}

type transitGatewayPeeringAttachmentArgs struct {
	// Options for transit gateway peering attachment
	Options *TransitGatewayPeeringAttachmentOptions `pulumi:"options"`
	// The ID of the peer account
	PeerAccountId string `pulumi:"peerAccountId"`
	// Peer Region
	PeerRegion string `pulumi:"peerRegion"`
	// The ID of the peer transit gateway.
	PeerTransitGatewayId string `pulumi:"peerTransitGatewayId"`
	// The tags for the transit gateway peering attachment.
	Tags []TransitGatewayPeeringAttachmentTag `pulumi:"tags"`
	// The ID of the transit gateway.
	TransitGatewayId string `pulumi:"transitGatewayId"`
}

// The set of arguments for constructing a TransitGatewayPeeringAttachment resource.
type TransitGatewayPeeringAttachmentArgs struct {
	// Options for transit gateway peering attachment
	Options TransitGatewayPeeringAttachmentOptionsPtrInput
	// The ID of the peer account
	PeerAccountId pulumi.StringInput
	// Peer Region
	PeerRegion pulumi.StringInput
	// The ID of the peer transit gateway.
	PeerTransitGatewayId pulumi.StringInput
	// The tags for the transit gateway peering attachment.
	Tags TransitGatewayPeeringAttachmentTagArrayInput
	// The ID of the transit gateway.
	TransitGatewayId pulumi.StringInput
}

func (TransitGatewayPeeringAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transitGatewayPeeringAttachmentArgs)(nil)).Elem()
}

type TransitGatewayPeeringAttachmentInput interface {
	pulumi.Input

	ToTransitGatewayPeeringAttachmentOutput() TransitGatewayPeeringAttachmentOutput
	ToTransitGatewayPeeringAttachmentOutputWithContext(ctx context.Context) TransitGatewayPeeringAttachmentOutput
}

func (*TransitGatewayPeeringAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayPeeringAttachment)(nil))
}

func (i *TransitGatewayPeeringAttachment) ToTransitGatewayPeeringAttachmentOutput() TransitGatewayPeeringAttachmentOutput {
	return i.ToTransitGatewayPeeringAttachmentOutputWithContext(context.Background())
}

func (i *TransitGatewayPeeringAttachment) ToTransitGatewayPeeringAttachmentOutputWithContext(ctx context.Context) TransitGatewayPeeringAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransitGatewayPeeringAttachmentOutput)
}

type TransitGatewayPeeringAttachmentOutput struct{ *pulumi.OutputState }

func (TransitGatewayPeeringAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TransitGatewayPeeringAttachment)(nil))
}

func (o TransitGatewayPeeringAttachmentOutput) ToTransitGatewayPeeringAttachmentOutput() TransitGatewayPeeringAttachmentOutput {
	return o
}

func (o TransitGatewayPeeringAttachmentOutput) ToTransitGatewayPeeringAttachmentOutputWithContext(ctx context.Context) TransitGatewayPeeringAttachmentOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransitGatewayPeeringAttachmentInput)(nil)).Elem(), &TransitGatewayPeeringAttachment{})
	pulumi.RegisterOutputType(TransitGatewayPeeringAttachmentOutput{})
}
