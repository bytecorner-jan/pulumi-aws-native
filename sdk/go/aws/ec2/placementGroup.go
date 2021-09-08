// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html
type PlacementGroup struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-strategy
	Strategy pulumi.StringPtrOutput `pulumi:"strategy"`
}

// NewPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewPlacementGroup(ctx *pulumi.Context,
	name string, args *PlacementGroupArgs, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	if args == nil {
		args = &PlacementGroupArgs{}
	}

	var resource PlacementGroup
	err := ctx.RegisterResource("aws-native:ec2:PlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlacementGroup gets an existing PlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlacementGroupState, opts ...pulumi.ResourceOption) (*PlacementGroup, error) {
	var resource PlacementGroup
	err := ctx.ReadResource("aws-native:ec2:PlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlacementGroup resources.
type placementGroupState struct {
}

type PlacementGroupState struct {
}

func (PlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupState)(nil)).Elem()
}

type placementGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-strategy
	Strategy *string `pulumi:"strategy"`
}

// The set of arguments for constructing a PlacementGroup resource.
type PlacementGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-placementgroup.html#cfn-ec2-placementgroup-strategy
	Strategy pulumi.StringPtrInput
}

func (PlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*placementGroupArgs)(nil)).Elem()
}

type PlacementGroupInput interface {
	pulumi.Input

	ToPlacementGroupOutput() PlacementGroupOutput
	ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput
}

func (*PlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroup)(nil))
}

func (i *PlacementGroup) ToPlacementGroupOutput() PlacementGroupOutput {
	return i.ToPlacementGroupOutputWithContext(context.Background())
}

func (i *PlacementGroup) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlacementGroupOutput)
}

type PlacementGroupOutput struct{ *pulumi.OutputState }

func (PlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlacementGroup)(nil))
}

func (o PlacementGroupOutput) ToPlacementGroupOutput() PlacementGroupOutput {
	return o
}

func (o PlacementGroupOutput) ToPlacementGroupOutputWithContext(ctx context.Context) PlacementGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PlacementGroupOutput{})
}
