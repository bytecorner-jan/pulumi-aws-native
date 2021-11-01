// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    [OutputType]
    public sealed class SpotFleetLaunchTemplateOverrides
    {
        public readonly string? AvailabilityZone;
        public readonly Outputs.SpotFleetInstanceRequirementsRequest? InstanceRequirements;
        public readonly string? InstanceType;
        public readonly string? SpotPrice;
        public readonly string? SubnetId;
        public readonly double? WeightedCapacity;

        [OutputConstructor]
        private SpotFleetLaunchTemplateOverrides(
            string? availabilityZone,

            Outputs.SpotFleetInstanceRequirementsRequest? instanceRequirements,

            string? instanceType,

            string? spotPrice,

            string? subnetId,

            double? weightedCapacity)
        {
            AvailabilityZone = availabilityZone;
            InstanceRequirements = instanceRequirements;
            InstanceType = instanceType;
            SpotPrice = spotPrice;
            SubnetId = subnetId;
            WeightedCapacity = weightedCapacity;
        }
    }
}
