// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package greengrass

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrass-resourcedefinitionversion.html
type ResourceDefinitionVersion struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ResourceDefinitionVersionAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResourceDefinitionVersionPropertiesOutput `pulumi:"properties"`
}

// NewResourceDefinitionVersion registers a new resource with the given unique name, arguments, and options.
func NewResourceDefinitionVersion(ctx *pulumi.Context,
	name string, args *ResourceDefinitionVersionArgs, opts ...pulumi.ResourceOption) (*ResourceDefinitionVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ResourceDefinitionVersion
	err := ctx.RegisterResource("cloudformation:Greengrass:ResourceDefinitionVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResourceDefinitionVersion gets an existing ResourceDefinitionVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResourceDefinitionVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceDefinitionVersionState, opts ...pulumi.ResourceOption) (*ResourceDefinitionVersion, error) {
	var resource ResourceDefinitionVersion
	err := ctx.ReadResource("cloudformation:Greengrass:ResourceDefinitionVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResourceDefinitionVersion resources.
type resourceDefinitionVersionState struct {
	// The attributes associated with the resource
	Attributes *ResourceDefinitionVersionAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ResourceDefinitionVersionProperties `pulumi:"properties"`
}

type ResourceDefinitionVersionState struct {
	// The attributes associated with the resource
	Attributes ResourceDefinitionVersionAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResourceDefinitionVersionPropertiesPtrInput
}

func (ResourceDefinitionVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDefinitionVersionState)(nil)).Elem()
}

type resourceDefinitionVersionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResourceDefinitionVersionProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ResourceDefinitionVersion resource.
type ResourceDefinitionVersionArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResourceDefinitionVersionPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ResourceDefinitionVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceDefinitionVersionArgs)(nil)).Elem()
}

type ResourceDefinitionVersionInput interface {
	pulumi.Input

	ToResourceDefinitionVersionOutput() ResourceDefinitionVersionOutput
	ToResourceDefinitionVersionOutputWithContext(ctx context.Context) ResourceDefinitionVersionOutput
}

func (*ResourceDefinitionVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDefinitionVersion)(nil))
}

func (i *ResourceDefinitionVersion) ToResourceDefinitionVersionOutput() ResourceDefinitionVersionOutput {
	return i.ToResourceDefinitionVersionOutputWithContext(context.Background())
}

func (i *ResourceDefinitionVersion) ToResourceDefinitionVersionOutputWithContext(ctx context.Context) ResourceDefinitionVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceDefinitionVersionOutput)
}

type ResourceDefinitionVersionOutput struct {
	*pulumi.OutputState
}

func (ResourceDefinitionVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceDefinitionVersion)(nil))
}

func (o ResourceDefinitionVersionOutput) ToResourceDefinitionVersionOutput() ResourceDefinitionVersionOutput {
	return o
}

func (o ResourceDefinitionVersionOutput) ToResourceDefinitionVersionOutputWithContext(ctx context.Context) ResourceDefinitionVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceDefinitionVersionOutput{})
}
