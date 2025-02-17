// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DataSync.Inputs
{

    /// <summary>
    /// The mount options used by DataSync to access the SMB server.
    /// </summary>
    public sealed class LocationSMBMountOptionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The specific SMB version that you want DataSync to use to mount your SMB share.
        /// </summary>
        [Input("version")]
        public Input<Pulumi.AwsNative.DataSync.LocationSMBMountOptionsVersion>? Version { get; set; }

        public LocationSMBMountOptionsArgs()
        {
        }
    }
}
