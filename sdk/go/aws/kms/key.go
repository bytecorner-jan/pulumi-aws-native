// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::KMS::Key resource specifies a customer master key (CMK) in AWS Key Management Service (AWS KMS). Authorized users can use the CMK to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use CMKs to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.
type Key struct {
	pulumi.CustomResourceState

	Arn pulumi.StringOutput `pulumi:"arn"`
	// A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Enables automatic rotation of the key material for the specified customer master key (CMK). By default, automation key rotation is not enabled.
	EnableKeyRotation pulumi.BoolPtrOutput `pulumi:"enableKeyRotation"`
	// Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	KeyId   pulumi.StringOutput  `pulumi:"keyId"`
	// The key policy that authorizes use of the CMK. The key policy must observe the following rules.
	KeyPolicy pulumi.AnyOutput `pulumi:"keyPolicy"`
	// Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.
	KeySpec KeySpecPtrOutput `pulumi:"keySpec"`
	// Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.
	KeyUsage KeyUsagePtrOutput `pulumi:"keyUsage"`
	// Specifies whether the CMK should be Multi-Region. You can't change the MultiRegion value after the CMK is created.
	MultiRegion pulumi.BoolPtrOutput `pulumi:"multiRegion"`
	// Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
	PendingWindowInDays pulumi.IntPtrOutput `pulumi:"pendingWindowInDays"`
	// An array of key-value pairs to apply to this resource.
	Tags KeyTagArrayOutput `pulumi:"tags"`
}

// NewKey registers a new resource with the given unique name, arguments, and options.
func NewKey(ctx *pulumi.Context,
	name string, args *KeyArgs, opts ...pulumi.ResourceOption) (*Key, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyPolicy == nil {
		return nil, errors.New("invalid value for required argument 'KeyPolicy'")
	}
	var resource Key
	err := ctx.RegisterResource("aws-native:kms:Key", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKey gets an existing Key resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyState, opts ...pulumi.ResourceOption) (*Key, error) {
	var resource Key
	err := ctx.ReadResource("aws-native:kms:Key", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Key resources.
type keyState struct {
}

type KeyState struct {
}

func (KeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyState)(nil)).Elem()
}

type keyArgs struct {
	// A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.
	Description *string `pulumi:"description"`
	// Enables automatic rotation of the key material for the specified customer master key (CMK). By default, automation key rotation is not enabled.
	EnableKeyRotation *bool `pulumi:"enableKeyRotation"`
	// Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.
	Enabled *bool `pulumi:"enabled"`
	// The key policy that authorizes use of the CMK. The key policy must observe the following rules.
	KeyPolicy interface{} `pulumi:"keyPolicy"`
	// Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.
	KeySpec *KeySpec `pulumi:"keySpec"`
	// Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.
	KeyUsage *KeyUsage `pulumi:"keyUsage"`
	// Specifies whether the CMK should be Multi-Region. You can't change the MultiRegion value after the CMK is created.
	MultiRegion *bool `pulumi:"multiRegion"`
	// Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
	PendingWindowInDays *int `pulumi:"pendingWindowInDays"`
	// An array of key-value pairs to apply to this resource.
	Tags []KeyTag `pulumi:"tags"`
}

// The set of arguments for constructing a Key resource.
type KeyArgs struct {
	// A description of the CMK. Use a description that helps you to distinguish this CMK from others in the account, such as its intended use.
	Description pulumi.StringPtrInput
	// Enables automatic rotation of the key material for the specified customer master key (CMK). By default, automation key rotation is not enabled.
	EnableKeyRotation pulumi.BoolPtrInput
	// Specifies whether the customer master key (CMK) is enabled. Disabled CMKs cannot be used in cryptographic operations.
	Enabled pulumi.BoolPtrInput
	// The key policy that authorizes use of the CMK. The key policy must observe the following rules.
	KeyPolicy pulumi.Input
	// Specifies the type of CMK to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric CMKs. You can't change the KeySpec value after the CMK is created.
	KeySpec KeySpecPtrInput
	// Determines the cryptographic operations for which you can use the CMK. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric CMKs. You can't change the KeyUsage value after the CMK is created.
	KeyUsage KeyUsagePtrInput
	// Specifies whether the CMK should be Multi-Region. You can't change the MultiRegion value after the CMK is created.
	MultiRegion pulumi.BoolPtrInput
	// Specifies the number of days in the waiting period before AWS KMS deletes a CMK that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
	PendingWindowInDays pulumi.IntPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags KeyTagArrayInput
}

func (KeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyArgs)(nil)).Elem()
}

type KeyInput interface {
	pulumi.Input

	ToKeyOutput() KeyOutput
	ToKeyOutputWithContext(ctx context.Context) KeyOutput
}

func (*Key) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (i *Key) ToKeyOutput() KeyOutput {
	return i.ToKeyOutputWithContext(context.Background())
}

func (i *Key) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyOutput)
}

type KeyOutput struct{ *pulumi.OutputState }

func (KeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Key)(nil))
}

func (o KeyOutput) ToKeyOutput() KeyOutput {
	return o
}

func (o KeyOutput) ToKeyOutputWithContext(ctx context.Context) KeyOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyInput)(nil)).Elem(), &Key{})
	pulumi.RegisterOutputType(KeyOutput{})
}
