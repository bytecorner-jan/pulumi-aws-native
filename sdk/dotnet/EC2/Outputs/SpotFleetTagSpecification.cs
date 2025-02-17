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
    public sealed class SpotFleetTagSpecification
    {
        public readonly Pulumi.AwsNative.EC2.SpotFleetTagSpecificationResourceType? ResourceType;
        public readonly ImmutableArray<Outputs.SpotFleetTag> Tags;

        [OutputConstructor]
        private SpotFleetTagSpecification(
            Pulumi.AwsNative.EC2.SpotFleetTagSpecificationResourceType? resourceType,

            ImmutableArray<Outputs.SpotFleetTag> tags)
        {
            ResourceType = resourceType;
            Tags = tags;
        }
    }
}
