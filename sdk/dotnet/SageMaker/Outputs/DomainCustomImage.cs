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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customimage.html
    /// </summary>
    [OutputType]
    public sealed class DomainCustomImage
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customimage.html#cfn-sagemaker-domain-customimage-appimageconfigname
        /// </summary>
        public readonly string AppImageConfigName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customimage.html#cfn-sagemaker-domain-customimage-imagename
        /// </summary>
        public readonly string ImageName;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-customimage.html#cfn-sagemaker-domain-customimage-imageversionnumber
        /// </summary>
        public readonly int? ImageVersionNumber;

        [OutputConstructor]
        private DomainCustomImage(
            string appImageConfigName,

            string imageName,

            int? imageVersionNumber)
        {
            AppImageConfigName = appImageConfigName;
            ImageName = imageName;
            ImageVersionNumber = imageVersionNumber;
        }
    }
}
