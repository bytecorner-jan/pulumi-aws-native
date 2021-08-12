// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ImageBuilder.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-additionalinstanceconfiguration.html
    /// </summary>
    public sealed class ImageRecipeAdditionalInstanceConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-additionalinstanceconfiguration.html#cfn-imagebuilder-imagerecipe-additionalinstanceconfiguration-systemsmanageragent
        /// </summary>
        [Input("systemsManagerAgent")]
        public Input<Inputs.ImageRecipeSystemsManagerAgentArgs>? SystemsManagerAgent { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-imagerecipe-additionalinstanceconfiguration.html#cfn-imagebuilder-imagerecipe-additionalinstanceconfiguration-userdataoverride
        /// </summary>
        [Input("userDataOverride", required: true)]
        public Input<string> UserDataOverride { get; set; } = null!;

        public ImageRecipeAdditionalInstanceConfigurationArgs()
        {
        }
    }
}
