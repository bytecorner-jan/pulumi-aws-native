// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appconfig

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html
type ApplicationTags struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html#cfn-appconfig-application-tags-key
	Key *string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html#cfn-appconfig-application-tags-value
	Value *string `pulumi:"value"`
}

// ApplicationTagsInput is an input type that accepts ApplicationTagsArgs and ApplicationTagsOutput values.
// You can construct a concrete instance of `ApplicationTagsInput` via:
//
//          ApplicationTagsArgs{...}
type ApplicationTagsInput interface {
	pulumi.Input

	ToApplicationTagsOutput() ApplicationTagsOutput
	ToApplicationTagsOutputWithContext(context.Context) ApplicationTagsOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html
type ApplicationTagsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html#cfn-appconfig-application-tags-key
	Key pulumi.StringPtrInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html#cfn-appconfig-application-tags-value
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ApplicationTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTags)(nil)).Elem()
}

func (i ApplicationTagsArgs) ToApplicationTagsOutput() ApplicationTagsOutput {
	return i.ToApplicationTagsOutputWithContext(context.Background())
}

func (i ApplicationTagsArgs) ToApplicationTagsOutputWithContext(ctx context.Context) ApplicationTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTagsOutput)
}

// ApplicationTagsArrayInput is an input type that accepts ApplicationTagsArray and ApplicationTagsArrayOutput values.
// You can construct a concrete instance of `ApplicationTagsArrayInput` via:
//
//          ApplicationTagsArray{ ApplicationTagsArgs{...} }
type ApplicationTagsArrayInput interface {
	pulumi.Input

	ToApplicationTagsArrayOutput() ApplicationTagsArrayOutput
	ToApplicationTagsArrayOutputWithContext(context.Context) ApplicationTagsArrayOutput
}

type ApplicationTagsArray []ApplicationTagsInput

func (ApplicationTagsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationTags)(nil)).Elem()
}

func (i ApplicationTagsArray) ToApplicationTagsArrayOutput() ApplicationTagsArrayOutput {
	return i.ToApplicationTagsArrayOutputWithContext(context.Background())
}

func (i ApplicationTagsArray) ToApplicationTagsArrayOutputWithContext(ctx context.Context) ApplicationTagsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationTagsArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html
type ApplicationTagsOutput struct{ *pulumi.OutputState }

func (ApplicationTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationTags)(nil)).Elem()
}

func (o ApplicationTagsOutput) ToApplicationTagsOutput() ApplicationTagsOutput {
	return o
}

func (o ApplicationTagsOutput) ToApplicationTagsOutputWithContext(ctx context.Context) ApplicationTagsOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html#cfn-appconfig-application-tags-key
func (o ApplicationTagsOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationTags) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-application-tags.html#cfn-appconfig-application-tags-value
func (o ApplicationTagsOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApplicationTags) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ApplicationTagsArrayOutput struct{ *pulumi.OutputState }

func (ApplicationTagsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApplicationTags)(nil)).Elem()
}

func (o ApplicationTagsArrayOutput) ToApplicationTagsArrayOutput() ApplicationTagsArrayOutput {
	return o
}

func (o ApplicationTagsArrayOutput) ToApplicationTagsArrayOutputWithContext(ctx context.Context) ApplicationTagsArrayOutput {
	return o
}

func (o ApplicationTagsArrayOutput) Index(i pulumi.IntInput) ApplicationTagsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApplicationTags {
		return vs[0].([]ApplicationTags)[vs[1].(int)]
	}).(ApplicationTagsOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html
type ConfigurationProfileTags struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-key
	Key *string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-value
	Value *string `pulumi:"value"`
}

