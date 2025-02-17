// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.NimbleStudio.Inputs
{

    /// <summary>
    /// &lt;p&gt;Initialization scripts for studio components.&lt;/p&gt;
    /// </summary>
    public sealed class StudioComponentInitializationScriptArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// &lt;p&gt;The version number of the protocol that is used by the launch profile. The only valid version is "2021-03-31".&lt;/p&gt;
        /// </summary>
        [Input("launchProfileProtocolVersion")]
        public Input<string>? LaunchProfileProtocolVersion { get; set; }

        [Input("platform")]
        public Input<Pulumi.AwsNative.NimbleStudio.StudioComponentLaunchProfilePlatform>? Platform { get; set; }

        [Input("runContext")]
        public Input<Pulumi.AwsNative.NimbleStudio.StudioComponentInitializationScriptRunContext>? RunContext { get; set; }

        /// <summary>
        /// &lt;p&gt;The initialization script.&lt;/p&gt;
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        public StudioComponentInitializationScriptArgs()
        {
        }
    }
}
