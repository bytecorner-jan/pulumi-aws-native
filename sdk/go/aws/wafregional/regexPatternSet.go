// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package wafregional

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html
type RegexPatternSet struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-name
	Name pulumi.StringOutput `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-regexpatternstrings
	RegexPatternStrings pulumi.StringArrayOutput `pulumi:"regexPatternStrings"`
}

// NewRegexPatternSet registers a new resource with the given unique name, arguments, and options.
func NewRegexPatternSet(ctx *pulumi.Context,
	name string, args *RegexPatternSetArgs, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.RegexPatternStrings == nil {
		return nil, errors.New("invalid value for required argument 'RegexPatternStrings'")
	}
	var resource RegexPatternSet
	err := ctx.RegisterResource("aws-native:wafregional:RegexPatternSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegexPatternSet gets an existing RegexPatternSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegexPatternSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegexPatternSetState, opts ...pulumi.ResourceOption) (*RegexPatternSet, error) {
	var resource RegexPatternSet
	err := ctx.ReadResource("aws-native:wafregional:RegexPatternSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegexPatternSet resources.
type regexPatternSetState struct {
}

type RegexPatternSetState struct {
}

func (RegexPatternSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetState)(nil)).Elem()
}

type regexPatternSetArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-name
	Name string `pulumi:"name"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-regexpatternstrings
	RegexPatternStrings []string `pulumi:"regexPatternStrings"`
}

// The set of arguments for constructing a RegexPatternSet resource.
type RegexPatternSetArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-name
	Name pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-regexpatternset.html#cfn-wafregional-regexpatternset-regexpatternstrings
	RegexPatternStrings pulumi.StringArrayInput
}

func (RegexPatternSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regexPatternSetArgs)(nil)).Elem()
}

type RegexPatternSetInput interface {
	pulumi.Input

	ToRegexPatternSetOutput() RegexPatternSetOutput
	ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput
}

func (*RegexPatternSet) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexPatternSet)(nil))
}

func (i *RegexPatternSet) ToRegexPatternSetOutput() RegexPatternSetOutput {
	return i.ToRegexPatternSetOutputWithContext(context.Background())
}

func (i *RegexPatternSet) ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegexPatternSetOutput)
}

type RegexPatternSetOutput struct{ *pulumi.OutputState }

func (RegexPatternSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegexPatternSet)(nil))
}

func (o RegexPatternSetOutput) ToRegexPatternSetOutput() RegexPatternSetOutput {
	return o
}

func (o RegexPatternSetOutput) ToRegexPatternSetOutputWithContext(ctx context.Context) RegexPatternSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RegexPatternSetOutput{})
}
