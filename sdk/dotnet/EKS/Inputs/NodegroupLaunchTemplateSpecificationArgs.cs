// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EKS.Inputs
{

    public sealed class NodegroupLaunchTemplateSpecificationArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public NodegroupLaunchTemplateSpecificationArgs()
        {
        }
    }
}
