// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint.Inputs
{

    public sealed class CampaignLimitsArgs : Pulumi.ResourceArgs
    {
        [Input("daily")]
        public Input<int>? Daily { get; set; }

        [Input("maximumDuration")]
        public Input<int>? MaximumDuration { get; set; }

        [Input("messagesPerSecond")]
        public Input<int>? MessagesPerSecond { get; set; }

        [Input("session")]
        public Input<int>? Session { get; set; }

        [Input("total")]
        public Input<int>? Total { get; set; }

        public CampaignLimitsArgs()
        {
        }
    }
}
