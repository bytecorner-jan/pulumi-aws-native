// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html
type Resource struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-parentid
	ParentId pulumi.StringOutput `pulumi:"parentId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-pathpart
	PathPart   pulumi.StringOutput `pulumi:"pathPart"`
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-restapiid
	RestApiId pulumi.StringOutput `pulumi:"restApiId"`
}

// NewResource registers a new resource with the given unique name, arguments, and options.
func NewResource(ctx *pulumi.Context,
	name string, args *ResourceArgs, opts ...pulumi.ResourceOption) (*Resource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	if args.PathPart == nil {
		return nil, errors.New("invalid value for required argument 'PathPart'")
	}
	if args.RestApiId == nil {
		return nil, errors.New("invalid value for required argument 'RestApiId'")
	}
	var resource Resource
	err := ctx.RegisterResource("aws-native:ApiGateway:Resource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResource gets an existing Resource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResourceState, opts ...pulumi.ResourceOption) (*Resource, error) {
	var resource Resource
	err := ctx.ReadResource("aws-native:ApiGateway:Resource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Resource resources.
type resourceState struct {
}

type ResourceState struct {
}

func (ResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceState)(nil)).Elem()
}

type resourceArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-parentid
	ParentId string `pulumi:"parentId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-pathpart
	PathPart string `pulumi:"pathPart"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-restapiid
	RestApiId string `pulumi:"restApiId"`
}

// The set of arguments for constructing a Resource resource.
type ResourceArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-parentid
	ParentId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-pathpart
	PathPart pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-restapiid
	RestApiId pulumi.StringInput
}

func (ResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*resourceArgs)(nil)).Elem()
}

type ResourceInput interface {
	pulumi.Input

	ToResourceOutput() ResourceOutput
	ToResourceOutputWithContext(ctx context.Context) ResourceOutput
}

func (*Resource) ElementType() reflect.Type {
	return reflect.TypeOf((*Resource)(nil))
}

func (i *Resource) ToResourceOutput() ResourceOutput {
	return i.ToResourceOutputWithContext(context.Background())
}

func (i *Resource) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceOutput)
}

type ResourceOutput struct{ *pulumi.OutputState }

func (ResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Resource)(nil))
}

func (o ResourceOutput) ToResourceOutput() ResourceOutput {
	return o
}

func (o ResourceOutput) ToResourceOutputWithContext(ctx context.Context) ResourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ResourceOutput{})
}
