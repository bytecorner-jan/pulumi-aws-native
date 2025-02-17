// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift
{
    /// <summary>
    /// Resource Type definition for AWS::GameLift::Script
    /// </summary>
    [Obsolete(@"Script is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")]
    [AwsNativeResourceType("aws-native:gamelift:Script")]
    public partial class Script : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("storageLocation")]
        public Output<Outputs.ScriptS3Location> StorageLocation { get; private set; } = null!;

        [Output("version")]
        public Output<string?> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Script resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Script(string name, ScriptArgs args, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:Script", name, args ?? new ScriptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Script(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:gamelift:Script", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Script resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Script Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Script(name, id, options);
        }
    }

    public sealed class ScriptArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("storageLocation", required: true)]
        public Input<Inputs.ScriptS3LocationArgs> StorageLocation { get; set; } = null!;

        [Input("version")]
        public Input<string>? Version { get; set; }

        public ScriptArgs()
        {
        }
    }
}
