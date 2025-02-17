// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package finspace

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Additional parameters to identify Federation mode
type EnvironmentFederationParameters struct {
	// SAML metadata URL to link with the Environment
	ApplicationCallBackURL *string `pulumi:"applicationCallBackURL"`
	// Attribute map for SAML configuration
	AttributeMap interface{} `pulumi:"attributeMap"`
	// Federation provider name to link with the Environment
	FederationProviderName *string `pulumi:"federationProviderName"`
	// SAML metadata URL to link with the Environment
	FederationURN *string `pulumi:"federationURN"`
	// SAML metadata document to link the federation provider to the Environment
	SamlMetadataDocument *string `pulumi:"samlMetadataDocument"`
	// SAML metadata URL to link with the Environment
	SamlMetadataURL *string `pulumi:"samlMetadataURL"`
}

// EnvironmentFederationParametersInput is an input type that accepts EnvironmentFederationParametersArgs and EnvironmentFederationParametersOutput values.
// You can construct a concrete instance of `EnvironmentFederationParametersInput` via:
//
//          EnvironmentFederationParametersArgs{...}
type EnvironmentFederationParametersInput interface {
	pulumi.Input

	ToEnvironmentFederationParametersOutput() EnvironmentFederationParametersOutput
	ToEnvironmentFederationParametersOutputWithContext(context.Context) EnvironmentFederationParametersOutput
}

// Additional parameters to identify Federation mode
type EnvironmentFederationParametersArgs struct {
	// SAML metadata URL to link with the Environment
	ApplicationCallBackURL pulumi.StringPtrInput `pulumi:"applicationCallBackURL"`
	// Attribute map for SAML configuration
	AttributeMap pulumi.Input `pulumi:"attributeMap"`
	// Federation provider name to link with the Environment
	FederationProviderName pulumi.StringPtrInput `pulumi:"federationProviderName"`
	// SAML metadata URL to link with the Environment
	FederationURN pulumi.StringPtrInput `pulumi:"federationURN"`
	// SAML metadata document to link the federation provider to the Environment
	SamlMetadataDocument pulumi.StringPtrInput `pulumi:"samlMetadataDocument"`
	// SAML metadata URL to link with the Environment
	SamlMetadataURL pulumi.StringPtrInput `pulumi:"samlMetadataURL"`
}

func (EnvironmentFederationParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationParameters)(nil)).Elem()
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersOutput() EnvironmentFederationParametersOutput {
	return i.ToEnvironmentFederationParametersOutputWithContext(context.Background())
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersOutputWithContext(ctx context.Context) EnvironmentFederationParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersOutput)
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return i.ToEnvironmentFederationParametersPtrOutputWithContext(context.Background())
}

func (i EnvironmentFederationParametersArgs) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersOutput).ToEnvironmentFederationParametersPtrOutputWithContext(ctx)
}

// EnvironmentFederationParametersPtrInput is an input type that accepts EnvironmentFederationParametersArgs, EnvironmentFederationParametersPtr and EnvironmentFederationParametersPtrOutput values.
// You can construct a concrete instance of `EnvironmentFederationParametersPtrInput` via:
//
//          EnvironmentFederationParametersArgs{...}
//
//  or:
//
//          nil
type EnvironmentFederationParametersPtrInput interface {
	pulumi.Input

	ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput
	ToEnvironmentFederationParametersPtrOutputWithContext(context.Context) EnvironmentFederationParametersPtrOutput
}

type environmentFederationParametersPtrType EnvironmentFederationParametersArgs

func EnvironmentFederationParametersPtr(v *EnvironmentFederationParametersArgs) EnvironmentFederationParametersPtrInput {
	return (*environmentFederationParametersPtrType)(v)
}

func (*environmentFederationParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentFederationParameters)(nil)).Elem()
}

func (i *environmentFederationParametersPtrType) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return i.ToEnvironmentFederationParametersPtrOutputWithContext(context.Background())
}

func (i *environmentFederationParametersPtrType) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentFederationParametersPtrOutput)
}

