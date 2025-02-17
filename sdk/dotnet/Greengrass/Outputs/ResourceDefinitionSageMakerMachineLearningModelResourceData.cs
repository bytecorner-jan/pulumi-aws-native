// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Greengrass.Outputs
{

    [OutputType]
    public sealed class ResourceDefinitionSageMakerMachineLearningModelResourceData
    {
        public readonly string DestinationPath;
        public readonly Outputs.ResourceDefinitionResourceDownloadOwnerSetting? OwnerSetting;
        public readonly string SageMakerJobArn;

        [OutputConstructor]
        private ResourceDefinitionSageMakerMachineLearningModelResourceData(
            string destinationPath,

            Outputs.ResourceDefinitionResourceDownloadOwnerSetting? ownerSetting,

            string sageMakerJobArn)
        {
            DestinationPath = destinationPath;
            OwnerSetting = ownerSetting;
            SageMakerJobArn = sageMakerJobArn;
        }
    }
}
