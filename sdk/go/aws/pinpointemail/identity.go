// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pinpointemail

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html
type Identity struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-dkimsigningenabled
	DkimSigningEnabled pulumi.BoolPtrOutput `pulumi:"dkimSigningEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-feedbackforwardingenabled
	FeedbackForwardingEnabled pulumi.BoolPtrOutput `pulumi:"feedbackForwardingEnabled"`
	IdentityDNSRecordName1    pulumi.StringOutput  `pulumi:"identityDNSRecordName1"`
	IdentityDNSRecordName2    pulumi.StringOutput  `pulumi:"identityDNSRecordName2"`
	IdentityDNSRecordName3    pulumi.StringOutput  `pulumi:"identityDNSRecordName3"`
	IdentityDNSRecordValue1   pulumi.StringOutput  `pulumi:"identityDNSRecordValue1"`
	IdentityDNSRecordValue2   pulumi.StringOutput  `pulumi:"identityDNSRecordValue2"`
	IdentityDNSRecordValue3   pulumi.StringOutput  `pulumi:"identityDNSRecordValue3"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-mailfromattributes
	MailFromAttributes IdentityMailFromAttributesPtrOutput `pulumi:"mailFromAttributes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-name
	Name pulumi.StringOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-tags
	Tags IdentityTagsArrayOutput `pulumi:"tags"`
}

// NewIdentity registers a new resource with the given unique name, arguments, and options.
func NewIdentity(ctx *pulumi.Context,
	name string, args *IdentityArgs, opts ...pulumi.ResourceOption) (*Identity, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	var resource Identity
	err := ctx.RegisterResource("aws-native:pinpointemail:Identity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentity gets an existing Identity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityState, opts ...pulumi.ResourceOption) (*Identity, error) {
	var resource Identity
	err := ctx.ReadResource("aws-native:pinpointemail:Identity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Identity resources.
type identityState struct {
}

type IdentityState struct {
}

func (IdentityState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityState)(nil)).Elem()
}

type identityArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-dkimsigningenabled
	DkimSigningEnabled *bool `pulumi:"dkimSigningEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-feedbackforwardingenabled
	FeedbackForwardingEnabled *bool `pulumi:"feedbackForwardingEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-mailfromattributes
	MailFromAttributes *IdentityMailFromAttributes `pulumi:"mailFromAttributes"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-name
	Name string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-tags
	Tags []IdentityTags `pulumi:"tags"`
}

// The set of arguments for constructing a Identity resource.
type IdentityArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-dkimsigningenabled
	DkimSigningEnabled pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-feedbackforwardingenabled
	FeedbackForwardingEnabled pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-mailfromattributes
	MailFromAttributes IdentityMailFromAttributesPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-name
	Name pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-pinpointemail-identity.html#cfn-pinpointemail-identity-tags
	Tags IdentityTagsArrayInput
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityArgs)(nil)).Elem()
}

type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(ctx context.Context) IdentityOutput
}

func (*Identity) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil))
}

func (i *Identity) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i *Identity) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil))
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
}
