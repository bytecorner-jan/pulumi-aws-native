// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class ChannelUdpGroupSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("inputLossAction")]
        public Input<string>? InputLossAction { get; set; }

        [Input("timedMetadataId3Frame")]
        public Input<string>? TimedMetadataId3Frame { get; set; }

        [Input("timedMetadataId3Period")]
        public Input<int>? TimedMetadataId3Period { get; set; }

        public ChannelUdpGroupSettingsArgs()
        {
        }
    }
}
