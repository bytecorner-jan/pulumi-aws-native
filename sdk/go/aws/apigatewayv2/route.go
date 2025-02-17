// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigatewayv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGatewayV2::Route
//
// Deprecated: Route is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Route struct {
	pulumi.CustomResourceState

	ApiId                            pulumi.StringOutput      `pulumi:"apiId"`
	ApiKeyRequired                   pulumi.BoolPtrOutput     `pulumi:"apiKeyRequired"`
	AuthorizationScopes              pulumi.StringArrayOutput `pulumi:"authorizationScopes"`
	AuthorizationType                pulumi.StringPtrOutput   `pulumi:"authorizationType"`
	AuthorizerId                     pulumi.StringPtrOutput   `pulumi:"authorizerId"`
	ModelSelectionExpression         pulumi.StringPtrOutput   `pulumi:"modelSelectionExpression"`
	OperationName                    pulumi.StringPtrOutput   `pulumi:"operationName"`
	RequestModels                    pulumi.AnyOutput         `pulumi:"requestModels"`
	RequestParameters                pulumi.AnyOutput         `pulumi:"requestParameters"`
	RouteKey                         pulumi.StringOutput      `pulumi:"routeKey"`
	RouteResponseSelectionExpression pulumi.StringPtrOutput   `pulumi:"routeResponseSelectionExpression"`
	Target                           pulumi.StringPtrOutput   `pulumi:"target"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiId == nil {
		return nil, errors.New("invalid value for required argument 'ApiId'")
	}
	if args.RouteKey == nil {
		return nil, errors.New("invalid value for required argument 'RouteKey'")
	}
	var resource Route
	err := ctx.RegisterResource("aws-native:apigatewayv2:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("aws-native:apigatewayv2:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
}

type RouteState struct {
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	ApiId                            string      `pulumi:"apiId"`
	ApiKeyRequired                   *bool       `pulumi:"apiKeyRequired"`
	AuthorizationScopes              []string    `pulumi:"authorizationScopes"`
	AuthorizationType                *string     `pulumi:"authorizationType"`
	AuthorizerId                     *string     `pulumi:"authorizerId"`
	ModelSelectionExpression         *string     `pulumi:"modelSelectionExpression"`
	OperationName                    *string     `pulumi:"operationName"`
	RequestModels                    interface{} `pulumi:"requestModels"`
	RequestParameters                interface{} `pulumi:"requestParameters"`
	RouteKey                         string      `pulumi:"routeKey"`
	RouteResponseSelectionExpression *string     `pulumi:"routeResponseSelectionExpression"`
	Target                           *string     `pulumi:"target"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	ApiId                            pulumi.StringInput
	ApiKeyRequired                   pulumi.BoolPtrInput
	AuthorizationScopes              pulumi.StringArrayInput
	AuthorizationType                pulumi.StringPtrInput
	AuthorizerId                     pulumi.StringPtrInput
	ModelSelectionExpression         pulumi.StringPtrInput
	OperationName                    pulumi.StringPtrInput
	RequestModels                    pulumi.Input
	RequestParameters                pulumi.Input
	RouteKey                         pulumi.StringInput
	RouteResponseSelectionExpression pulumi.StringPtrInput
	Target                           pulumi.StringPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterOutputType(RouteOutput{})
}
