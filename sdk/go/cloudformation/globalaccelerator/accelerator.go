// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package globalaccelerator

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-globalaccelerator-accelerator.html
type Accelerator struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes AcceleratorAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties AcceleratorPropertiesOutput `pulumi:"properties"`
}

// NewAccelerator registers a new resource with the given unique name, arguments, and options.
func NewAccelerator(ctx *pulumi.Context,
	name string, args *AcceleratorArgs, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Accelerator
	err := ctx.RegisterResource("cloudformation:GlobalAccelerator:Accelerator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccelerator gets an existing Accelerator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccelerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AcceleratorState, opts ...pulumi.ResourceOption) (*Accelerator, error) {
	var resource Accelerator
	err := ctx.ReadResource("cloudformation:GlobalAccelerator:Accelerator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Accelerator resources.
type acceleratorState struct {
	// The attributes associated with the resource
	Attributes *AcceleratorAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *AcceleratorProperties `pulumi:"properties"`
}

type AcceleratorState struct {
	// The attributes associated with the resource
	Attributes AcceleratorAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties AcceleratorPropertiesPtrInput
}

func (AcceleratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorState)(nil)).Elem()
}

type acceleratorArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties AcceleratorProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Accelerator resource.
type AcceleratorArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties AcceleratorPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (AcceleratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*acceleratorArgs)(nil)).Elem()
}

type AcceleratorInput interface {
	pulumi.Input

	ToAcceleratorOutput() AcceleratorOutput
	ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput
}

func (*Accelerator) ElementType() reflect.Type {
	return reflect.TypeOf((*Accelerator)(nil))
}

func (i *Accelerator) ToAcceleratorOutput() AcceleratorOutput {
	return i.ToAcceleratorOutputWithContext(context.Background())
}

func (i *Accelerator) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AcceleratorOutput)
}

type AcceleratorOutput struct {
	*pulumi.OutputState
}

func (AcceleratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Accelerator)(nil))
}

func (o AcceleratorOutput) ToAcceleratorOutput() AcceleratorOutput {
	return o
}

func (o AcceleratorOutput) ToAcceleratorOutputWithContext(ctx context.Context) AcceleratorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AcceleratorOutput{})
}
