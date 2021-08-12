// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html
type SAMLProvider struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-samlmetadatadocument
	SamlMetadataDocument pulumi.StringOutput `pulumi:"samlMetadataDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewSAMLProvider registers a new resource with the given unique name, arguments, and options.
func NewSAMLProvider(ctx *pulumi.Context,
	name string, args *SAMLProviderArgs, opts ...pulumi.ResourceOption) (*SAMLProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SamlMetadataDocument == nil {
		return nil, errors.New("invalid value for required argument 'SamlMetadataDocument'")
	}
	var resource SAMLProvider
	err := ctx.RegisterResource("aws-native:IAM:SAMLProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSAMLProvider gets an existing SAMLProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSAMLProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SAMLProviderState, opts ...pulumi.ResourceOption) (*SAMLProvider, error) {
	var resource SAMLProvider
	err := ctx.ReadResource("aws-native:IAM:SAMLProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SAMLProvider resources.
type samlproviderState struct {
}

type SAMLProviderState struct {
}

func (SAMLProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlproviderState)(nil)).Elem()
}

type samlproviderArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-name
	Name *string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-samlmetadatadocument
	SamlMetadataDocument string `pulumi:"samlMetadataDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a SAMLProvider resource.
type SAMLProviderArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-name
	Name pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-samlmetadatadocument
	SamlMetadataDocument pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-samlprovider.html#cfn-iam-samlprovider-tags
	Tags aws.TagArrayInput
}

func (SAMLProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlproviderArgs)(nil)).Elem()
}

type SAMLProviderInput interface {
	pulumi.Input

	ToSAMLProviderOutput() SAMLProviderOutput
	ToSAMLProviderOutputWithContext(ctx context.Context) SAMLProviderOutput
}

func (*SAMLProvider) ElementType() reflect.Type {
	return reflect.TypeOf((*SAMLProvider)(nil))
}

func (i *SAMLProvider) ToSAMLProviderOutput() SAMLProviderOutput {
	return i.ToSAMLProviderOutputWithContext(context.Background())
}

func (i *SAMLProvider) ToSAMLProviderOutputWithContext(ctx context.Context) SAMLProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SAMLProviderOutput)
}

type SAMLProviderOutput struct{ *pulumi.OutputState }

func (SAMLProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SAMLProvider)(nil))
}

func (o SAMLProviderOutput) ToSAMLProviderOutput() SAMLProviderOutput {
	return o
}

func (o SAMLProviderOutput) ToSAMLProviderOutputWithContext(ctx context.Context) SAMLProviderOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SAMLProviderOutput{})
}
