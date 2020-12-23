// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbproxy.html
type DBProxy struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes DBProxyAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DBProxyPropertiesOutput `pulumi:"properties"`
}

// NewDBProxy registers a new resource with the given unique name, arguments, and options.
func NewDBProxy(ctx *pulumi.Context,
	name string, args *DBProxyArgs, opts ...pulumi.ResourceOption) (*DBProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource DBProxy
	err := ctx.RegisterResource("cloudformation:RDS:DBProxy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDBProxy gets an existing DBProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDBProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DBProxyState, opts ...pulumi.ResourceOption) (*DBProxy, error) {
	var resource DBProxy
	err := ctx.ReadResource("cloudformation:RDS:DBProxy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DBProxy resources.
type dbproxyState struct {
	// The attributes associated with the resource
	Attributes *DBProxyAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *DBProxyProperties `pulumi:"properties"`
}

type DBProxyState struct {
	// The attributes associated with the resource
	Attributes DBProxyAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DBProxyPropertiesPtrInput
}

func (DBProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dbproxyState)(nil)).Elem()
}

type dbproxyArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties DBProxyProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a DBProxy resource.
type DBProxyArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties DBProxyPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (DBProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dbproxyArgs)(nil)).Elem()
}

type DBProxyInput interface {
	pulumi.Input

	ToDBProxyOutput() DBProxyOutput
	ToDBProxyOutputWithContext(ctx context.Context) DBProxyOutput
}

func (*DBProxy) ElementType() reflect.Type {
	return reflect.TypeOf((*DBProxy)(nil))
}

func (i *DBProxy) ToDBProxyOutput() DBProxyOutput {
	return i.ToDBProxyOutputWithContext(context.Background())
}

func (i *DBProxy) ToDBProxyOutputWithContext(ctx context.Context) DBProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DBProxyOutput)
}

type DBProxyOutput struct {
	*pulumi.OutputState
}

func (DBProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DBProxy)(nil))
}

func (o DBProxyOutput) ToDBProxyOutput() DBProxyOutput {
	return o
}

func (o DBProxyOutput) ToDBProxyOutputWithContext(ctx context.Context) DBProxyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DBProxyOutput{})
}
