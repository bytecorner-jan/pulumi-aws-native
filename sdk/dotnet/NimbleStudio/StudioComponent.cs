// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio
{
    /// <summary>
    /// Represents a studio component which connects a non-Nimble Studio resource in your account to your studio
    /// </summary>
    [AwsNativeResourceType("aws-native:nimblestudio:StudioComponent")]
    public partial class StudioComponent : Pulumi.CustomResource
    {
        [Output("configuration")]
        public Output<Outputs.StudioComponentConfiguration?> Configuration { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The description.&lt;/p&gt;
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The EC2 security groups that control access to the studio component.&lt;/p&gt;
        /// </summary>
        [Output("ec2SecurityGroupIds")]
        public Output<ImmutableArray<string>> Ec2SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;Initialization scripts for studio components.&lt;/p&gt;
        /// </summary>
        [Output("initializationScripts")]
        public Output<ImmutableArray<Outputs.StudioComponentInitializationScript>> InitializationScripts { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The name for the studio component.&lt;/p&gt;
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;Parameters for the studio component scripts.&lt;/p&gt;
        /// </summary>
        [Output("scriptParameters")]
        public Output<ImmutableArray<Outputs.StudioComponentScriptParameterKeyValue>> ScriptParameters { get; private set; } = null!;

        [Output("studioComponentId")]
        public Output<string> StudioComponentId { get; private set; } = null!;

        /// <summary>
        /// &lt;p&gt;The studioId. &lt;/p&gt;
        /// </summary>
        [Output("studioId")]
        public Output<string> StudioId { get; private set; } = null!;

        [Output("subtype")]
        public Output<Pulumi.AwsNative.NimbleStudio.StudioComponentSubtype?> Subtype { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.StudioComponentTags?> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<Pulumi.AwsNative.NimbleStudio.StudioComponentType> Type { get; private set; } = null!;


        /// <summary>
        /// Create a StudioComponent resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public StudioComponent(string name, StudioComponentArgs args, CustomResourceOptions? options = null)
            : base("aws-native:nimblestudio:StudioComponent", name, args ?? new StudioComponentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private StudioComponent(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:nimblestudio:StudioComponent", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing StudioComponent resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static StudioComponent Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new StudioComponent(name, id, options);
        }
    }

    public sealed class StudioComponentArgs : Pulumi.ResourceArgs
    {
        [Input("configuration")]
        public Input<Inputs.StudioComponentConfigurationArgs>? Configuration { get; set; }

        /// <summary>
        /// &lt;p&gt;The description.&lt;/p&gt;
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("ec2SecurityGroupIds")]
        private InputList<string>? _ec2SecurityGroupIds;

        /// <summary>
        /// &lt;p&gt;The EC2 security groups that control access to the studio component.&lt;/p&gt;
        /// </summary>
        public InputList<string> Ec2SecurityGroupIds
        {
            get => _ec2SecurityGroupIds ?? (_ec2SecurityGroupIds = new InputList<string>());
            set => _ec2SecurityGroupIds = value;
        }

        [Input("initializationScripts")]
        private InputList<Inputs.StudioComponentInitializationScriptArgs>? _initializationScripts;

        /// <summary>
        /// &lt;p&gt;Initialization scripts for studio components.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.StudioComponentInitializationScriptArgs> InitializationScripts
        {
            get => _initializationScripts ?? (_initializationScripts = new InputList<Inputs.StudioComponentInitializationScriptArgs>());
            set => _initializationScripts = value;
        }

        /// <summary>
        /// &lt;p&gt;The name for the studio component.&lt;/p&gt;
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("scriptParameters")]
        private InputList<Inputs.StudioComponentScriptParameterKeyValueArgs>? _scriptParameters;

        /// <summary>
        /// &lt;p&gt;Parameters for the studio component scripts.&lt;/p&gt;
        /// </summary>
        public InputList<Inputs.StudioComponentScriptParameterKeyValueArgs> ScriptParameters
        {
            get => _scriptParameters ?? (_scriptParameters = new InputList<Inputs.StudioComponentScriptParameterKeyValueArgs>());
            set => _scriptParameters = value;
        }

        /// <summary>
        /// &lt;p&gt;The studioId. &lt;/p&gt;
        /// </summary>
        [Input("studioId", required: true)]
        public Input<string> StudioId { get; set; } = null!;

        [Input("subtype")]
        public Input<Pulumi.AwsNative.NimbleStudio.StudioComponentSubtype>? Subtype { get; set; }

        [Input("tags")]
        public Input<Inputs.StudioComponentTagsArgs>? Tags { get; set; }

        [Input("type", required: true)]
        public Input<Pulumi.AwsNative.NimbleStudio.StudioComponentType> Type { get; set; } = null!;

        public StudioComponentArgs()
        {
        }
    }
}
