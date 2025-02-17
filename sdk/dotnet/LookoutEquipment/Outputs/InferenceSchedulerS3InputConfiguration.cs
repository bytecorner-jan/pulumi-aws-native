// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutEquipment.Outputs
{

    /// <summary>
    /// Specifies configuration information for the input data for the inference, including input data S3 location.
    /// </summary>
    [OutputType]
    public sealed class InferenceSchedulerS3InputConfiguration
    {
        public readonly string Bucket;
        public readonly string? Prefix;

        [OutputConstructor]
        private InferenceSchedulerS3InputConfiguration(
            string bucket,

            string? prefix)
        {
            Bucket = bucket;
            Prefix = prefix;
        }
    }
}
