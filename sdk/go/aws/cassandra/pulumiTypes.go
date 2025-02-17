// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cassandra

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KeyspaceTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// KeyspaceTagInput is an input type that accepts KeyspaceTagArgs and KeyspaceTagOutput values.
// You can construct a concrete instance of `KeyspaceTagInput` via:
//
//          KeyspaceTagArgs{...}
type KeyspaceTagInput interface {
	pulumi.Input

	ToKeyspaceTagOutput() KeyspaceTagOutput
	ToKeyspaceTagOutputWithContext(context.Context) KeyspaceTagOutput
}

type KeyspaceTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (KeyspaceTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyspaceTag)(nil)).Elem()
}

func (i KeyspaceTagArgs) ToKeyspaceTagOutput() KeyspaceTagOutput {
	return i.ToKeyspaceTagOutputWithContext(context.Background())
}

func (i KeyspaceTagArgs) ToKeyspaceTagOutputWithContext(ctx context.Context) KeyspaceTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyspaceTagOutput)
}

// KeyspaceTagArrayInput is an input type that accepts KeyspaceTagArray and KeyspaceTagArrayOutput values.
// You can construct a concrete instance of `KeyspaceTagArrayInput` via:
//
//          KeyspaceTagArray{ KeyspaceTagArgs{...} }
type KeyspaceTagArrayInput interface {
	pulumi.Input

	ToKeyspaceTagArrayOutput() KeyspaceTagArrayOutput
	ToKeyspaceTagArrayOutputWithContext(context.Context) KeyspaceTagArrayOutput
}

type KeyspaceTagArray []KeyspaceTagInput

func (KeyspaceTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyspaceTag)(nil)).Elem()
}

func (i KeyspaceTagArray) ToKeyspaceTagArrayOutput() KeyspaceTagArrayOutput {
	return i.ToKeyspaceTagArrayOutputWithContext(context.Background())
}

func (i KeyspaceTagArray) ToKeyspaceTagArrayOutputWithContext(ctx context.Context) KeyspaceTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyspaceTagArrayOutput)
}

type KeyspaceTagOutput struct{ *pulumi.OutputState }

func (KeyspaceTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyspaceTag)(nil)).Elem()
}

func (o KeyspaceTagOutput) ToKeyspaceTagOutput() KeyspaceTagOutput {
	return o
}

func (o KeyspaceTagOutput) ToKeyspaceTagOutputWithContext(ctx context.Context) KeyspaceTagOutput {
	return o
}

func (o KeyspaceTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v KeyspaceTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o KeyspaceTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyspaceTag) string { return v.Value }).(pulumi.StringOutput)
}

type KeyspaceTagArrayOutput struct{ *pulumi.OutputState }

func (KeyspaceTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyspaceTag)(nil)).Elem()
}

func (o KeyspaceTagArrayOutput) ToKeyspaceTagArrayOutput() KeyspaceTagArrayOutput {
	return o
}

func (o KeyspaceTagArrayOutput) ToKeyspaceTagArrayOutputWithContext(ctx context.Context) KeyspaceTagArrayOutput {
	return o
}

func (o KeyspaceTagArrayOutput) Index(i pulumi.IntInput) KeyspaceTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyspaceTag {
		return vs[0].([]KeyspaceTag)[vs[1].(int)]
	}).(KeyspaceTagOutput)
}

type TableBillingMode struct {
	Mode                  TableMode                   `pulumi:"mode"`
	ProvisionedThroughput *TableProvisionedThroughput `pulumi:"provisionedThroughput"`
}

// TableBillingModeInput is an input type that accepts TableBillingModeArgs and TableBillingModeOutput values.
// You can construct a concrete instance of `TableBillingModeInput` via:
//
//          TableBillingModeArgs{...}
type TableBillingModeInput interface {
	pulumi.Input

	ToTableBillingModeOutput() TableBillingModeOutput
	ToTableBillingModeOutputWithContext(context.Context) TableBillingModeOutput
}

type TableBillingModeArgs struct {
	Mode                  TableModeInput                     `pulumi:"mode"`
	ProvisionedThroughput TableProvisionedThroughputPtrInput `pulumi:"provisionedThroughput"`
}

func (TableBillingModeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableBillingMode)(nil)).Elem()
}

func (i TableBillingModeArgs) ToTableBillingModeOutput() TableBillingModeOutput {
	return i.ToTableBillingModeOutputWithContext(context.Background())
}

