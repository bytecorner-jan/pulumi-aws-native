// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RDS::DBSubnetGroup
//
// Deprecated: DBSubnetGroup is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DBSubnetGroup struct {
	pulumi.CustomResourceState

	DBSubnetGroupDescription pulumi.StringOutput         `pulumi:"dBSubnetGroupDescription"`
	DBSubnetGroupName        pulumi.StringPtrOutput      `pulumi:"dBSubnetGroupName"`
	SubnetIds                pulumi.StringArrayOutput    `pulumi:"subnetIds"`
	Tags                     DBSubnetGroupTagArrayOutput `pulumi:"tags"`
}

// NewDBSubnetGroup registers a new resource with the given unique name, arguments, and options.
func NewDBSubnetGroup(ctx *pulumi.Context,
	name string, args *DBSubnetGroupArgs, opts ...pulumi.ResourceOption) (*DBSubnetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DBSubnetGroupDescription == nil {
		return nil, errors.New("invalid value for required argument 'DBSubnetGroupDescription'")
	}
	if args.SubnetIds == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIds'")
	}
	var resource DBSubnetGroup
	err := ctx.RegisterResource("aws-native:rds:DBSubnetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBSubnetGroup gets an existing DBSubnetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBSubnetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBSubnetGroupState, opts ...pulumi.ResourceOption) (*DBSubnetGroup, error) {
	var resource DBSubnetGroup
	err := ctx.ReadResource("aws-native:rds:DBSubnetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBSubnetGroup resources.
type dbsubnetGroupState struct {
}

type DBSubnetGroupState struct {
}

func (DBSubnetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbsubnetGroupState)(nil)).Elem()
}

type dbsubnetGroupArgs struct {
	DBSubnetGroupDescription string             `pulumi:"dBSubnetGroupDescription"`
	DBSubnetGroupName        *string            `pulumi:"dBSubnetGroupName"`
	SubnetIds                []string           `pulumi:"subnetIds"`
	Tags                     []DBSubnetGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a DBSubnetGroup resource.
type DBSubnetGroupArgs struct {
	DBSubnetGroupDescription pulumi.StringInput
	DBSubnetGroupName        pulumi.StringPtrInput
	SubnetIds                pulumi.StringArrayInput
	Tags                     DBSubnetGroupTagArrayInput
}

func (DBSubnetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbsubnetGroupArgs)(nil)).Elem()
}

type DBSubnetGroupInput interface {
	pulumi.Input

	ToDBSubnetGroupOutput() DBSubnetGroupOutput
	ToDBSubnetGroupOutputWithContext(ctx context.Context) DBSubnetGroupOutput
}

func (*DBSubnetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSubnetGroup)(nil))
}

func (i *DBSubnetGroup) ToDBSubnetGroupOutput() DBSubnetGroupOutput {
	return i.ToDBSubnetGroupOutputWithContext(context.Background())
}

func (i *DBSubnetGroup) ToDBSubnetGroupOutputWithContext(ctx context.Context) DBSubnetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBSubnetGroupOutput)
}

type DBSubnetGroupOutput struct{ *pulumi.OutputState }

func (DBSubnetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBSubnetGroup)(nil))
}

func (o DBSubnetGroupOutput) ToDBSubnetGroupOutput() DBSubnetGroupOutput {
	return o
}

func (o DBSubnetGroupOutput) ToDBSubnetGroupOutputWithContext(ctx context.Context) DBSubnetGroupOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DBSubnetGroupInput)(nil)).Elem(), &DBSubnetGroup{})
	pulumi.RegisterOutputType(DBSubnetGroupOutput{})
}
