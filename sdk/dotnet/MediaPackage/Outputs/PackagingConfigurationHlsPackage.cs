// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// An HTTP Live Streaming (HLS) packaging configuration.
    /// </summary>
    [OutputType]
    public sealed class PackagingConfigurationHlsPackage
    {
        public readonly Outputs.PackagingConfigurationHlsEncryption? Encryption;
        /// <summary>
        /// A list of HLS manifest configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.PackagingConfigurationHlsManifest> HlsManifests;
        public readonly int? SegmentDurationSeconds;
        /// <summary>
        /// When enabled, audio streams will be placed in rendition groups in the output.
        /// </summary>
        public readonly bool? UseAudioRenditionGroup;

        [OutputConstructor]
        private PackagingConfigurationHlsPackage(
            Outputs.PackagingConfigurationHlsEncryption? encryption,

            ImmutableArray<Outputs.PackagingConfigurationHlsManifest> hlsManifests,

            int? segmentDurationSeconds,

            bool? useAudioRenditionGroup)
        {
            Encryption = encryption;
            HlsManifests = hlsManifests;
            SegmentDurationSeconds = segmentDurationSeconds;
            UseAudioRenditionGroup = useAudioRenditionGroup;
        }
    }
}
