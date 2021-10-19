// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    /// <summary>
    /// &lt;p&gt;A configuration for a streaming session.&lt;/p&gt;
    /// </summary>
    public sealed class LaunchProfileStreamConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("clipboardMode", required: true)]
        public Input<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingClipboardMode> ClipboardMode { get; set; } = null!;

        [Input("ec2InstanceTypes", required: true)]
        private InputList<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingInstanceType>? _ec2InstanceTypes;

        /// <summary>
        /// &lt;p&gt;The EC2 instance types that users can select from when launching a streaming session with this launch profile.&lt;/p&gt;
        /// </summary>
        public InputList<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingInstanceType> Ec2InstanceTypes
        {
            get => _ec2InstanceTypes ?? (_ec2InstanceTypes = new InputList<Pulumi.AwsNative.NimbleStudio.LaunchProfileStreamingInstanceType>());
            set => _ec2InstanceTypes = value;
        }

        /// <summary>
        /// &lt;p&gt;The length of time, in minutes, that a streaming session can run. After this point, Nimble Studio automatically terminates the session.&lt;/p&gt;
        /// </summary>
        [Input("maxSessionLengthInMinutes")]
        public Input<double>? MaxSessionLengthInMinutes { get; set; }

        [Input("streamingImageIds", required: true)]
        private InputList<string>? _streamingImageIds;

        /// <summary>
        /// &lt;p&gt;The streaming images that users can select from when launching a streaming session with this launch profile.&lt;/p&gt;
        /// </summary>
        public InputList<string> StreamingImageIds
        {
            get => _streamingImageIds ?? (_streamingImageIds = new InputList<string>());
            set => _streamingImageIds = value;
        }

        public LaunchProfileStreamConfigurationArgs()
        {
        }
    }
}
