// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResourceGroups.Inputs
{

    public sealed class GroupResourceQueryArgs : Pulumi.ResourceArgs
    {
        [Input("query")]
        public Input<Inputs.GroupQueryArgs>? Query { get; set; }

        [Input("type")]
        public Input<Pulumi.AwsNative.ResourceGroups.GroupResourceQueryType>? Type { get; set; }

        public GroupResourceQueryArgs()
        {
        }
    }
}