// ConfigurationProfileTagsInput is an input type that accepts ConfigurationProfileTagsArgs and ConfigurationProfileTagsOutput values.
// You can construct a concrete instance of `ConfigurationProfileTagsInput` via:
//
//          ConfigurationProfileTagsArgs{...}
type ConfigurationProfileTagsInput interface {
	pulumi.Input

	ToConfigurationProfileTagsOutput() ConfigurationProfileTagsOutput
	ToConfigurationProfileTagsOutputWithContext(context.Context) ConfigurationProfileTagsOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html
type ConfigurationProfileTagsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-key
	Key pulumi.StringPtrInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-value
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ConfigurationProfileTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileTags)(nil)).Elem()
}

func (i ConfigurationProfileTagsArgs) ToConfigurationProfileTagsOutput() ConfigurationProfileTagsOutput {
	return i.ToConfigurationProfileTagsOutputWithContext(context.Background())
}

func (i ConfigurationProfileTagsArgs) ToConfigurationProfileTagsOutputWithContext(ctx context.Context) ConfigurationProfileTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileTagsOutput)
}

// ConfigurationProfileTagsArrayInput is an input type that accepts ConfigurationProfileTagsArray and ConfigurationProfileTagsArrayOutput values.
// You can construct a concrete instance of `ConfigurationProfileTagsArrayInput` via:
//
//          ConfigurationProfileTagsArray{ ConfigurationProfileTagsArgs{...} }
type ConfigurationProfileTagsArrayInput interface {
	pulumi.Input

	ToConfigurationProfileTagsArrayOutput() ConfigurationProfileTagsArrayOutput
	ToConfigurationProfileTagsArrayOutputWithContext(context.Context) ConfigurationProfileTagsArrayOutput
}

type ConfigurationProfileTagsArray []ConfigurationProfileTagsInput

func (ConfigurationProfileTagsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationProfileTags)(nil)).Elem()
}

func (i ConfigurationProfileTagsArray) ToConfigurationProfileTagsArrayOutput() ConfigurationProfileTagsArrayOutput {
	return i.ToConfigurationProfileTagsArrayOutputWithContext(context.Background())
}

func (i ConfigurationProfileTagsArray) ToConfigurationProfileTagsArrayOutputWithContext(ctx context.Context) ConfigurationProfileTagsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileTagsArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html
type ConfigurationProfileTagsOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileTags)(nil)).Elem()
}

func (o ConfigurationProfileTagsOutput) ToConfigurationProfileTagsOutput() ConfigurationProfileTagsOutput {
	return o
}

func (o ConfigurationProfileTagsOutput) ToConfigurationProfileTagsOutputWithContext(ctx context.Context) ConfigurationProfileTagsOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-key
func (o ConfigurationProfileTagsOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileTags) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-tags.html#cfn-appconfig-configurationprofile-tags-value
func (o ConfigurationProfileTagsOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileTags) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileTagsArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileTagsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationProfileTags)(nil)).Elem()
}

func (o ConfigurationProfileTagsArrayOutput) ToConfigurationProfileTagsArrayOutput() ConfigurationProfileTagsArrayOutput {
	return o
}

func (o ConfigurationProfileTagsArrayOutput) ToConfigurationProfileTagsArrayOutputWithContext(ctx context.Context) ConfigurationProfileTagsArrayOutput {
	return o
}

func (o ConfigurationProfileTagsArrayOutput) Index(i pulumi.IntInput) ConfigurationProfileTagsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationProfileTags {
		return vs[0].([]ConfigurationProfileTags)[vs[1].(int)]
	}).(ConfigurationProfileTagsOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html
type ConfigurationProfileValidators struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-content
	Content *string `pulumi:"content"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-type
	Type *string `pulumi:"type"`
}

// ConfigurationProfileValidatorsInput is an input type that accepts ConfigurationProfileValidatorsArgs and ConfigurationProfileValidatorsOutput values.
// You can construct a concrete instance of `ConfigurationProfileValidatorsInput` via:
//
//          ConfigurationProfileValidatorsArgs{...}
type ConfigurationProfileValidatorsInput interface {
	pulumi.Input

	ToConfigurationProfileValidatorsOutput() ConfigurationProfileValidatorsOutput
	ToConfigurationProfileValidatorsOutputWithContext(context.Context) ConfigurationProfileValidatorsOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html
type ConfigurationProfileValidatorsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-content
	Content pulumi.StringPtrInput `pulumi:"content"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-type
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ConfigurationProfileValidatorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileValidators)(nil)).Elem()
}

