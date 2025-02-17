// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Outputs
{

    [OutputType]
    public sealed class ChannelUdpOutputSettings
    {
        public readonly int? BufferMsec;
        public readonly Outputs.ChannelUdpContainerSettings? ContainerSettings;
        public readonly Outputs.ChannelOutputLocationRef? Destination;
        public readonly Outputs.ChannelFecOutputSettings? FecOutputSettings;

        [OutputConstructor]
        private ChannelUdpOutputSettings(
            int? bufferMsec,

            Outputs.ChannelUdpContainerSettings? containerSettings,

            Outputs.ChannelOutputLocationRef? destination,

            Outputs.ChannelFecOutputSettings? fecOutputSettings)
        {
            BufferMsec = bufferMsec;
            ContainerSettings = containerSettings;
            Destination = destination;
            FecOutputSettings = fecOutputSettings;
        }
    }
}
