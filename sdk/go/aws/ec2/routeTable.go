// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::RouteTable
//
// Deprecated: RouteTable is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type RouteTable struct {
	pulumi.CustomResourceState

	Tags  RouteTableTagArrayOutput `pulumi:"tags"`
	VpcId pulumi.StringOutput      `pulumi:"vpcId"`
}

// NewRouteTable registers a new resource with the given unique name, arguments, and options.
func NewRouteTable(ctx *pulumi.Context,
	name string, args *RouteTableArgs, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource RouteTable
	err := ctx.RegisterResource("aws-native:ec2:RouteTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouteTable gets an existing RouteTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouteTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteTableState, opts ...pulumi.ResourceOption) (*RouteTable, error) {
	var resource RouteTable
	err := ctx.ReadResource("aws-native:ec2:RouteTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouteTable resources.
type routeTableState struct {
}

type RouteTableState struct {
}

func (RouteTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableState)(nil)).Elem()
}

type routeTableArgs struct {
	Tags  []RouteTableTag `pulumi:"tags"`
	VpcId string          `pulumi:"vpcId"`
}

// The set of arguments for constructing a RouteTable resource.
type RouteTableArgs struct {
	Tags  RouteTableTagArrayInput
	VpcId pulumi.StringInput
}

func (RouteTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeTableArgs)(nil)).Elem()
}

type RouteTableInput interface {
	pulumi.Input

	ToRouteTableOutput() RouteTableOutput
	ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput
}

func (*RouteTable) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTable)(nil))
}

func (i *RouteTable) ToRouteTableOutput() RouteTableOutput {
	return i.ToRouteTableOutputWithContext(context.Background())
}

func (i *RouteTable) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteTableOutput)
}

type RouteTableOutput struct{ *pulumi.OutputState }

func (RouteTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteTable)(nil))
}

func (o RouteTableOutput) ToRouteTableOutput() RouteTableOutput {
	return o
}

func (o RouteTableOutput) ToRouteTableOutputWithContext(ctx context.Context) RouteTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RouteTableOutput{})
}
