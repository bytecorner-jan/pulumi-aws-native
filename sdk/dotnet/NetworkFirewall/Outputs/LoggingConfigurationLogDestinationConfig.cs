// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NetworkFirewall.Outputs
{

    [OutputType]
    public sealed class LoggingConfigurationLogDestinationConfig
    {
        /// <summary>
        /// A key-value pair to configure the logDestinations.
        /// </summary>
        public readonly object LogDestination;
        public readonly Pulumi.AwsNative.NetworkFirewall.LoggingConfigurationLogDestinationConfigLogDestinationType LogDestinationType;
        public readonly Pulumi.AwsNative.NetworkFirewall.LoggingConfigurationLogDestinationConfigLogType LogType;

        [OutputConstructor]
        private LoggingConfigurationLogDestinationConfig(
            object logDestination,

            Pulumi.AwsNative.NetworkFirewall.LoggingConfigurationLogDestinationConfigLogDestinationType logDestinationType,

            Pulumi.AwsNative.NetworkFirewall.LoggingConfigurationLogDestinationConfigLogType logType)
        {
            LogDestination = logDestination;
            LogDestinationType = logDestinationType;
            LogType = logType;
        }
    }
}
