// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html
type VPCEndpoint struct {
	pulumi.CustomResourceState

	CreationTimestamp   pulumi.StringOutput      `pulumi:"creationTimestamp"`
	DnsEntries          pulumi.StringArrayOutput `pulumi:"dnsEntries"`
	NetworkInterfaceIds pulumi.StringArrayOutput `pulumi:"networkInterfaceIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-policydocument
	PolicyDocument pulumi.AnyOutput `pulumi:"policyDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-privatednsenabled
	PrivateDnsEnabled pulumi.BoolPtrOutput `pulumi:"privateDnsEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-routetableids
	RouteTableIds pulumi.StringArrayOutput `pulumi:"routeTableIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-securitygroupids
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-servicename
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-subnetids
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcendpointtype
	VpcEndpointType pulumi.StringPtrOutput `pulumi:"vpcEndpointType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcid
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewVPCEndpoint registers a new resource with the given unique name, arguments, and options.
func NewVPCEndpoint(ctx *pulumi.Context,
	name string, args *VPCEndpointArgs, opts ...pulumi.ResourceOption) (*VPCEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	var resource VPCEndpoint
	err := ctx.RegisterResource("aws-native:ec2:VPCEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVPCEndpoint gets an existing VPCEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVPCEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VPCEndpointState, opts ...pulumi.ResourceOption) (*VPCEndpoint, error) {
	var resource VPCEndpoint
	err := ctx.ReadResource("aws-native:ec2:VPCEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VPCEndpoint resources.
type vpcendpointState struct {
}

type VPCEndpointState struct {
}

func (VPCEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcendpointState)(nil)).Elem()
}

type vpcendpointArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-policydocument
	PolicyDocument interface{} `pulumi:"policyDocument"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-privatednsenabled
	PrivateDnsEnabled *bool `pulumi:"privateDnsEnabled"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-routetableids
	RouteTableIds []string `pulumi:"routeTableIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-securitygroupids
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-servicename
	ServiceName string `pulumi:"serviceName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-subnetids
	SubnetIds []string `pulumi:"subnetIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcendpointtype
	VpcEndpointType *string `pulumi:"vpcEndpointType"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcid
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a VPCEndpoint resource.
type VPCEndpointArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-policydocument
	PolicyDocument pulumi.Input
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-privatednsenabled
	PrivateDnsEnabled pulumi.BoolPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-routetableids
	RouteTableIds pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-securitygroupids
	SecurityGroupIds pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-servicename
	ServiceName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-subnetids
	SubnetIds pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcendpointtype
	VpcEndpointType pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcid
	VpcId pulumi.StringInput
}

func (VPCEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcendpointArgs)(nil)).Elem()
}

type VPCEndpointInput interface {
	pulumi.Input

	ToVPCEndpointOutput() VPCEndpointOutput
	ToVPCEndpointOutputWithContext(ctx context.Context) VPCEndpointOutput
}

func (*VPCEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((*VPCEndpoint)(nil))
}

func (i *VPCEndpoint) ToVPCEndpointOutput() VPCEndpointOutput {
	return i.ToVPCEndpointOutputWithContext(context.Background())
}

func (i *VPCEndpoint) ToVPCEndpointOutputWithContext(ctx context.Context) VPCEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VPCEndpointOutput)
}

type VPCEndpointOutput struct{ *pulumi.OutputState }

func (VPCEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VPCEndpoint)(nil))
}

func (o VPCEndpointOutput) ToVPCEndpointOutput() VPCEndpointOutput {
	return o
}

func (o VPCEndpointOutput) ToVPCEndpointOutputWithContext(ctx context.Context) VPCEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VPCEndpointOutput{})
}
