// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package configuration

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html
type ConfigurationAggregator struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-accountaggregationsources
	AccountAggregationSources  ConfigurationAggregatorAccountAggregationSourceArrayOutput `pulumi:"accountAggregationSources"`
	ConfigurationAggregatorArn pulumi.StringOutput                                        `pulumi:"configurationAggregatorArn"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-configurationaggregatorname
	ConfigurationAggregatorName pulumi.StringPtrOutput `pulumi:"configurationAggregatorName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-organizationaggregationsource
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrOutput `pulumi:"organizationAggregationSource"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewConfigurationAggregator registers a new resource with the given unique name, arguments, and options.
func NewConfigurationAggregator(ctx *pulumi.Context,
	name string, args *ConfigurationAggregatorArgs, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	if args == nil {
		args = &ConfigurationAggregatorArgs{}
	}

	var resource ConfigurationAggregator
	err := ctx.RegisterResource("aws-native:Configuration:ConfigurationAggregator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationAggregator gets an existing ConfigurationAggregator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationAggregator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationAggregatorState, opts ...pulumi.ResourceOption) (*ConfigurationAggregator, error) {
	var resource ConfigurationAggregator
	err := ctx.ReadResource("aws-native:Configuration:ConfigurationAggregator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationAggregator resources.
type configurationAggregatorState struct {
}

type ConfigurationAggregatorState struct {
}

func (ConfigurationAggregatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorState)(nil)).Elem()
}

type configurationAggregatorArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-accountaggregationsources
	AccountAggregationSources []ConfigurationAggregatorAccountAggregationSource `pulumi:"accountAggregationSources"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-configurationaggregatorname
	ConfigurationAggregatorName *string `pulumi:"configurationAggregatorName"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-organizationaggregationsource
	OrganizationAggregationSource *ConfigurationAggregatorOrganizationAggregationSource `pulumi:"organizationAggregationSource"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-tags
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationAggregator resource.
type ConfigurationAggregatorArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-accountaggregationsources
	AccountAggregationSources ConfigurationAggregatorAccountAggregationSourceArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-configurationaggregatorname
	ConfigurationAggregatorName pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-organizationaggregationsource
	OrganizationAggregationSource ConfigurationAggregatorOrganizationAggregationSourcePtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configurationaggregator.html#cfn-config-configurationaggregator-tags
	Tags aws.TagArrayInput
}

func (ConfigurationAggregatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationAggregatorArgs)(nil)).Elem()
}

type ConfigurationAggregatorInput interface {
	pulumi.Input

	ToConfigurationAggregatorOutput() ConfigurationAggregatorOutput
	ToConfigurationAggregatorOutputWithContext(ctx context.Context) ConfigurationAggregatorOutput
}

func (*ConfigurationAggregator) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationAggregator)(nil))
}

func (i *ConfigurationAggregator) ToConfigurationAggregatorOutput() ConfigurationAggregatorOutput {
	return i.ToConfigurationAggregatorOutputWithContext(context.Background())
}

func (i *ConfigurationAggregator) ToConfigurationAggregatorOutputWithContext(ctx context.Context) ConfigurationAggregatorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationAggregatorOutput)
}

type ConfigurationAggregatorOutput struct{ *pulumi.OutputState }

func (ConfigurationAggregatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationAggregator)(nil))
}

func (o ConfigurationAggregatorOutput) ToConfigurationAggregatorOutput() ConfigurationAggregatorOutput {
	return o
}

func (o ConfigurationAggregatorOutput) ToConfigurationAggregatorOutputWithContext(ctx context.Context) ConfigurationAggregatorOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigurationAggregatorOutput{})
}
