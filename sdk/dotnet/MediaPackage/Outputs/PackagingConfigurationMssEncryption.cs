// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.MediaPackage.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-mssencryption.html
    /// </summary>
    [OutputType]
    public sealed class PackagingConfigurationMssEncryption
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-mssencryption.html#cfn-mediapackage-packagingconfiguration-mssencryption-spekekeyprovider
        /// </summary>
        public readonly Outputs.PackagingConfigurationSpekeKeyProvider SpekeKeyProvider;

        [OutputConstructor]
        private PackagingConfigurationMssEncryption(Outputs.PackagingConfigurationSpekeKeyProvider spekeKeyProvider)
        {
            SpekeKeyProvider = spekeKeyProvider;
        }
    }
}
