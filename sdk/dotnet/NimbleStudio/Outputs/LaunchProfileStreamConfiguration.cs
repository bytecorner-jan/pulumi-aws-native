// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Outputs
{

    /// <summary>
    /// &lt;p&gt;A configuration for a streaming session.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class LaunchProfileStreamConfiguration
    {
        public readonly Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingClipboardMode ClipboardMode;
        /// <summary>
        /// &lt;p&gt;The EC2 instance types that users can select from when launching a streaming session with this launch profile.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingInstanceType> Ec2InstanceTypes;
        /// <summary>
        /// &lt;p&gt;The length of time, in minutes, that a streaming session can run. After this point, Nimble Studio automatically terminates the session.&lt;/p&gt;
        /// </summary>
        public readonly double? MaxSessionLengthInMinutes;
        /// <summary>
        /// &lt;p&gt;The streaming images that users can select from when launching a streaming session with this launch profile.&lt;/p&gt;
        /// </summary>
        public readonly ImmutableArray<string> StreamingImageIds;

        [OutputConstructor]
        private LaunchProfileStreamConfiguration(
            Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingClipboardMode clipboardMode,

            ImmutableArray<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingInstanceType> ec2InstanceTypes,

            double? maxSessionLengthInMinutes,

            ImmutableArray<string> streamingImageIds)
        {
            ClipboardMode = clipboardMode;
            Ec2InstanceTypes = ec2InstanceTypes;
            MaxSessionLengthInMinutes = maxSessionLengthInMinutes;
            StreamingImageIds = streamingImageIds;
        }
    }
}
