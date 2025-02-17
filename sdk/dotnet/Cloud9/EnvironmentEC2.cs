// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cloud9
{
    /// <summary>
    /// Resource Type definition for AWS::Cloud9::EnvironmentEC2
    /// </summary>
    [Obsolete(@"EnvironmentEC2 is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:cloud9:EnvironmentEC2")]
    public partial class EnvironmentEC2 : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("automaticStopTimeMinutes")]
        public Output<int?> AutomaticStopTimeMinutes { get; private set; } = null!;

        [Output("connectionType")]
        public Output<string?> ConnectionType { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("imageId")]
        public Output<string?> ImageId { get; private set; } = null!;

        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("ownerArn")]
        public Output<string?> OwnerArn { get; private set; } = null!;

        [Output("repositories")]
        public Output<ImmutableArray<Outputs.EnvironmentEC2Repository>> Repositories { get; private set; } = null!;

        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.EnvironmentEC2Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EnvironmentEC2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvironmentEC2(string name, EnvironmentEC2Args args, CustomResourceOptions? options = null)
            : base("aws-native:cloud9:EnvironmentEC2", name, args ?? new EnvironmentEC2Args(), MakeResourceOptions(options, ""))
        {
        }

        private EnvironmentEC2(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:cloud9:EnvironmentEC2", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing EnvironmentEC2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvironmentEC2 Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EnvironmentEC2(name, id, options);
        }
    }

    public sealed class EnvironmentEC2Args : Pulumi.ResourceArgs
    {
        [Input("automaticStopTimeMinutes")]
        public Input<int>? AutomaticStopTimeMinutes { get; set; }

        [Input("connectionType")]
        public Input<string>? ConnectionType { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("imageId")]
        public Input<string>? ImageId { get; set; }

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerArn")]
        public Input<string>? OwnerArn { get; set; }

        [Input("repositories")]
        private InputList<Inputs.EnvironmentEC2RepositoryArgs>? _repositories;
        public InputList<Inputs.EnvironmentEC2RepositoryArgs> Repositories
        {
            get => _repositories ?? (_repositories = new InputList<Inputs.EnvironmentEC2RepositoryArgs>());
            set => _repositories = value;
        }

        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.EnvironmentEC2TagArgs>? _tags;
        public InputList<Inputs.EnvironmentEC2TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.EnvironmentEC2TagArgs>());
            set => _tags = value;
        }

        public EnvironmentEC2Args()
        {
        }
    }
}
