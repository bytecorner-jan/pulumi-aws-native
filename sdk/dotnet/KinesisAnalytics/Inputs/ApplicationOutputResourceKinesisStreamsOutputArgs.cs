// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalytics.Inputs
{

    public sealed class ApplicationOutputResourceKinesisStreamsOutputArgs : Pulumi.ResourceArgs
    {
        [Input("resourceARN", required: true)]
        public Input<string> ResourceARN { get; set; } = null!;

        [Input("roleARN", required: true)]
        public Input<string> RoleARN { get; set; } = null!;

        public ApplicationOutputResourceKinesisStreamsOutputArgs()
        {
        }
    }
}
