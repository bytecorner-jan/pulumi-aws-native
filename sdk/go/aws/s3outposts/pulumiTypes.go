// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package s3outposts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-accesspoint-vpcconfiguration.html
type AccessPointVpcConfiguration struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-accesspoint-vpcconfiguration.html#cfn-s3outposts-accesspoint-vpcconfiguration-vpcid
	VpcId *string `pulumi:"vpcId"`
}

// AccessPointVpcConfigurationInput is an input type that accepts AccessPointVpcConfigurationArgs and AccessPointVpcConfigurationOutput values.
// You can construct a concrete instance of `AccessPointVpcConfigurationInput` via:
//
//          AccessPointVpcConfigurationArgs{...}
type AccessPointVpcConfigurationInput interface {
	pulumi.Input

	ToAccessPointVpcConfigurationOutput() AccessPointVpcConfigurationOutput
	ToAccessPointVpcConfigurationOutputWithContext(context.Context) AccessPointVpcConfigurationOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-accesspoint-vpcconfiguration.html
type AccessPointVpcConfigurationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-accesspoint-vpcconfiguration.html#cfn-s3outposts-accesspoint-vpcconfiguration-vpcid
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (AccessPointVpcConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPointVpcConfiguration)(nil)).Elem()
}

func (i AccessPointVpcConfigurationArgs) ToAccessPointVpcConfigurationOutput() AccessPointVpcConfigurationOutput {
	return i.ToAccessPointVpcConfigurationOutputWithContext(context.Background())
}

func (i AccessPointVpcConfigurationArgs) ToAccessPointVpcConfigurationOutputWithContext(ctx context.Context) AccessPointVpcConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointVpcConfigurationOutput)
}

func (i AccessPointVpcConfigurationArgs) ToAccessPointVpcConfigurationPtrOutput() AccessPointVpcConfigurationPtrOutput {
	return i.ToAccessPointVpcConfigurationPtrOutputWithContext(context.Background())
}

func (i AccessPointVpcConfigurationArgs) ToAccessPointVpcConfigurationPtrOutputWithContext(ctx context.Context) AccessPointVpcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointVpcConfigurationOutput).ToAccessPointVpcConfigurationPtrOutputWithContext(ctx)
}

// AccessPointVpcConfigurationPtrInput is an input type that accepts AccessPointVpcConfigurationArgs, AccessPointVpcConfigurationPtr and AccessPointVpcConfigurationPtrOutput values.
// You can construct a concrete instance of `AccessPointVpcConfigurationPtrInput` via:
//
//          AccessPointVpcConfigurationArgs{...}
//
//  or:
//
//          nil
type AccessPointVpcConfigurationPtrInput interface {
	pulumi.Input

	ToAccessPointVpcConfigurationPtrOutput() AccessPointVpcConfigurationPtrOutput
	ToAccessPointVpcConfigurationPtrOutputWithContext(context.Context) AccessPointVpcConfigurationPtrOutput
}

type accessPointVpcConfigurationPtrType AccessPointVpcConfigurationArgs

func AccessPointVpcConfigurationPtr(v *AccessPointVpcConfigurationArgs) AccessPointVpcConfigurationPtrInput {
	return (*accessPointVpcConfigurationPtrType)(v)
}

func (*accessPointVpcConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPointVpcConfiguration)(nil)).Elem()
}

func (i *accessPointVpcConfigurationPtrType) ToAccessPointVpcConfigurationPtrOutput() AccessPointVpcConfigurationPtrOutput {
	return i.ToAccessPointVpcConfigurationPtrOutputWithContext(context.Background())
}

func (i *accessPointVpcConfigurationPtrType) ToAccessPointVpcConfigurationPtrOutputWithContext(ctx context.Context) AccessPointVpcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessPointVpcConfigurationPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-accesspoint-vpcconfiguration.html
type AccessPointVpcConfigurationOutput struct{ *pulumi.OutputState }

func (AccessPointVpcConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessPointVpcConfiguration)(nil)).Elem()
}

