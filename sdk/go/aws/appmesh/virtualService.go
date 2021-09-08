// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appmesh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html
type VirtualService struct {
	pulumi.CustomResourceState

	Arn           pulumi.StringOutput `pulumi:"arn"`
	MeshName      pulumi.StringOutput `pulumi:"meshName"`
	MeshOwner     pulumi.StringOutput `pulumi:"meshOwner"`
	ResourceOwner pulumi.StringOutput `pulumi:"resourceOwner"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-spec
	Spec VirtualServiceVirtualServiceSpecOutput `pulumi:"spec"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-tags
	Tags               aws.TagArrayOutput  `pulumi:"tags"`
	Uid                pulumi.StringOutput `pulumi:"uid"`
	VirtualServiceName pulumi.StringOutput `pulumi:"virtualServiceName"`
}

// NewVirtualService registers a new resource with the given unique name, arguments, and options.
func NewVirtualService(ctx *pulumi.Context,
	name string, args *VirtualServiceArgs, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MeshName == nil {
		return nil, errors.New("invalid value for required argument 'MeshName'")
	}
	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	if args.VirtualServiceName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualServiceName'")
	}
	var resource VirtualService
	err := ctx.RegisterResource("aws-native:appmesh:VirtualService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualService gets an existing VirtualService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualServiceState, opts ...pulumi.ResourceOption) (*VirtualService, error) {
	var resource VirtualService
	err := ctx.ReadResource("aws-native:appmesh:VirtualService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualService resources.
type virtualServiceState struct {
}

type VirtualServiceState struct {
}

func (VirtualServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceState)(nil)).Elem()
}

type virtualServiceArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-meshname
	MeshName string `pulumi:"meshName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-meshowner
	MeshOwner *string `pulumi:"meshOwner"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-spec
	Spec VirtualServiceVirtualServiceSpec `pulumi:"spec"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-virtualservicename
	VirtualServiceName string `pulumi:"virtualServiceName"`
}

// The set of arguments for constructing a VirtualService resource.
type VirtualServiceArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-meshname
	MeshName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-meshowner
	MeshOwner pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-spec
	Spec VirtualServiceVirtualServiceSpecInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appmesh-virtualservice.html#cfn-appmesh-virtualservice-virtualservicename
	VirtualServiceName pulumi.StringInput
}

func (VirtualServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualServiceArgs)(nil)).Elem()
}

type VirtualServiceInput interface {
	pulumi.Input

	ToVirtualServiceOutput() VirtualServiceOutput
	ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput
}

func (*VirtualService) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualService)(nil))
}

func (i *VirtualService) ToVirtualServiceOutput() VirtualServiceOutput {
	return i.ToVirtualServiceOutputWithContext(context.Background())
}

func (i *VirtualService) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualServiceOutput)
}

type VirtualServiceOutput struct{ *pulumi.OutputState }

func (VirtualServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualService)(nil))
}

func (o VirtualServiceOutput) ToVirtualServiceOutput() VirtualServiceOutput {
	return o
}

func (o VirtualServiceOutput) ToVirtualServiceOutputWithContext(ctx context.Context) VirtualServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualServiceOutput{})
}
