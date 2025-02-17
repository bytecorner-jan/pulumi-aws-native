// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleSqsActionArgs : Pulumi.ResourceArgs
    {
        [Input("queueUrl", required: true)]
        public Input<string> QueueUrl { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("useBase64")]
        public Input<bool>? UseBase64 { get; set; }

        public TopicRuleSqsActionArgs()
        {
        }
    }
}
