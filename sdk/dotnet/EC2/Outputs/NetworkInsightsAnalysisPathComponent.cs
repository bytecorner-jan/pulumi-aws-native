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
    public sealed class NetworkInsightsAnalysisPathComponent
    {
        public readonly Outputs.NetworkInsightsAnalysisAnalysisAclRule? AclRule;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? Component;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? DestinationVpc;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? InboundHeader;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? OutboundHeader;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisRouteTableRoute? RouteTableRoute;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisSecurityGroupRule? SecurityGroupRule;
        public readonly int? SequenceNumber;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? SourceVpc;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? Subnet;
        public readonly Outputs.NetworkInsightsAnalysisAnalysisComponent? Vpc;

        [OutputConstructor]
        private NetworkInsightsAnalysisPathComponent(
            Outputs.NetworkInsightsAnalysisAnalysisAclRule? aclRule,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? component,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? destinationVpc,

            Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? inboundHeader,

            Outputs.NetworkInsightsAnalysisAnalysisPacketHeader? outboundHeader,

            Outputs.NetworkInsightsAnalysisAnalysisRouteTableRoute? routeTableRoute,

            Outputs.NetworkInsightsAnalysisAnalysisSecurityGroupRule? securityGroupRule,

            int? sequenceNumber,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? sourceVpc,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? subnet,

            Outputs.NetworkInsightsAnalysisAnalysisComponent? vpc)
        {
            AclRule = aclRule;
            Component = component;
            DestinationVpc = destinationVpc;
            InboundHeader = inboundHeader;
            OutboundHeader = outboundHeader;
            RouteTableRoute = routeTableRoute;
            SecurityGroupRule = securityGroupRule;
            SequenceNumber = sequenceNumber;
            SourceVpc = sourceVpc;
            Subnet = subnet;
            Vpc = vpc;
        }
    }
}