// Additional parameters to identify Federation mode
type EnvironmentFederationParametersOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationParameters)(nil)).Elem()
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersOutput() EnvironmentFederationParametersOutput {
	return o
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersOutputWithContext(ctx context.Context) EnvironmentFederationParametersOutput {
	return o
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return o.ToEnvironmentFederationParametersPtrOutputWithContext(context.Background())
}

func (o EnvironmentFederationParametersOutput) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentFederationParameters) *EnvironmentFederationParameters {
		return &v
	}).(EnvironmentFederationParametersPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersOutput) ApplicationCallBackURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.ApplicationCallBackURL }).(pulumi.StringPtrOutput)
}

// Attribute map for SAML configuration
func (o EnvironmentFederationParametersOutput) AttributeMap() pulumi.AnyOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) interface{} { return v.AttributeMap }).(pulumi.AnyOutput)
}

// Federation provider name to link with the Environment
func (o EnvironmentFederationParametersOutput) FederationProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.FederationProviderName }).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersOutput) FederationURN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.FederationURN }).(pulumi.StringPtrOutput)
}

// SAML metadata document to link the federation provider to the Environment
func (o EnvironmentFederationParametersOutput) SamlMetadataDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.SamlMetadataDocument }).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersOutput) SamlMetadataURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.SamlMetadataURL }).(pulumi.StringPtrOutput)
}

type EnvironmentFederationParametersPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentFederationParameters)(nil)).Elem()
}

func (o EnvironmentFederationParametersPtrOutput) ToEnvironmentFederationParametersPtrOutput() EnvironmentFederationParametersPtrOutput {
	return o
}

func (o EnvironmentFederationParametersPtrOutput) ToEnvironmentFederationParametersPtrOutputWithContext(ctx context.Context) EnvironmentFederationParametersPtrOutput {
	return o
}

func (o EnvironmentFederationParametersPtrOutput) Elem() EnvironmentFederationParametersOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) EnvironmentFederationParameters {
		if v != nil {
			return *v
		}
		var ret EnvironmentFederationParameters
		return ret
	}).(EnvironmentFederationParametersOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) ApplicationCallBackURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationCallBackURL
	}).(pulumi.StringPtrOutput)
}

// Attribute map for SAML configuration
func (o EnvironmentFederationParametersPtrOutput) AttributeMap() pulumi.AnyOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) interface{} {
		if v == nil {
			return nil
		}
		return v.AttributeMap
	}).(pulumi.AnyOutput)
}

// Federation provider name to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) FederationProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.FederationProviderName
	}).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) FederationURN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.FederationURN
	}).(pulumi.StringPtrOutput)
}

// SAML metadata document to link the federation provider to the Environment
func (o EnvironmentFederationParametersPtrOutput) SamlMetadataDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.SamlMetadataDocument
	}).(pulumi.StringPtrOutput)
}

// SAML metadata URL to link with the Environment
func (o EnvironmentFederationParametersPtrOutput) SamlMetadataURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.SamlMetadataURL
	}).(pulumi.StringPtrOutput)
}

// Parameters of the first Superuser for the FinSpace Environment
type EnvironmentSuperuserParameters struct {
	// Email address
	EmailAddress *string `pulumi:"emailAddress"`
	// First name
	FirstName *string `pulumi:"firstName"`
	// Last name
	LastName *string `pulumi:"lastName"`
}

// EnvironmentSuperuserParametersInput is an input type that accepts EnvironmentSuperuserParametersArgs and EnvironmentSuperuserParametersOutput values.
// You can construct a concrete instance of `EnvironmentSuperuserParametersInput` via:
//
//          EnvironmentSuperuserParametersArgs{...}
type EnvironmentSuperuserParametersInput interface {
	pulumi.Input

	ToEnvironmentSuperuserParametersOutput() EnvironmentSuperuserParametersOutput
	ToEnvironmentSuperuserParametersOutputWithContext(context.Context) EnvironmentSuperuserParametersOutput
}

// Parameters of the first Superuser for the FinSpace Environment
type EnvironmentSuperuserParametersArgs struct {
	// Email address
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	// First name
	FirstName pulumi.StringPtrInput `pulumi:"firstName"`
	// Last name
	LastName pulumi.StringPtrInput `pulumi:"lastName"`
}

