// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    public sealed class WebACLLabelMatchStatementArgs : Pulumi.ResourceArgs
    {
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        [Input("scope", required: true)]
        public Input<Pulumi.AwsNative.WAFv2.WebACLLabelMatchScope> Scope { get; set; } = null!;

        public WebACLLabelMatchStatementArgs()
        {
        }
    }
}
