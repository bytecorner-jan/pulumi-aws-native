// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
type IdentityPool struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowclassicflow
	AllowClassicFlow pulumi.BoolPtrOutput `pulumi:"allowClassicFlow"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	AllowUnauthenticatedIdentities pulumi.BoolOutput `pulumi:"allowUnauthenticatedIdentities"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	CognitoEvents pulumi.AnyOutput `pulumi:"cognitoEvents"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	CognitoIdentityProviders IdentityPoolCognitoIdentityProviderArrayOutput `pulumi:"cognitoIdentityProviders"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	CognitoStreams IdentityPoolCognitoStreamsPtrOutput `pulumi:"cognitoStreams"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	DeveloperProviderName pulumi.StringPtrOutput `pulumi:"developerProviderName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	IdentityPoolName pulumi.StringPtrOutput `pulumi:"identityPoolName"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	OpenIdConnectProviderARNs pulumi.StringArrayOutput `pulumi:"openIdConnectProviderARNs"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	PushSync IdentityPoolPushSyncPtrOutput `pulumi:"pushSync"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	SamlProviderARNs pulumi.StringArrayOutput `pulumi:"samlProviderARNs"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	SupportedLoginProviders pulumi.AnyOutput `pulumi:"supportedLoginProviders"`
}

// NewIdentityPool registers a new resource with the given unique name, arguments, and options.
func NewIdentityPool(ctx *pulumi.Context,
	name string, args *IdentityPoolArgs, opts ...pulumi.ResourceOption) (*IdentityPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowUnauthenticatedIdentities == nil {
		return nil, errors.New("invalid value for required argument 'AllowUnauthenticatedIdentities'")
	}
	var resource IdentityPool
	err := ctx.RegisterResource("aws-native:cognito:IdentityPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityPool gets an existing IdentityPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityPoolState, opts ...pulumi.ResourceOption) (*IdentityPool, error) {
	var resource IdentityPool
	err := ctx.ReadResource("aws-native:cognito:IdentityPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityPool resources.
type identityPoolState struct {
}

type IdentityPoolState struct {
}

func (IdentityPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolState)(nil)).Elem()
}

type identityPoolArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowclassicflow
	AllowClassicFlow *bool `pulumi:"allowClassicFlow"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	AllowUnauthenticatedIdentities bool `pulumi:"allowUnauthenticatedIdentities"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	CognitoEvents interface{} `pulumi:"cognitoEvents"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	CognitoIdentityProviders []IdentityPoolCognitoIdentityProvider `pulumi:"cognitoIdentityProviders"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	CognitoStreams *IdentityPoolCognitoStreams `pulumi:"cognitoStreams"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	DeveloperProviderName *string `pulumi:"developerProviderName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	IdentityPoolName *string `pulumi:"identityPoolName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	OpenIdConnectProviderARNs []string `pulumi:"openIdConnectProviderARNs"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	PushSync *IdentityPoolPushSync `pulumi:"pushSync"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	SamlProviderARNs []string `pulumi:"samlProviderARNs"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	SupportedLoginProviders interface{} `pulumi:"supportedLoginProviders"`
}

// The set of arguments for constructing a IdentityPool resource.
type IdentityPoolArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowclassicflow
	AllowClassicFlow pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-allowunauthenticatedidentities
	AllowUnauthenticatedIdentities pulumi.BoolInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoevents
	CognitoEvents pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitoidentityproviders
	CognitoIdentityProviders IdentityPoolCognitoIdentityProviderArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-cognitostreams
	CognitoStreams IdentityPoolCognitoStreamsPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-developerprovidername
	DeveloperProviderName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-identitypoolname
	IdentityPoolName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-openidconnectproviderarns
	OpenIdConnectProviderARNs pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-pushsync
	PushSync IdentityPoolPushSyncPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-samlproviderarns
	SamlProviderARNs pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html#cfn-cognito-identitypool-supportedloginproviders
	SupportedLoginProviders pulumi.Input
}

func (IdentityPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityPoolArgs)(nil)).Elem()
}

type IdentityPoolInput interface {
	pulumi.Input

	ToIdentityPoolOutput() IdentityPoolOutput
	ToIdentityPoolOutputWithContext(ctx context.Context) IdentityPoolOutput
}

func (*IdentityPool) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPool)(nil))
}

func (i *IdentityPool) ToIdentityPoolOutput() IdentityPoolOutput {
	return i.ToIdentityPoolOutputWithContext(context.Background())
}

func (i *IdentityPool) ToIdentityPoolOutputWithContext(ctx context.Context) IdentityPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPoolOutput)
}

type IdentityPoolOutput struct{ *pulumi.OutputState }

func (IdentityPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPool)(nil))
}

func (o IdentityPoolOutput) ToIdentityPoolOutput() IdentityPoolOutput {
	return o
}

func (o IdentityPoolOutput) ToIdentityPoolOutputWithContext(ctx context.Context) IdentityPoolOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IdentityPoolOutput{})
}
