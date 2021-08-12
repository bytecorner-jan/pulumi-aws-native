// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html
    /// </summary>
    [AwsNativeResourceType("aws-native:QuickSight:Template")]
    public partial class Template : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-awsaccountid
        /// </summary>
        [Output("awsAccountId")]
        public Output<string> AwsAccountId { get; private set; } = null!;

        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-name
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-permissions
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<Outputs.TemplateResourcePermission>> Permissions { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-sourceentity
        /// </summary>
        [Output("sourceEntity")]
        public Output<Outputs.TemplateTemplateSourceEntity?> SourceEntity { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-templateid
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-versiondescription
        /// </summary>
        [Output("versionDescription")]
        public Output<string?> VersionDescription { get; private set; } = null!;


        /// <summary>
        /// Create a Template resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Template(string name, TemplateArgs args, CustomResourceOptions? options = null)
            : base("aws-native:QuickSight:Template", name, args ?? new TemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Template(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:QuickSight:Template", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Template resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Template Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Template(name, id, options);
        }
    }

    public sealed class TemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-awsaccountid
        /// </summary>
        [Input("awsAccountId", required: true)]
        public Input<string> AwsAccountId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissions")]
        private InputList<Inputs.TemplateResourcePermissionArgs>? _permissions;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-permissions
        /// </summary>
        public InputList<Inputs.TemplateResourcePermissionArgs> Permissions
        {
            get => _permissions ?? (_permissions = new InputList<Inputs.TemplateResourcePermissionArgs>());
            set => _permissions = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-sourceentity
        /// </summary>
        [Input("sourceEntity")]
        public Input<Inputs.TemplateTemplateSourceEntityArgs>? SourceEntity { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-templateid
        /// </summary>
        [Input("templateId", required: true)]
        public Input<string> TemplateId { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-template.html#cfn-quicksight-template-versiondescription
        /// </summary>
        [Input("versionDescription")]
        public Input<string>? VersionDescription { get; set; }

        public TemplateArgs()
        {
        }
    }
}
