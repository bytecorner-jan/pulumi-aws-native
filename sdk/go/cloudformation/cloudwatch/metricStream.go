// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudwatch-metricstream.html
type MetricStream struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes MetricStreamAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties MetricStreamPropertiesOutput `pulumi:"properties"`
}

// NewMetricStream registers a new resource with the given unique name, arguments, and options.
func NewMetricStream(ctx *pulumi.Context,
	name string, args *MetricStreamArgs, opts ...pulumi.ResourceOption) (*MetricStream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource MetricStream
	err := ctx.RegisterResource("cloudformation:CloudWatch:MetricStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricStream gets an existing MetricStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricStreamState, opts ...pulumi.ResourceOption) (*MetricStream, error) {
	var resource MetricStream
	err := ctx.ReadResource("cloudformation:CloudWatch:MetricStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricStream resources.
type metricStreamState struct {
	// The attributes associated with the resource
	Attributes *MetricStreamAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *MetricStreamProperties `pulumi:"properties"`
}

type MetricStreamState struct {
	// The attributes associated with the resource
	Attributes MetricStreamAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties MetricStreamPropertiesPtrInput
}

func (MetricStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricStreamState)(nil)).Elem()
}

type metricStreamArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties MetricStreamProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a MetricStream resource.
type MetricStreamArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties MetricStreamPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (MetricStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricStreamArgs)(nil)).Elem()
}

type MetricStreamInput interface {
	pulumi.Input

	ToMetricStreamOutput() MetricStreamOutput
	ToMetricStreamOutputWithContext(ctx context.Context) MetricStreamOutput
}

func (*MetricStream) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricStream)(nil))
}

func (i *MetricStream) ToMetricStreamOutput() MetricStreamOutput {
	return i.ToMetricStreamOutputWithContext(context.Background())
}

func (i *MetricStream) ToMetricStreamOutputWithContext(ctx context.Context) MetricStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricStreamOutput)
}

type MetricStreamOutput struct {
	*pulumi.OutputState
}

func (MetricStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricStream)(nil))
}

func (o MetricStreamOutput) ToMetricStreamOutput() MetricStreamOutput {
	return o
}

func (o MetricStreamOutput) ToMetricStreamOutputWithContext(ctx context.Context) MetricStreamOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MetricStreamOutput{})
}
