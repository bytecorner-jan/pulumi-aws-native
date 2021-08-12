// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class ImageRecipeComponentConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-componentarn
        /// </summary>
        public readonly string? ComponentArn;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html#cfn-imagebuilder-imagerecipe-componentconfiguration-parameters
        /// </summary>
        public readonly ImmutableArray<Outputs.ImageRecipeComponentParameter> Parameters;

        [OutputConstructor]
        private ImageRecipeComponentConfiguration(
            string? componentArn,

            ImmutableArray<Outputs.ImageRecipeComponentParameter> parameters)
        {
            ComponentArn = componentArn;
            Parameters = parameters;
        }
    }
}
