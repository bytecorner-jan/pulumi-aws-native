// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class TargetGroupTargetDescription
    {
        public readonly string? AvailabilityZone;
        public readonly string Id;
        public readonly int? Port;

        [OutputConstructor]
        private TargetGroupTargetDescription(
            string? availabilityZone,

            string id,

            int? port)
        {
            AvailabilityZone = availabilityZone;
            Id = id;
            Port = port;
        }
    }
}