func (i TableBillingModeArgs) ToTableBillingModeOutputWithContext(ctx context.Context) TableBillingModeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableBillingModeOutput)
}

func (i TableBillingModeArgs) ToTableBillingModePtrOutput() TableBillingModePtrOutput {
	return i.ToTableBillingModePtrOutputWithContext(context.Background())
}

func (i TableBillingModeArgs) ToTableBillingModePtrOutputWithContext(ctx context.Context) TableBillingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableBillingModeOutput).ToTableBillingModePtrOutputWithContext(ctx)
}

// TableBillingModePtrInput is an input type that accepts TableBillingModeArgs, TableBillingModePtr and TableBillingModePtrOutput values.
// You can construct a concrete instance of `TableBillingModePtrInput` via:
//
//          TableBillingModeArgs{...}
//
//  or:
//
//          nil
type TableBillingModePtrInput interface {
	pulumi.Input

	ToTableBillingModePtrOutput() TableBillingModePtrOutput
	ToTableBillingModePtrOutputWithContext(context.Context) TableBillingModePtrOutput
}

type tableBillingModePtrType TableBillingModeArgs

func TableBillingModePtr(v *TableBillingModeArgs) TableBillingModePtrInput {
	return (*tableBillingModePtrType)(v)
}

func (*tableBillingModePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableBillingMode)(nil)).Elem()
}

func (i *tableBillingModePtrType) ToTableBillingModePtrOutput() TableBillingModePtrOutput {
	return i.ToTableBillingModePtrOutputWithContext(context.Background())
}

func (i *tableBillingModePtrType) ToTableBillingModePtrOutputWithContext(ctx context.Context) TableBillingModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableBillingModePtrOutput)
}

type TableBillingModeOutput struct{ *pulumi.OutputState }

func (TableBillingModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableBillingMode)(nil)).Elem()
}

func (o TableBillingModeOutput) ToTableBillingModeOutput() TableBillingModeOutput {
	return o
}

func (o TableBillingModeOutput) ToTableBillingModeOutputWithContext(ctx context.Context) TableBillingModeOutput {
	return o
}

func (o TableBillingModeOutput) ToTableBillingModePtrOutput() TableBillingModePtrOutput {
	return o.ToTableBillingModePtrOutputWithContext(context.Background())
}

func (o TableBillingModeOutput) ToTableBillingModePtrOutputWithContext(ctx context.Context) TableBillingModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableBillingMode) *TableBillingMode {
		return &v
	}).(TableBillingModePtrOutput)
}

func (o TableBillingModeOutput) Mode() TableModeOutput {
	return o.ApplyT(func(v TableBillingMode) TableMode { return v.Mode }).(TableModeOutput)
}

func (o TableBillingModeOutput) ProvisionedThroughput() TableProvisionedThroughputPtrOutput {
	return o.ApplyT(func(v TableBillingMode) *TableProvisionedThroughput { return v.ProvisionedThroughput }).(TableProvisionedThroughputPtrOutput)
}

type TableBillingModePtrOutput struct{ *pulumi.OutputState }

func (TableBillingModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableBillingMode)(nil)).Elem()
}

func (o TableBillingModePtrOutput) ToTableBillingModePtrOutput() TableBillingModePtrOutput {
	return o
}

func (o TableBillingModePtrOutput) ToTableBillingModePtrOutputWithContext(ctx context.Context) TableBillingModePtrOutput {
	return o
}

func (o TableBillingModePtrOutput) Elem() TableBillingModeOutput {
	return o.ApplyT(func(v *TableBillingMode) TableBillingMode {
		if v != nil {
			return *v
		}
		var ret TableBillingMode
		return ret
	}).(TableBillingModeOutput)
}

func (o TableBillingModePtrOutput) Mode() TableModePtrOutput {
	return o.ApplyT(func(v *TableBillingMode) *TableMode {
		if v == nil {
			return nil
		}
		return &v.Mode
	}).(TableModePtrOutput)
}

func (o TableBillingModePtrOutput) ProvisionedThroughput() TableProvisionedThroughputPtrOutput {
	return o.ApplyT(func(v *TableBillingMode) *TableProvisionedThroughput {
		if v == nil {
			return nil
		}
		return v.ProvisionedThroughput
	}).(TableProvisionedThroughputPtrOutput)
}

