// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html
    /// </summary>
    public sealed class ThemeSheetStyleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html#cfn-quicksight-theme-sheetstyle-tile
        /// </summary>
        [Input("tile")]
        public Input<Inputs.ThemeTileStyleArgs>? Tile { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-sheetstyle.html#cfn-quicksight-theme-sheetstyle-tilelayout
        /// </summary>
        [Input("tileLayout")]
        public Input<Inputs.ThemeTileLayoutStyleArgs>? TileLayout { get; set; }

        public ThemeSheetStyleArgs()
        {
        }
    }
}
