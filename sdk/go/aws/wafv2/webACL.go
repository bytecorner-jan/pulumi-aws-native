// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains the Rules that identify the requests that you want to allow, block, or count. In a WebACL, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a WebACL, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the WebACL with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a WebACL, a request needs to match only one of the specifications to be allowed, blocked, or counted.
type WebACL struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput                 `pulumi:"arn"`
	Capacity             pulumi.IntOutput                    `pulumi:"capacity"`
	CustomResponseBodies WebACLCustomResponseBodiesPtrOutput `pulumi:"customResponseBodies"`
	DefaultAction        WebACLDefaultActionOutput           `pulumi:"defaultAction"`
	Description          pulumi.StringPtrOutput              `pulumi:"description"`
	LabelNamespace       pulumi.StringOutput                 `pulumi:"labelNamespace"`
	Name                 pulumi.StringPtrOutput              `pulumi:"name"`
	// Collection of Rules.
	Rules            WebACLRuleArrayOutput        `pulumi:"rules"`
	Scope            WebACLScopeOutput            `pulumi:"scope"`
	Tags             WebACLTagArrayOutput         `pulumi:"tags"`
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
	err := ctx.RegisterResource("aws-native:wafv2:WebACL", name, args, &resource, opts...)
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
	err := ctx.ReadResource("aws-native:wafv2:WebACL", name, id, state, &resource, opts...)
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
	CustomResponseBodies *WebACLCustomResponseBodies `pulumi:"customResponseBodies"`
	DefaultAction        WebACLDefaultAction         `pulumi:"defaultAction"`
	Description          *string                     `pulumi:"description"`
	Name                 *string                     `pulumi:"name"`
	// Collection of Rules.
	Rules            []WebACLRule           `pulumi:"rules"`
	Scope            WebACLScope            `pulumi:"scope"`
	Tags             []WebACLTag            `pulumi:"tags"`
	VisibilityConfig WebACLVisibilityConfig `pulumi:"visibilityConfig"`
}

// The set of arguments for constructing a WebACL resource.
type WebACLArgs struct {
	CustomResponseBodies WebACLCustomResponseBodiesPtrInput
	DefaultAction        WebACLDefaultActionInput
	Description          pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	// Collection of Rules.
	Rules            WebACLRuleArrayInput
	Scope            WebACLScopeInput
	Tags             WebACLTagArrayInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*WebACLInput)(nil)).Elem(), &WebACL{})
	pulumi.RegisterOutputType(WebACLOutput{})
}
