// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html
    /// </summary>
    public sealed class AppImageConfigKernelSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html#cfn-sagemaker-appimageconfig-kernelspec-displayname
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-kernelspec.html#cfn-sagemaker-appimageconfig-kernelspec-name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public AppImageConfigKernelSpecArgs()
        {
        }
    }
}
