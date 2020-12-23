// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html
type ConfigRule struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ConfigRuleAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ConfigRulePropertiesOutput `pulumi:"properties"`
}

// NewConfigRule registers a new resource with the given unique name, arguments, and options.
func NewConfigRule(ctx *pulumi.Context,
	name string, args *ConfigRuleArgs, opts ...pulumi.ResourceOption) (*ConfigRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ConfigRule
	err := ctx.RegisterResource("cloudformation:Configuration:ConfigRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigRule gets an existing ConfigRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigRuleState, opts ...pulumi.ResourceOption) (*ConfigRule, error) {
	var resource ConfigRule
	err := ctx.ReadResource("cloudformation:Configuration:ConfigRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigRule resources.
type configRuleState struct {
	// The attributes associated with the resource
	Attributes *ConfigRuleAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ConfigRuleProperties `pulumi:"properties"`
}

type ConfigRuleState struct {
	// The attributes associated with the resource
	Attributes ConfigRuleAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ConfigRulePropertiesPtrInput
}

func (ConfigRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*configRuleState)(nil)).Elem()
}

type configRuleArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ConfigRuleProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ConfigRule resource.
type ConfigRuleArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ConfigRulePropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ConfigRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configRuleArgs)(nil)).Elem()
}

type ConfigRuleInput interface {
	pulumi.Input

	ToConfigRuleOutput() ConfigRuleOutput
	ToConfigRuleOutputWithContext(ctx context.Context) ConfigRuleOutput
}

func (*ConfigRule) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigRule)(nil))
}

func (i *ConfigRule) ToConfigRuleOutput() ConfigRuleOutput {
	return i.ToConfigRuleOutputWithContext(context.Background())
}

func (i *ConfigRule) ToConfigRuleOutputWithContext(ctx context.Context) ConfigRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigRuleOutput)
}

type ConfigRuleOutput struct {
	*pulumi.OutputState
}

func (ConfigRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigRule)(nil))
}

func (o ConfigRuleOutput) ToConfigRuleOutput() ConfigRuleOutput {
	return o
}

func (o ConfigRuleOutput) ToConfigRuleOutputWithContext(ctx context.Context) ConfigRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigRuleOutput{})
}