type TableClusteringKeyColumn struct {
	Column  TableColumn                      `pulumi:"column"`
	OrderBy *TableClusteringKeyColumnOrderBy `pulumi:"orderBy"`
}

// TableClusteringKeyColumnInput is an input type that accepts TableClusteringKeyColumnArgs and TableClusteringKeyColumnOutput values.
// You can construct a concrete instance of `TableClusteringKeyColumnInput` via:
//
//          TableClusteringKeyColumnArgs{...}
type TableClusteringKeyColumnInput interface {
	pulumi.Input

	ToTableClusteringKeyColumnOutput() TableClusteringKeyColumnOutput
	ToTableClusteringKeyColumnOutputWithContext(context.Context) TableClusteringKeyColumnOutput
}

type TableClusteringKeyColumnArgs struct {
	Column  TableColumnInput                        `pulumi:"column"`
	OrderBy TableClusteringKeyColumnOrderByPtrInput `pulumi:"orderBy"`
}

func (TableClusteringKeyColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableClusteringKeyColumn)(nil)).Elem()
}

func (i TableClusteringKeyColumnArgs) ToTableClusteringKeyColumnOutput() TableClusteringKeyColumnOutput {
	return i.ToTableClusteringKeyColumnOutputWithContext(context.Background())
}

func (i TableClusteringKeyColumnArgs) ToTableClusteringKeyColumnOutputWithContext(ctx context.Context) TableClusteringKeyColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableClusteringKeyColumnOutput)
}

// TableClusteringKeyColumnArrayInput is an input type that accepts TableClusteringKeyColumnArray and TableClusteringKeyColumnArrayOutput values.
// You can construct a concrete instance of `TableClusteringKeyColumnArrayInput` via:
//
//          TableClusteringKeyColumnArray{ TableClusteringKeyColumnArgs{...} }
type TableClusteringKeyColumnArrayInput interface {
	pulumi.Input

	ToTableClusteringKeyColumnArrayOutput() TableClusteringKeyColumnArrayOutput
	ToTableClusteringKeyColumnArrayOutputWithContext(context.Context) TableClusteringKeyColumnArrayOutput
}

type TableClusteringKeyColumnArray []TableClusteringKeyColumnInput

func (TableClusteringKeyColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableClusteringKeyColumn)(nil)).Elem()
}

func (i TableClusteringKeyColumnArray) ToTableClusteringKeyColumnArrayOutput() TableClusteringKeyColumnArrayOutput {
	return i.ToTableClusteringKeyColumnArrayOutputWithContext(context.Background())
}

func (i TableClusteringKeyColumnArray) ToTableClusteringKeyColumnArrayOutputWithContext(ctx context.Context) TableClusteringKeyColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableClusteringKeyColumnArrayOutput)
}

type TableClusteringKeyColumnOutput struct{ *pulumi.OutputState }

func (TableClusteringKeyColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableClusteringKeyColumn)(nil)).Elem()
}

func (o TableClusteringKeyColumnOutput) ToTableClusteringKeyColumnOutput() TableClusteringKeyColumnOutput {
	return o
}

func (o TableClusteringKeyColumnOutput) ToTableClusteringKeyColumnOutputWithContext(ctx context.Context) TableClusteringKeyColumnOutput {
	return o
}

func (o TableClusteringKeyColumnOutput) Column() TableColumnOutput {
	return o.ApplyT(func(v TableClusteringKeyColumn) TableColumn { return v.Column }).(TableColumnOutput)
}

func (o TableClusteringKeyColumnOutput) OrderBy() TableClusteringKeyColumnOrderByPtrOutput {
	return o.ApplyT(func(v TableClusteringKeyColumn) *TableClusteringKeyColumnOrderBy { return v.OrderBy }).(TableClusteringKeyColumnOrderByPtrOutput)
}

type TableClusteringKeyColumnArrayOutput struct{ *pulumi.OutputState }

func (TableClusteringKeyColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableClusteringKeyColumn)(nil)).Elem()
}

func (o TableClusteringKeyColumnArrayOutput) ToTableClusteringKeyColumnArrayOutput() TableClusteringKeyColumnArrayOutput {
	return o
}

