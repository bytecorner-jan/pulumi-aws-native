// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html
type Policy struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policydocument
	PolicyDocument pulumi.AnyOutput `pulumi:"policyDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policyname
	PolicyName pulumi.StringPtrOutput `pulumi:"policyName"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	var resource Policy
	err := ctx.RegisterResource("aws-native:iot:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("aws-native:iot:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
}

type PolicyState struct {
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policydocument
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policyname
	PolicyName *string `pulumi:"policyName"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policydocument
	PolicyDocument pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policy.html#cfn-iot-policy-policyname
	PolicyName pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
}
