// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaLive.Inputs
{

    public sealed class InputSecurityGroupInputWhitelistRuleCidrArgs : Pulumi.ResourceArgs
    {
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        public InputSecurityGroupInputWhitelistRuleCidrArgs()
        {
        }
    }
}
