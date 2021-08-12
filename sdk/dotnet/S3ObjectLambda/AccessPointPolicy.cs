// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.S3ObjectLambda
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html
    /// </summary>
    [AwsNativeResourceType("aws-native:S3ObjectLambda:AccessPointPolicy")]
    public partial class AccessPointPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-objectlambdaaccesspoint
        /// </summary>
        [Output("objectLambdaAccessPoint")]
        public Output<string> ObjectLambdaAccessPoint { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-policydocument
        /// </summary>
        [Output("policyDocument")]
        public Output<Union<System.Text.Json.JsonElement, string>> PolicyDocument { get; private set; } = null!;


        /// <summary>
        /// Create a AccessPointPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccessPointPolicy(string name, AccessPointPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:S3ObjectLambda:AccessPointPolicy", name, args ?? new AccessPointPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccessPointPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:S3ObjectLambda:AccessPointPolicy", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AccessPointPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccessPointPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccessPointPolicy(name, id, options);
        }
    }

    public sealed class AccessPointPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-objectlambdaaccesspoint
        /// </summary>
        [Input("objectLambdaAccessPoint", required: true)]
        public Input<string> ObjectLambdaAccessPoint { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3objectlambda-accesspointpolicy.html#cfn-s3objectlambda-accesspointpolicy-policydocument
        /// </summary>
        [Input("policyDocument", required: true)]
        public InputUnion<System.Text.Json.JsonElement, string> PolicyDocument { get; set; } = null!;

        public AccessPointPolicyArgs()
        {
        }
    }
}
