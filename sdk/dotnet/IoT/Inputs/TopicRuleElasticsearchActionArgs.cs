// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleElasticsearchActionArgs : Pulumi.ResourceArgs
    {
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("index", required: true)]
        public Input<string> Index { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TopicRuleElasticsearchActionArgs()
        {
        }
    }
}
