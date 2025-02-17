// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    /// <summary>
    /// Rule of RuleGroup that contains condition and action.
    /// </summary>
    public sealed class RuleGroupRuleArgs : Pulumi.ResourceArgs
    {
        [Input("action")]
        public Input<Inputs.RuleGroupRuleActionArgs>? Action { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("ruleLabels")]
        private InputList<Inputs.RuleGroupLabelArgs>? _ruleLabels;

        /// <summary>
        /// Collection of Rule Labels.
        /// </summary>
        public InputList<Inputs.RuleGroupLabelArgs> RuleLabels
        {
            get => _ruleLabels ?? (_ruleLabels = new InputList<Inputs.RuleGroupLabelArgs>());
            set => _ruleLabels = value;
        }

        [Input("statement", required: true)]
        public Input<Inputs.RuleGroupStatementArgs> Statement { get; set; } = null!;

        [Input("visibilityConfig", required: true)]
        public Input<Inputs.RuleGroupVisibilityConfigArgs> VisibilityConfig { get; set; } = null!;

        public RuleGroupRuleArgs()
        {
        }
    }
}
