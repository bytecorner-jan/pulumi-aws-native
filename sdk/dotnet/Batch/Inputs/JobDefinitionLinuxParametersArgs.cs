// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionLinuxParametersArgs : Pulumi.ResourceArgs
    {
        [Input("devices")]
        private InputList<Inputs.JobDefinitionDeviceArgs>? _devices;
        public InputList<Inputs.JobDefinitionDeviceArgs> Devices
        {
            get => _devices ?? (_devices = new InputList<Inputs.JobDefinitionDeviceArgs>());
            set => _devices = value;
        }

        [Input("initProcessEnabled")]
        public Input<bool>? InitProcessEnabled { get; set; }

        [Input("maxSwap")]
        public Input<int>? MaxSwap { get; set; }

        [Input("sharedMemorySize")]
        public Input<int>? SharedMemorySize { get; set; }

        [Input("swappiness")]
        public Input<int>? Swappiness { get; set; }

        [Input("tmpfs")]
        private InputList<Inputs.JobDefinitionTmpfsArgs>? _tmpfs;
        public InputList<Inputs.JobDefinitionTmpfsArgs> Tmpfs
        {
            get => _tmpfs ?? (_tmpfs = new InputList<Inputs.JobDefinitionTmpfsArgs>());
            set => _tmpfs = value;
        }

        public JobDefinitionLinuxParametersArgs()
        {
        }
    }
}
