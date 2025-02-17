// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Pinpoint
{
    /// <summary>
    /// Resource Type definition for AWS::Pinpoint::InAppTemplate
    /// </summary>
    [AwsNativeResourceType("aws-native:pinpoint:InAppTemplate")]
    public partial class InAppTemplate : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("content")]
        public Output<ImmutableArray<Outputs.InAppTemplateInAppMessageContent>> Content { get; private set; } = null!;

        [Output("customConfig")]
        public Output<object?> CustomConfig { get; private set; } = null!;

        [Output("layout")]
        public Output<Pulumi.AwsNative.Pinpoint.InAppTemplateLayout?> Layout { get; private set; } = null!;

        [Output("tags")]
        public Output<object?> Tags { get; private set; } = null!;

        [Output("templateDescription")]
        public Output<string?> TemplateDescription { get; private set; } = null!;

        [Output("templateName")]
        public Output<string> TemplateName { get; private set; } = null!;


        /// <summary>
        /// Create a InAppTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InAppTemplate(string name, InAppTemplateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:pinpoint:InAppTemplate", name, args ?? new InAppTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InAppTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:pinpoint:InAppTemplate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing InAppTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InAppTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InAppTemplate(name, id, options);
        }
    }

    public sealed class InAppTemplateArgs : Pulumi.ResourceArgs
    {
        [Input("content")]
        private InputList<Inputs.InAppTemplateInAppMessageContentArgs>? _content;
        public InputList<Inputs.InAppTemplateInAppMessageContentArgs> Content
        {
            get => _content ?? (_content = new InputList<Inputs.InAppTemplateInAppMessageContentArgs>());
            set => _content = value;
        }

        [Input("customConfig")]
        public Input<object>? CustomConfig { get; set; }

        [Input("layout")]
        public Input<Pulumi.AwsNative.Pinpoint.InAppTemplateLayout>? Layout { get; set; }

        [Input("tags")]
        public Input<object>? Tags { get; set; }

        [Input("templateDescription")]
        public Input<string>? TemplateDescription { get; set; }

        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public InAppTemplateArgs()
        {
        }
    }
}
