// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package emr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Specifies whether the Studio authenticates users using single sign-on (SSO) or IAM. Amazon EMR Studio currently only supports SSO authentication.
type StudioAuthMode string

const (
	StudioAuthModeSso = StudioAuthMode("SSO")
	StudioAuthModeIam = StudioAuthMode("IAM")
)

func (StudioAuthMode) ElementType() reflect.Type {
	return reflect.TypeOf((*StudioAuthMode)(nil)).Elem()
}

func (e StudioAuthMode) ToStudioAuthModeOutput() StudioAuthModeOutput {
	return pulumi.ToOutput(e).(StudioAuthModeOutput)
}

func (e StudioAuthMode) ToStudioAuthModeOutputWithContext(ctx context.Context) StudioAuthModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StudioAuthModeOutput)
}

func (e StudioAuthMode) ToStudioAuthModePtrOutput() StudioAuthModePtrOutput {
	return e.ToStudioAuthModePtrOutputWithContext(context.Background())
}

func (e StudioAuthMode) ToStudioAuthModePtrOutputWithContext(ctx context.Context) StudioAuthModePtrOutput {
	return StudioAuthMode(e).ToStudioAuthModeOutputWithContext(ctx).ToStudioAuthModePtrOutputWithContext(ctx)
}

func (e StudioAuthMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StudioAuthMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StudioAuthMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StudioAuthMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StudioAuthModeOutput struct{ *pulumi.OutputState }

func (StudioAuthModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StudioAuthMode)(nil)).Elem()
}

func (o StudioAuthModeOutput) ToStudioAuthModeOutput() StudioAuthModeOutput {
	return o
}

func (o StudioAuthModeOutput) ToStudioAuthModeOutputWithContext(ctx context.Context) StudioAuthModeOutput {
	return o
}

func (o StudioAuthModeOutput) ToStudioAuthModePtrOutput() StudioAuthModePtrOutput {
	return o.ToStudioAuthModePtrOutputWithContext(context.Background())
}

func (o StudioAuthModeOutput) ToStudioAuthModePtrOutputWithContext(ctx context.Context) StudioAuthModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StudioAuthMode) *StudioAuthMode {
		return &v
	}).(StudioAuthModePtrOutput)
}

func (o StudioAuthModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StudioAuthModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StudioAuthMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StudioAuthModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StudioAuthModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StudioAuthMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StudioAuthModePtrOutput struct{ *pulumi.OutputState }

func (StudioAuthModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioAuthMode)(nil)).Elem()
}

func (o StudioAuthModePtrOutput) ToStudioAuthModePtrOutput() StudioAuthModePtrOutput {
	return o
}

func (o StudioAuthModePtrOutput) ToStudioAuthModePtrOutputWithContext(ctx context.Context) StudioAuthModePtrOutput {
	return o
}

func (o StudioAuthModePtrOutput) Elem() StudioAuthModeOutput {
	return o.ApplyT(func(v *StudioAuthMode) StudioAuthMode {
		if v != nil {
			return *v
		}
		var ret StudioAuthMode
		return ret
	}).(StudioAuthModeOutput)
}

func (o StudioAuthModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StudioAuthModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StudioAuthMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StudioAuthModeInput is an input type that accepts StudioAuthModeArgs and StudioAuthModeOutput values.
// You can construct a concrete instance of `StudioAuthModeInput` via:
//
//          StudioAuthModeArgs{...}
type StudioAuthModeInput interface {
	pulumi.Input

	ToStudioAuthModeOutput() StudioAuthModeOutput
	ToStudioAuthModeOutputWithContext(context.Context) StudioAuthModeOutput
}

var studioAuthModePtrType = reflect.TypeOf((**StudioAuthMode)(nil)).Elem()

type StudioAuthModePtrInput interface {
	pulumi.Input

	ToStudioAuthModePtrOutput() StudioAuthModePtrOutput
	ToStudioAuthModePtrOutputWithContext(context.Context) StudioAuthModePtrOutput
}

type studioAuthModePtr string

func StudioAuthModePtr(v string) StudioAuthModePtrInput {
	return (*studioAuthModePtr)(&v)
}

func (*studioAuthModePtr) ElementType() reflect.Type {
	return studioAuthModePtrType
}

func (in *studioAuthModePtr) ToStudioAuthModePtrOutput() StudioAuthModePtrOutput {
	return pulumi.ToOutput(in).(StudioAuthModePtrOutput)
}

func (in *studioAuthModePtr) ToStudioAuthModePtrOutputWithContext(ctx context.Context) StudioAuthModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StudioAuthModePtrOutput)
}