func (EnvironmentSuperuserParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSuperuserParameters)(nil)).Elem()
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersOutput() EnvironmentSuperuserParametersOutput {
	return i.ToEnvironmentSuperuserParametersOutputWithContext(context.Background())
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSuperuserParametersOutput)
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return i.ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Background())
}

func (i EnvironmentSuperuserParametersArgs) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSuperuserParametersOutput).ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx)
}

// EnvironmentSuperuserParametersPtrInput is an input type that accepts EnvironmentSuperuserParametersArgs, EnvironmentSuperuserParametersPtr and EnvironmentSuperuserParametersPtrOutput values.
// You can construct a concrete instance of `EnvironmentSuperuserParametersPtrInput` via:
//
//          EnvironmentSuperuserParametersArgs{...}
//
//  or:
//
//          nil
type EnvironmentSuperuserParametersPtrInput interface {
	pulumi.Input

	ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput
	ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Context) EnvironmentSuperuserParametersPtrOutput
}

type environmentSuperuserParametersPtrType EnvironmentSuperuserParametersArgs

func EnvironmentSuperuserParametersPtr(v *EnvironmentSuperuserParametersArgs) EnvironmentSuperuserParametersPtrInput {
	return (*environmentSuperuserParametersPtrType)(v)
}

func (*environmentSuperuserParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSuperuserParameters)(nil)).Elem()
}

func (i *environmentSuperuserParametersPtrType) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return i.ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Background())
}

func (i *environmentSuperuserParametersPtrType) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentSuperuserParametersPtrOutput)
}

// Parameters of the first Superuser for the FinSpace Environment
type EnvironmentSuperuserParametersOutput struct{ *pulumi.OutputState }

func (EnvironmentSuperuserParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSuperuserParameters)(nil)).Elem()
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersOutput() EnvironmentSuperuserParametersOutput {
	return o
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersOutput {
	return o
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return o.ToEnvironmentSuperuserParametersPtrOutputWithContext(context.Background())
}

func (o EnvironmentSuperuserParametersOutput) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentSuperuserParameters) *EnvironmentSuperuserParameters {
		return &v
	}).(EnvironmentSuperuserParametersPtrOutput)
}

// Email address
func (o EnvironmentSuperuserParametersOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSuperuserParameters) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

// First name
func (o EnvironmentSuperuserParametersOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSuperuserParameters) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

// Last name
func (o EnvironmentSuperuserParametersOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSuperuserParameters) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type EnvironmentSuperuserParametersPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentSuperuserParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentSuperuserParameters)(nil)).Elem()
}

func (o EnvironmentSuperuserParametersPtrOutput) ToEnvironmentSuperuserParametersPtrOutput() EnvironmentSuperuserParametersPtrOutput {
	return o
}

func (o EnvironmentSuperuserParametersPtrOutput) ToEnvironmentSuperuserParametersPtrOutputWithContext(ctx context.Context) EnvironmentSuperuserParametersPtrOutput {
	return o
}

func (o EnvironmentSuperuserParametersPtrOutput) Elem() EnvironmentSuperuserParametersOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) EnvironmentSuperuserParameters {
		if v != nil {
			return *v
		}
		var ret EnvironmentSuperuserParameters
		return ret
	}).(EnvironmentSuperuserParametersOutput)
}

// Email address
func (o EnvironmentSuperuserParametersPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

// First name
func (o EnvironmentSuperuserParametersPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

// Last name
func (o EnvironmentSuperuserParametersPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentSuperuserParameters) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationParametersInput)(nil)).Elem(), EnvironmentFederationParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationParametersPtrInput)(nil)).Elem(), EnvironmentFederationParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentSuperuserParametersInput)(nil)).Elem(), EnvironmentSuperuserParametersArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentSuperuserParametersPtrInput)(nil)).Elem(), EnvironmentSuperuserParametersArgs{})
	pulumi.RegisterOutputType(EnvironmentFederationParametersOutput{})
	pulumi.RegisterOutputType(EnvironmentFederationParametersPtrOutput{})
	pulumi.RegisterOutputType(EnvironmentSuperuserParametersOutput{})
	pulumi.RegisterOutputType(EnvironmentSuperuserParametersPtrOutput{})
}
