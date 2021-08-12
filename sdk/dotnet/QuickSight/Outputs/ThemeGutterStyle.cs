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
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-gutterstyle.html
    /// </summary>
    [OutputType]
    public sealed class ThemeGutterStyle
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-theme-gutterstyle.html#cfn-quicksight-theme-gutterstyle-show
        /// </summary>
        public readonly bool? Show;

        [OutputConstructor]
        private ThemeGutterStyle(bool? show)
        {
            Show = show;
        }
    }
}
