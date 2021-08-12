// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html
type WebACL struct {
	pulumi.CustomResourceState

	Arn      pulumi.StringOutput `pulumi:"arn"`
	Capacity pulumi.IntOutput    `pulumi:"capacity"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-customresponsebodies
	CustomResponseBodies WebACLCustomResponseBodyMapOutput `pulumi:"customResponseBodies"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-defaultaction
	DefaultAction WebACLDefaultActionOutput `pulumi:"defaultAction"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-description
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	Id             pulumi.StringOutput    `pulumi:"id"`
	LabelNamespace pulumi.StringOutput    `pulumi:"labelNamespace"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-rules
	Rules WebACLRuleArrayOutput `pulumi:"rules"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-scope
	Scope pulumi.StringOutput `pulumi:"scope"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-visibilityconfig
	VisibilityConfig WebACLVisibilityConfigOutput `pulumi:"visibilityConfig"`
}

// NewWebACL registers a new resource with the given unique name, arguments, and options.
func NewWebACL(ctx *pulumi.Context,
	name string, args *WebACLArgs, opts ...pulumi.ResourceOption) (*WebACL, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultAction == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAction'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.VisibilityConfig == nil {
		return nil, errors.New("invalid value for required argument 'VisibilityConfig'")
	}
	var resource WebACL
	err := ctx.RegisterResource("aws-native:WAFv2:WebACL", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:WAFv2:WebACL", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebACL resources.
type webACLState struct {
}

type WebACLState struct {
}

func (WebACLState) ElementType() reflect.Type {
	return reflect.TypeOf((*webACLState)(nil)).Elem()
}

type webACLArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-customresponsebodies
	CustomResponseBodies map[string]WebACLCustomResponseBody `pulumi:"customResponseBodies"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-defaultaction
	DefaultAction WebACLDefaultAction `pulumi:"defaultAction"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-name
	Name *string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-rules
	Rules []WebACLRule `pulumi:"rules"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-scope
	Scope string `pulumi:"scope"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-visibilityconfig
	VisibilityConfig WebACLVisibilityConfig `pulumi:"visibilityConfig"`
}

// The set of arguments for constructing a WebACL resource.
type WebACLArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-customresponsebodies
	CustomResponseBodies WebACLCustomResponseBodyMapInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-defaultaction
	DefaultAction WebACLDefaultActionInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-name
	Name pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-rules
	Rules WebACLRuleArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-scope
	Scope pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafv2-webacl.html#cfn-wafv2-webacl-visibilityconfig
	VisibilityConfig WebACLVisibilityConfigInput
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

type WebACLOutput struct{ *pulumi.OutputState }

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