func (o AccessPointVpcConfigurationOutput) ToAccessPointVpcConfigurationOutput() AccessPointVpcConfigurationOutput {
	return o
}

func (o AccessPointVpcConfigurationOutput) ToAccessPointVpcConfigurationOutputWithContext(ctx context.Context) AccessPointVpcConfigurationOutput {
	return o
}

func (o AccessPointVpcConfigurationOutput) ToAccessPointVpcConfigurationPtrOutput() AccessPointVpcConfigurationPtrOutput {
	return o.ToAccessPointVpcConfigurationPtrOutputWithContext(context.Background())
}

func (o AccessPointVpcConfigurationOutput) ToAccessPointVpcConfigurationPtrOutputWithContext(ctx context.Context) AccessPointVpcConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessPointVpcConfiguration) *AccessPointVpcConfiguration {
		return &v
	}).(AccessPointVpcConfigurationPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-accesspoint-vpcconfiguration.html#cfn-s3outposts-accesspoint-vpcconfiguration-vpcid
func (o AccessPointVpcConfigurationOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccessPointVpcConfiguration) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

type AccessPointVpcConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AccessPointVpcConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessPointVpcConfiguration)(nil)).Elem()
}

func (o AccessPointVpcConfigurationPtrOutput) ToAccessPointVpcConfigurationPtrOutput() AccessPointVpcConfigurationPtrOutput {
	return o
}

func (o AccessPointVpcConfigurationPtrOutput) ToAccessPointVpcConfigurationPtrOutputWithContext(ctx context.Context) AccessPointVpcConfigurationPtrOutput {
	return o
}

func (o AccessPointVpcConfigurationPtrOutput) Elem() AccessPointVpcConfigurationOutput {
	return o.ApplyT(func(v *AccessPointVpcConfiguration) AccessPointVpcConfiguration {
		if v != nil {
			return *v
		}
		var ret AccessPointVpcConfiguration
		return ret
	}).(AccessPointVpcConfigurationOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-accesspoint-vpcconfiguration.html#cfn-s3outposts-accesspoint-vpcconfiguration-vpcid
func (o AccessPointVpcConfigurationPtrOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessPointVpcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.VpcId
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html
type BucketAbortIncompleteMultipartUpload struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html#cfn-s3outposts-bucket-abortincompletemultipartupload-daysafterinitiation
	DaysAfterInitiation int `pulumi:"daysAfterInitiation"`
}

// BucketAbortIncompleteMultipartUploadInput is an input type that accepts BucketAbortIncompleteMultipartUploadArgs and BucketAbortIncompleteMultipartUploadOutput values.
// You can construct a concrete instance of `BucketAbortIncompleteMultipartUploadInput` via:
//
//          BucketAbortIncompleteMultipartUploadArgs{...}
type BucketAbortIncompleteMultipartUploadInput interface {
	pulumi.Input

	ToBucketAbortIncompleteMultipartUploadOutput() BucketAbortIncompleteMultipartUploadOutput
	ToBucketAbortIncompleteMultipartUploadOutputWithContext(context.Context) BucketAbortIncompleteMultipartUploadOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html
type BucketAbortIncompleteMultipartUploadArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html#cfn-s3outposts-bucket-abortincompletemultipartupload-daysafterinitiation
	DaysAfterInitiation pulumi.IntInput `pulumi:"daysAfterInitiation"`
}

func (BucketAbortIncompleteMultipartUploadArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketAbortIncompleteMultipartUpload)(nil)).Elem()
}

func (i BucketAbortIncompleteMultipartUploadArgs) ToBucketAbortIncompleteMultipartUploadOutput() BucketAbortIncompleteMultipartUploadOutput {
	return i.ToBucketAbortIncompleteMultipartUploadOutputWithContext(context.Background())
}

func (i BucketAbortIncompleteMultipartUploadArgs) ToBucketAbortIncompleteMultipartUploadOutputWithContext(ctx context.Context) BucketAbortIncompleteMultipartUploadOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAbortIncompleteMultipartUploadOutput)
}

