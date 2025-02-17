// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Outputs
{

    [OutputType]
    public sealed class ClusterScalingRule
    {
        public readonly Outputs.ClusterScalingAction Action;
        public readonly string? Description;
        public readonly string Name;
        public readonly Outputs.ClusterScalingTrigger Trigger;

        [OutputConstructor]
        private ClusterScalingRule(
            Outputs.ClusterScalingAction action,

            string? description,

            string name,

            Outputs.ClusterScalingTrigger trigger)
        {
            Action = action;
            Description = description;
            Name = name;
            Trigger = trigger;
        }
    }
}
