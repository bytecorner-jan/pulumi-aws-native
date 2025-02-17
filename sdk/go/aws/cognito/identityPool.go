// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::IdentityPool
//
// Deprecated: IdentityPool is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type IdentityPool struct {
	pulumi.CustomResourceState

	AllowClassicFlow               pulumi.BoolPtrOutput                           `pulumi:"allowClassicFlow"`
	AllowUnauthenticatedIdentities pulumi.BoolOutput                              `pulumi:"allowUnauthenticatedIdentities"`
	CognitoEvents                  pulumi.AnyOutput                               `pulumi:"cognitoEvents"`
	CognitoIdentityProviders       IdentityPoolCognitoIdentityProviderArrayOutput `pulumi:"cognitoIdentityProviders"`
	CognitoStreams                 IdentityPoolCognitoStreamsPtrOutput            `pulumi:"cognitoStreams"`
	DeveloperProviderName          pulumi.StringPtrOutput                         `pulumi:"developerProviderName"`
	IdentityPoolName               pulumi.StringPtrOutput                         `pulumi:"identityPoolName"`
	Name                           pulumi.StringOutput                            `pulumi:"name"`
	OpenIdConnectProviderARNs      pulumi.StringArrayOutput                       `pulumi:"openIdConnectProviderARNs"`
	PushSync                       IdentityPoolPushSyncPtrOutput                  `pulumi:"pushSync"`
	SamlProviderARNs               pulumi.StringArrayOutput                       `pulumi:"samlProviderARNs"`
	SupportedLoginProviders        pulumi.AnyOutput                               `pulumi:"supportedLoginProviders"`
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
	AllowClassicFlow               *bool                                 `pulumi:"allowClassicFlow"`
	AllowUnauthenticatedIdentities bool                                  `pulumi:"allowUnauthenticatedIdentities"`
	CognitoEvents                  interface{}                           `pulumi:"cognitoEvents"`
	CognitoIdentityProviders       []IdentityPoolCognitoIdentityProvider `pulumi:"cognitoIdentityProviders"`
	CognitoStreams                 *IdentityPoolCognitoStreams           `pulumi:"cognitoStreams"`
	DeveloperProviderName          *string                               `pulumi:"developerProviderName"`
	IdentityPoolName               *string                               `pulumi:"identityPoolName"`
	OpenIdConnectProviderARNs      []string                              `pulumi:"openIdConnectProviderARNs"`
	PushSync                       *IdentityPoolPushSync                 `pulumi:"pushSync"`
	SamlProviderARNs               []string                              `pulumi:"samlProviderARNs"`
	SupportedLoginProviders        interface{}                           `pulumi:"supportedLoginProviders"`
}

// The set of arguments for constructing a IdentityPool resource.
type IdentityPoolArgs struct {
	AllowClassicFlow               pulumi.BoolPtrInput
	AllowUnauthenticatedIdentities pulumi.BoolInput
	CognitoEvents                  pulumi.Input
	CognitoIdentityProviders       IdentityPoolCognitoIdentityProviderArrayInput
	CognitoStreams                 IdentityPoolCognitoStreamsPtrInput
	DeveloperProviderName          pulumi.StringPtrInput
	IdentityPoolName               pulumi.StringPtrInput
	OpenIdConnectProviderARNs      pulumi.StringArrayInput
	PushSync                       IdentityPoolPushSyncPtrInput
	SamlProviderARNs               pulumi.StringArrayInput
	SupportedLoginProviders        pulumi.Input
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
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPoolInput)(nil)).Elem(), &IdentityPool{})
	pulumi.RegisterOutputType(IdentityPoolOutput{})
}
