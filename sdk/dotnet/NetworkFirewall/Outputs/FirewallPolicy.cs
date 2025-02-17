// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class FirewallPolicy
    {
        public readonly ImmutableArray<string> StatefulDefaultActions;
        public readonly Outputs.FirewallPolicyStatefulEngineOptions? StatefulEngineOptions;
        public readonly ImmutableArray<Outputs.FirewallPolicyStatefulRuleGroupReference> StatefulRuleGroupReferences;
        public readonly ImmutableArray<Outputs.FirewallPolicyCustomAction> StatelessCustomActions;
        public readonly ImmutableArray<string> StatelessDefaultActions;
        public readonly ImmutableArray<string> StatelessFragmentDefaultActions;
        public readonly ImmutableArray<Outputs.FirewallPolicyStatelessRuleGroupReference> StatelessRuleGroupReferences;

        [OutputConstructor]
        private FirewallPolicy(
            ImmutableArray<string> statefulDefaultActions,

            Outputs.FirewallPolicyStatefulEngineOptions? statefulEngineOptions,

            ImmutableArray<Outputs.FirewallPolicyStatefulRuleGroupReference> statefulRuleGroupReferences,

            ImmutableArray<Outputs.FirewallPolicyCustomAction> statelessCustomActions,

            ImmutableArray<string> statelessDefaultActions,

            ImmutableArray<string> statelessFragmentDefaultActions,

            ImmutableArray<Outputs.FirewallPolicyStatelessRuleGroupReference> statelessRuleGroupReferences)
        {
            StatefulDefaultActions = statefulDefaultActions;
            StatefulEngineOptions = statefulEngineOptions;
            StatefulRuleGroupReferences = statefulRuleGroupReferences;
            StatelessCustomActions = statelessCustomActions;
            StatelessDefaultActions = statelessDefaultActions;
            StatelessFragmentDefaultActions = statelessFragmentDefaultActions;
            StatelessRuleGroupReferences = statelessRuleGroupReferences;
        }
    }
}
