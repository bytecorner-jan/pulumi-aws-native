// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisAnalyticsV2.Inputs
{

    public sealed class ApplicationReferenceDataSourceS3ReferenceDataSourceArgs : Pulumi.ResourceArgs
    {
        [Input("bucketARN", required: true)]
        public Input<string> BucketARN { get; set; } = null!;

        [Input("fileKey", required: true)]
        public Input<string> FileKey { get; set; } = null!;

        public ApplicationReferenceDataSourceS3ReferenceDataSourceArgs()
        {
        }
    }
}
