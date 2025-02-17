// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CloudFront
{
    /// <summary>
    /// Resource Type definition for AWS::CloudFront::RealtimeLogConfig
    /// </summary>
    [AwsNativeResourceType("aws-native:cloudfront:RealtimeLogConfig")]
    public partial class RealtimeLogConfig : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("endPoints")]
        public Output<ImmutableArray<Outputs.RealtimeLogConfigEndPoint>> EndPoints { get; private set; } = null!;

        [Output("fields")]
        public Output<ImmutableArray<string>> Fields { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("samplingRate")]
        public Output<double> SamplingRate { get; private set; } = null!;


        /// <summary>
        /// Create a RealtimeLogConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RealtimeLogConfig(string name, RealtimeLogConfigArgs args, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:RealtimeLogConfig", name, args ?? new RealtimeLogConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RealtimeLogConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloudfront:RealtimeLogConfig", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing RealtimeLogConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RealtimeLogConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RealtimeLogConfig(name, id, options);
        }
    }

    public sealed class RealtimeLogConfigArgs : Pulumi.ResourceArgs
    {
        [Input("endPoints", required: true)]
        private InputList<Inputs.RealtimeLogConfigEndPointArgs>? _endPoints;
        public InputList<Inputs.RealtimeLogConfigEndPointArgs> EndPoints
        {
            get => _endPoints ?? (_endPoints = new InputList<Inputs.RealtimeLogConfigEndPointArgs>());
            set => _endPoints = value;
        }

        [Input("fields", required: true)]
        private InputList<string>? _fields;
        public InputList<string> Fields
        {
            get => _fields ?? (_fields = new InputList<string>());
            set => _fields = value;
        }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("samplingRate", required: true)]
        public Input<double> SamplingRate { get; set; } = null!;

        public RealtimeLogConfigArgs()
        {
        }
    }
}
