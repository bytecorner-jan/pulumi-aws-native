// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutEquipment.Inputs
{

    /// <summary>
    /// A tag is a key-value pair that can be added to a resource as metadata.
    /// </summary>
    public sealed class InferenceSchedulerTagArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key for the specified tag.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value for the specified tag.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public InferenceSchedulerTagArgs()
        {
        }
    }
}