func (i ConfigurationProfileValidatorsArgs) ToConfigurationProfileValidatorsOutput() ConfigurationProfileValidatorsOutput {
	return i.ToConfigurationProfileValidatorsOutputWithContext(context.Background())
}

func (i ConfigurationProfileValidatorsArgs) ToConfigurationProfileValidatorsOutputWithContext(ctx context.Context) ConfigurationProfileValidatorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileValidatorsOutput)
}

// ConfigurationProfileValidatorsArrayInput is an input type that accepts ConfigurationProfileValidatorsArray and ConfigurationProfileValidatorsArrayOutput values.
// You can construct a concrete instance of `ConfigurationProfileValidatorsArrayInput` via:
//
//          ConfigurationProfileValidatorsArray{ ConfigurationProfileValidatorsArgs{...} }
type ConfigurationProfileValidatorsArrayInput interface {
	pulumi.Input

	ToConfigurationProfileValidatorsArrayOutput() ConfigurationProfileValidatorsArrayOutput
	ToConfigurationProfileValidatorsArrayOutputWithContext(context.Context) ConfigurationProfileValidatorsArrayOutput
}

type ConfigurationProfileValidatorsArray []ConfigurationProfileValidatorsInput

func (ConfigurationProfileValidatorsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationProfileValidators)(nil)).Elem()
}

func (i ConfigurationProfileValidatorsArray) ToConfigurationProfileValidatorsArrayOutput() ConfigurationProfileValidatorsArrayOutput {
	return i.ToConfigurationProfileValidatorsArrayOutputWithContext(context.Background())
}

func (i ConfigurationProfileValidatorsArray) ToConfigurationProfileValidatorsArrayOutputWithContext(ctx context.Context) ConfigurationProfileValidatorsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationProfileValidatorsArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html
type ConfigurationProfileValidatorsOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileValidatorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationProfileValidators)(nil)).Elem()
}

func (o ConfigurationProfileValidatorsOutput) ToConfigurationProfileValidatorsOutput() ConfigurationProfileValidatorsOutput {
	return o
}

func (o ConfigurationProfileValidatorsOutput) ToConfigurationProfileValidatorsOutputWithContext(ctx context.Context) ConfigurationProfileValidatorsOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-content
func (o ConfigurationProfileValidatorsOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileValidators) *string { return v.Content }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-configurationprofile-validators.html#cfn-appconfig-configurationprofile-validators-type
func (o ConfigurationProfileValidatorsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationProfileValidators) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConfigurationProfileValidatorsArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationProfileValidatorsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationProfileValidators)(nil)).Elem()
}

func (o ConfigurationProfileValidatorsArrayOutput) ToConfigurationProfileValidatorsArrayOutput() ConfigurationProfileValidatorsArrayOutput {
	return o
}

func (o ConfigurationProfileValidatorsArrayOutput) ToConfigurationProfileValidatorsArrayOutputWithContext(ctx context.Context) ConfigurationProfileValidatorsArrayOutput {
	return o
}

func (o ConfigurationProfileValidatorsArrayOutput) Index(i pulumi.IntInput) ConfigurationProfileValidatorsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationProfileValidators {
		return vs[0].([]ConfigurationProfileValidators)[vs[1].(int)]
	}).(ConfigurationProfileValidatorsOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html
type DeploymentStrategyTags struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html#cfn-appconfig-deploymentstrategy-tags-key
	Key *string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html#cfn-appconfig-deploymentstrategy-tags-value
	Value *string `pulumi:"value"`
}

