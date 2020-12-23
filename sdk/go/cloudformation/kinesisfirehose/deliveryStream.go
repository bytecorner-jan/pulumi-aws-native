// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesisfirehose

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisfirehose-deliverystream.html
type DeliveryStream struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes DeliveryStreamAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DeliveryStreamPropertiesOutput `pulumi:"properties"`
}

// NewDeliveryStream registers a new resource with the given unique name, arguments, and options.
func NewDeliveryStream(ctx *pulumi.Context,
	name string, args *DeliveryStreamArgs, opts ...pulumi.ResourceOption) (*DeliveryStream, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource DeliveryStream
	err := ctx.RegisterResource("cloudformation:KinesisFirehose:DeliveryStream", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeliveryStream gets an existing DeliveryStream resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeliveryStream(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeliveryStreamState, opts ...pulumi.ResourceOption) (*DeliveryStream, error) {
	var resource DeliveryStream
	err := ctx.ReadResource("cloudformation:KinesisFirehose:DeliveryStream", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeliveryStream resources.
type deliveryStreamState struct {
	// The attributes associated with the resource
	Attributes *DeliveryStreamAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *DeliveryStreamProperties `pulumi:"properties"`
}

type DeliveryStreamState struct {
	// The attributes associated with the resource
	Attributes DeliveryStreamAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DeliveryStreamPropertiesPtrInput
}

func (DeliveryStreamState) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryStreamState)(nil)).Elem()
}

type deliveryStreamArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DeliveryStreamProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a DeliveryStream resource.
type DeliveryStreamArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DeliveryStreamPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (DeliveryStreamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deliveryStreamArgs)(nil)).Elem()
}

type DeliveryStreamInput interface {
	pulumi.Input

	ToDeliveryStreamOutput() DeliveryStreamOutput
	ToDeliveryStreamOutputWithContext(ctx context.Context) DeliveryStreamOutput
}

func (*DeliveryStream) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryStream)(nil))
}

func (i *DeliveryStream) ToDeliveryStreamOutput() DeliveryStreamOutput {
	return i.ToDeliveryStreamOutputWithContext(context.Background())
}

func (i *DeliveryStream) ToDeliveryStreamOutputWithContext(ctx context.Context) DeliveryStreamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeliveryStreamOutput)
}

type DeliveryStreamOutput struct {
	*pulumi.OutputState
}

func (DeliveryStreamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeliveryStream)(nil))
}

func (o DeliveryStreamOutput) ToDeliveryStreamOutput() DeliveryStreamOutput {
	return o
}

func (o DeliveryStreamOutput) ToDeliveryStreamOutputWithContext(ctx context.Context) DeliveryStreamOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeliveryStreamOutput{})
}
