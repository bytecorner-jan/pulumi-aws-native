// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalog

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html
type TagOptionAssociation struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-resourceid
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-tagoptionid
	TagOptionId pulumi.StringOutput `pulumi:"tagOptionId"`
}

// NewTagOptionAssociation registers a new resource with the given unique name, arguments, and options.
func NewTagOptionAssociation(ctx *pulumi.Context,
	name string, args *TagOptionAssociationArgs, opts ...pulumi.ResourceOption) (*TagOptionAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceId'")
	}
	if args.TagOptionId == nil {
		return nil, errors.New("invalid value for required argument 'TagOptionId'")
	}
	var resource TagOptionAssociation
	err := ctx.RegisterResource("aws-native:servicecatalog:TagOptionAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTagOptionAssociation gets an existing TagOptionAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTagOptionAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TagOptionAssociationState, opts ...pulumi.ResourceOption) (*TagOptionAssociation, error) {
	var resource TagOptionAssociation
	err := ctx.ReadResource("aws-native:servicecatalog:TagOptionAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TagOptionAssociation resources.
type tagOptionAssociationState struct {
}

type TagOptionAssociationState struct {
}

func (TagOptionAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionAssociationState)(nil)).Elem()
}

type tagOptionAssociationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-resourceid
	ResourceId string `pulumi:"resourceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-tagoptionid
	TagOptionId string `pulumi:"tagOptionId"`
}

// The set of arguments for constructing a TagOptionAssociation resource.
type TagOptionAssociationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-resourceid
	ResourceId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoptionassociation.html#cfn-servicecatalog-tagoptionassociation-tagoptionid
	TagOptionId pulumi.StringInput
}

func (TagOptionAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tagOptionAssociationArgs)(nil)).Elem()
}

type TagOptionAssociationInput interface {
	pulumi.Input

	ToTagOptionAssociationOutput() TagOptionAssociationOutput
	ToTagOptionAssociationOutputWithContext(ctx context.Context) TagOptionAssociationOutput
}

func (*TagOptionAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOptionAssociation)(nil))
}

func (i *TagOptionAssociation) ToTagOptionAssociationOutput() TagOptionAssociationOutput {
	return i.ToTagOptionAssociationOutputWithContext(context.Background())
}

func (i *TagOptionAssociation) ToTagOptionAssociationOutputWithContext(ctx context.Context) TagOptionAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOptionAssociationOutput)
}

type TagOptionAssociationOutput struct{ *pulumi.OutputState }

func (TagOptionAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagOptionAssociation)(nil))
}

func (o TagOptionAssociationOutput) ToTagOptionAssociationOutput() TagOptionAssociationOutput {
	return o
}

func (o TagOptionAssociationOutput) ToTagOptionAssociationOutputWithContext(ctx context.Context) TagOptionAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TagOptionAssociationOutput{})
}
