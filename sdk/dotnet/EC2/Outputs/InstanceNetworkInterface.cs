// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.EC2.Outputs
{

    [OutputType]
    public sealed class InstanceNetworkInterface
    {
        public readonly bool? AssociatePublicIpAddress;
        public readonly bool? DeleteOnTermination;
        public readonly string? Description;
        public readonly string DeviceIndex;
        public readonly ImmutableArray<string> GroupSet;
        public readonly int? Ipv6AddressCount;
        public readonly ImmutableArray<Outputs.InstanceIpv6Address> Ipv6Addresses;
        public readonly string? NetworkInterfaceId;
        public readonly string? PrivateIpAddress;
        public readonly ImmutableArray<Outputs.InstancePrivateIpAddressSpecification> PrivateIpAddresses;
        public readonly int? SecondaryPrivateIpAddressCount;
        public readonly string? SubnetId;

        [OutputConstructor]
        private InstanceNetworkInterface(
            bool? associatePublicIpAddress,

            bool? deleteOnTermination,

            string? description,

            string deviceIndex,

            ImmutableArray<string> groupSet,

            int? ipv6AddressCount,

            ImmutableArray<Outputs.InstanceIpv6Address> ipv6Addresses,

            string? networkInterfaceId,

            string? privateIpAddress,

            ImmutableArray<Outputs.InstancePrivateIpAddressSpecification> privateIpAddresses,

            int? secondaryPrivateIpAddressCount,

            string? subnetId)
        {
            AssociatePublicIpAddress = associatePublicIpAddress;
            DeleteOnTermination = deleteOnTermination;
            Description = description;
            DeviceIndex = deviceIndex;
            GroupSet = groupSet;
            Ipv6AddressCount = ipv6AddressCount;
            Ipv6Addresses = ipv6Addresses;
            NetworkInterfaceId = networkInterfaceId;
            PrivateIpAddress = privateIpAddress;
            PrivateIpAddresses = privateIpAddresses;
            SecondaryPrivateIpAddressCount = secondaryPrivateIpAddressCount;
            SubnetId = subnetId;
        }
    }
}
