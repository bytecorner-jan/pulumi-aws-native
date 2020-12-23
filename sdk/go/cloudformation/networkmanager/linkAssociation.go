// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-linkassociation.html
type LinkAssociation struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes LinkAssociationAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties LinkAssociationPropertiesOutput `pulumi:"properties"`
}

// NewLinkAssociation registers a new resource with the given unique name, arguments, and options.
func NewLinkAssociation(ctx *pulumi.Context,
	name string, args *LinkAssociationArgs, opts ...pulumi.ResourceOption) (*LinkAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource LinkAssociation
	err := ctx.RegisterResource("cloudformation:NetworkManager:LinkAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkAssociation gets an existing LinkAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkAssociationState, opts ...pulumi.ResourceOption) (*LinkAssociation, error) {
	var resource LinkAssociation
	err := ctx.ReadResource("cloudformation:NetworkManager:LinkAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkAssociation resources.
type linkAssociationState struct {
	// The attributes associated with the resource
	Attributes *LinkAssociationAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *LinkAssociationProperties `pulumi:"properties"`
}

type LinkAssociationState struct {
	// The attributes associated with the resource
	Attributes LinkAssociationAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties LinkAssociationPropertiesPtrInput
}

func (LinkAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAssociationState)(nil)).Elem()
}

type linkAssociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties LinkAssociationProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a LinkAssociation resource.
type LinkAssociationArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties LinkAssociationPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (LinkAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkAssociationArgs)(nil)).Elem()
}

type LinkAssociationInput interface {
	pulumi.Input

	ToLinkAssociationOutput() LinkAssociationOutput
	ToLinkAssociationOutputWithContext(ctx context.Context) LinkAssociationOutput
}

func (*LinkAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkAssociation)(nil))
}

func (i *LinkAssociation) ToLinkAssociationOutput() LinkAssociationOutput {
	return i.ToLinkAssociationOutputWithContext(context.Background())
}

func (i *LinkAssociation) ToLinkAssociationOutputWithContext(ctx context.Context) LinkAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkAssociationOutput)
}

type LinkAssociationOutput struct {
	*pulumi.OutputState
}

func (LinkAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkAssociation)(nil))
}

func (o LinkAssociationOutput) ToLinkAssociationOutput() LinkAssociationOutput {
	return o
}

func (o LinkAssociationOutput) ToLinkAssociationOutputWithContext(ctx context.Context) LinkAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkAssociationOutput{})
}