func (o TableClusteringKeyColumnArrayOutput) ToTableClusteringKeyColumnArrayOutputWithContext(ctx context.Context) TableClusteringKeyColumnArrayOutput {
	return o
}

func (o TableClusteringKeyColumnArrayOutput) Index(i pulumi.IntInput) TableClusteringKeyColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TableClusteringKeyColumn {
		return vs[0].([]TableClusteringKeyColumn)[vs[1].(int)]
	}).(TableClusteringKeyColumnOutput)
}

type TableColumn struct {
	ColumnName string `pulumi:"columnName"`
	ColumnType string `pulumi:"columnType"`
}

// TableColumnInput is an input type that accepts TableColumnArgs and TableColumnOutput values.
// You can construct a concrete instance of `TableColumnInput` via:
//
//          TableColumnArgs{...}
type TableColumnInput interface {
	pulumi.Input

	ToTableColumnOutput() TableColumnOutput
	ToTableColumnOutputWithContext(context.Context) TableColumnOutput
}

type TableColumnArgs struct {
	ColumnName pulumi.StringInput `pulumi:"columnName"`
	ColumnType pulumi.StringInput `pulumi:"columnType"`
}

func (TableColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableColumn)(nil)).Elem()
}

func (i TableColumnArgs) ToTableColumnOutput() TableColumnOutput {
	return i.ToTableColumnOutputWithContext(context.Background())
}

func (i TableColumnArgs) ToTableColumnOutputWithContext(ctx context.Context) TableColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableColumnOutput)
}

// TableColumnArrayInput is an input type that accepts TableColumnArray and TableColumnArrayOutput values.
// You can construct a concrete instance of `TableColumnArrayInput` via:
//
//          TableColumnArray{ TableColumnArgs{...} }
type TableColumnArrayInput interface {
	pulumi.Input

	ToTableColumnArrayOutput() TableColumnArrayOutput
	ToTableColumnArrayOutputWithContext(context.Context) TableColumnArrayOutput
}

type TableColumnArray []TableColumnInput

func (TableColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableColumn)(nil)).Elem()
}

func (i TableColumnArray) ToTableColumnArrayOutput() TableColumnArrayOutput {
	return i.ToTableColumnArrayOutputWithContext(context.Background())
}

func (i TableColumnArray) ToTableColumnArrayOutputWithContext(ctx context.Context) TableColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableColumnArrayOutput)
}

type TableColumnOutput struct{ *pulumi.OutputState }

func (TableColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableColumn)(nil)).Elem()
}

func (o TableColumnOutput) ToTableColumnOutput() TableColumnOutput {
	return o
}

func (o TableColumnOutput) ToTableColumnOutputWithContext(ctx context.Context) TableColumnOutput {
	return o
}

func (o TableColumnOutput) ColumnName() pulumi.StringOutput {
	return o.ApplyT(func(v TableColumn) string { return v.ColumnName }).(pulumi.StringOutput)
}

func (o TableColumnOutput) ColumnType() pulumi.StringOutput {
	return o.ApplyT(func(v TableColumn) string { return v.ColumnType }).(pulumi.StringOutput)
}

type TableColumnArrayOutput struct{ *pulumi.OutputState }

func (TableColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableColumn)(nil)).Elem()
}

func (o TableColumnArrayOutput) ToTableColumnArrayOutput() TableColumnArrayOutput {
	return o
}

func (o TableColumnArrayOutput) ToTableColumnArrayOutputWithContext(ctx context.Context) TableColumnArrayOutput {
	return o
}

func (o TableColumnArrayOutput) Index(i pulumi.IntInput) TableColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TableColumn {
		return vs[0].([]TableColumn)[vs[1].(int)]
	}).(TableColumnOutput)
}

// Represents the settings used to enable server-side encryption
type TableEncryptionSpecification struct {
	EncryptionType   TableEncryptionType `pulumi:"encryptionType"`
	KmsKeyIdentifier *string             `pulumi:"kmsKeyIdentifier"`
}

// TableEncryptionSpecificationInput is an input type that accepts TableEncryptionSpecificationArgs and TableEncryptionSpecificationOutput values.
// You can construct a concrete instance of `TableEncryptionSpecificationInput` via:
//
//          TableEncryptionSpecificationArgs{...}
type TableEncryptionSpecificationInput interface {
	pulumi.Input

	ToTableEncryptionSpecificationOutput() TableEncryptionSpecificationOutput
	ToTableEncryptionSpecificationOutputWithContext(context.Context) TableEncryptionSpecificationOutput
}

