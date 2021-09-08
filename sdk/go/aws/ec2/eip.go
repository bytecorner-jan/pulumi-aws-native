// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
type EIP struct {
	pulumi.CustomResourceState

	AllocationId pulumi.StringOutput `pulumi:"allocationId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-domain
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-instanceid
	InstanceId pulumi.StringPtrOutput `pulumi:"instanceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-publicipv4pool
	PublicIpv4Pool pulumi.StringPtrOutput `pulumi:"publicIpv4Pool"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewEIP registers a new resource with the given unique name, arguments, and options.
func NewEIP(ctx *pulumi.Context,
	name string, args *EIPArgs, opts ...pulumi.ResourceOption) (*EIP, error) {
	if args == nil {
		args = &EIPArgs{}
	}

	var resource EIP
	err := ctx.RegisterResource("aws-native:ec2:EIP", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEIP gets an existing EIP resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEIP(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EIPState, opts ...pulumi.ResourceOption) (*EIP, error) {
	var resource EIP
	err := ctx.ReadResource("aws-native:ec2:EIP", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EIP resources.
type eipState struct {
}

type EIPState struct {
}

func (EIPState) ElementType() reflect.Type {
	return reflect.TypeOf((*eipState)(nil)).Elem()
}

type eipArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-domain
	Domain *string `pulumi:"domain"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-instanceid
	InstanceId *string `pulumi:"instanceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-publicipv4pool
	PublicIpv4Pool *string `pulumi:"publicIpv4Pool"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a EIP resource.
type EIPArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-domain
	Domain pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-instanceid
	InstanceId pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-publicipv4pool
	PublicIpv4Pool pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-tags
	Tags aws.TagArrayInput
}

func (EIPArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eipArgs)(nil)).Elem()
}

type EIPInput interface {
	pulumi.Input

	ToEIPOutput() EIPOutput
	ToEIPOutputWithContext(ctx context.Context) EIPOutput
}

func (*EIP) ElementType() reflect.Type {
	return reflect.TypeOf((*EIP)(nil))
}

func (i *EIP) ToEIPOutput() EIPOutput {
	return i.ToEIPOutputWithContext(context.Background())
}

func (i *EIP) ToEIPOutputWithContext(ctx context.Context) EIPOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EIPOutput)
}

type EIPOutput struct{ *pulumi.OutputState }

func (EIPOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EIP)(nil))
}

func (o EIPOutput) ToEIPOutput() EIPOutput {
	return o
}

func (o EIPOutput) ToEIPOutputWithContext(ctx context.Context) EIPOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EIPOutput{})
}
