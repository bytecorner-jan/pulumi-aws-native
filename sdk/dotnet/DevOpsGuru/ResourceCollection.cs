// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DevOpsGuru
{
    /// <summary>
    /// This resource schema represents the ResourceCollection resource in the Amazon DevOps Guru.
    /// </summary>
    [AwsNativeResourceType("aws-native:devopsguru:ResourceCollection")]
    public partial class ResourceCollection : Pulumi.CustomResource
    {
        [Output("resourceCollectionFilter")]
        public Output<Outputs.ResourceCollectionFilter> ResourceCollectionFilter { get; private set; } = null!;

        /// <summary>
        /// The type of ResourceCollection
        /// </summary>
        [Output("resourceCollectionType")]
        public Output<Pulumi.AwsNative.DevOpsGuru.ResourceCollectionType> ResourceCollectionType { get; private set; } = null!;


        /// <summary>
        /// Create a ResourceCollection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ResourceCollection(string name, ResourceCollectionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:devopsguru:ResourceCollection", name, args ?? new ResourceCollectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ResourceCollection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:devopsguru:ResourceCollection", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ResourceCollection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ResourceCollection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ResourceCollection(name, id, options);
        }
    }

    public sealed class ResourceCollectionArgs : Pulumi.ResourceArgs
    {
        [Input("resourceCollectionFilter", required: true)]
        public Input<Inputs.ResourceCollectionFilterArgs> ResourceCollectionFilter { get; set; } = null!;

        public ResourceCollectionArgs()
        {
        }
    }
}
