// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html
    /// </summary>
    [OutputType]
    public sealed class ServiceImageRepository
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html#cfn-apprunner-service-imagerepository-imageconfiguration
        /// </summary>
        public readonly Outputs.ServiceImageConfiguration? ImageConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html#cfn-apprunner-service-imagerepository-imageidentifier
        /// </summary>
        public readonly string ImageIdentifier;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-imagerepository.html#cfn-apprunner-service-imagerepository-imagerepositorytype
        /// </summary>
        public readonly string ImageRepositoryType;

        [OutputConstructor]
        private ServiceImageRepository(
            Outputs.ServiceImageConfiguration? imageConfiguration,

            string imageIdentifier,

            string imageRepositoryType)
        {
            ImageConfiguration = imageConfiguration;
            ImageIdentifier = imageIdentifier;
            ImageRepositoryType = imageRepositoryType;
        }
    }
}
