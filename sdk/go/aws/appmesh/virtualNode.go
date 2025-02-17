// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppMesh::VirtualNode
//
// Deprecated: VirtualNode is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type VirtualNode struct {
	pulumi.CustomResourceState

	Arn             pulumi.StringOutput       `pulumi:"arn"`
	MeshName        pulumi.StringOutput       `pulumi:"meshName"`
	MeshOwner       pulumi.StringPtrOutput    `pulumi:"meshOwner"`
	ResourceOwner   pulumi.StringOutput       `pulumi:"resourceOwner"`
	Spec            VirtualNodeSpecOutput     `pulumi:"spec"`
	Tags            VirtualNodeTagArrayOutput `pulumi:"tags"`
	Uid             pulumi.StringOutput       `pulumi:"uid"`
	VirtualNodeName pulumi.StringPtrOutput    `pulumi:"virtualNodeName"`
}

// NewVirtualNode registers a new resource with the given unique name, arguments, and options.
func NewVirtualNode(ctx *pulumi.Context,
	name string, args *VirtualNodeArgs, opts ...pulumi.ResourceOption) (*VirtualNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	var resource VirtualNode
	err := ctx.RegisterResource("aws-native:appmesh:VirtualNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNode gets an existing VirtualNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNodeState, opts ...pulumi.ResourceOption) (*VirtualNode, error) {
	var resource VirtualNode
	err := ctx.ReadResource("aws-native:appmesh:VirtualNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNode resources.
type virtualNodeState struct {
}

type VirtualNodeState struct {
}

func (VirtualNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNodeState)(nil)).Elem()
}

type virtualNodeArgs struct {
	MeshName        string           `pulumi:"meshName"`
	MeshOwner       *string          `pulumi:"meshOwner"`
	Spec            VirtualNodeSpec  `pulumi:"spec"`
	Tags            []VirtualNodeTag `pulumi:"tags"`
	VirtualNodeName *string          `pulumi:"virtualNodeName"`
}

// The set of arguments for constructing a VirtualNode resource.
type VirtualNodeArgs struct {
	MeshName        pulumi.StringInput
	MeshOwner       pulumi.StringPtrInput
	Spec            VirtualNodeSpecInput
	Tags            VirtualNodeTagArrayInput
	VirtualNodeName pulumi.StringPtrInput
}

func (VirtualNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNodeArgs)(nil)).Elem()
}

type VirtualNodeInput interface {
	pulumi.Input

	ToVirtualNodeOutput() VirtualNodeOutput
	ToVirtualNodeOutputWithContext(ctx context.Context) VirtualNodeOutput
}

func (*VirtualNode) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNode)(nil))
}

func (i *VirtualNode) ToVirtualNodeOutput() VirtualNodeOutput {
	return i.ToVirtualNodeOutputWithContext(context.Background())
}

func (i *VirtualNode) ToVirtualNodeOutputWithContext(ctx context.Context) VirtualNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNodeOutput)
}

type VirtualNodeOutput struct{ *pulumi.OutputState }

func (VirtualNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNode)(nil))
}

func (o VirtualNodeOutput) ToVirtualNodeOutput() VirtualNodeOutput {
	return o
}

func (o VirtualNodeOutput) ToVirtualNodeOutputWithContext(ctx context.Context) VirtualNodeOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNodeInput)(nil)).Elem(), &VirtualNode{})
	pulumi.RegisterOutputType(VirtualNodeOutput{})
}
