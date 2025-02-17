// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Lambda::Version
//
// Deprecated: Version is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Version struct {
	pulumi.CustomResourceState

	CodeSha256                   pulumi.StringPtrOutput                              `pulumi:"codeSha256"`
	Description                  pulumi.StringPtrOutput                              `pulumi:"description"`
	FunctionName                 pulumi.StringOutput                                 `pulumi:"functionName"`
	ProvisionedConcurrencyConfig VersionProvisionedConcurrencyConfigurationPtrOutput `pulumi:"provisionedConcurrencyConfig"`
	Version                      pulumi.StringOutput                                 `pulumi:"version"`
}

// NewVersion registers a new resource with the given unique name, arguments, and options.
func NewVersion(ctx *pulumi.Context,
	name string, args *VersionArgs, opts ...pulumi.ResourceOption) (*Version, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FunctionName == nil {
		return nil, errors.New("invalid value for required argument 'FunctionName'")
	}
	var resource Version
	err := ctx.RegisterResource("aws-native:lambda:Version", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVersion gets an existing Version resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VersionState, opts ...pulumi.ResourceOption) (*Version, error) {
	var resource Version
	err := ctx.ReadResource("aws-native:lambda:Version", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Version resources.
type versionState struct {
}

type VersionState struct {
}

func (VersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*versionState)(nil)).Elem()
}

type versionArgs struct {
	CodeSha256                   *string                                     `pulumi:"codeSha256"`
	Description                  *string                                     `pulumi:"description"`
	FunctionName                 string                                      `pulumi:"functionName"`
	ProvisionedConcurrencyConfig *VersionProvisionedConcurrencyConfiguration `pulumi:"provisionedConcurrencyConfig"`
}

// The set of arguments for constructing a Version resource.
type VersionArgs struct {
	CodeSha256                   pulumi.StringPtrInput
	Description                  pulumi.StringPtrInput
	FunctionName                 pulumi.StringInput
	ProvisionedConcurrencyConfig VersionProvisionedConcurrencyConfigurationPtrInput
}

func (VersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*versionArgs)(nil)).Elem()
}

type VersionInput interface {
	pulumi.Input

	ToVersionOutput() VersionOutput
	ToVersionOutputWithContext(ctx context.Context) VersionOutput
}

func (*Version) ElementType() reflect.Type {
	return reflect.TypeOf((*Version)(nil))
}

func (i *Version) ToVersionOutput() VersionOutput {
	return i.ToVersionOutputWithContext(context.Background())
}

func (i *Version) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionOutput)
}

type VersionOutput struct{ *pulumi.OutputState }

func (VersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Version)(nil))
}

func (o VersionOutput) ToVersionOutput() VersionOutput {
	return o
}

func (o VersionOutput) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VersionInput)(nil)).Elem(), &Version{})
	pulumi.RegisterOutputType(VersionOutput{})
}
