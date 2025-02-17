// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::Host
type Host struct {
	pulumi.CustomResourceState

	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement HostAutoPlacementPtrOutput `pulumi:"autoPlacement"`
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Id of the host created.
	HostId pulumi.StringOutput `pulumi:"hostId"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery HostRecoveryPtrOutput `pulumi:"hostRecovery"`
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
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
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement *HostAutoPlacement `pulumi:"autoPlacement"`
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery *HostRecovery `pulumi:"hostRecovery"`
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
	InstanceType string `pulumi:"instanceType"`
}

// The set of arguments for constructing a Host resource.
type HostArgs struct {
	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID.
	AutoPlacement HostAutoPlacementPtrInput
	// The Availability Zone in which to allocate the Dedicated Host.
	AvailabilityZone pulumi.StringInput
	// Indicates whether to enable or disable host recovery for the Dedicated Host. Host recovery is disabled by default.
	HostRecovery HostRecoveryPtrInput
	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.
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
	pulumi.RegisterInputType(reflect.TypeOf((*HostInput)(nil)).Elem(), &Host{})
	pulumi.RegisterOutputType(HostOutput{})
}