// DeploymentStrategyTagsInput is an input type that accepts DeploymentStrategyTagsArgs and DeploymentStrategyTagsOutput values.
// You can construct a concrete instance of `DeploymentStrategyTagsInput` via:
//
//          DeploymentStrategyTagsArgs{...}
type DeploymentStrategyTagsInput interface {
	pulumi.Input

	ToDeploymentStrategyTagsOutput() DeploymentStrategyTagsOutput
	ToDeploymentStrategyTagsOutputWithContext(context.Context) DeploymentStrategyTagsOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html
type DeploymentStrategyTagsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html#cfn-appconfig-deploymentstrategy-tags-key
	Key pulumi.StringPtrInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html#cfn-appconfig-deploymentstrategy-tags-value
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (DeploymentStrategyTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentStrategyTags)(nil)).Elem()
}

func (i DeploymentStrategyTagsArgs) ToDeploymentStrategyTagsOutput() DeploymentStrategyTagsOutput {
	return i.ToDeploymentStrategyTagsOutputWithContext(context.Background())
}

func (i DeploymentStrategyTagsArgs) ToDeploymentStrategyTagsOutputWithContext(ctx context.Context) DeploymentStrategyTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentStrategyTagsOutput)
}

// DeploymentStrategyTagsArrayInput is an input type that accepts DeploymentStrategyTagsArray and DeploymentStrategyTagsArrayOutput values.
// You can construct a concrete instance of `DeploymentStrategyTagsArrayInput` via:
//
//          DeploymentStrategyTagsArray{ DeploymentStrategyTagsArgs{...} }
type DeploymentStrategyTagsArrayInput interface {
	pulumi.Input

	ToDeploymentStrategyTagsArrayOutput() DeploymentStrategyTagsArrayOutput
	ToDeploymentStrategyTagsArrayOutputWithContext(context.Context) DeploymentStrategyTagsArrayOutput
}

type DeploymentStrategyTagsArray []DeploymentStrategyTagsInput

func (DeploymentStrategyTagsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentStrategyTags)(nil)).Elem()
}

func (i DeploymentStrategyTagsArray) ToDeploymentStrategyTagsArrayOutput() DeploymentStrategyTagsArrayOutput {
	return i.ToDeploymentStrategyTagsArrayOutputWithContext(context.Background())
}

func (i DeploymentStrategyTagsArray) ToDeploymentStrategyTagsArrayOutputWithContext(ctx context.Context) DeploymentStrategyTagsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentStrategyTagsArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html
type DeploymentStrategyTagsOutput struct{ *pulumi.OutputState }

func (DeploymentStrategyTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentStrategyTags)(nil)).Elem()
}

func (o DeploymentStrategyTagsOutput) ToDeploymentStrategyTagsOutput() DeploymentStrategyTagsOutput {
	return o
}

func (o DeploymentStrategyTagsOutput) ToDeploymentStrategyTagsOutputWithContext(ctx context.Context) DeploymentStrategyTagsOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html#cfn-appconfig-deploymentstrategy-tags-key
func (o DeploymentStrategyTagsOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentStrategyTags) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deploymentstrategy-tags.html#cfn-appconfig-deploymentstrategy-tags-value
func (o DeploymentStrategyTagsOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentStrategyTags) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DeploymentStrategyTagsArrayOutput struct{ *pulumi.OutputState }

func (DeploymentStrategyTagsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentStrategyTags)(nil)).Elem()
}

func (o DeploymentStrategyTagsArrayOutput) ToDeploymentStrategyTagsArrayOutput() DeploymentStrategyTagsArrayOutput {
	return o
}

func (o DeploymentStrategyTagsArrayOutput) ToDeploymentStrategyTagsArrayOutputWithContext(ctx context.Context) DeploymentStrategyTagsArrayOutput {
	return o
}

