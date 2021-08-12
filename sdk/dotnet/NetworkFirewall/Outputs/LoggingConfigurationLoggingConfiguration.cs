// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-loggingconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class LoggingConfigurationLoggingConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-networkfirewall-loggingconfiguration-loggingconfiguration.html#cfn-networkfirewall-loggingconfiguration-loggingconfiguration-logdestinationconfigs
        /// </summary>
        public readonly ImmutableArray<Outputs.LoggingConfigurationLogDestinationConfig> LogDestinationConfigs;

        [OutputConstructor]
        private LoggingConfigurationLoggingConfiguration(ImmutableArray<Outputs.LoggingConfigurationLogDestinationConfig> logDestinationConfigs)
        {
            LogDestinationConfigs = logDestinationConfigs;
        }
    }
}