// Represents the settings used to enable server-side encryption
type TableEncryptionSpecificationArgs struct {
	EncryptionType   TableEncryptionTypeInput `pulumi:"encryptionType"`
	KmsKeyIdentifier pulumi.StringPtrInput    `pulumi:"kmsKeyIdentifier"`
}

func (TableEncryptionSpecificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEncryptionSpecification)(nil)).Elem()
}

func (i TableEncryptionSpecificationArgs) ToTableEncryptionSpecificationOutput() TableEncryptionSpecificationOutput {
	return i.ToTableEncryptionSpecificationOutputWithContext(context.Background())
}

func (i TableEncryptionSpecificationArgs) ToTableEncryptionSpecificationOutputWithContext(ctx context.Context) TableEncryptionSpecificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEncryptionSpecificationOutput)
}

func (i TableEncryptionSpecificationArgs) ToTableEncryptionSpecificationPtrOutput() TableEncryptionSpecificationPtrOutput {
	return i.ToTableEncryptionSpecificationPtrOutputWithContext(context.Background())
}

func (i TableEncryptionSpecificationArgs) ToTableEncryptionSpecificationPtrOutputWithContext(ctx context.Context) TableEncryptionSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEncryptionSpecificationOutput).ToTableEncryptionSpecificationPtrOutputWithContext(ctx)
}

// TableEncryptionSpecificationPtrInput is an input type that accepts TableEncryptionSpecificationArgs, TableEncryptionSpecificationPtr and TableEncryptionSpecificationPtrOutput values.
// You can construct a concrete instance of `TableEncryptionSpecificationPtrInput` via:
//
//          TableEncryptionSpecificationArgs{...}
//
//  or:
//
//          nil
type TableEncryptionSpecificationPtrInput interface {
	pulumi.Input

	ToTableEncryptionSpecificationPtrOutput() TableEncryptionSpecificationPtrOutput
	ToTableEncryptionSpecificationPtrOutputWithContext(context.Context) TableEncryptionSpecificationPtrOutput
}

type tableEncryptionSpecificationPtrType TableEncryptionSpecificationArgs

func TableEncryptionSpecificationPtr(v *TableEncryptionSpecificationArgs) TableEncryptionSpecificationPtrInput {
	return (*tableEncryptionSpecificationPtrType)(v)
}

func (*tableEncryptionSpecificationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableEncryptionSpecification)(nil)).Elem()
}

func (i *tableEncryptionSpecificationPtrType) ToTableEncryptionSpecificationPtrOutput() TableEncryptionSpecificationPtrOutput {
	return i.ToTableEncryptionSpecificationPtrOutputWithContext(context.Background())
}

func (i *tableEncryptionSpecificationPtrType) ToTableEncryptionSpecificationPtrOutputWithContext(ctx context.Context) TableEncryptionSpecificationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableEncryptionSpecificationPtrOutput)
}

// Represents the settings used to enable server-side encryption
type TableEncryptionSpecificationOutput struct{ *pulumi.OutputState }

func (TableEncryptionSpecificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableEncryptionSpecification)(nil)).Elem()
}

func (o TableEncryptionSpecificationOutput) ToTableEncryptionSpecificationOutput() TableEncryptionSpecificationOutput {
	return o
}

func (o TableEncryptionSpecificationOutput) ToTableEncryptionSpecificationOutputWithContext(ctx context.Context) TableEncryptionSpecificationOutput {
	return o
}

func (o TableEncryptionSpecificationOutput) ToTableEncryptionSpecificationPtrOutput() TableEncryptionSpecificationPtrOutput {
	return o.ToTableEncryptionSpecificationPtrOutputWithContext(context.Background())
}

func (o TableEncryptionSpecificationOutput) ToTableEncryptionSpecificationPtrOutputWithContext(ctx context.Context) TableEncryptionSpecificationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableEncryptionSpecification) *TableEncryptionSpecification {
		return &v
	}).(TableEncryptionSpecificationPtrOutput)
}

func (o TableEncryptionSpecificationOutput) EncryptionType() TableEncryptionTypeOutput {
	return o.ApplyT(func(v TableEncryptionSpecification) TableEncryptionType { return v.EncryptionType }).(TableEncryptionTypeOutput)
}

