// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::GatewayResponse
//
// Deprecated: GatewayResponse is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type GatewayResponse struct {
	pulumi.CustomResourceState

	// The response parameters (paths, query strings, and headers) for the response.
	ResponseParameters pulumi.AnyOutput `pulumi:"responseParameters"`
	// The response templates for the response.
	ResponseTemplates pulumi.AnyOutput `pulumi:"responseTemplates"`
	// The type of the Gateway Response.
	ResponseType pulumi.StringOutput `pulumi:"responseType"`
	// The identifier of the API.
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
	// The HTTP status code for the response.
	StatusCode pulumi.StringPtrOutput `pulumi:"statusCode"`
}

// NewGatewayResponse registers a new resource with the given unique name, arguments, and options.
func NewGatewayResponse(ctx *pulumi.Context,
	name string, args *GatewayResponseArgs, opts ...pulumi.ResourceOption) (*GatewayResponse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResponseType == nil {
		return nil, errors.New("invalid value for required argument 'ResponseType'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource GatewayResponse
	err := ctx.RegisterResource("aws-native:apigateway:GatewayResponse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGatewayResponse gets an existing GatewayResponse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGatewayResponse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayResponseState, opts ...pulumi.ResourceOption) (*GatewayResponse, error) {
	var resource GatewayResponse
	err := ctx.ReadResource("aws-native:apigateway:GatewayResponse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GatewayResponse resources.
type gatewayResponseState struct {
}

type GatewayResponseState struct {
}

func (GatewayResponseState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayResponseState)(nil)).Elem()
}

type gatewayResponseArgs struct {
	// The response parameters (paths, query strings, and headers) for the response.
	ResponseParameters interface{} `pulumi:"responseParameters"`
	// The response templates for the response.
	ResponseTemplates interface{} `pulumi:"responseTemplates"`
	// The type of the Gateway Response.
	ResponseType string `pulumi:"responseType"`
	// The identifier of the API.
	RestApiId string `pulumi:"restApiId"`
	// The HTTP status code for the response.
	StatusCode *string `pulumi:"statusCode"`
}

// The set of arguments for constructing a GatewayResponse resource.
type GatewayResponseArgs struct {
	// The response parameters (paths, query strings, and headers) for the response.
	ResponseParameters pulumi.Input
	// The response templates for the response.
	ResponseTemplates pulumi.Input
	// The type of the Gateway Response.
	ResponseType pulumi.StringInput
	// The identifier of the API.
	RestApiId pulumi.StringInput
	// The HTTP status code for the response.
	StatusCode pulumi.StringPtrInput
}

func (GatewayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayResponseArgs)(nil)).Elem()
}

type GatewayResponseInput interface {
	pulumi.Input

	ToGatewayResponseOutput() GatewayResponseOutput
	ToGatewayResponseOutputWithContext(ctx context.Context) GatewayResponseOutput
}

func (*GatewayResponse) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayResponse)(nil))
}

func (i *GatewayResponse) ToGatewayResponseOutput() GatewayResponseOutput {
	return i.ToGatewayResponseOutputWithContext(context.Background())
}

func (i *GatewayResponse) ToGatewayResponseOutputWithContext(ctx context.Context) GatewayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayResponseOutput)
}

type GatewayResponseOutput struct{ *pulumi.OutputState }

func (GatewayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GatewayResponse)(nil))
}

func (o GatewayResponseOutput) ToGatewayResponseOutput() GatewayResponseOutput {
	return o
}

func (o GatewayResponseOutput) ToGatewayResponseOutputWithContext(ctx context.Context) GatewayResponseOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayResponseInput)(nil)).Elem(), &GatewayResponse{})
	pulumi.RegisterOutputType(GatewayResponseOutput{})
}
