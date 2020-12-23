// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-devicedefinition.html
type DeviceDefinition struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes DeviceDefinitionAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DeviceDefinitionPropertiesOutput `pulumi:"properties"`
}

// NewDeviceDefinition registers a new resource with the given unique name, arguments, and options.
func NewDeviceDefinition(ctx *pulumi.Context,
	name string, args *DeviceDefinitionArgs, opts ...pulumi.ResourceOption) (*DeviceDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource DeviceDefinition
	err := ctx.RegisterResource("cloudformation:Greengrass:DeviceDefinition", name, args, &resource, opts...)
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
	err := ctx.ReadResource("cloudformation:Greengrass:DeviceDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceDefinition resources.
type deviceDefinitionState struct {
	// The attributes associated with the resource
	Attributes *DeviceDefinitionAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *DeviceDefinitionProperties `pulumi:"properties"`
}

type DeviceDefinitionState struct {
	// The attributes associated with the resource
	Attributes DeviceDefinitionAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DeviceDefinitionPropertiesPtrInput
}

func (DeviceDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceDefinitionState)(nil)).Elem()
}

type deviceDefinitionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DeviceDefinitionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a DeviceDefinition resource.
type DeviceDefinitionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DeviceDefinitionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
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

type DeviceDefinitionOutput struct {
	*pulumi.OutputState
}

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
	pulumi.RegisterOutputType(DeviceDefinitionOutput{})
}
