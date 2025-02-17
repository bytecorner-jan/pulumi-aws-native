// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    /// <summary>
    /// &lt;p&gt;The typeface for the theme.&lt;/p&gt;
    /// </summary>
    public sealed class ThemeTypographyArgs : Pulumi.ResourceArgs
    {
        [Input("fontFamilies")]
        private InputList<Inputs.ThemeFontArgs>? _fontFamilies;
        public InputList<Inputs.ThemeFontArgs> FontFamilies
        {
            get => _fontFamilies ?? (_fontFamilies = new InputList<Inputs.ThemeFontArgs>());
            set => _fontFamilies = value;
        }

        public ThemeTypographyArgs()
        {
        }
    }
}
