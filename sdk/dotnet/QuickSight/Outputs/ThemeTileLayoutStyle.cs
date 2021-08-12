// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilelayoutstyle.html
    /// </summary>
    [OutputType]
    public sealed class ThemeTileLayoutStyle
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilelayoutstyle.html#cfn-quicksight-theme-tilelayoutstyle-gutter
        /// </summary>
        public readonly Outputs.ThemeGutterStyle? Gutter;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-tilelayoutstyle.html#cfn-quicksight-theme-tilelayoutstyle-margin
        /// </summary>
        public readonly Outputs.ThemeMarginStyle? Margin;

        [OutputConstructor]
        private ThemeTileLayoutStyle(
            Outputs.ThemeGutterStyle? gutter,

            Outputs.ThemeMarginStyle? margin)
        {
            Gutter = gutter;
            Margin = margin;
        }
    }
}
