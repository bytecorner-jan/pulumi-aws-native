// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFRegional.Outputs
{

    [OutputType]
    public sealed class WebACLRule
    {
        public readonly Outputs.WebACLAction Action;
        public readonly int Priority;
        public readonly string RuleId;

        [OutputConstructor]
        private WebACLRule(
            Outputs.WebACLAction action,

            int priority,

            string ruleId)
        {
            Action = action;
            Priority = priority;
            RuleId = ruleId;
        }
    }
}
