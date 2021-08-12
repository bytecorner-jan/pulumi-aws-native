// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package redshift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html
type ClusterEndpoint struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-address
	Address *string `pulumi:"address"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-port
	Port *string `pulumi:"port"`
}

// ClusterEndpointInput is an input type that accepts ClusterEndpointArgs and ClusterEndpointOutput values.
// You can construct a concrete instance of `ClusterEndpointInput` via:
//
//          ClusterEndpointArgs{...}
type ClusterEndpointInput interface {
	pulumi.Input

	ToClusterEndpointOutput() ClusterEndpointOutput
	ToClusterEndpointOutputWithContext(context.Context) ClusterEndpointOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html
type ClusterEndpointArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-address
	Address pulumi.StringPtrInput `pulumi:"address"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-port
	Port pulumi.StringPtrInput `pulumi:"port"`
}

func (ClusterEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEndpoint)(nil)).Elem()
}

func (i ClusterEndpointArgs) ToClusterEndpointOutput() ClusterEndpointOutput {
	return i.ToClusterEndpointOutputWithContext(context.Background())
}

func (i ClusterEndpointArgs) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointOutput)
}

func (i ClusterEndpointArgs) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return i.ToClusterEndpointPtrOutputWithContext(context.Background())
}

func (i ClusterEndpointArgs) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointOutput).ToClusterEndpointPtrOutputWithContext(ctx)
}

// ClusterEndpointPtrInput is an input type that accepts ClusterEndpointArgs, ClusterEndpointPtr and ClusterEndpointPtrOutput values.
// You can construct a concrete instance of `ClusterEndpointPtrInput` via:
//
//          ClusterEndpointArgs{...}
//
//  or:
//
//          nil
type ClusterEndpointPtrInput interface {
	pulumi.Input

	ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput
	ToClusterEndpointPtrOutputWithContext(context.Context) ClusterEndpointPtrOutput
}

type clusterEndpointPtrType ClusterEndpointArgs

func ClusterEndpointPtr(v *ClusterEndpointArgs) ClusterEndpointPtrInput {
	return (*clusterEndpointPtrType)(v)
}

func (*clusterEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (i *clusterEndpointPtrType) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return i.ToClusterEndpointPtrOutputWithContext(context.Background())
}

func (i *clusterEndpointPtrType) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterEndpointPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html
type ClusterEndpointOutput struct{ *pulumi.OutputState }

func (ClusterEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointOutput) ToClusterEndpointOutput() ClusterEndpointOutput {
	return o
}

func (o ClusterEndpointOutput) ToClusterEndpointOutputWithContext(ctx context.Context) ClusterEndpointOutput {
	return o
}

func (o ClusterEndpointOutput) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return o.ToClusterEndpointPtrOutputWithContext(context.Background())
}

func (o ClusterEndpointOutput) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterEndpoint) *ClusterEndpoint {
		return &v
	}).(ClusterEndpointPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-address
func (o ClusterEndpointOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterEndpoint) *string { return v.Address }).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-port
func (o ClusterEndpointOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterEndpoint) *string { return v.Port }).(pulumi.StringPtrOutput)
}

type ClusterEndpointPtrOutput struct{ *pulumi.OutputState }

func (ClusterEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterEndpoint)(nil)).Elem()
}

func (o ClusterEndpointPtrOutput) ToClusterEndpointPtrOutput() ClusterEndpointPtrOutput {
	return o
}

func (o ClusterEndpointPtrOutput) ToClusterEndpointPtrOutputWithContext(ctx context.Context) ClusterEndpointPtrOutput {
	return o
}

func (o ClusterEndpointPtrOutput) Elem() ClusterEndpointOutput {
	return o.ApplyT(func(v *ClusterEndpoint) ClusterEndpoint {
		if v != nil {
			return *v
		}
		var ret ClusterEndpoint
		return ret
	}).(ClusterEndpointOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-address
func (o ClusterEndpointPtrOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Address
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-endpoint.html#cfn-redshift-cluster-endpoint-port
func (o ClusterEndpointPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type ClusterLoggingProperties struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
	BucketName string `pulumi:"bucketName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
}

// ClusterLoggingPropertiesInput is an input type that accepts ClusterLoggingPropertiesArgs and ClusterLoggingPropertiesOutput values.
// You can construct a concrete instance of `ClusterLoggingPropertiesInput` via:
//
//          ClusterLoggingPropertiesArgs{...}
type ClusterLoggingPropertiesInput interface {
	pulumi.Input

	ToClusterLoggingPropertiesOutput() ClusterLoggingPropertiesOutput
	ToClusterLoggingPropertiesOutputWithContext(context.Context) ClusterLoggingPropertiesOutput
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type ClusterLoggingPropertiesArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
	BucketName pulumi.StringInput `pulumi:"bucketName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
	S3KeyPrefix pulumi.StringPtrInput `pulumi:"s3KeyPrefix"`
}

func (ClusterLoggingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLoggingProperties)(nil)).Elem()
}

func (i ClusterLoggingPropertiesArgs) ToClusterLoggingPropertiesOutput() ClusterLoggingPropertiesOutput {
	return i.ToClusterLoggingPropertiesOutputWithContext(context.Background())
}

func (i ClusterLoggingPropertiesArgs) ToClusterLoggingPropertiesOutputWithContext(ctx context.Context) ClusterLoggingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingPropertiesOutput)
}

