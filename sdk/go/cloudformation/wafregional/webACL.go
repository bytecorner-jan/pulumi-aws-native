// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html
type WebACL struct {
	pulumi.CustomResourceState

	// The attributes associated with the resource
	Attributes WebACLAttributesOutput `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrOutput `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata pulumi.AnyOutput `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties WebACLPropertiesOutput `pulumi:"properties"`
}

// NewWebACL registers a new resource with the given unique name, arguments, and options.
func NewWebACL(ctx *pulumi.Context,
	name string, args *WebACLArgs, opts ...pulumi.ResourceOption) (*WebACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	var resource WebACL
	err := ctx.RegisterResource("cloudformation:WAFRegional:WebACL", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebACL gets an existing WebACL resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebACL(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebACLState, opts ...pulumi.ResourceOption) (*WebACL, error) {
	var resource WebACL
	err := ctx.ReadResource("cloudformation:WAFRegional:WebACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebACL resources.
type webACLState struct {
	// The attributes associated with the resource
	Attributes *WebACLAttributes `pulumi:"attributes"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties *WebACLProperties `pulumi:"properties"`
}

type WebACLState struct {
	// The attributes associated with the resource
	Attributes WebACLAttributesPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties WebACLPropertiesPtrInput
}

func (WebACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*webACLState)(nil)).Elem()
}

type webACLArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// An explicit logical ID for the resource
	LogicalId *string `pulumi:"logicalId"`
	// Arbitrary structured data associated with the resource
	Metadata interface{} `pulumi:"metadata"`
	// The input properties associated with the resource
	Properties WebACLProperties `pulumi:"properties"`
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy *string `pulumi:"updateReplacePolicy"`
}

// The set of arguments for constructing a WebACL resource.
type WebACLArgs struct {
	// With the deletionPolicy attribute you can preserve or (in some cases) backup a resource when its stack is deleted. You can specify a deletionPolicy attribute for each resource that you want to control. If a resource has no deletionPolicy attribute, AWS CloudFormation deletes the resource by default.
	DeletionPolicy pulumi.StringPtrInput
	// An explicit logical ID for the resource
	LogicalId pulumi.StringPtrInput
	// Arbitrary structured data associated with the resource
	Metadata pulumi.Input
	// The input properties associated with the resource
	Properties WebACLPropertiesInput
	// Use the updateReplacePolicy attribute to retain or (in some cases) backup the existing physical instance of a resource when it is replaced during a stack update operation.
	UpdateReplacePolicy pulumi.StringPtrInput
}

func (WebACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webACLArgs)(nil)).Elem()
}

type WebACLInput interface {
	pulumi.Input

	ToWebACLOutput() WebACLOutput
	ToWebACLOutputWithContext(ctx context.Context) WebACLOutput
}

func (*WebACL) ElementType() reflect.Type {
	return reflect.TypeOf((*WebACL)(nil))
}

func (i *WebACL) ToWebACLOutput() WebACLOutput {
	return i.ToWebACLOutputWithContext(context.Background())
}

func (i *WebACL) ToWebACLOutputWithContext(ctx context.Context) WebACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebACLOutput)
}

type WebACLOutput struct {
	*pulumi.OutputState
}

func (WebACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebACL)(nil))
}

func (o WebACLOutput) ToWebACLOutput() WebACLOutput {
	return o
}

func (o WebACLOutput) ToWebACLOutputWithContext(ctx context.Context) WebACLOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebACLOutput{})
}
