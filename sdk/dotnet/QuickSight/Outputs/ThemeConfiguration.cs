// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    /// <summary>
    /// &lt;p&gt;The theme configuration. This configuration contains all of the display properties for
    ///             a theme.&lt;/p&gt;
    /// </summary>
    [OutputType]
    public sealed class ThemeConfiguration
    {
        public readonly Outputs.ThemeDataColorPalette? DataColorPalette;
        public readonly Outputs.ThemeSheetStyle? Sheet;
        public readonly Outputs.ThemeTypography? Typography;
        public readonly Outputs.ThemeUIColorPalette? UIColorPalette;

        [OutputConstructor]
        private ThemeConfiguration(
            Outputs.ThemeDataColorPalette? dataColorPalette,

            Outputs.ThemeSheetStyle? sheet,

            Outputs.ThemeTypography? typography,

            Outputs.ThemeUIColorPalette? uIColorPalette)
        {
            DataColorPalette = dataColorPalette;
            Sheet = sheet;
            Typography = typography;
            UIColorPalette = uIColorPalette;
        }
    }
}
