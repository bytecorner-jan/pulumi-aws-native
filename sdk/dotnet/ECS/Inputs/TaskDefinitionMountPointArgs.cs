// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ECS.Inputs
{

    public sealed class TaskDefinitionMountPointArgs : Pulumi.ResourceArgs
    {
        [Input("containerPath")]
        public Input<string>? ContainerPath { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("sourceVolume")]
        public Input<string>? SourceVolume { get; set; }

        public TaskDefinitionMountPointArgs()
        {
        }
    }
}
