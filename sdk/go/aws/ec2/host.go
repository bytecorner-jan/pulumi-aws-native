// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html
type Host struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-autoplacement
	AutoPlacement pulumi.StringPtrOutput `pulumi:"autoPlacement"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-availabilityzone
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-hostrecovery
	HostRecovery pulumi.StringPtrOutput `pulumi:"hostRecovery"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancetype
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
}

// NewHost registers a new resource with the given unique name, arguments, and options.
func NewHost(ctx *pulumi.Context,
	name string, args *HostArgs, opts ...pulumi.ResourceOption) (*Host, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource Host
	err := ctx.RegisterResource("aws-native:ec2:Host", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHost gets an existing Host resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHost(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostState, opts ...pulumi.ResourceOption) (*Host, error) {
	var resource Host
	err := ctx.ReadResource("aws-native:ec2:Host", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Host resources.
type hostState struct {
}

type HostState struct {
}

func (HostState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostState)(nil)).Elem()
}

type hostArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-autoplacement
	AutoPlacement *string `pulumi:"autoPlacement"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-availabilityzone
	AvailabilityZone string `pulumi:"availabilityZone"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-hostrecovery
	HostRecovery *string `pulumi:"hostRecovery"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancetype
	InstanceType string `pulumi:"instanceType"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-autoplacement
	AutoPlacement pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-availabilityzone
	AvailabilityZone pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-hostrecovery
	HostRecovery pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancetype
	InstanceType pulumi.StringInput
}

func (HostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostArgs)(nil)).Elem()
}

type HostInput interface {
	pulumi.Input

	ToHostOutput() HostOutput
	ToHostOutputWithContext(ctx context.Context) HostOutput
}

func (*Host) ElementType() reflect.Type {
	return reflect.TypeOf((*Host)(nil))
}

func (i *Host) ToHostOutput() HostOutput {
	return i.ToHostOutputWithContext(context.Background())
}

func (i *Host) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostOutput)
}

type HostOutput struct{ *pulumi.OutputState }

func (HostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Host)(nil))
}

func (o HostOutput) ToHostOutput() HostOutput {
	return o
}

func (o HostOutput) ToHostOutputWithContext(ctx context.Context) HostOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HostOutput{})
}
