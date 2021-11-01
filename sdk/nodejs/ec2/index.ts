// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./capacityReservation";
export * from "./capacityReservationFleet";
export * from "./carrierGateway";
export * from "./clientVpnAuthorizationRule";
export * from "./clientVpnEndpoint";
export * from "./clientVpnRoute";
export * from "./clientVpnTargetNetworkAssociation";
export * from "./customerGateway";
export * from "./dhcpoptions";
export * from "./ec2fleet";
export * from "./egressOnlyInternetGateway";
export * from "./eip";
export * from "./eipassociation";
export * from "./enclaveCertificateIamRoleAssociation";
export * from "./flowLog";
export * from "./gatewayRouteTableAssociation";
export * from "./host";
export * from "./instance";
export * from "./internetGateway";
export * from "./launchTemplate";
export * from "./localGatewayRoute";
export * from "./localGatewayRouteTableVPCAssociation";
export * from "./natGateway";
export * from "./networkAcl";
export * from "./networkAclEntry";
export * from "./networkInsightsAnalysis";
export * from "./networkInsightsPath";
export * from "./networkInterface";
export * from "./networkInterfaceAttachment";
export * from "./networkInterfacePermission";
export * from "./placementGroup";
export * from "./prefixList";
export * from "./route";
export * from "./routeTable";
export * from "./securityGroup";
export * from "./securityGroupEgress";
export * from "./securityGroupIngress";
export * from "./spotFleet";
export * from "./subnet";
export * from "./subnetCidrBlock";
export * from "./subnetNetworkAclAssociation";
export * from "./subnetRouteTableAssociation";
export * from "./trafficMirrorFilter";
export * from "./trafficMirrorFilterRule";
export * from "./trafficMirrorSession";
export * from "./trafficMirrorTarget";
export * from "./transitGateway";
export * from "./transitGatewayAttachment";
export * from "./transitGatewayConnect";
export * from "./transitGatewayMulticastDomain";
export * from "./transitGatewayMulticastDomainAssociation";
export * from "./transitGatewayMulticastGroupMember";
export * from "./transitGatewayMulticastGroupSource";
export * from "./transitGatewayPeeringAttachment";
export * from "./transitGatewayRoute";
export * from "./transitGatewayRouteTable";
export * from "./transitGatewayRouteTableAssociation";
export * from "./transitGatewayRouteTablePropagation";
export * from "./transitGatewayVpcAttachment";
export * from "./volume";
export * from "./volumeAttachment";
export * from "./vpc";
export * from "./vpccidrBlock";
export * from "./vpcdhcpoptionsAssociation";
export * from "./vpcendpoint";
export * from "./vpcendpointConnectionNotification";
export * from "./vpcendpointService";
export * from "./vpcendpointServicePermissions";
export * from "./vpcgatewayAttachment";
export * from "./vpcpeeringConnection";
export * from "./vpnconnection";
export * from "./vpnconnectionRoute";
export * from "./vpngateway";
export * from "./vpngatewayRoutePropagation";

// Export enums:
export * from "../types/enums/ec2";

