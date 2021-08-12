// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KMS
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
    /// </summary>
    [AwsNativeResourceType("aws-native:KMS:Key")]
    public partial class Key : Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
        /// </summary>
        [Output("enableKeyRotation")]
        public Output<bool?> EnableKeyRotation { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        [Output("keyId")]
        public Output<string> KeyId { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
        /// </summary>
        [Output("keyPolicy")]
        public Output<Union<System.Text.Json.JsonElement, string>> KeyPolicy { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyspec
        /// </summary>
        [Output("keySpec")]
        public Output<string?> KeySpec { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
        /// </summary>
        [Output("keyUsage")]
        public Output<string?> KeyUsage { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-multiregion
        /// </summary>
        [Output("multiRegion")]
        public Output<bool?> MultiRegion { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
        /// </summary>
        [Output("pendingWindowInDays")]
        public Output<int?> PendingWindowInDays { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Pulumi.AwsNative.Outputs.Tag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Key resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Key(string name, KeyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:KMS:Key", name, args ?? new KeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Key(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:KMS:Key", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Key resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Key Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Key(name, id, options);
        }
    }

    public sealed class KeyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
        /// </summary>
        [Input("enableKeyRotation")]
        public Input<bool>? EnableKeyRotation { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
        /// </summary>
        [Input("keyPolicy", required: true)]
        public InputUnion<System.Text.Json.JsonElement, string> KeyPolicy { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyspec
        /// </summary>
        [Input("keySpec")]
        public Input<string>? KeySpec { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
        /// </summary>
        [Input("keyUsage")]
        public Input<string>? KeyUsage { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-multiregion
        /// </summary>
        [Input("multiRegion")]
        public Input<bool>? MultiRegion { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
        /// </summary>
        [Input("pendingWindowInDays")]
        public Input<int>? PendingWindowInDays { get; set; }

        [Input("tags")]
        private InputList<Pulumi.AwsNative.Inputs.TagArgs>? _tags;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
        /// </summary>
        public InputList<Pulumi.AwsNative.Inputs.TagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Pulumi.AwsNative.Inputs.TagArgs>());
            set => _tags = value;
        }

        public KeyArgs()
        {
        }
    }
}
