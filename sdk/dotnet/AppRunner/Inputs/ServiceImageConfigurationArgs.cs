// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppRunner.Inputs
{

    /// <summary>
    /// Image Configuration
    /// </summary>
    public sealed class ServiceImageConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("runtimeEnvironmentVariables")]
        private InputList<Inputs.ServiceKeyValuePairArgs>? _runtimeEnvironmentVariables;
        public InputList<Inputs.ServiceKeyValuePairArgs> RuntimeEnvironmentVariables
        {
            get => _runtimeEnvironmentVariables ?? (_runtimeEnvironmentVariables = new InputList<Inputs.ServiceKeyValuePairArgs>());
            set => _runtimeEnvironmentVariables = value;
        }

        /// <summary>
        /// Start Command
        /// </summary>
        [Input("startCommand")]
        public Input<string>? StartCommand { get; set; }

        public ServiceImageConfigurationArgs()
        {
        }
    }
}
