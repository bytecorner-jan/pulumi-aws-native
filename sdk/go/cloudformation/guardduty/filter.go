// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package guardduty

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html
type Filter struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes FilterAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties FilterPropertiesOutput `pulumi:"properties"`
}

// NewFilter registers a new resource with the given unique name, arguments, and options.
func NewFilter(ctx *pulumi.Context,
	name string, args *FilterArgs, opts ...pulumi.ResourceOption) (*Filter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource Filter
	err := ctx.RegisterResource("cloudformation:GuardDuty:Filter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFilter gets an existing Filter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FilterState, opts ...pulumi.ResourceOption) (*Filter, error) {
	var resource Filter
	err := ctx.ReadResource("cloudformation:GuardDuty:Filter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Filter resources.
type filterState struct {
	// The attributes associated with the resource
	Attributes *FilterAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *FilterProperties `pulumi:"properties"`
}

type FilterState struct {
	// The attributes associated with the resource
	Attributes FilterAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties FilterPropertiesPtrInput
}

func (FilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*filterState)(nil)).Elem()
}

type filterArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties FilterProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a Filter resource.
type FilterArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties FilterPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (FilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*filterArgs)(nil)).Elem()
}

type FilterInput interface {
	pulumi.Input

	ToFilterOutput() FilterOutput
	ToFilterOutputWithContext(ctx context.Context) FilterOutput
}

func (*Filter) ElementType() reflect.Type {
	return reflect.TypeOf((*Filter)(nil))
}

func (i *Filter) ToFilterOutput() FilterOutput {
	return i.ToFilterOutputWithContext(context.Background())
}

func (i *Filter) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilterOutput)
}

type FilterOutput struct {
	*pulumi.OutputState
}

func (FilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Filter)(nil))
}

func (o FilterOutput) ToFilterOutput() FilterOutput {
	return o
}

func (o FilterOutput) ToFilterOutputWithContext(ctx context.Context) FilterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FilterOutput{})
}
