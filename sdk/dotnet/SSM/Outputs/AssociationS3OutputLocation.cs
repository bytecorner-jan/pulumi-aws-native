// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM.Outputs
{

    [OutputType]
    public sealed class AssociationS3OutputLocation
    {
        public readonly string? OutputS3BucketName;
        public readonly string? OutputS3KeyPrefix;
        public readonly string? OutputS3Region;

        [OutputConstructor]
        private AssociationS3OutputLocation(
            string? outputS3BucketName,

            string? outputS3KeyPrefix,

            string? outputS3Region)
        {
            OutputS3BucketName = outputS3BucketName;
            OutputS3KeyPrefix = outputS3KeyPrefix;
            OutputS3Region = outputS3Region;
        }
    }
}