func (o DeploymentStrategyTagsArrayOutput) Index(i pulumi.IntInput) DeploymentStrategyTagsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeploymentStrategyTags {
		return vs[0].([]DeploymentStrategyTags)[vs[1].(int)]
	}).(DeploymentStrategyTagsOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html
type DeploymentTags struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-key
	Key *string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-value
	Value *string `pulumi:"value"`
}

// DeploymentTagsInput is an input type that accepts DeploymentTagsArgs and DeploymentTagsOutput values.
// You can construct a concrete instance of `DeploymentTagsInput` via:
//
//          DeploymentTagsArgs{...}
type DeploymentTagsInput interface {
	pulumi.Input

	ToDeploymentTagsOutput() DeploymentTagsOutput
	ToDeploymentTagsOutputWithContext(context.Context) DeploymentTagsOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html
type DeploymentTagsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-key
	Key pulumi.StringPtrInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-value
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (DeploymentTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTags)(nil)).Elem()
}

func (i DeploymentTagsArgs) ToDeploymentTagsOutput() DeploymentTagsOutput {
	return i.ToDeploymentTagsOutputWithContext(context.Background())
}

func (i DeploymentTagsArgs) ToDeploymentTagsOutputWithContext(ctx context.Context) DeploymentTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTagsOutput)
}

// DeploymentTagsArrayInput is an input type that accepts DeploymentTagsArray and DeploymentTagsArrayOutput values.
// You can construct a concrete instance of `DeploymentTagsArrayInput` via:
//
//          DeploymentTagsArray{ DeploymentTagsArgs{...} }
type DeploymentTagsArrayInput interface {
	pulumi.Input

	ToDeploymentTagsArrayOutput() DeploymentTagsArrayOutput
	ToDeploymentTagsArrayOutputWithContext(context.Context) DeploymentTagsArrayOutput
}

type DeploymentTagsArray []DeploymentTagsInput

func (DeploymentTagsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentTags)(nil)).Elem()
}

func (i DeploymentTagsArray) ToDeploymentTagsArrayOutput() DeploymentTagsArrayOutput {
	return i.ToDeploymentTagsArrayOutputWithContext(context.Background())
}

func (i DeploymentTagsArray) ToDeploymentTagsArrayOutputWithContext(ctx context.Context) DeploymentTagsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentTagsArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html
type DeploymentTagsOutput struct{ *pulumi.OutputState }

func (DeploymentTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentTags)(nil)).Elem()
}

func (o DeploymentTagsOutput) ToDeploymentTagsOutput() DeploymentTagsOutput {
	return o
}

func (o DeploymentTagsOutput) ToDeploymentTagsOutputWithContext(ctx context.Context) DeploymentTagsOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-key
func (o DeploymentTagsOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentTags) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-deployment-tags.html#cfn-appconfig-deployment-tags-value
func (o DeploymentTagsOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentTags) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DeploymentTagsArrayOutput struct{ *pulumi.OutputState }

func (DeploymentTagsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeploymentTags)(nil)).Elem()
}

func (o DeploymentTagsArrayOutput) ToDeploymentTagsArrayOutput() DeploymentTagsArrayOutput {
	return o
}

func (o DeploymentTagsArrayOutput) ToDeploymentTagsArrayOutputWithContext(ctx context.Context) DeploymentTagsArrayOutput {
	return o
}

func (o DeploymentTagsArrayOutput) Index(i pulumi.IntInput) DeploymentTagsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeploymentTags {
		return vs[0].([]DeploymentTags)[vs[1].(int)]
	}).(DeploymentTagsOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html
type EnvironmentMonitors struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmarn
	AlarmArn *string `pulumi:"alarmArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmrolearn
	AlarmRoleArn *string `pulumi:"alarmRoleArn"`
}

// EnvironmentMonitorsInput is an input type that accepts EnvironmentMonitorsArgs and EnvironmentMonitorsOutput values.
// You can construct a concrete instance of `EnvironmentMonitorsInput` via:
//
//          EnvironmentMonitorsArgs{...}
type EnvironmentMonitorsInput interface {
	pulumi.Input

	ToEnvironmentMonitorsOutput() EnvironmentMonitorsOutput
	ToEnvironmentMonitorsOutputWithContext(context.Context) EnvironmentMonitorsOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html
type EnvironmentMonitorsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmarn
	AlarmArn pulumi.StringPtrInput `pulumi:"alarmArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmrolearn
	AlarmRoleArn pulumi.StringPtrInput `pulumi:"alarmRoleArn"`
}

