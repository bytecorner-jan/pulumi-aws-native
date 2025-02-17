// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IVS.Inputs
{

    /// <summary>
    /// Recording S3 Destination Configuration.
    /// </summary>
    public sealed class RecordingConfigurationS3DestinationConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        public RecordingConfigurationS3DestinationConfigurationArgs()
        {
        }
    }
}
