// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53recoverycontrol

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html
type RoutingControl struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-clusterarn
	ClusterArn pulumi.StringPtrOutput `pulumi:"clusterArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-controlpanelarn
	ControlPanelArn pulumi.StringPtrOutput `pulumi:"controlPanelArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-name
	Name              pulumi.StringOutput `pulumi:"name"`
	RoutingControlArn pulumi.StringOutput `pulumi:"routingControlArn"`
	Status            pulumi.StringOutput `pulumi:"status"`
}

// NewRoutingControl registers a new resource with the given unique name, arguments, and options.
func NewRoutingControl(ctx *pulumi.Context,
	name string, args *RoutingControlArgs, opts ...pulumi.ResourceOption) (*RoutingControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource RoutingControl
	err := ctx.RegisterResource("aws-native:Route53RecoveryControl:RoutingControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoutingControl gets an existing RoutingControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoutingControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoutingControlState, opts ...pulumi.ResourceOption) (*RoutingControl, error) {
	var resource RoutingControl
	err := ctx.ReadResource("aws-native:Route53RecoveryControl:RoutingControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoutingControl resources.
type routingControlState struct {
}

type RoutingControlState struct {
}

func (RoutingControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*routingControlState)(nil)).Elem()
}

type routingControlArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-clusterarn
	ClusterArn *string `pulumi:"clusterArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-controlpanelarn
	ControlPanelArn *string `pulumi:"controlPanelArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-name
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a RoutingControl resource.
type RoutingControlArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-clusterarn
	ClusterArn pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-controlpanelarn
	ControlPanelArn pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53recoverycontrol-routingcontrol.html#cfn-route53recoverycontrol-routingcontrol-name
	Name pulumi.StringInput
}

func (RoutingControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routingControlArgs)(nil)).Elem()
}

type RoutingControlInput interface {
	pulumi.Input

	ToRoutingControlOutput() RoutingControlOutput
	ToRoutingControlOutputWithContext(ctx context.Context) RoutingControlOutput
}

func (*RoutingControl) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingControl)(nil))
}

func (i *RoutingControl) ToRoutingControlOutput() RoutingControlOutput {
	return i.ToRoutingControlOutputWithContext(context.Background())
}

func (i *RoutingControl) ToRoutingControlOutputWithContext(ctx context.Context) RoutingControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingControlOutput)
}

type RoutingControlOutput struct{ *pulumi.OutputState }

func (RoutingControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingControl)(nil))
}

func (o RoutingControlOutput) ToRoutingControlOutput() RoutingControlOutput {
	return o
}

func (o RoutingControlOutput) ToRoutingControlOutputWithContext(ctx context.Context) RoutingControlOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RoutingControlOutput{})
}