func (o TableEncryptionSpecificationOutput) KmsKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TableEncryptionSpecification) *string { return v.KmsKeyIdentifier }).(pulumi.StringPtrOutput)
}

type TableEncryptionSpecificationPtrOutput struct{ *pulumi.OutputState }

func (TableEncryptionSpecificationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableEncryptionSpecification)(nil)).Elem()
}

func (o TableEncryptionSpecificationPtrOutput) ToTableEncryptionSpecificationPtrOutput() TableEncryptionSpecificationPtrOutput {
	return o
}

func (o TableEncryptionSpecificationPtrOutput) ToTableEncryptionSpecificationPtrOutputWithContext(ctx context.Context) TableEncryptionSpecificationPtrOutput {
	return o
}

func (o TableEncryptionSpecificationPtrOutput) Elem() TableEncryptionSpecificationOutput {
	return o.ApplyT(func(v *TableEncryptionSpecification) TableEncryptionSpecification {
		if v != nil {
			return *v
		}
		var ret TableEncryptionSpecification
		return ret
	}).(TableEncryptionSpecificationOutput)
}

func (o TableEncryptionSpecificationPtrOutput) EncryptionType() TableEncryptionTypePtrOutput {
	return o.ApplyT(func(v *TableEncryptionSpecification) *TableEncryptionType {
		if v == nil {
			return nil
		}
		return &v.EncryptionType
	}).(TableEncryptionTypePtrOutput)
}

func (o TableEncryptionSpecificationPtrOutput) KmsKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableEncryptionSpecification) *string {
		if v == nil {
			return nil
		}
		return v.KmsKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

// Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits
type TableProvisionedThroughput struct {
	ReadCapacityUnits  int `pulumi:"readCapacityUnits"`
	WriteCapacityUnits int `pulumi:"writeCapacityUnits"`
}

// TableProvisionedThroughputInput is an input type that accepts TableProvisionedThroughputArgs and TableProvisionedThroughputOutput values.
// You can construct a concrete instance of `TableProvisionedThroughputInput` via:
//
//          TableProvisionedThroughputArgs{...}
type TableProvisionedThroughputInput interface {
	pulumi.Input

	ToTableProvisionedThroughputOutput() TableProvisionedThroughputOutput
	ToTableProvisionedThroughputOutputWithContext(context.Context) TableProvisionedThroughputOutput
}

// Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits
type TableProvisionedThroughputArgs struct {
	ReadCapacityUnits  pulumi.IntInput `pulumi:"readCapacityUnits"`
	WriteCapacityUnits pulumi.IntInput `pulumi:"writeCapacityUnits"`
}

func (TableProvisionedThroughputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableProvisionedThroughput)(nil)).Elem()
}

func (i TableProvisionedThroughputArgs) ToTableProvisionedThroughputOutput() TableProvisionedThroughputOutput {
	return i.ToTableProvisionedThroughputOutputWithContext(context.Background())
}

func (i TableProvisionedThroughputArgs) ToTableProvisionedThroughputOutputWithContext(ctx context.Context) TableProvisionedThroughputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableProvisionedThroughputOutput)
}

func (i TableProvisionedThroughputArgs) ToTableProvisionedThroughputPtrOutput() TableProvisionedThroughputPtrOutput {
	return i.ToTableProvisionedThroughputPtrOutputWithContext(context.Background())
}

func (i TableProvisionedThroughputArgs) ToTableProvisionedThroughputPtrOutputWithContext(ctx context.Context) TableProvisionedThroughputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableProvisionedThroughputOutput).ToTableProvisionedThroughputPtrOutputWithContext(ctx)
}

// TableProvisionedThroughputPtrInput is an input type that accepts TableProvisionedThroughputArgs, TableProvisionedThroughputPtr and TableProvisionedThroughputPtrOutput values.
// You can construct a concrete instance of `TableProvisionedThroughputPtrInput` via:
//
//          TableProvisionedThroughputArgs{...}
//
//  or:
//
//          nil
type TableProvisionedThroughputPtrInput interface {
	pulumi.Input

	ToTableProvisionedThroughputPtrOutput() TableProvisionedThroughputPtrOutput
	ToTableProvisionedThroughputPtrOutputWithContext(context.Context) TableProvisionedThroughputPtrOutput
}

