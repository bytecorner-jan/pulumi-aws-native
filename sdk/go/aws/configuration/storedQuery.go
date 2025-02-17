// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Config::StoredQuery
type StoredQuery struct {
	pulumi.CustomResourceState

	QueryArn         pulumi.StringOutput    `pulumi:"queryArn"`
	QueryDescription pulumi.StringPtrOutput `pulumi:"queryDescription"`
	QueryExpression  pulumi.StringOutput    `pulumi:"queryExpression"`
	QueryId          pulumi.StringOutput    `pulumi:"queryId"`
	QueryName        pulumi.StringOutput    `pulumi:"queryName"`
	// The tags for the stored query.
	Tags StoredQueryTagArrayOutput `pulumi:"tags"`
}

// NewStoredQuery registers a new resource with the given unique name, arguments, and options.
func NewStoredQuery(ctx *pulumi.Context,
	name string, args *StoredQueryArgs, opts ...pulumi.ResourceOption) (*StoredQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.QueryExpression == nil {
		return nil, errors.New("invalid value for required argument 'QueryExpression'")
	}
	if args.QueryName == nil {
		return nil, errors.New("invalid value for required argument 'QueryName'")
	}
	var resource StoredQuery
	err := ctx.RegisterResource("aws-native:configuration:StoredQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStoredQuery gets an existing StoredQuery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStoredQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoredQueryState, opts ...pulumi.ResourceOption) (*StoredQuery, error) {
	var resource StoredQuery
	err := ctx.ReadResource("aws-native:configuration:StoredQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StoredQuery resources.
type storedQueryState struct {
}

type StoredQueryState struct {
}

func (StoredQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*storedQueryState)(nil)).Elem()
}

type storedQueryArgs struct {
	QueryDescription *string `pulumi:"queryDescription"`
	QueryExpression  string  `pulumi:"queryExpression"`
	QueryName        string  `pulumi:"queryName"`
	// The tags for the stored query.
	Tags []StoredQueryTag `pulumi:"tags"`
}

// The set of arguments for constructing a StoredQuery resource.
type StoredQueryArgs struct {
	QueryDescription pulumi.StringPtrInput
	QueryExpression  pulumi.StringInput
	QueryName        pulumi.StringInput
	// The tags for the stored query.
	Tags StoredQueryTagArrayInput
}

func (StoredQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storedQueryArgs)(nil)).Elem()
}

type StoredQueryInput interface {
	pulumi.Input

	ToStoredQueryOutput() StoredQueryOutput
	ToStoredQueryOutputWithContext(ctx context.Context) StoredQueryOutput
}

func (*StoredQuery) ElementType() reflect.Type {
	return reflect.TypeOf((*StoredQuery)(nil))
}

func (i *StoredQuery) ToStoredQueryOutput() StoredQueryOutput {
	return i.ToStoredQueryOutputWithContext(context.Background())
}

func (i *StoredQuery) ToStoredQueryOutputWithContext(ctx context.Context) StoredQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoredQueryOutput)
}

type StoredQueryOutput struct{ *pulumi.OutputState }

func (StoredQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoredQuery)(nil))
}

func (o StoredQueryOutput) ToStoredQueryOutput() StoredQueryOutput {
	return o
}

func (o StoredQueryOutput) ToStoredQueryOutputWithContext(ctx context.Context) StoredQueryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StoredQueryInput)(nil)).Elem(), &StoredQuery{})
	pulumi.RegisterOutputType(StoredQueryOutput{})
}
