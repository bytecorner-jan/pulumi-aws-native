// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package athena

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html
type DataCatalog struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-name
	Name pulumi.StringOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-parameters
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDataCatalog registers a new resource with the given unique name, arguments, and options.
func NewDataCatalog(ctx *pulumi.Context,
	name string, args *DataCatalogArgs, opts ...pulumi.ResourceOption) (*DataCatalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource DataCatalog
	err := ctx.RegisterResource("aws-native:Athena:DataCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCatalog gets an existing DataCatalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCatalogState, opts ...pulumi.ResourceOption) (*DataCatalog, error) {
	var resource DataCatalog
	err := ctx.ReadResource("aws-native:Athena:DataCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCatalog resources.
type dataCatalogState struct {
}

type DataCatalogState struct {
}

func (DataCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogState)(nil)).Elem()
}

type dataCatalogArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-name
	Name string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-parameters
	Parameters map[string]string `pulumi:"parameters"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-type
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DataCatalog resource.
type DataCatalogArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-name
	Name pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-parameters
	Parameters pulumi.StringMapInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-athena-datacatalog.html#cfn-athena-datacatalog-type
	Type pulumi.StringInput
}

func (DataCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogArgs)(nil)).Elem()
}

type DataCatalogInput interface {
	pulumi.Input

	ToDataCatalogOutput() DataCatalogOutput
	ToDataCatalogOutputWithContext(ctx context.Context) DataCatalogOutput
}

func (*DataCatalog) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCatalog)(nil))
}

func (i *DataCatalog) ToDataCatalogOutput() DataCatalogOutput {
	return i.ToDataCatalogOutputWithContext(context.Background())
}

func (i *DataCatalog) ToDataCatalogOutputWithContext(ctx context.Context) DataCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCatalogOutput)
}

type DataCatalogOutput struct{ *pulumi.OutputState }

func (DataCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCatalog)(nil))
}

func (o DataCatalogOutput) ToDataCatalogOutput() DataCatalogOutput {
	return o
}

func (o DataCatalogOutput) ToDataCatalogOutputWithContext(ctx context.Context) DataCatalogOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataCatalogOutput{})
}