type tableProvisionedThroughputPtrType TableProvisionedThroughputArgs

func TableProvisionedThroughputPtr(v *TableProvisionedThroughputArgs) TableProvisionedThroughputPtrInput {
	return (*tableProvisionedThroughputPtrType)(v)
}

func (*tableProvisionedThroughputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableProvisionedThroughput)(nil)).Elem()
}

func (i *tableProvisionedThroughputPtrType) ToTableProvisionedThroughputPtrOutput() TableProvisionedThroughputPtrOutput {
	return i.ToTableProvisionedThroughputPtrOutputWithContext(context.Background())
}

func (i *tableProvisionedThroughputPtrType) ToTableProvisionedThroughputPtrOutputWithContext(ctx context.Context) TableProvisionedThroughputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableProvisionedThroughputPtrOutput)
}

// Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits
type TableProvisionedThroughputOutput struct{ *pulumi.OutputState }

func (TableProvisionedThroughputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableProvisionedThroughput)(nil)).Elem()
}

func (o TableProvisionedThroughputOutput) ToTableProvisionedThroughputOutput() TableProvisionedThroughputOutput {
	return o
}

func (o TableProvisionedThroughputOutput) ToTableProvisionedThroughputOutputWithContext(ctx context.Context) TableProvisionedThroughputOutput {
	return o
}

func (o TableProvisionedThroughputOutput) ToTableProvisionedThroughputPtrOutput() TableProvisionedThroughputPtrOutput {
	return o.ToTableProvisionedThroughputPtrOutputWithContext(context.Background())
}

func (o TableProvisionedThroughputOutput) ToTableProvisionedThroughputPtrOutputWithContext(ctx context.Context) TableProvisionedThroughputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableProvisionedThroughput) *TableProvisionedThroughput {
		return &v
	}).(TableProvisionedThroughputPtrOutput)
}

func (o TableProvisionedThroughputOutput) ReadCapacityUnits() pulumi.IntOutput {
	return o.ApplyT(func(v TableProvisionedThroughput) int { return v.ReadCapacityUnits }).(pulumi.IntOutput)
}

func (o TableProvisionedThroughputOutput) WriteCapacityUnits() pulumi.IntOutput {
	return o.ApplyT(func(v TableProvisionedThroughput) int { return v.WriteCapacityUnits }).(pulumi.IntOutput)
}

type TableProvisionedThroughputPtrOutput struct{ *pulumi.OutputState }

func (TableProvisionedThroughputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableProvisionedThroughput)(nil)).Elem()
}

func (o TableProvisionedThroughputPtrOutput) ToTableProvisionedThroughputPtrOutput() TableProvisionedThroughputPtrOutput {
	return o
}

func (o TableProvisionedThroughputPtrOutput) ToTableProvisionedThroughputPtrOutputWithContext(ctx context.Context) TableProvisionedThroughputPtrOutput {
	return o
}

func (o TableProvisionedThroughputPtrOutput) Elem() TableProvisionedThroughputOutput {
	return o.ApplyT(func(v *TableProvisionedThroughput) TableProvisionedThroughput {
		if v != nil {
			return *v
		}
		var ret TableProvisionedThroughput
		return ret
	}).(TableProvisionedThroughputOutput)
}

func (o TableProvisionedThroughputPtrOutput) ReadCapacityUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableProvisionedThroughput) *int {
		if v == nil {
			return nil
		}
		return &v.ReadCapacityUnits
	}).(pulumi.IntPtrOutput)
}

func (o TableProvisionedThroughputPtrOutput) WriteCapacityUnits() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableProvisionedThroughput) *int {
		if v == nil {
			return nil
		}
		return &v.WriteCapacityUnits
	}).(pulumi.IntPtrOutput)
}

// A key-value pair to apply to the resource
type TableTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

// TableTagInput is an input type that accepts TableTagArgs and TableTagOutput values.
// You can construct a concrete instance of `TableTagInput` via:
//
//          TableTagArgs{...}
type TableTagInput interface {
	pulumi.Input

	ToTableTagOutput() TableTagOutput
	ToTableTagOutputWithContext(context.Context) TableTagOutput
}

// A key-value pair to apply to the resource
type TableTagArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (TableTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableTag)(nil)).Elem()
}

func (i TableTagArgs) ToTableTagOutput() TableTagOutput {
	return i.ToTableTagOutputWithContext(context.Background())
}

