// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-imagebuilder-imagepipeline.html
type ImagePipeline struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ImagePipelineAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ImagePipelinePropertiesOutput `pulumi:"properties"`
}

// NewImagePipeline registers a new resource with the given unique name, arguments, and options.
func NewImagePipeline(ctx *pulumi.Context,
	name string, args *ImagePipelineArgs, opts ...pulumi.ResourceOption) (*ImagePipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ImagePipeline
	err := ctx.RegisterResource("cloudformation:ImageBuilder:ImagePipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImagePipeline gets an existing ImagePipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImagePipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImagePipelineState, opts ...pulumi.ResourceOption) (*ImagePipeline, error) {
	var resource ImagePipeline
	err := ctx.ReadResource("cloudformation:ImageBuilder:ImagePipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImagePipeline resources.
type imagePipelineState struct {
	// The attributes associated with the resource
	Attributes *ImagePipelineAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ImagePipelineProperties `pulumi:"properties"`
}

type ImagePipelineState struct {
	// The attributes associated with the resource
	Attributes ImagePipelineAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ImagePipelinePropertiesPtrInput
}

func (ImagePipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*imagePipelineState)(nil)).Elem()
}

type imagePipelineArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ImagePipelineProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ImagePipeline resource.
type ImagePipelineArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ImagePipelinePropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ImagePipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imagePipelineArgs)(nil)).Elem()
}

type ImagePipelineInput interface {
	pulumi.Input

	ToImagePipelineOutput() ImagePipelineOutput
	ToImagePipelineOutputWithContext(ctx context.Context) ImagePipelineOutput
}

func (*ImagePipeline) ElementType() reflect.Type {
	return reflect.TypeOf((*ImagePipeline)(nil))
}

func (i *ImagePipeline) ToImagePipelineOutput() ImagePipelineOutput {
	return i.ToImagePipelineOutputWithContext(context.Background())
}

func (i *ImagePipeline) ToImagePipelineOutputWithContext(ctx context.Context) ImagePipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePipelineOutput)
}

type ImagePipelineOutput struct {
	*pulumi.OutputState
}

func (ImagePipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImagePipeline)(nil))
}

func (o ImagePipelineOutput) ToImagePipelineOutput() ImagePipelineOutput {
	return o
}

func (o ImagePipelineOutput) ToImagePipelineOutputWithContext(ctx context.Context) ImagePipelineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ImagePipelineOutput{})
}
