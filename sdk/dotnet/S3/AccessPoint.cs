// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html
    /// </summary>
    [AwsNativeResourceType("aws-native:S3:AccessPoint")]
    public partial class AccessPoint : Pulumi.CustomResource
    {
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-bucket
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("networkOrigin")]
        public Output<string> NetworkOrigin { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-policy
        /// </summary>
        [Output("policy")]
        public Output<Union<System.Text.Json.JsonElement, string>?> Policy { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-policystatus
        /// </summary>
        [Output("policyStatus")]
        public Output<Union<System.Text.Json.JsonElement, string>?> PolicyStatus { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-publicaccessblockconfiguration
        /// </summary>
        [Output("publicAccessBlockConfiguration")]
        public Output<Outputs.AccessPointPublicAccessBlockConfiguration?> PublicAccessBlockConfiguration { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-vpcconfiguration
        /// </summary>
        [Output("vpcConfiguration")]
        public Output<Outputs.AccessPointVpcConfiguration?> VpcConfiguration { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPoint resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPoint(string name, AccessPointArgs args, CustomResourceOptions? options = null)
            : base("aws-native:S3:AccessPoint", name, args ?? new AccessPointArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPoint(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:S3:AccessPoint", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPoint resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPoint Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPoint(name, id, options);
        }
    }

    public sealed class AccessPointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-bucket
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-policy
        /// </summary>
        [Input("policy")]
        public InputUnion<System.Text.Json.JsonElement, string>? Policy { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-policystatus
        /// </summary>
        [Input("policyStatus")]
        public InputUnion<System.Text.Json.JsonElement, string>? PolicyStatus { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-publicaccessblockconfiguration
        /// </summary>
        [Input("publicAccessBlockConfiguration")]
        public Input<Inputs.AccessPointPublicAccessBlockConfigurationArgs>? PublicAccessBlockConfiguration { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accesspoint.html#cfn-s3-accesspoint-vpcconfiguration
        /// </summary>
        [Input("vpcConfiguration")]
        public Input<Inputs.AccessPointVpcConfigurationArgs>? VpcConfiguration { get; set; }

        public AccessPointArgs()
        {
        }
    }
}