func (i BucketAbortIncompleteMultipartUploadArgs) ToBucketAbortIncompleteMultipartUploadPtrOutput() BucketAbortIncompleteMultipartUploadPtrOutput {
	return i.ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(context.Background())
}

func (i BucketAbortIncompleteMultipartUploadArgs) ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(ctx context.Context) BucketAbortIncompleteMultipartUploadPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAbortIncompleteMultipartUploadOutput).ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(ctx)
}

// BucketAbortIncompleteMultipartUploadPtrInput is an input type that accepts BucketAbortIncompleteMultipartUploadArgs, BucketAbortIncompleteMultipartUploadPtr and BucketAbortIncompleteMultipartUploadPtrOutput values.
// You can construct a concrete instance of `BucketAbortIncompleteMultipartUploadPtrInput` via:
//
//          BucketAbortIncompleteMultipartUploadArgs{...}
//
//  or:
//
//          nil
type BucketAbortIncompleteMultipartUploadPtrInput interface {
	pulumi.Input

	ToBucketAbortIncompleteMultipartUploadPtrOutput() BucketAbortIncompleteMultipartUploadPtrOutput
	ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(context.Context) BucketAbortIncompleteMultipartUploadPtrOutput
}

type bucketAbortIncompleteMultipartUploadPtrType BucketAbortIncompleteMultipartUploadArgs

func BucketAbortIncompleteMultipartUploadPtr(v *BucketAbortIncompleteMultipartUploadArgs) BucketAbortIncompleteMultipartUploadPtrInput {
	return (*bucketAbortIncompleteMultipartUploadPtrType)(v)
}

func (*bucketAbortIncompleteMultipartUploadPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAbortIncompleteMultipartUpload)(nil)).Elem()
}

func (i *bucketAbortIncompleteMultipartUploadPtrType) ToBucketAbortIncompleteMultipartUploadPtrOutput() BucketAbortIncompleteMultipartUploadPtrOutput {
	return i.ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(context.Background())
}

func (i *bucketAbortIncompleteMultipartUploadPtrType) ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(ctx context.Context) BucketAbortIncompleteMultipartUploadPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketAbortIncompleteMultipartUploadPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html
type BucketAbortIncompleteMultipartUploadOutput struct{ *pulumi.OutputState }

func (BucketAbortIncompleteMultipartUploadOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketAbortIncompleteMultipartUpload)(nil)).Elem()
}

func (o BucketAbortIncompleteMultipartUploadOutput) ToBucketAbortIncompleteMultipartUploadOutput() BucketAbortIncompleteMultipartUploadOutput {
	return o
}

func (o BucketAbortIncompleteMultipartUploadOutput) ToBucketAbortIncompleteMultipartUploadOutputWithContext(ctx context.Context) BucketAbortIncompleteMultipartUploadOutput {
	return o
}

func (o BucketAbortIncompleteMultipartUploadOutput) ToBucketAbortIncompleteMultipartUploadPtrOutput() BucketAbortIncompleteMultipartUploadPtrOutput {
	return o.ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(context.Background())
}

func (o BucketAbortIncompleteMultipartUploadOutput) ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(ctx context.Context) BucketAbortIncompleteMultipartUploadPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BucketAbortIncompleteMultipartUpload) *BucketAbortIncompleteMultipartUpload {
		return &v
	}).(BucketAbortIncompleteMultipartUploadPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html#cfn-s3outposts-bucket-abortincompletemultipartupload-daysafterinitiation
func (o BucketAbortIncompleteMultipartUploadOutput) DaysAfterInitiation() pulumi.IntOutput {
	return o.ApplyT(func(v BucketAbortIncompleteMultipartUpload) int { return v.DaysAfterInitiation }).(pulumi.IntOutput)
}

type BucketAbortIncompleteMultipartUploadPtrOutput struct{ *pulumi.OutputState }

func (BucketAbortIncompleteMultipartUploadPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketAbortIncompleteMultipartUpload)(nil)).Elem()
}

