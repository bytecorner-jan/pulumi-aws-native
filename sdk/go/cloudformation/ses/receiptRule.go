// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ses-receiptrule.html
type ReceiptRule struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ReceiptRuleAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ReceiptRulePropertiesOutput `pulumi:"properties"`
}

// NewReceiptRule registers a new resource with the given unique name, arguments, and options.
func NewReceiptRule(ctx *pulumi.Context,
	name string, args *ReceiptRuleArgs, opts ...pulumi.ResourceOption) (*ReceiptRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ReceiptRule
	err := ctx.RegisterResource("cloudformation:SES:ReceiptRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReceiptRule gets an existing ReceiptRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReceiptRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReceiptRuleState, opts ...pulumi.ResourceOption) (*ReceiptRule, error) {
	var resource ReceiptRule
	err := ctx.ReadResource("cloudformation:SES:ReceiptRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReceiptRule resources.
type receiptRuleState struct {
	// The attributes associated with the resource
	Attributes *ReceiptRuleAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ReceiptRuleProperties `pulumi:"properties"`
}

type ReceiptRuleState struct {
	// The attributes associated with the resource
	Attributes ReceiptRuleAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ReceiptRulePropertiesPtrInput
}

func (ReceiptRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleState)(nil)).Elem()
}

type receiptRuleArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ReceiptRuleProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ReceiptRule resource.
type ReceiptRuleArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ReceiptRulePropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ReceiptRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*receiptRuleArgs)(nil)).Elem()
}

type ReceiptRuleInput interface {
	pulumi.Input

	ToReceiptRuleOutput() ReceiptRuleOutput
	ToReceiptRuleOutputWithContext(ctx context.Context) ReceiptRuleOutput
}

func (*ReceiptRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRule)(nil))
}

func (i *ReceiptRule) ToReceiptRuleOutput() ReceiptRuleOutput {
	return i.ToReceiptRuleOutputWithContext(context.Background())
}

func (i *ReceiptRule) ToReceiptRuleOutputWithContext(ctx context.Context) ReceiptRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReceiptRuleOutput)
}

type ReceiptRuleOutput struct {
	*pulumi.OutputState
}

func (ReceiptRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReceiptRule)(nil))
}

func (o ReceiptRuleOutput) ToReceiptRuleOutput() ReceiptRuleOutput {
	return o
}

func (o ReceiptRuleOutput) ToReceiptRuleOutputWithContext(ctx context.Context) ReceiptRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReceiptRuleOutput{})
}
