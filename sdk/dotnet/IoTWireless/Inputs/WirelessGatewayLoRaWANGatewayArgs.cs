// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTWireless.Inputs
{

    public sealed class WirelessGatewayLoRaWANGatewayArgs : Pulumi.ResourceArgs
    {
        [Input("gatewayEui", required: true)]
        public Input<string> GatewayEui { get; set; } = null!;

        [Input("rfRegion", required: true)]
        public Input<string> RfRegion { get; set; } = null!;

        public WirelessGatewayLoRaWANGatewayArgs()
        {
        }
    }
}
