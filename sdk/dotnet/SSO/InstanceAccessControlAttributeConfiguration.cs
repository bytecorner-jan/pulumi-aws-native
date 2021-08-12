// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SSO
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html
    /// </summary>
    [AwsNativeResourceType("aws-native:SSO:InstanceAccessControlAttributeConfiguration")]
    public partial class InstanceAccessControlAttributeConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
        /// </summary>
        [Output("accessControlAttributes")]
        public Output<ImmutableArray<Outputs.InstanceAccessControlAttributeConfigurationAccessControlAttribute>> AccessControlAttributes { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instancearn
        /// </summary>
        [Output("instanceArn")]
        public Output<string> InstanceArn { get; private set; } = null!;


        /// <summary>
        /// Create a InstanceAccessControlAttributeConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public InstanceAccessControlAttributeConfiguration(string name, InstanceAccessControlAttributeConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws-native:SSO:InstanceAccessControlAttributeConfiguration", name, args ?? new InstanceAccessControlAttributeConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private InstanceAccessControlAttributeConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:SSO:InstanceAccessControlAttributeConfiguration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing InstanceAccessControlAttributeConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static InstanceAccessControlAttributeConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new InstanceAccessControlAttributeConfiguration(name, id, options);
        }
    }

    public sealed class InstanceAccessControlAttributeConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("accessControlAttributes")]
        private InputList<Inputs.InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs>? _accessControlAttributes;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-accesscontrolattributes
        /// </summary>
        public InputList<Inputs.InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs> AccessControlAttributes
        {
            get => _accessControlAttributes ?? (_accessControlAttributes = new InputList<Inputs.InstanceAccessControlAttributeConfigurationAccessControlAttributeArgs>());
            set => _accessControlAttributes = value;
        }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sso-instanceaccesscontrolattributeconfiguration.html#cfn-sso-instanceaccesscontrolattributeconfiguration-instancearn
        /// </summary>
        [Input("instanceArn", required: true)]
        public Input<string> InstanceArn { get; set; } = null!;

        public InstanceAccessControlAttributeConfigurationArgs()
        {
        }
    }
}
