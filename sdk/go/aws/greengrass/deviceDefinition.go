// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Greengrass::DeviceDefinition
//
// Deprecated: DeviceDefinition is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DeviceDefinition struct {
	pulumi.CustomResourceState

	Arn              pulumi.StringOutput                  `pulumi:"arn"`
	InitialVersion   DeviceDefinitionVersionTypePtrOutput `pulumi:"initialVersion"`
	LatestVersionArn pulumi.StringOutput                  `pulumi:"latestVersionArn"`
	Name             pulumi.StringOutput                  `pulumi:"name"`
	Tags             pulumi.AnyOutput                     `pulumi:"tags"`
}

// NewDeviceDefinition registers a new resource with the given unique name, arguments, and options.
func NewDeviceDefinition(ctx *pulumi.Context,
	name string, args *DeviceDefinitionArgs, opts ...pulumi.ResourceOption) (*DeviceDefinition, error) {
	if args == nil {
		args = &DeviceDefinitionArgs{}
	}

	var resource DeviceDefinition
	err := ctx.RegisterResource("aws-native:greengrass:DeviceDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceDefinition gets an existing DeviceDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceDefinitionState, opts ...pulumi.ResourceOption) (*DeviceDefinition, error) {
	var resource DeviceDefinition
	err := ctx.ReadResource("aws-native:greengrass:DeviceDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceDefinition resources.
type deviceDefinitionState struct {
}

type DeviceDefinitionState struct {
}

func (DeviceDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceDefinitionState)(nil)).Elem()
}

type deviceDefinitionArgs struct {
	InitialVersion *DeviceDefinitionVersionType `pulumi:"initialVersion"`
	Name           *string                      `pulumi:"name"`
	Tags           interface{}                  `pulumi:"tags"`
}

// The set of arguments for constructing a DeviceDefinition resource.
type DeviceDefinitionArgs struct {
	InitialVersion DeviceDefinitionVersionTypePtrInput
	Name           pulumi.StringPtrInput
	Tags           pulumi.Input
}

func (DeviceDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceDefinitionArgs)(nil)).Elem()
}

type DeviceDefinitionInput interface {
	pulumi.Input

	ToDeviceDefinitionOutput() DeviceDefinitionOutput
	ToDeviceDefinitionOutputWithContext(ctx context.Context) DeviceDefinitionOutput
}

func (*DeviceDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceDefinition)(nil))
}

func (i *DeviceDefinition) ToDeviceDefinitionOutput() DeviceDefinitionOutput {
	return i.ToDeviceDefinitionOutputWithContext(context.Background())
}

func (i *DeviceDefinition) ToDeviceDefinitionOutputWithContext(ctx context.Context) DeviceDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceDefinitionOutput)
}

type DeviceDefinitionOutput struct{ *pulumi.OutputState }

func (DeviceDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeviceDefinition)(nil))
}

func (o DeviceDefinitionOutput) ToDeviceDefinitionOutput() DeviceDefinitionOutput {
	return o
}

func (o DeviceDefinitionOutput) ToDeviceDefinitionOutputWithContext(ctx context.Context) DeviceDefinitionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceDefinitionInput)(nil)).Elem(), &DeviceDefinition{})
	pulumi.RegisterOutputType(DeviceDefinitionOutput{})
}
