// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html
type LayerVersionPermission struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-action
	Action pulumi.StringOutput `pulumi:"action"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-layerversionarn
	LayerVersionArn pulumi.StringOutput `pulumi:"layerVersionArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-organizationid
	OrganizationId pulumi.StringPtrOutput `pulumi:"organizationId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-principal
	Principal pulumi.StringOutput `pulumi:"principal"`
}

// NewLayerVersionPermission registers a new resource with the given unique name, arguments, and options.
func NewLayerVersionPermission(ctx *pulumi.Context,
	name string, args *LayerVersionPermissionArgs, opts ...pulumi.ResourceOption) (*LayerVersionPermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.LayerVersionArn == nil {
		return nil, errors.New("invalid value for required argument 'LayerVersionArn'")
	}
	if args.Principal == nil {
		return nil, errors.New("invalid value for required argument 'Principal'")
	}
	var resource LayerVersionPermission
	err := ctx.RegisterResource("aws-native:lambda:LayerVersionPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLayerVersionPermission gets an existing LayerVersionPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLayerVersionPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LayerVersionPermissionState, opts ...pulumi.ResourceOption) (*LayerVersionPermission, error) {
	var resource LayerVersionPermission
	err := ctx.ReadResource("aws-native:lambda:LayerVersionPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LayerVersionPermission resources.
type layerVersionPermissionState struct {
}

type LayerVersionPermissionState struct {
}

func (LayerVersionPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionPermissionState)(nil)).Elem()
}

type layerVersionPermissionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-action
	Action string `pulumi:"action"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-layerversionarn
	LayerVersionArn string `pulumi:"layerVersionArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-organizationid
	OrganizationId *string `pulumi:"organizationId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-principal
	Principal string `pulumi:"principal"`
}

// The set of arguments for constructing a LayerVersionPermission resource.
type LayerVersionPermissionArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-action
	Action pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-layerversionarn
	LayerVersionArn pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-organizationid
	OrganizationId pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-layerversionpermission.html#cfn-lambda-layerversionpermission-principal
	Principal pulumi.StringInput
}

func (LayerVersionPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*layerVersionPermissionArgs)(nil)).Elem()
}

type LayerVersionPermissionInput interface {
	pulumi.Input

	ToLayerVersionPermissionOutput() LayerVersionPermissionOutput
	ToLayerVersionPermissionOutputWithContext(ctx context.Context) LayerVersionPermissionOutput
}

func (*LayerVersionPermission) ElementType() reflect.Type {
	return reflect.TypeOf((*LayerVersionPermission)(nil))
}

func (i *LayerVersionPermission) ToLayerVersionPermissionOutput() LayerVersionPermissionOutput {
	return i.ToLayerVersionPermissionOutputWithContext(context.Background())
}

func (i *LayerVersionPermission) ToLayerVersionPermissionOutputWithContext(ctx context.Context) LayerVersionPermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LayerVersionPermissionOutput)
}

type LayerVersionPermissionOutput struct{ *pulumi.OutputState }

func (LayerVersionPermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LayerVersionPermission)(nil))
}

func (o LayerVersionPermissionOutput) ToLayerVersionPermissionOutput() LayerVersionPermissionOutput {
	return o
}

func (o LayerVersionPermissionOutput) ToLayerVersionPermissionOutputWithContext(ctx context.Context) LayerVersionPermissionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LayerVersionPermissionOutput{})
}