func (EnvironmentMonitorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentMonitors)(nil)).Elem()
}

func (i EnvironmentMonitorsArgs) ToEnvironmentMonitorsOutput() EnvironmentMonitorsOutput {
	return i.ToEnvironmentMonitorsOutputWithContext(context.Background())
}

func (i EnvironmentMonitorsArgs) ToEnvironmentMonitorsOutputWithContext(ctx context.Context) EnvironmentMonitorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMonitorsOutput)
}

// EnvironmentMonitorsArrayInput is an input type that accepts EnvironmentMonitorsArray and EnvironmentMonitorsArrayOutput values.
// You can construct a concrete instance of `EnvironmentMonitorsArrayInput` via:
//
//          EnvironmentMonitorsArray{ EnvironmentMonitorsArgs{...} }
type EnvironmentMonitorsArrayInput interface {
	pulumi.Input

	ToEnvironmentMonitorsArrayOutput() EnvironmentMonitorsArrayOutput
	ToEnvironmentMonitorsArrayOutputWithContext(context.Context) EnvironmentMonitorsArrayOutput
}

type EnvironmentMonitorsArray []EnvironmentMonitorsInput

func (EnvironmentMonitorsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentMonitors)(nil)).Elem()
}

func (i EnvironmentMonitorsArray) ToEnvironmentMonitorsArrayOutput() EnvironmentMonitorsArrayOutput {
	return i.ToEnvironmentMonitorsArrayOutputWithContext(context.Background())
}

func (i EnvironmentMonitorsArray) ToEnvironmentMonitorsArrayOutputWithContext(ctx context.Context) EnvironmentMonitorsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMonitorsArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html
type EnvironmentMonitorsOutput struct{ *pulumi.OutputState }

func (EnvironmentMonitorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentMonitors)(nil)).Elem()
}

func (o EnvironmentMonitorsOutput) ToEnvironmentMonitorsOutput() EnvironmentMonitorsOutput {
	return o
}

func (o EnvironmentMonitorsOutput) ToEnvironmentMonitorsOutputWithContext(ctx context.Context) EnvironmentMonitorsOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmarn
func (o EnvironmentMonitorsOutput) AlarmArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentMonitors) *string { return v.AlarmArn }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-monitors.html#cfn-appconfig-environment-monitors-alarmrolearn
func (o EnvironmentMonitorsOutput) AlarmRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentMonitors) *string { return v.AlarmRoleArn }).(pulumi.StringPtrOutput)
}

type EnvironmentMonitorsArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentMonitorsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentMonitors)(nil)).Elem()
}

func (o EnvironmentMonitorsArrayOutput) ToEnvironmentMonitorsArrayOutput() EnvironmentMonitorsArrayOutput {
	return o
}

func (o EnvironmentMonitorsArrayOutput) ToEnvironmentMonitorsArrayOutputWithContext(ctx context.Context) EnvironmentMonitorsArrayOutput {
	return o
}

func (o EnvironmentMonitorsArrayOutput) Index(i pulumi.IntInput) EnvironmentMonitorsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentMonitors {
		return vs[0].([]EnvironmentMonitors)[vs[1].(int)]
	}).(EnvironmentMonitorsOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html
type EnvironmentTags struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-key
	Key *string `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-value
	Value *string `pulumi:"value"`
}

// EnvironmentTagsInput is an input type that accepts EnvironmentTagsArgs and EnvironmentTagsOutput values.
// You can construct a concrete instance of `EnvironmentTagsInput` via:
//
//          EnvironmentTagsArgs{...}
type EnvironmentTagsInput interface {
	pulumi.Input

	ToEnvironmentTagsOutput() EnvironmentTagsOutput
	ToEnvironmentTagsOutputWithContext(context.Context) EnvironmentTagsOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html
type EnvironmentTagsArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-key
	Key pulumi.StringPtrInput `pulumi:"key"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-value
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EnvironmentTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentTags)(nil)).Elem()
}

