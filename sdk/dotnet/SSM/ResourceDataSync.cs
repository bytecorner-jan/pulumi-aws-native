// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSM
{
    /// <summary>
    /// Resource Type definition for AWS::SSM::ResourceDataSync
    /// </summary>
    [AwsNativeResourceType("aws-native:ssm:ResourceDataSync")]
    public partial class ResourceDataSync : Pulumi.CustomResource
    {
        [Output("bucketName")]
        public Output<string?> BucketName { get; private set; } = null!;

        [Output("bucketPrefix")]
        public Output<string?> BucketPrefix { get; private set; } = null!;

        [Output("bucketRegion")]
        public Output<string?> BucketRegion { get; private set; } = null!;

        [Output("kMSKeyArn")]
        public Output<string?> KMSKeyArn { get; private set; } = null!;

        [Output("s3Destination")]
        public Output<Outputs.ResourceDataSyncS3Destination?> S3Destination { get; private set; } = null!;

        [Output("syncFormat")]
        public Output<string?> SyncFormat { get; private set; } = null!;

        [Output("syncName")]
        public Output<string> SyncName { get; private set; } = null!;

        [Output("syncSource")]
        public Output<Outputs.ResourceDataSyncSyncSource?> SyncSource { get; private set; } = null!;

        [Output("syncType")]
        public Output<string?> SyncType { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceDataSync resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceDataSync(string name, ResourceDataSyncArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:ssm:ResourceDataSync", name, args ?? new ResourceDataSyncArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceDataSync(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ssm:ResourceDataSync", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ResourceDataSync resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceDataSync Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceDataSync(name, id, options);
        }
    }

    public sealed class ResourceDataSyncArgs : Pulumi.ResourceArgs
    {
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        [Input("bucketPrefix")]
        public Input<string>? BucketPrefix { get; set; }

        [Input("bucketRegion")]
        public Input<string>? BucketRegion { get; set; }

        [Input("kMSKeyArn")]
        public Input<string>? KMSKeyArn { get; set; }

        [Input("s3Destination")]
        public Input<Inputs.ResourceDataSyncS3DestinationArgs>? S3Destination { get; set; }

        [Input("syncFormat")]
        public Input<string>? SyncFormat { get; set; }

        [Input("syncSource")]
        public Input<Inputs.ResourceDataSyncSyncSourceArgs>? SyncSource { get; set; }

        [Input("syncType")]
        public Input<string>? SyncType { get; set; }

        public ResourceDataSyncArgs()
        {
        }
    }
}