// Specifies whether the identity to map to the Studio is a user or a group.
type StudioSessionMappingIdentityType string

const (
	StudioSessionMappingIdentityTypeUser  = StudioSessionMappingIdentityType("USER")
	StudioSessionMappingIdentityTypeGroup = StudioSessionMappingIdentityType("GROUP")
)

func (StudioSessionMappingIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*StudioSessionMappingIdentityType)(nil)).Elem()
}

func (e StudioSessionMappingIdentityType) ToStudioSessionMappingIdentityTypeOutput() StudioSessionMappingIdentityTypeOutput {
	return pulumi.ToOutput(e).(StudioSessionMappingIdentityTypeOutput)
}

func (e StudioSessionMappingIdentityType) ToStudioSessionMappingIdentityTypeOutputWithContext(ctx context.Context) StudioSessionMappingIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StudioSessionMappingIdentityTypeOutput)
}

func (e StudioSessionMappingIdentityType) ToStudioSessionMappingIdentityTypePtrOutput() StudioSessionMappingIdentityTypePtrOutput {
	return e.ToStudioSessionMappingIdentityTypePtrOutputWithContext(context.Background())
}

func (e StudioSessionMappingIdentityType) ToStudioSessionMappingIdentityTypePtrOutputWithContext(ctx context.Context) StudioSessionMappingIdentityTypePtrOutput {
	return StudioSessionMappingIdentityType(e).ToStudioSessionMappingIdentityTypeOutputWithContext(ctx).ToStudioSessionMappingIdentityTypePtrOutputWithContext(ctx)
}

func (e StudioSessionMappingIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StudioSessionMappingIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StudioSessionMappingIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StudioSessionMappingIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StudioSessionMappingIdentityTypeOutput struct{ *pulumi.OutputState }

func (StudioSessionMappingIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StudioSessionMappingIdentityType)(nil)).Elem()
}

func (o StudioSessionMappingIdentityTypeOutput) ToStudioSessionMappingIdentityTypeOutput() StudioSessionMappingIdentityTypeOutput {
	return o
}

func (o StudioSessionMappingIdentityTypeOutput) ToStudioSessionMappingIdentityTypeOutputWithContext(ctx context.Context) StudioSessionMappingIdentityTypeOutput {
	return o
}

func (o StudioSessionMappingIdentityTypeOutput) ToStudioSessionMappingIdentityTypePtrOutput() StudioSessionMappingIdentityTypePtrOutput {
	return o.ToStudioSessionMappingIdentityTypePtrOutputWithContext(context.Background())
}

func (o StudioSessionMappingIdentityTypeOutput) ToStudioSessionMappingIdentityTypePtrOutputWithContext(ctx context.Context) StudioSessionMappingIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StudioSessionMappingIdentityType) *StudioSessionMappingIdentityType {
		return &v
	}).(StudioSessionMappingIdentityTypePtrOutput)
}

func (o StudioSessionMappingIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StudioSessionMappingIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StudioSessionMappingIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StudioSessionMappingIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StudioSessionMappingIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StudioSessionMappingIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StudioSessionMappingIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (StudioSessionMappingIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioSessionMappingIdentityType)(nil)).Elem()
}

