// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html
type ReplicationSubnetGroup struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupdescription
	ReplicationSubnetGroupDescription pulumi.StringOutput `pulumi:"replicationSubnetGroupDescription"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupidentifier
	ReplicationSubnetGroupIdentifier pulumi.StringPtrOutput `pulumi:"replicationSubnetGroupIdentifier"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-subnetids
	SubnetIds pulumi.StringArrayOutput `pulumi:"subnetIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewReplicationSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewReplicationSubnetGroup(ctx *pulumi.Context,
	name string, args *ReplicationSubnetGroupArgs, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ReplicationSubnetGroupDescription == nil {
		return nil, errors.New("invalid value for required argument 'ReplicationSubnetGroupDescription'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource ReplicationSubnetGroup
	err := ctx.RegisterResource("aws-native:dms:ReplicationSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReplicationSubnetGroup gets an existing ReplicationSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReplicationSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReplicationSubnetGroupState, opts ...pulumi.ResourceOption) (*ReplicationSubnetGroup, error) {
	var resource ReplicationSubnetGroup
	err := ctx.ReadResource("aws-native:dms:ReplicationSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReplicationSubnetGroup resources.
type replicationSubnetGroupState struct {
}

type ReplicationSubnetGroupState struct {
}

func (ReplicationSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupState)(nil)).Elem()
}

type replicationSubnetGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupdescription
	ReplicationSubnetGroupDescription string `pulumi:"replicationSubnetGroupDescription"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupidentifier
	ReplicationSubnetGroupIdentifier *string `pulumi:"replicationSubnetGroupIdentifier"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-subnetids
	SubnetIds []string `pulumi:"subnetIds"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a ReplicationSubnetGroup resource.
type ReplicationSubnetGroupArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupdescription
	ReplicationSubnetGroupDescription pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupidentifier
	ReplicationSubnetGroupIdentifier pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-subnetids
	SubnetIds pulumi.StringArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-tags
	Tags aws.TagArrayInput
}

func (ReplicationSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*replicationSubnetGroupArgs)(nil)).Elem()
}

type ReplicationSubnetGroupInput interface {
	pulumi.Input

	ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput
	ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput
}

func (*ReplicationSubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationSubnetGroup)(nil))
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput {
	return i.ToReplicationSubnetGroupOutputWithContext(context.Background())
}

func (i *ReplicationSubnetGroup) ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationSubnetGroupOutput)
}

type ReplicationSubnetGroupOutput struct{ *pulumi.OutputState }

func (ReplicationSubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationSubnetGroup)(nil))
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupOutput() ReplicationSubnetGroupOutput {
	return o
}

func (o ReplicationSubnetGroupOutput) ToReplicationSubnetGroupOutputWithContext(ctx context.Context) ReplicationSubnetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ReplicationSubnetGroupOutput{})
}
