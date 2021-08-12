// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html
    /// </summary>
    [AwsNativeResourceType("aws-native:SageMaker:ImageVersion")]
    public partial class ImageVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-baseimage
        /// </summary>
        [Output("baseImage")]
        public Output<string> BaseImage { get; private set; } = null!;

        [Output("containerImage")]
        public Output<string> ContainerImage { get; private set; } = null!;

        [Output("imageArn")]
        public Output<string> ImageArn { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-imagename
        /// </summary>
        [Output("imageName")]
        public Output<string> ImageName { get; private set; } = null!;

        [Output("imageVersionArn")]
        public Output<string> ImageVersionArn { get; private set; } = null!;

        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a ImageVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ImageVersion(string name, ImageVersionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:SageMaker:ImageVersion", name, args ?? new ImageVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ImageVersion(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:SageMaker:ImageVersion", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ImageVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ImageVersion Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ImageVersion(name, id, options);
        }
    }

    public sealed class ImageVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-baseimage
        /// </summary>
        [Input("baseImage", required: true)]
        public Input<string> BaseImage { get; set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-imageversion.html#cfn-sagemaker-imageversion-imagename
        /// </summary>
        [Input("imageName", required: true)]
        public Input<string> ImageName { get; set; } = null!;

        public ImageVersionArgs()
        {
        }
    }
}