func (i TableTagArgs) ToTableTagOutputWithContext(ctx context.Context) TableTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableTagOutput)
}

// TableTagArrayInput is an input type that accepts TableTagArray and TableTagArrayOutput values.
// You can construct a concrete instance of `TableTagArrayInput` via:
//
//          TableTagArray{ TableTagArgs{...} }
type TableTagArrayInput interface {
	pulumi.Input

	ToTableTagArrayOutput() TableTagArrayOutput
	ToTableTagArrayOutputWithContext(context.Context) TableTagArrayOutput
}

type TableTagArray []TableTagInput

func (TableTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableTag)(nil)).Elem()
}

func (i TableTagArray) ToTableTagArrayOutput() TableTagArrayOutput {
	return i.ToTableTagArrayOutputWithContext(context.Background())
}

func (i TableTagArray) ToTableTagArrayOutputWithContext(ctx context.Context) TableTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableTagArrayOutput)
}

// A key-value pair to apply to the resource
type TableTagOutput struct{ *pulumi.OutputState }

func (TableTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableTag)(nil)).Elem()
}

func (o TableTagOutput) ToTableTagOutput() TableTagOutput {
	return o
}

func (o TableTagOutput) ToTableTagOutputWithContext(ctx context.Context) TableTagOutput {
	return o
}

func (o TableTagOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v TableTag) string { return v.Key }).(pulumi.StringOutput)
}

func (o TableTagOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TableTag) string { return v.Value }).(pulumi.StringOutput)
}

type TableTagArrayOutput struct{ *pulumi.OutputState }

func (TableTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TableTag)(nil)).Elem()
}

func (o TableTagArrayOutput) ToTableTagArrayOutput() TableTagArrayOutput {
	return o
}

func (o TableTagArrayOutput) ToTableTagArrayOutputWithContext(ctx context.Context) TableTagArrayOutput {
	return o
}

func (o TableTagArrayOutput) Index(i pulumi.IntInput) TableTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TableTag {
		return vs[0].([]TableTag)[vs[1].(int)]
	}).(TableTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyspaceTagInput)(nil)).Elem(), KeyspaceTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyspaceTagArrayInput)(nil)).Elem(), KeyspaceTagArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableBillingModeInput)(nil)).Elem(), TableBillingModeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableBillingModePtrInput)(nil)).Elem(), TableBillingModeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableClusteringKeyColumnInput)(nil)).Elem(), TableClusteringKeyColumnArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableClusteringKeyColumnArrayInput)(nil)).Elem(), TableClusteringKeyColumnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableColumnInput)(nil)).Elem(), TableColumnArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableColumnArrayInput)(nil)).Elem(), TableColumnArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableEncryptionSpecificationInput)(nil)).Elem(), TableEncryptionSpecificationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableEncryptionSpecificationPtrInput)(nil)).Elem(), TableEncryptionSpecificationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableProvisionedThroughputInput)(nil)).Elem(), TableProvisionedThroughputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableProvisionedThroughputPtrInput)(nil)).Elem(), TableProvisionedThroughputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableTagInput)(nil)).Elem(), TableTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableTagArrayInput)(nil)).Elem(), TableTagArray{})
	pulumi.RegisterOutputType(KeyspaceTagOutput{})
	pulumi.RegisterOutputType(KeyspaceTagArrayOutput{})
	pulumi.RegisterOutputType(TableBillingModeOutput{})
	pulumi.RegisterOutputType(TableBillingModePtrOutput{})
	pulumi.RegisterOutputType(TableClusteringKeyColumnOutput{})
	pulumi.RegisterOutputType(TableClusteringKeyColumnArrayOutput{})
	pulumi.RegisterOutputType(TableColumnOutput{})
	pulumi.RegisterOutputType(TableColumnArrayOutput{})
	pulumi.RegisterOutputType(TableEncryptionSpecificationOutput{})
	pulumi.RegisterOutputType(TableEncryptionSpecificationPtrOutput{})
	pulumi.RegisterOutputType(TableProvisionedThroughputOutput{})
	pulumi.RegisterOutputType(TableProvisionedThroughputPtrOutput{})
	pulumi.RegisterOutputType(TableTagOutput{})
	pulumi.RegisterOutputType(TableTagArrayOutput{})
}
