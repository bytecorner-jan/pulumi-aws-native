// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html
    /// </summary>
    [OutputType]
    public sealed class ClusterExecuteCommandConfiguration
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-kmskeyid
        /// </summary>
        public readonly string? KmsKeyId;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-logconfiguration
        /// </summary>
        public readonly Outputs.ClusterExecuteCommandLogConfiguration? LogConfiguration;
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-cluster-executecommandconfiguration.html#cfn-ecs-cluster-executecommandconfiguration-logging
        /// </summary>
        public readonly string? Logging;

        [OutputConstructor]
        private ClusterExecuteCommandConfiguration(
            string? kmsKeyId,

            Outputs.ClusterExecuteCommandLogConfiguration? logConfiguration,

            string? logging)
        {
            KmsKeyId = kmsKeyId;
            LogConfiguration = logConfiguration;
            Logging = logging;
        }
    }
}
