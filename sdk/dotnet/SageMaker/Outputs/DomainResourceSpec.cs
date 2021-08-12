// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-resourcespec.html
    /// </summary>
    [OutputType]
    public sealed class DomainResourceSpec
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-resourcespec.html#cfn-sagemaker-domain-resourcespec-instancetype
        /// </summary>
        public readonly string? InstanceType;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-resourcespec.html#cfn-sagemaker-domain-resourcespec-sagemakerimagearn
        /// </summary>
        public readonly string? SageMakerImageArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-resourcespec.html#cfn-sagemaker-domain-resourcespec-sagemakerimageversionarn
        /// </summary>
        public readonly string? SageMakerImageVersionArn;

        [OutputConstructor]
        private DomainResourceSpec(
            string? instanceType,

            string? sageMakerImageArn,

            string? sageMakerImageVersionArn)
        {
            InstanceType = instanceType;
            SageMakerImageArn = sageMakerImageArn;
            SageMakerImageVersionArn = sageMakerImageVersionArn;
        }
    }
}
