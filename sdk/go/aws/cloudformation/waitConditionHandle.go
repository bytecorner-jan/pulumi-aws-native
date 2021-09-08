// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
type WaitConditionHandle struct {
	pulumi.CustomResourceState
}

// NewWaitConditionHandle registers a new resource with the given unique name, arguments, and options.
func NewWaitConditionHandle(ctx *pulumi.Context,
	name string, args *WaitConditionHandleArgs, opts ...pulumi.ResourceOption) (*WaitConditionHandle, error) {
	if args == nil {
		args = &WaitConditionHandleArgs{}
	}

	var resource WaitConditionHandle
	err := ctx.RegisterResource("aws-native:cloudformation:WaitConditionHandle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWaitConditionHandle gets an existing WaitConditionHandle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWaitConditionHandle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WaitConditionHandleState, opts ...pulumi.ResourceOption) (*WaitConditionHandle, error) {
	var resource WaitConditionHandle
	err := ctx.ReadResource("aws-native:cloudformation:WaitConditionHandle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WaitConditionHandle resources.
type waitConditionHandleState struct {
}

type WaitConditionHandleState struct {
}

func (WaitConditionHandleState) ElementType() reflect.Type {
	return reflect.TypeOf((*waitConditionHandleState)(nil)).Elem()
}

type waitConditionHandleArgs struct {
}

// The set of arguments for constructing a WaitConditionHandle resource.
type WaitConditionHandleArgs struct {
}

func (WaitConditionHandleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*waitConditionHandleArgs)(nil)).Elem()
}

type WaitConditionHandleInput interface {
	pulumi.Input

	ToWaitConditionHandleOutput() WaitConditionHandleOutput
	ToWaitConditionHandleOutputWithContext(ctx context.Context) WaitConditionHandleOutput
}

func (*WaitConditionHandle) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitConditionHandle)(nil))
}

func (i *WaitConditionHandle) ToWaitConditionHandleOutput() WaitConditionHandleOutput {
	return i.ToWaitConditionHandleOutputWithContext(context.Background())
}

func (i *WaitConditionHandle) ToWaitConditionHandleOutputWithContext(ctx context.Context) WaitConditionHandleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WaitConditionHandleOutput)
}

type WaitConditionHandleOutput struct{ *pulumi.OutputState }

func (WaitConditionHandleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WaitConditionHandle)(nil))
}

func (o WaitConditionHandleOutput) ToWaitConditionHandleOutput() WaitConditionHandleOutput {
	return o
}

func (o WaitConditionHandleOutput) ToWaitConditionHandleOutputWithContext(ctx context.Context) WaitConditionHandleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WaitConditionHandleOutput{})
}
