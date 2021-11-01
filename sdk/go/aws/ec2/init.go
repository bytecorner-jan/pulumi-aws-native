// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "aws-native:ec2:CapacityReservation":
		r = &CapacityReservation{}
	case "aws-native:ec2:CapacityReservationFleet":
		r = &CapacityReservationFleet{}
	case "aws-native:ec2:CarrierGateway":
		r = &CarrierGateway{}
	case "aws-native:ec2:ClientVpnAuthorizationRule":
		r = &ClientVpnAuthorizationRule{}
	case "aws-native:ec2:ClientVpnEndpoint":
		r = &ClientVpnEndpoint{}
	case "aws-native:ec2:ClientVpnRoute":
		r = &ClientVpnRoute{}
	case "aws-native:ec2:ClientVpnTargetNetworkAssociation":
		r = &ClientVpnTargetNetworkAssociation{}
	case "aws-native:ec2:CustomerGateway":
		r = &CustomerGateway{}
	case "aws-native:ec2:DHCPOptions":
		r = &DHCPOptions{}
	case "aws-native:ec2:EC2Fleet":
		r = &EC2Fleet{}
	case "aws-native:ec2:EIP":
		r = &EIP{}
	case "aws-native:ec2:EIPAssociation":
		r = &EIPAssociation{}
	case "aws-native:ec2:EgressOnlyInternetGateway":
		r = &EgressOnlyInternetGateway{}
	case "aws-native:ec2:EnclaveCertificateIamRoleAssociation":
		r = &EnclaveCertificateIamRoleAssociation{}
	case "aws-native:ec2:FlowLog":
		r = &FlowLog{}
	case "aws-native:ec2:GatewayRouteTableAssociation":
		r = &GatewayRouteTableAssociation{}
	case "aws-native:ec2:Host":
		r = &Host{}
	case "aws-native:ec2:Instance":
		r = &Instance{}
	case "aws-native:ec2:InternetGateway":
		r = &InternetGateway{}
	case "aws-native:ec2:LaunchTemplate":
		r = &LaunchTemplate{}
	case "aws-native:ec2:LocalGatewayRoute":
		r = &LocalGatewayRoute{}
	case "aws-native:ec2:LocalGatewayRouteTableVPCAssociation":
		r = &LocalGatewayRouteTableVPCAssociation{}
	case "aws-native:ec2:NatGateway":
		r = &NatGateway{}
	case "aws-native:ec2:NetworkAcl":
		r = &NetworkAcl{}
	case "aws-native:ec2:NetworkAclEntry":
		r = &NetworkAclEntry{}
	case "aws-native:ec2:NetworkInsightsAnalysis":
		r = &NetworkInsightsAnalysis{}
	case "aws-native:ec2:NetworkInsightsPath":
		r = &NetworkInsightsPath{}
	case "aws-native:ec2:NetworkInterface":
		r = &NetworkInterface{}
	case "aws-native:ec2:NetworkInterfaceAttachment":
		r = &NetworkInterfaceAttachment{}
	case "aws-native:ec2:NetworkInterfacePermission":
		r = &NetworkInterfacePermission{}
	case "aws-native:ec2:PlacementGroup":
		r = &PlacementGroup{}
	case "aws-native:ec2:PrefixList":
		r = &PrefixList{}
	case "aws-native:ec2:Route":
		r = &Route{}
	case "aws-native:ec2:RouteTable":
		r = &RouteTable{}
	case "aws-native:ec2:SecurityGroup":
		r = &SecurityGroup{}
	case "aws-native:ec2:SecurityGroupEgress":
		r = &SecurityGroupEgress{}
	case "aws-native:ec2:SecurityGroupIngress":
		r = &SecurityGroupIngress{}
	case "aws-native:ec2:SpotFleet":
		r = &SpotFleet{}
	case "aws-native:ec2:Subnet":
		r = &Subnet{}
	case "aws-native:ec2:SubnetCidrBlock":
		r = &SubnetCidrBlock{}
	case "aws-native:ec2:SubnetNetworkAclAssociation":
		r = &SubnetNetworkAclAssociation{}
	case "aws-native:ec2:SubnetRouteTableAssociation":
		r = &SubnetRouteTableAssociation{}
	case "aws-native:ec2:TrafficMirrorFilter":
		r = &TrafficMirrorFilter{}
	case "aws-native:ec2:TrafficMirrorFilterRule":
		r = &TrafficMirrorFilterRule{}
	case "aws-native:ec2:TrafficMirrorSession":
		r = &TrafficMirrorSession{}
	case "aws-native:ec2:TrafficMirrorTarget":
		r = &TrafficMirrorTarget{}
	case "aws-native:ec2:TransitGateway":
		r = &TransitGateway{}
	case "aws-native:ec2:TransitGatewayAttachment":
		r = &TransitGatewayAttachment{}
	case "aws-native:ec2:TransitGatewayConnect":
		r = &TransitGatewayConnect{}
	case "aws-native:ec2:TransitGatewayMulticastDomain":
		r = &TransitGatewayMulticastDomain{}
	case "aws-native:ec2:TransitGatewayMulticastDomainAssociation":
		r = &TransitGatewayMulticastDomainAssociation{}
	case "aws-native:ec2:TransitGatewayMulticastGroupMember":
		r = &TransitGatewayMulticastGroupMember{}
	case "aws-native:ec2:TransitGatewayMulticastGroupSource":
		r = &TransitGatewayMulticastGroupSource{}
	case "aws-native:ec2:TransitGatewayPeeringAttachment":
		r = &TransitGatewayPeeringAttachment{}
	case "aws-native:ec2:TransitGatewayRoute":
		r = &TransitGatewayRoute{}
	case "aws-native:ec2:TransitGatewayRouteTable":
		r = &TransitGatewayRouteTable{}
	case "aws-native:ec2:TransitGatewayRouteTableAssociation":
		r = &TransitGatewayRouteTableAssociation{}
	case "aws-native:ec2:TransitGatewayRouteTablePropagation":
		r = &TransitGatewayRouteTablePropagation{}
	case "aws-native:ec2:TransitGatewayVpcAttachment":
		r = &TransitGatewayVpcAttachment{}
	case "aws-native:ec2:VPC":
		r = &VPC{}
	case "aws-native:ec2:VPCCidrBlock":
		r = &VPCCidrBlock{}
	case "aws-native:ec2:VPCDHCPOptionsAssociation":
		r = &VPCDHCPOptionsAssociation{}
	case "aws-native:ec2:VPCEndpoint":
		r = &VPCEndpoint{}
	case "aws-native:ec2:VPCEndpointConnectionNotification":
		r = &VPCEndpointConnectionNotification{}
	case "aws-native:ec2:VPCEndpointService":
		r = &VPCEndpointService{}
	case "aws-native:ec2:VPCEndpointServicePermissions":
		r = &VPCEndpointServicePermissions{}
	case "aws-native:ec2:VPCGatewayAttachment":
		r = &VPCGatewayAttachment{}
	case "aws-native:ec2:VPCPeeringConnection":
		r = &VPCPeeringConnection{}
	case "aws-native:ec2:VPNConnection":
		r = &VPNConnection{}
	case "aws-native:ec2:VPNConnectionRoute":
		r = &VPNConnectionRoute{}
	case "aws-native:ec2:VPNGateway":
		r = &VPNGateway{}
	case "aws-native:ec2:VPNGatewayRoutePropagation":
		r = &VPNGatewayRoutePropagation{}
	case "aws-native:ec2:Volume":
		r = &Volume{}
	case "aws-native:ec2:VolumeAttachment":
		r = &VolumeAttachment{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := aws.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"aws-native",
		"ec2",
		&module{version},
	)
}