// Import resources to register:
import { CapacityReservation } from "./capacityReservation";
import { CapacityReservationFleet } from "./capacityReservationFleet";
import { CarrierGateway } from "./carrierGateway";
import { ClientVpnAuthorizationRule } from "./clientVpnAuthorizationRule";
import { ClientVpnEndpoint } from "./clientVpnEndpoint";
import { ClientVpnRoute } from "./clientVpnRoute";
import { ClientVpnTargetNetworkAssociation } from "./clientVpnTargetNetworkAssociation";
import { CustomerGateway } from "./customerGateway";
import { DHCPOptions } from "./dhcpoptions";
import { EC2Fleet } from "./ec2fleet";
import { EIP } from "./eip";
import { EIPAssociation } from "./eipassociation";
import { EgressOnlyInternetGateway } from "./egressOnlyInternetGateway";
import { EnclaveCertificateIamRoleAssociation } from "./enclaveCertificateIamRoleAssociation";
import { FlowLog } from "./flowLog";
import { GatewayRouteTableAssociation } from "./gatewayRouteTableAssociation";
import { Host } from "./host";
import { Instance } from "./instance";
import { InternetGateway } from "./internetGateway";
import { LaunchTemplate } from "./launchTemplate";
import { LocalGatewayRoute } from "./localGatewayRoute";
import { LocalGatewayRouteTableVPCAssociation } from "./localGatewayRouteTableVPCAssociation";
import { NatGateway } from "./natGateway";
import { NetworkAcl } from "./networkAcl";
import { NetworkAclEntry } from "./networkAclEntry";
import { NetworkInsightsAnalysis } from "./networkInsightsAnalysis";
import { NetworkInsightsPath } from "./networkInsightsPath";
import { NetworkInterface } from "./networkInterface";
import { NetworkInterfaceAttachment } from "./networkInterfaceAttachment";
import { NetworkInterfacePermission } from "./networkInterfacePermission";
import { PlacementGroup } from "./placementGroup";
import { PrefixList } from "./prefixList";
import { Route } from "./route";
import { RouteTable } from "./routeTable";
import { SecurityGroup } from "./securityGroup";
import { SecurityGroupEgress } from "./securityGroupEgress";
import { SecurityGroupIngress } from "./securityGroupIngress";
import { SpotFleet } from "./spotFleet";
import { Subnet } from "./subnet";
import { SubnetCidrBlock } from "./subnetCidrBlock";
import { SubnetNetworkAclAssociation } from "./subnetNetworkAclAssociation";
import { SubnetRouteTableAssociation } from "./subnetRouteTableAssociation";
import { TrafficMirrorFilter } from "./trafficMirrorFilter";
import { TrafficMirrorFilterRule } from "./trafficMirrorFilterRule";
import { TrafficMirrorSession } from "./trafficMirrorSession";
import { TrafficMirrorTarget } from "./trafficMirrorTarget";
import { TransitGateway } from "./transitGateway";
import { TransitGatewayAttachment } from "./transitGatewayAttachment";
import { TransitGatewayConnect } from "./transitGatewayConnect";
import { TransitGatewayMulticastDomain } from "./transitGatewayMulticastDomain";
import { TransitGatewayMulticastDomainAssociation } from "./transitGatewayMulticastDomainAssociation";
import { TransitGatewayMulticastGroupMember } from "./transitGatewayMulticastGroupMember";
import { TransitGatewayMulticastGroupSource } from "./transitGatewayMulticastGroupSource";
import { TransitGatewayPeeringAttachment } from "./transitGatewayPeeringAttachment";
import { TransitGatewayRoute } from "./transitGatewayRoute";
import { TransitGatewayRouteTable } from "./transitGatewayRouteTable";
import { TransitGatewayRouteTableAssociation } from "./transitGatewayRouteTableAssociation";
import { TransitGatewayRouteTablePropagation } from "./transitGatewayRouteTablePropagation";
import { TransitGatewayVpcAttachment } from "./transitGatewayVpcAttachment";
import { VPC } from "./vpc";
import { VPCCidrBlock } from "./vpccidrBlock";
import { VPCDHCPOptionsAssociation } from "./vpcdhcpoptionsAssociation";
import { VPCEndpoint } from "./vpcendpoint";
import { VPCEndpointConnectionNotification } from "./vpcendpointConnectionNotification";
import { VPCEndpointService } from "./vpcendpointService";
import { VPCEndpointServicePermissions } from "./vpcendpointServicePermissions";
import { VPCGatewayAttachment } from "./vpcgatewayAttachment";
import { VPCPeeringConnection } from "./vpcpeeringConnection";
import { VPNConnection } from "./vpnconnection";
import { VPNConnectionRoute } from "./vpnconnectionRoute";
import { VPNGateway } from "./vpngateway";
import { VPNGatewayRoutePropagation } from "./vpngatewayRoutePropagation";
import { Volume } from "./volume";
import { VolumeAttachment } from "./volumeAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws-native:ec2:CapacityReservation":
                return new CapacityReservation(name, <any>undefined, { urn })
            case "aws-native:ec2:CapacityReservationFleet":
                return new CapacityReservationFleet(name, <any>undefined, { urn })
            case "aws-native:ec2:CarrierGateway":
                return new CarrierGateway(name, <any>undefined, { urn })
            case "aws-native:ec2:ClientVpnAuthorizationRule":
                return new ClientVpnAuthorizationRule(name, <any>undefined, { urn })
            case "aws-native:ec2:ClientVpnEndpoint":
                return new ClientVpnEndpoint(name, <any>undefined, { urn })
            case "aws-native:ec2:ClientVpnRoute":
                return new ClientVpnRoute(name, <any>undefined, { urn })
            case "aws-native:ec2:ClientVpnTargetNetworkAssociation":
                return new ClientVpnTargetNetworkAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:CustomerGateway":
                return new CustomerGateway(name, <any>undefined, { urn })
            case "aws-native:ec2:DHCPOptions":
                return new DHCPOptions(name, <any>undefined, { urn })
            case "aws-native:ec2:EC2Fleet":
                return new EC2Fleet(name, <any>undefined, { urn })
            case "aws-native:ec2:EIP":
                return new EIP(name, <any>undefined, { urn })
            case "aws-native:ec2:EIPAssociation":
                return new EIPAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:EgressOnlyInternetGateway":
                return new EgressOnlyInternetGateway(name, <any>undefined, { urn })
            case "aws-native:ec2:EnclaveCertificateIamRoleAssociation":
                return new EnclaveCertificateIamRoleAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:FlowLog":
                return new FlowLog(name, <any>undefined, { urn })
            case "aws-native:ec2:GatewayRouteTableAssociation":
                return new GatewayRouteTableAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:Host":
                return new Host(name, <any>undefined, { urn })
            case "aws-native:ec2:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "aws-native:ec2:InternetGateway":
                return new InternetGateway(name, <any>undefined, { urn })
            case "aws-native:ec2:LaunchTemplate":
                return new LaunchTemplate(name, <any>undefined, { urn })
            case "aws-native:ec2:LocalGatewayRoute":
                return new LocalGatewayRoute(name, <any>undefined, { urn })
            case "aws-native:ec2:LocalGatewayRouteTableVPCAssociation":
                return new LocalGatewayRouteTableVPCAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:NatGateway":
                return new NatGateway(name, <any>undefined, { urn })
            case "aws-native:ec2:NetworkAcl":
                return new NetworkAcl(name, <any>undefined, { urn })
            case "aws-native:ec2:NetworkAclEntry":
                return new NetworkAclEntry(name, <any>undefined, { urn })
            case "aws-native:ec2:NetworkInsightsAnalysis":
                return new NetworkInsightsAnalysis(name, <any>undefined, { urn })
            case "aws-native:ec2:NetworkInsightsPath":
                return new NetworkInsightsPath(name, <any>undefined, { urn })
            case "aws-native:ec2:NetworkInterface":
                return new NetworkInterface(name, <any>undefined, { urn })
            case "aws-native:ec2:NetworkInterfaceAttachment":
                return new NetworkInterfaceAttachment(name, <any>undefined, { urn })
            case "aws-native:ec2:NetworkInterfacePermission":
                return new NetworkInterfacePermission(name, <any>undefined, { urn })
            case "aws-native:ec2:PlacementGroup":
                return new PlacementGroup(name, <any>undefined, { urn })
            case "aws-native:ec2:PrefixList":
                return new PrefixList(name, <any>undefined, { urn })
            case "aws-native:ec2:Route":
                return new Route(name, <any>undefined, { urn })
            case "aws-native:ec2:RouteTable":
                return new RouteTable(name, <any>undefined, { urn })
            case "aws-native:ec2:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "aws-native:ec2:SecurityGroupEgress":
                return new SecurityGroupEgress(name, <any>undefined, { urn })
            case "aws-native:ec2:SecurityGroupIngress":
                return new SecurityGroupIngress(name, <any>undefined, { urn })
            case "aws-native:ec2:SpotFleet":
                return new SpotFleet(name, <any>undefined, { urn })
            case "aws-native:ec2:Subnet":
                return new Subnet(name, <any>undefined, { urn })
            case "aws-native:ec2:SubnetCidrBlock":
                return new SubnetCidrBlock(name, <any>undefined, { urn })
            case "aws-native:ec2:SubnetNetworkAclAssociation":
                return new SubnetNetworkAclAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:SubnetRouteTableAssociation":
                return new SubnetRouteTableAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:TrafficMirrorFilter":
                return new TrafficMirrorFilter(name, <any>undefined, { urn })
            case "aws-native:ec2:TrafficMirrorFilterRule":
                return new TrafficMirrorFilterRule(name, <any>undefined, { urn })
            case "aws-native:ec2:TrafficMirrorSession":
                return new TrafficMirrorSession(name, <any>undefined, { urn })
            case "aws-native:ec2:TrafficMirrorTarget":
                return new TrafficMirrorTarget(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGateway":
                return new TransitGateway(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayAttachment":
                return new TransitGatewayAttachment(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayConnect":
                return new TransitGatewayConnect(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayMulticastDomain":
                return new TransitGatewayMulticastDomain(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayMulticastDomainAssociation":
                return new TransitGatewayMulticastDomainAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayMulticastGroupMember":
                return new TransitGatewayMulticastGroupMember(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayMulticastGroupSource":
                return new TransitGatewayMulticastGroupSource(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayPeeringAttachment":
                return new TransitGatewayPeeringAttachment(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayRoute":
                return new TransitGatewayRoute(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayRouteTable":
                return new TransitGatewayRouteTable(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayRouteTableAssociation":
                return new TransitGatewayRouteTableAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayRouteTablePropagation":
                return new TransitGatewayRouteTablePropagation(name, <any>undefined, { urn })
            case "aws-native:ec2:TransitGatewayVpcAttachment":
                return new TransitGatewayVpcAttachment(name, <any>undefined, { urn })
            case "aws-native:ec2:VPC":
                return new VPC(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCCidrBlock":
                return new VPCCidrBlock(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCDHCPOptionsAssociation":
                return new VPCDHCPOptionsAssociation(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCEndpoint":
                return new VPCEndpoint(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCEndpointConnectionNotification":
                return new VPCEndpointConnectionNotification(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCEndpointService":
                return new VPCEndpointService(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCEndpointServicePermissions":
                return new VPCEndpointServicePermissions(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCGatewayAttachment":
                return new VPCGatewayAttachment(name, <any>undefined, { urn })
            case "aws-native:ec2:VPCPeeringConnection":
                return new VPCPeeringConnection(name, <any>undefined, { urn })
            case "aws-native:ec2:VPNConnection":
                return new VPNConnection(name, <any>undefined, { urn })
            case "aws-native:ec2:VPNConnectionRoute":
                return new VPNConnectionRoute(name, <any>undefined, { urn })
            case "aws-native:ec2:VPNGateway":
                return new VPNGateway(name, <any>undefined, { urn })
            case "aws-native:ec2:VPNGatewayRoutePropagation":
                return new VPNGatewayRoutePropagation(name, <any>undefined, { urn })
            case "aws-native:ec2:Volume":
                return new Volume(name, <any>undefined, { urn })
            case "aws-native:ec2:VolumeAttachment":
                return new VolumeAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws-native", "ec2", _module)