func (o BucketAbortIncompleteMultipartUploadPtrOutput) ToBucketAbortIncompleteMultipartUploadPtrOutput() BucketAbortIncompleteMultipartUploadPtrOutput {
	return o
}

func (o BucketAbortIncompleteMultipartUploadPtrOutput) ToBucketAbortIncompleteMultipartUploadPtrOutputWithContext(ctx context.Context) BucketAbortIncompleteMultipartUploadPtrOutput {
	return o
}

func (o BucketAbortIncompleteMultipartUploadPtrOutput) Elem() BucketAbortIncompleteMultipartUploadOutput {
	return o.ApplyT(func(v *BucketAbortIncompleteMultipartUpload) BucketAbortIncompleteMultipartUpload {
		if v != nil {
			return *v
		}
		var ret BucketAbortIncompleteMultipartUpload
		return ret
	}).(BucketAbortIncompleteMultipartUploadOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-abortincompletemultipartupload.html#cfn-s3outposts-bucket-abortincompletemultipartupload-daysafterinitiation
func (o BucketAbortIncompleteMultipartUploadPtrOutput) DaysAfterInitiation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BucketAbortIncompleteMultipartUpload) *int {
		if v == nil {
			return nil
		}
		return &v.DaysAfterInitiation
	}).(pulumi.IntPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html
type BucketLifecycleConfiguration struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html#cfn-s3outposts-bucket-lifecycleconfiguration-rules
	Rules []BucketRule `pulumi:"rules"`
}

// BucketLifecycleConfigurationInput is an input type that accepts BucketLifecycleConfigurationArgs and BucketLifecycleConfigurationOutput values.
// You can construct a concrete instance of `BucketLifecycleConfigurationInput` via:
//
//          BucketLifecycleConfigurationArgs{...}
type BucketLifecycleConfigurationInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput
	ToBucketLifecycleConfigurationOutputWithContext(context.Context) BucketLifecycleConfigurationOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html
type BucketLifecycleConfigurationArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html#cfn-s3outposts-bucket-lifecycleconfiguration-rules
	Rules BucketRuleArrayInput `pulumi:"rules"`
}

func (BucketLifecycleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleConfiguration)(nil)).Elem()
}

func (i BucketLifecycleConfigurationArgs) ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput {
	return i.ToBucketLifecycleConfigurationOutputWithContext(context.Background())
}

func (i BucketLifecycleConfigurationArgs) ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationOutput)
}

func (i BucketLifecycleConfigurationArgs) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return i.ToBucketLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (i BucketLifecycleConfigurationArgs) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationOutput).ToBucketLifecycleConfigurationPtrOutputWithContext(ctx)
}

// BucketLifecycleConfigurationPtrInput is an input type that accepts BucketLifecycleConfigurationArgs, BucketLifecycleConfigurationPtr and BucketLifecycleConfigurationPtrOutput values.
// You can construct a concrete instance of `BucketLifecycleConfigurationPtrInput` via:
//
//          BucketLifecycleConfigurationArgs{...}
//
//  or:
//
//          nil
type BucketLifecycleConfigurationPtrInput interface {
	pulumi.Input

	ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput
	ToBucketLifecycleConfigurationPtrOutputWithContext(context.Context) BucketLifecycleConfigurationPtrOutput
}

type bucketLifecycleConfigurationPtrType BucketLifecycleConfigurationArgs

func BucketLifecycleConfigurationPtr(v *BucketLifecycleConfigurationArgs) BucketLifecycleConfigurationPtrInput {
	return (*bucketLifecycleConfigurationPtrType)(v)
}

func (*bucketLifecycleConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleConfiguration)(nil)).Elem()
}

func (i *bucketLifecycleConfigurationPtrType) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return i.ToBucketLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (i *bucketLifecycleConfigurationPtrType) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketLifecycleConfigurationPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html
type BucketLifecycleConfigurationOutput struct{ *pulumi.OutputState }

