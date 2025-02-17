// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QLDB
{
    /// <summary>
    /// Resource schema for AWS::QLDB::Stream.
    /// </summary>
    [AwsNativeResourceType("aws-native:qldb:Stream")]
    public partial class Stream : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("exclusiveEndTime")]
        public Output<string?> ExclusiveEndTime { get; private set; } = null!;

        [Output("inclusiveStartTime")]
        public Output<string> InclusiveStartTime { get; private set; } = null!;

        [Output("kinesisConfiguration")]
        public Output<Outputs.StreamKinesisConfiguration> KinesisConfiguration { get; private set; } = null!;

        [Output("ledgerName")]
        public Output<string> LedgerName { get; private set; } = null!;

        [Output("roleArn")]
        public Output<string> RoleArn { get; private set; } = null!;

        [Output("streamName")]
        public Output<string> StreamName { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.StreamTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Stream resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Stream(string name, StreamArgs args, CustomResourceOptions? options = null)
            : base("aws-native:qldb:Stream", name, args ?? new StreamArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Stream(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:qldb:Stream", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Stream resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Stream Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Stream(name, id, options);
        }
    }

    public sealed class StreamArgs : Pulumi.ResourceArgs
    {
        [Input("exclusiveEndTime")]
        public Input<string>? ExclusiveEndTime { get; set; }

        [Input("inclusiveStartTime", required: true)]
        public Input<string> InclusiveStartTime { get; set; } = null!;

        [Input("kinesisConfiguration", required: true)]
        public Input<Inputs.StreamKinesisConfigurationArgs> KinesisConfiguration { get; set; } = null!;

        [Input("ledgerName", required: true)]
        public Input<string> LedgerName { get; set; } = null!;

        [Input("roleArn", required: true)]
        public Input<string> RoleArn { get; set; } = null!;

        [Input("streamName")]
        public Input<string>? StreamName { get; set; }

        [Input("tags")]
        private InputList<Inputs.StreamTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.StreamTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.StreamTagArgs>());
            set => _tags = value;
        }

        public StreamArgs()
        {
        }
    }
}
