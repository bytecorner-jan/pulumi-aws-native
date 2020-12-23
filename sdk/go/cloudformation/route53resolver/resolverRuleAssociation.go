// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverruleassociation.html
type ResolverRuleAssociation struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes ResolverRuleAssociationAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResolverRuleAssociationPropertiesOutput `pulumi:"properties"`
}

// NewResolverRuleAssociation registers a new resource with the given unique name, arguments, and options.
func NewResolverRuleAssociation(ctx *pulumi.Context,
	name string, args *ResolverRuleAssociationArgs, opts ...pulumi.ResourceOption) (*ResolverRuleAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource ResolverRuleAssociation
	err := ctx.RegisterResource("cloudformation:Route53Resolver:ResolverRuleAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResolverRuleAssociation gets an existing ResolverRuleAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResolverRuleAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResolverRuleAssociationState, opts ...pulumi.ResourceOption) (*ResolverRuleAssociation, error) {
	var resource ResolverRuleAssociation
	err := ctx.ReadResource("cloudformation:Route53Resolver:ResolverRuleAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResolverRuleAssociation resources.
type resolverRuleAssociationState struct {
	// The attributes associated with the resource
	Attributes *ResolverRuleAssociationAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *ResolverRuleAssociationProperties `pulumi:"properties"`
}

type ResolverRuleAssociationState struct {
	// The attributes associated with the resource
	Attributes ResolverRuleAssociationAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResolverRuleAssociationPropertiesPtrInput
}

func (ResolverRuleAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleAssociationState)(nil)).Elem()
}

type resolverRuleAssociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties ResolverRuleAssociationProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a ResolverRuleAssociation resource.
type ResolverRuleAssociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties ResolverRuleAssociationPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (ResolverRuleAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resolverRuleAssociationArgs)(nil)).Elem()
}

type ResolverRuleAssociationInput interface {
	pulumi.Input

	ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput
	ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput
}

func (*ResolverRuleAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRuleAssociation)(nil))
}

func (i *ResolverRuleAssociation) ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput {
	return i.ToResolverRuleAssociationOutputWithContext(context.Background())
}

func (i *ResolverRuleAssociation) ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResolverRuleAssociationOutput)
}

type ResolverRuleAssociationOutput struct {
	*pulumi.OutputState
}

func (ResolverRuleAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResolverRuleAssociation)(nil))
}

func (o ResolverRuleAssociationOutput) ToResolverRuleAssociationOutput() ResolverRuleAssociationOutput {
	return o
}

func (o ResolverRuleAssociationOutput) ToResolverRuleAssociationOutputWithContext(ctx context.Context) ResolverRuleAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResolverRuleAssociationOutput{})
}
