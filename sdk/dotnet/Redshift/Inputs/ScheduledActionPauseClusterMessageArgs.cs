// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Redshift.Inputs
{

    /// <summary>
    /// Describes a pause cluster operation. For example, a scheduled action to run the `PauseCluster` API operation.
    /// </summary>
    public sealed class ScheduledActionPauseClusterMessageArgs : Pulumi.ResourceArgs
    {
        [Input("clusterIdentifier", required: true)]
        public Input<string> ClusterIdentifier { get; set; } = null!;

        public ScheduledActionPauseClusterMessageArgs()
        {
        }
    }
}