func (i ClusterLoggingPropertiesArgs) ToClusterLoggingPropertiesPtrOutput() ClusterLoggingPropertiesPtrOutput {
	return i.ToClusterLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterLoggingPropertiesArgs) ToClusterLoggingPropertiesPtrOutputWithContext(ctx context.Context) ClusterLoggingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingPropertiesOutput).ToClusterLoggingPropertiesPtrOutputWithContext(ctx)
}

// ClusterLoggingPropertiesPtrInput is an input type that accepts ClusterLoggingPropertiesArgs, ClusterLoggingPropertiesPtr and ClusterLoggingPropertiesPtrOutput values.
// You can construct a concrete instance of `ClusterLoggingPropertiesPtrInput` via:
//
//          ClusterLoggingPropertiesArgs{...}
//
//  or:
//
//          nil
type ClusterLoggingPropertiesPtrInput interface {
	pulumi.Input

	ToClusterLoggingPropertiesPtrOutput() ClusterLoggingPropertiesPtrOutput
	ToClusterLoggingPropertiesPtrOutputWithContext(context.Context) ClusterLoggingPropertiesPtrOutput
}

type clusterLoggingPropertiesPtrType ClusterLoggingPropertiesArgs

func ClusterLoggingPropertiesPtr(v *ClusterLoggingPropertiesArgs) ClusterLoggingPropertiesPtrInput {
	return (*clusterLoggingPropertiesPtrType)(v)
}

func (*clusterLoggingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterLoggingProperties)(nil)).Elem()
}

func (i *clusterLoggingPropertiesPtrType) ToClusterLoggingPropertiesPtrOutput() ClusterLoggingPropertiesPtrOutput {
	return i.ToClusterLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterLoggingPropertiesPtrType) ToClusterLoggingPropertiesPtrOutputWithContext(ctx context.Context) ClusterLoggingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterLoggingPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html
type ClusterLoggingPropertiesOutput struct{ *pulumi.OutputState }

func (ClusterLoggingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterLoggingProperties)(nil)).Elem()
}

func (o ClusterLoggingPropertiesOutput) ToClusterLoggingPropertiesOutput() ClusterLoggingPropertiesOutput {
	return o
}

func (o ClusterLoggingPropertiesOutput) ToClusterLoggingPropertiesOutputWithContext(ctx context.Context) ClusterLoggingPropertiesOutput {
	return o
}

func (o ClusterLoggingPropertiesOutput) ToClusterLoggingPropertiesPtrOutput() ClusterLoggingPropertiesPtrOutput {
	return o.ToClusterLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterLoggingPropertiesOutput) ToClusterLoggingPropertiesPtrOutputWithContext(ctx context.Context) ClusterLoggingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterLoggingProperties) *ClusterLoggingProperties {
		return &v
	}).(ClusterLoggingPropertiesPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
func (o ClusterLoggingPropertiesOutput) BucketName() pulumi.StringOutput {
	return o.ApplyT(func(v ClusterLoggingProperties) string { return v.BucketName }).(pulumi.StringOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
func (o ClusterLoggingPropertiesOutput) S3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterLoggingProperties) *string { return v.S3KeyPrefix }).(pulumi.StringPtrOutput)
}

type ClusterLoggingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterLoggingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterLoggingProperties)(nil)).Elem()
}

func (o ClusterLoggingPropertiesPtrOutput) ToClusterLoggingPropertiesPtrOutput() ClusterLoggingPropertiesPtrOutput {
	return o
}

func (o ClusterLoggingPropertiesPtrOutput) ToClusterLoggingPropertiesPtrOutputWithContext(ctx context.Context) ClusterLoggingPropertiesPtrOutput {
	return o
}

func (o ClusterLoggingPropertiesPtrOutput) Elem() ClusterLoggingPropertiesOutput {
	return o.ApplyT(func(v *ClusterLoggingProperties) ClusterLoggingProperties {
		if v != nil {
			return *v
		}
		var ret ClusterLoggingProperties
		return ret
	}).(ClusterLoggingPropertiesOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-bucketname
func (o ClusterLoggingPropertiesPtrOutput) BucketName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterLoggingProperties) *string {
		if v == nil {
			return nil
		}
		return &v.BucketName
	}).(pulumi.StringPtrOutput)
}

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-cluster-loggingproperties.html#cfn-redshift-cluster-loggingproperties-s3keyprefix
func (o ClusterLoggingPropertiesPtrOutput) S3KeyPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterLoggingProperties) *string {
		if v == nil {
			return nil
		}
		return v.S3KeyPrefix
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ClusterEndpointOutput{})
	pulumi.RegisterOutputType(ClusterEndpointPtrOutput{})
	pulumi.RegisterOutputType(ClusterLoggingPropertiesOutput{})
	pulumi.RegisterOutputType(ClusterLoggingPropertiesPtrOutput{})
}
