// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    public sealed class ConfigDataflowEndpointConfigArgs : Pulumi.ResourceArgs
    {
        [Input("dataflowEndpointName")]
        public Input<string>? DataflowEndpointName { get; set; }

        [Input("dataflowEndpointRegion")]
        public Input<string>? DataflowEndpointRegion { get; set; }

        public ConfigDataflowEndpointConfigArgs()
        {
        }
    }
}