func (BucketLifecycleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketLifecycleConfiguration)(nil)).Elem()
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationOutput() BucketLifecycleConfigurationOutput {
	return o
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationOutputWithContext(ctx context.Context) BucketLifecycleConfigurationOutput {
	return o
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return o.ToBucketLifecycleConfigurationPtrOutputWithContext(context.Background())
}

func (o BucketLifecycleConfigurationOutput) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BucketLifecycleConfiguration) *BucketLifecycleConfiguration {
		return &v
	}).(BucketLifecycleConfigurationPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html#cfn-s3outposts-bucket-lifecycleconfiguration-rules
func (o BucketLifecycleConfigurationOutput) Rules() BucketRuleArrayOutput {
	return o.ApplyT(func(v BucketLifecycleConfiguration) []BucketRule { return v.Rules }).(BucketRuleArrayOutput)
}

type BucketLifecycleConfigurationPtrOutput struct{ *pulumi.OutputState }

func (BucketLifecycleConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketLifecycleConfiguration)(nil)).Elem()
}

func (o BucketLifecycleConfigurationPtrOutput) ToBucketLifecycleConfigurationPtrOutput() BucketLifecycleConfigurationPtrOutput {
	return o
}

func (o BucketLifecycleConfigurationPtrOutput) ToBucketLifecycleConfigurationPtrOutputWithContext(ctx context.Context) BucketLifecycleConfigurationPtrOutput {
	return o
}

func (o BucketLifecycleConfigurationPtrOutput) Elem() BucketLifecycleConfigurationOutput {
	return o.ApplyT(func(v *BucketLifecycleConfiguration) BucketLifecycleConfiguration {
		if v != nil {
			return *v
		}
		var ret BucketLifecycleConfiguration
		return ret
	}).(BucketLifecycleConfigurationOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-lifecycleconfiguration.html#cfn-s3outposts-bucket-lifecycleconfiguration-rules
func (o BucketLifecycleConfigurationPtrOutput) Rules() BucketRuleArrayOutput {
	return o.ApplyT(func(v *BucketLifecycleConfiguration) []BucketRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(BucketRuleArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html
type BucketRule struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-abortincompletemultipartupload
	AbortIncompleteMultipartUpload *BucketAbortIncompleteMultipartUpload `pulumi:"abortIncompleteMultipartUpload"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationdate
	ExpirationDate *string `pulumi:"expirationDate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationindays
	ExpirationInDays *int `pulumi:"expirationInDays"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-filter
	Filter interface{} `pulumi:"filter"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-id
	Id *string `pulumi:"id"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-status
	Status *string `pulumi:"status"`
}

// BucketRuleInput is an input type that accepts BucketRuleArgs and BucketRuleOutput values.
// You can construct a concrete instance of `BucketRuleInput` via:
//
//          BucketRuleArgs{...}
type BucketRuleInput interface {
	pulumi.Input

	ToBucketRuleOutput() BucketRuleOutput
	ToBucketRuleOutputWithContext(context.Context) BucketRuleOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html
type BucketRuleArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-abortincompletemultipartupload
	AbortIncompleteMultipartUpload BucketAbortIncompleteMultipartUploadPtrInput `pulumi:"abortIncompleteMultipartUpload"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationdate
	ExpirationDate pulumi.StringPtrInput `pulumi:"expirationDate"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationindays
	ExpirationInDays pulumi.IntPtrInput `pulumi:"expirationInDays"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-filter
	Filter pulumi.Input `pulumi:"filter"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-id
	Id pulumi.StringPtrInput `pulumi:"id"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-status
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (BucketRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketRule)(nil)).Elem()
}

func (i BucketRuleArgs) ToBucketRuleOutput() BucketRuleOutput {
	return i.ToBucketRuleOutputWithContext(context.Background())
}

func (i BucketRuleArgs) ToBucketRuleOutputWithContext(ctx context.Context) BucketRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketRuleOutput)
}

// BucketRuleArrayInput is an input type that accepts BucketRuleArray and BucketRuleArrayOutput values.
// You can construct a concrete instance of `BucketRuleArrayInput` via:
//
//          BucketRuleArray{ BucketRuleArgs{...} }
type BucketRuleArrayInput interface {
	pulumi.Input

	ToBucketRuleArrayOutput() BucketRuleArrayOutput
	ToBucketRuleArrayOutputWithContext(context.Context) BucketRuleArrayOutput
}

type BucketRuleArray []BucketRuleInput

func (BucketRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketRule)(nil)).Elem()
}

