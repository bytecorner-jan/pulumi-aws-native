// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EMR
{
    /// <summary>
    /// Resource Type definition for AWS::EMR::SecurityConfiguration
    /// </summary>
    [Obsolete(@"SecurityConfiguration is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:emr:SecurityConfiguration")]
    public partial class SecurityConfiguration : Pulumi.CustomResource
    {
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("securityConfiguration")]
        public Output<object> SecurityConfigurationValue { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityConfiguration(string name, SecurityConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:emr:SecurityConfiguration", name, args ?? new SecurityConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:emr:SecurityConfiguration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing SecurityConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityConfiguration(name, id, options);
        }
    }

    public sealed class SecurityConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("securityConfiguration", required: true)]
        public Input<object> SecurityConfigurationValue { get; set; } = null!;

        public SecurityConfigurationArgs()
        {
        }
    }
}
