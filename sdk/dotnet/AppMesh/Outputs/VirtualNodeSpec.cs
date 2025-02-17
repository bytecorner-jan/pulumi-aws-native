// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Outputs
{

    [OutputType]
    public sealed class VirtualNodeSpec
    {
        public readonly Outputs.VirtualNodeBackendDefaults? BackendDefaults;
        public readonly ImmutableArray<Outputs.VirtualNodeBackend> Backends;
        public readonly ImmutableArray<Outputs.VirtualNodeListener> Listeners;
        public readonly Outputs.VirtualNodeLogging? Logging;
        public readonly Outputs.VirtualNodeServiceDiscovery? ServiceDiscovery;

        [OutputConstructor]
        private VirtualNodeSpec(
            Outputs.VirtualNodeBackendDefaults? backendDefaults,

            ImmutableArray<Outputs.VirtualNodeBackend> backends,

            ImmutableArray<Outputs.VirtualNodeListener> listeners,

            Outputs.VirtualNodeLogging? logging,

            Outputs.VirtualNodeServiceDiscovery? serviceDiscovery)
        {
            BackendDefaults = backendDefaults;
            Backends = backends;
            Listeners = listeners;
            Logging = logging;
            ServiceDiscovery = serviceDiscovery;
        }
    }
}
