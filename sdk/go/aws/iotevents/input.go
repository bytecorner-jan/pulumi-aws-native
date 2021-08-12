// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iotevents

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html
type Input struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
	InputDefinition InputInputDefinitionOutput `pulumi:"inputDefinition"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
	InputDescription pulumi.StringPtrOutput `pulumi:"inputDescription"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
	InputName pulumi.StringPtrOutput `pulumi:"inputName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewInput registers a new resource with the given unique name, arguments, and options.
func NewInput(ctx *pulumi.Context,
	name string, args *InputArgs, opts ...pulumi.ResourceOption) (*Input, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InputDefinition == nil {
		return nil, errors.New("invalid value for required argument 'InputDefinition'")
	}
	var resource Input
	err := ctx.RegisterResource("aws-native:IoTEvents:Input", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInput gets an existing Input resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInput(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InputState, opts ...pulumi.ResourceOption) (*Input, error) {
	var resource Input
	err := ctx.ReadResource("aws-native:IoTEvents:Input", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Input resources.
type inputState struct {
}

type InputState struct {
}

func (InputState) ElementType() reflect.Type {
	return reflect.TypeOf((*inputState)(nil)).Elem()
}

type inputArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
	InputDefinition InputInputDefinition `pulumi:"inputDefinition"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
	InputDescription *string `pulumi:"inputDescription"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
	InputName *string `pulumi:"inputName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a Input resource.
type InputArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdefinition
	InputDefinition InputInputDefinitionInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputdescription
	InputDescription pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-inputname
	InputName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotevents-input.html#cfn-iotevents-input-tags
	Tags aws.TagArrayInput
}

func (InputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inputArgs)(nil)).Elem()
}

type InputInput interface {
	pulumi.Input

	ToInputOutput() InputOutput
	ToInputOutputWithContext(ctx context.Context) InputOutput
}

func (*Input) ElementType() reflect.Type {
	return reflect.TypeOf((*Input)(nil))
}

func (i *Input) ToInputOutput() InputOutput {
	return i.ToInputOutputWithContext(context.Background())
}

func (i *Input) ToInputOutputWithContext(ctx context.Context) InputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputOutput)
}

type InputOutput struct{ *pulumi.OutputState }

func (InputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Input)(nil))
}

func (o InputOutput) ToInputOutput() InputOutput {
	return o
}

func (o InputOutput) ToInputOutputWithContext(ctx context.Context) InputOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InputOutput{})
}
