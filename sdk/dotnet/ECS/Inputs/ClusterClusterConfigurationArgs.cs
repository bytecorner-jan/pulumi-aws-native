// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Inputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clusterconfiguration.html
    /// </summary>
    public sealed class ClusterClusterConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-clusterconfiguration.html#cfn-ecs-cluster-clusterconfiguration-executecommandconfiguration
        /// </summary>
        [Input("executeCommandConfiguration")]
        public Input<Inputs.ClusterExecuteCommandConfigurationArgs>? ExecuteCommandConfiguration { get; set; }

        public ClusterClusterConfigurationArgs()
        {
        }
    }
}
