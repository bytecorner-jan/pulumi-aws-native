// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MetricFilterMetricTransformation struct {
	DefaultValue    *float64 `pulumi:"defaultValue"`
	MetricName      string   `pulumi:"metricName"`
	MetricNamespace string   `pulumi:"metricNamespace"`
	MetricValue     string   `pulumi:"metricValue"`
}

// MetricFilterMetricTransformationInput is an input type that accepts MetricFilterMetricTransformationArgs and MetricFilterMetricTransformationOutput values.
// You can construct a concrete instance of `MetricFilterMetricTransformationInput` via:
//
//          MetricFilterMetricTransformationArgs{...}
type MetricFilterMetricTransformationInput interface {
	pulumi.Input

	ToMetricFilterMetricTransformationOutput() MetricFilterMetricTransformationOutput
	ToMetricFilterMetricTransformationOutputWithContext(context.Context) MetricFilterMetricTransformationOutput
}

type MetricFilterMetricTransformationArgs struct {
	DefaultValue    pulumi.Float64PtrInput `pulumi:"defaultValue"`
	MetricName      pulumi.StringInput     `pulumi:"metricName"`
	MetricNamespace pulumi.StringInput     `pulumi:"metricNamespace"`
	MetricValue     pulumi.StringInput     `pulumi:"metricValue"`
}

func (MetricFilterMetricTransformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricFilterMetricTransformation)(nil)).Elem()
}

func (i MetricFilterMetricTransformationArgs) ToMetricFilterMetricTransformationOutput() MetricFilterMetricTransformationOutput {
	return i.ToMetricFilterMetricTransformationOutputWithContext(context.Background())
}

func (i MetricFilterMetricTransformationArgs) ToMetricFilterMetricTransformationOutputWithContext(ctx context.Context) MetricFilterMetricTransformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricFilterMetricTransformationOutput)
}

// MetricFilterMetricTransformationArrayInput is an input type that accepts MetricFilterMetricTransformationArray and MetricFilterMetricTransformationArrayOutput values.
// You can construct a concrete instance of `MetricFilterMetricTransformationArrayInput` via:
//
//          MetricFilterMetricTransformationArray{ MetricFilterMetricTransformationArgs{...} }
type MetricFilterMetricTransformationArrayInput interface {
	pulumi.Input

	ToMetricFilterMetricTransformationArrayOutput() MetricFilterMetricTransformationArrayOutput
	ToMetricFilterMetricTransformationArrayOutputWithContext(context.Context) MetricFilterMetricTransformationArrayOutput
}

type MetricFilterMetricTransformationArray []MetricFilterMetricTransformationInput

func (MetricFilterMetricTransformationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricFilterMetricTransformation)(nil)).Elem()
}

func (i MetricFilterMetricTransformationArray) ToMetricFilterMetricTransformationArrayOutput() MetricFilterMetricTransformationArrayOutput {
	return i.ToMetricFilterMetricTransformationArrayOutputWithContext(context.Background())
}

func (i MetricFilterMetricTransformationArray) ToMetricFilterMetricTransformationArrayOutputWithContext(ctx context.Context) MetricFilterMetricTransformationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricFilterMetricTransformationArrayOutput)
}

type MetricFilterMetricTransformationOutput struct{ *pulumi.OutputState }

func (MetricFilterMetricTransformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricFilterMetricTransformation)(nil)).Elem()
}

func (o MetricFilterMetricTransformationOutput) ToMetricFilterMetricTransformationOutput() MetricFilterMetricTransformationOutput {
	return o
}

func (o MetricFilterMetricTransformationOutput) ToMetricFilterMetricTransformationOutputWithContext(ctx context.Context) MetricFilterMetricTransformationOutput {
	return o
}

func (o MetricFilterMetricTransformationOutput) DefaultValue() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v MetricFilterMetricTransformation) *float64 { return v.DefaultValue }).(pulumi.Float64PtrOutput)
}

func (o MetricFilterMetricTransformationOutput) MetricName() pulumi.StringOutput {
	return o.ApplyT(func(v MetricFilterMetricTransformation) string { return v.MetricName }).(pulumi.StringOutput)
}

func (o MetricFilterMetricTransformationOutput) MetricNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v MetricFilterMetricTransformation) string { return v.MetricNamespace }).(pulumi.StringOutput)
}

func (o MetricFilterMetricTransformationOutput) MetricValue() pulumi.StringOutput {
	return o.ApplyT(func(v MetricFilterMetricTransformation) string { return v.MetricValue }).(pulumi.StringOutput)
}

type MetricFilterMetricTransformationArrayOutput struct{ *pulumi.OutputState }

func (MetricFilterMetricTransformationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricFilterMetricTransformation)(nil)).Elem()
}

func (o MetricFilterMetricTransformationArrayOutput) ToMetricFilterMetricTransformationArrayOutput() MetricFilterMetricTransformationArrayOutput {
	return o
}

func (o MetricFilterMetricTransformationArrayOutput) ToMetricFilterMetricTransformationArrayOutputWithContext(ctx context.Context) MetricFilterMetricTransformationArrayOutput {
	return o
}

func (o MetricFilterMetricTransformationArrayOutput) Index(i pulumi.IntInput) MetricFilterMetricTransformationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricFilterMetricTransformation {
		return vs[0].([]MetricFilterMetricTransformation)[vs[1].(int)]
	}).(MetricFilterMetricTransformationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricFilterMetricTransformationInput)(nil)).Elem(), MetricFilterMetricTransformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MetricFilterMetricTransformationArrayInput)(nil)).Elem(), MetricFilterMetricTransformationArray{})
	pulumi.RegisterOutputType(MetricFilterMetricTransformationOutput{})
	pulumi.RegisterOutputType(MetricFilterMetricTransformationArrayOutput{})
}
