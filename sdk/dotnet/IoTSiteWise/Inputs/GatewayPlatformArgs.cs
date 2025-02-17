// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTSiteWise.Inputs
{

    /// <summary>
    /// Contains a gateway's platform information.
    /// </summary>
    public sealed class GatewayPlatformArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A gateway that runs on AWS IoT Greengrass.
        /// </summary>
        [Input("greengrass", required: true)]
        public Input<Inputs.GatewayGreengrassArgs> Greengrass { get; set; } = null!;

        public GatewayPlatformArgs()
        {
        }
    }
}
