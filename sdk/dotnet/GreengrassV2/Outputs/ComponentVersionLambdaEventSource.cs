// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GreengrassV2.Outputs
{

    [OutputType]
    public sealed class ComponentVersionLambdaEventSource
    {
        public readonly string? Topic;
        public readonly Pulumi.AwsNative.GreengrassV2.ComponentVersionLambdaEventSourceType? Type;

        [OutputConstructor]
        private ComponentVersionLambdaEventSource(
            string? topic,

            Pulumi.AwsNative.GreengrassV2.ComponentVersionLambdaEventSourceType? type)
        {
            Topic = topic;
            Type = type;
        }
    }
}
