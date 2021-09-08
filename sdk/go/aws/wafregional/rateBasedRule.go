// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html
type RateBasedRule struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-matchpredicates
	MatchPredicates RateBasedRulePredicateArrayOutput `pulumi:"matchPredicates"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-metricname
	MetricName pulumi.StringOutput `pulumi:"metricName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-name
	Name pulumi.StringOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratekey
	RateKey pulumi.StringOutput `pulumi:"rateKey"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratelimit
	RateLimit pulumi.IntOutput `pulumi:"rateLimit"`
}

// NewRateBasedRule registers a new resource with the given unique name, arguments, and options.
func NewRateBasedRule(ctx *pulumi.Context,
	name string, args *RateBasedRuleArgs, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricName == nil {
		return nil, errors.New("invalid value for required argument 'MetricName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.RateKey == nil {
		return nil, errors.New("invalid value for required argument 'RateKey'")
	}
	if args.RateLimit == nil {
		return nil, errors.New("invalid value for required argument 'RateLimit'")
	}
	var resource RateBasedRule
	err := ctx.RegisterResource("aws-native:wafregional:RateBasedRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRateBasedRule gets an existing RateBasedRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRateBasedRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RateBasedRuleState, opts ...pulumi.ResourceOption) (*RateBasedRule, error) {
	var resource RateBasedRule
	err := ctx.ReadResource("aws-native:wafregional:RateBasedRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RateBasedRule resources.
type rateBasedRuleState struct {
}

type RateBasedRuleState struct {
}

func (RateBasedRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleState)(nil)).Elem()
}

type rateBasedRuleArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-matchpredicates
	MatchPredicates []RateBasedRulePredicate `pulumi:"matchPredicates"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-metricname
	MetricName string `pulumi:"metricName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-name
	Name string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratekey
	RateKey string `pulumi:"rateKey"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratelimit
	RateLimit int `pulumi:"rateLimit"`
}

// The set of arguments for constructing a RateBasedRule resource.
type RateBasedRuleArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-matchpredicates
	MatchPredicates RateBasedRulePredicateArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-metricname
	MetricName pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-name
	Name pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratekey
	RateKey pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-ratebasedrule.html#cfn-wafregional-ratebasedrule-ratelimit
	RateLimit pulumi.IntInput
}

func (RateBasedRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rateBasedRuleArgs)(nil)).Elem()
}

type RateBasedRuleInput interface {
	pulumi.Input

	ToRateBasedRuleOutput() RateBasedRuleOutput
	ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput
}

func (*RateBasedRule) ElementType() reflect.Type {
	return reflect.TypeOf((*RateBasedRule)(nil))
}

func (i *RateBasedRule) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return i.ToRateBasedRuleOutputWithContext(context.Background())
}

func (i *RateBasedRule) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RateBasedRuleOutput)
}

type RateBasedRuleOutput struct{ *pulumi.OutputState }

func (RateBasedRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RateBasedRule)(nil))
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutput() RateBasedRuleOutput {
	return o
}

func (o RateBasedRuleOutput) ToRateBasedRuleOutputWithContext(ctx context.Context) RateBasedRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RateBasedRuleOutput{})
}
