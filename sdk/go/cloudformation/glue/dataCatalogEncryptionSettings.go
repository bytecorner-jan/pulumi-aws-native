// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-glue-datacatalogencryptionsettings.html
type DataCatalogEncryptionSettings struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes DataCatalogEncryptionSettingsAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DataCatalogEncryptionSettingsPropertiesOutput `pulumi:"properties"`
}

// NewDataCatalogEncryptionSettings registers a new resource with the given unique name, arguments, and options.
func NewDataCatalogEncryptionSettings(ctx *pulumi.Context,
	name string, args *DataCatalogEncryptionSettingsArgs, opts ...pulumi.ResourceOption) (*DataCatalogEncryptionSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource DataCatalogEncryptionSettings
	err := ctx.RegisterResource("cloudformation:Glue:DataCatalogEncryptionSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCatalogEncryptionSettings gets an existing DataCatalogEncryptionSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCatalogEncryptionSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCatalogEncryptionSettingsState, opts ...pulumi.ResourceOption) (*DataCatalogEncryptionSettings, error) {
	var resource DataCatalogEncryptionSettings
	err := ctx.ReadResource("cloudformation:Glue:DataCatalogEncryptionSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCatalogEncryptionSettings resources.
type dataCatalogEncryptionSettingsState struct {
	// The attributes associated with the resource
	Attributes *DataCatalogEncryptionSettingsAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *DataCatalogEncryptionSettingsProperties `pulumi:"properties"`
}

type DataCatalogEncryptionSettingsState struct {
	// The attributes associated with the resource
	Attributes DataCatalogEncryptionSettingsAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DataCatalogEncryptionSettingsPropertiesPtrInput
}

func (DataCatalogEncryptionSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogEncryptionSettingsState)(nil)).Elem()
}

type dataCatalogEncryptionSettingsArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DataCatalogEncryptionSettingsProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a DataCatalogEncryptionSettings resource.
type DataCatalogEncryptionSettingsArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DataCatalogEncryptionSettingsPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (DataCatalogEncryptionSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogEncryptionSettingsArgs)(nil)).Elem()
}

type DataCatalogEncryptionSettingsInput interface {
	pulumi.Input

	ToDataCatalogEncryptionSettingsOutput() DataCatalogEncryptionSettingsOutput
	ToDataCatalogEncryptionSettingsOutputWithContext(ctx context.Context) DataCatalogEncryptionSettingsOutput
}

func (*DataCatalogEncryptionSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCatalogEncryptionSettings)(nil))
}

func (i *DataCatalogEncryptionSettings) ToDataCatalogEncryptionSettingsOutput() DataCatalogEncryptionSettingsOutput {
	return i.ToDataCatalogEncryptionSettingsOutputWithContext(context.Background())
}

func (i *DataCatalogEncryptionSettings) ToDataCatalogEncryptionSettingsOutputWithContext(ctx context.Context) DataCatalogEncryptionSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCatalogEncryptionSettingsOutput)
}

type DataCatalogEncryptionSettingsOutput struct {
	*pulumi.OutputState
}

func (DataCatalogEncryptionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCatalogEncryptionSettings)(nil))
}

func (o DataCatalogEncryptionSettingsOutput) ToDataCatalogEncryptionSettingsOutput() DataCatalogEncryptionSettingsOutput {
	return o
}

func (o DataCatalogEncryptionSettingsOutput) ToDataCatalogEncryptionSettingsOutputWithContext(ctx context.Context) DataCatalogEncryptionSettingsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataCatalogEncryptionSettingsOutput{})
}
