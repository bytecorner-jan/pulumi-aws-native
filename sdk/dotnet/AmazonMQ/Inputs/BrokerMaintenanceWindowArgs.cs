// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMQ.Inputs
{

    public sealed class BrokerMaintenanceWindowArgs : Pulumi.ResourceArgs
    {
        [Input("dayOfWeek", required: true)]
        public Input<string> DayOfWeek { get; set; } = null!;

        [Input("timeOfDay", required: true)]
        public Input<string> TimeOfDay { get; set; } = null!;

        [Input("timeZone", required: true)]
        public Input<string> TimeZone { get; set; } = null!;

        public BrokerMaintenanceWindowArgs()
        {
        }
    }
}