func (i BucketRuleArray) ToBucketRuleArrayOutput() BucketRuleArrayOutput {
	return i.ToBucketRuleArrayOutputWithContext(context.Background())
}

func (i BucketRuleArray) ToBucketRuleArrayOutputWithContext(ctx context.Context) BucketRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketRuleArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html
type BucketRuleOutput struct{ *pulumi.OutputState }

func (BucketRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketRule)(nil)).Elem()
}

func (o BucketRuleOutput) ToBucketRuleOutput() BucketRuleOutput {
	return o
}

func (o BucketRuleOutput) ToBucketRuleOutputWithContext(ctx context.Context) BucketRuleOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-abortincompletemultipartupload
func (o BucketRuleOutput) AbortIncompleteMultipartUpload() BucketAbortIncompleteMultipartUploadPtrOutput {
	return o.ApplyT(func(v BucketRule) *BucketAbortIncompleteMultipartUpload { return v.AbortIncompleteMultipartUpload }).(BucketAbortIncompleteMultipartUploadPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationdate
func (o BucketRuleOutput) ExpirationDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketRule) *string { return v.ExpirationDate }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-expirationindays
func (o BucketRuleOutput) ExpirationInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v BucketRule) *int { return v.ExpirationInDays }).(pulumi.IntPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-filter
func (o BucketRuleOutput) Filter() pulumi.AnyOutput {
	return o.ApplyT(func(v BucketRule) interface{} { return v.Filter }).(pulumi.AnyOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-id
func (o BucketRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-bucket-rule.html#cfn-s3outposts-bucket-rule-status
func (o BucketRuleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BucketRule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type BucketRuleArrayOutput struct{ *pulumi.OutputState }

func (BucketRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BucketRule)(nil)).Elem()
}

func (o BucketRuleArrayOutput) ToBucketRuleArrayOutput() BucketRuleArrayOutput {
	return o
}

func (o BucketRuleArrayOutput) ToBucketRuleArrayOutputWithContext(ctx context.Context) BucketRuleArrayOutput {
	return o
}

func (o BucketRuleArrayOutput) Index(i pulumi.IntInput) BucketRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BucketRule {
		return vs[0].([]BucketRule)[vs[1].(int)]
	}).(BucketRuleOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-networkinterface.html
type EndpointNetworkInterface struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-networkinterface.html#cfn-s3outposts-endpoint-networkinterface-networkinterfaceid
	NetworkInterfaceId string `pulumi:"networkInterfaceId"`
}

// EndpointNetworkInterfaceInput is an input type that accepts EndpointNetworkInterfaceArgs and EndpointNetworkInterfaceOutput values.
// You can construct a concrete instance of `EndpointNetworkInterfaceInput` via:
//
//          EndpointNetworkInterfaceArgs{...}
type EndpointNetworkInterfaceInput interface {
	pulumi.Input

	ToEndpointNetworkInterfaceOutput() EndpointNetworkInterfaceOutput
	ToEndpointNetworkInterfaceOutputWithContext(context.Context) EndpointNetworkInterfaceOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-networkinterface.html
type EndpointNetworkInterfaceArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-networkinterface.html#cfn-s3outposts-endpoint-networkinterface-networkinterfaceid
	NetworkInterfaceId pulumi.StringInput `pulumi:"networkInterfaceId"`
}

func (EndpointNetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointNetworkInterface)(nil)).Elem()
}

