// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WAFv2.Inputs
{

    public sealed class WebACLNotStatementArgs : Pulumi.ResourceArgs
    {
        [Input("statement", required: true)]
        public Input<Inputs.WebACLStatementArgs> Statement { get; set; } = null!;

        public WebACLNotStatementArgs()
        {
        }
    }
}
