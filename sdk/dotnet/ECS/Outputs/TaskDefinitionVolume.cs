// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Outputs
{

    [OutputType]
    public sealed class TaskDefinitionVolume
    {
        public readonly Outputs.TaskDefinitionDockerVolumeConfiguration? DockerVolumeConfiguration;
        public readonly Outputs.TaskDefinitionEFSVolumeConfiguration? EFSVolumeConfiguration;
        public readonly Outputs.TaskDefinitionHostVolumeProperties? Host;
        public readonly string? Name;

        [OutputConstructor]
        private TaskDefinitionVolume(
            Outputs.TaskDefinitionDockerVolumeConfiguration? dockerVolumeConfiguration,

            Outputs.TaskDefinitionEFSVolumeConfiguration? eFSVolumeConfiguration,

            Outputs.TaskDefinitionHostVolumeProperties? host,

            string? name)
        {
            DockerVolumeConfiguration = dockerVolumeConfiguration;
            EFSVolumeConfiguration = eFSVolumeConfiguration;
            Host = host;
            Name = name;
        }
    }
}