func (i EndpointNetworkInterfaceArgs) ToEndpointNetworkInterfaceOutput() EndpointNetworkInterfaceOutput {
	return i.ToEndpointNetworkInterfaceOutputWithContext(context.Background())
}

func (i EndpointNetworkInterfaceArgs) ToEndpointNetworkInterfaceOutputWithContext(ctx context.Context) EndpointNetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointNetworkInterfaceOutput)
}

// EndpointNetworkInterfaceArrayInput is an input type that accepts EndpointNetworkInterfaceArray and EndpointNetworkInterfaceArrayOutput values.
// You can construct a concrete instance of `EndpointNetworkInterfaceArrayInput` via:
//
//          EndpointNetworkInterfaceArray{ EndpointNetworkInterfaceArgs{...} }
type EndpointNetworkInterfaceArrayInput interface {
	pulumi.Input

	ToEndpointNetworkInterfaceArrayOutput() EndpointNetworkInterfaceArrayOutput
	ToEndpointNetworkInterfaceArrayOutputWithContext(context.Context) EndpointNetworkInterfaceArrayOutput
}

type EndpointNetworkInterfaceArray []EndpointNetworkInterfaceInput

func (EndpointNetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointNetworkInterface)(nil)).Elem()
}

func (i EndpointNetworkInterfaceArray) ToEndpointNetworkInterfaceArrayOutput() EndpointNetworkInterfaceArrayOutput {
	return i.ToEndpointNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i EndpointNetworkInterfaceArray) ToEndpointNetworkInterfaceArrayOutputWithContext(ctx context.Context) EndpointNetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointNetworkInterfaceArrayOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-networkinterface.html
type EndpointNetworkInterfaceOutput struct{ *pulumi.OutputState }

func (EndpointNetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointNetworkInterface)(nil)).Elem()
}

func (o EndpointNetworkInterfaceOutput) ToEndpointNetworkInterfaceOutput() EndpointNetworkInterfaceOutput {
	return o
}

func (o EndpointNetworkInterfaceOutput) ToEndpointNetworkInterfaceOutputWithContext(ctx context.Context) EndpointNetworkInterfaceOutput {
	return o
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3outposts-endpoint-networkinterface.html#cfn-s3outposts-endpoint-networkinterface-networkinterfaceid
func (o EndpointNetworkInterfaceOutput) NetworkInterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointNetworkInterface) string { return v.NetworkInterfaceId }).(pulumi.StringOutput)
}

type EndpointNetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (EndpointNetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointNetworkInterface)(nil)).Elem()
}

func (o EndpointNetworkInterfaceArrayOutput) ToEndpointNetworkInterfaceArrayOutput() EndpointNetworkInterfaceArrayOutput {
	return o
}

func (o EndpointNetworkInterfaceArrayOutput) ToEndpointNetworkInterfaceArrayOutputWithContext(ctx context.Context) EndpointNetworkInterfaceArrayOutput {
	return o
}

func (o EndpointNetworkInterfaceArrayOutput) Index(i pulumi.IntInput) EndpointNetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointNetworkInterface {
		return vs[0].([]EndpointNetworkInterface)[vs[1].(int)]
	}).(EndpointNetworkInterfaceOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessPointVpcConfigurationOutput{})
	pulumi.RegisterOutputType(AccessPointVpcConfigurationPtrOutput{})
	pulumi.RegisterOutputType(BucketAbortIncompleteMultipartUploadOutput{})
	pulumi.RegisterOutputType(BucketAbortIncompleteMultipartUploadPtrOutput{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationOutput{})
	pulumi.RegisterOutputType(BucketLifecycleConfigurationPtrOutput{})
	pulumi.RegisterOutputType(BucketRuleOutput{})
	pulumi.RegisterOutputType(BucketRuleArrayOutput{})
	pulumi.RegisterOutputType(EndpointNetworkInterfaceOutput{})
	pulumi.RegisterOutputType(EndpointNetworkInterfaceArrayOutput{})
}
