// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GroundStation.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html
    /// </summary>
    public sealed class ConfigAntennaUplinkConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html#cfn-groundstation-config-antennauplinkconfig-spectrumconfig
        /// </summary>
        [Input("spectrumConfig")]
        public Input<Inputs.ConfigUplinkSpectrumConfigArgs>? SpectrumConfig { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html#cfn-groundstation-config-antennauplinkconfig-targeteirp
        /// </summary>
        [Input("targetEirp")]
        public Input<Inputs.ConfigEirpArgs>? TargetEirp { get; set; }

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-groundstation-config-antennauplinkconfig.html#cfn-groundstation-config-antennauplinkconfig-transmitdisabled
        /// </summary>
        [Input("transmitDisabled")]
        public Input<bool>? TransmitDisabled { get; set; }

        public ConfigAntennaUplinkConfigArgs()
        {
        }
    }
}