func (o StudioSessionMappingIdentityTypePtrOutput) ToStudioSessionMappingIdentityTypePtrOutput() StudioSessionMappingIdentityTypePtrOutput {
	return o
}

func (o StudioSessionMappingIdentityTypePtrOutput) ToStudioSessionMappingIdentityTypePtrOutputWithContext(ctx context.Context) StudioSessionMappingIdentityTypePtrOutput {
	return o
}

func (o StudioSessionMappingIdentityTypePtrOutput) Elem() StudioSessionMappingIdentityTypeOutput {
	return o.ApplyT(func(v *StudioSessionMappingIdentityType) StudioSessionMappingIdentityType {
		if v != nil {
			return *v
		}
		var ret StudioSessionMappingIdentityType
		return ret
	}).(StudioSessionMappingIdentityTypeOutput)
}

func (o StudioSessionMappingIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StudioSessionMappingIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StudioSessionMappingIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// StudioSessionMappingIdentityTypeInput is an input type that accepts StudioSessionMappingIdentityTypeArgs and StudioSessionMappingIdentityTypeOutput values.
// You can construct a concrete instance of `StudioSessionMappingIdentityTypeInput` via:
//
//          StudioSessionMappingIdentityTypeArgs{...}
type StudioSessionMappingIdentityTypeInput interface {
	pulumi.Input

	ToStudioSessionMappingIdentityTypeOutput() StudioSessionMappingIdentityTypeOutput
	ToStudioSessionMappingIdentityTypeOutputWithContext(context.Context) StudioSessionMappingIdentityTypeOutput
}

var studioSessionMappingIdentityTypePtrType = reflect.TypeOf((**StudioSessionMappingIdentityType)(nil)).Elem()

type StudioSessionMappingIdentityTypePtrInput interface {
	pulumi.Input

	ToStudioSessionMappingIdentityTypePtrOutput() StudioSessionMappingIdentityTypePtrOutput
	ToStudioSessionMappingIdentityTypePtrOutputWithContext(context.Context) StudioSessionMappingIdentityTypePtrOutput
}

type studioSessionMappingIdentityTypePtr string

func StudioSessionMappingIdentityTypePtr(v string) StudioSessionMappingIdentityTypePtrInput {
	return (*studioSessionMappingIdentityTypePtr)(&v)
}

func (*studioSessionMappingIdentityTypePtr) ElementType() reflect.Type {
	return studioSessionMappingIdentityTypePtrType
}

func (in *studioSessionMappingIdentityTypePtr) ToStudioSessionMappingIdentityTypePtrOutput() StudioSessionMappingIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(StudioSessionMappingIdentityTypePtrOutput)
}

func (in *studioSessionMappingIdentityTypePtr) ToStudioSessionMappingIdentityTypePtrOutputWithContext(ctx context.Context) StudioSessionMappingIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StudioSessionMappingIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StudioAuthModeInput)(nil)).Elem(), StudioAuthMode("SSO"))
	pulumi.RegisterInputType(reflect.TypeOf((*StudioAuthModePtrInput)(nil)).Elem(), StudioAuthMode("SSO"))
	pulumi.RegisterInputType(reflect.TypeOf((*StudioSessionMappingIdentityTypeInput)(nil)).Elem(), StudioSessionMappingIdentityType("USER"))
	pulumi.RegisterInputType(reflect.TypeOf((*StudioSessionMappingIdentityTypePtrInput)(nil)).Elem(), StudioSessionMappingIdentityType("USER"))
	pulumi.RegisterOutputType(StudioAuthModeOutput{})
	pulumi.RegisterOutputType(StudioAuthModePtrOutput{})
	pulumi.RegisterOutputType(StudioSessionMappingIdentityTypeOutput{})
	pulumi.RegisterOutputType(StudioSessionMappingIdentityTypePtrOutput{})
}
