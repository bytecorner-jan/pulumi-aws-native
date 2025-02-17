// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DynamoDB::Table
//
// Deprecated: Table is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Table struct {
	pulumi.CustomResourceState

	Arn                              pulumi.StringOutput                            `pulumi:"arn"`
	AttributeDefinitions             TableAttributeDefinitionArrayOutput            `pulumi:"attributeDefinitions"`
	BillingMode                      pulumi.StringPtrOutput                         `pulumi:"billingMode"`
	ContributorInsightsSpecification TableContributorInsightsSpecificationPtrOutput `pulumi:"contributorInsightsSpecification"`
	GlobalSecondaryIndexes           TableGlobalSecondaryIndexArrayOutput           `pulumi:"globalSecondaryIndexes"`
	KeySchema                        TableKeySchemaArrayOutput                      `pulumi:"keySchema"`
	KinesisStreamSpecification       TableKinesisStreamSpecificationPtrOutput       `pulumi:"kinesisStreamSpecification"`
	LocalSecondaryIndexes            TableLocalSecondaryIndexArrayOutput            `pulumi:"localSecondaryIndexes"`
	PointInTimeRecoverySpecification TablePointInTimeRecoverySpecificationPtrOutput `pulumi:"pointInTimeRecoverySpecification"`
	ProvisionedThroughput            TableProvisionedThroughputPtrOutput            `pulumi:"provisionedThroughput"`
	SSESpecification                 TableSSESpecificationPtrOutput                 `pulumi:"sSESpecification"`
	StreamArn                        pulumi.StringOutput                            `pulumi:"streamArn"`
	StreamSpecification              TableStreamSpecificationPtrOutput              `pulumi:"streamSpecification"`
	TableName                        pulumi.StringPtrOutput                         `pulumi:"tableName"`
	Tags                             TableTagArrayOutput                            `pulumi:"tags"`
	TimeToLiveSpecification          TableTimeToLiveSpecificationPtrOutput          `pulumi:"timeToLiveSpecification"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeySchema == nil {
		return nil, errors.New("invalid value for required argument 'KeySchema'")
	}
	var resource Table
	err := ctx.RegisterResource("aws-native:dynamodb:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("aws-native:dynamodb:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	AttributeDefinitions             []TableAttributeDefinition             `pulumi:"attributeDefinitions"`
	BillingMode                      *string                                `pulumi:"billingMode"`
	ContributorInsightsSpecification *TableContributorInsightsSpecification `pulumi:"contributorInsightsSpecification"`
	GlobalSecondaryIndexes           []TableGlobalSecondaryIndex            `pulumi:"globalSecondaryIndexes"`
	KeySchema                        []TableKeySchema                       `pulumi:"keySchema"`
	KinesisStreamSpecification       *TableKinesisStreamSpecification       `pulumi:"kinesisStreamSpecification"`
	LocalSecondaryIndexes            []TableLocalSecondaryIndex             `pulumi:"localSecondaryIndexes"`
	PointInTimeRecoverySpecification *TablePointInTimeRecoverySpecification `pulumi:"pointInTimeRecoverySpecification"`
	ProvisionedThroughput            *TableProvisionedThroughput            `pulumi:"provisionedThroughput"`
	SSESpecification                 *TableSSESpecification                 `pulumi:"sSESpecification"`
	StreamSpecification              *TableStreamSpecification              `pulumi:"streamSpecification"`
	TableName                        *string                                `pulumi:"tableName"`
	Tags                             []TableTag                             `pulumi:"tags"`
	TimeToLiveSpecification          *TableTimeToLiveSpecification          `pulumi:"timeToLiveSpecification"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	AttributeDefinitions             TableAttributeDefinitionArrayInput
	BillingMode                      pulumi.StringPtrInput
	ContributorInsightsSpecification TableContributorInsightsSpecificationPtrInput
	GlobalSecondaryIndexes           TableGlobalSecondaryIndexArrayInput
	KeySchema                        TableKeySchemaArrayInput
	KinesisStreamSpecification       TableKinesisStreamSpecificationPtrInput
	LocalSecondaryIndexes            TableLocalSecondaryIndexArrayInput
	PointInTimeRecoverySpecification TablePointInTimeRecoverySpecificationPtrInput
	ProvisionedThroughput            TableProvisionedThroughputPtrInput
	SSESpecification                 TableSSESpecificationPtrInput
	StreamSpecification              TableStreamSpecificationPtrInput
	TableName                        pulumi.StringPtrInput
	Tags                             TableTagArrayInput
	TimeToLiveSpecification          TableTimeToLiveSpecificationPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((*Table)(nil))
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Table)(nil))
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterOutputType(TableOutput{})
}
