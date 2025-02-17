// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EKS
{
    /// <summary>
    /// Resource Schema for AWS::EKS::FargateProfile
    /// </summary>
    [AwsNativeResourceType("aws-native:eks:FargateProfile")]
    public partial class FargateProfile : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Name of the Cluster
        /// </summary>
        [Output("clusterName")]
        public Output<string> ClusterName { get; private set; } = null!;

        /// <summary>
        /// Name of FargateProfile
        /// </summary>
        [Output("fargateProfileName")]
        public Output<string?> FargateProfileName { get; private set; } = null!;

        /// <summary>
        /// The IAM policy arn for pods
        /// </summary>
        [Output("podExecutionRoleArn")]
        public Output<string> PodExecutionRoleArn { get; private set; } = null!;

        [Output("selectors")]
        public Output<ImmutableArray<Outputs.FargateProfileSelector>> Selectors { get; private set; } = null!;

        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.FargateProfileTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a FargateProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FargateProfile(string name, FargateProfileArgs args, CustomResourceOptions? options = null)
            : base("aws-native:eks:FargateProfile", name, args ?? new FargateProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FargateProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:eks:FargateProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing FargateProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FargateProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FargateProfile(name, id, options);
        }
    }

    public sealed class FargateProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the Cluster
        /// </summary>
        [Input("clusterName", required: true)]
        public Input<string> ClusterName { get; set; } = null!;

        /// <summary>
        /// Name of FargateProfile
        /// </summary>
        [Input("fargateProfileName")]
        public Input<string>? FargateProfileName { get; set; }

        /// <summary>
        /// The IAM policy arn for pods
        /// </summary>
        [Input("podExecutionRoleArn", required: true)]
        public Input<string> PodExecutionRoleArn { get; set; } = null!;

        [Input("selectors", required: true)]
        private InputList<Inputs.FargateProfileSelectorArgs>? _selectors;
        public InputList<Inputs.FargateProfileSelectorArgs> Selectors
        {
            get => _selectors ?? (_selectors = new InputList<Inputs.FargateProfileSelectorArgs>());
            set => _selectors = value;
        }

        [Input("subnets")]
        private InputList<string>? _subnets;
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        [Input("tags")]
        private InputList<Inputs.FargateProfileTagArgs>? _tags;

        /// <summary>
        /// An array of key-value pairs to apply to this resource.
        /// </summary>
        public InputList<Inputs.FargateProfileTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.FargateProfileTagArgs>());
            set => _tags = value;
        }

        public FargateProfileArgs()
        {
        }
    }
}
