// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// Current State of the Instance.
    /// </summary>
    public sealed class InstanceStateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Status code of the Instance.
        /// </summary>
        [Input("code")]
        public Input<int>? Code { get; set; }

        /// <summary>
        /// Status code of the Instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public InstanceStateArgs()
        {
        }
    }
}
