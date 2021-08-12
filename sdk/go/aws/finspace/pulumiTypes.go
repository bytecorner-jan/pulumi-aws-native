// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package finspace

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html
type EnvironmentFederationParameters struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-applicationcallbackurl
	ApplicationCallBackURL *string `pulumi:"applicationCallBackURL"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-attributemap
	AttributeMap interface{} `pulumi:"attributeMap"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationprovidername
	FederationProviderName *string `pulumi:"federationProviderName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationurn
	FederationURN *string `pulumi:"federationURN"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadatadocument
	SamlMetadataDocument *string `pulumi:"samlMetadataDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadataurl
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

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html
type EnvironmentFederationParametersArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-applicationcallbackurl
	ApplicationCallBackURL pulumi.StringPtrInput `pulumi:"applicationCallBackURL"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-attributemap
	AttributeMap pulumi.Input `pulumi:"attributeMap"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationprovidername
	FederationProviderName pulumi.StringPtrInput `pulumi:"federationProviderName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationurn
	FederationURN pulumi.StringPtrInput `pulumi:"federationURN"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadatadocument
	SamlMetadataDocument pulumi.StringPtrInput `pulumi:"samlMetadataDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadataurl
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

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html
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

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-applicationcallbackurl
func (o EnvironmentFederationParametersOutput) ApplicationCallBackURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.ApplicationCallBackURL }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-attributemap
func (o EnvironmentFederationParametersOutput) AttributeMap() pulumi.AnyOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) interface{} { return v.AttributeMap }).(pulumi.AnyOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationprovidername
func (o EnvironmentFederationParametersOutput) FederationProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.FederationProviderName }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationurn
func (o EnvironmentFederationParametersOutput) FederationURN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.FederationURN }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadatadocument
func (o EnvironmentFederationParametersOutput) SamlMetadataDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentFederationParameters) *string { return v.SamlMetadataDocument }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadataurl
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

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-applicationcallbackurl
func (o EnvironmentFederationParametersPtrOutput) ApplicationCallBackURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationCallBackURL
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-attributemap
func (o EnvironmentFederationParametersPtrOutput) AttributeMap() pulumi.AnyOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) interface{} {
		if v == nil {
			return nil
		}
		return v.AttributeMap
	}).(pulumi.AnyOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationprovidername
func (o EnvironmentFederationParametersPtrOutput) FederationProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.FederationProviderName
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-federationurn
func (o EnvironmentFederationParametersPtrOutput) FederationURN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.FederationURN
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadatadocument
func (o EnvironmentFederationParametersPtrOutput) SamlMetadataDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.SamlMetadataDocument
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-finspace-environment-federationparameters.html#cfn-finspace-environment-federationparameters-samlmetadataurl
func (o EnvironmentFederationParametersPtrOutput) SamlMetadataURL() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnvironmentFederationParameters) *string {
		if v == nil {
			return nil
		}
		return v.SamlMetadataURL
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentFederationParametersOutput{})
	pulumi.RegisterOutputType(EnvironmentFederationParametersPtrOutput{})
}
