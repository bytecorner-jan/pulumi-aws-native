// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::EC2::SubnetNetworkAclAssociation
type SubnetNetworkAclAssociation struct {
	pulumi.CustomResourceState

	AssociationId pulumi.StringOutput `pulumi:"associationId"`
	// The ID of the network ACL
	NetworkAclId pulumi.StringOutput `pulumi:"networkAclId"`
	// The ID of the subnet
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
}

// NewSubnetNetworkAclAssociation registers a new resource with the given unique name, arguments, and options.
func NewSubnetNetworkAclAssociation(ctx *pulumi.Context,
	name string, args *SubnetNetworkAclAssociationArgs, opts ...pulumi.ResourceOption) (*SubnetNetworkAclAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkAclId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkAclId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource SubnetNetworkAclAssociation
	err := ctx.RegisterResource("aws-native:ec2:SubnetNetworkAclAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetNetworkAclAssociation gets an existing SubnetNetworkAclAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetNetworkAclAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetNetworkAclAssociationState, opts ...pulumi.ResourceOption) (*SubnetNetworkAclAssociation, error) {
	var resource SubnetNetworkAclAssociation
	err := ctx.ReadResource("aws-native:ec2:SubnetNetworkAclAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetNetworkAclAssociation resources.
type subnetNetworkAclAssociationState struct {
}

type SubnetNetworkAclAssociationState struct {
}

func (SubnetNetworkAclAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetNetworkAclAssociationState)(nil)).Elem()
}

type subnetNetworkAclAssociationArgs struct {
	// The ID of the network ACL
	NetworkAclId string `pulumi:"networkAclId"`
	// The ID of the subnet
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SubnetNetworkAclAssociation resource.
type SubnetNetworkAclAssociationArgs struct {
	// The ID of the network ACL
	NetworkAclId pulumi.StringInput
	// The ID of the subnet
	SubnetId pulumi.StringInput
}

func (SubnetNetworkAclAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetNetworkAclAssociationArgs)(nil)).Elem()
}

type SubnetNetworkAclAssociationInput interface {
	pulumi.Input

	ToSubnetNetworkAclAssociationOutput() SubnetNetworkAclAssociationOutput
	ToSubnetNetworkAclAssociationOutputWithContext(ctx context.Context) SubnetNetworkAclAssociationOutput
}

func (*SubnetNetworkAclAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetNetworkAclAssociation)(nil))
}

func (i *SubnetNetworkAclAssociation) ToSubnetNetworkAclAssociationOutput() SubnetNetworkAclAssociationOutput {
	return i.ToSubnetNetworkAclAssociationOutputWithContext(context.Background())
}

func (i *SubnetNetworkAclAssociation) ToSubnetNetworkAclAssociationOutputWithContext(ctx context.Context) SubnetNetworkAclAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetNetworkAclAssociationOutput)
}

type SubnetNetworkAclAssociationOutput struct{ *pulumi.OutputState }

func (SubnetNetworkAclAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetNetworkAclAssociation)(nil))
}

func (o SubnetNetworkAclAssociationOutput) ToSubnetNetworkAclAssociationOutput() SubnetNetworkAclAssociationOutput {
	return o
}

func (o SubnetNetworkAclAssociationOutput) ToSubnetNetworkAclAssociationOutputWithContext(ctx context.Context) SubnetNetworkAclAssociationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetNetworkAclAssociationOutput{})
}
