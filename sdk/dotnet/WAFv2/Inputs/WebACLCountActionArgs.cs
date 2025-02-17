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
    /// Allow traffic towards application.
    /// </summary>
    public sealed class WebACLCountActionArgs : Pulumi.ResourceArgs
    {
        [Input("customRequestHandling")]
        public Input<Inputs.WebACLCustomRequestHandlingArgs>? CustomRequestHandling { get; set; }

        public WebACLCountActionArgs()
        {
        }
    }
}
