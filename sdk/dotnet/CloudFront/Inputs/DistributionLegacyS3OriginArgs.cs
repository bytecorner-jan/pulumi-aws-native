// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront.Inputs
{

    public sealed class DistributionLegacyS3OriginArgs : Pulumi.ResourceArgs
    {
        [Input("dNSName", required: true)]
        public Input<string> DNSName { get; set; } = null!;

        [Input("originAccessIdentity")]
        public Input<string>? OriginAccessIdentity { get; set; }

        public DistributionLegacyS3OriginArgs()
        {
        }
    }
}
