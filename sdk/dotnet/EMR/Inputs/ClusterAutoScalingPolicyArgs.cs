// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR.Inputs
{

    public sealed class ClusterAutoScalingPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("constraints", required: true)]
        public Input<Inputs.ClusterScalingConstraintsArgs> Constraints { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.ClusterScalingRuleArgs>? _rules;
        public InputList<Inputs.ClusterScalingRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ClusterScalingRuleArgs>());
            set => _rules = value;
        }

        public ClusterAutoScalingPolicyArgs()
        {
        }
    }
}
