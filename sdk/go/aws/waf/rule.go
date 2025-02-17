// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::WAF::Rule
//
// Deprecated: Rule is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Rule struct {
	pulumi.CustomResourceState

	MetricName pulumi.StringOutput      `pulumi:"metricName"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Predicates RulePredicateArrayOutput `pulumi:"predicates"`
}

// NewRule registers a new resource with the given unique name, arguments, and options.
func NewRule(ctx *pulumi.Context,
	name string, args *RuleArgs, opts ...pulumi.ResourceOption) (*Rule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	var resource Rule
	err := ctx.RegisterResource("aws-native:waf:Rule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRule gets an existing Rule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleState, opts ...pulumi.ResourceOption) (*Rule, error) {
	var resource Rule
	err := ctx.ReadResource("aws-native:waf:Rule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rule resources.
type ruleState struct {
}

type RuleState struct {
}

func (RuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleState)(nil)).Elem()
}

type ruleArgs struct {
	MetricName string          `pulumi:"metricName"`
	Name       *string         `pulumi:"name"`
	Predicates []RulePredicate `pulumi:"predicates"`
}

// The set of arguments for constructing a Rule resource.
type RuleArgs struct {
	MetricName pulumi.StringInput
	Name       pulumi.StringPtrInput
	Predicates RulePredicateArrayInput
}

func (RuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleArgs)(nil)).Elem()
}

type RuleInput interface {
	pulumi.Input

	ToRuleOutput() RuleOutput
	ToRuleOutputWithContext(ctx context.Context) RuleOutput
}

func (*Rule) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (i *Rule) ToRuleOutput() RuleOutput {
	return i.ToRuleOutputWithContext(context.Background())
}

func (i *Rule) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleOutput)
}

type RuleOutput struct{ *pulumi.OutputState }

func (RuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Rule)(nil))
}

func (o RuleOutput) ToRuleOutput() RuleOutput {
	return o
}

func (o RuleOutput) ToRuleOutputWithContext(ctx context.Context) RuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleInput)(nil)).Elem(), &Rule{})
	pulumi.RegisterOutputType(RuleOutput{})
}
