// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
type MetricFilter struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes MetricFilterAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties MetricFilterPropertiesOutput `pulumi:"properties"`
}

// NewMetricFilter registers a new resource with the given unique name, arguments, and options.
func NewMetricFilter(ctx *pulumi.Context,
	name string, args *MetricFilterArgs, opts ...pulumi.ResourceOption) (*MetricFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource MetricFilter
	err := ctx.RegisterResource("cloudformation:Logs:MetricFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricFilter gets an existing MetricFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricFilterState, opts ...pulumi.ResourceOption) (*MetricFilter, error) {
	var resource MetricFilter
	err := ctx.ReadResource("cloudformation:Logs:MetricFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricFilter resources.
type metricFilterState struct {
	// The attributes associated with the resource
	Attributes *MetricFilterAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *MetricFilterProperties `pulumi:"properties"`
}

type MetricFilterState struct {
	// The attributes associated with the resource
	Attributes MetricFilterAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties MetricFilterPropertiesPtrInput
}

func (MetricFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricFilterState)(nil)).Elem()
}

type metricFilterArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties MetricFilterProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a MetricFilter resource.
type MetricFilterArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties MetricFilterPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (MetricFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricFilterArgs)(nil)).Elem()
}

type MetricFilterInput interface {
	pulumi.Input

	ToMetricFilterOutput() MetricFilterOutput
	ToMetricFilterOutputWithContext(ctx context.Context) MetricFilterOutput
}

func (*MetricFilter) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricFilter)(nil))
}

func (i *MetricFilter) ToMetricFilterOutput() MetricFilterOutput {
	return i.ToMetricFilterOutputWithContext(context.Background())
}

func (i *MetricFilter) ToMetricFilterOutputWithContext(ctx context.Context) MetricFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricFilterOutput)
}

type MetricFilterOutput struct {
	*pulumi.OutputState
}

func (MetricFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricFilter)(nil))
}

func (o MetricFilterOutput) ToMetricFilterOutput() MetricFilterOutput {
	return o
}

func (o MetricFilterOutput) ToMetricFilterOutputWithContext(ctx context.Context) MetricFilterOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MetricFilterOutput{})
}
