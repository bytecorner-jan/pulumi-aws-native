// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    public sealed class RuleGroupAndStatementArgs : Pulumi.ResourceArgs
    {
        [Input("statements", required: true)]
        private InputList<Inputs.RuleGroupStatementArgs>? _statements;
        public InputList<Inputs.RuleGroupStatementArgs> Statements
        {
            get => _statements ?? (_statements = new InputList<Inputs.RuleGroupStatementArgs>());
            set => _statements = value;
        }

        public RuleGroupAndStatementArgs()
        {
        }
    }
}