func (i EnvironmentTagsArgs) ToEnvironmentTagsOutput() EnvironmentTagsOutput {
	return i.ToEnvironmentTagsOutputWithContext(context.Background())
}

func (i EnvironmentTagsArgs) ToEnvironmentTagsOutputWithContext(ctx context.Context) EnvironmentTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentTagsOutput)
}

// EnvironmentTagsArrayInput is an input type that accepts EnvironmentTagsArray and EnvironmentTagsArrayOutput values.
// You can construct a concrete instance of `EnvironmentTagsArrayInput` via:
//
//          EnvironmentTagsArray{ EnvironmentTagsArgs{...} }
type EnvironmentTagsArrayInput interface {
	pulumi.Input

	ToEnvironmentTagsArrayOutput() EnvironmentTagsArrayOutput
	ToEnvironmentTagsArrayOutputWithContext(context.Context) EnvironmentTagsArrayOutput
}

type EnvironmentTagsArray []EnvironmentTagsInput

func (EnvironmentTagsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentTags)(nil)).Elem()
}

func (i EnvironmentTagsArray) ToEnvironmentTagsArrayOutput() EnvironmentTagsArrayOutput {
	return i.ToEnvironmentTagsArrayOutputWithContext(context.Background())
}

func (i EnvironmentTagsArray) ToEnvironmentTagsArrayOutputWithContext(ctx context.Context) EnvironmentTagsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentTagsArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html
type EnvironmentTagsOutput struct{ *pulumi.OutputState }

func (EnvironmentTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentTags)(nil)).Elem()
}

func (o EnvironmentTagsOutput) ToEnvironmentTagsOutput() EnvironmentTagsOutput {
	return o
}

func (o EnvironmentTagsOutput) ToEnvironmentTagsOutputWithContext(ctx context.Context) EnvironmentTagsOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-key
func (o EnvironmentTagsOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentTags) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appconfig-environment-tags.html#cfn-appconfig-environment-tags-value
func (o EnvironmentTagsOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentTags) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentTagsArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentTagsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentTags)(nil)).Elem()
}

func (o EnvironmentTagsArrayOutput) ToEnvironmentTagsArrayOutput() EnvironmentTagsArrayOutput {
	return o
}

func (o EnvironmentTagsArrayOutput) ToEnvironmentTagsArrayOutputWithContext(ctx context.Context) EnvironmentTagsArrayOutput {
	return o
}

func (o EnvironmentTagsArrayOutput) Index(i pulumi.IntInput) EnvironmentTagsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentTags {
		return vs[0].([]EnvironmentTags)[vs[1].(int)]
	}).(EnvironmentTagsOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationTagsOutput{})
	pulumi.RegisterOutputType(ApplicationTagsArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileTagsOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileTagsArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileValidatorsOutput{})
	pulumi.RegisterOutputType(ConfigurationProfileValidatorsArrayOutput{})
	pulumi.RegisterOutputType(DeploymentStrategyTagsOutput{})
	pulumi.RegisterOutputType(DeploymentStrategyTagsArrayOutput{})
	pulumi.RegisterOutputType(DeploymentTagsOutput{})
	pulumi.RegisterOutputType(DeploymentTagsArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentMonitorsOutput{})
	pulumi.RegisterOutputType(EnvironmentMonitorsArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentTagsOutput{})
	pulumi.RegisterOutputType(EnvironmentTagsArrayOutput{})
}
